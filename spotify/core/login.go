package core

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/chfanghr/tdr/spotify/authblob"
	"github.com/chfanghr/tdr/spotify/connection"
	spot "github.com/chfanghr/tdr/spotify/proto"
	. "github.com/chfanghr/tdr/spotify/utils"
	"github.com/golang/protobuf/proto"
	"log"
)

var Version string
var BuildID string

func Login(username string, password string, deviceName string) (*Session, error) {
	if data, err := UnwrapResultFromJob(func() {
		s, err := setupSession()
		ThrowIfError(err)
		ThrowResult(Result{Data: s, Err: s.loginSession(username, password, deviceName)})
	}); err != nil {
		return nil, err
	} else {
		return data.(*Session), nil
	}
}

func (s *Session) loginSession(username string, password string, deviceName string) error {
	return ResultFromJob(func() {
		s.deviceId = GenerateDeviceId(deviceName)
		s.deviceName = deviceName
		err := s.startConnection()
		ThrowIfError(err)
		loginPacket, err := makeLoginPasswordPacket(username, password, s.deviceId)
		ThrowIfError(err)
		ThrowIfError(s.doLogin(loginPacket, username))
	}).Err
}

// Login to Spotify using an existing authBlob
func LoginSaved(username string, authData []byte, deviceName string) (*Session, error) {
	if data, err := UnwrapResultFromJob(func() {
		s, err := setupSession()
		ThrowIfError(err)
		s.deviceId = GenerateDeviceId(deviceName)
		s.deviceName = deviceName
		err = s.startConnection()
		ThrowIfError(err)
		packet, err := makeLoginBlobPacket(username, authData,
			spot.AuthenticationType_AUTHENTICATION_STORED_SPOTIFY_CREDENTIALS.Enum(), s.deviceId)
		ThrowIfError(err)
		ThrowResult(Result{Data: s, Err: s.doLogin(packet, username)})
	}); err != nil {
		return nil, err
	} else {
		return data.(*Session), nil
	}
}

//TODO
// Registers as a Spotify Connect device via mdns. When user connects, logs on to Spotify and saves
// credentials in file at cacheBlobPath. Once saved, the blob credentials allow the program to connect to other
// Spotify Connect devices and control them.
//func LoginDiscovery(cacheBlobPath string, deviceName string) (*Session, error) {
//	deviceId := GenerateDeviceId(deviceName)
//	disc := discovery.LoginFromConnect(cacheBlobPath, deviceId, deviceName)
//	return sessionFromDiscovery(disc)
//}

//TODO
// Login using an authentication blob through Spotify Connect discovery system, reading an existing blob data. To read
// from a file, see LoginDiscoveryBlobFile.
//func LoginDiscoveryBlob(username string, blob string, deviceName string) (*Session, error) {
//	deviceId := GenerateDeviceId(deviceName)
//	disc := discovery.CreateFromBlob(BlobInfo{
//		Username:    username,
//		DecodedBlob: blob,
//	}, "", deviceId, deviceName)
//	return sessionFromDiscovery(disc)
//}

//TODO
// Login from credentials at cacheBlobPath previously saved by LoginDiscovery. Similar to LoginDiscoveryBlob, except
// it reads it directly from a file.
//func LoginDiscoveryBlobFile(cacheBlobPath, deviceName string) (*Session, error) {
//	deviceId := GenerateDeviceId(deviceName)
//	disc := discovery.CreateFromFile(cacheBlobPath, deviceId, deviceName)
//	return sessionFromDiscovery(disc)
//}

// Login to Spotify using the OAuth method
func LoginOAuth(deviceName string, clientId string, clientSecret string) (*Session, error) {
	token := getOAuthToken(clientId, clientSecret)
	return loginOAuthToken(token.AccessToken, deviceName)
}

