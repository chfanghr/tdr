package authblob

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"github.com/chfanghr/tdr/spotify/crypto"
	"golang.org/x/crypto/pbkdf2"

	. "github.com/chfanghr/tdr/spotify/utils"
	"math/big"
)

func decodeBlob(blob64 string, client64 string, keys crypto.PrivateKeys) (string, error) {
	if data, err := UnwrapResultFromJob(func() {
		clientKey, err := base64.StdEncoding.DecodeString(client64)
		ThrowIfError(err)
		blobBytes, err := base64.StdEncoding.DecodeString(blob64)
		ThrowIfError(err)

		clientkeyBe := new(big.Int)
		clientkeyBe.SetBytes(clientKey)
		sharedKey := crypto.Powm(clientkeyBe, keys.PrivateKey(), keys.Prime())
		iv := blobBytes[0:16]
		encryptedPart := blobBytes[16 : len(blobBytes)-20]
		ckSum := blobBytes[len(blobBytes)-20:]
		key := sha1.Sum(sharedKey.Bytes())
		baseKey := key[:16]
		hash := hmac.New(sha1.New, baseKey)
		hash.Write([]byte("checksum"))
		checksumKey := hash.Sum(nil)
		hash.Reset()
		hash.Write([]byte("encryption"))
		encryptionKey := hash.Sum(nil)
		hash.Reset()

		macHash := hmac.New(sha1.New, checksumKey)
		macHash.Write(encryptedPart)
		mac := macHash.Sum(nil)

		if !bytes.Equal(mac, ckSum) {
			ThrowIfError(errors.New("mac mismatch"))
		}

		block, _ := aes.NewCipher(encryptionKey[0:16])
		stream := cipher.NewCTR(block, iv)
		stream.XORKeyStream(encryptedPart, encryptedPart)
		ThrowData(string(encryptedPart))
	}); err != nil {
		return "", err
	} else {
		return data.(string), nil
	}
}

func decodeBlobSecondary(blob64 string, username string, deviceId string) []byte {
	blob, _ := base64.StdEncoding.DecodeString(blob64)
	secret := sha1.Sum([]byte(deviceId))
	key := blobKey(username, secret[:])

	data := decryptBlob(blob, key)
	return data
}

func decryptBlob(blob []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	bs := block.BlockSize()
	if len(blob)%bs != 0 {
		panic("Need a multiple of the blksize")
	}

	plaintext := make([]byte, len(blob))

	plain := plaintext
	for len(blob) > 0 {
		block.Decrypt(plaintext, blob)
		plaintext = plaintext[bs:]
		blob = blob[bs:]
	}

	l := len(plain)
	for i := 0; i < l-0x10; i++ {
		plain[l-i-1] = plain[l-i-1] ^ plain[l-i-0x11]
	}

	return plain
}

func blobKey(username string, secret []byte) []byte {
	data := pbkdf2.Key(secret, []byte(username), 0x100, 20, sha1.New)[0:20]

	hash := sha1.Sum(data)
	length := make([]byte, 4)
	binary.BigEndian.PutUint32(length, 20)
	return append(hash[:], length...)
}

func encryptBlob(blob []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	bs := block.BlockSize()
	if len(blob)%bs != 0 {
		panic("Need a multiple of the blocksize")
	}

	l := len(blob)
	for i := l - 0x11; i >= 0; i-- {
		blob[l-i-1] = blob[l-i-1] ^ blob[l-i-0x11]
	}

	ciphertext := make([]byte, len(blob))
	encoded := ciphertext
	for len(blob) > 0 {
		block.Encrypt(ciphertext, blob)
		ciphertext = ciphertext[bs:]
		blob = blob[bs:]
	}

	return encoded
}

func makeSpotBlob(blobPart []byte, keys crypto.PrivateKeys, publicKey string) string {
	part := []byte(base64.StdEncoding.EncodeToString(blobPart))

	sharedKey := keys.SharedKey(publicKey)
	iv := crypto.RandomVec(16)

	key := sha1.Sum(sharedKey)
	baseKey := key[:16]
	hash := hmac.New(sha1.New, baseKey)

	hash.Write([]byte("checksum"))
	checksumKey := hash.Sum(nil)
	hash.Reset()

	hash.Write([]byte("encryption"))
	encryptionKey := hash.Sum(nil)
	hash.Reset()

	block, _ := aes.NewCipher(encryptionKey[0:16])
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(part, part)

	macHash := hmac.New(sha1.New, checksumKey)
	macHash.Write(part)
	mac := macHash.Sum(nil)

	part = append(iv, part...)
	part = append(part, mac...)

	return base64.StdEncoding.EncodeToString(part)
}
