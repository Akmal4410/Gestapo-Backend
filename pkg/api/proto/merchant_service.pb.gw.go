// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: api/proto/merchant_service.proto

/*
Package proto is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package proto

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

func request_MerchantService_GetProfile_0(ctx context.Context, marshaler runtime.Marshaler, client MerchantServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetMerchantProfileRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["user_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "user_id")
	}

	protoReq.UserId, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "user_id", err)
	}

	msg, err := client.GetProfile(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_MerchantService_GetProfile_0(ctx context.Context, marshaler runtime.Marshaler, server MerchantServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetMerchantProfileRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["user_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "user_id")
	}

	protoReq.UserId, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "user_id", err)
	}

	msg, err := server.GetProfile(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterMerchantServiceHandlerServer registers the http handlers for service MerchantService to "mux".
// UnaryRPC     :call MerchantServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterMerchantServiceHandlerFromEndpoint instead.
func RegisterMerchantServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server MerchantServiceServer) error {

	mux.Handle("GET", pattern_MerchantService_GetProfile_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/pb.MerchantService/GetProfile", runtime.WithHTTPPathPattern("/merchant/profile/{user_id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_MerchantService_GetProfile_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_MerchantService_GetProfile_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterMerchantServiceHandlerFromEndpoint is same as RegisterMerchantServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterMerchantServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterMerchantServiceHandler(ctx, mux, conn)
}

// RegisterMerchantServiceHandler registers the http handlers for service MerchantService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterMerchantServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterMerchantServiceHandlerClient(ctx, mux, NewMerchantServiceClient(conn))
}

// RegisterMerchantServiceHandlerClient registers the http handlers for service MerchantService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "MerchantServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "MerchantServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "MerchantServiceClient" to call the correct interceptors.
func RegisterMerchantServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client MerchantServiceClient) error {

	mux.Handle("GET", pattern_MerchantService_GetProfile_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/pb.MerchantService/GetProfile", runtime.WithHTTPPathPattern("/merchant/profile/{user_id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_MerchantService_GetProfile_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_MerchantService_GetProfile_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_MerchantService_GetProfile_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 4, 1, 5, 2}, []string{"merchant", "profile", "user_id"}, ""))
)

var (
	forward_MerchantService_GetProfile_0 = runtime.ForwardResponseMessage
)
