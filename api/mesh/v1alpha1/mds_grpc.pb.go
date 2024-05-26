// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1alpha1

import (
	context "context"
)

import (
	v3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MDSSyncServiceClient is the client API for MDSSyncService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MDSSyncServiceClient interface {
	// MappingRegister from dp to cp, data plane register snp information to
	// control plane.
	MappingRegister(ctx context.Context, in *MappingRegisterRequest, opts ...grpc.CallOption) (*MappingRegisterResponse, error)
	MetadataRegister(ctx context.Context, in *MetaDataRegisterRequest, opts ...grpc.CallOption) (*MetaDataRegisterResponse, error)
	// sync mapping and metadate resource
	ZoneToDubboInstance(ctx context.Context, opts ...grpc.CallOption) (MDSSyncService_ZoneToDubboInstanceClient, error)
}

type mDSSyncServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMDSSyncServiceClient(cc grpc.ClientConnInterface) MDSSyncServiceClient {
	return &mDSSyncServiceClient{cc}
}

func (c *mDSSyncServiceClient) MappingRegister(ctx context.Context, in *MappingRegisterRequest, opts ...grpc.CallOption) (*MappingRegisterResponse, error) {
	out := new(MappingRegisterResponse)
	err := c.cc.Invoke(ctx, "/dubbo.mesh.v1alpha1.MDSSyncService/MappingRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mDSSyncServiceClient) MetadataRegister(ctx context.Context, in *MetaDataRegisterRequest, opts ...grpc.CallOption) (*MetaDataRegisterResponse, error) {
	out := new(MetaDataRegisterResponse)
	err := c.cc.Invoke(ctx, "/dubbo.mesh.v1alpha1.MDSSyncService/MetadataRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mDSSyncServiceClient) ZoneToDubboInstance(ctx context.Context, opts ...grpc.CallOption) (MDSSyncService_ZoneToDubboInstanceClient, error) {
	stream, err := c.cc.NewStream(ctx, &MDSSyncService_ServiceDesc.Streams[0], "/dubbo.mesh.v1alpha1.MDSSyncService/ZoneToDubboInstance", opts...)
	if err != nil {
		return nil, err
	}
	x := &mDSSyncServiceZoneToDubboInstanceClient{stream}
	return x, nil
}

type MDSSyncService_ZoneToDubboInstanceClient interface {
	Send(*v3.DeltaDiscoveryResponse) error
	Recv() (*v3.DeltaDiscoveryRequest, error)
	grpc.ClientStream
}

type mDSSyncServiceZoneToDubboInstanceClient struct {
	grpc.ClientStream
}

func (x *mDSSyncServiceZoneToDubboInstanceClient) Send(m *v3.DeltaDiscoveryResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *mDSSyncServiceZoneToDubboInstanceClient) Recv() (*v3.DeltaDiscoveryRequest, error) {
	m := new(v3.DeltaDiscoveryRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MDSSyncServiceServer is the server API for MDSSyncService service.
// All implementations must embed UnimplementedMDSSyncServiceServer
// for forward compatibility
type MDSSyncServiceServer interface {
	// MappingRegister from dp to cp, data plane register snp information to
	// control plane.
	MappingRegister(context.Context, *MappingRegisterRequest) (*MappingRegisterResponse, error)
	MetadataRegister(context.Context, *MetaDataRegisterRequest) (*MetaDataRegisterResponse, error)
	// sync mapping and metadate resource
	ZoneToDubboInstance(MDSSyncService_ZoneToDubboInstanceServer) error
	mustEmbedUnimplementedMDSSyncServiceServer()
}

// UnimplementedMDSSyncServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMDSSyncServiceServer struct {
}

func (UnimplementedMDSSyncServiceServer) MappingRegister(context.Context, *MappingRegisterRequest) (*MappingRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MappingRegister not implemented")
}
func (UnimplementedMDSSyncServiceServer) MetadataRegister(context.Context, *MetaDataRegisterRequest) (*MetaDataRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MetadataRegister not implemented")
}
func (UnimplementedMDSSyncServiceServer) ZoneToDubboInstance(MDSSyncService_ZoneToDubboInstanceServer) error {
	return status.Errorf(codes.Unimplemented, "method ZoneToDubboInstance not implemented")
}
func (UnimplementedMDSSyncServiceServer) mustEmbedUnimplementedMDSSyncServiceServer() {}

// UnsafeMDSSyncServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MDSSyncServiceServer will
// result in compilation errors.
type UnsafeMDSSyncServiceServer interface {
	mustEmbedUnimplementedMDSSyncServiceServer()
}

func RegisterMDSSyncServiceServer(s grpc.ServiceRegistrar, srv MDSSyncServiceServer) {
	s.RegisterService(&MDSSyncService_ServiceDesc, srv)
}

func _MDSSyncService_MappingRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MappingRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MDSSyncServiceServer).MappingRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dubbo.mesh.v1alpha1.MDSSyncService/MappingRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MDSSyncServiceServer).MappingRegister(ctx, req.(*MappingRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MDSSyncService_MetadataRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetaDataRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MDSSyncServiceServer).MetadataRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dubbo.mesh.v1alpha1.MDSSyncService/MetadataRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MDSSyncServiceServer).MetadataRegister(ctx, req.(*MetaDataRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MDSSyncService_ZoneToDubboInstance_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MDSSyncServiceServer).ZoneToDubboInstance(&mDSSyncServiceZoneToDubboInstanceServer{stream})
}

type MDSSyncService_ZoneToDubboInstanceServer interface {
	Send(*v3.DeltaDiscoveryRequest) error
	Recv() (*v3.DeltaDiscoveryResponse, error)
	grpc.ServerStream
}

type mDSSyncServiceZoneToDubboInstanceServer struct {
	grpc.ServerStream
}

func (x *mDSSyncServiceZoneToDubboInstanceServer) Send(m *v3.DeltaDiscoveryRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *mDSSyncServiceZoneToDubboInstanceServer) Recv() (*v3.DeltaDiscoveryResponse, error) {
	m := new(v3.DeltaDiscoveryResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MDSSyncService_ServiceDesc is the grpc.ServiceDesc for MDSSyncService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MDSSyncService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dubbo.mesh.v1alpha1.MDSSyncService",
	HandlerType: (*MDSSyncServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MappingRegister",
			Handler:    _MDSSyncService_MappingRegister_Handler,
		},
		{
			MethodName: "MetadataRegister",
			Handler:    _MDSSyncService_MetadataRegister_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ZoneToDubboInstance",
			Handler:       _MDSSyncService_ZoneToDubboInstance_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/mesh/v1alpha1/mds.proto",
}