// Code generated by protoc-gen-go.
// source: github.com/micro/discovery-srv/proto/discovery/discovery.proto
// DO NOT EDIT!

/*
Package go_micro_srv_discovery_discovery is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/discovery-srv/proto/discovery/discovery.proto

It has these top-level messages:
	ServiceEndpoint
	EndpointsRequest
	EndpointsResponse
	HeartbeatsRequest
	HeartbeatsResponse
	WatchResultsRequest
	WatchResultsResponse
*/
package go_micro_srv_discovery_discovery

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type ServiceEndpoint struct {
	Service  string              `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	Version  string              `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Endpoint *discovery.Endpoint `protobuf:"bytes,3,opt,name=endpoint" json:"endpoint,omitempty"`
}

func (m *ServiceEndpoint) Reset()                    { *m = ServiceEndpoint{} }
func (m *ServiceEndpoint) String() string            { return proto.CompactTextString(m) }
func (*ServiceEndpoint) ProtoMessage()               {}
func (*ServiceEndpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ServiceEndpoint) GetEndpoint() *discovery.Endpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

type EndpointsRequest struct {
	Service string `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Limit   uint64 `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	Offset  uint64 `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
}

func (m *EndpointsRequest) Reset()                    { *m = EndpointsRequest{} }
func (m *EndpointsRequest) String() string            { return proto.CompactTextString(m) }
func (*EndpointsRequest) ProtoMessage()               {}
func (*EndpointsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type EndpointsResponse struct {
	Endpoints []*ServiceEndpoint `protobuf:"bytes,1,rep,name=endpoints" json:"endpoints,omitempty"`
}

func (m *EndpointsResponse) Reset()                    { *m = EndpointsResponse{} }
func (m *EndpointsResponse) String() string            { return proto.CompactTextString(m) }
func (*EndpointsResponse) ProtoMessage()               {}
func (*EndpointsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *EndpointsResponse) GetEndpoints() []*ServiceEndpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

type HeartbeatsRequest struct {
	Id     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	After  uint64 `protobuf:"varint,2,opt,name=after" json:"after,omitempty"`
	Limit  uint64 `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	Offset uint64 `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
}

func (m *HeartbeatsRequest) Reset()                    { *m = HeartbeatsRequest{} }
func (m *HeartbeatsRequest) String() string            { return proto.CompactTextString(m) }
func (*HeartbeatsRequest) ProtoMessage()               {}
func (*HeartbeatsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type HeartbeatsResponse struct {
	Heartbeats []*discovery.Heartbeat `protobuf:"bytes,1,rep,name=heartbeats" json:"heartbeats,omitempty"`
}

func (m *HeartbeatsResponse) Reset()                    { *m = HeartbeatsResponse{} }
func (m *HeartbeatsResponse) String() string            { return proto.CompactTextString(m) }
func (*HeartbeatsResponse) ProtoMessage()               {}
func (*HeartbeatsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *HeartbeatsResponse) GetHeartbeats() []*discovery.Heartbeat {
	if m != nil {
		return m.Heartbeats
	}
	return nil
}

type WatchResultsRequest struct {
	Service string `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	After   uint64 `protobuf:"varint,2,opt,name=after" json:"after,omitempty"`
	Limit   uint64 `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	Offset  uint64 `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
}

