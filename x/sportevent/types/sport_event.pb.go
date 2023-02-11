// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/sportevent/sport_event.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// SportEventStatus is the sport-event status enumeration
type SportEventStatus int32

const (
	// event is pending
	SportEventStatus_SPORT_EVENT_STATUS_UNSPECIFIED SportEventStatus = 0
	// invalid event
	SportEventStatus_SPORT_EVENT_STATUS_INVALID SportEventStatus = 1
	// event canceled
	SportEventStatus_SPORT_EVENT_STATUS_CANCELLED SportEventStatus = 2
	// event aborted
	SportEventStatus_SPORT_EVENT_STATUS_ABORTED SportEventStatus = 3
	// result of the event is declared
	SportEventStatus_SPORT_EVENT_STATUS_RESULT_DECLARED SportEventStatus = 4
)

var SportEventStatus_name = map[int32]string{
	0: "SPORT_EVENT_STATUS_UNSPECIFIED",
	1: "SPORT_EVENT_STATUS_INVALID",
	2: "SPORT_EVENT_STATUS_CANCELLED",
	3: "SPORT_EVENT_STATUS_ABORTED",
	4: "SPORT_EVENT_STATUS_RESULT_DECLARED",
}

var SportEventStatus_value = map[string]int32{
	"SPORT_EVENT_STATUS_UNSPECIFIED":     0,
	"SPORT_EVENT_STATUS_INVALID":         1,
	"SPORT_EVENT_STATUS_CANCELLED":       2,
	"SPORT_EVENT_STATUS_ABORTED":         3,
	"SPORT_EVENT_STATUS_RESULT_DECLARED": 4,
}

func (x SportEventStatus) String() string {
	return proto.EnumName(SportEventStatus_name, int32(x))
}

func (SportEventStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f4c38f73099259f8, []int{0}
}

