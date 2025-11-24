package gptpType

import "encoding/binary"

type CorrectionField struct {
	CorrectionNs    uint64
	CorrectionSubNs uint16
}

func EncodeCorrectionField(val *CorrectionField) []byte {

	b := []byte{}
	u_int64 := make([]byte, 8)
	binary.BigEndian.PutUint64(u_int64, val.CorrectionNs)
	b = append(b, u_int64[2:]...)
	u_int16 := make([]byte, 2)
	binary.BigEndian.PutUint16(u_int16, val.GetCorrectionSubNs())
	b = append(b, u_int16...)

	return b
}

func DecodeCorrectionField(b []byte, val *CorrectionField) {
	var ns [8]byte
	copy(ns[2:], b[0:6])
	val.CorrectionNs = binary.BigEndian.Uint64(ns[:])
	val.CorrectionSubNs = binary.BigEndian.Uint16(b[6:8])
}

func (c *CorrectionField) GetCorrectionNs() uint64 {
	return c.CorrectionNs
}

func (c *CorrectionField) GetCorrectionSubNs() uint16 {
	return c.CorrectionSubNs
}
