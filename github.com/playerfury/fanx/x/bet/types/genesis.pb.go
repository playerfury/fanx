// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fanx/bet/genesis.proto

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
	return fileDescriptor_073b1176f79d0704, []int{0}
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

func init() { proto.RegisterFile("fanx/bet/genesis.proto", fileDescriptor_073b1176f79d0704) }

var fileDescriptor_073b1176f79d0704 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0xdb, 0xfd, 0x73, 0x64, 0x2a, 0xa3, 0x4c, 0x19, 0x3b, 0xd4, 0xe1, 0x41, 0x76, 0x4a,
	0xa1, 0x1e, 0xbc, 0xca, 0x18, 0xc8, 0xc0, 0x83, 0x38, 0xbc, 0x78, 0x19, 0x89, 0x7d, 0x57, 0x03,
	0x5b, 0x1b, 0x9a, 0x77, 0xe0, 0xbe, 0x85, 0x1f, 0x6b, 0xc7, 0x1d, 0x3d, 0x89, 0xac, 0x5f, 0x44,
	0x96, 0x64, 0xd9, 0xf4, 0x56, 0x9e, 0x3e, 0xbf, 0x5f, 0x92, 0xf7, 0x25, 0x97, 0x33, 0x96, 0x7d,
	0x44, 0x1c, 0x30, 0x4a, 0x21, 0x03, 0x25, 0x14, 0x95, 0x45, 0x8e, 0x79, 0xd0, 0xdc, 0xe5, 0x94,
	0x03, 0xf6, 0x3a, 0x69, 0x9e, 0xe6, 0x3a, 0x8c, 0x76, 0x5f, 0xe6, 0x7f, 0xef, 0xc2, 0x71, 0x92,
	0x15, 0x6c, 0x61, 0xb1, 0x5e, 0xe0, 0x62, 0x0e, 0x68, 0xb3, 0x8e, 0xcb, 0x14, 0x32, 0xb4, 0xcd,
	0xeb, 0xb2, 0x42, 0x4e, 0x1f, 0xcc, 0x91, 0x13, 0x64, 0x08, 0x01, 0x25, 0x0d, 0xa3, 0xea, 0xfa,
	0x7d, 0x7f, 0xd0, 0x8a, 0xdb, 0x74, 0x7f, 0x05, 0xfa, 0xa4, 0xf3, 0x61, 0x6d, 0xfd, 0x7d, 0xe5,
	0x3d, 0xdb, 0x56, 0x40, 0x49, 0x93, 0x03, 0x4e, 0xe7, 0x42, 0x61, 0xb7, 0xd2, 0xaf, 0x0e, 0x5a,
	0xf1, 0xd9, 0x81, 0x18, 0x02, 0xda, 0xfa, 0x09, 0x07, 0x7c, 0x14, 0x0a, 0x83, 0x11, 0x69, 0x4b,
	0xc8, 0x12, 0x91, 0xa5, 0x53, 0xc7, 0x55, 0x35, 0xd7, 0x39, 0x3a, 0xc9, 0x34, 0x0e, 0xf8, 0xb9,
	0x74, 0xc9, 0xde, 0xa2, 0x00, 0x71, 0x0e, 0xc9, 0xc1, 0x52, 0xfb, 0x6f, 0x99, 0x98, 0xc6, 0x91,
	0x45, 0xb9, 0x44, 0x5b, 0xee, 0x48, 0x6b, 0x29, 0x92, 0x58, 0x24, 0x46, 0x50, 0xd7, 0x82, 0xa3,
	0x07, 0xbf, 0x8c, 0x47, 0xf1, 0x78, 0x64, 0x61, 0x62, 0xaa, 0x1a, 0xa4, 0xa4, 0xae, 0x87, 0xd8,
	0x6d, 0xe8, 0x19, 0x05, 0x7f, 0x5e, 0xbc, 0x9b, 0xe3, 0x7e, 0x4a, 0xa6, 0x36, 0xbc, 0x5f, 0x6f,
	0x43, 0x7f, 0xb3, 0x0d, 0xfd, 0x9f, 0x6d, 0xe8, 0x7f, 0x96, 0xa1, 0xb7, 0x29, 0x43, 0xef, 0xab,
	0x0c, 0xbd, 0xd7, 0x9b, 0x54, 0xe0, 0xfb, 0x92, 0xd3, 0xb7, 0x7c, 0x11, 0xc9, 0x39, 0x5b, 0x41,
	0x31, 0x5b, 0x16, 0xab, 0x48, 0xef, 0xca, 0x6c, 0x0b, 0x57, 0x12, 0x14, 0x6f, 0xe8, 0x75, 0xdd,
	0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x06, 0x25, 0xea, 0xac, 0x29, 0x02, 0x00, 0x00,
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