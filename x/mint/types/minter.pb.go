// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fanx/mint/minter.proto

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

// Minter represents the minting state.
type Minter struct {
	// inflation is the current annual inflation rate.
	Inflation github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=inflation,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation"`
	// phase_step is the index of phases slice + 1.
	PhaseStep int32 `protobuf:"varint,2,opt,name=phase_step,json=phaseStep,proto3" json:"phase_step,omitempty"`
	// phase_provisions is the current phase expected provisions.
	PhaseProvisions github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=phase_provisions,json=phaseProvisions,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"phase_provisions" yaml:"phase_provisions"`
	// truncated_tokens holds current truncated tokens because of Dec to Int
	// conversion in the minting.
	TruncatedTokens github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=truncated_tokens,json=truncatedTokens,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"truncated_tokens"`
}

func (m *Minter) Reset()         { *m = Minter{} }
func (m *Minter) String() string { return proto.CompactTextString(m) }
func (*Minter) ProtoMessage()    {}
func (*Minter) Descriptor() ([]byte, []int) {
	return fileDescriptor_382275596952eec1, []int{0}
}
func (m *Minter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Minter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Minter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Minter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Minter.Merge(m, src)
}
func (m *Minter) XXX_Size() int {
	return m.Size()
}
func (m *Minter) XXX_DiscardUnknown() {
	xxx_messageInfo_Minter.DiscardUnknown(m)
}

var xxx_messageInfo_Minter proto.InternalMessageInfo

func (m *Minter) GetPhaseStep() int32 {
	if m != nil {
		return m.PhaseStep
	}
	return 0
}

func init() {
	proto.RegisterType((*Minter)(nil), "fanx.mint.Minter")
}

func init() { proto.RegisterFile("fanx/mint/minter.proto", fileDescriptor_382275596952eec1) }

var fileDescriptor_382275596952eec1 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0x4b, 0xcc, 0xab,
	0xd0, 0xcf, 0xcd, 0xcc, 0x2b, 0x01, 0x13, 0xa9, 0x45, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42,
	0x9c, 0x20, 0x71, 0x3d, 0x90, 0x90, 0x94, 0x48, 0x7a, 0x7e, 0x7a, 0x3e, 0x58, 0x54, 0x1f, 0xc4,
	0x82, 0x28, 0x50, 0x3a, 0xc6, 0xc4, 0xc5, 0xe6, 0x0b, 0xd6, 0x21, 0xe4, 0xc3, 0xc5, 0x99, 0x99,
	0x97, 0x96, 0x93, 0x58, 0x92, 0x99, 0x9f, 0x27, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0xe9, 0xa4, 0x77,
	0xe2, 0x9e, 0x3c, 0xc3, 0xad, 0x7b, 0xf2, 0x6a, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9,
	0xf9, 0xb9, 0xfa, 0xc9, 0xf9, 0xc5, 0xb9, 0xf9, 0xc5, 0x50, 0x4a, 0xb7, 0x38, 0x25, 0x5b, 0xbf,
	0xa4, 0xb2, 0x20, 0xb5, 0x58, 0xcf, 0x25, 0x35, 0x39, 0x08, 0x61, 0x80, 0x90, 0x2c, 0x17, 0x57,
	0x41, 0x46, 0x62, 0x71, 0x6a, 0x7c, 0x71, 0x49, 0x6a, 0x81, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x6b,
	0x10, 0x27, 0x58, 0x24, 0xb8, 0x24, 0xb5, 0x40, 0xa8, 0x84, 0x4b, 0x00, 0x22, 0x5d, 0x50, 0x94,
	0x5f, 0x96, 0x59, 0x9c, 0x99, 0x9f, 0x57, 0x2c, 0xc1, 0x0c, 0xb6, 0xd3, 0x93, 0x34, 0x3b, 0x3f,
	0xdd, 0x93, 0x17, 0xaf, 0x4c, 0xcc, 0xcd, 0xb1, 0x52, 0x42, 0x37, 0x4f, 0x29, 0x88, 0x1f, 0x2c,
	0x14, 0x00, 0x17, 0x11, 0x8a, 0xe4, 0x12, 0x28, 0x29, 0x2a, 0xcd, 0x4b, 0x4e, 0x2c, 0x49, 0x4d,
	0x89, 0x2f, 0xc9, 0xcf, 0x4e, 0xcd, 0x2b, 0x96, 0x60, 0x21, 0xcb, 0xa7, 0xfc, 0x70, 0x73, 0x42,
	0xc0, 0xc6, 0x38, 0x39, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72,
	0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x3a,
	0x92, 0x91, 0x05, 0x39, 0x89, 0x95, 0xa9, 0x45, 0x69, 0xa5, 0x45, 0x95, 0xfa, 0xe0, 0x18, 0x83,
	0xc6, 0x19, 0xd8, 0xdc, 0x24, 0x36, 0x70, 0x94, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x06,
	0x3f, 0x52, 0xe9, 0xcd, 0x01, 0x00, 0x00,
}

func (m *Minter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Minter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Minter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TruncatedTokens.Size()
		i -= size
		if _, err := m.TruncatedTokens.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMinter(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.PhaseProvisions.Size()
		i -= size
		if _, err := m.PhaseProvisions.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMinter(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.PhaseStep != 0 {
		i = encodeVarintMinter(dAtA, i, uint64(m.PhaseStep))
		i--
		dAtA[i] = 0x10
	}
	{
		size := m.Inflation.Size()
		i -= size
		if _, err := m.Inflation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMinter(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintMinter(dAtA []byte, offset int, v uint64) int {
	offset -= sovMinter(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Minter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Inflation.Size()
	n += 1 + l + sovMinter(uint64(l))
	if m.PhaseStep != 0 {
		n += 1 + sovMinter(uint64(m.PhaseStep))
	}
	l = m.PhaseProvisions.Size()
	n += 1 + l + sovMinter(uint64(l))
	l = m.TruncatedTokens.Size()
	n += 1 + l + sovMinter(uint64(l))
	return n
}

func sovMinter(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMinter(x uint64) (n int) {
	return sovMinter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Minter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMinter
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
			return fmt.Errorf("proto: Minter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Minter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inflation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
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
				return ErrInvalidLengthMinter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMinter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Inflation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PhaseStep", wireType)
			}
			m.PhaseStep = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PhaseStep |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PhaseProvisions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
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
				return ErrInvalidLengthMinter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMinter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PhaseProvisions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TruncatedTokens", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
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
				return ErrInvalidLengthMinter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMinter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TruncatedTokens.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMinter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMinter
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
func skipMinter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMinter
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
					return 0, ErrIntOverflowMinter
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
					return 0, ErrIntOverflowMinter
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
				return 0, ErrInvalidLengthMinter
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMinter
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMinter
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMinter        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMinter          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMinter = fmt.Errorf("proto: unexpected end of group")
)
