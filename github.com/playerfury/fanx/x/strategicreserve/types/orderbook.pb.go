// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fanx/strategicreserve/orderbook.proto

package types

import (
	fmt "fmt"
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

// OrderBookStatus is the enum type for the status of the order book.
type OrderBookStatus int32

const (
	// invalid
	OrderBookStatus_ORDER_BOOK_STATUS_UNSPECIFIED OrderBookStatus = 0
	// active
	OrderBookStatus_ORDER_BOOK_STATUS_STATUS_ACTIVE OrderBookStatus = 1
	// resolved not settled
	OrderBookStatus_ORDER_BOOK_STATUS_STATUS_RESOLVED OrderBookStatus = 2
	// resolved and settled
	OrderBookStatus_ORDER_BOOK_STATUS_STATUS_SETTLED OrderBookStatus = 3
)

var OrderBookStatus_name = map[int32]string{
	0: "ORDER_BOOK_STATUS_UNSPECIFIED",
	1: "ORDER_BOOK_STATUS_STATUS_ACTIVE",
	2: "ORDER_BOOK_STATUS_STATUS_RESOLVED",
	3: "ORDER_BOOK_STATUS_STATUS_SETTLED",
}

var OrderBookStatus_value = map[string]int32{
	"ORDER_BOOK_STATUS_UNSPECIFIED":     0,
	"ORDER_BOOK_STATUS_STATUS_ACTIVE":   1,
	"ORDER_BOOK_STATUS_STATUS_RESOLVED": 2,
	"ORDER_BOOK_STATUS_STATUS_SETTLED":  3,
}

func (x OrderBookStatus) String() string {
	return proto.EnumName(OrderBookStatus_name, int32(x))
}

func (OrderBookStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_69fde81e7f5eb06e, []int{0}
}

// OrderBook represents the order book maintained against a market.
type OrderBook struct {
	// uid is the universal unique identifier of the order book.
	UID string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid"`
	// participation_count is the count of participations in the order book.
	ParticipationCount uint64 `protobuf:"varint,2,opt,name=participation_count,json=participationCount,proto3" json:"participation_count,omitempty" yaml:"participation_count"`
	// odds_count is the count of the odds in the order book.
	OddsCount uint64 `protobuf:"varint,3,opt,name=odds_count,json=oddsCount,proto3" json:"odds_count,omitempty" yaml:"odds_count"`
	// status represents the status of the order book.
	Status OrderBookStatus `protobuf:"varint,4,opt,name=status,proto3,enum=fanx.strategicreserve.OrderBookStatus" json:"status,omitempty"`
}

func (m *OrderBook) Reset()      { *m = OrderBook{} }
func (*OrderBook) ProtoMessage() {}
func (*OrderBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_69fde81e7f5eb06e, []int{0}
}
func (m *OrderBook) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OrderBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OrderBook.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OrderBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderBook.Merge(m, src)
}
func (m *OrderBook) XXX_Size() int {
	return m.Size()
}
func (m *OrderBook) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderBook.DiscardUnknown(m)
}

var xxx_messageInfo_OrderBook proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("fanx.strategicreserve.OrderBookStatus", OrderBookStatus_name, OrderBookStatus_value)
	proto.RegisterType((*OrderBook)(nil), "fanx.strategicreserve.OrderBook")
}

func init() {
	proto.RegisterFile("fanx/strategicreserve/orderbook.proto", fileDescriptor_69fde81e7f5eb06e)
}

