// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fanx/bet/genesis.proto

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

// GenesisState defines the bet module's genesis state.
type GenesisState struct {
	// params contains parameters of bet module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// bet_list contains the bet list in the genesis init.
	BetList []Bet `protobuf:"bytes,2,rep,name=bet_list,json=betList,proto3" json:"bet_list"`
	// pending_bet_list contains the pending bet list in the genesis init.
	PendingBetList []PendingBet `protobuf:"bytes,3,rep,name=pending_bet_list,json=pendingBetList,proto3" json:"pending_bet_list"`
	// settled_bet_list contains the settled bet list in the genesis init.
	SettledBetList []SettledBet `protobuf:"bytes,4,rep,name=settled_bet_list,json=settledBetList,proto3" json:"settled_bet_list"`
	// uid2id_list contains bet to id list in the genesis init.
	Uid2IdList []UID2ID `protobuf:"bytes,5,rep,name=uid2id_list,json=uid2idList,proto3" json:"uid2id_list"`
	// stats contains statistics in the genesis init.
	Stats BetStats `protobuf:"bytes,6,opt,name=stats,proto3" json:"stats"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c49ebc0f2678a09, []int{0}
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

func (m *GenesisState) GetBetList() []Bet {
	if m != nil {
		return m.BetList
	}
	return nil
}

func (m *GenesisState) GetPendingBetList() []PendingBet {
	if m != nil {
		return m.PendingBetList
	}
	return nil
}

func (m *GenesisState) GetSettledBetList() []SettledBet {
	if m != nil {
		return m.SettledBetList
	}
	return nil
}

func (m *GenesisState) GetUid2IdList() []UID2ID {
	if m != nil {
		return m.Uid2IdList
	}
	return nil
}

func (m *GenesisState) GetStats() BetStats {
	if m != nil {
		return m.Stats
	}
	return BetStats{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "fanx.bet.GenesisState")
}

func init() { proto.RegisterFile("fanx/bet/genesis.proto", fileDescriptor_6c49ebc0f2678a09) }

var fileDescriptor_6c49ebc0f2678a09 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x93, 0xaf, 0x7f, 0x3e, 0x99, 0x8a, 0x68, 0xac, 0x58, 0x82, 0x8c, 0xc5, 0x85, 0x74,
	0x63, 0x02, 0x71, 0xd3, 0xa5, 0x96, 0x82, 0x14, 0x44, 0xc4, 0xe2, 0xc6, 0x4d, 0xc9, 0x98, 0xcb,
	0x38, 0xd8, 0x76, 0x42, 0xe7, 0x16, 0xf5, 0x2d, 0x7c, 0x03, 0x5f, 0xa7, 0xcb, 0x2e, 0x5d, 0x89,
	0xb4, 0x2f, 0x22, 0x99, 0x99, 0xb4, 0x82, 0xe9, 0xee, 0xe6, 0xe4, 0x9c, 0xdf, 0x1d, 0xce, 0x25,
	0x07, 0x8a, 0x43, 0xc8, 0x00, 0x43, 0x0e, 0x63, 0x50, 0x42, 0x05, 0xe9, 0x44, 0xa2, 0xf4, 0x3c,
	0x95, 0x7d, 0xe3, 0x8b, 0x9c, 0x3c, 0x07, 0x8a, 0x43, 0xc0, 0x00, 0xfd, 0x3a, 0x97, 0x5c, 0xea,
	0xdf, 0x61, 0x36, 0x19, 0xa7, 0x5f, 0xcf, 0x01, 0x69, 0x3c, 0x89, 0x47, 0x36, 0xef, 0xef, 0xe5,
	0x2a, 0x03, 0xb4, 0xd2, 0x7e, 0x2e, 0x29, 0x8c, 0xd1, 0xfa, 0x4e, 0x3e, 0x4a, 0x64, 0xfb, 0xca,
	0x6c, 0xee, 0x63, 0x8c, 0xe0, 0xb5, 0x49, 0xd5, 0x80, 0x1a, 0x6e, 0xd3, 0x6d, 0xd5, 0x22, 0x3f,
	0xf8, 0xfb, 0x92, 0xe0, 0x56, 0x3b, 0x3a, 0xe5, 0xd9, 0xd7, 0xb1, 0x73, 0x67, 0xfd, 0x5e, 0x9b,
	0x6c, 0x31, 0xc0, 0xc1, 0x50, 0x28, 0x6c, 0xfc, 0x6b, 0x96, 0x5a, 0xb5, 0xe8, 0xb0, 0x28, 0xdb,
	0x01, 0xb4, 0xc1, 0xff, 0x0c, 0xf0, 0x5a, 0x28, 0xf4, 0x6e, 0xc8, 0x6e, 0x0a, 0xe3, 0x44, 0x8c,
	0xf9, 0x60, 0x45, 0x28, 0x69, 0x02, 0x2d, 0xdc, 0x6e, 0xbc, 0x6b, 0xd0, 0x4e, 0xba, 0x52, 0x72,
	0x9e, 0x02, 0xc4, 0x21, 0x24, 0x6b, 0x5e, 0x79, 0x33, 0xaf, 0x6f, 0xbc, 0xbf, 0x78, 0x6a, 0xa5,
	0x68, 0xde, 0x25, 0xa9, 0x4d, 0x45, 0x12, 0x89, 0xc4, 0xa0, 0x2a, 0x1a, 0x55, 0x58, 0xcc, 0x7d,
	0xaf, 0x1b, 0xf5, 0xba, 0x16, 0x43, 0x4c, 0x48, 0x23, 0xda, 0xa4, 0xa2, 0x6b, 0x6f, 0x54, 0x75,
	0xab, 0x47, 0x1b, 0x9a, 0xc9, 0x6e, 0x90, 0xf7, 0x6a, 0x02, 0x9d, 0x8b, 0xd9, 0x82, 0xba, 0xf3,
	0x05, 0x75, 0xbf, 0x17, 0xd4, 0x7d, 0x5f, 0x52, 0x67, 0xbe, 0xa4, 0xce, 0xe7, 0x92, 0x3a, 0x0f,
	0xa7, 0x5c, 0xe0, 0xd3, 0x94, 0x05, 0x8f, 0x72, 0x14, 0x2a, 0x0e, 0x67, 0x96, 0x97, 0xcd, 0xe1,
	0xab, 0xbe, 0x34, 0xbe, 0xa5, 0xa0, 0x58, 0x55, 0x9f, 0xfa, 0xfc, 0x27, 0x00, 0x00, 0xff, 0xff,
	0xe5, 0xda, 0xaa, 0x89, 0x6b, 0x02, 0x00, 0x00,
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
	dAtA[i] = 0x32
	if len(m.Uid2IdList) > 0 {
		for iNdEx := len(m.Uid2IdList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Uid2IdList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.SettledBetList) > 0 {
		for iNdEx := len(m.SettledBetList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SettledBetList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.PendingBetList) > 0 {
		for iNdEx := len(m.PendingBetList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PendingBetList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.BetList) > 0 {
		for iNdEx := len(m.BetList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BetList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.BetList) > 0 {
		for _, e := range m.BetList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PendingBetList) > 0 {
		for _, e := range m.PendingBetList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.SettledBetList) > 0 {
		for _, e := range m.SettledBetList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Uid2IdList) > 0 {
		for _, e := range m.Uid2IdList {
			l = e.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field BetList", wireType)
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
			m.BetList = append(m.BetList, Bet{})
			if err := m.BetList[len(m.BetList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingBetList", wireType)
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
			m.PendingBetList = append(m.PendingBetList, PendingBet{})
			if err := m.PendingBetList[len(m.PendingBetList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SettledBetList", wireType)
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
			m.SettledBetList = append(m.SettledBetList, SettledBet{})
			if err := m.SettledBetList[len(m.SettledBetList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid2IdList", wireType)
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
			m.Uid2IdList = append(m.Uid2IdList, UID2ID{})
			if err := m.Uid2IdList[len(m.Uid2IdList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
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
