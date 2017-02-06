// Code generated by protoc-gen-go.
// source: google/api/usage.proto
// DO NOT EDIT!

package google_api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Configuration controlling usage of a service.
type Usage struct {
	// Requirements that must be satisfied before a consumer project can use the
	// service. Each requirement is of the form <service.name>/<requirement-id>;
	// for example 'serviceusage.googleapis.com/billing-enabled'.
	Requirements []string `protobuf:"bytes,1,rep,name=requirements" json:"requirements,omitempty"`
	// A list of usage rules that apply to individual API methods.
	//
	// **NOTE:** All service configuration rules follow "last one wins" order.
	Rules []*UsageRule `protobuf:"bytes,6,rep,name=rules" json:"rules,omitempty"`
}

func (m *Usage) Reset()                    { *m = Usage{} }
func (m *Usage) String() string            { return proto.CompactTextString(m) }
func (*Usage) ProtoMessage()               {}
func (*Usage) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{0} }

func (m *Usage) GetRequirements() []string {
	if m != nil {
		return m.Requirements
	}
	return nil
}

func (m *Usage) GetRules() []*UsageRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// Usage configuration rules for the service.
//
// NOTE: Under development.
//
//
// Use this rule to configure unregistered calls for the service. Unregistered
// calls are calls that do not contain consumer project identity.
// (Example: calls that do not contain an API key).
// By default, API methods do not allow unregistered calls, and each method call
// must be identified by a consumer project identity. Use this rule to
// allow/disallow unregistered calls.
//
// Example of an API that wants to allow unregistered calls for entire service.
//
//     usage:
//       rules:
//       - selector: "*"
//         allow_unregistered_calls: true
//
// Example of a method that wants to allow unregistered calls.
//
//     usage:
//       rules:
//       - selector: "google.example.library.v1.LibraryService.CreateBook"
//         allow_unregistered_calls: true
type UsageRule struct {
	// Selects the methods to which this rule applies. Use '*' to indicate all
	// methods in all APIs.
	//
	// Refer to [selector][google.api.DocumentationRule.selector] for syntax details.
	Selector string `protobuf:"bytes,1,opt,name=selector" json:"selector,omitempty"`
	// True, if the method allows unregistered calls; false otherwise.
	AllowUnregisteredCalls bool `protobuf:"varint,2,opt,name=allow_unregistered_calls,json=allowUnregisteredCalls" json:"allow_unregistered_calls,omitempty"`
}

func (m *UsageRule) Reset()                    { *m = UsageRule{} }
func (m *UsageRule) String() string            { return proto.CompactTextString(m) }
func (*UsageRule) ProtoMessage()               {}
func (*UsageRule) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{1} }

func (m *UsageRule) GetSelector() string {
	if m != nil {
		return m.Selector
	}
	return ""
}

func (m *UsageRule) GetAllowUnregisteredCalls() bool {
	if m != nil {
		return m.AllowUnregisteredCalls
	}
	return false
}

func init() {
	proto.RegisterType((*Usage)(nil), "google.api.Usage")
	proto.RegisterType((*UsageRule)(nil), "google.api.UsageRule")
}

func init() { proto.RegisterFile("google/api/usage.proto", fileDescriptor21) }

var fileDescriptor21 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x8f, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0xd9, 0xd6, 0x96, 0x66, 0x14, 0x0f, 0x0b, 0x96, 0xa5, 0x78, 0x08, 0x01, 0x21, 0x20,
	0xa4, 0xa0, 0x17, 0xaf, 0xd6, 0x83, 0x78, 0x0b, 0x0b, 0x05, 0x6f, 0x65, 0x8d, 0x43, 0x58, 0x98,
	0xee, 0xc4, 0xfd, 0x83, 0xdf, 0xc7, 0x4f, 0x2a, 0x49, 0x20, 0xc6, 0xe3, 0xbc, 0xf7, 0x9b, 0x79,
	0xf3, 0x60, 0xdb, 0x32, 0xb7, 0x84, 0x7b, 0xd3, 0xd9, 0x7d, 0x0a, 0xa6, 0xc5, 0xaa, 0xf3, 0x1c,
	0x59, 0xc2, 0xa8, 0x57, 0xa6, 0xb3, 0xbb, 0xdb, 0x19, 0x63, 0x9c, 0xe3, 0x68, 0xa2, 0x65, 0x17,
	0x46, 0xb2, 0x78, 0x87, 0xd5, 0xb1, 0x5f, 0x94, 0x05, 0x5c, 0x79, 0xfc, 0x4a, 0xd6, 0xe3, 0x19,
	0x5d, 0x0c, 0x4a, 0xe4, 0xcb, 0x32, 0xd3, 0xff, 0x34, 0x79, 0x0f, 0x2b, 0x9f, 0x08, 0x83, 0x5a,
	0xe7, 0xcb, 0xf2, 0xf2, 0xe1, 0xa6, 0xfa, 0x8b, 0xa9, 0x86, 0x2b, 0x3a, 0x11, 0xea, 0x91, 0x29,
	0x0c, 0x64, 0x93, 0x26, 0x77, 0xb0, 0x09, 0x48, 0xd8, 0x44, 0xf6, 0x4a, 0xe4, 0xa2, 0xcc, 0xf4,
	0x34, 0xcb, 0x27, 0x50, 0x86, 0x88, 0xbf, 0x4f, 0xc9, 0x79, 0x6c, 0x6d, 0x88, 0xe8, 0xf1, 0xf3,
	0xd4, 0x18, 0xa2, 0xa0, 0x16, 0xb9, 0x28, 0x37, 0x7a, 0x3b, 0xf8, 0xc7, 0x99, 0xfd, 0xd2, 0xbb,
	0x87, 0x3b, 0xb8, 0x6e, 0xf8, 0x3c, 0xfb, 0xe2, 0x00, 0x43, 0x64, 0xdd, 0x57, 0xab, 0xc5, 0xcf,
	0xe2, 0xe2, 0xf5, 0xb9, 0x7e, 0xfb, 0x58, 0x0f, 0x55, 0x1f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xcd, 0xad, 0x83, 0x99, 0x2e, 0x01, 0x00, 0x00,
}