func (m *WatchResultsRequest) Reset()                    { *m = WatchResultsRequest{} }
func (m *WatchResultsRequest) String() string            { return proto.CompactTextString(m) }
func (*WatchResultsRequest) ProtoMessage()               {}
func (*WatchResultsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type WatchResultsResponse struct {
	Results []*discovery.Result `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *WatchResultsResponse) Reset()                    { *m = WatchResultsResponse{} }
func (m *WatchResultsResponse) String() string            { return proto.CompactTextString(m) }
func (*WatchResultsResponse) ProtoMessage()               {}
func (*WatchResultsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *WatchResultsResponse) GetResults() []*discovery.Result {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
	proto.RegisterType((*ServiceEndpoint)(nil), "go.micro.srv.discovery.discovery.ServiceEndpoint")
	proto.RegisterType((*EndpointsRequest)(nil), "go.micro.srv.discovery.discovery.EndpointsRequest")
	proto.RegisterType((*EndpointsResponse)(nil), "go.micro.srv.discovery.discovery.EndpointsResponse")
	proto.RegisterType((*HeartbeatsRequest)(nil), "go.micro.srv.discovery.discovery.HeartbeatsRequest")
	proto.RegisterType((*HeartbeatsResponse)(nil), "go.micro.srv.discovery.discovery.HeartbeatsResponse")
	proto.RegisterType((*WatchResultsRequest)(nil), "go.micro.srv.discovery.discovery.WatchResultsRequest")
	proto.RegisterType((*WatchResultsResponse)(nil), "go.micro.srv.discovery.discovery.WatchResultsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Discovery service

type DiscoveryClient interface {
	Endpoints(ctx context.Context, in *EndpointsRequest, opts ...client.CallOption) (*EndpointsResponse, error)
	Heartbeats(ctx context.Context, in *HeartbeatsRequest, opts ...client.CallOption) (*HeartbeatsResponse, error)
	WatchResults(ctx context.Context, in *WatchResultsRequest, opts ...client.CallOption) (*WatchResultsResponse, error)
}

type discoveryClient struct {
	c           client.Client
	serviceName string
}

func NewDiscoveryClient(serviceName string, c client.Client) DiscoveryClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.discovery.discovery"
	}
	return &discoveryClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *discoveryClient) Endpoints(ctx context.Context, in *EndpointsRequest, opts ...client.CallOption) (*EndpointsResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Discovery.Endpoints", in)
	out := new(EndpointsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryClient) Heartbeats(ctx context.Context, in *HeartbeatsRequest, opts ...client.CallOption) (*HeartbeatsResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Discovery.Heartbeats", in)
	out := new(HeartbeatsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryClient) WatchResults(ctx context.Context, in *WatchResultsRequest, opts ...client.CallOption) (*WatchResultsResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Discovery.WatchResults", in)
	out := new(WatchResultsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Discovery service

type DiscoveryHandler interface {
	Endpoints(context.Context, *EndpointsRequest, *EndpointsResponse) error
	Heartbeats(context.Context, *HeartbeatsRequest, *HeartbeatsResponse) error
	WatchResults(context.Context, *WatchResultsRequest, *WatchResultsResponse) error
}

func RegisterDiscoveryHandler(s server.Server, hdlr DiscoveryHandler) {
	s.Handle(s.NewHandler(&Discovery{hdlr}))
}

type Discovery struct {
	DiscoveryHandler
}

func (h *Discovery) Endpoints(ctx context.Context, in *EndpointsRequest, out *EndpointsResponse) error {
	return h.DiscoveryHandler.Endpoints(ctx, in, out)
}

func (h *Discovery) Heartbeats(ctx context.Context, in *HeartbeatsRequest, out *HeartbeatsResponse) error {
	return h.DiscoveryHandler.Heartbeats(ctx, in, out)
}

func (h *Discovery) WatchResults(ctx context.Context, in *WatchResultsRequest, out *WatchResultsResponse) error {
	return h.DiscoveryHandler.WatchResults(ctx, in, out)
}

var fileDescriptor0 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x93, 0xcd, 0x4f, 0xc2, 0x30,
	0x18, 0xc6, 0xe5, 0x43, 0x70, 0xaf, 0x1f, 0x48, 0xf5, 0xb0, 0xcc, 0xc4, 0x90, 0x9d, 0xbc, 0x50,
	0x14, 0xd0, 0x93, 0xf1, 0x84, 0x89, 0x17, 0x13, 0x23, 0x26, 0xc6, 0xe3, 0x18, 0x1d, 0x34, 0x61,
	0xeb, 0x6c, 0xcb, 0x12, 0x4e, 0xfe, 0xe9, 0x5a, 0xf6, 0x2d, 0x18, 0xc7, 0x6e, 0xe5, 0xe9, 0xfb,
	0x3c, 0xfd, 0xf1, 0xbc, 0x19, 0x3c, 0xcc, 0xa8, 0x9c, 0x2f, 0x27, 0xd8, 0x66, 0x6e, 0xcf, 0xa5,
	0x36, 0x67, 0xbd, 0x29, 0x15, 0x36, 0x0b, 0x08, 0x5f, 0x75, 0x05, 0x0f, 0x7a, 0x3e, 0x67, 0x32,
	0xa7, 0x65, 0x27, 0x1c, 0xde, 0xa0, 0xce, 0x8c, 0xe1, 0xd0, 0x87, 0xd5, 0x34, 0xce, 0x6e, 0xd3,
	0x93, 0x71, 0xbf, 0xf5, 0xc2, 0x8c, 0x75, 0xfd, 0x85, 0x25, 0x1d, 0xc6, 0xdd, 0x5c, 0xf2, 0xc6,
	0x4b, 0x51, 0xbe, 0xf9, 0x02, 0xad, 0x31, 0xe1, 0x01, 0xb5, 0xc9, 0xa3, 0x37, 0xf5, 0x19, 0xf5,
	0x24, 0x6a, 0x41, 0x53, 0x44, 0x92, 0x5e, 0xe9, 0x54, 0xae, 0xb4, 0xb5, 0xa0, 0x1c, 0x82, 0x32,
	0x4f, 0xaf, 0x86, 0xc2, 0x05, 0x1c, 0x90, 0x78, 0x5a, 0xaf, 0x29, 0xe5, 0xb0, 0xaf, 0xe1, 0xc4,
	0x6e, 0x8e, 0xe1, 0x34, 0x39, 0x8b, 0x57, 0xf2, 0xb9, 0x24, 0x62, 0x97, 0xc8, 0x63, 0xd8, 0x5f,
	0x50, 0x97, 0x46, 0x79, 0x75, 0x74, 0x02, 0x0d, 0xe6, 0x38, 0x82, 0x48, 0xbd, 0xbe, 0xfe, 0x6d,
	0x7e, 0x40, 0x3b, 0x17, 0x2a, 0x7c, 0xe6, 0x09, 0x82, 0x46, 0xa0, 0x25, 0x18, 0x42, 0xe5, 0xd6,
	0x14, 0xc7, 0x0d, 0x2e, 0xea, 0x0b, 0x6f, 0xfc, 0x5d, 0xf3, 0x19, 0xda, 0x4f, 0xc4, 0xe2, 0x72,
	0x42, 0xac, 0x0c, 0x18, 0xa0, 0x4a, 0xa7, 0x31, 0xab, 0x42, 0xb3, 0x1c, 0x49, 0x78, 0x48, 0x5a,
	0x2f, 0x22, 0x1d, 0x02, 0xca, 0xc7, 0xc5, 0xa8, 0x97, 0x00, 0xf3, 0x54, 0x8d, 0x59, 0x01, 0xa7,
	0x83, 0xe6, 0x1b, 0x9c, 0xbd, 0x5b, 0xd2, 0x9e, 0x2b, 0xc3, 0x72, 0xf1, 0x4f, 0x6f, 0xe5, 0x58,
	0xae, 0xe1, 0xfc, 0x77, 0x6a, 0x4c, 0xa3, 0x43, 0x93, 0x47, 0x52, 0x8c, 0xd2, 0xc4, 0xd1, 0x48,
	0xff, 0xbb, 0x0a, 0xda, 0x28, 0xa9, 0x0a, 0x05, 0xa0, 0xa5, 0xad, 0xa3, 0x7e, 0x71, 0xb5, 0x9b,
	0x7b, 0x37, 0x06, 0xa5, 0x3c, 0x11, 0x9d, 0xb9, 0x87, 0x56, 0x00, 0x59, 0x87, 0x68, 0x87, 0x90,
	0xad, 0x05, 0x1a, 0xc3, 0x72, 0xa6, 0xf4, 0xe9, 0x2f, 0x38, 0xca, 0x57, 0x86, 0x6e, 0x8b, 0x73,
	0xfe, 0x58, 0x9c, 0x71, 0x57, 0xd6, 0x96, 0x00, 0x4c, 0x1a, 0xe1, 0x77, 0x39, 0xf8, 0x09, 0x00,
	0x00, 0xff, 0xff, 0x58, 0x40, 0xf7, 0x67, 0x39, 0x04, 0x00, 0x00,
}
