// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fanx/market/ticket.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// MarketAddTicketPayload indicates data of add market ticket
type MarketAddTicketPayload struct {
	// uid is the universal unique identifier of the market.
	UID string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid"`
	// start_ts is the start timestamp of the market.
	StartTS uint64 `protobuf:"varint,2,opt,name=start_ts,proto3" json:"start_ts"`
	// end_ts is the end timestamp of the market.
	EndTS uint64 `protobuf:"varint,3,opt,name=end_ts,proto3" json:"end_ts"`
	// odds is the list of odds of the market.
	Odds []*Odds `protobuf:"bytes,4,rep,name=odds,proto3" json:"odds,omitempty"`
	// status is the current status of the market.
	Status MarketStatus `protobuf:"varint,5,opt,name=status,proto3,enum=fanx.fanx.market.MarketStatus" json:"status,omitempty"`
	// creator is the address of the creator of the market.
	Creator string `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	// min_bet_amount is the minimum allowed bet amount for a market.
	MinBetAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=min_bet_amount,json=minBetAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_bet_amount"`
	// bet_fee is the fee that the bettor needs to pay to bet on the market.
	BetFee github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,8,opt,name=bet_fee,json=betFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"bet_fee"`
	// meta contains human-readable metadata of the market.
	Meta string `protobuf:"bytes,9,opt,name=meta,proto3" json:"meta,omitempty"`
	// sr_contribution_for_house is the amount of contribution by sr for the house
	SrContributionForHouse github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,10,opt,name=sr_contribution_for_house,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"sr_contribution_for_house"`
}

func (m *MarketAddTicketPayload) Reset()         { *m = MarketAddTicketPayload{} }
func (m *MarketAddTicketPayload) String() string { return proto.CompactTextString(m) }
func (*MarketAddTicketPayload) ProtoMessage()    {}
func (*MarketAddTicketPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dc46cd902954700, []int{0}
}
func (m *MarketAddTicketPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MarketAddTicketPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MarketAddTicketPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MarketAddTicketPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketAddTicketPayload.Merge(m, src)
}
func (m *MarketAddTicketPayload) XXX_Size() int {
	return m.Size()
}
func (m *MarketAddTicketPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketAddTicketPayload.DiscardUnknown(m)
}

var xxx_messageInfo_MarketAddTicketPayload proto.InternalMessageInfo

func (m *MarketAddTicketPayload) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *MarketAddTicketPayload) GetStartTS() uint64 {
	if m != nil {
		return m.StartTS
	}
	return 0
}

func (m *MarketAddTicketPayload) GetEndTS() uint64 {
	if m != nil {
		return m.EndTS
	}
	return 0
}

func (m *MarketAddTicketPayload) GetOdds() []*Odds {
	if m != nil {
		return m.Odds
	}
	return nil
}

func (m *MarketAddTicketPayload) GetStatus() MarketStatus {
	if m != nil {
		return m.Status
	}
	return MarketStatus_MARKET_STATUS_UNSPECIFIED
}

func (m *MarketAddTicketPayload) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MarketAddTicketPayload) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

