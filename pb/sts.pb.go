// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: sts.proto

package pb

import (
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

type GenCosStsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *GenCosStsReq) Reset() {
	*x = GenCosStsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenCosStsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenCosStsReq) ProtoMessage() {}

func (x *GenCosStsReq) ProtoReflect() protoreflect.Message {
	mi := &file_sts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenCosStsReq.ProtoReflect.Descriptor instead.
func (*GenCosStsReq) Descriptor() ([]byte, []int) {
	return file_sts_proto_rawDescGZIP(), []int{0}
}

func (x *GenCosStsReq) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type GenCosStsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecretId     string `protobuf:"bytes,1,opt,name=secretId,proto3" json:"secretId,omitempty"`
	SecretKey    string `protobuf:"bytes,2,opt,name=secretKey,proto3" json:"secretKey,omitempty"`
	SessionToken string `protobuf:"bytes,3,opt,name=sessionToken,proto3" json:"sessionToken,omitempty"`
	ExpiredTime  int64  `protobuf:"varint,4,opt,name=expiredTime,proto3" json:"expiredTime,omitempty"`
	StartTime    int64  `protobuf:"varint,5,opt,name=startTime,proto3" json:"startTime,omitempty"`
}

func (x *GenCosStsResp) Reset() {
	*x = GenCosStsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenCosStsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenCosStsResp) ProtoMessage() {}

func (x *GenCosStsResp) ProtoReflect() protoreflect.Message {
	mi := &file_sts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenCosStsResp.ProtoReflect.Descriptor instead.
func (*GenCosStsResp) Descriptor() ([]byte, []int) {
	return file_sts_proto_rawDescGZIP(), []int{1}
}

func (x *GenCosStsResp) GetSecretId() string {
	if x != nil {
		return x.SecretId
	}
	return ""
}

func (x *GenCosStsResp) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *GenCosStsResp) GetSessionToken() string {
	if x != nil {
		return x.SessionToken
	}
	return ""
}

func (x *GenCosStsResp) GetExpiredTime() int64 {
	if x != nil {
		return x.ExpiredTime
	}
	return 0
}

func (x *GenCosStsResp) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

type GenSignedUrlReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecretId  string `protobuf:"bytes,1,opt,name=secretId,proto3" json:"secretId,omitempty"`
	SecretKey string `protobuf:"bytes,2,opt,name=secretKey,proto3" json:"secretKey,omitempty"`
	Method    string `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	Path      string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *GenSignedUrlReq) Reset() {
	*x = GenSignedUrlReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenSignedUrlReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenSignedUrlReq) ProtoMessage() {}

func (x *GenSignedUrlReq) ProtoReflect() protoreflect.Message {
	mi := &file_sts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenSignedUrlReq.ProtoReflect.Descriptor instead.
func (*GenSignedUrlReq) Descriptor() ([]byte, []int) {
	return file_sts_proto_rawDescGZIP(), []int{2}
}

func (x *GenSignedUrlReq) GetSecretId() string {
	if x != nil {
		return x.SecretId
	}
	return ""
}

func (x *GenSignedUrlReq) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *GenSignedUrlReq) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *GenSignedUrlReq) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type GenSignedUrlResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SignedUrl string `protobuf:"bytes,1,opt,name=signedUrl,proto3" json:"signedUrl,omitempty"`
}

func (x *GenSignedUrlResp) Reset() {
	*x = GenSignedUrlResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenSignedUrlResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenSignedUrlResp) ProtoMessage() {}

func (x *GenSignedUrlResp) ProtoReflect() protoreflect.Message {
	mi := &file_sts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenSignedUrlResp.ProtoReflect.Descriptor instead.
func (*GenSignedUrlResp) Descriptor() ([]byte, []int) {
	return file_sts_proto_rawDescGZIP(), []int{3}
}

func (x *GenSignedUrlResp) GetSignedUrl() string {
	if x != nil {
		return x.SignedUrl
	}
	return ""
}

type DeleteObjectReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *DeleteObjectReq) Reset() {
	*x = DeleteObjectReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteObjectReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteObjectReq) ProtoMessage() {}

func (x *DeleteObjectReq) ProtoReflect() protoreflect.Message {
	mi := &file_sts_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteObjectReq.ProtoReflect.Descriptor instead.
func (*DeleteObjectReq) Descriptor() ([]byte, []int) {
	return file_sts_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteObjectReq) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type DeleteObjectResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteObjectResp) Reset() {
	*x = DeleteObjectResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sts_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteObjectResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteObjectResp) ProtoMessage() {}

func (x *DeleteObjectResp) ProtoReflect() protoreflect.Message {
	mi := &file_sts_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteObjectResp.ProtoReflect.Descriptor instead.
func (*DeleteObjectResp) Descriptor() ([]byte, []int) {
	return file_sts_proto_rawDescGZIP(), []int{5}
}

// app值为meowchat,old,manager
type GetAccessTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App string `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
}

func (x *GetAccessTokenReq) Reset() {
	*x = GetAccessTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sts_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccessTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessTokenReq) ProtoMessage() {}

