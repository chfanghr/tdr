package mercury

import (
	"encoding/binary"
	"fmt"
	"github.com/chfanghr/tdr/spotify/connection"
	spot "github.com/chfanghr/tdr/spotify/proto"
	"github.com/golang/protobuf/proto"
	"io"
	"sync"
)

type internal struct {
	seqLock    sync.Mutex
	nextSeqNum uint32
	pending    map[string]Pending
	stream     connection.PacketStream
}

func (m *internal) nextSeq() (uint32, []byte) {
	m.seqLock.Lock()

	seq := make([]byte, 4)
	seqInt := m.nextSeqNum
	binary.BigEndian.PutUint32(seq, seqInt)
	m.nextSeqNum += 1
	m.seqLock.Unlock()

	return seqInt, seq
}

func (m *internal) request(req Request) (seqKey string, err error) {
	_, seq := m.nextSeq()
	data, err := encodeRequest(seq, req)
	if err != nil {
		return "", err
	}

	var cmd uint8
	switch {
	case req.Method == "SUB":
		cmd = 0xb3
	case req.Method == "UNSUB":
		cmd = 0xb4
	default:
		cmd = 0xb2
	}

	err = m.stream.SendPacket(cmd, data)
	if err != nil {
		return "", err
	}

	return string(seq), nil
}

func (m *internal) parseResponse(cmd uint8, reader io.Reader) (response *Response, err error) {
	seq, flags, count, err := handleHead(reader)
	if err != nil {
		fmt.Println("error handling response", err)
		return
	}

	seqKey := string(seq)
	pending, ok := m.pending[seqKey]

	if !ok && cmd == 0xb5 {
		pending = Pending{}
	} else if !ok {
		//log.Print("ignoring seq ", SeqKey)
	}

	for i := uint16(0); i < count; i++ {
		part, err := parsePart(reader)
		if err != nil {
			fmt.Println("read part")
			return nil, err
		}

		if pending.partial != nil {
			part = append(pending.partial, part...)
			pending.partial = nil
		}

		if i == count-1 && (flags == 2) {
			pending.partial = part
		} else {
			pending.parts = append(pending.parts, part)
		}
	}

	if flags == 1 {
		delete(m.pending, seqKey)
		return m.completeRequest(cmd, pending, seqKey)
	} else {
		m.pending[seqKey] = pending
	}
	return nil, nil
}

func (m *internal) completeRequest(cmd uint8, pending Pending, seqKey string) (response *Response, err error) {
	headerData := pending.parts[0]
	header := &spot.Header{}
	err = proto.Unmarshal(headerData, header)
	if err != nil {
		return nil, err
	}

	return &Response{
		HeaderData: headerData,
		Uri:        *header.Uri,
		Payload:    pending.parts[1:],
		StatusCode: header.GetStatusCode(),
		SeqKey:     seqKey,
	}, nil

}