// SportEvent the representation of the sport-event to be stored in sport-event.
// state
type SportEvent struct {
	// uid is the universal unique identifier of the sport-event.
	UID string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid"`
	// start_ts is the start timestamp of the sport-event.
	StartTS uint64 `protobuf:"varint,2,opt,name=start_ts,proto3" json:"start_ts"`
	// end_ts is the end timestamp of the sport-event.
	EndTS uint64 `protobuf:"varint,3,opt,name=end_ts,proto3" json:"end_ts"`
	// odds is the list of odds of the sport-event.
	Odds []*Odds `protobuf:"bytes,4,rep,name=odds,proto3" json:"odds,omitempty"`
	// winner_odds_uids is the list of winner odds universal unique identifiers.
	WinnerOddsUIDs []string `protobuf:"bytes,5,rep,name=winner_odds_uids,proto3" json:"winner_odds_uids"`
	// status is the current status of the sport-event.
	Status SportEventStatus `protobuf:"varint,6,opt,name=status,proto3,enum=sgenetwork.sge.sportevent.SportEventStatus" json:"status,omitempty"`
	// resolution_ts is the timestamp of the resolution of sport-event.
	ResolutionTS uint64 `protobuf:"varint,7,opt,name=resolution_ts,proto3" json:"resolution_ts"`
	// creator is the address of the creator of sport-event.
	Creator string `protobuf:"bytes,8,opt,name=creator,proto3" json:"creator,omitempty"`
	// bet_constraints holds the constraints of sport-event to accept bets.
	BetConstraints *EventBetConstraints `protobuf:"bytes,9,opt,name=bet_constraints,json=betConstraints,proto3" json:"bet_constraints,omitempty"`
	// active is the status of active or inactive sport-event.
	Active bool `protobuf:"varint,10,opt,name=active,proto3" json:"active,omitempty"`
	// meta contains human-readable metadata of the sport-event.
	Meta                   string                                 `protobuf:"bytes,11,opt,name=meta,proto3" json:"meta,omitempty"`
	SrContributionForHouse github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,12,opt,name=sr_contribution_for_house,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"sr_contribution_for_house"`
	BookId                 string                                 `protobuf:"bytes,13,opt,name=book_id,proto3" json:"book_id"`
}

func (m *SportEvent) Reset()         { *m = SportEvent{} }
func (m *SportEvent) String() string { return proto.CompactTextString(m) }
func (*SportEvent) ProtoMessage()    {}
func (*SportEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4c38f73099259f8, []int{0}
}
func (m *SportEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SportEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SportEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SportEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SportEvent.Merge(m, src)
}
func (m *SportEvent) XXX_Size() int {
	return m.Size()
}
func (m *SportEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_SportEvent.DiscardUnknown(m)
}

var xxx_messageInfo_SportEvent proto.InternalMessageInfo

func (m *SportEvent) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *SportEvent) GetStartTS() uint64 {
	if m != nil {
		return m.StartTS
	}
	return 0
}

func (m *SportEvent) GetEndTS() uint64 {
	if m != nil {
		return m.EndTS
	}
	return 0
}

func (m *SportEvent) GetOdds() []*Odds {
	if m != nil {
		return m.Odds
	}
	return nil
}

func (m *SportEvent) GetWinnerOddsUIDs() []string {
	if m != nil {
		return m.WinnerOddsUIDs
	}
	return nil
}

func (m *SportEvent) GetStatus() SportEventStatus {
	if m != nil {
		return m.Status
	}
	return SportEventStatus_SPORT_EVENT_STATUS_UNSPECIFIED
}

func (m *SportEvent) GetResolutionTS() uint64 {
	if m != nil {
		return m.ResolutionTS
	}
	return 0
}

func (m *SportEvent) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *SportEvent) GetBetConstraints() *EventBetConstraints {
	if m != nil {
		return m.BetConstraints
	}
	return nil
}

func (m *SportEvent) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *SportEvent) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

func (m *SportEvent) GetBookId() string {
	if m != nil {
		return m.BookId
	}
	return ""
}

// Bet constraints parent group for a sport-event
type EventBetConstraints struct {
	// min_amount is the minimum allowed bet amount for a sport-event.
	MinAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=min_amount,json=minAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_amount"`
	// bet_fee is the fee that thebettor needs to pay to bet on the sport-event.
	BetFee types.Coin `protobuf:"bytes,3,opt,name=bet_fee,json=betFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coin" json:"bet_fee"`
}

func (m *EventBetConstraints) Reset()         { *m = EventBetConstraints{} }
func (m *EventBetConstraints) String() string { return proto.CompactTextString(m) }
func (*EventBetConstraints) ProtoMessage()    {}
func (*EventBetConstraints) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4c38f73099259f8, []int{1}
}
func (m *EventBetConstraints) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventBetConstraints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventBetConstraints.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventBetConstraints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventBetConstraints.Merge(m, src)
}
func (m *EventBetConstraints) XXX_Size() int {
	return m.Size()
}
func (m *EventBetConstraints) XXX_DiscardUnknown() {
	xxx_messageInfo_EventBetConstraints.DiscardUnknown(m)
}

var xxx_messageInfo_EventBetConstraints proto.InternalMessageInfo

func (m *EventBetConstraints) GetBetFee() types.Coin {
	if m != nil {
		return m.BetFee
	}
	return types.Coin{}
}

func init() {
	proto.RegisterEnum("sgenetwork.sge.sportevent.SportEventStatus", SportEventStatus_name, SportEventStatus_value)
	proto.RegisterType((*SportEvent)(nil), "sgenetwork.sge.sportevent.SportEvent")
	proto.RegisterType((*EventBetConstraints)(nil), "sgenetwork.sge.sportevent.EventBetConstraints")
}

func init() { proto.RegisterFile("sge/sportevent/sport_event.proto", fileDescriptor_f4c38f73099259f8) }

