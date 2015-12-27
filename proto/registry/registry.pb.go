// Code generated by protoc-gen-go.
// source: github.com/micro/discovery-srv/proto/registry/registry.proto
// DO NOT EDIT!

/*
Package registry is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/discovery-srv/proto/registry/registry.proto

It has these top-level messages:
	RegisterRequest
	RegisterResponse
	DeregisterRequest
	DeregisterResponse
	GetServiceRequest
	GetServiceResponse
	ListServicesRequest
	ListServicesResponse
	WatchRequest
	WatchResponse
*/
package registry

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import discovery "github.com/micro/go-platform/discovery/proto"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RegisterRequest struct {
	Service *discovery.Service `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
}

func (m *RegisterRequest) Reset()                    { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()               {}
func (*RegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RegisterRequest) GetService() *discovery.Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type RegisterResponse struct {
}

func (m *RegisterResponse) Reset()                    { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()               {}
func (*RegisterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type DeregisterRequest struct {
	Service *discovery.Service `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
}

func (m *DeregisterRequest) Reset()                    { *m = DeregisterRequest{} }
func (m *DeregisterRequest) String() string            { return proto.CompactTextString(m) }
func (*DeregisterRequest) ProtoMessage()               {}
func (*DeregisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DeregisterRequest) GetService() *discovery.Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type DeregisterResponse struct {
}

