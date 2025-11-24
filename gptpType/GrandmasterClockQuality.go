package gptpType

import "encoding/binary"

type GrandmasterClockQuality struct {
	Class    uint8
	Accuracy uint8
	Variance uint16
}

func EncodeGrandmasterClockQuality(val *GrandmasterClockQuality) []byte {

	b := []byte{}
	b = append(b, val.Class)
	b = append(b, val.Accuracy)
	u_int16 := make([]byte, 2)
	binary.BigEndian.PutUint16(u_int16, val.Variance)
	b = append(b, u_int16...)

	return b
}

func DecodeGrandmasterClockQuality(b []byte, val *GrandmasterClockQuality) {
	val.Class = b[0]
	val.Accuracy = b[1]
	val.Variance = binary.BigEndian.Uint16(b[2:4])
}

func (g *GrandmasterClockQuality) GetClass() uint8 {
	return g.Class
}

func (g *GrandmasterClockQuality) GetAccuracy() uint8 {
	return g.Accuracy
}

func (g *GrandmasterClockQuality) GetVariance() uint16 {
	return g.Variance
}
