// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v5.26.1
// source: api/proto/user_service.proto

package proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetHomeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32         `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Status  bool          `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Message string        `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Data    *HomeResponse `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetHomeResponse) Reset() {
	*x = GetHomeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_user_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHomeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHomeResponse) ProtoMessage() {}

func (x *GetHomeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_user_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHomeResponse.ProtoReflect.Descriptor instead.
func (*GetHomeResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_user_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetHomeResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetHomeResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *GetHomeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetHomeResponse) GetData() *HomeResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

type HomeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Discount  *DiscountResponse  `protobuf:"bytes,1,opt,name=discount,proto3" json:"discount,omitempty"`
	Merchants []*UserResponse    `protobuf:"bytes,2,rep,name=merchants,proto3" json:"merchants,omitempty"`
	Products  []*ProductResponse `protobuf:"bytes,3,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *HomeResponse) Reset() {
	*x = HomeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_user_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeResponse) ProtoMessage() {}

func (x *HomeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_user_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeResponse.ProtoReflect.Descriptor instead.
func (*HomeResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_user_service_proto_rawDescGZIP(), []int{1}
}

func (x *HomeResponse) GetDiscount() *DiscountResponse {
	if x != nil {
		return x.Discount
	}
	return nil
}

func (x *HomeResponse) GetMerchants() []*UserResponse {
	if x != nil {
		return x.Merchants
	}
	return nil
}

func (x *HomeResponse) GetProducts() []*ProductResponse {
	if x != nil {
		return x.Products
	}
	return nil
}

var File_api_proto_user_service_proto protoreflect.FileDescriptor

var file_api_proto_user_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x1a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x26, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xa1, 0x01, 0x0a, 0x0c, 0x48, 0x6f,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70,
	0x62, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x09,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x09, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x2f, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x32, 0x4d, 0x0a,
	0x0a, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c,
	0x12, 0x0a, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x42, 0x0b, 0x5a, 0x09,
	0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_proto_user_service_proto_rawDescOnce sync.Once
	file_api_proto_user_service_proto_rawDescData = file_api_proto_user_service_proto_rawDesc
)

func file_api_proto_user_service_proto_rawDescGZIP() []byte {
	file_api_proto_user_service_proto_rawDescOnce.Do(func() {
		file_api_proto_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_user_service_proto_rawDescData)
	})
	return file_api_proto_user_service_proto_rawDescData
}

var file_api_proto_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_proto_user_service_proto_goTypes = []interface{}{
	(*GetHomeResponse)(nil),  // 0: pb.GetHomeResponse
	(*HomeResponse)(nil),     // 1: pb.HomeResponse
	(*DiscountResponse)(nil), // 2: pb.DiscountResponse
	(*UserResponse)(nil),     // 3: pb.UserResponse
	(*ProductResponse)(nil),  // 4: pb.ProductResponse
	(*Request)(nil),          // 5: pb.Request
}
var file_api_proto_user_service_proto_depIdxs = []int32{
	1, // 0: pb.GetHomeResponse.data:type_name -> pb.HomeResponse
	2, // 1: pb.HomeResponse.discount:type_name -> pb.DiscountResponse
	3, // 2: pb.HomeResponse.merchants:type_name -> pb.UserResponse
	4, // 3: pb.HomeResponse.products:type_name -> pb.ProductResponse
	5, // 4: pb.UserServie.GetHome:input_type -> pb.Request
	0, // 5: pb.UserServie.GetHome:output_type -> pb.GetHomeResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_proto_user_service_proto_init() }
func file_api_proto_user_service_proto_init() {
	if File_api_proto_user_service_proto != nil {
		return
	}
	file_api_proto_common_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_proto_user_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHomeResponse); i {
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
		file_api_proto_user_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_user_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_user_service_proto_goTypes,
		DependencyIndexes: file_api_proto_user_service_proto_depIdxs,
		MessageInfos:      file_api_proto_user_service_proto_msgTypes,
	}.Build()
	File_api_proto_user_service_proto = out.File
	file_api_proto_user_service_proto_rawDesc = nil
	file_api_proto_user_service_proto_goTypes = nil
	file_api_proto_user_service_proto_depIdxs = nil
}
