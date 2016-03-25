// Code generated by protoc-gen-gogo.
// source: google/type/money.proto
// DO NOT EDIT!

package google_type

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Represents an amount of money with its currency type.
type Money struct {
	// The 3-letter currency code defined in ISO 4217.
	CurrencyCode string `protobuf:"bytes,1,opt,name=currency_code,proto3" json:"currency_code,omitempty"`
	// The whole units of the amount.
	// For example if `currencyCode` is `"USD"`, then 1 unit is one US dollar.
	Units int64 `protobuf:"varint,2,opt,name=units,proto3" json:"units,omitempty"`
	// Number of nano (10^-9) units of the amount.
	// The value must be between -999,999,999 and +999,999,999 inclusive.
	// If `units` is positive, `nanos` must be positive or zero.
	// If `units` is zero, `nanos` can be positive, zero, or negative.
	// If `units` is negative, `nanos` must be negative or zero.
	// For example $-1.75 is represented as `units`=-1 and `nanos`=-750,000,000.
	Nanos int32 `protobuf:"varint,3,opt,name=nanos,proto3" json:"nanos,omitempty"`
}

func (m *Money) Reset()                    { *m = Money{} }
func (m *Money) String() string            { return proto.CompactTextString(m) }
func (*Money) ProtoMessage()               {}
func (*Money) Descriptor() ([]byte, []int) { return fileDescriptorMoney, []int{0} }

func init() {
	proto.RegisterType((*Money)(nil), "google.type.Money")
}

var fileDescriptorMoney = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0xcf, 0xcd, 0xcf, 0x4b, 0xad, 0xd4, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x86, 0x48, 0xe8, 0x81, 0x24, 0x94, 0x6c, 0xb8, 0x58, 0x7d, 0x41,
	0x72, 0x42, 0xa2, 0x5c, 0xbc, 0xc9, 0xa5, 0x45, 0x45, 0xa9, 0x79, 0xc9, 0x95, 0xf1, 0xc9, 0xf9,
	0x29, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x42, 0xbc, 0x5c, 0xac, 0xa5, 0x79, 0x99, 0x25,
	0xc5, 0x12, 0x4c, 0x40, 0x2e, 0x33, 0x88, 0x9b, 0x97, 0x98, 0x97, 0x5f, 0x2c, 0xc1, 0x0c, 0xe4,
	0xb2, 0x3a, 0xc9, 0x73, 0xf1, 0x27, 0xe7, 0xe7, 0xea, 0x21, 0x19, 0xe8, 0xc4, 0x05, 0x36, 0x2e,
	0x00, 0x64, 0x53, 0x00, 0x63, 0x12, 0x1b, 0xd8, 0x4a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x38, 0x21, 0x56, 0x5a, 0x8d, 0x00, 0x00, 0x00,
}
