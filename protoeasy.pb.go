// Code generated by protoc-gen-go.
// source: protoeasy.proto
// DO NOT EDIT!

package protoeasy

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/pb/go/google/protobuf"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GoPluginType int32

const (
	GoPluginType_GO_PLUGIN_TYPE_NONE   GoPluginType = 0
	GoPluginType_GO_PLUGIN_TYPE_GO     GoPluginType = 1
	GoPluginType_GO_PLUGIN_TYPE_GOFAST GoPluginType = 2
)

var GoPluginType_name = map[int32]string{
	0: "GO_PLUGIN_TYPE_NONE",
	1: "GO_PLUGIN_TYPE_GO",
	2: "GO_PLUGIN_TYPE_GOFAST",
}
var GoPluginType_value = map[string]int32{
	"GO_PLUGIN_TYPE_NONE":   0,
	"GO_PLUGIN_TYPE_GO":     1,
	"GO_PLUGIN_TYPE_GOFAST": 2,
}

func (x GoPluginType) String() string {
	return proto.EnumName(GoPluginType_name, int32(x))
}
func (GoPluginType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GogoPluginType int32

const (
	GogoPluginType_GOGO_PLUGIN_TYPE_NONE       GogoPluginType = 0
	GogoPluginType_GOGO_PLUGIN_TYPE_GOGO       GogoPluginType = 1
	GogoPluginType_GOGO_PLUGIN_TYPE_GOGOFAST   GogoPluginType = 2
	GogoPluginType_GOGO_PLUGIN_TYPE_GOGOFASTER GogoPluginType = 3
	GogoPluginType_GOGO_PLUGIN_TYPE_GOGOSLICK  GogoPluginType = 4
)

var GogoPluginType_name = map[int32]string{
	0: "GOGO_PLUGIN_TYPE_NONE",
	1: "GOGO_PLUGIN_TYPE_GOGO",
	2: "GOGO_PLUGIN_TYPE_GOGOFAST",
	3: "GOGO_PLUGIN_TYPE_GOGOFASTER",
	4: "GOGO_PLUGIN_TYPE_GOGOSLICK",
}
var GogoPluginType_value = map[string]int32{
	"GOGO_PLUGIN_TYPE_NONE":       0,
	"GOGO_PLUGIN_TYPE_GOGO":       1,
	"GOGO_PLUGIN_TYPE_GOGOFAST":   2,
	"GOGO_PLUGIN_TYPE_GOGOFASTER": 3,
	"GOGO_PLUGIN_TYPE_GOGOSLICK":  4,
}

func (x GogoPluginType) String() string {
	return proto.EnumName(GogoPluginType_name, int32(x))
}
func (GogoPluginType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type CompileOptions struct {
	Grpc                        bool              `protobuf:"varint,1,opt,name=grpc" json:"grpc,omitempty"`
	GrpcGateway                 bool              `protobuf:"varint,2,opt,name=grpc_gateway,json=grpcGateway" json:"grpc_gateway,omitempty"`
	NoDefaultIncludes           bool              `protobuf:"varint,3,opt,name=no_default_includes,json=noDefaultIncludes" json:"no_default_includes,omitempty"`
	ExcludePattern              []string          `protobuf:"bytes,4,rep,name=exclude_pattern,json=excludePattern" json:"exclude_pattern,omitempty"`
	RelContext                  string            `protobuf:"bytes,5,opt,name=rel_context,json=relContext" json:"rel_context,omitempty"`
	ExtraPluginFlag             []string          `protobuf:"bytes,6,rep,name=extra_plugin_flag,json=extraPluginFlag" json:"extra_plugin_flag,omitempty"`
	Cpp                         bool              `protobuf:"varint,20,opt,name=cpp" json:"cpp,omitempty"`
	CppRelOut                   string            `protobuf:"bytes,21,opt,name=cpp_rel_out,json=cppRelOut" json:"cpp_rel_out,omitempty"`
	Csharp                      bool              `protobuf:"varint,30,opt,name=csharp" json:"csharp,omitempty"`
	CsharpRelOut                string            `protobuf:"bytes,31,opt,name=csharp_rel_out,json=csharpRelOut" json:"csharp_rel_out,omitempty"`
	Go                          bool              `protobuf:"varint,40,opt,name=go" json:"go,omitempty"`
	GoPluginType                GoPluginType      `protobuf:"varint,41,opt,name=go_plugin_type,json=goPluginType,enum=protoeasy.GoPluginType" json:"go_plugin_type,omitempty"`
	GoRelOut                    string            `protobuf:"bytes,42,opt,name=go_rel_out,json=goRelOut" json:"go_rel_out,omitempty"`
	GoImportPath                string            `protobuf:"bytes,43,opt,name=go_import_path,json=goImportPath" json:"go_import_path,omitempty"`
	GoNoDefaultModifiers        bool              `protobuf:"varint,44,opt,name=go_no_default_modifiers,json=goNoDefaultModifiers" json:"go_no_default_modifiers,omitempty"`
	GoModifiers                 map[string]string `protobuf:"bytes,45,rep,name=go_modifiers,json=goModifiers" json:"go_modifiers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Gogo                        bool              `protobuf:"varint,50,opt,name=gogo" json:"gogo,omitempty"`
	GogoPluginType              GogoPluginType    `protobuf:"varint,51,opt,name=gogo_plugin_type,json=gogoPluginType,enum=protoeasy.GogoPluginType" json:"gogo_plugin_type,omitempty"`
	GogoRelOut                  string            `protobuf:"bytes,52,opt,name=gogo_rel_out,json=gogoRelOut" json:"gogo_rel_out,omitempty"`
	GogoImportPath              string            `protobuf:"bytes,53,opt,name=gogo_import_path,json=gogoImportPath" json:"gogo_import_path,omitempty"`
	GogoNoDefaultModifiers      bool              `protobuf:"varint,54,opt,name=gogo_no_default_modifiers,json=gogoNoDefaultModifiers" json:"gogo_no_default_modifiers,omitempty"`
	GogoModifiers               map[string]string `protobuf:"bytes,55,rep,name=gogo_modifiers,json=gogoModifiers" json:"gogo_modifiers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Objc                        bool              `protobuf:"varint,60,opt,name=objc" json:"objc,omitempty"`
	ObjcRelOut                  string            `protobuf:"bytes,61,opt,name=objc_rel_out,json=objcRelOut" json:"objc_rel_out,omitempty"`
	Python                      bool              `protobuf:"varint,70,opt,name=python" json:"python,omitempty"`
	PythonRelOut                string            `protobuf:"bytes,71,opt,name=python_rel_out,json=pythonRelOut" json:"python_rel_out,omitempty"`
	Ruby                        bool              `protobuf:"varint,80,opt,name=ruby" json:"ruby,omitempty"`
	RubyRelOut                  string            `protobuf:"bytes,81,opt,name=ruby_rel_out,json=rubyRelOut" json:"ruby_rel_out,omitempty"`
	DescriptorSet               bool              `protobuf:"varint,90,opt,name=descriptor_set,json=descriptorSet" json:"descriptor_set,omitempty"`
	DescriptorSetRelOut         string            `protobuf:"bytes,91,opt,name=descriptor_set_rel_out,json=descriptorSetRelOut" json:"descriptor_set_rel_out,omitempty"`
	DescriptorSetFileName       string            `protobuf:"bytes,92,opt,name=descriptor_set_file_name,json=descriptorSetFileName" json:"descriptor_set_file_name,omitempty"`
	DescriptorSetIncludeImports bool              `protobuf:"varint,93,opt,name=descriptor_set_include_imports,json=descriptorSetIncludeImports" json:"descriptor_set_include_imports,omitempty"`
}

func (m *CompileOptions) Reset()                    { *m = CompileOptions{} }
func (m *CompileOptions) String() string            { return proto.CompactTextString(m) }
func (*CompileOptions) ProtoMessage()               {}
func (*CompileOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CompileOptions) GetGrpc() bool {
	if m != nil {
		return m.Grpc
	}
	return false
}

func (m *CompileOptions) GetGrpcGateway() bool {
	if m != nil {
		return m.GrpcGateway
	}
	return false
}

func (m *CompileOptions) GetNoDefaultIncludes() bool {
	if m != nil {
		return m.NoDefaultIncludes
	}
	return false
}

func (m *CompileOptions) GetExcludePattern() []string {
	if m != nil {
		return m.ExcludePattern
	}
	return nil
}

func (m *CompileOptions) GetRelContext() string {
	if m != nil {
		return m.RelContext
	}
	return ""
}

func (m *CompileOptions) GetExtraPluginFlag() []string {
	if m != nil {
		return m.ExtraPluginFlag
	}
	return nil
}

func (m *CompileOptions) GetCpp() bool {
	if m != nil {
		return m.Cpp
	}
	return false
}

func (m *CompileOptions) GetCppRelOut() string {
	if m != nil {
		return m.CppRelOut
	}
	return ""
}

func (m *CompileOptions) GetCsharp() bool {
	if m != nil {
		return m.Csharp
	}
	return false
}

func (m *CompileOptions) GetCsharpRelOut() string {
	if m != nil {
		return m.CsharpRelOut
	}
	return ""
}

func (m *CompileOptions) GetGo() bool {
	if m != nil {
		return m.Go
	}
	return false
}

func (m *CompileOptions) GetGoPluginType() GoPluginType {
	if m != nil {
		return m.GoPluginType
	}
	return GoPluginType_GO_PLUGIN_TYPE_NONE
}

func (m *CompileOptions) GetGoRelOut() string {
	if m != nil {
		return m.GoRelOut
	}
	return ""
}

func (m *CompileOptions) GetGoImportPath() string {
	if m != nil {
		return m.GoImportPath
	}
	return ""
}

func (m *CompileOptions) GetGoNoDefaultModifiers() bool {
	if m != nil {
		return m.GoNoDefaultModifiers
	}
	return false
}

func (m *CompileOptions) GetGoModifiers() map[string]string {
	if m != nil {
		return m.GoModifiers
	}
	return nil
}

func (m *CompileOptions) GetGogo() bool {
	if m != nil {
		return m.Gogo
	}
	return false
}

func (m *CompileOptions) GetGogoPluginType() GogoPluginType {
	if m != nil {
		return m.GogoPluginType
	}
	return GogoPluginType_GOGO_PLUGIN_TYPE_NONE
}

func (m *CompileOptions) GetGogoRelOut() string {
	if m != nil {
		return m.GogoRelOut
	}
	return ""
}

func (m *CompileOptions) GetGogoImportPath() string {
	if m != nil {
		return m.GogoImportPath
	}
	return ""
}

func (m *CompileOptions) GetGogoNoDefaultModifiers() bool {
	if m != nil {
		return m.GogoNoDefaultModifiers
	}
	return false
}

func (m *CompileOptions) GetGogoModifiers() map[string]string {
	if m != nil {
		return m.GogoModifiers
	}
	return nil
}

func (m *CompileOptions) GetObjc() bool {
	if m != nil {
		return m.Objc
	}
	return false
}

func (m *CompileOptions) GetObjcRelOut() string {
	if m != nil {
		return m.ObjcRelOut
	}
	return ""
}

func (m *CompileOptions) GetPython() bool {
	if m != nil {
		return m.Python
	}
	return false
}

func (m *CompileOptions) GetPythonRelOut() string {
	if m != nil {
		return m.PythonRelOut
	}
	return ""
}

func (m *CompileOptions) GetRuby() bool {
	if m != nil {
		return m.Ruby
	}
	return false
}

func (m *CompileOptions) GetRubyRelOut() string {
	if m != nil {
		return m.RubyRelOut
	}
	return ""
}

func (m *CompileOptions) GetDescriptorSet() bool {
	if m != nil {
		return m.DescriptorSet
	}
	return false
}

func (m *CompileOptions) GetDescriptorSetRelOut() string {
	if m != nil {
		return m.DescriptorSetRelOut
	}
	return ""
}

func (m *CompileOptions) GetDescriptorSetFileName() string {
	if m != nil {
		return m.DescriptorSetFileName
	}
	return ""
}

func (m *CompileOptions) GetDescriptorSetIncludeImports() bool {
	if m != nil {
		return m.DescriptorSetIncludeImports
	}
	return false
}

type Command struct {
	Arg []string `protobuf:"bytes,1,rep,name=arg" json:"arg,omitempty"`
}

func (m *Command) Reset()                    { *m = Command{} }
func (m *Command) String() string            { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()               {}
func (*Command) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Command) GetArg() []string {
	if m != nil {
		return m.Arg
	}
	return nil
}

type CompileInfo struct {
	Command         []*Command                `protobuf:"bytes,1,rep,name=command" json:"command,omitempty"`
	InputSizeBytes  uint64                    `protobuf:"varint,2,opt,name=input_size_bytes,json=inputSizeBytes" json:"input_size_bytes,omitempty"`
	OutputSizeBytes uint64                    `protobuf:"varint,3,opt,name=output_size_bytes,json=outputSizeBytes" json:"output_size_bytes,omitempty"`
	Duration        *google_protobuf.Duration `protobuf:"bytes,4,opt,name=duration" json:"duration,omitempty"`
}

func (m *CompileInfo) Reset()                    { *m = CompileInfo{} }
func (m *CompileInfo) String() string            { return proto.CompactTextString(m) }
func (*CompileInfo) ProtoMessage()               {}
func (*CompileInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CompileInfo) GetCommand() []*Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *CompileInfo) GetInputSizeBytes() uint64 {
	if m != nil {
		return m.InputSizeBytes
	}
	return 0
}

func (m *CompileInfo) GetOutputSizeBytes() uint64 {
	if m != nil {
		return m.OutputSizeBytes
	}
	return 0
}

func (m *CompileInfo) GetDuration() *google_protobuf.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

type CompileRequest struct {
	Tar            []byte          `protobuf:"bytes,1,opt,name=tar,proto3" json:"tar,omitempty"`
	CompileOptions *CompileOptions `protobuf:"bytes,2,opt,name=compile_options,json=compileOptions" json:"compile_options,omitempty"`
}

func (m *CompileRequest) Reset()                    { *m = CompileRequest{} }
func (m *CompileRequest) String() string            { return proto.CompactTextString(m) }
func (*CompileRequest) ProtoMessage()               {}
func (*CompileRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CompileRequest) GetTar() []byte {
	if m != nil {
		return m.Tar
	}
	return nil
}

func (m *CompileRequest) GetCompileOptions() *CompileOptions {
	if m != nil {
		return m.CompileOptions
	}
	return nil
}

type CompileResponse struct {
	Tar         []byte       `protobuf:"bytes,1,opt,name=tar,proto3" json:"tar,omitempty"`
	CompileInfo *CompileInfo `protobuf:"bytes,2,opt,name=compile_info,json=compileInfo" json:"compile_info,omitempty"`
}

func (m *CompileResponse) Reset()                    { *m = CompileResponse{} }
func (m *CompileResponse) String() string            { return proto.CompactTextString(m) }
func (*CompileResponse) ProtoMessage()               {}
func (*CompileResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CompileResponse) GetTar() []byte {
	if m != nil {
		return m.Tar
	}
	return nil
}

func (m *CompileResponse) GetCompileInfo() *CompileInfo {
	if m != nil {
		return m.CompileInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*CompileOptions)(nil), "protoeasy.CompileOptions")
	proto.RegisterType((*Command)(nil), "protoeasy.Command")
	proto.RegisterType((*CompileInfo)(nil), "protoeasy.CompileInfo")
	proto.RegisterType((*CompileRequest)(nil), "protoeasy.CompileRequest")
	proto.RegisterType((*CompileResponse)(nil), "protoeasy.CompileResponse")
	proto.RegisterEnum("protoeasy.GoPluginType", GoPluginType_name, GoPluginType_value)
	proto.RegisterEnum("protoeasy.GogoPluginType", GogoPluginType_name, GogoPluginType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for API service

type APIClient interface {
	Compile(ctx context.Context, in *CompileRequest, opts ...grpc.CallOption) (*CompileResponse, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) Compile(ctx context.Context, in *CompileRequest, opts ...grpc.CallOption) (*CompileResponse, error) {
	out := new(CompileResponse)
	err := grpc.Invoke(ctx, "/protoeasy.API/Compile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for API service

type APIServer interface {
	Compile(context.Context, *CompileRequest) (*CompileResponse, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_Compile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Compile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoeasy.API/Compile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Compile(ctx, req.(*CompileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protoeasy.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Compile",
			Handler:    _API_Compile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protoeasy.proto",
}

func init() { proto.RegisterFile("protoeasy.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1030 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0xed, 0x72, 0xdb, 0x44,
	0x17, 0xae, 0x62, 0x37, 0x8d, 0x8f, 0x5c, 0xd9, 0xd9, 0x7c, 0x6d, 0x9c, 0xf7, 0x4d, 0x8c, 0xa7,
	0x0c, 0x26, 0x04, 0x77, 0xc6, 0x21, 0x94, 0x32, 0x94, 0xa1, 0x71, 0x13, 0x8f, 0x87, 0x36, 0x31,
	0x72, 0xf8, 0x51, 0xbe, 0x34, 0xb2, 0xbc, 0xde, 0x08, 0x64, 0xed, 0x22, 0xad, 0x20, 0xee, 0x05,
	0xf1, 0x9f, 0xbb, 0xe0, 0xb2, 0x98, 0xdd, 0x95, 0x65, 0xcb, 0x31, 0xcc, 0xf0, 0x4b, 0x67, 0x9f,
	0xf3, 0x9c, 0x0f, 0x3d, 0x7b, 0x74, 0x04, 0x15, 0x1e, 0x31, 0xc1, 0x88, 0x1b, 0x4f, 0x5b, 0xca,
	0x42, 0xa5, 0x0c, 0xa8, 0x1d, 0x52, 0xc6, 0x68, 0x40, 0x9e, 0x2a, 0x64, 0x98, 0x8c, 0x9f, 0x8e,
	0x92, 0xc8, 0x15, 0x3e, 0x0b, 0x35, 0xb5, 0xf1, 0xa7, 0x09, 0x56, 0x87, 0x4d, 0xb8, 0x1f, 0x90,
	0x6b, 0x2e, 0xf1, 0x18, 0x21, 0x28, 0xd2, 0x88, 0x7b, 0xd8, 0xa8, 0x1b, 0xcd, 0x0d, 0x5b, 0xd9,
	0xe8, 0x3d, 0x28, 0xcb, 0xa7, 0x43, 0x5d, 0x41, 0x7e, 0x77, 0xa7, 0x78, 0x4d, 0xf9, 0x4c, 0x89,
	0x75, 0x35, 0x84, 0x5a, 0xb0, 0x15, 0x32, 0x67, 0x44, 0xc6, 0x6e, 0x12, 0x08, 0xc7, 0x0f, 0xbd,
	0x20, 0x19, 0x91, 0x18, 0x17, 0x14, 0x73, 0x33, 0x64, 0xaf, 0xb4, 0xa7, 0x97, 0x3a, 0xd0, 0x07,
	0x50, 0x21, 0x77, 0xca, 0x76, 0xb8, 0x2b, 0x04, 0x89, 0x42, 0x5c, 0xac, 0x17, 0x9a, 0x25, 0xdb,
	0x4a, 0xe1, 0xbe, 0x46, 0xd1, 0x11, 0x98, 0x11, 0x09, 0x1c, 0x8f, 0x85, 0x82, 0xdc, 0x09, 0xfc,
	0xb0, 0x6e, 0x34, 0x4b, 0x36, 0x44, 0x24, 0xe8, 0x68, 0x04, 0x1d, 0xc3, 0x26, 0xb9, 0x13, 0x91,
	0xeb, 0xf0, 0x20, 0xa1, 0x7e, 0xe8, 0x8c, 0x03, 0x97, 0xe2, 0x75, 0x95, 0xab, 0xa2, 0x1c, 0x7d,
	0x85, 0x5f, 0x06, 0x2e, 0x45, 0x55, 0x28, 0x78, 0x9c, 0xe3, 0x6d, 0xd5, 0x95, 0x34, 0xd1, 0x21,
	0x98, 0x1e, 0xe7, 0x8e, 0x2c, 0xc1, 0x12, 0x81, 0x77, 0x54, 0xfa, 0x92, 0xc7, 0xb9, 0x4d, 0x82,
	0xeb, 0x44, 0xa0, 0x5d, 0x58, 0xf7, 0xe2, 0x5b, 0x37, 0xe2, 0xf8, 0x50, 0x05, 0xa5, 0x27, 0xf4,
	0x04, 0x2c, 0x6d, 0x65, 0xa1, 0x47, 0x2a, 0xb4, 0xac, 0xd1, 0x34, 0xda, 0x82, 0x35, 0xca, 0x70,
	0x53, 0x45, 0xae, 0x51, 0x86, 0x5e, 0x80, 0x45, 0xd9, 0xac, 0x51, 0x31, 0xe5, 0x04, 0x7f, 0x58,
	0x37, 0x9a, 0x56, 0x7b, 0xaf, 0x35, 0xbf, 0xc4, 0x2e, 0xd3, 0x0d, 0xdf, 0x4c, 0x39, 0xb1, 0xcb,
	0x74, 0xe1, 0x84, 0xfe, 0x07, 0x40, 0x59, 0x56, 0xf0, 0x58, 0x15, 0xdc, 0xa0, 0x2c, 0x2d, 0xf6,
	0x44, 0x25, 0xf7, 0x27, 0x9c, 0x45, 0x42, 0x8a, 0x7a, 0x8b, 0x3f, 0xd2, 0x2d, 0x51, 0xd6, 0x53,
	0x60, 0xdf, 0x15, 0xb7, 0xe8, 0x0c, 0xf6, 0x28, 0x73, 0x16, 0xee, 0x6a, 0xc2, 0x46, 0xfe, 0xd8,
	0x27, 0x51, 0x8c, 0x4f, 0x54, 0x9f, 0xdb, 0x94, 0x5d, 0xcd, 0xae, 0xeb, 0xcd, 0xcc, 0x87, 0xde,
	0x40, 0x99, 0xb2, 0x05, 0xee, 0xc7, 0xf5, 0x42, 0xd3, 0x6c, 0x1f, 0x2f, 0xf4, 0x9d, 0x9f, 0xa3,
	0x56, 0x97, 0x65, 0xc1, 0x17, 0xa1, 0x88, 0xa6, 0xb6, 0x49, 0xe7, 0x88, 0x9a, 0x32, 0x46, 0x19,
	0x6e, 0xa7, 0x53, 0xc6, 0x28, 0x43, 0x1d, 0xa8, 0xca, 0x67, 0x4e, 0x9e, 0x53, 0x25, 0xcf, 0x7e,
	0x4e, 0x9e, 0x45, 0x49, 0x6c, 0x8b, 0xe6, 0xce, 0xa8, 0x2e, 0xfb, 0x5c, 0x10, 0xe9, 0x13, 0x3d,
	0x2f, 0x12, 0x4b, 0x65, 0x6a, 0xa6, 0x65, 0x16, 0x85, 0x3a, 0x53, 0x2c, 0x95, 0x6b, 0x41, 0xaa,
	0xe7, 0xb0, 0xaf, 0x98, 0x2b, 0xc5, 0xfa, 0x54, 0x75, 0xbe, 0x2b, 0x09, 0x2b, 0xe4, 0x1a, 0x80,
	0x4a, 0xb6, 0xc0, 0x7f, 0xa6, 0x04, 0x3b, 0xf9, 0x37, 0xc1, 0xe8, 0xb2, 0x64, 0x8f, 0x29, 0x5b,
	0x12, 0x8d, 0x0d, 0x7f, 0xf6, 0xf0, 0x17, 0x5a, 0x34, 0x69, 0xcb, 0xf7, 0x95, 0xcf, 0xec, 0x7d,
	0x5f, 0xe8, 0xf7, 0x95, 0xd8, 0x7c, 0x82, 0xf9, 0x54, 0xdc, 0xb2, 0x10, 0x5f, 0xea, 0x09, 0xd6,
	0x27, 0x39, 0x2e, 0xda, 0xca, 0x62, 0xbb, 0x7a, 0x5c, 0x34, 0x9a, 0x46, 0x23, 0x28, 0x46, 0xc9,
	0x70, 0x8a, 0xfb, 0xba, 0xa6, 0xb4, 0x65, 0x4d, 0xf9, 0xcc, 0xe2, 0xbe, 0x49, 0xbf, 0xc9, 0x64,
	0x38, 0x4d, 0xa3, 0xde, 0x07, 0x6b, 0x44, 0x62, 0x2f, 0xf2, 0xb9, 0x60, 0x91, 0x13, 0x13, 0x81,
	0xbf, 0x53, 0xf1, 0x8f, 0xe7, 0xe8, 0x80, 0x08, 0x74, 0x0a, 0xbb, 0x79, 0x5a, 0x96, 0xf2, 0x7b,
	0x95, 0x72, 0x2b, 0x47, 0x4f, 0x73, 0x3f, 0x03, 0xbc, 0x14, 0x34, 0xf6, 0x03, 0xe2, 0x84, 0xee,
	0x84, 0xe0, 0x1f, 0x54, 0xd8, 0x4e, 0x2e, 0xec, 0xd2, 0x0f, 0xc8, 0x95, 0x3b, 0x21, 0xa8, 0x03,
	0x87, 0x4b, 0x81, 0xe9, 0x9a, 0x4a, 0x47, 0x21, 0xc6, 0x3f, 0xaa, 0x26, 0x0f, 0x72, 0xe1, 0xe9,
	0xc6, 0xd2, 0x63, 0x11, 0xd7, 0xbe, 0x84, 0xea, 0xf2, 0x64, 0xcb, 0xad, 0xf2, 0x0b, 0x99, 0xaa,
	0x8d, 0x59, 0xb2, 0xa5, 0x89, 0xb6, 0xe1, 0xe1, 0x6f, 0x6e, 0x90, 0x10, 0xb5, 0x29, 0x4b, 0xb6,
	0x3e, 0x7c, 0xbe, 0xf6, 0x99, 0x51, 0xfb, 0x0a, 0xd0, 0xfd, 0x8b, 0xfe, 0x2f, 0x19, 0x1a, 0x07,
	0xf0, 0xa8, 0xc3, 0x26, 0x13, 0x37, 0x1c, 0xc9, 0x30, 0x37, 0xa2, 0xd8, 0x50, 0xcb, 0x4e, 0x9a,
	0x8d, 0xbf, 0x0c, 0x30, 0xd3, 0xb9, 0xea, 0x85, 0x63, 0x86, 0x4e, 0xe0, 0x91, 0xa7, 0xc9, 0x8a,
	0x65, 0xb6, 0x51, 0x7e, 0x00, 0xa5, 0xc7, 0x9e, 0x51, 0xe4, 0xa7, 0xe1, 0x87, 0x3c, 0x11, 0x4e,
	0xec, 0xbf, 0x23, 0xce, 0x70, 0x2a, 0x48, 0xac, 0xea, 0x17, 0x6d, 0x4b, 0xe1, 0x03, 0xff, 0x1d,
	0x39, 0x97, 0xa8, 0x5c, 0xba, 0x2c, 0x11, 0x4b, 0xd4, 0x82, 0xa2, 0x56, 0xb4, 0x63, 0xce, 0x3d,
	0x83, 0x8d, 0xd9, 0x6f, 0x07, 0x17, 0xeb, 0x46, 0xd3, 0x6c, 0xef, 0xb7, 0xf4, 0x7f, 0xa9, 0x35,
	0xfb, 0x2f, 0xb5, 0x5e, 0xa5, 0x04, 0x3b, 0xa3, 0x36, 0xc6, 0xd9, 0xaf, 0xc9, 0x26, 0xbf, 0x26,
	0x24, 0x16, 0xf2, 0x75, 0x85, 0x1b, 0x29, 0x95, 0xca, 0xb6, 0x34, 0xd1, 0x39, 0x54, 0x3c, 0xcd,
	0x71, 0x98, 0xfe, 0x8c, 0x54, 0xbf, 0x66, 0x6e, 0x63, 0xe4, 0xbf, 0x33, 0xdb, 0xf2, 0x72, 0xe7,
	0xc6, 0x4f, 0x50, 0xc9, 0xea, 0xc4, 0x9c, 0x85, 0x31, 0x59, 0x51, 0xe8, 0x39, 0x94, 0x67, 0x85,
	0xfc, 0x70, 0xcc, 0xd2, 0x2a, 0xbb, 0xf7, 0xab, 0x48, 0xd5, 0x6d, 0xd3, 0x9b, 0x1f, 0x8e, 0xdf,
	0x42, 0x79, 0x71, 0xa5, 0xa3, 0x3d, 0xd8, 0xea, 0x5e, 0x3b, 0xfd, 0xd7, 0xdf, 0x76, 0x7b, 0x57,
	0xce, 0xcd, 0xdb, 0xfe, 0x85, 0x73, 0x75, 0x7d, 0x75, 0x51, 0x7d, 0x80, 0x76, 0x60, 0x73, 0xc9,
	0xd1, 0xbd, 0xae, 0x1a, 0x68, 0x1f, 0x76, 0xee, 0xc1, 0x97, 0x2f, 0x07, 0x37, 0xd5, 0xb5, 0xe3,
	0x3f, 0x0c, 0xb0, 0xf2, 0xfb, 0x50, 0xb3, 0x57, 0xe7, 0x5f, 0xe5, 0x92, 0x40, 0xd5, 0x40, 0xff,
	0x87, 0xfd, 0x95, 0x2e, 0x5d, 0x07, 0x1d, 0xc1, 0xc1, 0x3f, 0xba, 0x2f, 0xec, 0x6a, 0x01, 0x1d,
	0x42, 0x6d, 0x25, 0x61, 0xf0, 0xba, 0xd7, 0xf9, 0xba, 0x5a, 0x6c, 0xf7, 0xa0, 0xf0, 0xb2, 0xdf,
	0x43, 0xe7, 0x6a, 0x74, 0xa5, 0x32, 0x68, 0xc5, 0x05, 0xa5, 0xd7, 0x5c, 0xab, 0xad, 0x72, 0xe9,
	0x9b, 0x69, 0x3c, 0x18, 0xae, 0x2b, 0xe7, 0xe9, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x18, 0x90,
	0xf0, 0xc7, 0xf7, 0x08, 0x00, 0x00,
}
