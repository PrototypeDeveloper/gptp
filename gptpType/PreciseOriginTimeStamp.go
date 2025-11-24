package gptpType

import "encoding/binary"

type PreciseOriginTimeStamp struct {
	Seconds     uint64
	NanoSeconds uint32
}

func EncodePreciseOriginTimeStamp(val *PreciseOriginTimeStamp) []byte {

	b := []byte{}
	u_int64 := make([]byte, 8)
	binary.BigEndian.PutUint64(u_int64, val.Seconds)
	b = append(b, u_int64[2:]...)
	u_int32 := make([]byte, 4)
	binary.BigEndian.PutUint32(u_int32, val.NanoSeconds)
	b = append(b, u_int32...)

	return b
}

func DecodePreciseOriginTimeStamp(b []byte, val *PreciseOriginTimeStamp) {
	var seconds [8]byte
	copy(seconds[2:], b[0:6])
	val.Seconds = binary.BigEndian.Uint64(seconds[:])
	val.NanoSeconds = binary.BigEndian.Uint32(b[6:10])
}

func (p *PreciseOriginTimeStamp) GetSeconds() uint64 {
	return p.Seconds
}

func (p *PreciseOriginTimeStamp) GetNanoSeconds() uint32 {
	return p.NanoSeconds
}
