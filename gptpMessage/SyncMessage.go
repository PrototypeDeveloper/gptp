package gptpMessage

import (
	"encoding/binary"
	"fmt"
	"gptp/gptpType"
)

const SyncMessageLength = 58
const SyncPTPMessageLength = 44

type SyncMessage struct {
	gptpType.Ethernet
	PTPMessage SyncPTPMessage
}

type SyncPTPMessage struct {
	gptpType.Header
	Body SyncPTPMessageBody
}

type SyncPTPMessageBody struct {
	Reserved [10]byte
}

func EncodeSyncMessage(msg *SyncMessage) ([]byte, error) {

	result := make([]byte, 0)

	// Length Check
	if msg.GetHeader().MessageLength < SyncPTPMessageLength {
		return result, fmt.Errorf("illegal sync message length")
	}

	// Encode Ethernet
	result = append(result, gptpType.EncodeEthernet(msg.GetEthernet())...)

	// Encode Header
	result = append(result, gptpType.EncodeHeader(msg.GetHeader())...)

	// Encode Sync Message Body
	result = append(result, EncodeSyncMessageBody(msg.GetBody())...)

	// Update Message Length
	if binary.BigEndian.Uint16(result[MsgLength:MsgLength+2]) != uint16(len(result[14:])) {
		u_int16 := make([]byte, 2)
		binary.BigEndian.PutUint16(u_int16, uint16(len(result[14:])))
		result[MsgLength] = u_int16[0]
		result[MsgLength] = u_int16[1]
	}

	return result, nil
}

func DecodeSyncMessage(b []byte) (*SyncMessage, error) {

	msg := &SyncMessage{}

	// Length Check
	if len(b) < SyncMessageLength {
		return msg, fmt.Errorf("illegal sync message length")
	}

	// Decode Ethernet
	gptpType.DecodeEthernet(b, msg.GetEthernet())

	// Decode Header
	gptpType.DeocdeHeader(b[14:], msg.GetHeader())

	// Decode Sync Message Body
	DecodeSyncMessageBody(b[48:], msg.GetBody())

	return msg, nil
}

func EncodeSyncMessageBody(body *SyncPTPMessageBody) []byte {

	b := body.Reserved[:]
	return b
}

func DecodeSyncMessageBody(b []byte, body *SyncPTPMessageBody) {
	copy(body.Reserved[:], b[0:10])
}

func (m *SyncMessage) GetEthernet() *gptpType.Ethernet {
	return &m.Ethernet
}

func (m *SyncMessage) GetHeader() *gptpType.Header {
	return &m.PTPMessage.Header
}

func (m *SyncMessage) GetBody() *SyncPTPMessageBody {
	return &m.PTPMessage.Body
}

func (s *SyncPTPMessageBody) GetReserved() []byte {
	return s.Reserved[:]
}
