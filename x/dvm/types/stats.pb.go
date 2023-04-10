// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fanx/dvm/stats.proto

package types

import (
	fmt "fmt"
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

// ProposalStats is the type for the proposal statistics.
type ProposalStats struct {
	// pubkeys_change_count is the pubkeys change proposals count.
	PubkeysChangeCount uint64 `protobuf:"varint,1,opt,name=pubkeys_change_count,json=pubkeysChangeCount,proto3" json:"pubkeys_change_count,omitempty"`
}

func (m *ProposalStats) Reset()         { *m = ProposalStats{} }
func (m *ProposalStats) String() string { return proto.CompactTextString(m) }
func (*ProposalStats) ProtoMessage()    {}
func (*ProposalStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_18b1583139df99d6, []int{0}
}
func (m *ProposalStats) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProposalStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProposalStats.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProposalStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalStats.Merge(m, src)
}
func (m *ProposalStats) XXX_Size() int {
	return m.Size()
}
func (m *ProposalStats) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalStats.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalStats proto.InternalMessageInfo

func (m *ProposalStats) GetPubkeysChangeCount() uint64 {
	if m != nil {
		return m.PubkeysChangeCount
	}
	return 0
}

func init() {
	proto.RegisterType((*ProposalStats)(nil), "fanx.fanx.dvm.ProposalStats")
}

func init() { proto.RegisterFile("fanx/dvm/stats.proto", fileDescriptor_18b1583139df99d6) }

var fileDescriptor_18b1583139df99d6 = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x4e, 0x4f, 0xd5,
	0x4f, 0x29, 0xcb, 0xd5, 0x2f, 0x2e, 0x49, 0x2c, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0x2a, 0x4e, 0x4f, 0xcd, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x2b, 0x4e, 0x4f, 0xd5,
	0x4b, 0x29, 0xcb, 0x55, 0x72, 0xe4, 0xe2, 0x0d, 0x28, 0xca, 0x2f, 0xc8, 0x2f, 0x4e, 0xcc, 0x09,
	0x06, 0x29, 0x15, 0x32, 0xe0, 0x12, 0x29, 0x28, 0x4d, 0xca, 0x4e, 0xad, 0x2c, 0x8e, 0x4f, 0xce,
	0x48, 0xcc, 0x4b, 0x4f, 0x8d, 0x4f, 0xce, 0x2f, 0xcd, 0x2b, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60,
	0x09, 0x12, 0x82, 0xca, 0x39, 0x83, 0xa5, 0x9c, 0x41, 0x32, 0x4e, 0x0e, 0x27, 0x1e, 0xc9, 0x31,
	0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb,
	0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x96, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c,
	0x9f, 0xab, 0x5f, 0x9c, 0x9e, 0xaa, 0x0b, 0xb5, 0x1c, 0xc4, 0xd6, 0xaf, 0x00, 0x3b, 0xaf, 0xa4,
	0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x3e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5f,
	0x7b, 0x02, 0x61, 0xb6, 0x00, 0x00, 0x00,
}

func (m *ProposalStats) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProposalStats) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProposalStats) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PubkeysChangeCount != 0 {
		i = encodeVarintStats(dAtA, i, uint64(m.PubkeysChangeCount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintStats(dAtA []byte, offset int, v uint64) int {
	offset -= sovStats(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProposalStats) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PubkeysChangeCount != 0 {
		n += 1 + sovStats(uint64(m.PubkeysChangeCount))
	}
	return n
}

func sovStats(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStats(x uint64) (n int) {
	return sovStats(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProposalStats) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStats
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
			return fmt.Errorf("proto: ProposalStats: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProposalStats: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubkeysChangeCount", wireType)
			}
			m.PubkeysChangeCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PubkeysChangeCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStats
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
func skipStats(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStats
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
					return 0, ErrIntOverflowStats
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
					return 0, ErrIntOverflowStats
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
				return 0, ErrInvalidLengthStats
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStats
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStats
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStats        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStats          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStats = fmt.Errorf("proto: unexpected end of group")
)
