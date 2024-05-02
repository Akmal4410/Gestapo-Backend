// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v5.26.1
// source: api/proto/merchant_service.proto

package proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetMerchantProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetMerchantProfileRequest) Reset() {
	*x = GetMerchantProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_merchant_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMerchantProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMerchantProfileRequest) ProtoMessage() {}

func (x *GetMerchantProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_merchant_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMerchantProfileRequest.ProtoReflect.Descriptor instead.
func (*GetMerchantProfileRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_merchant_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetMerchantProfileRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type MerchantResponse struct {
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

func (x *MerchantResponse) Reset() {
	*x = MerchantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_merchant_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MerchantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MerchantResponse) ProtoMessage() {}

func (x *MerchantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_merchant_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MerchantResponse.ProtoReflect.Descriptor instead.
func (*MerchantResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_merchant_service_proto_rawDescGZIP(), []int{1}
}

func (x *MerchantResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MerchantResponse) GetProfileImage() string {
	if x != nil && x.ProfileImage != nil {
		return *x.ProfileImage
	}
	return ""
}

func (x *MerchantResponse) GetFullName() string {
	if x != nil && x.FullName != nil {
		return *x.FullName
	}
	return ""
}

func (x *MerchantResponse) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *MerchantResponse) GetPhone() string {
	if x != nil && x.Phone != nil {
		return *x.Phone
	}
	return ""
}

func (x *MerchantResponse) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *MerchantResponse) GetDob() *timestamppb.Timestamp {
	if x != nil {
		return x.Dob
	}
	return nil
}

func (x *MerchantResponse) GetGender() string {
	if x != nil && x.Gender != nil {
		return *x.Gender
	}
	return ""
}

func (x *MerchantResponse) GetUserType() string {
	if x != nil {
		return x.UserType
	}
	return ""
}

type GetMerchantProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Status  bool              `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Message string            `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Data    *MerchantResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetMerchantProfileResponse) Reset() {
	*x = GetMerchantProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_merchant_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMerchantProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMerchantProfileResponse) ProtoMessage() {}

func (x *GetMerchantProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_merchant_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMerchantProfileResponse.ProtoReflect.Descriptor instead.
func (*GetMerchantProfileResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_merchant_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetMerchantProfileResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetMerchantProfileResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *GetMerchantProfileResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetMerchantProfileResponse) GetData() *MerchantResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

type AddDiscountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId   string                 `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Percentage  uint64                 `protobuf:"fixed64,4,opt,name=percentage,proto3" json:"percentage,omitempty"`
	CardColor   string                 `protobuf:"bytes,5,opt,name=card_color,json=cardColor,proto3" json:"card_color,omitempty"`
	StartTime   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *AddDiscountRequest) Reset() {
	*x = AddDiscountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_merchant_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDiscountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDiscountRequest) ProtoMessage() {}

func (x *AddDiscountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_merchant_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDiscountRequest.ProtoReflect.Descriptor instead.
func (*AddDiscountRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_merchant_service_proto_rawDescGZIP(), []int{3}
}

func (x *AddDiscountRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *AddDiscountRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddDiscountRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AddDiscountRequest) GetPercentage() uint64 {
	if x != nil {
		return x.Percentage
	}
	return 0
}

func (x *AddDiscountRequest) GetCardColor() string {
	if x != nil {
		return x.CardColor
	}
	return ""
}

func (x *AddDiscountRequest) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *AddDiscountRequest) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

type EditDiscountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DiscountId  string                 `protobuf:"bytes,1,opt,name=discount_id,json=discountId,proto3" json:"discount_id,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Percentage  uint64                 `protobuf:"fixed64,4,opt,name=percentage,proto3" json:"percentage,omitempty"`
	CardColor   string                 `protobuf:"bytes,5,opt,name=card_color,json=cardColor,proto3" json:"card_color,omitempty"`
	StartTime   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *EditDiscountRequest) Reset() {
	*x = EditDiscountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_merchant_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditDiscountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditDiscountRequest) ProtoMessage() {}

func (x *EditDiscountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_merchant_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditDiscountRequest.ProtoReflect.Descriptor instead.
func (*EditDiscountRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_merchant_service_proto_rawDescGZIP(), []int{4}
}

func (x *EditDiscountRequest) GetDiscountId() string {
	if x != nil {
		return x.DiscountId
	}
	return ""
}

func (x *EditDiscountRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EditDiscountRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *EditDiscountRequest) GetPercentage() uint64 {
	if x != nil {
		return x.Percentage
	}
	return 0
}

func (x *EditDiscountRequest) GetCardColor() string {
	if x != nil {
		return x.CardColor
	}
	return ""
}

func (x *EditDiscountRequest) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *EditDiscountRequest) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

var File_api_proto_merchant_service_proto protoreflect.FileDescriptor

var file_api_proto_merchant_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x26, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x34, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xf5, 0x02, 0x0a, 0x10, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x0d, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x02, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x03, 0x64, 0x6f, 0x62,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x48, 0x04, 0x52, 0x03, 0x64, 0x6f, 0x62, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x66, 0x75, 0x6c,
	0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x64,
	0x6f, 0x62, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x8c, 0x01,
	0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x9a, 0x02, 0x0a,
	0x12, 0x41, 0x64, 0x64, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x06, 0x52, 0x0a, 0x70, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x64,
	0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61,
	0x72, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x9d, 0x02, 0x0a, 0x13, 0x45, 0x64,
	0x69, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x06, 0x52, 0x0a, 0x70, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x64,
	0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61,
	0x72, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xa9, 0x03, 0x0a, 0x0f, 0x4d, 0x65,
	0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x70, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2f, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x12,
	0x4e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x0b,
	0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x6d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12,
	0x61, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e,
	0x74, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x71, 0x0a, 0x13, 0x45, 0x64, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x45,
	0x64, 0x69, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x33, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x3a, 0x01, 0x2a, 0x32, 0x28, 0x2f, 0x6d, 0x65,
	0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x7b, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x0b, 0x5a, 0x09, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_merchant_service_proto_rawDescOnce sync.Once
	file_api_proto_merchant_service_proto_rawDescData = file_api_proto_merchant_service_proto_rawDesc
)

