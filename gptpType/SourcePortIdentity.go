package gptpType

import "encoding/binary"

type SourcePortIdentity struct {
	ClockIdentity uint64
	SourcePortId  uint16
}

func EncodeSourcePortIdentity(val *SourcePortIdentity) []byte {

	b := []byte{}
	u_int64 := make([]byte, 8)
	binary.BigEndian.PutUint64(u_int64, val.ClockIdentity)
	b = append(b, u_int64...)
	u_int16 := make([]byte, 2)
	binary.BigEndian.PutUint16(u_int16, val.SourcePortId)
	b = append(b, u_int16...)

	return b
}

func DecodeSourcePortIdentity(b []byte, val *SourcePortIdentity) {
	val.ClockIdentity = binary.BigEndian.Uint64(b[0:8])
	val.SourcePortId = binary.BigEndian.Uint16(b[8:10])
}

func (s *SourcePortIdentity) GetClockIdentity() uint64 {
	return s.ClockIdentity
}

func (s *SourcePortIdentity) GetSourcePortId() uint16 {
	return s.SourcePortId
}