// MarketUpdateTicketPayload indicates data of the market update ticket
type MarketUpdateTicketPayload struct {
	// uid is the uuid of the market
	UID string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid"`
	// start_ts is the start timestamp of the market
	StartTS uint64 `protobuf:"varint,2,opt,name=start_ts,proto3" json:"start_ts"`
	// end_ts is the end timestamp of the market
	EndTS uint64 `protobuf:"varint,3,opt,name=end_ts,proto3" json:"end_ts"`
	// min_bet_amount is the minimum allowed bet amount for a market.
	MinBetAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=min_bet_amount,json=minBetAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_bet_amount"`
	// bet_fee is the fee that the bettor needs to pay to bet on the market.
	BetFee github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=bet_fee,json=betFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"bet_fee"`
	// status is the status of the resolution.
	Status MarketStatus `protobuf:"varint,6,opt,name=status,proto3,enum=fanx.fanx.market.MarketStatus" json:"status,omitempty"`
}

func (m *MarketUpdateTicketPayload) Reset()         { *m = MarketUpdateTicketPayload{} }
func (m *MarketUpdateTicketPayload) String() string { return proto.CompactTextString(m) }
func (*MarketUpdateTicketPayload) ProtoMessage()    {}
func (*MarketUpdateTicketPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dc46cd902954700, []int{1}
}
func (m *MarketUpdateTicketPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MarketUpdateTicketPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MarketUpdateTicketPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MarketUpdateTicketPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketUpdateTicketPayload.Merge(m, src)
}
func (m *MarketUpdateTicketPayload) XXX_Size() int {
	return m.Size()
}
func (m *MarketUpdateTicketPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketUpdateTicketPayload.DiscardUnknown(m)
}

var xxx_messageInfo_MarketUpdateTicketPayload proto.InternalMessageInfo

func (m *MarketUpdateTicketPayload) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *MarketUpdateTicketPayload) GetStartTS() uint64 {
	if m != nil {
		return m.StartTS
	}
	return 0
}

func (m *MarketUpdateTicketPayload) GetEndTS() uint64 {
	if m != nil {
		return m.EndTS
	}
	return 0
}

func (m *MarketUpdateTicketPayload) GetStatus() MarketStatus {
	if m != nil {
		return m.Status
	}
	return MarketStatus_MARKET_STATUS_UNSPECIFIED
}

// MarketResolutionTicketPayload indicates data of the
// resolution of the market ticket.
type MarketResolutionTicketPayload struct {
	// uid is the universal unique identifier of the market.
	UID string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid"`
	// resolution_ts is the resolution timestamp of the market.
	ResolutionTS uint64 `protobuf:"varint,2,opt,name=resolution_ts,proto3" json:"resolution_ts"`
	// winner_odds_uids is the universal unique identifier list of the winner
	// odds.
	WinnerOddsUIDs []string `protobuf:"bytes,3,rep,name=winner_odds_uids,proto3" json:"winner_odds_uids"`
	// status is the status of the resolution.
	Status MarketStatus `protobuf:"varint,4,opt,name=status,proto3,enum=fanx.fanx.market.MarketStatus" json:"status,omitempty"`
}

func (m *MarketResolutionTicketPayload) Reset()         { *m = MarketResolutionTicketPayload{} }
func (m *MarketResolutionTicketPayload) String() string { return proto.CompactTextString(m) }
func (*MarketResolutionTicketPayload) ProtoMessage()    {}
func (*MarketResolutionTicketPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dc46cd902954700, []int{2}
}
func (m *MarketResolutionTicketPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MarketResolutionTicketPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MarketResolutionTicketPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MarketResolutionTicketPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketResolutionTicketPayload.Merge(m, src)
}
func (m *MarketResolutionTicketPayload) XXX_Size() int {
	return m.Size()
}
func (m *MarketResolutionTicketPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketResolutionTicketPayload.DiscardUnknown(m)
}

var xxx_messageInfo_MarketResolutionTicketPayload proto.InternalMessageInfo

func (m *MarketResolutionTicketPayload) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *MarketResolutionTicketPayload) GetResolutionTS() uint64 {
	if m != nil {
		return m.ResolutionTS
	}
	return 0
}

func (m *MarketResolutionTicketPayload) GetWinnerOddsUIDs() []string {
	if m != nil {
		return m.WinnerOddsUIDs
	}
	return nil
}

func (m *MarketResolutionTicketPayload) GetStatus() MarketStatus {
	if m != nil {
		return m.Status
	}
	return MarketStatus_MARKET_STATUS_UNSPECIFIED
}

func init() {
	proto.RegisterType((*MarketAddTicketPayload)(nil), "fanx.fanx.market.MarketAddTicketPayload")
	proto.RegisterType((*MarketUpdateTicketPayload)(nil), "fanx.fanx.market.MarketUpdateTicketPayload")
	proto.RegisterType((*MarketResolutionTicketPayload)(nil), "fanx.fanx.market.MarketResolutionTicketPayload")
}

func init() { proto.RegisterFile("fanx/market/ticket.proto", fileDescriptor_1dc46cd902954700) }