var fileDescriptor_f4c38f73099259f8 = []byte{
	// 763 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x8e, 0xda, 0x46,
	0x14, 0xc6, 0x81, 0x98, 0x65, 0x48, 0x28, 0x9a, 0xb4, 0xa9, 0x59, 0x55, 0x1e, 0x97, 0x8b, 0x2d,
	0x6a, 0xb4, 0xb6, 0xc2, 0x3e, 0x01, 0xfe, 0xd9, 0xd4, 0x12, 0x65, 0xa3, 0x31, 0x24, 0x52, 0x6f,
	0x2c, 0x1b, 0x4f, 0x88, 0x45, 0xf1, 0xac, 0x3c, 0xc3, 0xa6, 0x7d, 0x8b, 0xbe, 0x40, 0xaf, 0x7a,
	0xd7, 0xa7, 0xe8, 0x65, 0x7a, 0x97, 0xcb, 0xaa, 0x17, 0x6e, 0xe5, 0xbd, 0xe3, 0x29, 0xaa, 0x19,
	0xbc, 0x7f, 0x09, 0x44, 0xcd, 0x0d, 0x3e, 0x67, 0xbe, 0xf3, 0xcd, 0xf9, 0xf8, 0x7c, 0x8e, 0x81,
	0xc1, 0x16, 0xc4, 0x62, 0xe7, 0x34, 0xe7, 0xe4, 0x82, 0x64, 0x7c, 0x1b, 0x86, 0x32, 0x36, 0xcf,
	0x73, 0xca, 0x29, 0xec, 0xb1, 0x05, 0xc9, 0x08, 0x7f, 0x43, 0xf3, 0xa5, 0xc9, 0x16, 0xc4, 0xbc,
	0x29, 0x3e, 0xfc, 0x7c, 0x41, 0x17, 0x54, 0x56, 0x59, 0x22, 0xda, 0x12, 0x0e, 0xf5, 0x39, 0x65,
	0x2b, 0xca, 0xac, 0x38, 0x62, 0xc4, 0xba, 0x78, 0x1a, 0x13, 0x1e, 0x3d, 0xb5, 0xe6, 0x34, 0xcd,
	0x2a, 0xbc, 0xf7, 0x5e, 0x4b, 0x9a, 0x24, 0x6c, 0x0b, 0xf5, 0x7f, 0x55, 0x01, 0x08, 0x04, 0xe2,
	0x09, 0x04, 0x1a, 0xa0, 0xbe, 0x4e, 0x13, 0x4d, 0x31, 0x94, 0x41, 0xcb, 0xee, 0x94, 0x05, 0xaa,
	0xcf, 0x7c, 0x77, 0x53, 0x20, 0x71, 0x8a, 0xc5, 0x0f, 0x3c, 0x01, 0x07, 0x8c, 0x47, 0x39, 0x0f,
	0x39, 0xd3, 0xee, 0x19, 0xca, 0xa0, 0x61, 0x7f, 0x59, 0x16, 0xa8, 0x19, 0x88, 0xb3, 0x69, 0xb0,
	0x29, 0xd0, 0x35, 0x8c, 0xaf, 0x23, 0xf8, 0x04, 0xa8, 0x24, 0x4b, 0x04, 0xa5, 0x2e, 0x29, 0x8f,
	0xca, 0x02, 0xdd, 0xf7, 0xb2, 0x44, 0x12, 0x2a, 0x08, 0x57, 0x4f, 0x78, 0x02, 0x1a, 0x42, 0xa0,
	0xd6, 0x30, 0xea, 0x83, 0xf6, 0x10, 0x99, 0x7b, 0xdd, 0x30, 0xcf, 0x92, 0x84, 0x61, 0x59, 0x0c,
	0x31, 0xe8, 0xbe, 0x49, 0xb3, 0x8c, 0xe4, 0xa1, 0x48, 0xc3, 0x75, 0x9a, 0x30, 0xed, 0xbe, 0x51,
	0x1f, 0xb4, 0xec, 0xa3, 0xb2, 0x40, 0x9d, 0x97, 0x12, 0x13, 0xf5, 0x33, 0xdf, 0x65, 0x9b, 0x02,
	0x7d, 0x50, 0x8d, 0x3f, 0x38, 0x81, 0x0e, 0x50, 0x19, 0x8f, 0xf8, 0x9a, 0x69, 0xaa, 0xa1, 0x0c,
	0x3a, 0xc3, 0x27, 0x1f, 0x91, 0x72, 0xe3, 0x61, 0x20, 0x29, 0xb8, 0xa2, 0xc2, 0x67, 0xe0, 0x61,
	0x4e, 0x18, 0xfd, 0x71, 0xcd, 0x53, 0x9a, 0x09, 0x07, 0x9a, 0xd2, 0x81, 0xaf, 0xcb, 0x02, 0x3d,
	0xc0, 0xd7, 0x80, 0x34, 0xe2, 0x6e, 0x21, 0xbe, 0x9b, 0x42, 0x0d, 0x34, 0xe7, 0x39, 0x89, 0x38,
	0xcd, 0xb5, 0x03, 0xf1, 0x7a, 0xf0, 0x55, 0x0a, 0x5f, 0x82, 0xcf, 0x62, 0xc2, 0xc3, 0x39, 0xcd,
	0x18, 0xcf, 0xa3, 0x34, 0xe3, 0x4c, 0x6b, 0x19, 0xca, 0xa0, 0x3d, 0x34, 0x3f, 0x22, 0x58, 0x6a,
	0xb5, 0x09, 0x77, 0x6e, 0x58, 0xb8, 0x13, 0xdf, 0xc9, 0xe1, 0x63, 0xa0, 0x46, 0x73, 0x9e, 0x5e,
	0x10, 0x0d, 0x18, 0xca, 0xe0, 0x00, 0x57, 0x19, 0x84, 0xa0, 0xb1, 0x22, 0x3c, 0xd2, 0xda, 0x52,
	0x87, 0x8c, 0xe1, 0x6f, 0x0a, 0xe8, 0xb1, 0x5c, 0x88, 0xe0, 0x79, 0x1a, 0x6f, 0x55, 0xbf, 0xa2,
	0x79, 0xf8, 0x9a, 0xae, 0x19, 0xd1, 0x1e, 0xc8, 0x81, 0x22, 0x6f, 0x0b, 0x54, 0xfb, 0xbb, 0x40,
	0x47, 0x8b, 0x94, 0xbf, 0x5e, 0xc7, 0xe6, 0x9c, 0xae, 0xac, 0x6a, 0x74, 0xb7, 0x8f, 0x63, 0x96,
	0x2c, 0x2d, 0xfe, 0xf3, 0x39, 0x61, 0xa6, 0x9f, 0xf1, 0xb2, 0x40, 0x8f, 0x83, 0xdc, 0xb9, 0x75,
	0xe3, 0x29, 0xcd, 0xbf, 0x13, 0xf7, 0x6d, 0x0a, 0xb4, 0xbf, 0x19, 0xde, 0x0f, 0x41, 0x0b, 0x34,
	0x63, 0x4a, 0x97, 0x61, 0x9a, 0x68, 0x0f, 0xa5, 0xa4, 0x2f, 0xca, 0x02, 0xa9, 0x36, 0xa5, 0x4b,
	0x3f, 0xd9, 0x14, 0xe8, 0x0a, 0xc4, 0x57, 0x41, 0xff, 0x4f, 0x05, 0x3c, 0xda, 0x61, 0x15, 0xfc,
	0x1e, 0x80, 0x55, 0x9a, 0x85, 0xd1, 0x8a, 0xae, 0x33, 0x2e, 0x17, 0xa1, 0x65, 0x9b, 0x9f, 0xf6,
	0xf7, 0x70, 0x6b, 0x95, 0x66, 0x23, 0x79, 0x01, 0x9c, 0x83, 0xa6, 0x78, 0x85, 0xaf, 0x08, 0x91,
	0x1b, 0xd2, 0x1e, 0xf6, 0xcc, 0x2d, 0xc5, 0x14, 0x3b, 0x6d, 0x56, 0x3b, 0x6d, 0x3a, 0x34, 0xcd,
	0x6c, 0x4b, 0xb4, 0xf9, 0xfd, 0x1f, 0xf4, 0xcd, 0xff, 0x68, 0x23, 0x08, 0x58, 0x8d, 0x09, 0x3f,
	0x25, 0xe4, 0xdb, 0x3f, 0x14, 0xd0, 0x7d, 0x7f, 0x4e, 0x61, 0x1f, 0xe8, 0xc1, 0xf3, 0x33, 0x3c,
	0x0d, 0xbd, 0x17, 0xde, 0x64, 0x1a, 0x06, 0xd3, 0xd1, 0x74, 0x16, 0x84, 0xb3, 0x49, 0xf0, 0xdc,
	0x73, 0xfc, 0x53, 0xdf, 0x73, 0xbb, 0x35, 0xa8, 0x83, 0xc3, 0x1d, 0x35, 0xfe, 0xe4, 0xc5, 0x68,
	0xec, 0xbb, 0x5d, 0x05, 0x1a, 0xe0, 0xab, 0x1d, 0xb8, 0x33, 0x9a, 0x38, 0xde, 0x78, 0xec, 0xb9,
	0xdd, 0x7b, 0x7b, 0x6e, 0x18, 0xd9, 0x67, 0x78, 0xea, 0xb9, 0xdd, 0x3a, 0x3c, 0x02, 0xfd, 0x1d,
	0x38, 0xf6, 0x82, 0xd9, 0x78, 0x1a, 0xba, 0x9e, 0x33, 0x1e, 0x61, 0xcf, 0xed, 0x36, 0xec, 0x67,
	0x6f, 0x4b, 0x5d, 0x79, 0x57, 0xea, 0xca, 0xbf, 0xa5, 0xae, 0xfc, 0x72, 0xa9, 0xd7, 0xde, 0x5d,
	0xea, 0xb5, 0xbf, 0x2e, 0xf5, 0xda, 0x0f, 0xc7, 0xb7, 0xdc, 0x60, 0x0b, 0x72, 0x5c, 0x8d, 0xbd,
	0x88, 0xad, 0x9f, 0x6e, 0x7f, 0xfc, 0xa4, 0x31, 0xb1, 0x2a, 0x3f, 0x7f, 0x27, 0xff, 0x05, 0x00,
	0x00, 0xff, 0xff, 0x9a, 0x5e, 0xe3, 0x3e, 0x8e, 0x05, 0x00, 0x00,
}

