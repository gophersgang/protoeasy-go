// Code generated by protoc-gen-gogo.
// source: google/longrunning/operations.proto
// DO NOT EDIT!

package google_longrunning

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gengo/grpc-gateway/third_party/googleapis/google/api"
import google_protobuf1 "go.pedge.io/pb/gogo/google/protobuf"
import _ "go.pedge.io/pb/gogo/google/protobuf"
import google_rpc "go.pedge.io/pb/gogo/google/rpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

// This resource represents a long-running operation that is the result of a
// network API call.
type Operation struct {
	// The name of the operation resource, which is only unique within the same
	// service that originally returns it.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Some service-specific metadata associated with the operation.  It typically
	// contains progress information and common metadata such as create time.
	// Some services may not provide such metadata.  Any method that returns a
	// long-running operation should document the metadata type, if any.
	Metadata *google_protobuf1.Any `protobuf:"bytes,2,opt,name=metadata" json:"metadata,omitempty"`
	// If the value is false, it means the operation is still in progress.
	// If true, the operation is completed and the `result` is available.
	Done bool `protobuf:"varint,3,opt,name=done,proto3" json:"done,omitempty"`
	// Types that are valid to be assigned to Result:
	//	*Operation_Error
	//	*Operation_Response
	Result isOperation_Result `protobuf_oneof:"result"`
}

func (m *Operation) Reset()                    { *m = Operation{} }
func (m *Operation) String() string            { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()               {}
func (*Operation) Descriptor() ([]byte, []int) { return fileDescriptorOperations, []int{0} }

type isOperation_Result interface {
	isOperation_Result()
}

type Operation_Error struct {
	Error *google_rpc.Status `protobuf:"bytes,4,opt,name=error,oneof"`
}
type Operation_Response struct {
	Response *google_protobuf1.Any `protobuf:"bytes,5,opt,name=response,oneof"`
}

func (*Operation_Error) isOperation_Result()    {}
func (*Operation_Response) isOperation_Result() {}

func (m *Operation) GetResult() isOperation_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *Operation) GetMetadata() *google_protobuf1.Any {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Operation) GetError() *google_rpc.Status {
	if x, ok := m.GetResult().(*Operation_Error); ok {
		return x.Error
	}
	return nil
}

func (m *Operation) GetResponse() *google_protobuf1.Any {
	if x, ok := m.GetResult().(*Operation_Response); ok {
		return x.Response
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Operation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Operation_OneofMarshaler, _Operation_OneofUnmarshaler, _Operation_OneofSizer, []interface{}{
		(*Operation_Error)(nil),
		(*Operation_Response)(nil),
	}
}

func _Operation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Operation)
	// result
	switch x := m.Result.(type) {
	case *Operation_Error:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case *Operation_Response:
		_ = b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Response); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Operation.Result has unexpected type %T", x)
	}
	return nil
}

