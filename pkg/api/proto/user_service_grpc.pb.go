// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: api/proto/user_service.proto

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
	UserServie_GetHome_FullMethodName               = "/pb.UserServie/GetHome"
	UserServie_AddRemoveWishlist_FullMethodName     = "/pb.UserServie/AddRemoveWishlist"
	UserServie_GetWishlist_FullMethodName           = "/pb.UserServie/GetWishlist"
	UserServie_AddProductToCart_FullMethodName      = "/pb.UserServie/AddProductToCart"
	UserServie_GetCartItmes_FullMethodName          = "/pb.UserServie/GetCartItmes"
	UserServie_CheckoutCartItems_FullMethodName     = "/pb.UserServie/CheckoutCartItems"
	UserServie_RemoveProductFromCart_FullMethodName = "/pb.UserServie/RemoveProductFromCart"
	UserServie_AddAddress_FullMethodName            = "/pb.UserServie/AddAddress"
	UserServie_GetAddresses_FullMethodName          = "/pb.UserServie/GetAddresses"
	UserServie_GetAddressByID_FullMethodName        = "/pb.UserServie/GetAddressByID"
	UserServie_EditAddress_FullMethodName           = "/pb.UserServie/EditAddress"
	UserServie_DeleteAddress_FullMethodName         = "/pb.UserServie/DeleteAddress"
	UserServie_CreateOrder_FullMethodName           = "/pb.UserServie/CreateOrder"
	UserServie_GetUserOrders_FullMethodName         = "/pb.UserServie/GetUserOrders"
)

// UserServieClient is the client API for UserServie service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServieClient interface {
	GetHome(ctx context.Context, in *Request, opts ...grpc.CallOption) (*GetHomeResponse, error)
	AddRemoveWishlist(ctx context.Context, in *AddRemoveWishlistRequest, opts ...grpc.CallOption) (*Response, error)
	GetWishlist(ctx context.Context, in *Request, opts ...grpc.CallOption) (*GetWishlistResponse, error)
	// ------ Cart Related------------
	AddProductToCart(ctx context.Context, in *AddToCartRequest, opts ...grpc.CallOption) (*Response, error)
	GetCartItmes(ctx context.Context, in *Request, opts ...grpc.CallOption) (*GetCartItemsResponse, error)
	CheckoutCartItems(ctx context.Context, in *CheckoutCartItemsRequest, opts ...grpc.CallOption) (*Response, error)
	RemoveProductFromCart(ctx context.Context, in *RemoveFromCartRequest, opts ...grpc.CallOption) (*Response, error)
	// ------ Address Related------------
	AddAddress(ctx context.Context, in *AddAddressRequest, opts ...grpc.CallOption) (*Response, error)
	GetAddresses(ctx context.Context, in *Request, opts ...grpc.CallOption) (*GetAddressesResponse, error)
	GetAddressByID(ctx context.Context, in *AddressIdRequest, opts ...grpc.CallOption) (*GetAddressByIdResponse, error)
	EditAddress(ctx context.Context, in *EditAddressRequest, opts ...grpc.CallOption) (*Response, error)
	DeleteAddress(ctx context.Context, in *AddressIdRequest, opts ...grpc.CallOption) (*Response, error)
	// ------ Order Related------------
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*Response, error)
	GetUserOrders(ctx context.Context, in *GetUserOrdersRequest, opts ...grpc.CallOption) (*GetUserOrderResponse, error)
}

type userServieClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServieClient(cc grpc.ClientConnInterface) UserServieClient {
	return &userServieClient{cc}
}

