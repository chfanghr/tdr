package core

import (
	"bytes"
	"fmt"
	"github.com/chfanghr/tdr/spotify/connection"
	"github.com/chfanghr/tdr/spotify/crypto"
	"github.com/chfanghr/tdr/spotify/discovery"
	"github.com/chfanghr/tdr/spotify/mercury"
	"github.com/chfanghr/tdr/spotify/player"
	spot "github.com/chfanghr/tdr/spotify/proto"
	. "github.com/chfanghr/tdr/spotify/utils"
	"github.com/chfanghr/tdr/spotify/version"
	"github.com/golang/protobuf/proto"
	"io"
	"net"
	"time"
)

// Session represents an active Spotify connection
type Session struct {
	/// Constructor references

	// mercuryConstructor is the constructor that should be used to build a mercury connection
	mercuryConstructor func(conn connection.PacketStream) *mercury.Client

	// shannonConstructor is the constructor used to build the shannon-encrypted PacketStream connection
	shannonConstructor func(keys crypto.SharedKeys, conn connection.PlainConnection) connection.PacketStream

	/// Managers and helpers
	// stream is the encrypted connection to the Spotify server
	stream connection.PacketStream

	// mercury is the mercury client associated with this session
	mercury *mercury.Client

	// discovery is the discovery service used for Spotify Connect devices discovery
	discovery *discovery.Discovery

	// player is the player service used to load the audio data
	player *player.Player

	// tcpCon is the plain I/O network connection to the server
	tcpCon io.ReadWriter

	// keys are the encryption keys used to communicate with the server
	keys crypto.PrivateKeys

	/// State and variables

	// deviceId is the device identifier (computer name, Android serial number, ...) sent during auth to the Spotify
	// servers for this session
	deviceId string

	// deviceName is the device name (Android device model) sent during auth to the Spotify servers for this session
	deviceName string

	// username is the currently authenticated canonical username
	username string

	// reusableAuthBlob is the reusable authentication blob for Spotify Connect devices
	reusableAuthBlob []byte

	// country is the user country returned by the Spotify servers
	country string
}

func (s *Session) Stream() connection.PacketStream {
	return s.stream
}

func (s *Session) Discovery() *discovery.Discovery {
	return s.discovery
}

func (s *Session) Mercury() *mercury.Client {
	return s.mercury
}

func (s *Session) Player() *player.Player {
	return s.player
}

func (s *Session) Username() string {
	return s.username
}

func (s *Session) DeviceId() string {
	return s.deviceId
}

func (s *Session) ReusableAuthBlob() []byte {
	return s.reusableAuthBlob
}

func (s *Session) Country() string {
	return s.country
}

func (s *Session) startConnection() error {
	return ResultFromJob(func() {
		// First, start by performing a plaintext connection and send the Hello message
		conn := connection.MakePlainConnection(s.tcpCon, s.tcpCon)
		helloMessage := makeHelloMessage(s.keys.PubKey(), s.keys.ClientNonce())
		initClientPacket, err := conn.SendPrefixPacket([]byte{0, 4}, helloMessage)
		WrapAndThrowIfError("failed to write client hello", err)

		// Wait and read the hello reply
		initServerPacket, err := conn.RecvPacket()
		WrapAndThrowIfError("failed to receive packet for hello: ", err)

		response := spot.APResponseMessage{}
		err = proto.Unmarshal(initServerPacket[4:], &response)
		WrapAndThrowIfError("failed to unmarshal server hello", err)

		remoteKey := response.Challenge.LoginCryptoChallenge.DiffieHellman.Gs
		sharedKeys := s.keys.AddRemoteKey(remoteKey, initClientPacket, initServerPacket)

		plainResponse := &spot.ClientResponsePlaintext{
			LoginCryptoResponse: &spot.LoginCryptoResponseUnion{
				DiffieHellman: &spot.LoginCryptoDiffieHellmanResponse{
					Hmac: sharedKeys.Challenge(),
				},
			},
			PowResponse:    &spot.PoWResponseUnion{},
			CryptoResponse: &spot.CryptoResponseUnion{},
		}

		plainResponseMessage, err := proto.Marshal(plainResponse)
		ThrowIfError(err)

		_, err = conn.SendPrefixPacket([]byte{}, plainResponseMessage)
		WrapAndThrowIfError("failed to write client plain response", err)

		s.stream = s.shannonConstructor(sharedKeys, conn)
		s.mercury = s.mercuryConstructor(s.stream)

		s.player = player.CreatePlayer(s.stream, s.mercury)
	}).Err
}

func setupSession() (*Session, error) {
	session := &Session{
		keys:               crypto.GenerateKeys(),
		mercuryConstructor: mercury.CreateMercury,
		shannonConstructor: crypto.CreateStream,
	}
	return session, session.doConnect()
}

//TODO
//func sessionFromDiscovery(d *discovery.Discovery) (*Session, error) {
//	s, err := setupSession()
//	if err != nil {
//		return nil, err
//	}
//
//	s.discovery = d
//	s.deviceId = d.DeviceId()
//	s.deviceName = d.DeviceName()
//
//	err = s.startConnection()
//	if err != nil {
//		return s, err
//	}
//
//	loginPacket := s.getLoginBlobPacket(d.LoginBlob())
//	return s, s.doLogin(loginPacket, d.LoginBlob().Username)
//}

