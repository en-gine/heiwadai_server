// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: v1/user/Banner.proto

package user

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Banner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageURL string `protobuf:"bytes,1,opt,name=ImageURL,proto3" json:"ImageURL,omitempty"`
	URL      string `protobuf:"bytes,2,opt,name=URL,proto3" json:"URL,omitempty"`
}

func (x *Banner) Reset() {
	*x = Banner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_user_Banner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Banner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Banner) ProtoMessage() {}

func (x *Banner) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_Banner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Banner.ProtoReflect.Descriptor instead.
func (*Banner) Descriptor() ([]byte, []int) {
	return file_v1_user_Banner_proto_rawDescGZIP(), []int{0}
}

func (x *Banner) GetImageURL() string {
	if x != nil {
		return x.ImageURL
	}
	return ""
}

func (x *Banner) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

type BannerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Banners []*Banner `protobuf:"bytes,1,rep,name=banners,proto3" json:"banners,omitempty"`
}

func (x *BannerResponse) Reset() {
	*x = BannerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_user_Banner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BannerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BannerResponse) ProtoMessage() {}

func (x *BannerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_Banner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BannerResponse.ProtoReflect.Descriptor instead.
func (*BannerResponse) Descriptor() ([]byte, []int) {
	return file_v1_user_Banner_proto_rawDescGZIP(), []int{1}
}

func (x *BannerResponse) GetBanners() []*Banner {
	if x != nil {
		return x.Banners
	}
	return nil
}

var File_v1_user_Banner_proto protoreflect.FileDescriptor

var file_v1_user_Banner_proto_rawDesc = []byte{
	0x0a, 0x14, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x36, 0x0a, 0x06, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x22, 0x3f, 0x0a, 0x0e, 0x42, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x62, 0x61,
	0x6e, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x52, 0x07, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x32, 0x56, 0x0a, 0x10, 0x42, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x42, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x7f, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x42, 0x0b, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0xa2, 0x02, 0x03, 0x53, 0x55, 0x58, 0xaa, 0x02, 0x0b,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0xca, 0x02, 0x0b, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5c, 0x55, 0x73, 0x65, 0x72, 0xe2, 0x02, 0x17, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5c, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3a, 0x3a, 0x55, 0x73,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_user_Banner_proto_rawDescOnce sync.Once
	file_v1_user_Banner_proto_rawDescData = file_v1_user_Banner_proto_rawDesc
)

func file_v1_user_Banner_proto_rawDescGZIP() []byte {
	file_v1_user_Banner_proto_rawDescOnce.Do(func() {
		file_v1_user_Banner_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_user_Banner_proto_rawDescData)
	})
	return file_v1_user_Banner_proto_rawDescData
}

var file_v1_user_Banner_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_user_Banner_proto_goTypes = []interface{}{
	(*Banner)(nil),         // 0: server.user.Banner
	(*BannerResponse)(nil), // 1: server.user.BannerResponse
	(*emptypb.Empty)(nil),  // 2: google.protobuf.Empty
}
var file_v1_user_Banner_proto_depIdxs = []int32{
	0, // 0: server.user.BannerResponse.banners:type_name -> server.user.Banner
	2, // 1: server.user.BannerController.GetBanner:input_type -> google.protobuf.Empty
	1, // 2: server.user.BannerController.GetBanner:output_type -> server.user.BannerResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_user_Banner_proto_init() }
func file_v1_user_Banner_proto_init() {
	if File_v1_user_Banner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_user_Banner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Banner); i {
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
		file_v1_user_Banner_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BannerResponse); i {
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
			RawDescriptor: file_v1_user_Banner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_user_Banner_proto_goTypes,
		DependencyIndexes: file_v1_user_Banner_proto_depIdxs,
		MessageInfos:      file_v1_user_Banner_proto_msgTypes,
	}.Build()
	File_v1_user_Banner_proto = out.File
	file_v1_user_Banner_proto_rawDesc = nil
	file_v1_user_Banner_proto_goTypes = nil
	file_v1_user_Banner_proto_depIdxs = nil
}
