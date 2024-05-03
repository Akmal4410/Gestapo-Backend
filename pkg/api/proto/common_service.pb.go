// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v5.26.1
// source: api/proto/common_service.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_common_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_common_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_api_proto_common_service_proto_rawDescGZIP(), []int{0}
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Status  bool   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_common_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_common_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_api_proto_common_service_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32           `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Status  bool            `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Message string          `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*UserResponse `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetUsersResponse) Reset() {
	*x = GetUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_common_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersResponse) ProtoMessage() {}

func (x *GetUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_common_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersResponse.ProtoReflect.Descriptor instead.
func (*GetUsersResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_common_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetUsersResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetUsersResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *GetUsersResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetUsersResponse) GetData() []*UserResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProfileImage *string                `protobuf:"bytes,2,opt,name=profile_image,json=profileImage,proto3,oneof" json:"profile_image,omitempty"`
	FullName     *string                `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3,oneof" json:"full_name,omitempty"`
	UserName     string                 `protobuf:"bytes,4,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Phone        *string                `protobuf:"bytes,5,opt,name=phone,proto3,oneof" json:"phone,omitempty"`
	Email        *string                `protobuf:"bytes,6,opt,name=email,proto3,oneof" json:"email,omitempty"`
	Dob          *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=dob,proto3,oneof" json:"dob,omitempty"`
	Gender       *string                `protobuf:"bytes,8,opt,name=gender,proto3,oneof" json:"gender,omitempty"`
	UserType     string                 `protobuf:"bytes,9,opt,name=user_type,json=userType,proto3" json:"user_type,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_common_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_common_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_common_service_proto_rawDescGZIP(), []int{3}
}

func (x *UserResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserResponse) GetProfileImage() string {
	if x != nil && x.ProfileImage != nil {
		return *x.ProfileImage
	}
	return ""
}

func (x *UserResponse) GetFullName() string {
	if x != nil && x.FullName != nil {
		return *x.FullName
	}
	return ""
}

func (x *UserResponse) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserResponse) GetPhone() string {
	if x != nil && x.Phone != nil {
		return *x.Phone
	}
	return ""
}

func (x *UserResponse) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *UserResponse) GetDob() *timestamppb.Timestamp {
	if x != nil {
		return x.Dob
	}
	return nil
}

func (x *UserResponse) GetGender() string {
	if x != nil && x.Gender != nil {
		return *x.Gender
	}
	return ""
}

func (x *UserResponse) GetUserType() string {
	if x != nil {
		return x.UserType
	}
	return ""
}

type GetProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MerchantId string `protobuf:"bytes,1,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
}

func (x *GetProductRequest) Reset() {
	*x = GetProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_common_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductRequest) ProtoMessage() {}

func (x *GetProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_common_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductRequest.ProtoReflect.Descriptor instead.
func (*GetProductRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_common_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetProductRequest) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

type ProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MerchantId    string    `protobuf:"bytes,2,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	ProductImages []string  `protobuf:"bytes,3,rep,name=product_images,json=productImages,proto3" json:"product_images,omitempty"`
	ProductName   string    `protobuf:"bytes,4,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	Description   string    `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	CategoryName  string    `protobuf:"bytes,6,opt,name=category_name,json=categoryName,proto3" json:"category_name,omitempty"`
	Size          []float64 `protobuf:"fixed64,7,rep,packed,name=size,proto3" json:"size,omitempty"`
	Price         float64   `protobuf:"fixed64,8,opt,name=price,proto3" json:"price,omitempty"`
	DiscountPrice float64   `protobuf:"fixed64,9,opt,name=discount_price,json=discountPrice,proto3" json:"discount_price,omitempty"`
}

func (x *ProductResponse) Reset() {
	*x = ProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_common_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductResponse) ProtoMessage() {}

func (x *ProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_common_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductResponse.ProtoReflect.Descriptor instead.
func (*ProductResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_common_service_proto_rawDescGZIP(), []int{5}
}

func (x *ProductResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductResponse) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

func (x *ProductResponse) GetProductImages() []string {
	if x != nil {
		return x.ProductImages
	}
	return nil
}

func (x *ProductResponse) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *ProductResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductResponse) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

