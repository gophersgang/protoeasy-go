// Code generated by protoc-gen-go.
// source: pb/type/wrappers.proto
// DO NOT EDIT!

package pbtype

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DoubleValues struct {
	Value []float64 `protobuf:"fixed64,1,rep,name=value" json:"value,omitempty"`
}

func (m *DoubleValues) Reset()                    { *m = DoubleValues{} }
func (m *DoubleValues) String() string            { return proto.CompactTextString(m) }
func (*DoubleValues) ProtoMessage()               {}
func (*DoubleValues) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

type FloatValues struct {
	Value []float32 `protobuf:"fixed32,1,rep,name=value" json:"value,omitempty"`
}

func (m *FloatValues) Reset()                    { *m = FloatValues{} }
func (m *FloatValues) String() string            { return proto.CompactTextString(m) }
func (*FloatValues) ProtoMessage()               {}
func (*FloatValues) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

type Int64Values struct {
	Value []int64 `protobuf:"varint,1,rep,name=value" json:"value,omitempty"`
}

func (m *Int64Values) Reset()                    { *m = Int64Values{} }
func (m *Int64Values) String() string            { return proto.CompactTextString(m) }
func (*Int64Values) ProtoMessage()               {}
func (*Int64Values) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

type UInt64Values struct {
	Value []uint64 `protobuf:"varint,1,rep,name=value" json:"value,omitempty"`
}

func (m *UInt64Values) Reset()                    { *m = UInt64Values{} }
func (m *UInt64Values) String() string            { return proto.CompactTextString(m) }
func (*UInt64Values) ProtoMessage()               {}
func (*UInt64Values) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{3} }

type Int32Values struct {
	Value []int32 `protobuf:"varint,1,rep,name=value" json:"value,omitempty"`
}

func (m *Int32Values) Reset()                    { *m = Int32Values{} }
func (m *Int32Values) String() string            { return proto.CompactTextString(m) }
func (*Int32Values) ProtoMessage()               {}
func (*Int32Values) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{4} }

type UInt32Values struct {
	Value []uint32 `protobuf:"varint,1,rep,name=value" json:"value,omitempty"`
}

func (m *UInt32Values) Reset()                    { *m = UInt32Values{} }
func (m *UInt32Values) String() string            { return proto.CompactTextString(m) }
func (*UInt32Values) ProtoMessage()               {}
func (*UInt32Values) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{5} }

type BoolValues struct {
	Value []bool `protobuf:"varint,1,rep,name=value" json:"value,omitempty"`
}

func (m *BoolValues) Reset()                    { *m = BoolValues{} }
func (m *BoolValues) String() string            { return proto.CompactTextString(m) }
func (*BoolValues) ProtoMessage()               {}
func (*BoolValues) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{6} }

type StringValues struct {
	Value []string `protobuf:"bytes,1,rep,name=value" json:"value,omitempty"`
}

func (m *StringValues) Reset()                    { *m = StringValues{} }
func (m *StringValues) String() string            { return proto.CompactTextString(m) }
func (*StringValues) ProtoMessage()               {}
func (*StringValues) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{7} }

type BytesValues struct {
	Value [][]byte `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (m *BytesValues) Reset()                    { *m = BytesValues{} }
func (m *BytesValues) String() string            { return proto.CompactTextString(m) }
func (*BytesValues) ProtoMessage()               {}
func (*BytesValues) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{8} }

func init() {
	proto.RegisterType((*DoubleValues)(nil), "pb.type.DoubleValues")
	proto.RegisterType((*FloatValues)(nil), "pb.type.FloatValues")
	proto.RegisterType((*Int64Values)(nil), "pb.type.Int64Values")
	proto.RegisterType((*UInt64Values)(nil), "pb.type.UInt64Values")
	proto.RegisterType((*Int32Values)(nil), "pb.type.Int32Values")
	proto.RegisterType((*UInt32Values)(nil), "pb.type.UInt32Values")
	proto.RegisterType((*BoolValues)(nil), "pb.type.BoolValues")
	proto.RegisterType((*StringValues)(nil), "pb.type.StringValues")
	proto.RegisterType((*BytesValues)(nil), "pb.type.BytesValues")
}

var fileDescriptor8 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x48, 0xd2, 0x2f,
	0xa9, 0x2c, 0x48, 0xd5, 0x2f, 0x2f, 0x4a, 0x2c, 0x28, 0x48, 0x2d, 0x2a, 0xd6, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x48, 0xd2, 0x03, 0x89, 0x2b, 0xc9, 0x72, 0xf1, 0xb8, 0xe4, 0x97,
	0x26, 0xe5, 0xa4, 0x86, 0x25, 0xe6, 0x94, 0xa6, 0x16, 0x0b, 0xf1, 0x72, 0xb1, 0x96, 0x81, 0x58,
	0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0x8c, 0x4a, 0x32, 0x5c, 0xdc, 0x6e, 0x39, 0xf9, 0x89, 0x25, 0xd8,
	0x64, 0x99, 0x40, 0xb2, 0x9e, 0x79, 0x25, 0x66, 0x26, 0xd8, 0x64, 0x99, 0x41, 0x46, 0x87, 0xe2,
	0x94, 0x66, 0x81, 0x6a, 0x36, 0x36, 0xc2, 0x26, 0xcb, 0x0a, 0xd3, 0x8c, 0x5d, 0x9a, 0x57, 0x49,
	0x9a, 0x8b, 0xcb, 0x29, 0x3f, 0x3f, 0x07, 0x9b, 0x24, 0x07, 0x48, 0x6f, 0x70, 0x49, 0x51, 0x66,
	0x5e, 0x3a, 0x36, 0x69, 0x4e, 0x90, 0xc5, 0x4e, 0x95, 0x25, 0xa9, 0xc5, 0xd8, 0x64, 0x79, 0x9c,
	0x38, 0xa2, 0xd8, 0x0a, 0x92, 0x40, 0x41, 0x93, 0xc4, 0x06, 0x0e, 0x2a, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xf9, 0xc9, 0xc1, 0xdf, 0x44, 0x01, 0x00, 0x00,
}
