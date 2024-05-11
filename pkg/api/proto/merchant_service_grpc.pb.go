// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: api/proto/merchant_service.proto

package proto

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

const (
	MerchantService_GetProfile_FullMethodName          = "/pb.MerchantService/GetProfile"
	MerchantService_GetProducts_FullMethodName         = "/pb.MerchantService/GetProducts"
	MerchantService_DeleteProduct_FullMethodName       = "/pb.MerchantService/DeleteProduct"
	MerchantService_AddProductDiscount_FullMethodName  = "/pb.MerchantService/AddProductDiscount"
	MerchantService_EditProductDiscount_FullMethodName = "/pb.MerchantService/EditProductDiscount"
	MerchantService_GetAllDiscounts_FullMethodName     = "/pb.MerchantService/GetAllDiscounts"
	MerchantService_GetMerchantOrders_FullMethodName   = "/pb.MerchantService/GetMerchantOrders"
	MerchantService_UpdateOrderStatus_FullMethodName   = "/pb.MerchantService/UpdateOrderStatus"
)

// MerchantServiceClient is the client API for MerchantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MerchantServiceClient interface {
	GetProfile(ctx context.Context, in *GetMerchantProfileRequest, opts ...grpc.CallOption) (*GetMerchantProfileResponse, error)
	GetProducts(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductsResponse, error)
	DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*Response, error)
	AddProductDiscount(ctx context.Context, in *AddDiscountRequest, opts ...grpc.CallOption) (*Response, error)
	EditProductDiscount(ctx context.Context, in *EditDiscountRequest, opts ...grpc.CallOption) (*Response, error)
	GetAllDiscounts(ctx context.Context, in *GetDiscountsRequest, opts ...grpc.CallOption) (*GetDiscountsResponse, error)
	// ------ Order Related------------
	GetMerchantOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrderResponse, error)
	UpdateOrderStatus(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*Response, error)
}

type merchantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMerchantServiceClient(cc grpc.ClientConnInterface) MerchantServiceClient {
	return &merchantServiceClient{cc}
}