var fileDescriptor_1dc46cd902954700 = []byte{
	// 599 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0x41, 0x4f, 0xdb, 0x3e,
	0x14, 0x6f, 0x68, 0x68, 0xc1, 0x7f, 0xfe, 0xd5, 0xe4, 0x0d, 0x08, 0x4c, 0x8b, 0x3b, 0x26, 0xa1,
	0x4e, 0x13, 0x89, 0x04, 0xc7, 0x9d, 0x28, 0x0c, 0xc6, 0x61, 0xda, 0xe4, 0x82, 0x26, 0xed, 0x12,
	0xa5, 0xb5, 0x09, 0x51, 0x49, 0x8c, 0x6c, 0x47, 0x8c, 0x6f, 0xb1, 0xfb, 0x8e, 0xfb, 0x32, 0x1c,
	0x39, 0x4e, 0x3b, 0x58, 0x53, 0xb8, 0xf5, 0x2b, 0x4c, 0x93, 0x26, 0x3b, 0x59, 0x09, 0x62, 0x3d,
	0x74, 0xdb, 0x61, 0x97, 0xf8, 0xf9, 0xbd, 0xdf, 0xfb, 0xbd, 0xf8, 0xfd, 0x9e, 0x0d, 0x96, 0x45,
	0x44, 0xfd, 0x24, 0xe4, 0x43, 0x2a, 0x7d, 0x19, 0x0f, 0x86, 0x54, 0x7a, 0x67, 0x9c, 0x49, 0x06,
	0x17, 0x45, 0x44, 0x53, 0x2a, 0xcf, 0x19, 0x1f, 0x7a, 0x22, 0xa2, 0x5e, 0x81, 0x59, 0xad, 0xe2,
	0x8b, 0xa5, 0xc0, 0xaf, 0x2e, 0x56, 0x02, 0x8c, 0x10, 0x51, 0xba, 0x1f, 0x44, 0x2c, 0x62, 0xc6,
	0xf4, 0xb5, 0x55, 0x78, 0xd7, 0xbe, 0xdb, 0x60, 0xe9, 0x95, 0xc1, 0x6e, 0x13, 0x72, 0x68, 0xca,
	0xbe, 0x09, 0x2f, 0x4e, 0x59, 0x48, 0x60, 0x1b, 0xd4, 0xb3, 0x98, 0x38, 0x56, 0xdb, 0xea, 0xcc,
	0x77, 0x5b, 0xb9, 0x42, 0xf5, 0xa3, 0x83, 0xdd, 0x91, 0x42, 0xda, 0x8b, 0xf5, 0x07, 0x6e, 0x81,
	0x39, 0x21, 0x43, 0x2e, 0x03, 0x29, 0x9c, 0x99, 0xb6, 0xd5, 0xb1, 0xbb, 0xcb, 0xb9, 0x42, 0xcd,
	0x9e, 0xf6, 0x1d, 0xf6, 0x46, 0x0a, 0x8d, 0xc3, 0x78, 0x6c, 0xc1, 0x67, 0xa0, 0x41, 0x53, 0xa2,
	0x53, 0xea, 0x26, 0xe5, 0x7e, 0xae, 0xd0, 0xec, 0x8b, 0x94, 0x98, 0x84, 0x32, 0x84, 0xcb, 0x15,
	0xfa, 0xc0, 0xd6, 0x47, 0x70, 0xec, 0x76, 0xbd, 0xf3, 0xdf, 0xe6, 0x43, 0xef, 0x97, 0xad, 0xf0,
	0x5e, 0x13, 0x22, 0xb0, 0x01, 0xc2, 0xe7, 0xa0, 0x21, 0x64, 0x28, 0x33, 0xe1, 0xcc, 0xb6, 0xad,
	0x4e, 0x6b, 0xf3, 0xc9, 0x84, 0x94, 0xe2, 0xcc, 0x3d, 0x03, 0xc5, 0x65, 0x0a, 0x74, 0x40, 0x73,
	0xc0, 0x69, 0x28, 0x19, 0x77, 0x1a, 0xfa, 0xd4, 0xf8, 0xe7, 0x16, 0x1e, 0x82, 0x56, 0x12, 0xa7,
	0x41, 0x9f, 0xca, 0x20, 0x4c, 0x58, 0x96, 0x4a, 0xa7, 0x69, 0xda, 0xe2, 0x5d, 0x2a, 0x54, 0xfb,
	0xa2, 0xd0, 0x7a, 0x14, 0xcb, 0x93, 0xac, 0xef, 0x0d, 0x58, 0xe2, 0x0f, 0x98, 0x48, 0x98, 0x28,
	0x97, 0x0d, 0x41, 0x86, 0xbe, 0xbc, 0x38, 0xa3, 0xc2, 0x3b, 0x48, 0x25, 0x5e, 0x48, 0xe2, 0xb4,
	0x4b, 0xe5, 0xb6, 0xe1, 0x80, 0xfb, 0xa0, 0xa9, 0x19, 0x8f, 0x29, 0x75, 0xe6, 0x7e, 0x8b, 0xae,
	0xd1, 0xa7, 0x72, 0x8f, 0x52, 0x08, 0x81, 0x9d, 0x50, 0x19, 0x3a, 0xf3, 0xe6, 0xaf, 0x8d, 0x0d,
	0x3f, 0x59, 0x60, 0x45, 0xf0, 0x60, 0xc0, 0x52, 0xc9, 0xe3, 0x7e, 0x26, 0x63, 0x96, 0x06, 0xc7,
	0x8c, 0x07, 0x27, 0x2c, 0x13, 0xd4, 0x01, 0xa6, 0x1e, 0x9d, 0xae, 0x5e, 0xae, 0xd0, 0x52, 0x8f,
	0xef, 0x54, 0x18, 0xf7, 0x18, 0x7f, 0xa9, 0xf9, 0x46, 0x0a, 0x4d, 0x2e, 0x86, 0x27, 0x87, 0xd6,
	0xbe, 0xcd, 0x80, 0x95, 0x42, 0x8b, 0xa3, 0x33, 0x12, 0x4a, 0xfa, 0xef, 0x8d, 0xe0, 0x5d, 0xe9,
	0xed, 0xbf, 0x2b, 0xfd, 0xec, 0x1f, 0x49, 0x7f, 0x33, 0xf0, 0x8d, 0xa9, 0x07, 0x7e, 0xed, 0xe3,
	0x0c, 0x78, 0x54, 0x04, 0x30, 0x15, 0xec, 0xd4, 0x88, 0x33, 0xad, 0x02, 0xfb, 0xe0, 0x7f, 0x3e,
	0x4e, 0xbe, 0x91, 0xe1, 0x71, 0xae, 0xd0, 0x42, 0x85, 0x55, 0xb7, 0xf6, 0x36, 0x10, 0xdf, 0xde,
	0x42, 0x0c, 0xee, 0x9d, 0xc7, 0x69, 0x4a, 0x79, 0xa0, 0x6f, 0x72, 0x90, 0xc5, 0x44, 0xeb, 0x53,
	0xef, 0xcc, 0x77, 0xd7, 0x73, 0x85, 0x5a, 0x6f, 0x4d, 0x4c, 0x5f, 0xf5, 0xa3, 0x83, 0x5d, 0x31,
	0x52, 0xe8, 0x0e, 0x1a, 0xdf, 0xf1, 0x54, 0xba, 0x63, 0x4f, 0xdd, 0x9d, 0xee, 0xce, 0x65, 0xee,
	0x5a, 0x57, 0xb9, 0x6b, 0x7d, 0xcd, 0x5d, 0xeb, 0xc3, 0xb5, 0x5b, 0xbb, 0xba, 0x76, 0x6b, 0x9f,
	0xaf, 0xdd, 0xda, 0xbb, 0xa7, 0x15, 0x91, 0x44, 0x44, 0x37, 0x4a, 0x46, 0x6d, 0xfb, 0xef, 0xc7,
	0x8f, 0xb8, 0xd6, 0xaa, 0xdf, 0x30, 0xef, 0xec, 0xd6, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf2,
	0x0c, 0xf3, 0x9e, 0xdf, 0x05, 0x00, 0x00,
}

