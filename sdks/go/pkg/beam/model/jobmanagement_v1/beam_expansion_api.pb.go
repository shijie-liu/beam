// Code generated by protoc-gen-go. DO NOT EDIT.
// source: beam_expansion_api.proto

package jobmanagement_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import pipeline_v1 "github.com/apache/beam/sdks/go/pkg/beam/model/pipeline_v1"

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

type ExpansionRequest struct {
	// Set of components needed to interpret the transform, or which
	// may be useful for its expansion.  This includes the input
	// PCollections (if any) to the to-be-expanded transform, along
	// with their coders and windowing strategies.
	Components *pipeline_v1.Components `protobuf:"bytes,1,opt,name=components,proto3" json:"components,omitempty"`
	// The actual PTransform to be expaneded according to its spec.
	// Its input should be set, but its subtransforms and outputs
	// should not be.
	Transform *pipeline_v1.PTransform `protobuf:"bytes,2,opt,name=transform,proto3" json:"transform,omitempty"`
	// A namespace (prefix) to use for the id of any newly created
	// components.
	Namespace            string   `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExpansionRequest) Reset()         { *m = ExpansionRequest{} }
func (m *ExpansionRequest) String() string { return proto.CompactTextString(m) }
func (*ExpansionRequest) ProtoMessage()    {}
func (*ExpansionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_beam_expansion_api_864d3491b09f6fcb, []int{0}
}
func (m *ExpansionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExpansionRequest.Unmarshal(m, b)
}
func (m *ExpansionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExpansionRequest.Marshal(b, m, deterministic)
}
func (dst *ExpansionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExpansionRequest.Merge(dst, src)
}
func (m *ExpansionRequest) XXX_Size() int {
	return xxx_messageInfo_ExpansionRequest.Size(m)
}
func (m *ExpansionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExpansionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExpansionRequest proto.InternalMessageInfo

func (m *ExpansionRequest) GetComponents() *pipeline_v1.Components {
	if m != nil {
		return m.Components
	}
	return nil
}

func (m *ExpansionRequest) GetTransform() *pipeline_v1.PTransform {
	if m != nil {
		return m.Transform
	}
	return nil
}

func (m *ExpansionRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type ExpansionResponse struct {
	// Set of components needed to execute the expanded transform,
	// including the (original) inputs, outputs, and subtransforms.
	Components *pipeline_v1.Components `protobuf:"bytes,1,opt,name=components,proto3" json:"components,omitempty"`
	// The expanded transform itself, with references to its outputs
	// and subtransforms.
	Transform *pipeline_v1.PTransform `protobuf:"bytes,2,opt,name=transform,proto3" json:"transform,omitempty"`
	// (Optional) An string representation of any error encountered while
	// attempting to expand this transform.
	Error                string   `protobuf:"bytes,10,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExpansionResponse) Reset()         { *m = ExpansionResponse{} }
func (m *ExpansionResponse) String() string { return proto.CompactTextString(m) }
func (*ExpansionResponse) ProtoMessage()    {}
func (*ExpansionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_beam_expansion_api_864d3491b09f6fcb, []int{1}
}
func (m *ExpansionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExpansionResponse.Unmarshal(m, b)
}
func (m *ExpansionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExpansionResponse.Marshal(b, m, deterministic)
}
func (dst *ExpansionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExpansionResponse.Merge(dst, src)
}
func (m *ExpansionResponse) XXX_Size() int {
	return xxx_messageInfo_ExpansionResponse.Size(m)
}
func (m *ExpansionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExpansionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExpansionResponse proto.InternalMessageInfo

func (m *ExpansionResponse) GetComponents() *pipeline_v1.Components {
	if m != nil {
		return m.Components
	}
	return nil
}

func (m *ExpansionResponse) GetTransform() *pipeline_v1.PTransform {
	if m != nil {
		return m.Transform
	}
	return nil
}

func (m *ExpansionResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*ExpansionRequest)(nil), "org.apache.beam.model.expansion.v1.ExpansionRequest")
	proto.RegisterType((*ExpansionResponse)(nil), "org.apache.beam.model.expansion.v1.ExpansionResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExpansionServiceClient is the client API for ExpansionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExpansionServiceClient interface {
	Expand(ctx context.Context, in *ExpansionRequest, opts ...grpc.CallOption) (*ExpansionResponse, error)
}

type expansionServiceClient struct {
	cc *grpc.ClientConn
}

func NewExpansionServiceClient(cc *grpc.ClientConn) ExpansionServiceClient {
	return &expansionServiceClient{cc}
}

func (c *expansionServiceClient) Expand(ctx context.Context, in *ExpansionRequest, opts ...grpc.CallOption) (*ExpansionResponse, error) {
	out := new(ExpansionResponse)
	err := c.cc.Invoke(ctx, "/org.apache.beam.model.expansion.v1.ExpansionService/Expand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExpansionServiceServer is the server API for ExpansionService service.
type ExpansionServiceServer interface {
	Expand(context.Context, *ExpansionRequest) (*ExpansionResponse, error)
}

func RegisterExpansionServiceServer(s *grpc.Server, srv ExpansionServiceServer) {
	s.RegisterService(&_ExpansionService_serviceDesc, srv)
}

func _ExpansionService_Expand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpansionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpansionServiceServer).Expand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.apache.beam.model.expansion.v1.ExpansionService/Expand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpansionServiceServer).Expand(ctx, req.(*ExpansionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExpansionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "org.apache.beam.model.expansion.v1.ExpansionService",
	HandlerType: (*ExpansionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Expand",
			Handler:    _ExpansionService_Expand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "beam_expansion_api.proto",
}

func init() {
	proto.RegisterFile("beam_expansion_api.proto", fileDescriptor_beam_expansion_api_864d3491b09f6fcb)
}

var fileDescriptor_beam_expansion_api_864d3491b09f6fcb = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x89, 0x62, 0x21, 0xab, 0x87, 0xba, 0x28, 0x84, 0xe2, 0xa1, 0xe4, 0xd4, 0x8b, 0x03,
	0xa9, 0xfa, 0x00, 0x56, 0x3d, 0x89, 0x20, 0xd1, 0x93, 0x97, 0xb0, 0x49, 0xc7, 0xba, 0xd2, 0x9d,
	0x59, 0x77, 0x93, 0xe0, 0x2b, 0xf8, 0x60, 0xde, 0x7c, 0x28, 0x31, 0xa5, 0x49, 0x10, 0xc1, 0x7a,
	0xf3, 0xb8, 0xc3, 0xff, 0xcd, 0xf2, 0xcf, 0xff, 0x8b, 0x28, 0x47, 0x65, 0x32, 0x7c, 0xb5, 0x8a,
	0xbc, 0x66, 0xca, 0x94, 0xd5, 0x60, 0x1d, 0x97, 0x2c, 0x63, 0x76, 0x0b, 0x50, 0x56, 0x15, 0x4f,
	0x08, 0x5f, 0x22, 0x30, 0x3c, 0xc7, 0x25, 0xb4, 0x52, 0xa8, 0x93, 0xd1, 0x61, 0x43, 0xbb, 0x8a,
	0x08, 0x5d, 0x87, 0xc6, 0x1f, 0x81, 0x18, 0x5e, 0xad, 0x75, 0x29, 0xbe, 0x54, 0xe8, 0x4b, 0x79,
	0x23, 0x44, 0xc1, 0xc6, 0x32, 0x21, 0x95, 0x3e, 0x0a, 0xc6, 0xc1, 0x64, 0x77, 0x7a, 0x0c, 0x3f,
	0x7f, 0x62, 0xb5, 0xc5, 0xa5, 0x26, 0x84, 0x3a, 0x81, 0x8b, 0x16, 0x4a, 0x7b, 0x0b, 0xe4, 0xb5,
	0x08, 0x4b, 0xa7, 0xc8, 0x3f, 0xb2, 0x33, 0xd1, 0xd6, 0xc6, 0xdb, 0x6e, 0xef, 0xd7, 0x50, 0xda,
	0xf1, 0xf2, 0x48, 0x84, 0xa4, 0x0c, 0x7a, 0xab, 0x0a, 0x8c, 0xb6, 0xc7, 0xc1, 0x24, 0x4c, 0xbb,
	0x41, 0xfc, 0x1e, 0x88, 0xfd, 0x9e, 0x1d, 0x6f, 0x99, 0x3c, 0xfe, 0x6b, 0x3f, 0x07, 0x62, 0x07,
	0x9d, 0x63, 0x17, 0x89, 0xc6, 0xcb, 0xea, 0x31, 0x7d, 0xeb, 0xc7, 0x72, 0x87, 0xae, 0xd6, 0x05,
	0xca, 0x4a, 0x0c, 0x9a, 0xd9, 0x5c, 0x9e, 0xc2, 0xef, 0x89, 0xc3, 0xf7, 0x58, 0x47, 0x67, 0x7f,
	0xa4, 0x56, 0xd7, 0x9b, 0x5d, 0x8a, 0x0d, 0xfa, 0x35, 0xdb, 0x6b, 0xc1, 0x73, 0xab, 0x1f, 0x86,
	0xcf, 0x9c, 0x1b, 0x45, 0x6a, 0x81, 0x06, 0xa9, 0xcc, 0xea, 0x24, 0x1f, 0x34, 0x7d, 0x3b, 0xf9,
	0x0c, 0x00, 0x00, 0xff, 0xff, 0xb4, 0xc3, 0x3e, 0x66, 0xc6, 0x02, 0x00, 0x00,
}