func (m *DeregisterResponse) Reset()                    { *m = DeregisterResponse{} }
func (m *DeregisterResponse) String() string            { return proto.CompactTextString(m) }
func (*DeregisterResponse) ProtoMessage()               {}
func (*DeregisterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type GetServiceRequest struct {
	Service string `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
}

func (m *GetServiceRequest) Reset()                    { *m = GetServiceRequest{} }
func (m *GetServiceRequest) String() string            { return proto.CompactTextString(m) }
func (*GetServiceRequest) ProtoMessage()               {}
func (*GetServiceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type GetServiceResponse struct {
	Services []*discovery.Service `protobuf:"bytes,1,rep,name=services" json:"services,omitempty"`
}

func (m *GetServiceResponse) Reset()                    { *m = GetServiceResponse{} }
func (m *GetServiceResponse) String() string            { return proto.CompactTextString(m) }
func (*GetServiceResponse) ProtoMessage()               {}
func (*GetServiceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetServiceResponse) GetServices() []*discovery.Service {
	if m != nil {
		return m.Services
	}
	return nil
}

type ListServicesRequest struct {
}

func (m *ListServicesRequest) Reset()                    { *m = ListServicesRequest{} }
func (m *ListServicesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListServicesRequest) ProtoMessage()               {}
func (*ListServicesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type ListServicesResponse struct {
	Services []*discovery.Service `protobuf:"bytes,1,rep,name=services" json:"services,omitempty"`
}

func (m *ListServicesResponse) Reset()                    { *m = ListServicesResponse{} }
func (m *ListServicesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListServicesResponse) ProtoMessage()               {}
func (*ListServicesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ListServicesResponse) GetServices() []*discovery.Service {
	if m != nil {
		return m.Services
	}
	return nil
}

type WatchRequest struct {
	Service string `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
}

func (m *WatchRequest) Reset()                    { *m = WatchRequest{} }
func (m *WatchRequest) String() string            { return proto.CompactTextString(m) }
func (*WatchRequest) ProtoMessage()               {}
func (*WatchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type WatchResponse struct {
	Action  string             `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	Service *discovery.Service `protobuf:"bytes,2,opt,name=service" json:"service,omitempty"`
}

func (m *WatchResponse) Reset()                    { *m = WatchResponse{} }
func (m *WatchResponse) String() string            { return proto.CompactTextString(m) }
func (*WatchResponse) ProtoMessage()               {}
func (*WatchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *WatchResponse) GetService() *discovery.Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "RegisterResponse")
	proto.RegisterType((*DeregisterRequest)(nil), "DeregisterRequest")
	proto.RegisterType((*DeregisterResponse)(nil), "DeregisterResponse")
	proto.RegisterType((*GetServiceRequest)(nil), "GetServiceRequest")
	proto.RegisterType((*GetServiceResponse)(nil), "GetServiceResponse")
	proto.RegisterType((*ListServicesRequest)(nil), "ListServicesRequest")
	proto.RegisterType((*ListServicesResponse)(nil), "ListServicesResponse")
	proto.RegisterType((*WatchRequest)(nil), "WatchRequest")
	proto.RegisterType((*WatchResponse)(nil), "WatchResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Registry service

type RegistryClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterResponse, error)
	Deregister(ctx context.Context, in *DeregisterRequest, opts ...client.CallOption) (*DeregisterResponse, error)
	GetService(ctx context.Context, in *GetServiceRequest, opts ...client.CallOption) (*GetServiceResponse, error)
	ListServices(ctx context.Context, in *ListServicesRequest, opts ...client.CallOption) (*ListServicesResponse, error)
	Watch(ctx context.Context, in *WatchRequest, opts ...client.CallOption) (Registry_WatchClient, error)
}

type registryClient struct {
	c           client.Client
	serviceName string
}

func NewRegistryClient(serviceName string, c client.Client) RegistryClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "registry"
	}
	return &registryClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *registryClient) Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Registry.Register", in)
	out := new(RegisterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) Deregister(ctx context.Context, in *DeregisterRequest, opts ...client.CallOption) (*DeregisterResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Registry.Deregister", in)
	out := new(DeregisterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) GetService(ctx context.Context, in *GetServiceRequest, opts ...client.CallOption) (*GetServiceResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Registry.GetService", in)
	out := new(GetServiceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) ListServices(ctx context.Context, in *ListServicesRequest, opts ...client.CallOption) (*ListServicesResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Registry.ListServices", in)
	out := new(ListServicesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) Watch(ctx context.Context, in *WatchRequest, opts ...client.CallOption) (Registry_WatchClient, error) {
	req := c.c.NewRequest(c.serviceName, "Registry.Watch", &WatchRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &registryWatchClient{stream}, nil
}

type Registry_WatchClient interface {
	RecvR() (*WatchResponse, error)
	client.Streamer
}

type registryWatchClient struct {
	client.Streamer
}

func (x *registryWatchClient) RecvR() (*WatchResponse, error) {
	m := new(WatchResponse)
	err := x.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Registry service

type RegistryHandler interface {
	Register(context.Context, *RegisterRequest, *RegisterResponse) error
	Deregister(context.Context, *DeregisterRequest, *DeregisterResponse) error
	GetService(context.Context, *GetServiceRequest, *GetServiceResponse) error
	ListServices(context.Context, *ListServicesRequest, *ListServicesResponse) error
	Watch(context.Context, *WatchRequest, Registry_WatchStream) error
}

func RegisterRegistryHandler(s server.Server, hdlr RegistryHandler) {
	s.Handle(s.NewHandler(&Registry{hdlr}))
}

type Registry struct {
	RegistryHandler
}

func (h *Registry) Register(ctx context.Context, in *RegisterRequest, out *RegisterResponse) error {
	return h.RegistryHandler.Register(ctx, in, out)
}

func (h *Registry) Deregister(ctx context.Context, in *DeregisterRequest, out *DeregisterResponse) error {
	return h.RegistryHandler.Deregister(ctx, in, out)
}

func (h *Registry) GetService(ctx context.Context, in *GetServiceRequest, out *GetServiceResponse) error {
	return h.RegistryHandler.GetService(ctx, in, out)
}

func (h *Registry) ListServices(ctx context.Context, in *ListServicesRequest, out *ListServicesResponse) error {
	return h.RegistryHandler.ListServices(ctx, in, out)
}

func (h *Registry) Watch(ctx context.Context, stream server.Streamer) error {
	m := new(WatchRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.RegistryHandler.Watch(ctx, m, &registryWatchStream{stream})
}

type Registry_WatchStream interface {
	SendR(*WatchResponse) error
	server.Streamer
}

type registryWatchStream struct {
	server.Streamer
}

func (x *registryWatchStream) SendR(m *WatchResponse) error {
	return x.Streamer.Send(m)
}

var fileDescriptor0 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x92, 0x4d, 0x4f, 0x02, 0x31,
	0x10, 0x86, 0x01, 0x23, 0xe2, 0xc8, 0xe7, 0x00, 0x89, 0xee, 0x45, 0xd3, 0x78, 0x30, 0x46, 0x66,
	0x11, 0x0f, 0x26, 0x46, 0x6f, 0x26, 0x5e, 0x3c, 0xe1, 0xc1, 0x33, 0xd4, 0x0a, 0x4d, 0x84, 0x62,
	0x5b, 0x48, 0xf8, 0x2b, 0xfe, 0x5a, 0x9b, 0xdd, 0xc2, 0xee, 0xba, 0x98, 0x70, 0xeb, 0x34, 0xef,
	0xfb, 0x74, 0xf3, 0xcc, 0xc2, 0xe3, 0x44, 0xda, 0xe9, 0x72, 0x4c, 0x5c, 0xcd, 0xc2, 0x99, 0xe4,
	0x5a, 0x85, 0x1f, 0xd2, 0x70, 0xb5, 0x12, 0x7a, 0xdd, 0x33, 0x7a, 0x15, 0x2e, 0xb4, 0xb2, 0x2a,
	0xd4, 0x62, 0x22, 0x8d, 0xd5, 0xeb, 0xed, 0x81, 0xa2, 0xfb, 0x20, 0xdf, 0x9e, 0xa8, 0xde, 0xe2,
	0x6b, 0x64, 0x3f, 0x95, 0x9e, 0x25, 0x24, 0x4f, 0xd9, 0xce, 0x71, 0x9b, 0xdd, 0x40, 0x63, 0x18,
	0xf1, 0x84, 0x1e, 0x8a, 0xef, 0xa5, 0x30, 0x16, 0xcf, 0xe0, 0xc8, 0x08, 0xbd, 0x92, 0x5c, 0x9c,
	0x16, 0x2f, 0x8a, 0x57, 0x27, 0x83, 0x0a, 0xbd, 0xc5, 0x33, 0x43, 0x68, 0x26, 0x69, 0xb3, 0x50,
	0x73, 0x23, 0x18, 0x41, 0xeb, 0x59, 0xe8, 0xfd, 0x19, 0x1d, 0xc0, 0x74, 0xde, 0x53, 0x2e, 0xa1,
	0xf5, 0x22, 0xac, 0xcf, 0x6c, 0x28, 0x8d, 0x2c, 0xe5, 0x98, 0xf5, 0x01, 0xd3, 0xa9, 0xb8, 0x8b,
	0x01, 0x54, 0x7c, 0xcc, 0xb8, 0xdc, 0x41, 0xe6, 0xb5, 0x2e, 0xb4, 0x5f, 0xdd, 0x4b, 0x7e, 0x34,
	0x9e, 0xcc, 0x06, 0xd0, 0xc9, 0x5e, 0xef, 0x81, 0x3a, 0x87, 0xea, 0xfb, 0xc8, 0xf2, 0xe9, 0xbf,
	0x5f, 0xf7, 0x00, 0x35, 0x1f, 0xf0, 0xb4, 0x3a, 0x94, 0x47, 0xdc, 0x4a, 0x35, 0x8f, 0x03, 0x69,
	0x2b, 0xa5, 0xac, 0x95, 0xc1, 0x4f, 0x09, 0x2a, 0x43, 0xbf, 0x58, 0xbc, 0xdd, 0x9c, 0x85, 0xc6,
	0x26, 0xfd, 0xd9, 0x4f, 0xd0, 0xa2, 0xdc, 0x0e, 0x0a, 0x78, 0x0f, 0x90, 0x58, 0x45, 0xa4, 0xdc,
	0x4a, 0x82, 0x36, 0xed, 0xd0, 0x1e, 0x15, 0x13, 0xa5, 0xae, 0x98, 0xdb, 0x82, 0x2b, 0xe6, 0x9d,
	0xbb, 0xe2, 0x13, 0x54, 0xd3, 0x0a, 0xb1, 0x43, 0x3b, 0x44, 0x07, 0x5d, 0xda, 0xe5, 0xd9, 0xd5,
	0xaf, 0xe1, 0x30, 0x92, 0x85, 0x35, 0x4a, 0x5b, 0x0d, 0xea, 0x94, 0x71, 0xc8, 0x0a, 0xfd, 0xe2,
	0xb8, 0x1c, 0xfd, 0xab, 0x77, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x2e, 0xa1, 0x57, 0x29,
	0x03, 0x00, 0x00,
}