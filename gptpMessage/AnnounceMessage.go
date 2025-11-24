package gptpMessage

import (
	"encoding/binary"
	"fmt"
	"gptp/gptpType"
)

const AnnounceMessageLength = 90
const AnnouncePTPMessageLength = 76

type AnnounceMessage struct {
	gptpType.Ethernet
	PTPMessage AnnouncePTPMessage
}

type AnnouncePTPMessage struct {
	gptpType.Header
	Body AnnouncePTPMessageBody
}

type AnnouncePTPMessageBody struct {
	Reserved1               [10]byte
	CurrentUtcOffset        uint16
	Reserved2               uint8
	GrandmasterPriority1    uint8
	GrandmasterClockQuality gptpType.GrandmasterClockQuality
	GrandmasterPriority2    uint8
	GrandmasterIdentity     uint64
	StepsRemoved            uint16
	TimeSource              uint8
	PathTraceTLV            gptpType.PathTraceTLV
}

func EncodeAnnounceMessage(msg *AnnounceMessage) ([]byte, error) {

	result := make([]byte, 0)

	// Length Check
	if msg.GetHeader().MessageLength < AnnouncePTPMessageLength {
		return result, fmt.Errorf("illegal announce message length")
	}

	// Encode Ethernet
	result = append(result, gptpType.EncodeEthernet(msg.GetEthernet())...)

	// Encode Header
	result = append(result, gptpType.EncodeHeader(msg.GetHeader())...)

	// Encode Announce Message Body
	result = append(result, EncodeAnnounceMessageBody(msg.GetBody())...)

	// Update Message Length
	if binary.BigEndian.Uint16(result[MsgLength:MsgLength+2]) != uint16(len(result[14:])) {
		u_int16 := make([]byte, 2)
		binary.BigEndian.PutUint16(u_int16, uint16(len(result[14:])))
		result[MsgLength] = u_int16[0]
		result[MsgLength] = u_int16[1]
	}

	return result, nil
}

func DecodeAnnounceMessage(b []byte) (*AnnounceMessage, error) {

	msg := &AnnounceMessage{}

	// Length Check
	if len(b) < AnnounceMessageLength {
		return msg, fmt.Errorf("illegal announce message length")
	}

	// Decode Ethernet
	gptpType.DecodeEthernet(b, msg.GetEthernet())

	// Decode Header
	gptpType.DeocdeHeader(b[14:], msg.GetHeader())

	// Decode Announce Message Body
	DecodeAnnounceMessageBody(b[48:], msg.GetBody())

	return msg, nil
}

func EncodeAnnounceMessageBody(body *AnnouncePTPMessageBody) []byte {

	b := []byte{}
	b = append(b, body.Reserved1[:]...)
	u_int16 := make([]byte, 2)
	binary.BigEndian.PutUint16(u_int16, body.CurrentUtcOffset)
	b = append(b, u_int16...)
	b = append(b, body.Reserved2)
	b = append(b, body.GrandmasterPriority1)
	b = append(b, gptpType.EncodeGrandmasterClockQuality(body.GetGrandmasterClockQuality())...)
	b = append(b, body.GrandmasterPriority2)
	u_int64 := make([]byte, 8)
	binary.BigEndian.PutUint64(u_int64, body.GrandmasterIdentity)
	b = append(b, u_int64...)
	binary.BigEndian.PutUint16(u_int16, body.StepsRemoved)
	b = append(b, u_int16...)
	b = append(b, body.TimeSource)
	b = append(b, gptpType.EncodePathTraceTLV(body.GetPathTraceTLV())...)
	return b
}

func DecodeAnnounceMessageBody(b []byte, body *AnnouncePTPMessageBody) {
	copy(body.Reserved1[:], b[0:10])
	body.CurrentUtcOffset = binary.BigEndian.Uint16(b[10:12])
	body.Reserved2 = b[12]
	body.GrandmasterPriority1 = b[13]
	gptpType.DecodeGrandmasterClockQuality(b[14:18], body.GetGrandmasterClockQuality())
	body.GrandmasterPriority1 = b[18]
	body.GrandmasterIdentity = binary.BigEndian.Uint64(b[19:27])
	body.StepsRemoved = binary.BigEndian.Uint16(b[27:29])
	body.TimeSource = b[29]
	gptpType.DecodePathTraceTLV(b[30:], body.GetPathTraceTLV())
}

func (m *AnnounceMessage) GetEthernet() *gptpType.Ethernet {
	return &m.Ethernet
}

func (m *AnnounceMessage) GetHeader() *gptpType.Header {
	return &m.PTPMessage.Header
}

func (m *AnnounceMessage) GetBody() *AnnouncePTPMessageBody {
	return &m.PTPMessage.Body
}

func (a *AnnouncePTPMessageBody) GetReserved1() []byte {
	return a.Reserved1[:]
}

func (a *AnnouncePTPMessageBody) GetPCurrentUtcOffset() uint16 {
	return a.CurrentUtcOffset
}

func (a *AnnouncePTPMessageBody) GetPReserved2() uint8 {
	return a.Reserved2
}

func (a *AnnouncePTPMessageBody) GetGrandmasterPriority1() uint8 {
	return a.GrandmasterPriority1
}

func (a *AnnouncePTPMessageBody) GetGrandmasterClockQuality() *gptpType.GrandmasterClockQuality {
	return &a.GrandmasterClockQuality
}

func (a *AnnouncePTPMessageBody) GetGrandmasterPriority2() uint8 {
	return a.GrandmasterPriority2
}

func (a *AnnouncePTPMessageBody) GetGrandmasterIdentity() uint64 {
	return a.GrandmasterIdentity
}

func (a *AnnouncePTPMessageBody) GetStepsRemoved() uint16 {
	return a.StepsRemoved
}

func (a *AnnouncePTPMessageBody) GetTimeSource() uint8 {
	return a.TimeSource
}

func (a *AnnouncePTPMessageBody) GetPathTraceTLV() *gptpType.PathTraceTLV {
	return &a.PathTraceTLV
}
