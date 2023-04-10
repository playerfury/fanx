// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: furyx/dvm/ticket.proto

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

// PubkeysChangeProposalPayload indicates data of public keys changes proposal
// ticket.
type PubkeysChangeProposalPayload struct {
	// public_keys contain new pub keys to be added to public keys.
	PublicKeys []string `protobuf:"bytes,1,rep,name=public_keys,json=publicKeys,proto3" json:"public_keys,omitempty"`
	// leader_index is the universal unique identifier of the public key.
	LeaderIndex uint32 `protobuf:"varint,2,opt,name=leader_index,json=leaderIndex,proto3" json:"leader_index,omitempty"`
}

func (m *PubkeysChangeProposalPayload) Reset()         { *m = PubkeysChangeProposalPayload{} }
func (m *PubkeysChangeProposalPayload) String() string { return proto.CompactTextString(m) }
func (*PubkeysChangeProposalPayload) ProtoMessage()    {}
func (*PubkeysChangeProposalPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_7121f0f264448d06, []int{0}
}
func (m *PubkeysChangeProposalPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PubkeysChangeProposalPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PubkeysChangeProposalPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PubkeysChangeProposalPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubkeysChangeProposalPayload.Merge(m, src)
}
func (m *PubkeysChangeProposalPayload) XXX_Size() int {
	return m.Size()
}
func (m *PubkeysChangeProposalPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_PubkeysChangeProposalPayload.DiscardUnknown(m)
}

var xxx_messageInfo_PubkeysChangeProposalPayload proto.InternalMessageInfo

func (m *PubkeysChangeProposalPayload) GetPublicKeys() []string {
	if m != nil {
		return m.PublicKeys
	}
	return nil
}

func (m *PubkeysChangeProposalPayload) GetLeaderIndex() uint32 {
	if m != nil {
		return m.LeaderIndex
	}
	return 0
}

// ProposalVotePayload indicates vote data ticket.
type ProposalVotePayload struct {
	// proposal_id is the id of the proposal.
	ProposalId uint64 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	// vote is the chosen option for the vote.
	Vote ProposalVote `protobuf:"varint,2,opt,name=vote,proto3,enum=furyx.furyx.dvm.ProposalVote" json:"vote,omitempty"`
}

func (m *ProposalVotePayload) Reset()         { *m = ProposalVotePayload{} }
func (m *ProposalVotePayload) String() string { return proto.CompactTextString(m) }
func (*ProposalVotePayload) ProtoMessage()    {}
func (*ProposalVotePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_7121f0f264448d06, []int{1}
}
func (m *ProposalVotePayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProposalVotePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProposalVotePayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProposalVotePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalVotePayload.Merge(m, src)
}
func (m *ProposalVotePayload) XXX_Size() int {
	return m.Size()
}
func (m *ProposalVotePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalVotePayload.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalVotePayload proto.InternalMessageInfo

func (m *ProposalVotePayload) GetProposalId() uint64 {
	if m != nil {
		return m.ProposalId
	}
	return 0
}

func (m *ProposalVotePayload) GetVote() ProposalVote {
	if m != nil {
		return m.Vote
	}
	return ProposalVote_PROPOSAL_VOTE_UNSPECIFIED
}

func init() {
	proto.RegisterType((*PubkeysChangeProposalPayload)(nil), "furyx.furyx.dvm.PubkeysChangeProposalPayload")
	proto.RegisterType((*ProposalVotePayload)(nil), "furyx.furyx.dvm.ProposalVotePayload")
}

func init() { proto.RegisterFile("furyx/dvm/ticket.proto", fileDescriptor_7121f0f264448d06) }

var fileDescriptor_7121f0f264448d06 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x41, 0x4b, 0xb4, 0x40,
	0x1c, 0x87, 0x9d, 0xf7, 0x5d, 0x82, 0x66, 0xab, 0xc3, 0xd4, 0x41, 0x22, 0x26, 0xdb, 0x43, 0x78,
	0x69, 0x84, 0xea, 0x03, 0x44, 0x9d, 0x96, 0x2e, 0xe2, 0xa1, 0x43, 0x17, 0x51, 0xe7, 0xcf, 0xac,
	0x38, 0x3a, 0xe2, 0x8c, 0xb6, 0x7e, 0x8b, 0x3e, 0x56, 0xc7, 0x3d, 0x76, 0x0c, 0xfd, 0x22, 0xa1,
	0xae, 0xb0, 0xd0, 0x6d, 0x78, 0x78, 0x78, 0x98, 0xff, 0x0f, 0x5f, 0x68, 0x01, 0x1e, 0x6f, 0x72,
	0xcf, 0xa4, 0x49, 0x06, 0x86, 0x95, 0x95, 0x32, 0x8a, 0x10, 0x2d, 0xa0, 0x00, 0xf3, 0xa1, 0xaa,
	0x8c, 0x69, 0x01, 0x8c, 0x37, 0xf9, 0x25, 0x99, 0xcd, 0x46, 0x19, 0x98, 0xbc, 0x55, 0x8c, 0xaf,
	0xfc, 0x3a, 0xce, 0xa0, 0xd5, 0x2f, 0x9b, 0xa8, 0x10, 0xe0, 0x57, 0xaa, 0x54, 0x3a, 0x92, 0x7e,
	0xd4, 0x4a, 0x15, 0x71, 0x72, 0x8d, 0x97, 0x65, 0x1d, 0xcb, 0x34, 0x09, 0x07, 0xc7, 0x46, 0xce,
	0x7f, 0xf7, 0x38, 0xc0, 0x13, 0x7a, 0x85, 0x56, 0x93, 0x1b, 0x7c, 0x22, 0x21, 0xe2, 0x50, 0x85,
	0x69, 0xc1, 0x61, 0x6b, 0xff, 0x73, 0x90, 0x7b, 0x1a, 0x2c, 0x27, 0xb6, 0x1e, 0xd0, 0x4a, 0xe2,
	0xf3, 0x39, 0xfb, 0xa6, 0x0c, 0x1c, 0xa6, 0xf7, 0x38, 0x4c, 0xb9, 0x8d, 0x1c, 0xe4, 0x2e, 0x02,
	0x3c, 0xa3, 0x35, 0x27, 0x8f, 0x78, 0x31, 0xfc, 0x74, 0x4c, 0x9e, 0xdd, 0x3b, 0xec, 0xef, 0x49,
	0xec, 0xb0, 0x1b, 0x8c, 0xf6, 0xf3, 0xd3, 0x57, 0x47, 0xd1, 0xae, 0xa3, 0xe8, 0xa7, 0xa3, 0xe8,
	0xb3, 0xa7, 0xd6, 0xae, 0xa7, 0xd6, 0x77, 0x4f, 0xad, 0xf7, 0x5b, 0x91, 0x9a, 0x4d, 0x1d, 0xb3,
	0x44, 0xe5, 0x9e, 0x16, 0x70, 0xb7, 0x8f, 0x0d, 0x6f, 0x6f, 0x3b, 0x4d, 0xd8, 0x96, 0xa0, 0xe3,
	0xa3, 0x71, 0x9a, 0x87, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x94, 0x3e, 0x43, 0x5a, 0x01,
	0x00, 0x00,
}

