package gptpType

import "encoding/binary"

type PathTraceTLV struct {
	TlvType      uint16
	LengthField  uint16
	PathSequence []byte
}

func EncodePathTraceTLV(val *PathTraceTLV) []byte {

	b := []byte{}
	u_int16 := make([]byte, 2)
	binary.BigEndian.PutUint16(u_int16, val.TlvType)
	b = append(b, u_int16...)
	binary.BigEndian.PutUint16(u_int16, val.LengthField)
	b = append(b, u_int16...)
	b = append(b, val.PathSequence...)

	return b
}

func DecodePathTraceTLV(b []byte, val *PathTraceTLV) {
	val.TlvType = binary.BigEndian.Uint16(b[0:2])
	val.LengthField = binary.BigEndian.Uint16(b[2:4])
	val.PathSequence = make([]byte, val.LengthField)
	copy(val.PathSequence, b[4:(4+val.LengthField)])
}

func (p *PathTraceTLV) GetClass() uint16 {
	return p.TlvType
}

func (p *PathTraceTLV) GetLengthField() uint16 {
	return p.LengthField
}

func (p *PathTraceTLV) GetPathSequence() []byte {
	return p.PathSequence
}
