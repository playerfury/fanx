// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fanx/house/deposit.proto

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

// Deposit represents the deposit against a market held by an account.
type Deposit struct {
	// creator is the bech32-encoded address of the depositor.
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty" yaml:"creator"`
	// market_uid is the uid of market/order book against which deposit is being
	// made.
	MarketUID string `protobuf:"bytes,2,opt,name=market_uid,proto3" json:"market_uid"`
	// participation_index is the index corresponding to the order book
	// participation
	ParticipationIndex uint64 `protobuf:"varint,3,opt,name=participation_index,json=participationIndex,proto3" json:"participation_index,omitempty" yaml:"participation_index"`
	// amount is the amount being deposited on an order book to be a house
	Amount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount" yaml:"amount"`
	// fee is deducted from the deposited amount for participation in the order
	// book.
	Fee github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=fee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"fee" yaml:"fee"`
	// liquidity is the liquidity being provided to the order book after fee
	// deduction.
	Liquidity github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=liquidity,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"liquidity" yaml:"liquidity"`
	// withdrawal_count is the total count of the withdrawals from an order book
	WithdrawalCount uint64 `protobuf:"varint,7,opt,name=withdrawal_count,json=withdrawalCount,proto3" json:"withdrawal_count,omitempty" yaml:"withdrawals"`
	// total_withdrawal_amount is the total amount withdrawn from the liquidity
	// provided
	TotalWithdrawalAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,8,opt,name=total_withdrawal_amount,json=totalWithdrawalAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_withdrawal_amount" yaml:"total_withdrawal_amount"`
}

func (m *Deposit) Reset()      { *m = Deposit{} }
func (*Deposit) ProtoMessage() {}
func (*Deposit) Descriptor() ([]byte, []int) {
	return fileDescriptor_90b4fde25e30b2d5, []int{0}
}
func (m *Deposit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Deposit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Deposit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Deposit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deposit.Merge(m, src)
}
func (m *Deposit) XXX_Size() int {
	return m.Size()
}
func (m *Deposit) XXX_DiscardUnknown() {
	xxx_messageInfo_Deposit.DiscardUnknown(m)
}

var xxx_messageInfo_Deposit proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Deposit)(nil), "fanx.house.Deposit")
}

func init() { proto.RegisterFile("fanx/house/deposit.proto", fileDescriptor_90b4fde25e30b2d5) }

var fileDescriptor_90b4fde25e30b2d5 = []byte{
	// 460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x3d, 0x6f, 0xd3, 0x40,
	0x18, 0xc7, 0x6d, 0xda, 0x26, 0xe4, 0xc4, 0x4b, 0x75, 0xbc, 0x59, 0x45, 0xf2, 0x55, 0x37, 0xa0,
	0x0c, 0x10, 0x0f, 0x6c, 0x05, 0x09, 0xd5, 0x74, 0xc9, 0xc0, 0x8b, 0x2c, 0xa1, 0x4a, 0x2c, 0xe1,
	0x6a, 0x5f, 0x92, 0x53, 0x6d, 0x9f, 0x39, 0x9f, 0xd5, 0xfa, 0x1b, 0x74, 0x64, 0x64, 0xcc, 0x07,
	0xe0, 0x83, 0x74, 0xec, 0x88, 0x18, 0x4e, 0x28, 0x59, 0x10, 0x63, 0x3e, 0x01, 0xca, 0x73, 0x69,
	0x62, 0x24, 0x18, 0x32, 0xf9, 0xfc, 0xfc, 0xff, 0xcf, 0xef, 0xf9, 0xfb, 0x91, 0x0f, 0x79, 0x43,
	0x96, 0x9f, 0x07, 0x63, 0x59, 0x95, 0x3c, 0x48, 0x78, 0x21, 0x4b, 0xa1, 0x7b, 0x85, 0x92, 0x5a,
	0x62, 0xb4, 0x50, 0x7a, 0xa0, 0xec, 0xdd, 0x1f, 0xc9, 0x91, 0x84, 0x72, 0xb0, 0x38, 0x59, 0x07,
	0xfd, 0xb6, 0x83, 0xda, 0x47, 0xb6, 0x07, 0x3f, 0x45, 0xed, 0x58, 0x71, 0xa6, 0xa5, 0xf2, 0xdc,
	0x7d, 0xb7, 0xdb, 0x09, 0xf1, 0xdc, 0x90, 0x3b, 0x35, 0xcb, 0xd2, 0x03, 0xba, 0x14, 0x68, 0x74,
	0x6d, 0xc1, 0x2f, 0x10, 0xca, 0x98, 0x3a, 0xe5, 0x7a, 0x50, 0x89, 0xc4, 0xbb, 0x01, 0x0d, 0x8f,
	0xa7, 0x86, 0x74, 0xde, 0x40, 0xf5, 0x43, 0xff, 0xe8, 0xb7, 0x21, 0x0d, 0x4b, 0xd4, 0x38, 0xe3,
	0x77, 0xe8, 0x5e, 0xc1, 0x94, 0x16, 0xb1, 0x28, 0x98, 0x16, 0x32, 0x1f, 0x88, 0x3c, 0xe1, 0xe7,
	0xde, 0xd6, 0xbe, 0xdb, 0xdd, 0x0e, 0xfd, 0xb9, 0x21, 0x7b, 0x76, 0xec, 0x3f, 0x4c, 0x34, 0xc2,
	0x7f, 0x55, 0xfb, 0x8b, 0x22, 0x3e, 0x46, 0x2d, 0x96, 0xc9, 0x2a, 0xd7, 0xde, 0x36, 0x24, 0x79,
	0x75, 0x69, 0x88, 0xf3, 0xc3, 0x90, 0x27, 0x23, 0xa1, 0xc7, 0xd5, 0x49, 0x2f, 0x96, 0x59, 0x10,
	0xcb, 0x32, 0x93, 0xe5, 0xf2, 0xf1, 0xac, 0x4c, 0x4e, 0x03, 0x5d, 0x17, 0xbc, 0xec, 0xf5, 0x73,
	0x3d, 0x37, 0xe4, 0xb6, 0x9d, 0x68, 0x29, 0x34, 0x5a, 0xe2, 0xf0, 0x5b, 0xb4, 0x35, 0xe4, 0xdc,
	0xdb, 0x01, 0xea, 0xcb, 0x8d, 0xa9, 0xc8, 0x52, 0x87, 0x9c, 0xd3, 0x68, 0x01, 0xc2, 0x9f, 0x50,
	0x27, 0x15, 0x9f, 0x2b, 0x91, 0x08, 0x5d, 0x7b, 0x2d, 0xa0, 0x86, 0x1b, 0x53, 0x77, 0x2d, 0x75,
	0x05, 0xa2, 0xd1, 0x1a, 0x8a, 0x0f, 0xd1, 0xee, 0x99, 0xd0, 0xe3, 0x44, 0xb1, 0x33, 0x96, 0x0e,
	0x62, 0x58, 0x4a, 0x1b, 0x16, 0xfb, 0x70, 0x6e, 0x08, 0xb6, 0xad, 0x6b, 0x47, 0x49, 0xa3, 0xbb,
	0xeb, 0xb7, 0xd7, 0xf0, 0xd1, 0x17, 0x2e, 0x7a, 0xa4, 0xa5, 0x66, 0xe9, 0xa0, 0x41, 0x5a, 0xee,
	0xf7, 0x26, 0x64, 0x7e, 0xbf, 0x71, 0x66, 0xdf, 0x0e, 0xfe, 0x0f, 0x96, 0x46, 0x0f, 0x40, 0x39,
	0x5e, 0x09, 0x87, 0x50, 0x3f, 0xb8, 0x75, 0x31, 0x21, 0xce, 0xd7, 0x09, 0x71, 0x7e, 0x4d, 0x88,
	0x13, 0x86, 0x97, 0x53, 0xdf, 0xbd, 0x9a, 0xfa, 0xee, 0xcf, 0xa9, 0xef, 0x7e, 0x99, 0xf9, 0xce,
	0xd5, 0xcc, 0x77, 0xbe, 0xcf, 0x7c, 0xe7, 0x63, 0xb7, 0x11, 0xa4, 0x48, 0x59, 0xcd, 0xd5, 0xb0,
	0x52, 0x75, 0x00, 0x57, 0xe3, 0xfa, 0x72, 0x40, 0x9c, 0x93, 0x16, 0xfc, 0xf9, 0xcf, 0xff, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x7a, 0xac, 0x16, 0xdb, 0x37, 0x03, 0x00, 0x00,
}