func _Operation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Operation)
	switch tag {
	case 4: // result.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_rpc.Status)
		err := b.DecodeMessage(msg)
		m.Result = &Operation_Error{msg}
		return true, err
	case 5: // result.response
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf1.Any)
		err := b.DecodeMessage(msg)
		m.Result = &Operation_Response{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Operation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Operation)
	// result
	switch x := m.Result.(type) {
	case *Operation_Error:
		s := proto.Size(x.Error)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_Response:
		s := proto.Size(x.Response)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// The request message for [Operations.GetOperation][google.longrunning.Operations.GetOperation].
type GetOperationRequest struct {
	// The name of the operation resource.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *GetOperationRequest) Reset()                    { *m = GetOperationRequest{} }
func (m *GetOperationRequest) String() string            { return proto.CompactTextString(m) }
func (*GetOperationRequest) ProtoMessage()               {}
func (*GetOperationRequest) Descriptor() ([]byte, []int) { return fileDescriptorOperations, []int{1} }

// The request message for [Operations.ListOperations][google.longrunning.Operations.ListOperations].
type ListOperationsRequest struct {
	// The name of the operation collection.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The standard List filter.
	Filter string `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	// The standard List page size.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,proto3" json:"page_size,omitempty"`
	// The standard List page token.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,proto3" json:"page_token,omitempty"`
}

func (m *ListOperationsRequest) Reset()                    { *m = ListOperationsRequest{} }
func (m *ListOperationsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListOperationsRequest) ProtoMessage()               {}
func (*ListOperationsRequest) Descriptor() ([]byte, []int) { return fileDescriptorOperations, []int{2} }

// The response message for [Operations.ListOperations][google.longrunning.Operations.ListOperations].
type ListOperationsResponse struct {
	// A list of operations that match the specified filter in the request.
	Operations []*Operation `protobuf:"bytes,1,rep,name=operations" json:"operations,omitempty"`
	// The standard List next-page token.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,proto3" json:"next_page_token,omitempty"`
}

func (m *ListOperationsResponse) Reset()                    { *m = ListOperationsResponse{} }
func (m *ListOperationsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListOperationsResponse) ProtoMessage()               {}
func (*ListOperationsResponse) Descriptor() ([]byte, []int) { return fileDescriptorOperations, []int{3} }

func (m *ListOperationsResponse) GetOperations() []*Operation {
	if m != nil {
		return m.Operations
	}
	return nil
}

// The request message for [Operations.CancelOperation][google.longrunning.Operations.CancelOperation].
type CancelOperationRequest struct {
	// The name of the operation resource to be cancelled.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *CancelOperationRequest) Reset()                    { *m = CancelOperationRequest{} }
func (m *CancelOperationRequest) String() string            { return proto.CompactTextString(m) }
func (*CancelOperationRequest) ProtoMessage()               {}
func (*CancelOperationRequest) Descriptor() ([]byte, []int) { return fileDescriptorOperations, []int{4} }

// The request message for [Operations.DeleteOperation][google.longrunning.Operations.DeleteOperation].
type DeleteOperationRequest struct {
	// The name of the operation resource to be deleted.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *DeleteOperationRequest) Reset()                    { *m = DeleteOperationRequest{} }
func (m *DeleteOperationRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteOperationRequest) ProtoMessage()               {}
func (*DeleteOperationRequest) Descriptor() ([]byte, []int) { return fileDescriptorOperations, []int{5} }

func init() {
	proto.RegisterType((*Operation)(nil), "google.longrunning.Operation")
	proto.RegisterType((*GetOperationRequest)(nil), "google.longrunning.GetOperationRequest")
	proto.RegisterType((*ListOperationsRequest)(nil), "google.longrunning.ListOperationsRequest")
	proto.RegisterType((*ListOperationsResponse)(nil), "google.longrunning.ListOperationsResponse")
	proto.RegisterType((*CancelOperationRequest)(nil), "google.longrunning.CancelOperationRequest")
	proto.RegisterType((*DeleteOperationRequest)(nil), "google.longrunning.DeleteOperationRequest")
}

var fileDescriptorOperations = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0xd3, 0x24, 0x4a, 0x86, 0xaa, 0x11, 0x03, 0x4d, 0x8c, 0xa1, 0xa2, 0x72, 0x50, 0x29,
	0x46, 0xb2, 0xd5, 0x72, 0xab, 0xc4, 0x81, 0x00, 0xa2, 0x07, 0x24, 0x2a, 0xb8, 0x22, 0x55, 0xdb,
	0x64, 0x1a, 0x59, 0x38, 0xbb, 0x66, 0xbd, 0x86, 0x16, 0xd4, 0x03, 0x1c, 0x38, 0x71, 0xe3, 0x25,
	0x78, 0x1f, 0x5e, 0x81, 0x07, 0x61, 0xb3, 0x71, 0x70, 0xea, 0x6e, 0x0a, 0xc7, 0xec, 0x7c, 0xf3,
	0xcd, 0xf7, 0x13, 0x43, 0x7f, 0x2c, 0xc4, 0x38, 0xa1, 0x28, 0x11, 0x7c, 0x2c, 0x73, 0xce, 0x63,
	0x3e, 0x8e, 0x44, 0x4a, 0x92, 0xa9, 0x58, 0xf0, 0x2c, 0x4c, 0xa5, 0x50, 0x02, 0x71, 0x06, 0x0a,
	0x17, 0x40, 0xde, 0x9d, 0x62, 0x91, 0xa5, 0x71, 0xc4, 0x38, 0x17, 0x6a, 0x71, 0xc3, 0xbb, 0x55,
	0x4c, 0xcd, 0xaf, 0xe3, 0xfc, 0x44, 0x43, 0xce, 0x8a, 0xd1, 0xed, 0xea, 0x88, 0x26, 0xa9, 0x9a,
	0x0f, 0x7b, 0xc5, 0x50, 0xa6, 0xc3, 0x28, 0xd3, 0x94, 0x79, 0x41, 0xe8, 0xff, 0x74, 0xa0, 0xfd,
	0x6a, 0xae, 0x0b, 0xd7, 0xa0, 0xce, 0xd9, 0x84, 0x5c, 0x67, 0xcb, 0xd9, 0x69, 0xe3, 0x36, 0xb4,
	0x26, 0xa4, 0xd8, 0x88, 0x29, 0xe6, 0xd6, 0xf4, 0xcb, 0xb5, 0xbd, 0x9b, 0x61, 0xa1, 0x78, 0x7e,
	0x24, 0x7c, 0xc2, 0xcf, 0xa6, 0x5b, 0x23, 0xc1, 0xc9, 0x5d, 0xd5, 0x98, 0x16, 0xf6, 0xa1, 0x41,
	0x52, 0x0a, 0xe9, 0xd6, 0xcd, 0x0a, 0xce, 0x57, 0xf4, 0xe9, 0xf0, 0x8d, 0x39, 0x7d, 0xb0, 0x82,
	0x3b, 0xd0, 0x92, 0x94, 0xa5, 0xda, 0x18, 0xb9, 0x8d, 0xe5, 0xd4, 0x07, 0x2b, 0x83, 0x16, 0x34,
	0x35, 0x32, 0x4f, 0x94, 0xdf, 0x87, 0x1b, 0x2f, 0x48, 0xfd, 0x15, 0xfb, 0x9a, 0xde, 0xe7, 0x94,
	0xa9, 0x8b, 0x9a, 0xfd, 0xb7, 0xb0, 0xf1, 0x32, 0xce, 0x4a, 0x54, 0x56, 0x85, 0xd5, 0x8d, 0xb5,
	0x75, 0x68, 0x9e, 0xc4, 0x89, 0x22, 0x59, 0x58, 0xbd, 0x0e, 0xed, 0x94, 0x8d, 0xe9, 0x28, 0x8b,
	0x3f, 0x91, 0xf1, 0xda, 0x40, 0x04, 0x30, 0x4f, 0x4a, 0xbc, 0x23, 0x6e, 0xbc, 0xb5, 0xfd, 0x11,
	0x74, 0xab, 0xec, 0x33, 0x13, 0xb8, 0x0b, 0x50, 0xd6, 0xab, 0x49, 0x57, 0xb5, 0xa5, 0xcd, 0xf0,
	0x72, 0xbf, 0x61, 0x19, 0x76, 0x0f, 0x3a, 0x9c, 0x4e, 0xd5, 0xd1, 0xc2, 0x95, 0x9a, 0xb9, 0xb2,
	0x0d, 0xdd, 0xa7, 0x8c, 0x0f, 0x29, 0xf9, 0x87, 0x57, 0x8d, 0x7b, 0x46, 0x09, 0x29, 0xba, 0x1a,
	0xb7, 0xf7, 0xbd, 0x0e, 0x50, 0x4a, 0xc6, 0x53, 0x58, 0x5b, 0xcc, 0x11, 0xef, 0xdb, 0x64, 0x5a,
	0x92, 0xf6, 0xae, 0xf6, 0xe3, 0x6f, 0x7d, 0xfd, 0xf5, 0xfb, 0x47, 0xcd, 0x43, 0x37, 0xfa, 0xb0,
	0x1b, 0x7d, 0x9e, 0xde, 0x7f, 0x5c, 0x46, 0x12, 0x05, 0xc1, 0x39, 0x7e, 0x73, 0x60, 0xfd, 0x62,
	0x7e, 0xf8, 0xc0, 0xc6, 0x69, 0x6d, 0xd0, 0x0b, 0xfe, 0x07, 0x3a, 0xab, 0xc3, 0xdf, 0x34, 0x5a,
	0x7a, 0xb8, 0x61, 0xd3, 0x72, 0x8e, 0x5f, 0x1c, 0xe8, 0x54, 0x22, 0x46, 0x2b, 0xbd, 0xbd, 0x07,
	0xaf, 0x7b, 0xe9, 0xcf, 0xfa, 0x7c, 0xfa, 0xb1, 0xf9, 0x81, 0x39, 0x7b, 0xcf, 0xbf, 0xbb, 0x2c,
	0x82, 0xfd, 0xa1, 0x21, 0xdc, 0x77, 0x02, 0xfc, 0x08, 0x9d, 0x4a, 0x7b, 0x76, 0x09, 0xf6, 0x8a,
	0x97, 0x4a, 0x28, 0x5a, 0x08, 0x96, 0xb6, 0x30, 0x78, 0x08, 0xdd, 0xa1, 0x98, 0x58, 0x4e, 0x0d,
	0x3a, 0x65, 0x92, 0x87, 0x53, 0xd6, 0x43, 0xe7, 0xb8, 0x69, 0xe8, 0x1f, 0xfd, 0x09, 0x00, 0x00,
	0xff, 0xff, 0x8a, 0xbf, 0x59, 0x1b, 0xd0, 0x04, 0x00, 0x00,
}
