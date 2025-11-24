package gptpType

import "encoding/binary"

type FollowUpInformation struct {
	TlvType                    uint16
	LengthField                uint16
	OrganizationId             uint32
	OrganizationSubType        uint32
	CumulativeScaledRateOffset uint32
	GmTimeBaseIndicator        uint16
	LastGmPhaseChange          [12]byte
	ScaledLastGmFreqChange     uint32
}

func EncodeFollowUpInformation(val *FollowUpInformation) []byte {

	b := []byte{}
	u_int16 := make([]byte, 2)
	binary.BigEndian.PutUint16(u_int16, val.TlvType)
	b = append(b, u_int16...)
	binary.BigEndian.PutUint16(u_int16, val.LengthField)
	b = append(b, u_int16...)
	u_int32 := make([]byte, 4)
	binary.BigEndian.PutUint32(u_int32, val.OrganizationId)
	b = append(b, u_int32[1:]...)
	binary.BigEndian.PutUint32(u_int32, val.OrganizationSubType)
	b = append(b, u_int32[1:]...)
	binary.BigEndian.PutUint32(u_int32, val.CumulativeScaledRateOffset)
	b = append(b, u_int32...)
	binary.BigEndian.PutUint16(u_int16, val.GmTimeBaseIndicator)
	b = append(b, u_int16...)
	b = append(b, val.LastGmPhaseChange[:]...)
	binary.BigEndian.PutUint32(u_int32, val.ScaledLastGmFreqChange)
	b = append(b, u_int32...)

	return b
}

func DecodeFollowUpInformation(b []byte, val *FollowUpInformation) {
	val.TlvType = binary.BigEndian.Uint16(b[0:2])
	val.LengthField = binary.BigEndian.Uint16(b[2:4])
	var organizationId [4]byte
	copy(organizationId[1:], b[4:7])
	val.OrganizationId = binary.BigEndian.Uint32(organizationId[:])
	var organizationSubType [4]byte
	copy(organizationSubType[1:], b[7:10])
	val.OrganizationSubType = binary.BigEndian.Uint32(organizationSubType[:])
	val.CumulativeScaledRateOffset = binary.BigEndian.Uint32(b[10:14])
	val.GmTimeBaseIndicator = binary.BigEndian.Uint16(b[14:16])
	copy(val.LastGmPhaseChange[:], b[16:28])
	val.ScaledLastGmFreqChange = binary.BigEndian.Uint32(b[28:32])
}

func (f *FollowUpInformation) GetTlvType() uint16 {
	return f.TlvType
}

func (f *FollowUpInformation) GetLengthField() uint16 {
	return f.LengthField
}

func (f *FollowUpInformation) GetOrganizationId() uint32 {
	return f.OrganizationSubType
}

func (f *FollowUpInformation) GetOrganizationSubType() uint32 {
	return f.OrganizationId
}

func (f *FollowUpInformation) GetCumulativeScaledRateOffset() uint32 {
	return f.CumulativeScaledRateOffset
}

func (f *FollowUpInformation) GetGmTimeBaseIndicator() uint16 {
	return f.GmTimeBaseIndicator
}

func (f *FollowUpInformation) GetLastGmPhaseChange() []byte {
	return f.LastGmPhaseChange[:]
}

func (f *FollowUpInformation) GetScaledLastGmFreqChange() uint32 {
	return f.ScaledLastGmFreqChange
}
