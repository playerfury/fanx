// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fanx/bet/odds_type.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
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

// OddsType is the representation of the type of the odds.
type OddsType int32

const (
	// invalid odds type
	OddsType_ODDS_TYPE_UNSPECIFIED OddsType = 0
	// decimal odds type (european)
	OddsType_ODDS_TYPE_DECIMAL OddsType = 1
	// fractional odds type (british)
	OddsType_ODDS_TYPE_FRACTIONAL OddsType = 2
	// moneyline odds type (american)
	OddsType_ODDS_TYPE_MONEYLINE OddsType = 3
)

var OddsType_name = map[int32]string{
	0: "ODDS_TYPE_UNSPECIFIED",
	1: "ODDS_TYPE_DECIMAL",
	2: "ODDS_TYPE_FRACTIONAL",
	3: "ODDS_TYPE_MONEYLINE",
}

var OddsType_value = map[string]int32{
	"ODDS_TYPE_UNSPECIFIED": 0,
	"ODDS_TYPE_DECIMAL":     1,
	"ODDS_TYPE_FRACTIONAL":  2,
	"ODDS_TYPE_MONEYLINE":   3,
}

func (x OddsType) String() string {
	return proto.EnumName(OddsType_name, int32(x))
}

func (OddsType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e8f071f6e85b8e02, []int{0}
}

func init() {
	proto.RegisterEnum("fanx.bet.OddsType", OddsType_name, OddsType_value)
}

func init() { proto.RegisterFile("fanx/bet/odds_type.proto", fileDescriptor_e8f071f6e85b8e02) }

var fileDescriptor_e8f071f6e85b8e02 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0x4b, 0xcc, 0xab,
	0xd0, 0x4f, 0x4a, 0x2d, 0xd1, 0xcf, 0x4f, 0x49, 0x29, 0x8e, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0xc9, 0xe8, 0x25, 0xa5, 0x96, 0x68, 0xe5, 0x73, 0x71,
	0xf8, 0xa7, 0xa4, 0x14, 0x87, 0x54, 0x16, 0xa4, 0x0a, 0x49, 0x72, 0x89, 0xfa, 0xbb, 0xb8, 0x04,
	0xc7, 0x87, 0x44, 0x06, 0xb8, 0xc6, 0x87, 0xfa, 0x05, 0x07, 0xb8, 0x3a, 0x7b, 0xba, 0x79, 0xba,
	0xba, 0x08, 0x30, 0x08, 0x89, 0x72, 0x09, 0x22, 0xa4, 0x5c, 0x5c, 0x9d, 0x3d, 0x7d, 0x1d, 0x7d,
	0x04, 0x18, 0x85, 0x24, 0xb8, 0x44, 0x10, 0xc2, 0x6e, 0x41, 0x8e, 0xce, 0x21, 0x9e, 0xfe, 0x7e,
	0x8e, 0x3e, 0x02, 0x4c, 0x42, 0xe2, 0x5c, 0xc2, 0x08, 0x19, 0x5f, 0x7f, 0x3f, 0xd7, 0x48, 0x1f,
	0x4f, 0x3f, 0x57, 0x01, 0x66, 0x27, 0x87, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c,
	0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63,
	0x88, 0x52, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xc8, 0x49,
	0xac, 0x4c, 0x2d, 0x4a, 0x2b, 0x2d, 0xaa, 0xd4, 0x07, 0x7b, 0x02, 0xe2, 0x0d, 0x90, 0x0f, 0x8a,
	0x93, 0xd8, 0xc0, 0x7e, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x03, 0x89, 0xb2, 0x90, 0xdf,
	0x00, 0x00, 0x00,
}