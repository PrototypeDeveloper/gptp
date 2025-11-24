package gptpMessage

import (
	"encoding/binary"
	"fmt"
	"gptp/gptpType"
)

const PeerDelayReqMessageLength = 68
const PeerDelayReqPTPMessageLength = 54

type PeerDelayReqMessage struct {
	gptpType.Ethernet
	PTPMessage PeerDelayReqPTPMessage
}

type PeerDelayReqPTPMessage struct {
	gptpType.Header
	Body PeerDelayReqPTPMessageBody
}

type PeerDelayReqPTPMessageBody struct {
	Reserved1 [10]byte
	Reserved2 [10]byte
}

func EncodePeerDelayReqMessage(msg *PeerDelayReqMessage) ([]byte, error) {

	result := make([]byte, 0)

	// Length Check
	if msg.GetHeader().MessageLength < PeerDelayReqPTPMessageLength {
		return result, fmt.Errorf("illegal peer Delay req message length")
	}

	// Encode Ethernet
	result = append(result, gptpType.EncodeEthernet(msg.GetEthernet())...)

	// Encode Header
	result = append(result, gptpType.EncodeHeader(msg.GetHeader())...)

	// Encode Peer Delay Req Message Body
	result = append(result, EncodePeerDelayReqMessageBody(msg.GetBody())...)

	// Update Message Length
	if binary.BigEndian.Uint16(result[MsgLength:MsgLength+2]) != uint16(len(result[14:])) {
		u_int16 := make([]byte, 2)
		binary.BigEndian.PutUint16(u_int16, uint16(len(result[14:])))
		result[MsgLength] = u_int16[0]
		result[MsgLength] = u_int16[1]
	}

	return result, nil
}

func DecodePeerDelayReqMessage(b []byte) (*PeerDelayReqMessage, error) {

	msg := &PeerDelayReqMessage{}

	// Length Check
	if len(b) < PeerDelayReqMessageLength {
		return msg, fmt.Errorf("illegal peer Delay req message length")
	}

	// Decode Ethernet
	gptpType.DecodeEthernet(b, msg.GetEthernet())

	// Decode Header
	gptpType.DeocdeHeader(b[14:], msg.GetHeader())

	// Decode Peer Delay Req Message Body
	DecodePeerDelayReqMessageBody(b[48:], msg.GetBody())

	return msg, nil
}

func EncodePeerDelayReqMessageBody(body *PeerDelayReqPTPMessageBody) []byte {

	b := body.Reserved1[:]
	b = append(b, body.Reserved2[:]...)
	return b
}

func DecodePeerDelayReqMessageBody(b []byte, body *PeerDelayReqPTPMessageBody) {
	copy(body.Reserved1[:], b[0:10])
	copy(body.Reserved2[:], b[10:20])
}

func (m *PeerDelayReqMessage) GetEthernet() *gptpType.Ethernet {
	return &m.Ethernet
}

func (m *PeerDelayReqMessage) GetHeader() *gptpType.Header {
	return &m.PTPMessage.Header
}

func (m *PeerDelayReqMessage) GetBody() *PeerDelayReqPTPMessageBody {
	return &m.PTPMessage.Body
}

func (s *PeerDelayReqPTPMessageBody) GetReserved1() []byte {
	return s.Reserved1[:]
}

func (s *PeerDelayReqPTPMessageBody) GetReserved2() []byte {
	return s.Reserved2[:]
}