func (c *userServieClient) GetHome(ctx context.Context, in *Request, opts ...grpc.CallOption) (*GetHomeResponse, error) {
	out := new(GetHomeResponse)
	err := c.cc.Invoke(ctx, UserServie_GetHome_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServieClient) AddRemoveWishlist(ctx context.Context, in *AddRemoveWishlistRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, UserServie_AddRemoveWishlist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServieClient) GetWishlist(ctx context.Context, in *Request, opts ...grpc.CallOption) (*GetWishlistResponse, error) {
	out := new(GetWishlistResponse)
	err := c.cc.Invoke(ctx, UserServie_GetWishlist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServieClient) AddProductToCart(ctx context.Context, in *AddToCartRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, UserServie_AddProductToCart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServieClient) GetCartItmes(ctx context.Context, in *Request, opts ...grpc.CallOption) (*GetCartItemsResponse, error) {
	out := new(GetCartItemsResponse)
	err := c.cc.Invoke(ctx, UserServie_GetCartItmes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServieClient) CheckoutCartItems(ctx context.Context, in *CheckoutCartItemsRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, UserServie_CheckoutCartItems_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServieClient) RemoveProductFromCart(ctx context.Context, in *RemoveFromCartRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, UserServie_RemoveProductFromCart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServieClient) AddAddress(ctx context.Context, in *AddAddressRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, UserServie_AddAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServieClient) GetAddresses(ctx context.Context, in *Request, opts ...grpc.CallOption) (*GetAddressesResponse, error) {
	out := new(GetAddressesResponse)
	err := c.cc.Invoke(ctx, UserServie_GetAddresses_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServieClient) GetAddressByID(ctx context.Context, in *AddressIdRequest, opts ...grpc.CallOption) (*GetAddressByIdResponse, error) {
	out := new(GetAddressByIdResponse)
	err := c.cc.Invoke(ctx, UserServie_GetAddressByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServieClient) EditAddress(ctx context.Context, in *EditAddressRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, UserServie_EditAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServieClient) DeleteAddress(ctx context.Context, in *AddressIdRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, UserServie_DeleteAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServieClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, UserServie_CreateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServieClient) GetUserOrders(ctx context.Context, in *GetUserOrdersRequest, opts ...grpc.CallOption) (*GetUserOrderResponse, error) {
	out := new(GetUserOrderResponse)
	err := c.cc.Invoke(ctx, UserServie_GetUserOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServieServer is the server API for UserServie service.
// All implementations must embed UnimplementedUserServieServer
// for forward compatibility
type UserServieServer interface {
	GetHome(context.Context, *Request) (*GetHomeResponse, error)
	AddRemoveWishlist(context.Context, *AddRemoveWishlistRequest) (*Response, error)
	GetWishlist(context.Context, *Request) (*GetWishlistResponse, error)
	// ------ Cart Related------------
	AddProductToCart(context.Context, *AddToCartRequest) (*Response, error)
	GetCartItmes(context.Context, *Request) (*GetCartItemsResponse, error)
	CheckoutCartItems(context.Context, *CheckoutCartItemsRequest) (*Response, error)
	RemoveProductFromCart(context.Context, *RemoveFromCartRequest) (*Response, error)
	// ------ Address Related------------
	AddAddress(context.Context, *AddAddressRequest) (*Response, error)
	GetAddresses(context.Context, *Request) (*GetAddressesResponse, error)
	GetAddressByID(context.Context, *AddressIdRequest) (*GetAddressByIdResponse, error)
	EditAddress(context.Context, *EditAddressRequest) (*Response, error)
	DeleteAddress(context.Context, *AddressIdRequest) (*Response, error)
	// ------ Order Related------------
	CreateOrder(context.Context, *CreateOrderRequest) (*Response, error)
	GetUserOrders(context.Context, *GetUserOrdersRequest) (*GetUserOrderResponse, error)
	mustEmbedUnimplementedUserServieServer()
}

// UnimplementedUserServieServer must be embedded to have forward compatible implementations.
type UnimplementedUserServieServer struct {
}

func (UnimplementedUserServieServer) GetHome(context.Context, *Request) (*GetHomeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHome not implemented")
}
func (UnimplementedUserServieServer) AddRemoveWishlist(context.Context, *AddRemoveWishlistRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRemoveWishlist not implemented")
}
func (UnimplementedUserServieServer) GetWishlist(context.Context, *Request) (*GetWishlistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWishlist not implemented")
}
func (UnimplementedUserServieServer) AddProductToCart(context.Context, *AddToCartRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProductToCart not implemented")
}
func (UnimplementedUserServieServer) GetCartItmes(context.Context, *Request) (*GetCartItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCartItmes not implemented")
}
func (UnimplementedUserServieServer) CheckoutCartItems(context.Context, *CheckoutCartItemsRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckoutCartItems not implemented")
}
func (UnimplementedUserServieServer) RemoveProductFromCart(context.Context, *RemoveFromCartRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveProductFromCart not implemented")
}
func (UnimplementedUserServieServer) AddAddress(context.Context, *AddAddressRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAddress not implemented")
}
func (UnimplementedUserServieServer) GetAddresses(context.Context, *Request) (*GetAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddresses not implemented")
}
func (UnimplementedUserServieServer) GetAddressByID(context.Context, *AddressIdRequest) (*GetAddressByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddressByID not implemented")
}
func (UnimplementedUserServieServer) EditAddress(context.Context, *EditAddressRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditAddress not implemented")
}
func (UnimplementedUserServieServer) DeleteAddress(context.Context, *AddressIdRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAddress not implemented")
}
func (UnimplementedUserServieServer) CreateOrder(context.Context, *CreateOrderRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedUserServieServer) GetUserOrders(context.Context, *GetUserOrdersRequest) (*GetUserOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserOrders not implemented")
}
func (UnimplementedUserServieServer) mustEmbedUnimplementedUserServieServer() {}

// UnsafeUserServieServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServieServer will
// result in compilation errors.
type UnsafeUserServieServer interface {
	mustEmbedUnimplementedUserServieServer()
}

func RegisterUserServieServer(s grpc.ServiceRegistrar, srv UserServieServer) {
	s.RegisterService(&UserServie_ServiceDesc, srv)
}

func _UserServie_GetHome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServieServer).GetHome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServie_GetHome_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServieServer).GetHome(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServie_AddRemoveWishlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRemoveWishlistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServieServer).AddRemoveWishlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServie_AddRemoveWishlist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServieServer).AddRemoveWishlist(ctx, req.(*AddRemoveWishlistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServie_GetWishlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServieServer).GetWishlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServie_GetWishlist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServieServer).GetWishlist(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServie_AddProductToCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddToCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServieServer).AddProductToCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServie_AddProductToCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServieServer).AddProductToCart(ctx, req.(*AddToCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServie_GetCartItmes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServieServer).GetCartItmes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServie_GetCartItmes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServieServer).GetCartItmes(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServie_CheckoutCartItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckoutCartItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServieServer).CheckoutCartItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServie_CheckoutCartItems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServieServer).CheckoutCartItems(ctx, req.(*CheckoutCartItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServie_RemoveProductFromCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFromCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServieServer).RemoveProductFromCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServie_RemoveProductFromCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServieServer).RemoveProductFromCart(ctx, req.(*RemoveFromCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServie_AddAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServieServer).AddAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServie_AddAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServieServer).AddAddress(ctx, req.(*AddAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServie_GetAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServieServer).GetAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServie_GetAddresses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServieServer).GetAddresses(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServie_GetAddressByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServieServer).GetAddressByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServie_GetAddressByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServieServer).GetAddressByID(ctx, req.(*AddressIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServie_EditAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServieServer).EditAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServie_EditAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServieServer).EditAddress(ctx, req.(*EditAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServie_DeleteAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServieServer).DeleteAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServie_DeleteAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServieServer).DeleteAddress(ctx, req.(*AddressIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServie_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServieServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServie_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServieServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServie_GetUserOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServieServer).GetUserOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServie_GetUserOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServieServer).GetUserOrders(ctx, req.(*GetUserOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserServie_ServiceDesc is the grpc.ServiceDesc for UserServie service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserServie_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserServie",
	HandlerType: (*UserServieServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHome",
			Handler:    _UserServie_GetHome_Handler,
		},
		{
			MethodName: "AddRemoveWishlist",
			Handler:    _UserServie_AddRemoveWishlist_Handler,
		},
		{
			MethodName: "GetWishlist",
			Handler:    _UserServie_GetWishlist_Handler,
		},
		{
			MethodName: "AddProductToCart",
			Handler:    _UserServie_AddProductToCart_Handler,
		},
		{
			MethodName: "GetCartItmes",
			Handler:    _UserServie_GetCartItmes_Handler,
		},
		{
			MethodName: "CheckoutCartItems",
			Handler:    _UserServie_CheckoutCartItems_Handler,
		},
		{
			MethodName: "RemoveProductFromCart",
			Handler:    _UserServie_RemoveProductFromCart_Handler,
		},
		{
			MethodName: "AddAddress",
			Handler:    _UserServie_AddAddress_Handler,
		},
		{
			MethodName: "GetAddresses",
			Handler:    _UserServie_GetAddresses_Handler,
		},
		{
			MethodName: "GetAddressByID",
			Handler:    _UserServie_GetAddressByID_Handler,
		},
		{
			MethodName: "EditAddress",
			Handler:    _UserServie_EditAddress_Handler,
		},
		{
			MethodName: "DeleteAddress",
			Handler:    _UserServie_DeleteAddress_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _UserServie_CreateOrder_Handler,
		},
		{
			MethodName: "GetUserOrders",
			Handler:    _UserServie_GetUserOrders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/user_service.proto",
}
