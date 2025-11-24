package gptpMessage

import (
	"encoding/binary"
	"fmt"
	"gptp/gptpType"
)

const FollowUpMessageLength = 90
const FollowUpPTPMessageLength = 76

type FollowUpMessage struct {
	gptpType.Ethernet
	PTPMessage FollowUpPTPMessage
}

type FollowUpPTPMessage struct {
	gptpType.Header
	Body FollowUpPTPMessageBody
}

type FollowUpPTPMessageBody struct {
	PreciseOriginTimeStamp gptpType.PreciseOriginTimeStamp
	FollowUpInformation    gptpType.FollowUpInformation
}

func EncodeFollowUpMessage(msg *FollowUpMessage) ([]byte, error) {

	result := make([]byte, 0)

	// Length Check
	if msg.GetHeader().MessageLength < FollowUpPTPMessageLength {
		return result, fmt.Errorf("illegal follow up message length")
	}

	// Encode Ethernet
	result = append(result, gptpType.EncodeEthernet(msg.GetEthernet())...)

	// Encode Header
	result = append(result, gptpType.EncodeHeader(msg.GetHeader())...)

	// Encode Follow Up Message Body
	result = append(result, EncodeFollowUpMessageBody(msg.GetBody())...)

	// Update Message Length
	if binary.BigEndian.Uint16(result[MsgLength:MsgLength+2]) != uint16(len(result[14:])) {
		u_int16 := make([]byte, 2)
		binary.BigEndian.PutUint16(u_int16, uint16(len(result[14:])))
		result[MsgLength] = u_int16[0]
		result[MsgLength] = u_int16[1]
	}

	return result, nil
}

func DecodeFollowUpMessage(b []byte) (*FollowUpMessage, error) {

	msg := &FollowUpMessage{}

	// Length Check
	if len(b) < FollowUpMessageLength {
		return msg, fmt.Errorf("illegal follow up message length")
	}

	// Decode Ethernet
	gptpType.DecodeEthernet(b, msg.GetEthernet())

	// Decode Header
	gptpType.DeocdeHeader(b[14:], msg.GetHeader())

	// Decode Follow Up Message Body
	DecodeFollowUpMessageBody(b[48:], msg.GetBody())

	return msg, nil
}

func EncodeFollowUpMessageBody(body *FollowUpPTPMessageBody) []byte {

	b := []byte{}
	b = append(b, gptpType.EncodePreciseOriginTimeStamp(body.GetPreciseOriginTimeStamp())...)
	b = append(b, gptpType.EncodeFollowUpInformation(body.GetFollowUpInformation())...)
	return b
}

func DecodeFollowUpMessageBody(b []byte, body *FollowUpPTPMessageBody) {
	gptpType.DecodePreciseOriginTimeStamp(b[0:10], body.GetPreciseOriginTimeStamp())
	gptpType.DecodeFollowUpInformation(b[10:], body.GetFollowUpInformation())
}

func (m *FollowUpMessage) GetEthernet() *gptpType.Ethernet {
	return &m.Ethernet
}

func (m *FollowUpMessage) GetHeader() *gptpType.Header {
	return &m.PTPMessage.Header
}

func (m *FollowUpMessage) GetBody() *FollowUpPTPMessageBody {
	return &m.PTPMessage.Body
}

func (p *FollowUpPTPMessageBody) GetPreciseOriginTimeStamp() *gptpType.PreciseOriginTimeStamp {
	return &p.PreciseOriginTimeStamp
}

func (p *FollowUpPTPMessageBody) GetFollowUpInformation() *gptpType.FollowUpInformation {
	return &p.FollowUpInformation
}
