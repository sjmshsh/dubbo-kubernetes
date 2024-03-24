// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: audit/v1alpha1/service.proto

package auditv1alpha1

import (
	context "context"
)

import (
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AuditService_ListAuditedEvents_FullMethodName = "/bufman.dubbo.apache.org.audit.v1alpha1.AuditService/ListAuditedEvents"
)

// AuditServiceClient is the client API for AuditService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuditServiceClient interface {
	// ListAuditedEvents lists audited events recorded in the BSR instance.
	ListAuditedEvents(ctx context.Context, in *ListAuditedEventsRequest, opts ...grpc.CallOption) (*ListAuditedEventsResponse, error)
}

type auditServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuditServiceClient(cc grpc.ClientConnInterface) AuditServiceClient {
	return &auditServiceClient{cc}
}

func (c *auditServiceClient) ListAuditedEvents(ctx context.Context, in *ListAuditedEventsRequest, opts ...grpc.CallOption) (*ListAuditedEventsResponse, error) {
	out := new(ListAuditedEventsResponse)
	err := c.cc.Invoke(ctx, AuditService_ListAuditedEvents_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuditServiceServer is the server API for AuditService service.
// All implementations must embed UnimplementedAuditServiceServer
// for forward compatibility
type AuditServiceServer interface {
	// ListAuditedEvents lists audited events recorded in the BSR instance.
	ListAuditedEvents(context.Context, *ListAuditedEventsRequest) (*ListAuditedEventsResponse, error)
	mustEmbedUnimplementedAuditServiceServer()
}

// UnimplementedAuditServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuditServiceServer struct {
}

func (UnimplementedAuditServiceServer) ListAuditedEvents(context.Context, *ListAuditedEventsRequest) (*ListAuditedEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAuditedEvents not implemented")
}
func (UnimplementedAuditServiceServer) mustEmbedUnimplementedAuditServiceServer() {}

// UnsafeAuditServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuditServiceServer will
// result in compilation errors.
type UnsafeAuditServiceServer interface {
	mustEmbedUnimplementedAuditServiceServer()
}

func RegisterAuditServiceServer(s grpc.ServiceRegistrar, srv AuditServiceServer) {
	s.RegisterService(&AuditService_ServiceDesc, srv)
}

func _AuditService_ListAuditedEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAuditedEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditServiceServer).ListAuditedEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuditService_ListAuditedEvents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditServiceServer).ListAuditedEvents(ctx, req.(*ListAuditedEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuditService_ServiceDesc is the grpc.ServiceDesc for AuditService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuditService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bufman.dubbo.apache.org.audit.v1alpha1.AuditService",
	HandlerType: (*AuditServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAuditedEvents",
			Handler:    _AuditService_ListAuditedEvents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "audit/v1alpha1/service.proto",
}