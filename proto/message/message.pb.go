// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/micro/message-api/proto/message/message.proto

/*
Package go_micro_srv_message is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/message-api/proto/message/message.proto

It has these top-level messages:
*/
package go_micro_srv_message

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import go_micro_api "github.com/micro/micro/api/proto"

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Message service

type MessageClient interface {
	Create(ctx context.Context, in *go_micro_api.Request, opts ...grpc.CallOption) (*go_micro_api.Response, error)
	Delete(ctx context.Context, in *go_micro_api.Request, opts ...grpc.CallOption) (*go_micro_api.Response, error)
	Update(ctx context.Context, in *go_micro_api.Request, opts ...grpc.CallOption) (*go_micro_api.Response, error)
	Search(ctx context.Context, in *go_micro_api.Request, opts ...grpc.CallOption) (*go_micro_api.Response, error)
	Read(ctx context.Context, in *go_micro_api.Request, opts ...grpc.CallOption) (*go_micro_api.Response, error)
}

type messageClient struct {
	cc *grpc.ClientConn
}

func NewMessageClient(cc *grpc.ClientConn) MessageClient {
	return &messageClient{cc}
}

func (c *messageClient) Create(ctx context.Context, in *go_micro_api.Request, opts ...grpc.CallOption) (*go_micro_api.Response, error) {
	out := new(go_micro_api.Response)
	err := grpc.Invoke(ctx, "/go.micro.srv.message.Message/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) Delete(ctx context.Context, in *go_micro_api.Request, opts ...grpc.CallOption) (*go_micro_api.Response, error) {
	out := new(go_micro_api.Response)
	err := grpc.Invoke(ctx, "/go.micro.srv.message.Message/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) Update(ctx context.Context, in *go_micro_api.Request, opts ...grpc.CallOption) (*go_micro_api.Response, error) {
	out := new(go_micro_api.Response)
	err := grpc.Invoke(ctx, "/go.micro.srv.message.Message/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) Search(ctx context.Context, in *go_micro_api.Request, opts ...grpc.CallOption) (*go_micro_api.Response, error) {
	out := new(go_micro_api.Response)
	err := grpc.Invoke(ctx, "/go.micro.srv.message.Message/Search", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) Read(ctx context.Context, in *go_micro_api.Request, opts ...grpc.CallOption) (*go_micro_api.Response, error) {
	out := new(go_micro_api.Response)
	err := grpc.Invoke(ctx, "/go.micro.srv.message.Message/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Message service

type MessageServer interface {
	Create(context.Context, *go_micro_api.Request) (*go_micro_api.Response, error)
	Delete(context.Context, *go_micro_api.Request) (*go_micro_api.Response, error)
	Update(context.Context, *go_micro_api.Request) (*go_micro_api.Response, error)
	Search(context.Context, *go_micro_api.Request) (*go_micro_api.Response, error)
	Read(context.Context, *go_micro_api.Request) (*go_micro_api.Response, error)
}

func RegisterMessageServer(s *grpc.Server, srv MessageServer) {
	s.RegisterService(&_Message_serviceDesc, srv)
}

func _Message_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(go_micro_api.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.message.Message/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).Create(ctx, req.(*go_micro_api.Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(go_micro_api.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.message.Message/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).Delete(ctx, req.(*go_micro_api.Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(go_micro_api.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.message.Message/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).Update(ctx, req.(*go_micro_api.Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(go_micro_api.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.message.Message/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).Search(ctx, req.(*go_micro_api.Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(go_micro_api.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.message.Message/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).Read(ctx, req.(*go_micro_api.Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Message_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.srv.message.Message",
	HandlerType: (*MessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Message_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Message_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Message_Update_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Message_Search_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _Message_Read_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/micro/message-api/proto/message/message.proto",
}

func init() {
	proto.RegisterFile("github.com/micro/message-api/proto/message/message.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x48, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xcd, 0x4c, 0x2e, 0xca, 0xd7, 0xcf, 0x4d, 0x2d,
	0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x4d, 0x2c, 0xc8, 0xd4, 0x2f, 0x28, 0xca, 0x2f, 0x81, 0x8b, 0xc0,
	0x68, 0x3d, 0xb0, 0xa8, 0x90, 0x48, 0x7a, 0xbe, 0x1e, 0x58, 0x87, 0x5e, 0x71, 0x51, 0x99, 0x1e,
	0x54, 0x4e, 0x4a, 0x0b, 0xd3, 0x3c, 0x30, 0x89, 0x30, 0x2d, 0xb1, 0x20, 0x13, 0x62, 0x82, 0xd1,
	0x3a, 0x26, 0x2e, 0x76, 0x5f, 0x88, 0x3e, 0x21, 0x4b, 0x2e, 0x36, 0xe7, 0xa2, 0xd4, 0xc4, 0x92,
	0x54, 0x21, 0x51, 0x3d, 0xb8, 0xc1, 0x20, 0xa5, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x52,
	0x62, 0xe8, 0xc2, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x4a, 0x0c, 0x20, 0xad, 0x2e, 0xa9, 0x39,
	0xa9, 0x64, 0x6a, 0x0d, 0x2d, 0x48, 0x21, 0xd7, 0xd6, 0xe0, 0xd4, 0xc4, 0xa2, 0xe4, 0x0c, 0xd2,
	0xb5, 0x9a, 0x73, 0xb1, 0x04, 0xa5, 0x26, 0xa6, 0x90, 0xac, 0x31, 0x89, 0x0d, 0x1c, 0x6e, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8a, 0xf1, 0x80, 0xa7, 0xb5, 0x01, 0x00, 0x00,
}