func loginOAuthToken(accessToken string, deviceName string) (*Session, error) {
	if data, err := UnwrapResultFromJob(func() {
		s, err := setupSession()
		ThrowIfError(err)
		s.deviceId = GenerateDeviceId(deviceName)
		s.deviceName = deviceName

		err = s.startConnection()
		ThrowIfError(err)

		packet, err := makeLoginBlobPacket("", []byte(accessToken),
			spot.AuthenticationType_AUTHENTICATION_SPOTIFY_TOKEN.Enum(), s.deviceId)
		ThrowResult(Result{Data: s, Err: s.doLogin(packet, "")})
	}); err != nil {
		return nil, err
	} else {
		return data.(*Session), nil
	}
}

func (s *Session) doLogin(packet []byte, username string) error {
	err := s.stream.SendPacket(connection.PacketLogin, packet)
	if err != nil {
		log.Fatal("bad shannon write", err)
	}

	// Pll once for authentication response
	welcome, err := s.handleLogin()
	if err != nil {
		return err
	}

	// Store the few interesting values
	s.username = welcome.GetCanonicalUsername()
	if s.username == "" {
		// Spotify might not return a canonical username, so reuse the blob's one instead
		//s.username = s.discovery.LoginBlob().Username
	}
	s.reusableAuthBlob = welcome.GetReusableAuthCredentials()

	// Poll for acknowledge before loading - needed for gopherjs
	// s.poll()
	go s.runPollLoop()

	return nil
}

func (s *Session) handleLogin() (*spot.APWelcome, error) {
	if data, err := UnwrapResultFromJob(func() {
		cmd, data, err := s.stream.RecvPacket()
		WrapAndThrowIfError("authentication failed", err)
		if cmd == connection.PacketAuthFailure {
			ThrowError(fmt.Errorf("authentication failed"))
		} else if cmd == connection.PacketAPWelcome {
			welcome := &spot.APWelcome{}
			WrapAndThrowIfError("authentication failed", proto.Unmarshal(data, welcome))
			//log.Println("Authentication succeeded: Welcome,", welcome.GetCanonicalUsername())
			//log.Println("Blob type:", welcome.GetReusableAuthCredentialsType())
			ThrowData(welcome)
		} else {
			ThrowError(fmt.Errorf("authentication failed: unexpected cmd %v", cmd))
		}
	}); err != nil {
		return nil, err
	} else {
		return data.(*spot.APWelcome), nil
	}
}

func (s *Session) getLoginBlobPacket(blob *authblob.AuthBlob) ([]byte, error) {
	data, _ := base64.StdEncoding.DecodeString(blob.DecodedBlob)
	buffer := bytes.NewBuffer(data)
	_, _ = buffer.ReadByte()
	readBytes(buffer)
	_, _ = buffer.ReadByte()
	authNum := readInt(buffer)
	authType := spot.AuthenticationType(authNum)
	_, _ = buffer.ReadByte()
	authData := readBytes(buffer)
	return makeLoginBlobPacket(blob.Username, authData, &authType, s.deviceId)
}

func makeLoginPasswordPacket(username string, password string, deviceId string) ([]byte, error) {
	return makeLoginBlobPacket(username, []byte(password), spot.AuthenticationType_AUTHENTICATION_USER_PASS.Enum(), deviceId)
}

func makeLoginBlobPacket(username string, authData []byte, authType *spot.AuthenticationType, deviceId string) ([]byte, error) {
	if data, err := UnwrapResultFromJob(func() {
		versionString := fmt.Sprintf("tdr_core_%s_%s", Version, BuildID)
		packet := &spot.ClientResponseEncrypted{
			LoginCredentials: &spot.LoginCredentials{
				Username: proto.String(username),
				Typ:      authType,
				AuthData: authData,
			},
			SystemInfo: &spot.SystemInfo{
				CpuFamily:               spot.CpuFamily_CPU_UNKNOWN.Enum(),
				Os:                      spot.Os_OS_UNKNOWN.Enum(),
				SystemInformationString: proto.String("tdr_core"),
				DeviceId:                proto.String(deviceId),
			},
			VersionString: proto.String(versionString),
		}
		packetData, err := proto.Marshal(packet)
		ThrowResult(Result{Err: err, Data: packetData})
	}); err != nil {
		return nil, err
	} else {
		return data.([]byte), nil
	}
}
