// Mercury is the protocol implementation for Spotify Connect playback control and metadata fetching.It works as a
// PUB/SUB system, where you, as an audio sink, subscribes to the events of a specified user (playlist changes) but
// also access various metadata normally fetched by external players (tracks metadata, playlists, artists, etc).

package mercury

import (
	"github.com/chfanghr/tdr/spotify/connection"
)

// CreateMercury initializes a Connection for the specified session.
func CreateMercury(stream connection.PacketStream) *Client {
	client := &Client{
		callbacks:     make(map[string]Callback),
		subscriptions: make(map[string][]chan Response),
		internal: &internal{
			pending: make(map[string]Pending),
			stream:  stream,
		},
	}
	return client
}