func (c *merchantServiceClient) GetProfile(ctx context.Context, in *GetMerchantProfileRequest, opts ...grpc.CallOption) (*GetMerchantProfileResponse, error) {
	out := new(GetMerchantProfileResponse)
	err := c.cc.Invoke(ctx, MerchantService_GetProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServiceClient) GetProducts(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductsResponse, error) {
	out := new(GetProductsResponse)
	err := c.cc.Invoke(ctx, MerchantService_GetProducts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServiceClient) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, MerchantService_DeleteProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServiceClient) AddProductDiscount(ctx context.Context, in *AddDiscountRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, MerchantService_AddProductDiscount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServiceClient) EditProductDiscount(ctx context.Context, in *EditDiscountRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, MerchantService_EditProductDiscount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServiceClient) GetAllDiscounts(ctx context.Context, in *GetDiscountsRequest, opts ...grpc.CallOption) (*GetDiscountsResponse, error) {
	out := new(GetDiscountsResponse)
	err := c.cc.Invoke(ctx, MerchantService_GetAllDiscounts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServiceClient) GetMerchantOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrderResponse, error) {
	out := new(GetOrderResponse)
	err := c.cc.Invoke(ctx, MerchantService_GetMerchantOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServiceClient) UpdateOrderStatus(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, MerchantService_UpdateOrderStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MerchantServiceServer is the server API for MerchantService service.
// All implementations must embed UnimplementedMerchantServiceServer
// for forward compatibility
type MerchantServiceServer interface {
	GetProfile(context.Context, *GetMerchantProfileRequest) (*GetMerchantProfileResponse, error)
	GetProducts(context.Context, *GetProductRequest) (*GetProductsResponse, error)
	DeleteProduct(context.Context, *DeleteProductRequest) (*Response, error)
	AddProductDiscount(context.Context, *AddDiscountRequest) (*Response, error)
	EditProductDiscount(context.Context, *EditDiscountRequest) (*Response, error)
	GetAllDiscounts(context.Context, *GetDiscountsRequest) (*GetDiscountsResponse, error)
	// ------ Order Related------------
	GetMerchantOrders(context.Context, *GetOrdersRequest) (*GetOrderResponse, error)
	UpdateOrderStatus(context.Context, *UpdateOrderRequest) (*Response, error)
	mustEmbedUnimplementedMerchantServiceServer()
}

// UnimplementedMerchantServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMerchantServiceServer struct {
}

func (UnimplementedMerchantServiceServer) GetProfile(context.Context, *GetMerchantProfileRequest) (*GetMerchantProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedMerchantServiceServer) GetProducts(context.Context, *GetProductRequest) (*GetProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProducts not implemented")
}
func (UnimplementedMerchantServiceServer) DeleteProduct(context.Context, *DeleteProductRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedMerchantServiceServer) AddProductDiscount(context.Context, *AddDiscountRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProductDiscount not implemented")
}
func (UnimplementedMerchantServiceServer) EditProductDiscount(context.Context, *EditDiscountRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditProductDiscount not implemented")
}
func (UnimplementedMerchantServiceServer) GetAllDiscounts(context.Context, *GetDiscountsRequest) (*GetDiscountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllDiscounts not implemented")
}
func (UnimplementedMerchantServiceServer) GetMerchantOrders(context.Context, *GetOrdersRequest) (*GetOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMerchantOrders not implemented")
}
func (UnimplementedMerchantServiceServer) UpdateOrderStatus(context.Context, *UpdateOrderRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrderStatus not implemented")
}
func (UnimplementedMerchantServiceServer) mustEmbedUnimplementedMerchantServiceServer() {}

// UnsafeMerchantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MerchantServiceServer will
// result in compilation errors.
type UnsafeMerchantServiceServer interface {
	mustEmbedUnimplementedMerchantServiceServer()
}

func RegisterMerchantServiceServer(s grpc.ServiceRegistrar, srv MerchantServiceServer) {
	s.RegisterService(&MerchantService_ServiceDesc, srv)
}

func _MerchantService_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMerchantProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MerchantService_GetProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).GetProfile(ctx, req.(*GetMerchantProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantService_GetProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).GetProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MerchantService_GetProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).GetProducts(ctx, req.(*GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MerchantService_DeleteProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).DeleteProduct(ctx, req.(*DeleteProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantService_AddProductDiscount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDiscountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).AddProductDiscount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MerchantService_AddProductDiscount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).AddProductDiscount(ctx, req.(*AddDiscountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantService_EditProductDiscount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditDiscountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).EditProductDiscount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MerchantService_EditProductDiscount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).EditProductDiscount(ctx, req.(*EditDiscountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantService_GetAllDiscounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDiscountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).GetAllDiscounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MerchantService_GetAllDiscounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).GetAllDiscounts(ctx, req.(*GetDiscountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantService_GetMerchantOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).GetMerchantOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MerchantService_GetMerchantOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).GetMerchantOrders(ctx, req.(*GetOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantService_UpdateOrderStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).UpdateOrderStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MerchantService_UpdateOrderStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).UpdateOrderStatus(ctx, req.(*UpdateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MerchantService_ServiceDesc is the grpc.ServiceDesc for MerchantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MerchantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MerchantService",
	HandlerType: (*MerchantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProfile",
			Handler:    _MerchantService_GetProfile_Handler,
		},
		{
			MethodName: "GetProducts",
			Handler:    _MerchantService_GetProducts_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _MerchantService_DeleteProduct_Handler,
		},
		{
			MethodName: "AddProductDiscount",
			Handler:    _MerchantService_AddProductDiscount_Handler,
		},
		{
			MethodName: "EditProductDiscount",
			Handler:    _MerchantService_EditProductDiscount_Handler,
		},
		{
			MethodName: "GetAllDiscounts",
			Handler:    _MerchantService_GetAllDiscounts_Handler,
		},
		{
			MethodName: "GetMerchantOrders",
			Handler:    _MerchantService_GetMerchantOrders_Handler,
		},
		{
			MethodName: "UpdateOrderStatus",
			Handler:    _MerchantService_UpdateOrderStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/merchant_service.proto",
}
