package gptpMessage

import (
	"gptp/gptpType"
)

// Message Type definitions.
const (
	MsgTypeSyncMessage                  uint8 = 0x0
	MsgTypePeerDelayReqMessage          uint8 = 0x2
	MsgTypePeerDelayRespMessage         uint8 = 0x3
	MsgTypeFollowUpMessage              uint8 = 0x8
	MsgTypePeerDelayRespFollowUpMessage uint8 = 0xa
	MsgTypeAnnounceMessage              uint8 = 0xb
)

// Message Length
const (
	MsgLength uint16 = 16
)

type Message interface {
	GetEthernet() *gptpType.Ethernet
	GetHeader() *gptpType.Header
}
