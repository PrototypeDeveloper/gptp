package gptpType

import "encoding/binary"

type RequestingPortIdentity struct {
	Identity uint64
	Id       uint16
}

func EncodeRequestingPortIdentity(val *RequestingPortIdentity) []byte {

	b := []byte{}
	u_int64 := make([]byte, 8)
	binary.BigEndian.PutUint64(u_int64, val.Identity)
	b = append(b, u_int64...)
	u_int16 := make([]byte, 2)
	binary.BigEndian.PutUint16(u_int16, uint16(val.Id))
	b = append(b, u_int16...)

	return b
}

func DecodeRequestingPortIdentity(b []byte, val *RequestingPortIdentity) {
	val.Identity = binary.BigEndian.Uint64(b[0:8])
	val.Id = binary.BigEndian.Uint16(b[8:10])
}

func (r *RequestingPortIdentity) GetIdentity() uint64 {
	return r.Identity
}

func (r *RequestingPortIdentity) GetId() uint16 {
	return r.Id
}
