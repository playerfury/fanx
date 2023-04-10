// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fanx/bet/bet_odds.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// BetOdds is the type to store odds of a market.
type BetOdds struct {
	// uid is universal unique identifier of odds.
	// Required | Unique | uuid-v4 or code
	UID string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid"`
	// market_uid is the parent, used for fast retrieving.
	// Required | NonUnique | -
	MarketUID string `protobuf:"bytes,2,opt,name=market_uid,proto3" json:"market_uid"`
	// value of the odds in corresponding odds type proposed in bet placement
	// message. Required | NonUnique | "1.286" or "2/7" or "+500"
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	// max_loss_multiplier is the factor for calculating max loss for given odds
	MaxLossMultiplier github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=max_loss_multiplier,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"max_loss_multiplier"`
}

func (m *BetOdds) Reset()         { *m = BetOdds{} }
func (m *BetOdds) String() string { return proto.CompactTextString(m) }
func (*BetOdds) ProtoMessage()    {}
func (*BetOdds) Descriptor() ([]byte, []int) {
	return fileDescriptor_99fce9bc4226027a, []int{0}
}
func (m *BetOdds) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BetOdds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BetOdds.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BetOdds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BetOdds.Merge(m, src)
}
func (m *BetOdds) XXX_Size() int {
	return m.Size()
}
func (m *BetOdds) XXX_DiscardUnknown() {
	xxx_messageInfo_BetOdds.DiscardUnknown(m)
}

var xxx_messageInfo_BetOdds proto.InternalMessageInfo

func (m *BetOdds) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *BetOdds) GetMarketUID() string {
	if m != nil {
		return m.MarketUID
	}
	return ""
}

func (m *BetOdds) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*BetOdds)(nil), "fanx.bet.BetOdds")
}

func init() { proto.RegisterFile("fanx/bet/bet_odds.proto", fileDescriptor_99fce9bc4226027a) }

var fileDescriptor_99fce9bc4226027a = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0x4b, 0xcc, 0xab,
	0xd0, 0x4f, 0x4a, 0x2d, 0x01, 0xe1, 0xf8, 0xfc, 0x94, 0x94, 0x62, 0xbd, 0x82, 0xa2, 0xfc, 0x92,
	0x7c, 0x21, 0x0e, 0x90, 0x84, 0x5e, 0x52, 0x6a, 0x89, 0x94, 0x48, 0x7a, 0x7e, 0x7a, 0x3e, 0x58,
	0x50, 0x1f, 0xc4, 0x82, 0xc8, 0x2b, 0xf5, 0x32, 0x71, 0xb1, 0x3b, 0xa5, 0x96, 0xf8, 0xa7, 0xa4,
	0x14, 0x0b, 0x29, 0x70, 0x31, 0x97, 0x66, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x3a, 0xf1,
	0x3d, 0xba, 0x27, 0xcf, 0x1c, 0xea, 0xe9, 0xf2, 0xea, 0x9e, 0x3c, 0x48, 0x34, 0x08, 0x44, 0x08,
	0x59, 0x73, 0x71, 0xe5, 0x26, 0x16, 0x65, 0xa7, 0x96, 0xc4, 0x83, 0x14, 0x32, 0x81, 0x15, 0x4a,
	0x3f, 0xba, 0x27, 0xcf, 0xe9, 0x0b, 0x16, 0x85, 0x28, 0x47, 0x52, 0x12, 0x84, 0xc4, 0x16, 0x12,
	0xe1, 0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60, 0x06, 0xe9, 0x0b, 0x82, 0x70, 0x84, 0x7a,
	0x18, 0xb9, 0x84, 0x73, 0x13, 0x2b, 0xe2, 0x73, 0xf2, 0x8b, 0x8b, 0xe3, 0x73, 0x4b, 0x73, 0x4a,
	0x32, 0x0b, 0x72, 0x32, 0x53, 0x8b, 0x24, 0x58, 0xc0, 0x86, 0x47, 0x9d, 0xb8, 0x27, 0xcf, 0x70,
	0xeb, 0x9e, 0xbc, 0x5a, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x72,
	0x7e, 0x71, 0x6e, 0x7e, 0x31, 0x94, 0xd2, 0x2d, 0x4e, 0xc9, 0xd6, 0x2f, 0xa9, 0x2c, 0x48, 0x2d,
	0xd6, 0x73, 0x49, 0x4d, 0x7e, 0x74, 0x4f, 0x5e, 0xd0, 0x37, 0xb1, 0xc2, 0x27, 0xbf, 0xb8, 0xd8,
	0x17, 0x6e, 0xd4, 0xab, 0x7b, 0xf2, 0xd8, 0x6c, 0x08, 0xc2, 0x26, 0xe8, 0xe4, 0x70, 0xe2, 0x91,
	0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1,
	0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xc8, 0x4e, 0x28, 0xc8, 0x49, 0xac, 0x4c, 0x2d,
	0x4a, 0x2b, 0x2d, 0xaa, 0xd4, 0x07, 0x07, 0x3c, 0x24, 0xe8, 0xc1, 0xce, 0x48, 0x62, 0x03, 0x07,
	0xac, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x89, 0x27, 0xdd, 0x93, 0x01, 0x00, 0x00,
}

func (m *BetOdds) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BetOdds) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BetOdds) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MaxLossMultiplier.Size()
		i -= size
		if _, err := m.MaxLossMultiplier.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBetOdds(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintBetOdds(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.MarketUID) > 0 {
		i -= len(m.MarketUID)
		copy(dAtA[i:], m.MarketUID)
		i = encodeVarintBetOdds(dAtA, i, uint64(len(m.MarketUID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.UID) > 0 {
		i -= len(m.UID)
		copy(dAtA[i:], m.UID)
		i = encodeVarintBetOdds(dAtA, i, uint64(len(m.UID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBetOdds(dAtA []byte, offset int, v uint64) int {
	offset -= sovBetOdds(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BetOdds) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UID)
	if l > 0 {
		n += 1 + l + sovBetOdds(uint64(l))
	}
	l = len(m.MarketUID)
	if l > 0 {
		n += 1 + l + sovBetOdds(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovBetOdds(uint64(l))
	}
	l = m.MaxLossMultiplier.Size()
	n += 1 + l + sovBetOdds(uint64(l))
	return n
}

func sovBetOdds(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBetOdds(x uint64) (n int) {
	return sovBetOdds(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BetOdds) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBetOdds
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
			return fmt.Errorf("proto: BetOdds: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BetOdds: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBetOdds
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
				return ErrInvalidLengthBetOdds
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBetOdds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBetOdds
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
				return ErrInvalidLengthBetOdds
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBetOdds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MarketUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBetOdds
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
				return ErrInvalidLengthBetOdds
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBetOdds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxLossMultiplier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBetOdds
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
				return ErrInvalidLengthBetOdds
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBetOdds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxLossMultiplier.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBetOdds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBetOdds
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
func skipBetOdds(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBetOdds
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
					return 0, ErrIntOverflowBetOdds
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
					return 0, ErrIntOverflowBetOdds
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
				return 0, ErrInvalidLengthBetOdds
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBetOdds
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBetOdds
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBetOdds        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBetOdds          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBetOdds = fmt.Errorf("proto: unexpected end of group")
)