func (x *ProductResponse) GetSize() []float64 {
	if x != nil {
		return x.Size
	}
	return nil
}

func (x *ProductResponse) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ProductResponse) GetDiscountPrice() float64 {
	if x != nil {
		return x.DiscountPrice
	}
	return 0
}

type GetProductsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32              `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Status  bool               `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Message string             `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*ProductResponse `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetProductsResponse) Reset() {
	*x = GetProductsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_common_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductsResponse) ProtoMessage() {}

func (x *GetProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_common_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductsResponse.ProtoReflect.Descriptor instead.
func (*GetProductsResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_common_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetProductsResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetProductsResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *GetProductsResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetProductsResponse) GetData() []*ProductResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetProductByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32            `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Status  bool             `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Message string           `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Data    *ProductResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetProductByIdResponse) Reset() {
	*x = GetProductByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_common_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductByIdResponse) ProtoMessage() {}

func (x *GetProductByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_common_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductByIdResponse.ProtoReflect.Descriptor instead.
func (*GetProductByIdResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_common_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetProductByIdResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetProductByIdResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *GetProductByIdResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetProductByIdResponse) GetData() *ProductResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_api_proto_common_service_proto protoreflect.FileDescriptor

var file_api_proto_common_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x50, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x7e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0xf1, 0x02, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a,
	0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x88,
	0x01, 0x01, 0x12, 0x31, 0x0a, 0x03, 0x64, 0x6f, 0x62, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x04, 0x52, 0x03, 0x64,
	0x6f, 0x62, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x88,
	0x01, 0x01, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x42,
	0x10, 0x0a, 0x0e, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x64, 0x6f, 0x62, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x34, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x22, 0xa4, 0x02, 0x0a,
	0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a,
	0x0d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x03, 0x28, 0x01,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x22, 0x84, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x27, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x87, 0x01, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x42, 0x0b, 0x5a, 0x09, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_common_service_proto_rawDescOnce sync.Once
	file_api_proto_common_service_proto_rawDescData = file_api_proto_common_service_proto_rawDesc
)

func file_api_proto_common_service_proto_rawDescGZIP() []byte {
	file_api_proto_common_service_proto_rawDescOnce.Do(func() {
		file_api_proto_common_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_common_service_proto_rawDescData)
	})
	return file_api_proto_common_service_proto_rawDescData
}

var file_api_proto_common_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_proto_common_service_proto_goTypes = []interface{}{
	(*Request)(nil),                // 0: pb.Request
	(*Response)(nil),               // 1: pb.Response
	(*GetUsersResponse)(nil),       // 2: pb.GetUsersResponse
	(*UserResponse)(nil),           // 3: pb.UserResponse
	(*GetProductRequest)(nil),      // 4: pb.GetProductRequest
	(*ProductResponse)(nil),        // 5: pb.ProductResponse
	(*GetProductsResponse)(nil),    // 6: pb.GetProductsResponse
	(*GetProductByIdResponse)(nil), // 7: pb.GetProductByIdResponse
	(*timestamppb.Timestamp)(nil),  // 8: google.protobuf.Timestamp
}
var file_api_proto_common_service_proto_depIdxs = []int32{
	3, // 0: pb.GetUsersResponse.data:type_name -> pb.UserResponse
	8, // 1: pb.UserResponse.dob:type_name -> google.protobuf.Timestamp
	5, // 2: pb.GetProductsResponse.data:type_name -> pb.ProductResponse
	5, // 3: pb.GetProductByIdResponse.data:type_name -> pb.ProductResponse
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_proto_common_service_proto_init() }
func file_api_proto_common_service_proto_init() {
	if File_api_proto_common_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_common_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_common_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_common_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_common_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_common_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_common_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_common_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_common_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductByIdResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_api_proto_common_service_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_common_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_common_service_proto_goTypes,
		DependencyIndexes: file_api_proto_common_service_proto_depIdxs,
		MessageInfos:      file_api_proto_common_service_proto_msgTypes,
	}.Build()
	File_api_proto_common_service_proto = out.File
	file_api_proto_common_service_proto_rawDesc = nil
	file_api_proto_common_service_proto_goTypes = nil
	file_api_proto_common_service_proto_depIdxs = nil
}
