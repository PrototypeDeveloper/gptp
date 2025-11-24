package gptp

import (
	"encoding/binary"
	"fmt"
	"gptp/gptpMessage"
)

// Type definitions.
const (
	TypePTPv2OverEthernet = 0x88f7
)

// Length
const (
	MsgLength    = 18
	HeaderLength = 14
)

func Encoder(msg gptpMessage.Message) (result []byte, err error) {

	switch m := msg.(type) {
	case *gptpMessage.SyncMessage:
		result, err = gptpMessage.EncodeSyncMessage(m)
	case *gptpMessage.PeerDelayReqMessage:
		result, err = gptpMessage.EncodePeerDelayReqMessage(m)
	case *gptpMessage.PeerDelayRespMessage:
		result, err = gptpMessage.EncodePeerDelayRespMessage(m)
	case *gptpMessage.FollowUpMessage:
		result, err = gptpMessage.EncodeFollowUpMessage(m)
	case *gptpMessage.PeerDelayRespFollowUpMessage:
		result, err = gptpMessage.EncodePeerDelayRespFollowUpMessage(m)
	case *gptpMessage.AnnounceMessage:
		result, err = gptpMessage.EncodeAnnounceMessage(m)
	default:
		err = fmt.Errorf("illegal gptp message")
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

func Decoder(b []byte) (gptpMessage.Message, error) {

	if len(b) < MsgLength {
		return nil, fmt.Errorf("illegal message length")
	}

	headerType := binary.BigEndian.Uint16(b[12:14])
	if headerType != TypePTPv2OverEthernet {
		return nil, fmt.Errorf("illegal message types")
	}

	messageLength := binary.BigEndian.Uint16(b[16:18])
	sumLength := messageLength + HeaderLength
	if len(b) < int(sumLength) {
		return nil, fmt.Errorf("illegal gptp message length")
	}

	var m gptpMessage.Message
	var err error
	messageType := uint8(b[14] & 0x0f)
	switch messageType {
	case gptpMessage.MsgTypeSyncMessage:
		m, err = gptpMessage.DecodeSyncMessage(b)
	case gptpMessage.MsgTypePeerDelayReqMessage:
		m, err = gptpMessage.DecodePeerDelayReqMessage(b)
	case gptpMessage.MsgTypePeerDelayRespMessage:
		m, err = gptpMessage.DecodePeerDelayRespMessage(b)
	case gptpMessage.MsgTypeFollowUpMessage:
		m, err = gptpMessage.DecodeFollowUpMessage(b)
	case gptpMessage.MsgTypePeerDelayRespFollowUpMessage:
		m, err = gptpMessage.DecodePeerDelayRespFollowUpMessage(b)
	case gptpMessage.MsgTypeAnnounceMessage:
		m, err = gptpMessage.DecodeAnnounceMessage(b)
	default:
		err = fmt.Errorf("illegal gptp message type")
	}

	if err != nil {
		return nil, err
	}

	return m, nil
}