var fileDescriptor_69fde81e7f5eb06e = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0x4b, 0xcc, 0xab,
	0xd0, 0x2f, 0x2e, 0x29, 0x4a, 0x2c, 0x49, 0x4d, 0xcf, 0x4c, 0x2e, 0x4a, 0x2d, 0x4e, 0x2d, 0x2a,
	0x4b, 0xd5, 0xcf, 0x2f, 0x4a, 0x49, 0x2d, 0x4a, 0xca, 0xcf, 0xcf, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x12, 0x05, 0x29, 0xd3, 0x43, 0x57, 0x26, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x56,
	0xa1, 0x0f, 0x62, 0x41, 0x14, 0x2b, 0x35, 0x31, 0x71, 0x71, 0xfa, 0x83, 0x0c, 0x70, 0xca, 0xcf,
	0xcf, 0x16, 0x52, 0xe0, 0x62, 0x2e, 0xcd, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x74, 0xe2,
	0x7b, 0x74, 0x4f, 0x9e, 0x39, 0xd4, 0xd3, 0xe5, 0xd5, 0x3d, 0x79, 0x90, 0x68, 0x10, 0x88, 0x10,
	0xf2, 0xe7, 0x12, 0x2e, 0x48, 0x2c, 0x2a, 0xc9, 0x4c, 0xce, 0x2c, 0x48, 0x2c, 0xc9, 0xcc, 0xcf,
	0x8b, 0x4f, 0xce, 0x2f, 0xcd, 0x2b, 0x91, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x71, 0x92, 0xfb, 0x74,
	0x4f, 0x5e, 0xaa, 0x32, 0x31, 0x37, 0xc7, 0x4a, 0x09, 0x8b, 0x22, 0xa5, 0x20, 0x21, 0x14, 0x51,
	0x67, 0x90, 0xa0, 0x90, 0x09, 0x17, 0x57, 0x7e, 0x4a, 0x4a, 0x31, 0xd4, 0x1c, 0x66, 0xb0, 0x39,
	0xa2, 0x9f, 0xee, 0xc9, 0x0b, 0x42, 0xcc, 0x41, 0xc8, 0x29, 0x05, 0x71, 0x82, 0x38, 0x10, 0x5d,
	0x76, 0x5c, 0x6c, 0xc5, 0x25, 0x89, 0x25, 0xa5, 0xc5, 0x12, 0x2c, 0x0a, 0x8c, 0x1a, 0x7c, 0x46,
	0x6a, 0x7a, 0x58, 0x3d, 0xad, 0x07, 0xf7, 0x5a, 0x30, 0x58, 0x75, 0x10, 0x54, 0x97, 0x15, 0x4f,
	0xc7, 0x02, 0x79, 0x86, 0x19, 0x0b, 0xe4, 0x19, 0x5e, 0x2c, 0x90, 0x67, 0xd0, 0x5a, 0xc6, 0xc8,
	0xc5, 0x8f, 0xa6, 0x52, 0x48, 0x91, 0x4b, 0xd6, 0x3f, 0xc8, 0xc5, 0x35, 0x28, 0xde, 0xc9, 0xdf,
	0xdf, 0x3b, 0x3e, 0x38, 0xc4, 0x31, 0x24, 0x34, 0x38, 0x3e, 0xd4, 0x2f, 0x38, 0xc0, 0xd5, 0xd9,
	0xd3, 0xcd, 0xd3, 0xd5, 0x45, 0x80, 0x41, 0x48, 0x99, 0x4b, 0x1e, 0x53, 0x09, 0x94, 0x72, 0x74,
	0x0e, 0xf1, 0x0c, 0x73, 0x15, 0x60, 0x14, 0x52, 0xe5, 0x52, 0xc4, 0xa9, 0x28, 0xc8, 0x35, 0xd8,
	0xdf, 0x27, 0xcc, 0xd5, 0x45, 0x80, 0x49, 0x48, 0x85, 0x4b, 0x01, 0xa7, 0xb2, 0x60, 0xd7, 0x90,
	0x10, 0x1f, 0x57, 0x17, 0x01, 0x66, 0x27, 0xdf, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63,
	0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96,
	0x63, 0x88, 0x32, 0x4e, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xc8,
	0x49, 0xac, 0x4c, 0x2d, 0x4a, 0x2b, 0x2d, 0xaa, 0xd4, 0x07, 0xa7, 0x18, 0x2c, 0x69, 0xa6, 0xa4,
	0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0x9c, 0x06, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x75,
	0x2a, 0x64, 0xb5, 0x59, 0x02, 0x00, 0x00,
}

func (m *OrderBook) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrderBook) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OrderBook) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintOrderbook(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if m.OddsCount != 0 {
		i = encodeVarintOrderbook(dAtA, i, uint64(m.OddsCount))
		i--
		dAtA[i] = 0x18
	}
	if m.ParticipationCount != 0 {
		i = encodeVarintOrderbook(dAtA, i, uint64(m.ParticipationCount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.UID) > 0 {
		i -= len(m.UID)
		copy(dAtA[i:], m.UID)
		i = encodeVarintOrderbook(dAtA, i, uint64(len(m.UID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOrderbook(dAtA []byte, offset int, v uint64) int {
	offset -= sovOrderbook(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OrderBook) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UID)
	if l > 0 {
		n += 1 + l + sovOrderbook(uint64(l))
	}
	if m.ParticipationCount != 0 {
		n += 1 + sovOrderbook(uint64(m.ParticipationCount))
	}
	if m.OddsCount != 0 {
		n += 1 + sovOrderbook(uint64(m.OddsCount))
	}
	if m.Status != 0 {
		n += 1 + sovOrderbook(uint64(m.Status))
	}
	return n
}

func sovOrderbook(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOrderbook(x uint64) (n int) {
	return sovOrderbook(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OrderBook) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrderbook
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
			return fmt.Errorf("proto: OrderBook: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrderBook: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderbook
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
				return ErrInvalidLengthOrderbook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrderbook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParticipationCount", wireType)
			}
			m.ParticipationCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderbook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ParticipationCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OddsCount", wireType)
			}
			m.OddsCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderbook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OddsCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderbook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= OrderBookStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOrderbook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrderbook
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
func skipOrderbook(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrderbook
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
					return 0, ErrIntOverflowOrderbook
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
					return 0, ErrIntOverflowOrderbook
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
				return 0, ErrInvalidLengthOrderbook
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOrderbook
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOrderbook
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOrderbook        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrderbook          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOrderbook = fmt.Errorf("proto: unexpected end of group")
)
