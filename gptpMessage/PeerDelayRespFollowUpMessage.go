package gptpMessage

import (
	"encoding/binary"
	"fmt"
	"gptp/gptpType"
)

const PeerDelayRespFollowUpMessageLength = 68
const PeerDelayRespFollowUpPTPMessageLength = 54

type PeerDelayRespFollowUpMessage struct {
	gptpType.Ethernet
	PTPMessage PeerDelayRespFollowUpPTPMessage
}

type PeerDelayRespFollowUpPTPMessage struct {
	gptpType.Header
	Body PeerDelayRespFollowUpPTPMessageBody
}

type PeerDelayRespFollowUpPTPMessageBody struct {
	ResponseOriginTimestamp gptpType.ResponseOriginTimestamp
	RequestingPortIdentity  gptpType.RequestingPortIdentity
}

func EncodePeerDelayRespFollowUpMessage(msg *PeerDelayRespFollowUpMessage) ([]byte, error) {

	result := make([]byte, 0)

	// Length Check
	if msg.GetHeader().MessageLength < PeerDelayRespFollowUpPTPMessageLength {
		return result, fmt.Errorf("illegal peer Delay Resp follow up message length")
	}

	// Encode Ethernet
	result = append(result, gptpType.EncodeEthernet(msg.GetEthernet())...)

	// Encode Header
	result = append(result, gptpType.EncodeHeader(msg.GetHeader())...)

	// Encode Peer Delay Resp follow up Message Body
	result = append(result, EncodePeerDelayRespFollowUpMessageBody(msg.GetBody())...)

	// Update Message Length
	if binary.BigEndian.Uint16(result[MsgLength:MsgLength+2]) != uint16(len(result[14:])) {
		u_int16 := make([]byte, 2)
		binary.BigEndian.PutUint16(u_int16, uint16(len(result[14:])))
		result[MsgLength] = u_int16[0]
		result[MsgLength] = u_int16[1]
	}

	return result, nil
}

func DecodePeerDelayRespFollowUpMessage(b []byte) (*PeerDelayRespFollowUpMessage, error) {

	msg := &PeerDelayRespFollowUpMessage{}

	// Length Check
	if len(b) < PeerDelayRespFollowUpMessageLength {
		return msg, fmt.Errorf("illegal peer Delay Resp follow up message length")
	}

	// Decode Ethernet
	gptpType.DecodeEthernet(b, msg.GetEthernet())

	// Decode Header
	gptpType.DeocdeHeader(b[14:], msg.GetHeader())

	// Decode Peer Delay Resp follow up Message Body
	DecodePeerDelayRespFollowUpMessageBody(b[48:], msg.GetBody())

	return msg, nil
}

func EncodePeerDelayRespFollowUpMessageBody(body *PeerDelayRespFollowUpPTPMessageBody) []byte {

	b := []byte{}
	b = append(b, gptpType.EncodeResponseOriginTimestamp(body.GetResponseOriginTimestamp())...)
	b = append(b, gptpType.EncodeRequestingPortIdentity(body.GetRequestingPortIdentity())...)
	return b
}

func DecodePeerDelayRespFollowUpMessageBody(b []byte, body *PeerDelayRespFollowUpPTPMessageBody) {
	gptpType.DecodeResponseOriginTimestamp(b[0:10], body.GetResponseOriginTimestamp())
	gptpType.DecodeRequestingPortIdentity(b[10:20], body.GetRequestingPortIdentity())
}

func (m *PeerDelayRespFollowUpMessage) GetEthernet() *gptpType.Ethernet {
	return &m.Ethernet
}

func (m *PeerDelayRespFollowUpMessage) GetHeader() *gptpType.Header {
	return &m.PTPMessage.Header
}

func (m *PeerDelayRespFollowUpMessage) GetBody() *PeerDelayRespFollowUpPTPMessageBody {
	return &m.PTPMessage.Body
}

func (p *PeerDelayRespFollowUpPTPMessageBody) GetResponseOriginTimestamp() *gptpType.ResponseOriginTimestamp {
	return &p.ResponseOriginTimestamp
}

func (p *PeerDelayRespFollowUpPTPMessageBody) GetRequestingPortIdentity() *gptpType.RequestingPortIdentity {
	return &p.RequestingPortIdentity
}