func (m *SportEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SportEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SportEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BookId) > 0 {
		i -= len(m.BookId)
		copy(dAtA[i:], m.BookId)
		i = encodeVarintSportEvent(dAtA, i, uint64(len(m.BookId)))
		i--
		dAtA[i] = 0x6a
	}
	{
		size := m.SrContributionForHouse.Size()
		i -= size
		if _, err := m.SrContributionForHouse.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSportEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	if len(m.Meta) > 0 {
		i -= len(m.Meta)
		copy(dAtA[i:], m.Meta)
		i = encodeVarintSportEvent(dAtA, i, uint64(len(m.Meta)))
		i--
		dAtA[i] = 0x5a
	}
	if m.Active {
		i--
		if m.Active {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x50
	}
	if m.BetConstraints != nil {
		{
			size, err := m.BetConstraints.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSportEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintSportEvent(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x42
	}
	if m.ResolutionTS != 0 {
		i = encodeVarintSportEvent(dAtA, i, uint64(m.ResolutionTS))
		i--
		dAtA[i] = 0x38
	}
	if m.Status != 0 {
		i = encodeVarintSportEvent(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x30
	}
	if len(m.WinnerOddsUIDs) > 0 {
		for iNdEx := len(m.WinnerOddsUIDs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.WinnerOddsUIDs[iNdEx])
			copy(dAtA[i:], m.WinnerOddsUIDs[iNdEx])
			i = encodeVarintSportEvent(dAtA, i, uint64(len(m.WinnerOddsUIDs[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Odds) > 0 {
		for iNdEx := len(m.Odds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Odds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSportEvent(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.EndTS != 0 {
		i = encodeVarintSportEvent(dAtA, i, uint64(m.EndTS))
		i--
		dAtA[i] = 0x18
	}
	if m.StartTS != 0 {
		i = encodeVarintSportEvent(dAtA, i, uint64(m.StartTS))
		i--
		dAtA[i] = 0x10
	}
	if len(m.UID) > 0 {
		i -= len(m.UID)
		copy(dAtA[i:], m.UID)
		i = encodeVarintSportEvent(dAtA, i, uint64(len(m.UID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventBetConstraints) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventBetConstraints) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventBetConstraints) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.BetFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSportEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.MinAmount.Size()
		i -= size
		if _, err := m.MinAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSportEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}

func encodeVarintSportEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovSportEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SportEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UID)
	if l > 0 {
		n += 1 + l + sovSportEvent(uint64(l))
	}
	if m.StartTS != 0 {
		n += 1 + sovSportEvent(uint64(m.StartTS))
	}
	if m.EndTS != 0 {
		n += 1 + sovSportEvent(uint64(m.EndTS))
	}
	if len(m.Odds) > 0 {
		for _, e := range m.Odds {
			l = e.Size()
			n += 1 + l + sovSportEvent(uint64(l))
		}
	}
	if len(m.WinnerOddsUIDs) > 0 {
		for _, s := range m.WinnerOddsUIDs {
			l = len(s)
			n += 1 + l + sovSportEvent(uint64(l))
		}
	}
	if m.Status != 0 {
		n += 1 + sovSportEvent(uint64(m.Status))
	}
	if m.ResolutionTS != 0 {
		n += 1 + sovSportEvent(uint64(m.ResolutionTS))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovSportEvent(uint64(l))
	}
	if m.BetConstraints != nil {
		l = m.BetConstraints.Size()
		n += 1 + l + sovSportEvent(uint64(l))
	}
	if m.Active {
		n += 2
	}
	l = len(m.Meta)
	if l > 0 {
		n += 1 + l + sovSportEvent(uint64(l))
	}
	l = m.SrContributionForHouse.Size()
	n += 1 + l + sovSportEvent(uint64(l))
	l = len(m.BookId)
	if l > 0 {
		n += 1 + l + sovSportEvent(uint64(l))
	}
	return n
}

func (m *EventBetConstraints) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MinAmount.Size()
	n += 1 + l + sovSportEvent(uint64(l))
	l = m.BetFee.Size()
	n += 1 + l + sovSportEvent(uint64(l))
	return n
}

func sovSportEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSportEvent(x uint64) (n int) {
	return sovSportEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SportEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSportEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SportEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SportEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSportEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSportEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSportEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTS", wireType)
			}
			m.StartTS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSportEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTS |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTS", wireType)
			}
			m.EndTS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSportEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndTS |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Odds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSportEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSportEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSportEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Odds = append(m.Odds, &Odds{})
			if err := m.Odds[len(m.Odds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WinnerOddsUIDs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSportEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSportEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSportEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WinnerOddsUIDs = append(m.WinnerOddsUIDs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSportEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= SportEventStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResolutionTS", wireType)
			}
			m.ResolutionTS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSportEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResolutionTS |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSportEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSportEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSportEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BetConstraints", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSportEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSportEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSportEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BetConstraints == nil {
				m.BetConstraints = &EventBetConstraints{}
			}
			if err := m.BetConstraints.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Active", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSportEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Active = bool(v != 0)
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSportEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSportEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSportEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Meta = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrContributionForHouse", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSportEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSportEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSportEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SrContributionForHouse.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BookId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSportEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSportEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSportEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BookId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSportEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSportEvent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EventBetConstraints) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSportEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventBetConstraints: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventBetConstraints: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSportEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSportEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSportEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BetFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSportEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSportEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSportEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BetFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSportEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSportEvent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSportEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSportEvent
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSportEvent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSportEvent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthSportEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSportEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSportEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSportEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSportEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSportEvent = fmt.Errorf("proto: unexpected end of group")
)
