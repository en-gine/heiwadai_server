// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: v1/user/Messages.proto

package user

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

type MessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *MessageRequest) Reset() {
	*x = MessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_user_Messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageRequest) ProtoMessage() {}

func (x *MessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_Messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageRequest.ProtoReflect.Descriptor instead.
func (*MessageRequest) Descriptor() ([]byte, []int) {
	return file_v1_user_Messages_proto_rawDescGZIP(), []int{0}
}

func (x *MessageRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type MessagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*MessageResponse `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *MessagesResponse) Reset() {
	*x = MessagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_user_Messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessagesResponse) ProtoMessage() {}

func (x *MessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_Messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessagesResponse.ProtoReflect.Descriptor instead.
func (*MessagesResponse) Descriptor() ([]byte, []int) {
	return file_v1_user_Messages_proto_rawDescGZIP(), []int{1}
}

func (x *MessagesResponse) GetMessages() []*MessageResponse {
	if x != nil {
		return x.Messages
	}
	return nil
}

type MessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title       string                 `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Content     string                 `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty"`
	AuthorID    string                 `protobuf:"bytes,4,opt,name=AuthorID,proto3" json:"AuthorID,omitempty"`
	DisplayDate *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=DisplayDate,proto3" json:"DisplayDate,omitempty"`
}

func (x *MessageResponse) Reset() {
	*x = MessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_user_Messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageResponse) ProtoMessage() {}

func (x *MessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_Messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageResponse.ProtoReflect.Descriptor instead.
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return file_v1_user_Messages_proto_rawDescGZIP(), []int{2}
}

func (x *MessageResponse) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *MessageResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MessageResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *MessageResponse) GetAuthorID() string {
	if x != nil {
		return x.AuthorID
	}
	return ""
}

func (x *MessageResponse) GetDisplayDate() *timestamppb.Timestamp {
	if x != nil {
		return x.DisplayDate
	}
	return nil
}

var File_v1_user_Messages_proto protoreflect.FileDescriptor

var file_v1_user_Messages_proto_rawDesc = []byte{
	0x0a, 0x16, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x4c, 0x0a, 0x10, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x08,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0xab, 0x01, 0x0a, 0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x49, 0x44, 0x12, 0x3c, 0x0a, 0x0b, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x44, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x44, 0x61, 0x74, 0x65, 0x32, 0x65, 0x0a, 0x11, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x50, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x41, 0x66, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x81, 0x01, 0x0a, 0x0f,
	0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x42,
	0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0xa2, 0x02, 0x03, 0x53, 0x55, 0x58, 0xaa, 0x02, 0x0b, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0xca, 0x02, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5c, 0x55, 0x73, 0x65, 0x72, 0xe2, 0x02, 0x17, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c,
	0x55, 0x73, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3a, 0x3a, 0x55, 0x73, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_user_Messages_proto_rawDescOnce sync.Once
	file_v1_user_Messages_proto_rawDescData = file_v1_user_Messages_proto_rawDesc
)

func file_v1_user_Messages_proto_rawDescGZIP() []byte {
	file_v1_user_Messages_proto_rawDescOnce.Do(func() {
		file_v1_user_Messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_user_Messages_proto_rawDescData)
	})
	return file_v1_user_Messages_proto_rawDescData
}

var file_v1_user_Messages_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v1_user_Messages_proto_goTypes = []interface{}{
	(*MessageRequest)(nil),        // 0: server.user.MessageRequest
	(*MessagesResponse)(nil),      // 1: server.user.MessagesResponse
	(*MessageResponse)(nil),       // 2: server.user.MessageResponse
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_v1_user_Messages_proto_depIdxs = []int32{
	2, // 0: server.user.MessagesResponse.messages:type_name -> server.user.MessageResponse
	3, // 1: server.user.MessageResponse.DisplayDate:type_name -> google.protobuf.Timestamp
	0, // 2: server.user.MessageController.GetMessagesAfter:input_type -> server.user.MessageRequest
	1, // 3: server.user.MessageController.GetMessagesAfter:output_type -> server.user.MessagesResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_user_Messages_proto_init() }
func file_v1_user_Messages_proto_init() {
	if File_v1_user_Messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_user_Messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageRequest); i {
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
		file_v1_user_Messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessagesResponse); i {
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
		file_v1_user_Messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageResponse); i {
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
			RawDescriptor: file_v1_user_Messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_user_Messages_proto_goTypes,
		DependencyIndexes: file_v1_user_Messages_proto_depIdxs,
		MessageInfos:      file_v1_user_Messages_proto_msgTypes,
	}.Build()
	File_v1_user_Messages_proto = out.File
	file_v1_user_Messages_proto_rawDesc = nil
	file_v1_user_Messages_proto_goTypes = nil
	file_v1_user_Messages_proto_depIdxs = nil
}