func (m *MarketAddTicketPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MarketAddTicketPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MarketAddTicketPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.SrContributionForHouse.Size()
		i -= size
		if _, err := m.SrContributionForHouse.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTicket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	if len(m.Meta) > 0 {
		i -= len(m.Meta)
		copy(dAtA[i:], m.Meta)
		i = encodeVarintTicket(dAtA, i, uint64(len(m.Meta)))
		i--
		dAtA[i] = 0x4a
	}
	{
		size := m.BetFee.Size()
		i -= size
		if _, err := m.BetFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTicket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.MinBetAmount.Size()
		i -= size
		if _, err := m.MinBetAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTicket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTicket(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x32
	}
	if m.Status != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Odds) > 0 {
		for iNdEx := len(m.Odds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Odds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTicket(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.EndTS != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.EndTS))
		i--
		dAtA[i] = 0x18
	}
	if m.StartTS != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.StartTS))
		i--
		dAtA[i] = 0x10
	}
	if len(m.UID) > 0 {
		i -= len(m.UID)
		copy(dAtA[i:], m.UID)
		i = encodeVarintTicket(dAtA, i, uint64(len(m.UID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MarketUpdateTicketPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MarketUpdateTicketPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MarketUpdateTicketPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.BetFee.Size()
		i -= size
		if _, err := m.BetFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTicket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.MinBetAmount.Size()
		i -= size
		if _, err := m.MinBetAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTicket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.EndTS != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.EndTS))
		i--
		dAtA[i] = 0x18
	}
	if m.StartTS != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.StartTS))
		i--
		dAtA[i] = 0x10
	}
	if len(m.UID) > 0 {
		i -= len(m.UID)
		copy(dAtA[i:], m.UID)
		i = encodeVarintTicket(dAtA, i, uint64(len(m.UID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MarketResolutionTicketPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MarketResolutionTicketPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MarketResolutionTicketPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if len(m.WinnerOddsUIDs) > 0 {
		for iNdEx := len(m.WinnerOddsUIDs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.WinnerOddsUIDs[iNdEx])
			copy(dAtA[i:], m.WinnerOddsUIDs[iNdEx])
			i = encodeVarintTicket(dAtA, i, uint64(len(m.WinnerOddsUIDs[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.ResolutionTS != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.ResolutionTS))
		i--
		dAtA[i] = 0x10
	}
	if len(m.UID) > 0 {
		i -= len(m.UID)
		copy(dAtA[i:], m.UID)
		i = encodeVarintTicket(dAtA, i, uint64(len(m.UID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTicket(dAtA []byte, offset int, v uint64) int {
	offset -= sovTicket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MarketAddTicketPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UID)
	if l > 0 {
		n += 1 + l + sovTicket(uint64(l))
	}
	if m.StartTS != 0 {
		n += 1 + sovTicket(uint64(m.StartTS))
	}
	if m.EndTS != 0 {
		n += 1 + sovTicket(uint64(m.EndTS))
	}
	if len(m.Odds) > 0 {
		for _, e := range m.Odds {
			l = e.Size()
			n += 1 + l + sovTicket(uint64(l))
		}
	}
	if m.Status != 0 {
		n += 1 + sovTicket(uint64(m.Status))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTicket(uint64(l))
	}
	l = m.MinBetAmount.Size()
	n += 1 + l + sovTicket(uint64(l))
	l = m.BetFee.Size()
	n += 1 + l + sovTicket(uint64(l))
	l = len(m.Meta)
	if l > 0 {
		n += 1 + l + sovTicket(uint64(l))
	}
	l = m.SrContributionForHouse.Size()
	n += 1 + l + sovTicket(uint64(l))
	return n
}

func (m *MarketUpdateTicketPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UID)
	if l > 0 {
		n += 1 + l + sovTicket(uint64(l))
	}
	if m.StartTS != 0 {
		n += 1 + sovTicket(uint64(m.StartTS))
	}
	if m.EndTS != 0 {
		n += 1 + sovTicket(uint64(m.EndTS))
	}
	l = m.MinBetAmount.Size()
	n += 1 + l + sovTicket(uint64(l))
	l = m.BetFee.Size()
	n += 1 + l + sovTicket(uint64(l))
	if m.Status != 0 {
		n += 1 + sovTicket(uint64(m.Status))
	}
	return n
}

func (m *MarketResolutionTicketPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UID)
	if l > 0 {
		n += 1 + l + sovTicket(uint64(l))
	}
	if m.ResolutionTS != 0 {
		n += 1 + sovTicket(uint64(m.ResolutionTS))
	}
	if len(m.WinnerOddsUIDs) > 0 {
		for _, s := range m.WinnerOddsUIDs {
			l = len(s)
			n += 1 + l + sovTicket(uint64(l))
		}
	}
	if m.Status != 0 {
		n += 1 + sovTicket(uint64(m.Status))
	}
	return n
}

func sovTicket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTicket(x uint64) (n int) {
	return sovTicket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MarketAddTicketPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTicket
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
			return fmt.Errorf("proto: MarketAddTicketPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MarketAddTicketPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
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
					return ErrIntOverflowTicket
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
					return ErrIntOverflowTicket
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
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= MarketStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinBetAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinBetAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BetFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BetFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Meta = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrContributionForHouse", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SrContributionForHouse.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTicket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTicket
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
func (m *MarketUpdateTicketPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTicket
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
			return fmt.Errorf("proto: MarketUpdateTicketPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MarketUpdateTicketPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
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
					return ErrIntOverflowTicket
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
					return ErrIntOverflowTicket
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
				return fmt.Errorf("proto: wrong wireType = %d for field MinBetAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinBetAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BetFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BetFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= MarketStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTicket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTicket
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
func (m *MarketResolutionTicketPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTicket
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
			return fmt.Errorf("proto: MarketResolutionTicketPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MarketResolutionTicketPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResolutionTS", wireType)
			}
			m.ResolutionTS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WinnerOddsUIDs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WinnerOddsUIDs = append(m.WinnerOddsUIDs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= MarketStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTicket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTicket
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
func skipTicket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTicket
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
					return 0, ErrIntOverflowTicket
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
					return 0, ErrIntOverflowTicket
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
				return 0, ErrInvalidLengthTicket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTicket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTicket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTicket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTicket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTicket = fmt.Errorf("proto: unexpected end of group")
)