func (x *GetAccessTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_sts_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessTokenReq.ProtoReflect.Descriptor instead.
func (*GetAccessTokenReq) Descriptor() ([]byte, []int) {
	return file_sts_proto_rawDescGZIP(), []int{6}
}

func (x *GetAccessTokenReq) GetApp() string {
	if x != nil {
		return x.App
	}
	return ""
}

type GetAccessTokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	ExpiresIn   int64  `protobuf:"varint,2,opt,name=expiresIn,proto3" json:"expiresIn,omitempty"`
}

func (x *GetAccessTokenResp) Reset() {
	*x = GetAccessTokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sts_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccessTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessTokenResp) ProtoMessage() {}

func (x *GetAccessTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_sts_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessTokenResp.ProtoReflect.Descriptor instead.
func (*GetAccessTokenResp) Descriptor() ([]byte, []int) {
	return file_sts_proto_rawDescGZIP(), []int{7}
}

func (x *GetAccessTokenResp) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *GetAccessTokenResp) GetExpiresIn() int64 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

var File_sts_proto protoreflect.FileDescriptor

var file_sts_proto_rawDesc = []byte{
	0x0a, 0x09, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x73, 0x74, 0x73,
	0x22, 0x22, 0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x43, 0x6f, 0x73, 0x53, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x22, 0xad, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x6e, 0x43, 0x6f, 0x73, 0x53,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79,
	0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x77, 0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x30, 0x0a,
	0x10, 0x47, 0x65, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x22,
	0x25, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x12, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x25, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70,
	0x70, 0x22, 0x54, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x32, 0xfa, 0x01, 0x0a, 0x07, 0x73, 0x74, 0x73, 0x5f,
	0x72, 0x70, 0x63, 0x12, 0x32, 0x0a, 0x09, 0x67, 0x65, 0x6e, 0x43, 0x6f, 0x73, 0x53, 0x74, 0x73,
	0x12, 0x11, 0x2e, 0x73, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x6e, 0x43, 0x6f, 0x73, 0x53, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x73, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x6e, 0x43, 0x6f, 0x73,
	0x53, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3b, 0x0a, 0x0c, 0x67, 0x65, 0x6e, 0x53, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x2e, 0x73, 0x74, 0x73, 0x2e, 0x47, 0x65,
	0x6e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e,
	0x73, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x55, 0x72, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x3b, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x14, 0x2e, 0x73, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x73, 0x74, 0x73,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x41, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x16, 0x2e, 0x73, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x73, 0x74,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sts_proto_rawDescOnce sync.Once
	file_sts_proto_rawDescData = file_sts_proto_rawDesc
)

func file_sts_proto_rawDescGZIP() []byte {
	file_sts_proto_rawDescOnce.Do(func() {
		file_sts_proto_rawDescData = protoimpl.X.CompressGZIP(file_sts_proto_rawDescData)
	})
	return file_sts_proto_rawDescData
}

var file_sts_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_sts_proto_goTypes = []interface{}{
	(*GenCosStsReq)(nil),       // 0: sts.GenCosStsReq
	(*GenCosStsResp)(nil),      // 1: sts.GenCosStsResp
	(*GenSignedUrlReq)(nil),    // 2: sts.GenSignedUrlReq
	(*GenSignedUrlResp)(nil),   // 3: sts.GenSignedUrlResp
	(*DeleteObjectReq)(nil),    // 4: sts.DeleteObjectReq
	(*DeleteObjectResp)(nil),   // 5: sts.DeleteObjectResp
	(*GetAccessTokenReq)(nil),  // 6: sts.GetAccessTokenReq
	(*GetAccessTokenResp)(nil), // 7: sts.GetAccessTokenResp
}
var file_sts_proto_depIdxs = []int32{
	0, // 0: sts.sts_rpc.genCosSts:input_type -> sts.GenCosStsReq
	2, // 1: sts.sts_rpc.genSignedUrl:input_type -> sts.GenSignedUrlReq
	4, // 2: sts.sts_rpc.deleteObject:input_type -> sts.DeleteObjectReq
	6, // 3: sts.sts_rpc.GetAccessToken:input_type -> sts.GetAccessTokenReq
	1, // 4: sts.sts_rpc.genCosSts:output_type -> sts.GenCosStsResp
	3, // 5: sts.sts_rpc.genSignedUrl:output_type -> sts.GenSignedUrlResp
	5, // 6: sts.sts_rpc.deleteObject:output_type -> sts.DeleteObjectResp
	7, // 7: sts.sts_rpc.GetAccessToken:output_type -> sts.GetAccessTokenResp
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sts_proto_init() }
func file_sts_proto_init() {
	if File_sts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenCosStsReq); i {
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
		file_sts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenCosStsResp); i {
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
		file_sts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenSignedUrlReq); i {
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
		file_sts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenSignedUrlResp); i {
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
		file_sts_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteObjectReq); i {
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
		file_sts_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteObjectResp); i {
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
		file_sts_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccessTokenReq); i {
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
		file_sts_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccessTokenResp); i {
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
			RawDescriptor: file_sts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sts_proto_goTypes,
		DependencyIndexes: file_sts_proto_depIdxs,
		MessageInfos:      file_sts_proto_msgTypes,
	}.Build()
	File_sts_proto = out.File
	file_sts_proto_rawDesc = nil
	file_sts_proto_goTypes = nil
	file_sts_proto_depIdxs = nil
}