func (s *Session) doConnect() error {
	return ResultFromJob(func() {
		apUrl, err := GetAP()
		WrapAndThrowIfError("failed to get ap url", err)
		s.tcpCon, err = net.Dial("tcp", apUrl)
		WrapAndThrowIfError("failed to connect", err)
	}).Err
}

func (s *Session) disconnect() error {
	return ResultFromJob(func() {
		if s.tcpCon != nil {
			conn := s.tcpCon.(net.Conn)
			WrapAndThrowIfError("failed to close tcp connection", conn.Close())
			s.tcpCon = nil
		}
	}).Err
}

func (s *Session) doReconnect() error {
	return ResultFromJob(func() {
		ThrowIfError(s.disconnect())
		ThrowIfError(s.doConnect())
		ThrowIfError(s.startConnection())
		packet, err := makeLoginBlobPacket(s.username, s.reusableAuthBlob,
			spot.AuthenticationType_AUTHENTICATION_STORED_SPOTIFY_CREDENTIALS.Enum(), s.deviceId)
		ThrowIfError(err)
		ThrowIfError(s.doLogin(packet, s.username))
	}).Err
}

func (s *Session) planReconnect() {
	go func() {
		time.Sleep(1 * time.Second)

		if err := s.doReconnect(); err != nil {
			// Try to reconnect again in a second
			s.planReconnect()
		}
	}()
}

func (s *Session) runPollLoop() {
	for {
		cmd, data, err := s.stream.RecvPacket()
		if err != nil {
			//log.Println("Error during RecvPacket: ", err)
			if err == io.EOF {
				// We've been disconnected, reconnect
				s.planReconnect()
				break
			}
		} else {
			s.handle(cmd, data)
		}
	}
}

func (s *Session) handle(cmd uint8, data []byte) {
	version.Debug.Printf("handle, cmd=0x%x data=%x\n", cmd, data)
	switch {
	case cmd == connection.PacketPing:
		// Ping
		CrashProgramIfError("failed to handle PacketPing", s.stream.SendPacket(connection.PacketPong, data))

	case cmd == connection.PacketPongAck:
		// Pong reply, ignore

	case cmd == connection.PacketAesKey || cmd == connection.PacketAesKeyError ||
		cmd == connection.PacketStreamChunkRes:
		// Audio key and data responses
		IgnoreError(nil, s.player.HandleCmd(cmd, data))

	case cmd == connection.PacketCountryCode:
		// Handle country code
		s.country = fmt.Sprintf("%s", data)

	case 0xb2 <= cmd && cmd <= 0xb6:
		// Mercury responses
		WrapAndThrowIfError("handle 0xbx", s.mercury.Handle(cmd, bytes.NewReader(data))) //if error occurs,just panic

	case cmd == connection.PacketSecretBlock:
		// Old RSA public key

	case cmd == connection.PacketLegacyWelcome:
		// Empty welcome packet

	case cmd == connection.PacketProductInfo:
		// Has some info about A/B testing status, product setup, etc... in an XML fashion.

	case cmd == 0x1f:
		// Unknown, data is zeroes only

	case cmd == connection.PacketLicenseVersion:
		// This is a simple blob containing the current Spotify license version (e.g. 1.0.1-FR). Format of the blob
		// is [ uint16 id (= 0x001), uint8 len, string license ]

	default:
		version.Debug.Printf("unhandled cmd 0x%x\n", cmd)
	}
}

func (s *Session) poll() {
	cmd, data, err := s.stream.RecvPacket()
	WrapAndThrowIfError("poll error", err)
	s.handle(cmd, data)
}

func readInt(b *bytes.Buffer) uint32 {
	c, _ := b.ReadByte()
	lo := uint32(c)
	if lo&0x80 == 0 {
		return lo
	}

	c2, _ := b.ReadByte()
	hi := uint32(c2)
	return lo&0x7f | hi<<7
}

func readBytes(b *bytes.Buffer) []byte {
	length := readInt(b)
	data := make([]byte, length)
	_, _ = b.Read(data)
	return data
}

func makeHelloMessage(publicKey []byte, nonce []byte) []byte {
	hello := &spot.ClientHello{
		BuildInfo: &spot.BuildInfo{
			Product:  spot.Product_PRODUCT_PARTNER.Enum(),
			Platform: spot.Platform_PLATFORM_LINUX_X86.Enum(),
			Version:  proto.Uint64(0x10800000000),
		},
		CryptosuitesSupported: []spot.Cryptosuite{
			spot.Cryptosuite_CRYPTO_SUITE_SHANNON},
		LoginCryptoHello: &spot.LoginCryptoHelloUnion{
			DiffieHellman: &spot.LoginCryptoDiffieHellmanHello{
				Gc:              publicKey,
				ServerKeysKnown: proto.Uint32(1),
			},
		},
		ClientNonce: nonce,
		FeatureSet: &spot.FeatureSet{
			Autoupdate2: proto.Bool(true),
		},
		Padding: []byte{0x1e},
	}
	packetData, err := proto.Marshal(hello)
	CrashProgramIfError("login marshaling failed", err)
	return packetData
}