func (m *PubkeysChangeProposalPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PubkeysChangeProposalPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PubkeysChangeProposalPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LeaderIndex != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.LeaderIndex))
		i--
		dAtA[i] = 0x10
	}
	if len(m.PublicKeys) > 0 {
		for iNdEx := len(m.PublicKeys) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PublicKeys[iNdEx])
			copy(dAtA[i:], m.PublicKeys[iNdEx])
			i = encodeVarintTicket(dAtA, i, uint64(len(m.PublicKeys[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ProposalVotePayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProposalVotePayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProposalVotePayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Vote != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.Vote))
		i--
		dAtA[i] = 0x10
	}
	if m.ProposalId != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.ProposalId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTicket(dAtA []byte, offset int, v uint64) int {
	offset -= sovTicket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PubkeysChangeProposalPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PublicKeys) > 0 {
		for _, s := range m.PublicKeys {
			l = len(s)
			n += 1 + l + sovTicket(uint64(l))
		}
	}
	if m.LeaderIndex != 0 {
		n += 1 + sovTicket(uint64(m.LeaderIndex))
	}
	return n
}

func (m *ProposalVotePayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalId != 0 {
		n += 1 + sovTicket(uint64(m.ProposalId))
	}
	if m.Vote != 0 {
		n += 1 + sovTicket(uint64(m.Vote))
	}
	return n
}

func sovTicket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTicket(x uint64) (n int) {
	return sovTicket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PubkeysChangeProposalPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTicket
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
			return fmt.Errorf("proto: PubkeysChangeProposalPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PubkeysChangeProposalPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKeys", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKeys = append(m.PublicKeys, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaderIndex", wireType)
			}
			m.LeaderIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LeaderIndex |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTicket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTicket
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
func (m *ProposalVotePayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTicket
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
			return fmt.Errorf("proto: ProposalVotePayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProposalVotePayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
			}
			m.ProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vote", wireType)
			}
			m.Vote = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
			skippy, err := skipTicket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTicket
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
func skipTicket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTicket
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
					return 0, ErrIntOverflowTicket
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
					return 0, ErrIntOverflowTicket
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
				return 0, ErrInvalidLengthTicket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTicket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTicket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTicket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTicket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTicket = fmt.Errorf("proto: unexpected end of group")
)
