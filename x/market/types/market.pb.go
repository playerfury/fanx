// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fanx/market/market.proto

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

// MarketStatus is the market status enumeration
type MarketStatus int32

const (
	// unspecified market
	MarketStatus_MARKET_STATUS_UNSPECIFIED MarketStatus = 0
	// market is active
	MarketStatus_MARKET_STATUS_ACTIVE MarketStatus = 1
	// market is inactive
	MarketStatus_MARKET_STATUS_INACTIVE MarketStatus = 2
	// market is canceled
	MarketStatus_MARKET_STATUS_CANCELED MarketStatus = 3
	// market is aborted
	MarketStatus_MARKET_STATUS_ABORTED MarketStatus = 4
	// result of the market is declared
	MarketStatus_MARKET_STATUS_RESULT_DECLARED MarketStatus = 5
)

var MarketStatus_name = map[int32]string{
	0: "MARKET_STATUS_UNSPECIFIED",
	1: "MARKET_STATUS_ACTIVE",
	2: "MARKET_STATUS_INACTIVE",
	3: "MARKET_STATUS_CANCELED",
	4: "MARKET_STATUS_ABORTED",
	5: "MARKET_STATUS_RESULT_DECLARED",
}

var MarketStatus_value = map[string]int32{
	"MARKET_STATUS_UNSPECIFIED":     0,
	"MARKET_STATUS_ACTIVE":          1,
	"MARKET_STATUS_INACTIVE":        2,
	"MARKET_STATUS_CANCELED":        3,
	"MARKET_STATUS_ABORTED":         4,
	"MARKET_STATUS_RESULT_DECLARED": 5,
}

func (x MarketStatus) String() string {
	return proto.EnumName(MarketStatus_name, int32(x))
}

func (MarketStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_06d4443afc295240, []int{0}
}

