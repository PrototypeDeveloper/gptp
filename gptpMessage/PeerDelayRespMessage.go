package gptpMessage

import (
	"encoding/binary"
	"fmt"
	"gptp/gptpType"
)

const PeerDelayRespMessageLength = 68
const PeerDelayRespPTPMessageLength = 54

type PeerDelayRespMessage struct {
	gptpType.Ethernet
	PTPMessage PeerDelayRespPTPMessage
}

type PeerDelayRespPTPMessage struct {
	gptpType.Header
	Body PeerDelayRespPTPMessageBody
}

type PeerDelayRespPTPMessageBody struct {
	RequestReceiptTimeStamp gptpType.RequestReceiptTimeStamp
	RequestingPortIdentity  gptpType.RequestingPortIdentity
}

func EncodePeerDelayRespMessage(msg *PeerDelayRespMessage) ([]byte, error) {

	result := make([]byte, 0)

	// Length Check
	if msg.GetHeader().MessageLength < PeerDelayRespPTPMessageLength {
		return result, fmt.Errorf("illegal peer Delay Resp message length")
	}

	// Encode Ethernet
	result = append(result, gptpType.EncodeEthernet(msg.GetEthernet())...)

	// Encode Header
	result = append(result, gptpType.EncodeHeader(msg.GetHeader())...)

	// Encode Peer Delay Resp Message Body
	result = append(result, EncodePeerDelayRespMessageBody(msg.GetBody())...)

	// Update Message Length
	if binary.BigEndian.Uint16(result[MsgLength:MsgLength+2]) != uint16(len(result[14:])) {
		u_int16 := make([]byte, 2)
		binary.BigEndian.PutUint16(u_int16, uint16(len(result[14:])))
		result[MsgLength] = u_int16[0]
		result[MsgLength] = u_int16[1]
	}

	return result, nil
}

func DecodePeerDelayRespMessage(b []byte) (*PeerDelayRespMessage, error) {

	msg := &PeerDelayRespMessage{}

	// Length Check
	if len(b) < PeerDelayRespMessageLength {
		return msg, fmt.Errorf("illegal peer Delay Resp message length")
	}

	// Decode Ethernet
	gptpType.DecodeEthernet(b, msg.GetEthernet())

	// Decode Header
	gptpType.DeocdeHeader(b[14:], msg.GetHeader())

	// Decode Peer Delay Resp Message Body
	DecodePeerDelayRespMessageBody(b[48:], msg.GetBody())

	return msg, nil
}

func EncodePeerDelayRespMessageBody(body *PeerDelayRespPTPMessageBody) []byte {

	b := []byte{}
	b = append(b, gptpType.EncodeRequestReceiptTimeStamp(body.GetRequestReceiptTimeStamp())...)
	b = append(b, gptpType.EncodeRequestingPortIdentity(body.GetRequestingPortIdentity())...)
	return b
}

func DecodePeerDelayRespMessageBody(b []byte, body *PeerDelayRespPTPMessageBody) {
	gptpType.DecodeRequestReceiptTimeStamp(b[0:10], body.GetRequestReceiptTimeStamp())
	gptpType.DecodeRequestingPortIdentity(b[10:20], body.GetRequestingPortIdentity())
}

func (m *PeerDelayRespMessage) GetEthernet() *gptpType.Ethernet {
	return &m.Ethernet
}

func (m *PeerDelayRespMessage) GetHeader() *gptpType.Header {
	return &m.PTPMessage.Header
}

func (m *PeerDelayRespMessage) GetBody() *PeerDelayRespPTPMessageBody {
	return &m.PTPMessage.Body
}

func (p *PeerDelayRespPTPMessageBody) GetRequestReceiptTimeStamp() *gptpType.RequestReceiptTimeStamp {
	return &p.RequestReceiptTimeStamp
}

func (p *PeerDelayRespPTPMessageBody) GetRequestingPortIdentity() *gptpType.RequestingPortIdentity {
	return &p.RequestingPortIdentity
}
