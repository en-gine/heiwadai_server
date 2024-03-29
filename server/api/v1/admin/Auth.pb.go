// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: v1/admin/Auth.proto

package admin

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

type AdminRefreshTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
}

func (x *AdminRefreshTokenRequest) Reset() {
	*x = AdminRefreshTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_admin_Auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminRefreshTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminRefreshTokenRequest) ProtoMessage() {}

func (x *AdminRefreshTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_admin_Auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminRefreshTokenRequest.ProtoReflect.Descriptor instead.
func (*AdminRefreshTokenRequest) Descriptor() ([]byte, []int) {
	return file_v1_admin_Auth_proto_rawDescGZIP(), []int{0}
}

func (x *AdminRefreshTokenRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *AdminRefreshTokenRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type AdminAuthTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	ExpiresIn    int64  `protobuf:"varint,2,opt,name=expiresIn,proto3" json:"expiresIn,omitempty"`
	RefreshToken string `protobuf:"bytes,3,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
}

func (x *AdminAuthTokenResponse) Reset() {
	*x = AdminAuthTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_admin_Auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminAuthTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminAuthTokenResponse) ProtoMessage() {}

func (x *AdminAuthTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_admin_Auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminAuthTokenResponse.ProtoReflect.Descriptor instead.
func (*AdminAuthTokenResponse) Descriptor() ([]byte, []int) {
	return file_v1_admin_Auth_proto_rawDescGZIP(), []int{1}
}

func (x *AdminAuthTokenResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *AdminAuthTokenResponse) GetExpiresIn() int64 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

func (x *AdminAuthTokenResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type UpdatePasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UpdatePasswordRequest) Reset() {
	*x = UpdatePasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_admin_Auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePasswordRequest) ProtoMessage() {}

func (x *UpdatePasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_admin_Auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePasswordRequest.ProtoReflect.Descriptor instead.
func (*UpdatePasswordRequest) Descriptor() ([]byte, []int) {
	return file_v1_admin_Auth_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatePasswordRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UpdateEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mail string `protobuf:"bytes,1,opt,name=Mail,proto3" json:"Mail,omitempty"`
}

func (x *UpdateEmailRequest) Reset() {
	*x = UpdateEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_admin_Auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEmailRequest) ProtoMessage() {}

func (x *UpdateEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_admin_Auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEmailRequest.ProtoReflect.Descriptor instead.
func (*UpdateEmailRequest) Descriptor() ([]byte, []int) {
	return file_v1_admin_Auth_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateEmailRequest) GetMail() string {
	if x != nil {
		return x.Mail
	}
	return ""
}

type AdminRegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Mail          string `protobuf:"bytes,2,opt,name=Mail,proto3" json:"Mail,omitempty"`
	BelongStoreID string `protobuf:"bytes,4,opt,name=BelongStoreID,proto3" json:"BelongStoreID,omitempty"`
}

func (x *AdminRegisterRequest) Reset() {
	*x = AdminRegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_admin_Auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminRegisterRequest) ProtoMessage() {}

func (x *AdminRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_admin_Auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminRegisterRequest.ProtoReflect.Descriptor instead.
func (*AdminRegisterRequest) Descriptor() ([]byte, []int) {
	return file_v1_admin_Auth_proto_rawDescGZIP(), []int{4}
}

func (x *AdminRegisterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AdminRegisterRequest) GetMail() string {
	if x != nil {
		return x.Mail
	}
	return ""
}

func (x *AdminRegisterRequest) GetBelongStoreID() string {
	if x != nil {
		return x.BelongStoreID
	}
	return ""
}

type AdminRegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *AdminRegisterResponse) Reset() {
	*x = AdminRegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_admin_Auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminRegisterResponse) ProtoMessage() {}

func (x *AdminRegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_admin_Auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminRegisterResponse.ProtoReflect.Descriptor instead.
func (*AdminRegisterResponse) Descriptor() ([]byte, []int) {
	return file_v1_admin_Auth_proto_rawDescGZIP(), []int{5}
}

func (x *AdminRegisterResponse) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type ResendInviteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mail string `protobuf:"bytes,1,opt,name=Mail,proto3" json:"Mail,omitempty"`
}

func (x *ResendInviteRequest) Reset() {
	*x = ResendInviteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_admin_Auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResendInviteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResendInviteRequest) ProtoMessage() {}

func (x *ResendInviteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_admin_Auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResendInviteRequest.ProtoReflect.Descriptor instead.
func (*ResendInviteRequest) Descriptor() ([]byte, []int) {
	return file_v1_admin_Auth_proto_rawDescGZIP(), []int{6}
}

func (x *ResendInviteRequest) GetMail() string {
	if x != nil {
		return x.Mail
	}
	return ""
}

var File_v1_admin_Auth_proto protoreflect.FileDescriptor

var file_v1_admin_Auth_proto_rawDesc = []byte{
	0x0a, 0x13, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x41, 0x75, 0x74, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x60, 0x0a, 0x18, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22,
	0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x7c, 0x0a, 0x16, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c,
	0x0a, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x12, 0x22, 0x0a, 0x0c,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x33, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x28, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4d, 0x61, 0x69, 0x6c, 0x22,
	0x64, 0x0a, 0x14, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4d,
	0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4d, 0x61, 0x69, 0x6c, 0x12,
	0x24, 0x0a, 0x0d, 0x42, 0x65, 0x6c, 0x6f, 0x6e, 0x67, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x44,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x42, 0x65, 0x6c, 0x6f, 0x6e, 0x67, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x49, 0x44, 0x22, 0x27, 0x0a, 0x15, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x29,
	0x0a, 0x13, 0x52, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4d, 0x61, 0x69, 0x6c, 0x32, 0xec, 0x03, 0x0a, 0x0e, 0x41, 0x75,
	0x74, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x07,
	0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x07, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x12, 0x26, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x55, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x65, 0x6e,
	0x64, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x6e,
	0x64, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x83, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x42, 0x09, 0x41,
	0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x13, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0xa2,
	0x02, 0x03, 0x53, 0x41, 0x58, 0xaa, 0x02, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0xca, 0x02, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0xe2, 0x02, 0x18, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3a, 0x3a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_admin_Auth_proto_rawDescOnce sync.Once
	file_v1_admin_Auth_proto_rawDescData = file_v1_admin_Auth_proto_rawDesc
)

func file_v1_admin_Auth_proto_rawDescGZIP() []byte {
	file_v1_admin_Auth_proto_rawDescOnce.Do(func() {
		file_v1_admin_Auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_admin_Auth_proto_rawDescData)
	})
	return file_v1_admin_Auth_proto_rawDescData
}

var file_v1_admin_Auth_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_v1_admin_Auth_proto_goTypes = []interface{}{
	(*AdminRefreshTokenRequest)(nil), // 0: server.admin.AdminRefreshTokenRequest
	(*AdminAuthTokenResponse)(nil),   // 1: server.admin.AdminAuthTokenResponse
	(*UpdatePasswordRequest)(nil),    // 2: server.admin.UpdatePasswordRequest
	(*UpdateEmailRequest)(nil),       // 3: server.admin.UpdateEmailRequest
	(*AdminRegisterRequest)(nil),     // 4: server.admin.AdminRegisterRequest
	(*AdminRegisterResponse)(nil),    // 5: server.admin.AdminRegisterResponse
	(*ResendInviteRequest)(nil),      // 6: server.admin.ResendInviteRequest
	(*emptypb.Empty)(nil),            // 7: google.protobuf.Empty
}
var file_v1_admin_Auth_proto_depIdxs = []int32{
	7, // 0: server.admin.AuthController.SignOut:input_type -> google.protobuf.Empty
	0, // 1: server.admin.AuthController.Refresh:input_type -> server.admin.AdminRefreshTokenRequest
	2, // 2: server.admin.AuthController.UpdatePassword:input_type -> server.admin.UpdatePasswordRequest
	3, // 3: server.admin.AuthController.UpdateEmail:input_type -> server.admin.UpdateEmailRequest
	4, // 4: server.admin.AuthController.Register:input_type -> server.admin.AdminRegisterRequest
	6, // 5: server.admin.AuthController.ResendInviteMail:input_type -> server.admin.ResendInviteRequest
	7, // 6: server.admin.AuthController.SignOut:output_type -> google.protobuf.Empty
	1, // 7: server.admin.AuthController.Refresh:output_type -> server.admin.AdminAuthTokenResponse
	7, // 8: server.admin.AuthController.UpdatePassword:output_type -> google.protobuf.Empty
	7, // 9: server.admin.AuthController.UpdateEmail:output_type -> google.protobuf.Empty
	5, // 10: server.admin.AuthController.Register:output_type -> server.admin.AdminRegisterResponse
	7, // 11: server.admin.AuthController.ResendInviteMail:output_type -> google.protobuf.Empty
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_admin_Auth_proto_init() }
func file_v1_admin_Auth_proto_init() {
	if File_v1_admin_Auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_admin_Auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminRefreshTokenRequest); i {
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
		file_v1_admin_Auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminAuthTokenResponse); i {
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
		file_v1_admin_Auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePasswordRequest); i {
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
		file_v1_admin_Auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEmailRequest); i {
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
		file_v1_admin_Auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminRegisterRequest); i {
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
		file_v1_admin_Auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminRegisterResponse); i {
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
		file_v1_admin_Auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResendInviteRequest); i {
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
			RawDescriptor: file_v1_admin_Auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_admin_Auth_proto_goTypes,
		DependencyIndexes: file_v1_admin_Auth_proto_depIdxs,
		MessageInfos:      file_v1_admin_Auth_proto_msgTypes,
	}.Build()
	File_v1_admin_Auth_proto = out.File
	file_v1_admin_Auth_proto_rawDesc = nil
	file_v1_admin_Auth_proto_goTypes = nil
	file_v1_admin_Auth_proto_depIdxs = nil
}