// Market is the representation of the market to be stored in
// the market state.
type Market struct {
	// uid is the universal unique identifier of the market.
	UID string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid"`
	// start_ts is the start timestamp of the market.
	StartTS uint64 `protobuf:"varint,2,opt,name=start_ts,proto3" json:"start_ts"`
	// end_ts is the end timestamp of the market.
	EndTS uint64 `protobuf:"varint,3,opt,name=end_ts,proto3" json:"end_ts"`
	// odds is the list of odds of the market.
	Odds []*Odds `protobuf:"bytes,4,rep,name=odds,proto3" json:"odds,omitempty"`
	// winner_odds_uids is the list of winner odds universal unique identifiers.
	WinnerOddsUIDs []string `protobuf:"bytes,5,rep,name=winner_odds_uids,proto3" json:"winner_odds_uids"`
	// status is the current status of the market.
	Status MarketStatus `protobuf:"varint,6,opt,name=status,proto3,enum=fanx.market.MarketStatus" json:"status,omitempty"`
	// resolution_ts is the timestamp of the resolution of market.
	ResolutionTS uint64 `protobuf:"varint,7,opt,name=resolution_ts,proto3" json:"resolution_ts"`
	// creator is the address of the creator of market.
	Creator string `protobuf:"bytes,8,opt,name=creator,proto3" json:"creator,omitempty"`
	// bet_constraints holds the constraints of market to accept bets.
	BetConstraints *MarketBetConstraints `protobuf:"bytes,9,opt,name=bet_constraints,json=betConstraints,proto3" json:"bet_constraints,omitempty"`
	// meta contains human-readable metadata of the market.
	Meta string `protobuf:"bytes,10,opt,name=meta,proto3" json:"meta,omitempty"`
	// sr_contribution_for_house is the amount of contribution by sr for the house
	SrContributionForHouse github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,11,opt,name=sr_contribution_for_house,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"sr_contribution_for_house"`
	// book_uid is the unique identifier corresponding to the book
	BookUID string `protobuf:"bytes,12,opt,name=book_uid,proto3" json:"book_uid"`
}

func (m *Market) Reset()         { *m = Market{} }
func (m *Market) String() string { return proto.CompactTextString(m) }
func (*Market) ProtoMessage()    {}
func (*Market) Descriptor() ([]byte, []int) {
	return fileDescriptor_06d4443afc295240, []int{0}
}
func (m *Market) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Market) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Market.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Market) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Market.Merge(m, src)
}
func (m *Market) XXX_Size() int {
	return m.Size()
}
func (m *Market) XXX_DiscardUnknown() {
	xxx_messageInfo_Market.DiscardUnknown(m)
}

var xxx_messageInfo_Market proto.InternalMessageInfo

func (m *Market) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *Market) GetStartTS() uint64 {
	if m != nil {
		return m.StartTS
	}
	return 0
}

func (m *Market) GetEndTS() uint64 {
	if m != nil {
		return m.EndTS
	}
	return 0
}

func (m *Market) GetOdds() []*Odds {
	if m != nil {
		return m.Odds
	}
	return nil
}

func (m *Market) GetWinnerOddsUIDs() []string {
	if m != nil {
		return m.WinnerOddsUIDs
	}
	return nil
}

func (m *Market) GetStatus() MarketStatus {
	if m != nil {
		return m.Status
	}
	return MarketStatus_MARKET_STATUS_UNSPECIFIED
}

func (m *Market) GetResolutionTS() uint64 {
	if m != nil {
		return m.ResolutionTS
	}
	return 0
}

func (m *Market) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Market) GetBetConstraints() *MarketBetConstraints {
	if m != nil {
		return m.BetConstraints
	}
	return nil
}

func (m *Market) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

func (m *Market) GetBookUID() string {
	if m != nil {
		return m.BookUID
	}
	return ""
}

// MarketBetConstraints is the bet constrains type for the market
type MarketBetConstraints struct {
	// min_amount is the minimum allowed bet amount for a market.
	MinAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=min_amount,json=minAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_amount"`
	// bet_fee is the fee that the bettor needs to pay to bet on the market.
	BetFee github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=bet_fee,json=betFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"bet_fee"`
}

func (m *MarketBetConstraints) Reset()         { *m = MarketBetConstraints{} }
func (m *MarketBetConstraints) String() string { return proto.CompactTextString(m) }
func (*MarketBetConstraints) ProtoMessage()    {}
func (*MarketBetConstraints) Descriptor() ([]byte, []int) {
	return fileDescriptor_06d4443afc295240, []int{1}
}
func (m *MarketBetConstraints) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MarketBetConstraints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MarketBetConstraints.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MarketBetConstraints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketBetConstraints.Merge(m, src)
}
func (m *MarketBetConstraints) XXX_Size() int {
	return m.Size()
}
func (m *MarketBetConstraints) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketBetConstraints.DiscardUnknown(m)
}

var xxx_messageInfo_MarketBetConstraints proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("fanx.market.MarketStatus", MarketStatus_name, MarketStatus_value)
	proto.RegisterType((*Market)(nil), "fanx.market.Market")
	proto.RegisterType((*MarketBetConstraints)(nil), "fanx.market.MarketBetConstraints")
}

func init() { proto.RegisterFile("fanx/market/market.proto", fileDescriptor_06d4443afc295240) }

var fileDescriptor_06d4443afc295240 = []byte{
	// 681 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0x8e, 0x9b, 0x7f, 0xcd, 0xa6, 0xbf, 0xfc, 0xc2, 0x52, 0x8a, 0x53, 0xa9, 0x71, 0x5a, 0x89,
	0x2a, 0x80, 0x48, 0x44, 0xfb, 0x04, 0x89, 0xe3, 0x96, 0x40, 0xff, 0xa0, 0x75, 0x02, 0x12, 0x17,
	0xcb, 0x8e, 0x37, 0xad, 0x95, 0xda, 0x5b, 0xed, 0xae, 0x05, 0x7d, 0x0b, 0xde, 0x81, 0x3b, 0x6f,
	0xc0, 0xbd, 0xc7, 0x1e, 0x2b, 0x0e, 0x16, 0x72, 0x6f, 0x79, 0x0a, 0xb4, 0x1b, 0x13, 0x62, 0xda,
	0x1e, 0x7a, 0xc9, 0xce, 0x7c, 0xdf, 0x37, 0xb3, 0x33, 0xe3, 0xc9, 0x02, 0x75, 0x6c, 0x07, 0x5f,
	0xda, 0xbe, 0x4d, 0x27, 0x98, 0x27, 0x47, 0xeb, 0x9c, 0x12, 0x4e, 0x60, 0x59, 0x30, 0xad, 0x19,
	0xb4, 0xbe, 0x7a, 0x42, 0x4e, 0x88, 0xc4, 0xdb, 0xc2, 0x9a, 0x49, 0xd6, 0xd7, 0x16, 0x83, 0x89,
	0xeb, 0xb2, 0x19, 0xbe, 0x75, 0x9d, 0x07, 0x85, 0x43, 0x89, 0xc2, 0x06, 0xc8, 0x86, 0x9e, 0xab,
	0x2a, 0x0d, 0xa5, 0x59, 0xea, 0x56, 0xe2, 0x48, 0xcb, 0x0e, 0xfb, 0xbd, 0x69, 0xa4, 0x09, 0x14,
	0x89, 0x1f, 0xb8, 0x0b, 0x96, 0x19, 0xb7, 0x29, 0xb7, 0x38, 0x53, 0x97, 0x1a, 0x4a, 0x33, 0xd7,
	0x7d, 0x1a, 0x47, 0x5a, 0xd1, 0x14, 0xd8, 0xc0, 0x9c, 0x46, 0xda, 0x9c, 0x46, 0x73, 0x0b, 0xbe,
	0x04, 0x05, 0x1c, 0xb8, 0x22, 0x24, 0x2b, 0x43, 0x1e, 0xc7, 0x91, 0x96, 0x37, 0x02, 0x57, 0x06,
	0x24, 0x14, 0x4a, 0x4e, 0xf8, 0x0c, 0xe4, 0x44, 0x71, 0x6a, 0xae, 0x91, 0x6d, 0x96, 0x77, 0x1e,
	0xb5, 0x16, 0x1a, 0x6b, 0x1d, 0xbb, 0x2e, 0x43, 0x92, 0x86, 0x08, 0x54, 0x3f, 0x7b, 0x41, 0x80,
	0xa9, 0x25, 0x5c, 0x2b, 0xf4, 0x5c, 0xa6, 0xe6, 0x1b, 0xd9, 0x66, 0xa9, 0xbb, 0x1d, 0x47, 0x5a,
	0xe5, 0xa3, 0xe4, 0x84, 0x7e, 0xd8, 0xef, 0xb1, 0x69, 0xa4, 0xdd, 0x52, 0xa3, 0x5b, 0x08, 0x7c,
	0x0d, 0x0a, 0x8c, 0xdb, 0x3c, 0x64, 0x6a, 0xa1, 0xa1, 0x34, 0x2b, 0x3b, 0xb5, 0xd4, 0xe5, 0xb3,
	0x19, 0x99, 0x52, 0x80, 0x12, 0x21, 0xdc, 0x07, 0xff, 0x51, 0xcc, 0xc8, 0x59, 0xc8, 0x3d, 0x12,
	0x88, 0x0e, 0x8b, 0xb2, 0xc3, 0xcd, 0x38, 0xd2, 0x56, 0xd0, 0x9c, 0x90, 0x8d, 0xa6, 0x85, 0x28,
	0xed, 0x42, 0x15, 0x14, 0x47, 0x14, 0xdb, 0x9c, 0x50, 0x75, 0x59, 0x8c, 0x1f, 0xfd, 0x71, 0xe1,
	0x5b, 0xf0, 0xbf, 0x83, 0xb9, 0x35, 0x22, 0x01, 0xe3, 0xd4, 0xf6, 0x02, 0xce, 0xd4, 0x52, 0x43,
	0x69, 0x96, 0x77, 0x36, 0xef, 0x28, 0xaf, 0x8b, 0xb9, 0xfe, 0x57, 0x88, 0x2a, 0x4e, 0xca, 0x87,
	0x10, 0xe4, 0x7c, 0xcc, 0x6d, 0x15, 0xc8, 0x2b, 0xa4, 0x0d, 0xbf, 0x29, 0xa0, 0xc6, 0xa8, 0xc8,
	0xcf, 0xa9, 0xe7, 0xcc, 0x0a, 0x1a, 0x13, 0x6a, 0x9d, 0x92, 0x90, 0x61, 0xb5, 0x2c, 0x77, 0x01,
	0x5f, 0x46, 0x5a, 0xe6, 0x67, 0xa4, 0x6d, 0x9f, 0x78, 0xfc, 0x34, 0x74, 0x5a, 0x23, 0xe2, 0xb7,
	0x47, 0x84, 0xf9, 0x84, 0x25, 0xc7, 0x2b, 0xe6, 0x4e, 0xda, 0xfc, 0xe2, 0x1c, 0xb3, 0x56, 0x3f,
	0xe0, 0x71, 0xa4, 0xad, 0x99, 0x54, 0x5f, 0xc8, 0xb8, 0x47, 0xe8, 0x1b, 0x91, 0x6f, 0x1a, 0x69,
	0xf7, 0x5f, 0x86, 0xee, 0xa7, 0xc4, 0xe2, 0x39, 0x84, 0x4c, 0xc4, 0x87, 0x52, 0x57, 0x64, 0x4d,
	0x72, 0xf1, 0xba, 0x84, 0x4c, 0x66, 0x3b, 0x3a, 0xa7, 0xd1, 0xdc, 0xda, 0xfa, 0xae, 0x80, 0xd5,
	0xbb, 0xe6, 0x02, 0x0f, 0x01, 0xf0, 0xbd, 0xc0, 0xb2, 0x7d, 0x12, 0x06, 0x3c, 0xd9, 0xf7, 0xd6,
	0xc3, 0x7a, 0x44, 0x25, 0xdf, 0x0b, 0x3a, 0x32, 0x01, 0xdc, 0x07, 0x45, 0xf1, 0x89, 0xc6, 0x18,
	0xcb, 0x3f, 0xc5, 0xc3, 0x73, 0x15, 0x1c, 0xcc, 0xf7, 0x30, 0x7e, 0xf1, 0x43, 0x01, 0x2b, 0x8b,
	0x7b, 0x06, 0x37, 0x40, 0xed, 0xb0, 0x83, 0xde, 0x19, 0x03, 0xcb, 0x1c, 0x74, 0x06, 0x43, 0xd3,
	0x1a, 0x1e, 0x99, 0xef, 0x0d, 0xbd, 0xbf, 0xd7, 0x37, 0x7a, 0xd5, 0x0c, 0x54, 0xc1, 0x6a, 0x9a,
	0xee, 0xe8, 0x83, 0xfe, 0x07, 0xa3, 0xaa, 0xc0, 0x75, 0xb0, 0x96, 0x66, 0xfa, 0x47, 0x09, 0xb7,
	0x74, 0x9b, 0xd3, 0x3b, 0x47, 0xba, 0x71, 0x60, 0xf4, 0xaa, 0x59, 0x58, 0x03, 0x4f, 0xfe, 0xc9,
	0xd8, 0x3d, 0x46, 0x03, 0xa3, 0x57, 0xcd, 0xc1, 0x4d, 0xb0, 0x91, 0xa6, 0x90, 0x61, 0x0e, 0x0f,
	0x06, 0x56, 0xcf, 0xd0, 0x0f, 0x3a, 0xc8, 0xe8, 0x55, 0xf3, 0x5d, 0xfd, 0x32, 0xae, 0x2b, 0x57,
	0x71, 0x5d, 0xf9, 0x15, 0xd7, 0x95, 0xaf, 0x37, 0xf5, 0xcc, 0xd5, 0x4d, 0x3d, 0x73, 0x7d, 0x53,
	0xcf, 0x7c, 0x7a, 0xbe, 0x30, 0x89, 0xf3, 0x33, 0xfb, 0x02, 0xd3, 0x71, 0x48, 0x2f, 0xda, 0xf2,
	0x4d, 0x9a, 0xbf, 0x4a, 0x72, 0x20, 0x4e, 0x41, 0xbe, 0x4b, 0xbb, 0xbf, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x9d, 0x65, 0x24, 0x87, 0xee, 0x04, 0x00, 0x00,
}

func (m *Market) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Market) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Market) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BookUID) > 0 {
		i -= len(m.BookUID)
		copy(dAtA[i:], m.BookUID)
		i = encodeVarintMarket(dAtA, i, uint64(len(m.BookUID)))
		i--
		dAtA[i] = 0x62
	}
	{
		size := m.SrContributionForHouse.Size()
		i -= size
		if _, err := m.SrContributionForHouse.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMarket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	if len(m.Meta) > 0 {
		i -= len(m.Meta)
		copy(dAtA[i:], m.Meta)
		i = encodeVarintMarket(dAtA, i, uint64(len(m.Meta)))
		i--
		dAtA[i] = 0x52
	}
	if m.BetConstraints != nil {
		{
			size, err := m.BetConstraints.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMarket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintMarket(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x42
	}
	if m.ResolutionTS != 0 {
		i = encodeVarintMarket(dAtA, i, uint64(m.ResolutionTS))
		i--
		dAtA[i] = 0x38
	}
	if m.Status != 0 {
		i = encodeVarintMarket(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x30
	}
	if len(m.WinnerOddsUIDs) > 0 {
		for iNdEx := len(m.WinnerOddsUIDs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.WinnerOddsUIDs[iNdEx])
			copy(dAtA[i:], m.WinnerOddsUIDs[iNdEx])
			i = encodeVarintMarket(dAtA, i, uint64(len(m.WinnerOddsUIDs[iNdEx])))
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
				i = encodeVarintMarket(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.EndTS != 0 {
		i = encodeVarintMarket(dAtA, i, uint64(m.EndTS))
		i--
		dAtA[i] = 0x18
	}
	if m.StartTS != 0 {
		i = encodeVarintMarket(dAtA, i, uint64(m.StartTS))
		i--
		dAtA[i] = 0x10
	}
	if len(m.UID) > 0 {
		i -= len(m.UID)
		copy(dAtA[i:], m.UID)
		i = encodeVarintMarket(dAtA, i, uint64(len(m.UID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MarketBetConstraints) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MarketBetConstraints) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MarketBetConstraints) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.BetFee.Size()
		i -= size
		if _, err := m.BetFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMarket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.MinAmount.Size()
		i -= size
		if _, err := m.MinAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMarket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintMarket(dAtA []byte, offset int, v uint64) int {
	offset -= sovMarket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Market) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UID)
	if l > 0 {
		n += 1 + l + sovMarket(uint64(l))
	}
	if m.StartTS != 0 {
		n += 1 + sovMarket(uint64(m.StartTS))
	}
	if m.EndTS != 0 {
		n += 1 + sovMarket(uint64(m.EndTS))
	}
	if len(m.Odds) > 0 {
		for _, e := range m.Odds {
			l = e.Size()
			n += 1 + l + sovMarket(uint64(l))
		}
	}
	if len(m.WinnerOddsUIDs) > 0 {
		for _, s := range m.WinnerOddsUIDs {
			l = len(s)
			n += 1 + l + sovMarket(uint64(l))
		}
	}
	if m.Status != 0 {
		n += 1 + sovMarket(uint64(m.Status))
	}
	if m.ResolutionTS != 0 {
		n += 1 + sovMarket(uint64(m.ResolutionTS))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovMarket(uint64(l))
	}
	if m.BetConstraints != nil {
		l = m.BetConstraints.Size()
		n += 1 + l + sovMarket(uint64(l))
	}
	l = len(m.Meta)
	if l > 0 {
		n += 1 + l + sovMarket(uint64(l))
	}
	l = m.SrContributionForHouse.Size()
	n += 1 + l + sovMarket(uint64(l))
	l = len(m.BookUID)
	if l > 0 {
		n += 1 + l + sovMarket(uint64(l))
	}
	return n
}

func (m *MarketBetConstraints) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MinAmount.Size()
	n += 1 + l + sovMarket(uint64(l))
	l = m.BetFee.Size()
	n += 1 + l + sovMarket(uint64(l))
	return n
}

func sovMarket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMarket(x uint64) (n int) {
	return sovMarket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Market) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMarket
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
			return fmt.Errorf("proto: Market: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Market: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
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
				return ErrInvalidLengthMarket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarket
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
					return ErrIntOverflowMarket
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
					return ErrIntOverflowMarket
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
					return ErrIntOverflowMarket
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
				return ErrInvalidLengthMarket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMarket
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
					return ErrIntOverflowMarket
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
				return ErrInvalidLengthMarket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarket
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
					return ErrIntOverflowMarket
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
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResolutionTS", wireType)
			}
			m.ResolutionTS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
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
					return ErrIntOverflowMarket
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
				return ErrInvalidLengthMarket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarket
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
					return ErrIntOverflowMarket
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
				return ErrInvalidLengthMarket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMarket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BetConstraints == nil {
				m.BetConstraints = &MarketBetConstraints{}
			}
			if err := m.BetConstraints.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
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
				return ErrInvalidLengthMarket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Meta = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrContributionForHouse", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
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
				return ErrInvalidLengthMarket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SrContributionForHouse.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BookUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
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
				return ErrInvalidLengthMarket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BookUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMarket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMarket
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
func (m *MarketBetConstraints) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMarket
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
			return fmt.Errorf("proto: MarketBetConstraints: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MarketBetConstraints: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
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
				return ErrInvalidLengthMarket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BetFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
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
				return ErrInvalidLengthMarket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarket
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
			skippy, err := skipMarket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMarket
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
func skipMarket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMarket
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
					return 0, ErrIntOverflowMarket
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
					return 0, ErrIntOverflowMarket
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
				return 0, ErrInvalidLengthMarket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMarket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMarket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMarket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMarket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMarket = fmt.Errorf("proto: unexpected end of group")
)
