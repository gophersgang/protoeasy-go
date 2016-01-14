// Code generated by protoc-gen-go.
// source: pb/type/money.proto
// DO NOT EDIT!

package pbtype

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Money struct {
	CurrencyCode CurrencyCode `protobuf:"varint,1,opt,name=currency_code,enum=pb.type.CurrencyCode" json:"currency_code,omitempty"`
	ValueMicros  int64        `protobuf:"varint,2,opt,name=value_micros" json:"value_micros,omitempty"`
}

func (m *Money) Reset()                    { *m = Money{} }
func (m *Money) String() string            { return proto.CompactTextString(m) }
func (*Money) ProtoMessage()               {}
func (*Money) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func init() {
	proto.RegisterType((*Money)(nil), "pb.type.Money")
}

var fileDescriptor4 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x48, 0xd2, 0x2f,
	0xa9, 0x2c, 0x48, 0xd5, 0xcf, 0xcd, 0xcf, 0x4b, 0xad, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x2f, 0x48, 0xd2, 0x03, 0x09, 0x4a, 0x49, 0xc2, 0x64, 0x93, 0xf3, 0x4b, 0xf3, 0x4a, 0x8a,
	0x2a, 0xf5, 0xd2, 0x53, 0xf3, 0x20, 0x6a, 0x94, 0xbc, 0xb9, 0x58, 0x7d, 0x41, 0x5a, 0x84, 0x74,
	0xb8, 0x78, 0x93, 0x4b, 0x8b, 0x8a, 0x52, 0xf3, 0x92, 0x2b, 0xe3, 0x93, 0xf3, 0x53, 0x52, 0x25,
	0x18, 0x15, 0x18, 0x35, 0xf8, 0x8c, 0x44, 0xf5, 0xa0, 0x86, 0xe8, 0x39, 0x43, 0x65, 0x9d, 0x81,
	0x92, 0x42, 0x22, 0x5c, 0x3c, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0xf1, 0xb9, 0x99, 0xc9, 0x45, 0xf9,
	0xc5, 0x12, 0x4c, 0x40, 0xc5, 0xcc, 0x4e, 0x1c, 0x51, 0x6c, 0x05, 0x49, 0x20, 0xc5, 0x49, 0x6c,
	0x60, 0xd3, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xdb, 0xc9, 0x52, 0x5e, 0x98, 0x00, 0x00,
	0x00,
}