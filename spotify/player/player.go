package player

import (
	"github.com/chfanghr/tdr/spotify/connection"
	"github.com/chfanghr/tdr/spotify/mercury"
	. "github.com/chfanghr/tdr/spotify/utils"
)

type Player struct{}

func CreatePlayer(connection.PacketStream, *mercury.Client) *Player { ThrowEmptyResult(); return nil } //TODO

func (p *Player) HandleCmd(cmd uint8, data []byte) error { return nil } //TODO
