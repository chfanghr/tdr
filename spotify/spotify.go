package spotify

import (
	"github.com/chfanghr/tdr/spotify/core"
	"github.com/chfanghr/tdr/spotify/version"
)

func Version() string { return version.Version }

func BuildID() string { return version.BuildID }

//TODO
func Login(username, password string) (*core.Session, error) { panic(nil) }

//TODO
func LoginAuthBlob(buf []byte) (*core.Session, error) { panic(nil) }
