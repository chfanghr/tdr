package mercury

import (
	"encoding/json"
	"fmt"
	"github.com/chfanghr/tdr/spotify/metadata"
	spot "github.com/chfanghr/tdr/spotify/proto"
	. "github.com/chfanghr/tdr/spotify/utils"
	"github.com/chfanghr/tdr/spotify/version"
	"github.com/golang/protobuf/proto"
	"io"
	"io/ioutil"
	"net/url"
)

type Client struct {
	subscriptions map[string][]chan Response
	callbacks     map[string]Callback
	internal      *internal
}

func (m *Client) GetRootPlaylist(username string) (*spot.SelectedListContent, error) {
	uri := fmt.Sprintf("hm://playlist/user/%s/rootlist", username)

	result := &spot.SelectedListContent{}
	err := m.mercuryGetProto(uri, result)
	return result, err
}

func (m *Client) GetPlaylist(id string) (*spot.SelectedListContent, error) {
	uri := fmt.Sprintf("hm://playlist/%s", id)

	result := &spot.SelectedListContent{}
	err := m.mercuryGetProto(uri, result)
	return result, err
}

func (m *Client) GetToken(clientId string, scopes string) (*metadata.Token, error) {
	uri := fmt.Sprintf("hm://keymaster/token/authenticated?client_id=%s&scope=%s", url.QueryEscape(clientId),
		url.QueryEscape(scopes))

	token := &metadata.Token{}
	err := m.mercuryGetJson(uri, token)
	return token, err
}

func (m *Client) Search(search string, limit int, country string, username string) (*metadata.SearchResponse, error) {
	v := url.Values{}
	v.Set("entityVersion", "2")
	v.Set("limit", fmt.Sprintf("%d", limit))
	v.Set("imageSize", "large")
	v.Set("catalogue", "")
	v.Set("country", country)
	v.Set("platform", "zelda")
	v.Set("username", username)

	uri := fmt.Sprintf("hm://searchview/km/v4/search/%s?%s", url.QueryEscape(search), v.Encode())

	result := &metadata.SearchResponse{}
	err := m.mercuryGetJson(uri, result)
	return result, err
}

func (m *Client) Suggest(search string) (*metadata.SuggestResult, error) {
	uri := "hm://searchview/km/v3/suggest/" + url.QueryEscape(search) + "?limit=3&intent=2516516747764520149&sequence=0&catalogue=&country=&locale=&platform=zelda&username="
	data := m.mercuryGet(uri)

	return parseSuggest(data)
}

func (m *Client) GetTrack(id string) (*spot.Track, error) {
	uri := "hm://metadata/4/track/" + id
	result := &spot.Track{}
	err := m.mercuryGetProto(uri, result)
	return result, err
}

func (m *Client) GetArtist(id string) (*spot.Artist, error) {
	uri := "hm://metadata/4/artist/" + id
	result := &spot.Artist{}
	err := m.mercuryGetProto(uri, result)
	return result, err
}

func (m *Client) GetAlbum(id string) (*spot.Album, error) {
	uri := "hm://metadata/4/album/" + id
	result := &spot.Album{}
	err := m.mercuryGetProto(uri, result)
	return result, err
}

func (m *Client) NextSeq() []byte {
	_, seq := m.internal.nextSeq()
	return seq
}

func (m *Client) NextSeqWithInt() (uint32, []byte) {return m.internal.nextSeq()}

// Subscribe subscribes the specified receiving channel to the specified URI, and calls the callback function
// whenever there's an event happening.
func (m *Client) Subscribe(uri string, recv chan Response, cb Callback) error {
	m.addChannelSubscriber(uri, recv)
	return m.Request(Request{
		Method: "SUB",
		Uri:    uri,
	}, func(response Response) {
		for _, part := range response.Payload {
			sub := &spot.Subscription{}
			err := proto.Unmarshal(part, sub)
			if err == nil && *sub.Uri != uri {
				m.addChannelSubscriber(*sub.Uri, recv)
			}
		}
		cb(response)
	})
}

func (m *Client) Request(req Request, cb Callback) (err error) {
	seq, err := m.internal.request(req)
	if err != nil {
		// Call the callback with a 500 error-code so that the request doesn't remain pending in case of error
		if cb != nil {
			cb(Response{StatusCode: 500,})
		}
		return err
	}

	m.callbacks[string(seq)] = cb
	return nil
}

func (m *Client) Handle(cmd uint8, reader io.Reader) error {
	return ResultFromJob(func() {
		response, err := m.internal.parseResponse(cmd, reader)
		ThrowIfError(err)
		if response != nil {
			if cmd == 0xb5 {
				chList, ok := m.subscriptions[response.Uri]
				if ok {
					for _, ch := range chList {
						ch <- *response
					}
				}
			} else if cb, ok := m.callbacks[response.SeqKey]; ok {
				delete(m.callbacks, response.SeqKey)
				cb(*response)
			}
		}
	}).Err
}

func (m *Client) mercuryGet(url string) []byte {
	done := make(chan []byte)
	go func() {
		IgnoreError(nil, m.Request(Request{
			Method:  "GET",
			Uri:     url,
			Payload: [][]byte{},
		}, func(res Response) {
			done <- res.CombinePayload()
		}))
	}()

	result := <-done
	return result
}

func (m *Client) mercuryGetJson(url string, result interface{}) error {
	data := m.mercuryGet(url)
	version.Debug.Printf("%s", data)
	return json.Unmarshal(data, result)
}

func (m *Client) mercuryGetProto(url string, result proto.Message) error {
	data := m.mercuryGet(url)
	if version.IsDebugBuild() {
		IgnoreError(nil,ioutil.WriteFile("/tmp/proto.blob", data, 0644))
	}
	return proto.Unmarshal(data, result)
}

func (m *Client) addChannelSubscriber(uri string, recv chan Response) {
	chList, ok := m.subscriptions[uri]
	if !ok {
		chList = make([]chan Response, 0)
	}
	chList = append(chList, recv)
	m.subscriptions[uri] = chList
}
