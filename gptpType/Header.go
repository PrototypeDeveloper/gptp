package gptpType

import (
	"encoding/binary"
)

type Header struct {
	MajorSdoId          uint8
	MessageType         uint8
	MinorVersionPTP     uint8
	VersionPTP          uint8
	MessageLength       uint16
	DomainNumber        uint8
	MinorSdoId          uint8
	Flag                uint16
	CorrectionField     CorrectionField
	MessageTypeSpecific uint32
	SourcePortIdentity  SourcePortIdentity
	SequenceId          uint16
	ControlField        uint8
	LogMessageInterval  uint8
}

func EncodeHeader(header *Header) []byte {

	b := []byte{}
	b = append(b, (byte)((header.MajorSdoId<<4)+(header.MessageType&0x0f)))
	b = append(b, (byte)((header.MinorVersionPTP<<4)+(header.VersionPTP&0x0f)))
	u_int16 := make([]byte, 2)
	binary.BigEndian.PutUint16(u_int16, header.MessageLength)
	b = append(b, u_int16...)
	b = append(b, header.DomainNumber)
	b = append(b, header.MinorSdoId)
	binary.BigEndian.PutUint16(u_int16, header.Flag)
	b = append(b, u_int16...)
	b = append(b, EncodeCorrectionField(header.GetCorrectionField())...)
	u_int32 := make([]byte, 4)
	binary.BigEndian.PutUint32(u_int32, header.MessageTypeSpecific)
	b = append(b, u_int32...)
	b = append(b, EncodeSourcePortIdentity(header.GetSourcePortIdentity())...)
	binary.BigEndian.PutUint16(u_int16, header.SequenceId)
	b = append(b, u_int16...)
	b = append(b, header.ControlField)
	b = append(b, header.LogMessageInterval)

	return b
}

func DeocdeHeader(b []byte, header *Header) {

	header.MajorSdoId = b[0] >> 4
	header.MessageType = b[0] & 0x0f
	header.MinorVersionPTP = b[1] >> 4
	header.VersionPTP = b[1] & 0x0f
	header.MessageLength = binary.BigEndian.Uint16(b[2:4])
	header.DomainNumber = b[4]
	header.MinorSdoId = b[5]
	header.Flag = binary.BigEndian.Uint16(b[6:8])
	DecodeCorrectionField(b[8:16], header.GetCorrectionField())
	header.MessageTypeSpecific = binary.BigEndian.Uint32(b[16:20])
	DecodeSourcePortIdentity(b[20:30], header.GetSourcePortIdentity())
	header.SequenceId = binary.BigEndian.Uint16(b[30:32])
	header.ControlField = b[32]
	header.LogMessageInterval = b[33]
}

func (h *Header) GetMajorSdoId() uint8 {
	return h.MajorSdoId
}

func (h *Header) GetMessageType() uint8 {
	return h.MessageType
}

func (h *Header) GetMinorVersionPTP() uint8 {
	return h.MinorVersionPTP
}

func (h *Header) GetVersionPTP() uint8 {
	return h.VersionPTP
}

func (h *Header) GetMessageLength() uint16 {
	return h.MessageLength
}

func (h *Header) GetDomainNumber() uint8 {
	return h.DomainNumber
}

func (h *Header) GetMinorSdoId() uint8 {
	return h.MinorSdoId
}

func (h *Header) GetFlag() uint16 {
	return h.Flag
}

func (h *Header) GetCorrectionField() *CorrectionField {
	return &h.CorrectionField
}

func (h *Header) GetMessageTypeSpecific() uint32 {
	return h.MessageTypeSpecific
}

func (h *Header) GetSourcePortIdentity() *SourcePortIdentity {
	return &h.SourcePortIdentity
}

func (h *Header) GetSequenceId() uint16 {
	return h.SequenceId
}

func (h *Header) GetControlField() uint8 {
	return h.ControlField
}

func (h *Header) GetLogMessageInterval() uint8 {
	return h.LogMessageInterval
}
