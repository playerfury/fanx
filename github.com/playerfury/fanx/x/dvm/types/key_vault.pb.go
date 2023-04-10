// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fanx/dvm/key_vault.proto

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

// KeyVault is the information of important keys stored in dvm state.
type KeyVault struct {
	// public_keys contains allowed public keys.
	PublicKeys []string `protobuf:"bytes,1,rep,name=public_keys,json=publicKeys,proto3" json:"public_keys,omitempty"`
}

func (m *KeyVault) Reset()         { *m = KeyVault{} }
func (m *KeyVault) String() string { return proto.CompactTextString(m) }
func (*KeyVault) ProtoMessage()    {}
func (*KeyVault) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1fe327b9f1e0e3f, []int{0}
}
func (m *KeyVault) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeyVault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeyVault.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeyVault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyVault.Merge(m, src)
}
func (m *KeyVault) XXX_Size() int {
	return m.Size()
}
func (m *KeyVault) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyVault.DiscardUnknown(m)
}

var xxx_messageInfo_KeyVault proto.InternalMessageInfo

func (m *KeyVault) GetPublicKeys() []string {
	if m != nil {
		return m.PublicKeys
	}
	return nil
}

func init() {
	proto.RegisterType((*KeyVault)(nil), "fanx.dvm.KeyVault")
}

func init() { proto.RegisterFile("fanx/dvm/key_vault.proto", fileDescriptor_b1fe327b9f1e0e3f) }

var fileDescriptor_b1fe327b9f1e0e3f = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0x4b, 0xcc, 0xab,
	0xd0, 0x4f, 0x29, 0xcb, 0xd5, 0xcf, 0x4e, 0xad, 0x8c, 0x2f, 0x4b, 0x2c, 0xcd, 0x29, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0xc9, 0xe8, 0xa5, 0x94, 0xe5, 0x4a, 0x89, 0xa4, 0xe7,
	0xa7, 0xe7, 0x83, 0x05, 0xf5, 0x41, 0x2c, 0x88, 0xbc, 0x92, 0x21, 0x17, 0x87, 0x77, 0x6a, 0x65,
	0x18, 0x48, 0x87, 0x90, 0x2a, 0x17, 0x77, 0x41, 0x69, 0x52, 0x4e, 0x66, 0x72, 0x7c, 0x76, 0x6a,
	0x65, 0xb1, 0x04, 0xa3, 0x02, 0xb3, 0x06, 0xa7, 0x13, 0xcb, 0x89, 0x7b, 0xf2, 0x0c, 0x41, 0x5c,
	0x10, 0x09, 0xef, 0xd4, 0xca, 0x62, 0x27, 0x87, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63,
	0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96,
	0x63, 0x88, 0x52, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xc8,
	0x49, 0xac, 0x4c, 0x2d, 0x4a, 0x2b, 0x2d, 0xaa, 0xd4, 0x07, 0x3b, 0x0e, 0xe2, 0xbc, 0x92, 0xca,
	0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0xdd, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x20, 0x77,
	0xcd, 0xb8, 0xb7, 0x00, 0x00, 0x00,
}

func (m *KeyVault) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeyVault) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeyVault) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PublicKeys) > 0 {
		for iNdEx := len(m.PublicKeys) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PublicKeys[iNdEx])
			copy(dAtA[i:], m.PublicKeys[iNdEx])
			i = encodeVarintKeyVault(dAtA, i, uint64(len(m.PublicKeys[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintKeyVault(dAtA []byte, offset int, v uint64) int {
	offset -= sovKeyVault(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *KeyVault) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PublicKeys) > 0 {
		for _, s := range m.PublicKeys {
			l = len(s)
			n += 1 + l + sovKeyVault(uint64(l))
		}
	}
	return n
}

func sovKeyVault(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKeyVault(x uint64) (n int) {
	return sovKeyVault(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *KeyVault) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKeyVault
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
			return fmt.Errorf("proto: KeyVault: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyVault: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKeys", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyVault
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
				return ErrInvalidLengthKeyVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKeyVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKeys = append(m.PublicKeys, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKeyVault(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthKeyVault
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
func skipKeyVault(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKeyVault
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
					return 0, ErrIntOverflowKeyVault
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
					return 0, ErrIntOverflowKeyVault
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
				return 0, ErrInvalidLengthKeyVault
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupKeyVault
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthKeyVault
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthKeyVault        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKeyVault          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupKeyVault = fmt.Errorf("proto: unexpected end of group")
)
