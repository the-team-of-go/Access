// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ToA

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ReportClient is the client API for Report service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReportClient interface {
	Send(ctx context.Context, in *ReportReq, opts ...grpc.CallOption) (*ReportRsp, error)
}

type reportClient struct {
	cc grpc.ClientConnInterface
}

func NewReportClient(cc grpc.ClientConnInterface) ReportClient {
	return &reportClient{cc}
}

func (c *reportClient) Send(ctx context.Context, in *ReportReq, opts ...grpc.CallOption) (*ReportRsp, error) {
	out := new(ReportRsp)
	err := c.cc.Invoke(ctx, "/ToA.Report/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportServer is the server API for Report service.
// All implementations must embed UnimplementedReportServer
// for forward compatibility
type ReportServer interface {
	Send(context.Context, *ReportReq) (*ReportRsp, error)
	mustEmbedUnimplementedReportServer()
}

// UnimplementedReportServer must be embedded to have forward compatible implementations.
type UnimplementedReportServer struct {
}

func (UnimplementedReportServer) Send(context.Context, *ReportReq) (*ReportRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedReportServer) mustEmbedUnimplementedReportServer() {}

// UnsafeReportServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReportServer will
// result in compilation errors.
type UnsafeReportServer interface {
	mustEmbedUnimplementedReportServer()
}

func RegisterReportServer(s grpc.ServiceRegistrar, srv ReportServer) {
	s.RegisterService(&Report_ServiceDesc, srv)
}

func _Report_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ToA.Report/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServer).Send(ctx, req.(*ReportReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Report_ServiceDesc is the grpc.ServiceDesc for Report service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Report_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ToA.Report",
	HandlerType: (*ReportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Report_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ToA.proto",
}