func file_api_proto_merchant_service_proto_rawDescGZIP() []byte {
	file_api_proto_merchant_service_proto_rawDescOnce.Do(func() {
		file_api_proto_merchant_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_merchant_service_proto_rawDescData)
	})
	return file_api_proto_merchant_service_proto_rawDescData
}

var file_api_proto_merchant_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_proto_merchant_service_proto_goTypes = []interface{}{
	(*GetMerchantProfileRequest)(nil),  // 0: pb.GetMerchantProfileRequest
	(*MerchantResponse)(nil),           // 1: pb.MerchantResponse
	(*GetMerchantProfileResponse)(nil), // 2: pb.GetMerchantProfileResponse
	(*AddDiscountRequest)(nil),         // 3: pb.AddDiscountRequest
	(*EditDiscountRequest)(nil),        // 4: pb.EditDiscountRequest
	(*timestamppb.Timestamp)(nil),      // 5: google.protobuf.Timestamp
	(*Request)(nil),                    // 6: pb.Request
	(*GetProductsResponse)(nil),        // 7: pb.GetProductsResponse
	(*Response)(nil),                   // 8: pb.Response
}
var file_api_proto_merchant_service_proto_depIdxs = []int32{
	5,  // 0: pb.MerchantResponse.dob:type_name -> google.protobuf.Timestamp
	1,  // 1: pb.GetMerchantProfileResponse.data:type_name -> pb.MerchantResponse
	5,  // 2: pb.AddDiscountRequest.start_time:type_name -> google.protobuf.Timestamp
	5,  // 3: pb.AddDiscountRequest.end_time:type_name -> google.protobuf.Timestamp
	5,  // 4: pb.EditDiscountRequest.start_time:type_name -> google.protobuf.Timestamp
	5,  // 5: pb.EditDiscountRequest.end_time:type_name -> google.protobuf.Timestamp
	0,  // 6: pb.MerchantService.GetProfile:input_type -> pb.GetMerchantProfileRequest
	6,  // 7: pb.MerchantService.GetProducts:input_type -> pb.Request
	3,  // 8: pb.MerchantService.AddProductDiscount:input_type -> pb.AddDiscountRequest
	4,  // 9: pb.MerchantService.EditProductDiscount:input_type -> pb.EditDiscountRequest
	2,  // 10: pb.MerchantService.GetProfile:output_type -> pb.GetMerchantProfileResponse
	7,  // 11: pb.MerchantService.GetProducts:output_type -> pb.GetProductsResponse
	8,  // 12: pb.MerchantService.AddProductDiscount:output_type -> pb.Response
	8,  // 13: pb.MerchantService.EditProductDiscount:output_type -> pb.Response
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_api_proto_merchant_service_proto_init() }
func file_api_proto_merchant_service_proto_init() {
	if File_api_proto_merchant_service_proto != nil {
		return
	}
	file_api_proto_common_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_proto_merchant_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMerchantProfileRequest); i {
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
		file_api_proto_merchant_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MerchantResponse); i {
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
		file_api_proto_merchant_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMerchantProfileResponse); i {
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
		file_api_proto_merchant_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDiscountRequest); i {
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
		file_api_proto_merchant_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditDiscountRequest); i {
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
	file_api_proto_merchant_service_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_merchant_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_merchant_service_proto_goTypes,
		DependencyIndexes: file_api_proto_merchant_service_proto_depIdxs,
		MessageInfos:      file_api_proto_merchant_service_proto_msgTypes,
	}.Build()
	File_api_proto_merchant_service_proto = out.File
	file_api_proto_merchant_service_proto_rawDesc = nil
	file_api_proto_merchant_service_proto_goTypes = nil
	file_api_proto_merchant_service_proto_depIdxs = nil
}
