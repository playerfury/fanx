// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fanx/strategicreserve/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the strategicreserve module's genesis state.
type GenesisState struct {
	// params defines all the parameters related to order book.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// order_book_list defines the order books available at genesis.
	OrderBookList []OrderBook `protobuf:"bytes,2,rep,name=order_book_list,json=orderBookList,proto3" json:"order_book_list"`
	// order_book_participation_list defines the book participations available at
	// genesis.
	OrderBookParticipationList []OrderBookParticipation `protobuf:"bytes,3,rep,name=order_book_participation_list,json=orderBookParticipationList,proto3" json:"order_book_participation_list"`
	// order_book_exposure_list defines the order book exposures available at genesis.
	OrderBookExposureList []OrderBookOddsExposure `protobuf:"bytes,4,rep,name=order_book_exposure_list,json=orderBookExposureList,proto3" json:"order_book_exposure_list"`
	// participation_exposure_list defines the participation exposures available
	// at genesis.
	ParticipationExposureList []ParticipationExposure `protobuf:"bytes,5,rep,name=participation_exposure_list,json=participationExposureList,proto3" json:"participation_exposure_list"`
	// participation_exposure_by_index_list defines the participation exposures by
	// the indices available at genesis.
	ParticipationExposureByIndexList []ParticipationExposure `protobuf:"bytes,6,rep,name=participation_exposure_by_index_list,json=participationExposureByIndexList,proto3" json:"participation_exposure_by_index_list"`
	// historical_participation_exposure_list defines the historical participation
	// exposures available at genesis.
	HistoricalParticipationExposureList []ParticipationExposure `protobuf:"bytes,7,rep,name=historical_participation_exposure_list,json=historicalParticipationExposureList,proto3" json:"historical_participation_exposure_list"`
	// historical_participation_exposure_list defines the participation bet pair
	// exposures available at genesis.
	ParticipationBetPairExposureList []ParticipationBetPair `protobuf:"bytes,8,rep,name=participation_bet_pair_exposure_list,json=participationBetPairExposureList,proto3" json:"participation_bet_pair_exposure_list"`
	// payout_lock defines the current locks for the payouts.
	PayoutLock [][]byte `protobuf:"bytes,9,rep,name=payout_lock,json=payoutLock,proto3" json:"payout_lock,omitempty"`
	// stats is the statistics of the order book.
	Stats OrderBookStats `protobuf:"bytes,10,opt,name=stats,proto3" json:"stats"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_708845e0eb254a47, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetOrderBookList() []OrderBook {
	if m != nil {
		return m.OrderBookList
	}
	return nil
}

func (m *GenesisState) GetOrderBookParticipationList() []OrderBookParticipation {
	if m != nil {
		return m.OrderBookParticipationList
	}
	return nil
}

func (m *GenesisState) GetOrderBookExposureList() []OrderBookOddsExposure {
	if m != nil {
		return m.OrderBookExposureList
	}
	return nil
}

func (m *GenesisState) GetParticipationExposureList() []ParticipationExposure {
	if m != nil {
		return m.ParticipationExposureList
	}
	return nil
}

func (m *GenesisState) GetParticipationExposureByIndexList() []ParticipationExposure {
	if m != nil {
		return m.ParticipationExposureByIndexList
	}
	return nil
}

func (m *GenesisState) GetHistoricalParticipationExposureList() []ParticipationExposure {
	if m != nil {
		return m.HistoricalParticipationExposureList
	}
	return nil
}

func (m *GenesisState) GetParticipationBetPairExposureList() []ParticipationBetPair {
	if m != nil {
		return m.ParticipationBetPairExposureList
	}
	return nil
}

func (m *GenesisState) GetPayoutLock() [][]byte {
	if m != nil {
		return m.PayoutLock
	}
	return nil
}

func (m *GenesisState) GetStats() OrderBookStats {
	if m != nil {
		return m.Stats
	}
	return OrderBookStats{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "fanx.strategicreserve.GenesisState")
}

func init() {
	proto.RegisterFile("fanx/strategicreserve/genesis.proto", fileDescriptor_708845e0eb254a47)
}

var fileDescriptor_708845e0eb254a47 = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xc1, 0x8b, 0xd3, 0x40,
	0x14, 0xc6, 0x1b, 0xbb, 0x5b, 0xd7, 0xe9, 0x8a, 0x10, 0x14, 0x62, 0xc5, 0x6c, 0xdc, 0x5d, 0xb5,
	0x08, 0x26, 0xb0, 0x8b, 0x7a, 0x0f, 0x2c, 0x22, 0xae, 0x6c, 0xd1, 0x8b, 0x78, 0x09, 0x93, 0x74,
	0xc8, 0x0e, 0xe9, 0xf6, 0x85, 0x99, 0x57, 0x6d, 0xc1, 0x83, 0x77, 0x3d, 0xe8, 0x7f, 0xb5, 0x07,
	0x0f, 0x7b, 0xf4, 0x24, 0xd2, 0xfe, 0x23, 0x92, 0x99, 0x89, 0x36, 0xdd, 0xd4, 0xad, 0xb0, 0xb7,
	0xb4, 0xef, 0x9b, 0xef, 0xf7, 0xcd, 0x47, 0xf2, 0xc8, 0xb6, 0x4c, 0x59, 0x20, 0x51, 0x50, 0x64,
	0x29, 0x4f, 0x04, 0x93, 0x4c, 0xbc, 0x67, 0x41, 0xca, 0x86, 0x4c, 0x72, 0xe9, 0xe7, 0x02, 0x10,
	0xec, 0x2d, 0x59, 0xfc, 0xc6, 0x0f, 0x20, 0x32, 0x5f, 0xa6, 0xcc, 0x5f, 0x94, 0x77, 0x6e, 0xa6,
	0x90, 0x82, 0xd2, 0x06, 0xc5, 0x93, 0x3e, 0xd6, 0xb9, 0x57, 0x6b, 0x9d, 0x53, 0x41, 0x4f, 0x8c,
	0x73, 0x67, 0xb7, 0x56, 0x02, 0xa2, 0xcf, 0x44, 0x0c, 0x90, 0x19, 0x55, 0x77, 0x99, 0x11, 0xf2,
	0x84, 0xe7, 0x14, 0x39, 0x0c, 0x8d, 0xd2, 0xab, 0x55, 0x4a, 0xa4, 0x58, 0x12, 0x77, 0x6a, 0x15,
	0x6c, 0x9c, 0x83, 0x1c, 0x09, 0xa6, 0x45, 0xdb, 0xdf, 0x37, 0xc8, 0xe6, 0x73, 0x5d, 0xc1, 0x1b,
	0xa4, 0xc8, 0xec, 0x03, 0xd2, 0xd2, 0xb9, 0x1d, 0xcb, 0xb3, 0xba, 0xed, 0xbd, 0x87, 0xfe, 0x05,
	0x95, 0xf8, 0x3d, 0x25, 0x0f, 0xd7, 0x4e, 0x7f, 0x6e, 0x35, 0x5e, 0x9b, 0xc3, 0xf6, 0x5b, 0x72,
	0x43, 0xdd, 0x2d, 0x2a, 0x2e, 0x17, 0x0d, 0xb8, 0x44, 0xe7, 0x8a, 0xd7, 0xec, 0xb6, 0xf7, 0x1e,
	0x5d, 0xe8, 0x77, 0x54, 0x9c, 0x0b, 0x01, 0x32, 0x63, 0x79, 0x1d, 0xca, 0x3f, 0x0e, 0xb9, 0x44,
	0xfb, 0x93, 0x45, 0xee, 0xce, 0x59, 0x57, 0xba, 0xd1, 0xa0, 0xa6, 0x02, 0x3d, 0x5b, 0x1d, 0xd4,
	0x9b, 0xf7, 0x30, 0xd4, 0x0e, 0xd4, 0x4e, 0x55, 0x84, 0x11, 0x71, 0xe6, 0x12, 0x94, 0x8d, 0x6a,
	0xf8, 0x9a, 0x82, 0x3f, 0x5d, 0x1d, 0x7e, 0xd4, 0xef, 0xcb, 0x03, 0x63, 0x61, 0xd8, 0xb7, 0xfe,
	0xb0, 0xcb, 0x81, 0xc2, 0x7e, 0x24, 0x77, 0xaa, 0xb7, 0xad, 0x92, 0xd7, 0x57, 0x24, 0x57, 0xee,
	0xb3, 0x40, 0xbe, 0x9d, 0xd7, 0x0d, 0x15, 0xfd, 0x8b, 0x45, 0x76, 0x97, 0xe0, 0xe3, 0x49, 0xc4,
	0x87, 0x7d, 0x36, 0xd6, 0x39, 0x5a, 0x97, 0x90, 0xc3, 0xab, 0xcd, 0x11, 0x4e, 0x5e, 0x14, 0x18,
	0x15, 0xe7, 0x9b, 0x45, 0x1e, 0x1c, 0x73, 0x89, 0x20, 0x78, 0x42, 0x07, 0xd1, 0xbf, 0x8a, 0xb9,
	0x7a, 0x09, 0x81, 0x76, 0xfe, 0xb2, 0x7a, 0x4b, 0x2b, 0xfa, 0x7c, 0xae, 0xa2, 0x98, 0x61, 0x94,
	0x53, 0x2e, 0x16, 0x12, 0x6d, 0xa8, 0x44, 0x4f, 0xfe, 0x2f, 0x51, 0xc8, 0xb0, 0x47, 0xb9, 0xa8,
	0x6d, 0xc8, 0xcc, 0x2a, 0x69, 0xee, 0x93, 0x76, 0x4e, 0x27, 0x30, 0xc2, 0x68, 0x00, 0x49, 0xe6,
	0x5c, 0xf3, 0x9a, 0xdd, 0x4d, 0x73, 0x98, 0xe8, 0xc1, 0x21, 0x24, 0x99, 0xfd, 0x92, 0xac, 0xab,
	0xad, 0xe1, 0x10, 0xf5, 0xbd, 0x07, 0xab, 0xbf, 0xb9, 0xc5, 0xc2, 0x28, 0xbf, 0x7b, 0xed, 0x11,
	0xbe, 0x3a, 0x9d, 0xba, 0xd6, 0xd9, 0xd4, 0xb5, 0x7e, 0x4d, 0x5d, 0xeb, 0xeb, 0xcc, 0x6d, 0x9c,
	0xcd, 0xdc, 0xc6, 0x8f, 0x99, 0xdb, 0x78, 0xb7, 0x9f, 0x72, 0x3c, 0x1e, 0xc5, 0x7e, 0x02, 0x27,
	0x81, 0x4c, 0xd9, 0x63, 0x83, 0x28, 0x9e, 0x83, 0xf1, 0xf9, 0x35, 0x85, 0x93, 0x9c, 0xc9, 0xb8,
	0xa5, 0x96, 0xd4, 0xfe, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x14, 0x3b, 0xed, 0xab, 0xbb, 0x05,
	0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Stats.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	if len(m.PayoutLock) > 0 {
		for iNdEx := len(m.PayoutLock) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PayoutLock[iNdEx])
			copy(dAtA[i:], m.PayoutLock[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.PayoutLock[iNdEx])))
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.ParticipationBetPairExposureList) > 0 {
		for iNdEx := len(m.ParticipationBetPairExposureList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ParticipationBetPairExposureList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.HistoricalParticipationExposureList) > 0 {
		for iNdEx := len(m.HistoricalParticipationExposureList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.HistoricalParticipationExposureList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.ParticipationExposureByIndexList) > 0 {
		for iNdEx := len(m.ParticipationExposureByIndexList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ParticipationExposureByIndexList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.ParticipationExposureList) > 0 {
		for iNdEx := len(m.ParticipationExposureList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ParticipationExposureList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.OrderBookExposureList) > 0 {
		for iNdEx := len(m.OrderBookExposureList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OrderBookExposureList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.OrderBookParticipationList) > 0 {
		for iNdEx := len(m.OrderBookParticipationList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OrderBookParticipationList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.OrderBookList) > 0 {
		for iNdEx := len(m.OrderBookList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OrderBookList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.OrderBookList) > 0 {
		for _, e := range m.OrderBookList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.OrderBookParticipationList) > 0 {
		for _, e := range m.OrderBookParticipationList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.OrderBookExposureList) > 0 {
		for _, e := range m.OrderBookExposureList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ParticipationExposureList) > 0 {
		for _, e := range m.ParticipationExposureList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ParticipationExposureByIndexList) > 0 {
		for _, e := range m.ParticipationExposureByIndexList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.HistoricalParticipationExposureList) > 0 {
		for _, e := range m.HistoricalParticipationExposureList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ParticipationBetPairExposureList) > 0 {
		for _, e := range m.ParticipationBetPairExposureList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PayoutLock) > 0 {
		for _, b := range m.PayoutLock {
			l = len(b)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Stats.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderBookList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderBookList = append(m.OrderBookList, OrderBook{})
			if err := m.OrderBookList[len(m.OrderBookList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderBookParticipationList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderBookParticipationList = append(m.OrderBookParticipationList, OrderBookParticipation{})
			if err := m.OrderBookParticipationList[len(m.OrderBookParticipationList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderBookExposureList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderBookExposureList = append(m.OrderBookExposureList, OrderBookOddsExposure{})
			if err := m.OrderBookExposureList[len(m.OrderBookExposureList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParticipationExposureList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ParticipationExposureList = append(m.ParticipationExposureList, ParticipationExposure{})
			if err := m.ParticipationExposureList[len(m.ParticipationExposureList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParticipationExposureByIndexList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ParticipationExposureByIndexList = append(m.ParticipationExposureByIndexList, ParticipationExposure{})
			if err := m.ParticipationExposureByIndexList[len(m.ParticipationExposureByIndexList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HistoricalParticipationExposureList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HistoricalParticipationExposureList = append(m.HistoricalParticipationExposureList, ParticipationExposure{})
			if err := m.HistoricalParticipationExposureList[len(m.HistoricalParticipationExposureList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParticipationBetPairExposureList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ParticipationBetPairExposureList = append(m.ParticipationBetPairExposureList, ParticipationBetPair{})
			if err := m.ParticipationBetPairExposureList[len(m.ParticipationBetPairExposureList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PayoutLock", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PayoutLock = append(m.PayoutLock, make([]byte, postIndex-iNdEx))
			copy(m.PayoutLock[len(m.PayoutLock)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Stats.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
