// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: furyx/dvm/vote.proto

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

// ProposalVote is the enum type for the proposal vote.
type ProposalVote int32

const (
	// unchosen value
	ProposalVote_PROPOSAL_VOTE_UNSPECIFIED ProposalVote = 0
	// no
	ProposalVote_PROPOSAL_VOTE_NO ProposalVote = 1
	// yes
	ProposalVote_PROPOSAL_VOTE_YES ProposalVote = 2
)

var ProposalVote_name = map[int32]string{
	0: "PROPOSAL_VOTE_UNSPECIFIED",
	1: "PROPOSAL_VOTE_NO",
	2: "PROPOSAL_VOTE_YES",
}

var ProposalVote_value = map[string]int32{
	"PROPOSAL_VOTE_UNSPECIFIED": 0,
	"PROPOSAL_VOTE_NO":          1,
	"PROPOSAL_VOTE_YES":         2,
}

func (x ProposalVote) String() string {
	return proto.EnumName(ProposalVote_name, int32(x))
}

func (ProposalVote) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c881d83c5c5286a5, []int{0}
}

// Vote is the type for the proposal vote.
type Vote struct {
	// public_key is the public key of the voter.
	PublicKey string `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// vote is the vote enum value.
	Vote ProposalVote `protobuf:"varint,2,opt,name=vote,proto3,enum=furyx.furyx.dvm.ProposalVote" json:"vote,omitempty"`
}

func (m *Vote) Reset()         { *m = Vote{} }
func (m *Vote) String() string { return proto.CompactTextString(m) }
func (*Vote) ProtoMessage()    {}
func (*Vote) Descriptor() ([]byte, []int) {
	return fileDescriptor_c881d83c5c5286a5, []int{0}
}
func (m *Vote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Vote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Vote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Vote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vote.Merge(m, src)
}
func (m *Vote) XXX_Size() int {
	return m.Size()
}
func (m *Vote) XXX_DiscardUnknown() {
	xxx_messageInfo_Vote.DiscardUnknown(m)
}

var xxx_messageInfo_Vote proto.InternalMessageInfo

func (m *Vote) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func (m *Vote) GetVote() ProposalVote {
	if m != nil {
		return m.Vote
	}
	return ProposalVote_PROPOSAL_VOTE_UNSPECIFIED
}

func init() {
	proto.RegisterEnum("furyx.furyx.dvm.ProposalVote", ProposalVote_name, ProposalVote_value)
	proto.RegisterType((*Vote)(nil), "furyx.furyx.dvm.Vote")
}

func init() { proto.RegisterFile("furyx/dvm/vote.proto", fileDescriptor_c881d83c5c5286a5) }

var fileDescriptor_c881d83c5c5286a5 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x4e, 0x4f, 0xd5,
	0x4f, 0x29, 0xcb, 0xd5, 0x2f, 0xcb, 0x2f, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x02,
	0x89, 0xe5, 0xa5, 0x96, 0x94, 0xe7, 0x17, 0x65, 0xeb, 0x15, 0xa7, 0xa7, 0xea, 0xa5, 0x94, 0xe5,
	0x2a, 0x45, 0x73, 0xb1, 0x84, 0xe5, 0x97, 0xa4, 0x0a, 0xc9, 0x72, 0x71, 0x15, 0x94, 0x26, 0xe5,
	0x64, 0x26, 0xc7, 0x67, 0xa7, 0x56, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x71, 0x42, 0x44,
	0xbc, 0x53, 0x2b, 0x85, 0x4c, 0xb8, 0x58, 0x40, 0x06, 0x49, 0x30, 0x29, 0x30, 0x6a, 0xf0, 0x19,
	0x29, 0xe8, 0x61, 0x9a, 0xa4, 0x17, 0x50, 0x94, 0x5f, 0x90, 0x5f, 0x9c, 0x98, 0x03, 0x32, 0x2e,
	0x08, 0xac, 0x5a, 0x2b, 0x8a, 0x8b, 0x07, 0x59, 0x54, 0x48, 0x96, 0x4b, 0x32, 0x20, 0xc8, 0x3f,
	0xc0, 0x3f, 0xd8, 0xd1, 0x27, 0x3e, 0xcc, 0x3f, 0xc4, 0x35, 0x3e, 0xd4, 0x2f, 0x38, 0xc0, 0xd5,
	0xd9, 0xd3, 0xcd, 0xd3, 0xd5, 0x45, 0x80, 0x41, 0x48, 0x84, 0x4b, 0x00, 0x55, 0xda, 0xcf, 0x5f,
	0x80, 0x51, 0x48, 0x94, 0x4b, 0x10, 0x55, 0x34, 0xd2, 0x35, 0x58, 0x80, 0xc9, 0xc9, 0xe1, 0xc4,
	0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1,
	0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xd4, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93,
	0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x8b, 0xd3, 0x53, 0x75, 0xa1, 0x0e, 0x05, 0xb1, 0xf5, 0x2b, 0xc0,
	0x61, 0x52, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x0e, 0x15, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xcf, 0xc4, 0x8d, 0x37, 0x2b, 0x01, 0x00, 0x00,
}

func (m *Vote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Vote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Vote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Vote != 0 {
		i = encodeVarintVote(dAtA, i, uint64(m.Vote))
		i--
		dAtA[i] = 0x10
	}
	if len(m.PublicKey) > 0 {
		i -= len(m.PublicKey)
		copy(dAtA[i:], m.PublicKey)
		i = encodeVarintVote(dAtA, i, uint64(len(m.PublicKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVote(dAtA []byte, offset int, v uint64) int {
	offset -= sovVote(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Vote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PublicKey)
	if l > 0 {
		n += 1 + l + sovVote(uint64(l))
	}
	if m.Vote != 0 {
		n += 1 + sovVote(uint64(m.Vote))
	}
	return n
}

func sovVote(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVote(x uint64) (n int) {
	return sovVote(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Vote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVote
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
			return fmt.Errorf("proto: Vote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Vote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVote
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
				return ErrInvalidLengthVote
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vote", wireType)
			}
			m.Vote = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Vote |= ProposalVote(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVote(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVote
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
func skipVote(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVote
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
					return 0, ErrIntOverflowVote
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
					return 0, ErrIntOverflowVote
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
				return 0, ErrInvalidLengthVote
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVote
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVote
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVote        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVote          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVote = fmt.Errorf("proto: unexpected end of group")
)
