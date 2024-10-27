// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: audio-streaming/audiostreaming.proto

/*
Package audiostreamingv1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package audiostreamingv1

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_AudioStreamingService_StreamAudio_0(ctx context.Context, marshaler runtime.Marshaler, client AudioStreamingServiceClient, req *http.Request, pathParams map[string]string) (AudioStreamingService_StreamAudioClient, runtime.ServerMetadata, error) {
	var protoReq StreamAudioRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	stream, err := client.StreamAudio(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil

}

// RegisterAudioStreamingServiceHandlerServer registers the http handlers for service AudioStreamingService to "mux".
// UnaryRPC     :call AudioStreamingServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterAudioStreamingServiceHandlerFromEndpoint instead.
// GRPC interceptors will not work for this type of registration. To use interceptors, you must use the "runtime.WithMiddlewares" option in the "runtime.NewServeMux" call.
func RegisterAudioStreamingServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server AudioStreamingServiceServer) error {

	mux.Handle("POST", pattern_AudioStreamingService_StreamAudio_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	return nil
}

// RegisterAudioStreamingServiceHandlerFromEndpoint is same as RegisterAudioStreamingServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterAudioStreamingServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterAudioStreamingServiceHandler(ctx, mux, conn)
}

// RegisterAudioStreamingServiceHandler registers the http handlers for service AudioStreamingService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterAudioStreamingServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterAudioStreamingServiceHandlerClient(ctx, mux, NewAudioStreamingServiceClient(conn))
}

// RegisterAudioStreamingServiceHandlerClient registers the http handlers for service AudioStreamingService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "AudioStreamingServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "AudioStreamingServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "AudioStreamingServiceClient" to call the correct interceptors. This client ignores the HTTP middlewares.
func RegisterAudioStreamingServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client AudioStreamingServiceClient) error {

	mux.Handle("POST", pattern_AudioStreamingService_StreamAudio_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/audiostreaming.AudioStreamingService/StreamAudio", runtime.WithHTTPPathPattern("/v1/audio/stream"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_AudioStreamingService_StreamAudio_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_AudioStreamingService_StreamAudio_0(annotatedContext, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_AudioStreamingService_StreamAudio_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "audio", "stream"}, ""))
)

var (
	forward_AudioStreamingService_StreamAudio_0 = runtime.ForwardResponseStream
)
