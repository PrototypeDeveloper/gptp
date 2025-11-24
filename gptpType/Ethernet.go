package gptpType

import (
	"encoding/binary"
	"net"
)

type Ethernet struct {
	Distination net.HardwareAddr
	Source      net.HardwareAddr
	Type        uint16
}

func EncodeEthernet(eth *Ethernet) []byte {

	b := []byte{}
	b = append(b, eth.Distination...)
	b = append(b, eth.Source...)
	u_int16 := make([]byte, 2)
	binary.BigEndian.PutUint16(u_int16, eth.Type)
	b = append(b, u_int16...)

	return b
}

func DecodeEthernet(b []byte, eth *Ethernet) {
	eth.Distination = b[0:6]
	eth.Source = b[6:12]
	eth.Type = binary.BigEndian.Uint16(b[12:14])
}

func (e *Ethernet) GetDistination() net.HardwareAddr {
	return e.Distination
}

func (e *Ethernet) GetSource() net.HardwareAddr {
	return e.Source
}

func (e *Ethernet) GetType() uint16 {
	return e.Type
}