func (m *Deposit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Deposit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Deposit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TotalWithdrawalAmount.Size()
		i -= size
		if _, err := m.TotalWithdrawalAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDeposit(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if m.WithdrawalCount != 0 {
		i = encodeVarintDeposit(dAtA, i, uint64(m.WithdrawalCount))
		i--
		dAtA[i] = 0x38
	}
	{
		size := m.Liquidity.Size()
		i -= size
		if _, err := m.Liquidity.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDeposit(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.Fee.Size()
		i -= size
		if _, err := m.Fee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDeposit(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDeposit(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.ParticipationIndex != 0 {
		i = encodeVarintDeposit(dAtA, i, uint64(m.ParticipationIndex))
		i--
		dAtA[i] = 0x18
	}
	if len(m.MarketUID) > 0 {
		i -= len(m.MarketUID)
		copy(dAtA[i:], m.MarketUID)
		i = encodeVarintDeposit(dAtA, i, uint64(len(m.MarketUID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintDeposit(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDeposit(dAtA []byte, offset int, v uint64) int {
	offset -= sovDeposit(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Deposit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovDeposit(uint64(l))
	}
	l = len(m.MarketUID)
	if l > 0 {
		n += 1 + l + sovDeposit(uint64(l))
	}
	if m.ParticipationIndex != 0 {
		n += 1 + sovDeposit(uint64(m.ParticipationIndex))
	}
	l = m.Amount.Size()
	n += 1 + l + sovDeposit(uint64(l))
	l = m.Fee.Size()
	n += 1 + l + sovDeposit(uint64(l))
	l = m.Liquidity.Size()
	n += 1 + l + sovDeposit(uint64(l))
	if m.WithdrawalCount != 0 {
		n += 1 + sovDeposit(uint64(m.WithdrawalCount))
	}
	l = m.TotalWithdrawalAmount.Size()
	n += 1 + l + sovDeposit(uint64(l))
	return n
}

func sovDeposit(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDeposit(x uint64) (n int) {
	return sovDeposit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Deposit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeposit
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
			return fmt.Errorf("proto: Deposit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Deposit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
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
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
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
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MarketUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParticipationIndex", wireType)
			}
			m.ParticipationIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ParticipationIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
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
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
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
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Liquidity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
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
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Liquidity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawalCount", wireType)
			}
			m.WithdrawalCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WithdrawalCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalWithdrawalAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
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
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalWithdrawalAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeposit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDeposit
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
func skipDeposit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeposit
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
					return 0, ErrIntOverflowDeposit
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
					return 0, ErrIntOverflowDeposit
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
				return 0, ErrInvalidLengthDeposit
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDeposit
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDeposit
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDeposit        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeposit          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDeposit = fmt.Errorf("proto: unexpected end of group")
)
