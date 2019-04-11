package mercury

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"github.com/chfanghr/tdr/spotify/metadata"
	spot "github.com/chfanghr/tdr/spotify/proto"
	. "github.com/chfanghr/tdr/spotify/utils"
	"github.com/golang/protobuf/proto"
	"io"
)

type Request struct {
	Method      string
	Uri         string
	ContentType string
	Payload     [][]byte
}

type Callback func(Response)

type Pending struct {
	parts   [][]byte
	partial []byte
}

type Connection interface {
	Subscribe(uri string, recv chan Response, cb Callback) error
	Request(req Request, cb Callback) (err error)
	Handle(cmd uint8, reader io.Reader) (err error)
}

type Response struct {
	HeaderData []byte
	Uri        string
	Payload    [][]byte
	StatusCode int32
	SeqKey     string
}

func (res *Response) CombinePayload() []byte {
	body := make([]byte, 0)
	for _, p := range res.Payload {
		body = append(body, p...)
	}
	return body
}

func encodeMercuryHead(seq []byte, partsLength uint16, flags uint8) (*bytes.Buffer, error) {
	if data, err := UnwrapResultFromJob(func() {
		buf := new(bytes.Buffer)
		ThrowIfError(binary.Write(buf, binary.BigEndian, uint16(len(seq))))
		_, err := buf.Write(seq)
		ThrowIfError(err)
		ThrowIfError(binary.Write(buf, binary.BigEndian, uint8(flags)))
		ThrowIfError(binary.Write(buf, binary.BigEndian, partsLength))
		ThrowData(buf)
	}); err != nil {
		return nil, err
	} else {
		return data.(*bytes.Buffer), nil
	}
}

func encodeRequest(seq []byte, req Request) ([]byte, error) {
	if data, err := UnwrapResultFromJob(func() {
		buf, err := encodeMercuryHead(seq, uint16(1+len(req.Payload)), uint8(1))
		ThrowIfError(err)
		header := &spot.Header{
			Uri:    proto.String(req.Uri),
			Method: proto.String(req.Method),
		}
		if req.ContentType != "" {
			header.ContentType = proto.String(req.ContentType)
		}
		headerData, err := proto.Marshal(header)
		ThrowIfError(binary.Write(buf, binary.BigEndian, uint16(len(headerData))))
		_, err = buf.Write(headerData)
		ThrowIfError(err)
		for _, p := range req.Payload {
			ThrowIfError(binary.Write(buf, binary.BigEndian, uint16(len(p))))
			_, err = buf.Write(p)
			ThrowIfError(err)
		}
		ThrowData(buf.Bytes())
	}); err != nil {
		return nil, err
	} else {
		return data.([]byte), nil
	}
}

func handleHead(reader io.Reader) ([]byte, uint8, uint16, error) {
	type res struct {
		seq   []byte
		flags uint8
		count uint16
	}
	if data, resErr := UnwrapResultFromJob(func() {
		var seq []byte
		var flags uint8
		var count uint16
		var seqLength uint16
		ThrowIfError(binary.Read(reader, binary.BigEndian, &seqLength))
		seq = make([]byte, seqLength)
		_, err := io.ReadFull(reader, seq)
		ThrowIfError(err)
		ThrowIfError(binary.Read(reader, binary.BigEndian, &flags))
		ThrowIfError(binary.Read(reader, binary.BigEndian, &count))
		ThrowData(res{seq: seq, flags: flags, count: count})
	}); resErr != nil {
		return nil, 0, 0, resErr;
	} else {
		resData := data.(res)
		return resData.seq, resData.flags, resData.count, nil
	}
}

func parsePart(reader io.Reader) ([]byte, error) {
	if data, err := UnwrapResultFromJob(func() {
		var size uint16
		ThrowIfError(binary.Read(reader, binary.BigEndian, &size))
		buf := make([]byte, size)
		_, err := io.ReadFull(reader, buf)
		ThrowIfError(err)
		ThrowData(buf)
	}); err != nil {
		return nil, err
	} else {
		return data.([]byte), err
	}
}

func parseSuggest(body []byte) (*metadata.SuggestResult, error) {
	if data, err := UnwrapResultFromJob(func() {
		result := &metadata.SuggestResult{}
		ThrowIfError(json.Unmarshal(body, result))
		for _, s := range result.Sections {
			switch s.Typ {
			case "top-results":
				ThrowIfError(json.Unmarshal(s.RawItems, &result.TopHits))
			case "album-results":
				ThrowIfError(json.Unmarshal(s.RawItems, &result.Albums))
			case "artist-results":
				ThrowIfError(json.Unmarshal(s.RawItems, &result.Artists))
			case "track-results":
				ThrowIfError(json.Unmarshal(s.RawItems, &result.Tracks))
			}
		}
		ThrowData(result)
	}); err != nil {
		return nil, err
	} else {
		return data.(*metadata.SuggestResult), nil
	}
}
