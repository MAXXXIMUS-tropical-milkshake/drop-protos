// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: audiostreaming/audiostreaming.proto

package audiostreamingv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AudioStreamingService_Upload_FullMethodName = "/audiostreaming.AudioStreamingService/Upload"
)

// AudioStreamingServiceClient is the client API for AudioStreamingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AudioStreamingServiceClient interface {
	Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error)
}

type audioStreamingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAudioStreamingServiceClient(cc grpc.ClientConnInterface) AudioStreamingServiceClient {
	return &audioStreamingServiceClient{cc}
}

func (c *audioStreamingServiceClient) Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UploadResponse)
	err := c.cc.Invoke(ctx, AudioStreamingService_Upload_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AudioStreamingServiceServer is the server API for AudioStreamingService service.
// All implementations must embed UnimplementedAudioStreamingServiceServer
// for forward compatibility.
type AudioStreamingServiceServer interface {
	Upload(context.Context, *UploadRequest) (*UploadResponse, error)
	mustEmbedUnimplementedAudioStreamingServiceServer()
}

// UnimplementedAudioStreamingServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAudioStreamingServiceServer struct{}

func (UnimplementedAudioStreamingServiceServer) Upload(context.Context, *UploadRequest) (*UploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedAudioStreamingServiceServer) mustEmbedUnimplementedAudioStreamingServiceServer() {}
func (UnimplementedAudioStreamingServiceServer) testEmbeddedByValue()                               {}

// UnsafeAudioStreamingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AudioStreamingServiceServer will
// result in compilation errors.
type UnsafeAudioStreamingServiceServer interface {
	mustEmbedUnimplementedAudioStreamingServiceServer()
}

func RegisterAudioStreamingServiceServer(s grpc.ServiceRegistrar, srv AudioStreamingServiceServer) {
	// If the following call pancis, it indicates UnimplementedAudioStreamingServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AudioStreamingService_ServiceDesc, srv)
}

func _AudioStreamingService_Upload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudioStreamingServiceServer).Upload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AudioStreamingService_Upload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudioStreamingServiceServer).Upload(ctx, req.(*UploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AudioStreamingService_ServiceDesc is the grpc.ServiceDesc for AudioStreamingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AudioStreamingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "audiostreaming.AudioStreamingService",
	HandlerType: (*AudioStreamingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upload",
			Handler:    _AudioStreamingService_Upload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "audiostreaming/audiostreaming.proto",
}