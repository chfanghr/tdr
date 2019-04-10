package authblob

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"github.com/chfanghr/tdr/spotify/crypto"
	. "github.com/chfanghr/tdr/spotify/utils"
	"io/ioutil"
)

type AuthBlob struct {
	Username    string `json:"username"`
	DecodedBlob string `json:"decoded_blob"`
}

func FromFile(path string) (*AuthBlob, error) {
	if data, err := UnwrapResult(func() {
		buf, err := ioutil.ReadFile(path)
		ThrowIfError(err)
		tmp := &AuthBlob{}
		err = json.Unmarshal(buf, tmp)
		ThrowIfError(err)
	}); err != nil {
		return nil, err
	} else {
		return data.(*AuthBlob), nil
	}
}

func NewAuthBlob(blob64 string, client64 string, keys crypto.PrivateKeys, deviceId string, username string) (*AuthBlob, error) {
	if data, err := UnwrapResult(func() {
		partDecoded, err := decodeBlob(blob64, client64, keys)
		ThrowIfError(err)
		fullDecoded := decodeBlobSecondary(partDecoded, username,
			deviceId)
		ThrowData(&AuthBlob{
			Username:    username,
			DecodedBlob: base64.StdEncoding.EncodeToString(fullDecoded),
		})
	}); err != nil {
		return nil, err
	} else {
		return data.(*AuthBlob), nil
	}
}

func ToFile(path string, ab *AuthBlob) error {
	_, err := UnwrapResult(func() {
		buf, err := json.Marshal(ab)
		ThrowIfError(err)
		err = ioutil.WriteFile(path, buf, 0600)
		ThrowIfError(err)
	})
	return err
}

func (b *AuthBlob) MakeSpotBlob(deviceId string, client64 string, dhKeys crypto.PrivateKeys) (string, error) {
	if data, err := UnwrapResult(func() {
		secret := sha1.Sum([]byte(deviceId))
		key := blobKey(b.Username, secret[:])
		blobBytes, err := base64.StdEncoding.DecodeString(b.DecodedBlob)
		ThrowIfError(err)
		encoded := encryptBlob(blobBytes, key)
		fullEncoded := makeSpotBlob(encoded, dhKeys, client64)
		ThrowData(fullEncoded)
	}); err != nil {
		return "", err
	} else {
		return data.(string), nil
	}
}
