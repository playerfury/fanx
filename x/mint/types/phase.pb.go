// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fanx/mint/phase.proto

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

// Phase defines the phase parameters for the module.
type Phase struct {
	// inflation is the current phase inflation rate.
	Inflation github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=inflation,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation" yaml:"inflation"`
	// year_coefficient is the proportion of a complete year.
	YearCoefficient github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=year_coefficient,json=yearCoefficient,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"year_coefficient" yaml:"year_coefficient"`
}

func (m *Phase) Reset()      { *m = Phase{} }
func (*Phase) ProtoMessage() {}
func (*Phase) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fc2241b526dba74, []int{0}
}
func (m *Phase) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Phase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Phase.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Phase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Phase.Merge(m, src)
}
func (m *Phase) XXX_Size() int {
	return m.Size()
}
func (m *Phase) XXX_DiscardUnknown() {
	xxx_messageInfo_Phase.DiscardUnknown(m)
}

var xxx_messageInfo_Phase proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Phase)(nil), "fanx.mint.Phase")
}

func init() { proto.RegisterFile("fanx/mint/phase.proto", fileDescriptor_7fc2241b526dba74) }

var fileDescriptor_7fc2241b526dba74 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x4e, 0x4f, 0xd5,
	0xcf, 0xcd, 0xcc, 0x2b, 0xd1, 0x2f, 0xc8, 0x48, 0x2c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x12, 0x2e, 0x4e, 0x4f, 0xcd, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x2b, 0x4e, 0x4f,
	0xd5, 0x03, 0x29, 0x90, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xcb, 0xeb, 0x83, 0x58, 0x10, 0xa5,
	0x4a, 0x4f, 0x19, 0xb9, 0x58, 0x03, 0x40, 0x5a, 0x85, 0x12, 0xb8, 0x38, 0x33, 0xf3, 0xd2, 0x72,
	0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x9d, 0x9c, 0x4e, 0xdc, 0x93,
	0x67, 0xb8, 0x75, 0x4f, 0x5e, 0x2d, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57,
	0x3f, 0x39, 0xbf, 0x38, 0x37, 0xbf, 0x18, 0x4a, 0xe9, 0x16, 0xa7, 0x64, 0xeb, 0x97, 0x54, 0x16,
	0xa4, 0x16, 0xeb, 0xb9, 0xa4, 0x26, 0x7f, 0xba, 0x27, 0x2f, 0x50, 0x99, 0x98, 0x9b, 0x63, 0xa5,
	0x04, 0x37, 0x48, 0x29, 0x08, 0x61, 0xa8, 0x50, 0x09, 0x97, 0x40, 0x65, 0x6a, 0x62, 0x51, 0x7c,
	0x72, 0x7e, 0x6a, 0x5a, 0x5a, 0x66, 0x72, 0x66, 0x6a, 0x5e, 0x89, 0x04, 0x13, 0xd8, 0x22, 0x4f,
	0x92, 0x2d, 0x12, 0x87, 0x58, 0x84, 0x6e, 0x9e, 0x52, 0x10, 0x3f, 0x48, 0xc8, 0x19, 0x21, 0x62,
	0xc5, 0x32, 0x63, 0x81, 0x3c, 0x83, 0x93, 0xe3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31,
	0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb,
	0x31, 0x44, 0xa9, 0x23, 0xd9, 0x59, 0x9c, 0x9e, 0xaa, 0x0b, 0x0d, 0x38, 0x10, 0x5b, 0xbf, 0x02,
	0x12, 0xb6, 0x60, 0x8b, 0x93, 0xd8, 0xc0, 0x21, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x78,
	0xe0, 0xa9, 0xf5, 0x74, 0x01, 0x00, 0x00,
}

func (m *Phase) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Phase) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Phase) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.YearCoefficient.Size()
		i -= size
		if _, err := m.YearCoefficient.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPhase(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Inflation.Size()
		i -= size
		if _, err := m.Inflation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPhase(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintPhase(dAtA []byte, offset int, v uint64) int {
	offset -= sovPhase(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Phase) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Inflation.Size()
	n += 1 + l + sovPhase(uint64(l))
	l = m.YearCoefficient.Size()
	n += 1 + l + sovPhase(uint64(l))
	return n
}

func sovPhase(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPhase(x uint64) (n int) {
	return sovPhase(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Phase) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPhase
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
			return fmt.Errorf("proto: Phase: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Phase: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inflation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPhase
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
				return ErrInvalidLengthPhase
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPhase
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Inflation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field YearCoefficient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPhase
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
				return ErrInvalidLengthPhase
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPhase
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.YearCoefficient.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPhase(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPhase
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
func skipPhase(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPhase
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
					return 0, ErrIntOverflowPhase
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
					return 0, ErrIntOverflowPhase
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
				return 0, ErrInvalidLengthPhase
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPhase
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPhase
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPhase        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPhase          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPhase = fmt.Errorf("proto: unexpected end of group")
)
