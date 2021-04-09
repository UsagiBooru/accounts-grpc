// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: accounts.proto

package gen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Status int32

const (
	Status_SUCCESS Status = 0
	Status_FAILED  Status = 1
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "SUCCESS",
		1: "FAILED",
	}
	Status_value = map[string]int32{
		"SUCCESS": 0,
		"FAILED":  1,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_accounts_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_accounts_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{0}
}

type GetAccountNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID int32 `protobuf:"varint,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
}

func (x *GetAccountNameRequest) Reset() {
	*x = GetAccountNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountNameRequest) ProtoMessage() {}

func (x *GetAccountNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountNameRequest.ProtoReflect.Descriptor instead.
func (*GetAccountNameRequest) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{0}
}

func (x *GetAccountNameRequest) GetAccountID() int32 {
	if x != nil {
		return x.AccountID
	}
	return 0
}

type GetAccountNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code Status `protobuf:"varint,1,opt,name=code,proto3,enum=Status" json:"code,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetAccountNameResponse) Reset() {
	*x = GetAccountNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountNameResponse) ProtoMessage() {}

func (x *GetAccountNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountNameResponse.ProtoReflect.Descriptor instead.
func (*GetAccountNameResponse) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{1}
}

func (x *GetAccountNameResponse) GetCode() Status {
	if x != nil {
		return x.Code
	}
	return Status_SUCCESS
}

func (x *GetAccountNameResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AddUploadHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID int32 `protobuf:"varint,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
}

func (x *AddUploadHistoryRequest) Reset() {
	*x = AddUploadHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUploadHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUploadHistoryRequest) ProtoMessage() {}

func (x *AddUploadHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUploadHistoryRequest.ProtoReflect.Descriptor instead.
func (*AddUploadHistoryRequest) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{2}
}

func (x *AddUploadHistoryRequest) GetAccountID() int32 {
	if x != nil {
		return x.AccountID
	}
	return 0
}

type AddUploadHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      Status `protobuf:"varint,1,opt,name=code,proto3,enum=Status" json:"code,omitempty"`
	HistoryID int32  `protobuf:"varint,2,opt,name=historyID,proto3" json:"historyID,omitempty"`
}

func (x *AddUploadHistoryResponse) Reset() {
	*x = AddUploadHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUploadHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUploadHistoryResponse) ProtoMessage() {}

func (x *AddUploadHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUploadHistoryResponse.ProtoReflect.Descriptor instead.
func (*AddUploadHistoryResponse) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{3}
}

func (x *AddUploadHistoryResponse) GetCode() Status {
	if x != nil {
		return x.Code
	}
	return Status_SUCCESS
}

func (x *AddUploadHistoryResponse) GetHistoryID() int32 {
	if x != nil {
		return x.HistoryID
	}
	return 0
}

type EditUploadHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HistoryID     int32 `protobuf:"varint,1,opt,name=historyID,proto3" json:"historyID,omitempty"`
	HistoryStatus int32 `protobuf:"varint,2,opt,name=historyStatus,proto3" json:"historyStatus,omitempty"`
}

func (x *EditUploadHistoryRequest) Reset() {
	*x = EditUploadHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditUploadHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditUploadHistoryRequest) ProtoMessage() {}

func (x *EditUploadHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditUploadHistoryRequest.ProtoReflect.Descriptor instead.
func (*EditUploadHistoryRequest) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{4}
}

func (x *EditUploadHistoryRequest) GetHistoryID() int32 {
	if x != nil {
		return x.HistoryID
	}
	return 0
}

func (x *EditUploadHistoryRequest) GetHistoryStatus() int32 {
	if x != nil {
		return x.HistoryStatus
	}
	return 0
}

type EditUploadHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code Status `protobuf:"varint,1,opt,name=code,proto3,enum=Status" json:"code,omitempty"`
}

func (x *EditUploadHistoryResponse) Reset() {
	*x = EditUploadHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditUploadHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditUploadHistoryResponse) ProtoMessage() {}

func (x *EditUploadHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditUploadHistoryResponse.ProtoReflect.Descriptor instead.
func (*EditUploadHistoryResponse) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{5}
}

func (x *EditUploadHistoryResponse) GetCode() Status {
	if x != nil {
		return x.Code
	}
	return Status_SUCCESS
}

type DeleteUploadHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HistoryID int32 `protobuf:"varint,1,opt,name=historyID,proto3" json:"historyID,omitempty"`
}

func (x *DeleteUploadHistoryRequest) Reset() {
	*x = DeleteUploadHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUploadHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUploadHistoryRequest) ProtoMessage() {}

func (x *DeleteUploadHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUploadHistoryRequest.ProtoReflect.Descriptor instead.
func (*DeleteUploadHistoryRequest) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteUploadHistoryRequest) GetHistoryID() int32 {
	if x != nil {
		return x.HistoryID
	}
	return 0
}

type DeleteUploadHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code Status `protobuf:"varint,1,opt,name=code,proto3,enum=Status" json:"code,omitempty"`
}

func (x *DeleteUploadHistoryResponse) Reset() {
	*x = DeleteUploadHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUploadHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUploadHistoryResponse) ProtoMessage() {}

func (x *DeleteUploadHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUploadHistoryResponse.ProtoReflect.Descriptor instead.
func (*DeleteUploadHistoryResponse) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteUploadHistoryResponse) GetCode() Status {
	if x != nil {
		return x.Code
	}
	return Status_SUCCESS
}

var File_accounts_proto protoreflect.FileDescriptor

var file_accounts_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x35, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x49, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1b, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x37, 0x0a, 0x17, 0x41, 0x64, 0x64, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x55, 0x0a, 0x18, 0x41,
	0x64, 0x64, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x49, 0x44, 0x22, 0x5e, 0x0a, 0x18, 0x45, 0x64, 0x69, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d,
	0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0d, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x38, 0x0a, 0x19, 0x45, 0x64, 0x69, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1b, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x3a, 0x0a, 0x1a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x68,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x22, 0x3a, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x2a, 0x21, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b,
	0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x32, 0xbc, 0x02, 0x0a, 0x08, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x12, 0x43, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x10, 0x41, 0x64, 0x64,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x18, 0x2e,
	0x41, 0x64, 0x64, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x11, 0x45, 0x64, 0x69, 0x74, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x19, 0x2e, 0x45, 0x64, 0x69, 0x74,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x52, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1b, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_accounts_proto_rawDescOnce sync.Once
	file_accounts_proto_rawDescData = file_accounts_proto_rawDesc
)

func file_accounts_proto_rawDescGZIP() []byte {
	file_accounts_proto_rawDescOnce.Do(func() {
		file_accounts_proto_rawDescData = protoimpl.X.CompressGZIP(file_accounts_proto_rawDescData)
	})
	return file_accounts_proto_rawDescData
}

var file_accounts_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_accounts_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_accounts_proto_goTypes = []interface{}{
	(Status)(0),                         // 0: Status
	(*GetAccountNameRequest)(nil),       // 1: GetAccountNameRequest
	(*GetAccountNameResponse)(nil),      // 2: GetAccountNameResponse
	(*AddUploadHistoryRequest)(nil),     // 3: AddUploadHistoryRequest
	(*AddUploadHistoryResponse)(nil),    // 4: AddUploadHistoryResponse
	(*EditUploadHistoryRequest)(nil),    // 5: EditUploadHistoryRequest
	(*EditUploadHistoryResponse)(nil),   // 6: EditUploadHistoryResponse
	(*DeleteUploadHistoryRequest)(nil),  // 7: DeleteUploadHistoryRequest
	(*DeleteUploadHistoryResponse)(nil), // 8: DeleteUploadHistoryResponse
}
var file_accounts_proto_depIdxs = []int32{
	0, // 0: GetAccountNameResponse.code:type_name -> Status
	0, // 1: AddUploadHistoryResponse.code:type_name -> Status
	0, // 2: EditUploadHistoryResponse.code:type_name -> Status
	0, // 3: DeleteUploadHistoryResponse.code:type_name -> Status
	1, // 4: Accounts.GetAccountName:input_type -> GetAccountNameRequest
	3, // 5: Accounts.AddUploadHistory:input_type -> AddUploadHistoryRequest
	5, // 6: Accounts.EditUploadHistory:input_type -> EditUploadHistoryRequest
	7, // 7: Accounts.DeleteUploadHistory:input_type -> DeleteUploadHistoryRequest
	2, // 8: Accounts.GetAccountName:output_type -> GetAccountNameResponse
	4, // 9: Accounts.AddUploadHistory:output_type -> AddUploadHistoryResponse
	6, // 10: Accounts.EditUploadHistory:output_type -> EditUploadHistoryResponse
	8, // 11: Accounts.DeleteUploadHistory:output_type -> DeleteUploadHistoryResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_accounts_proto_init() }
func file_accounts_proto_init() {
	if File_accounts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_accounts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountNameRequest); i {
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
		file_accounts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountNameResponse); i {
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
		file_accounts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUploadHistoryRequest); i {
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
		file_accounts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUploadHistoryResponse); i {
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
		file_accounts_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditUploadHistoryRequest); i {
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
		file_accounts_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditUploadHistoryResponse); i {
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
		file_accounts_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUploadHistoryRequest); i {
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
		file_accounts_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUploadHistoryResponse); i {
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
			RawDescriptor: file_accounts_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_accounts_proto_goTypes,
		DependencyIndexes: file_accounts_proto_depIdxs,
		EnumInfos:         file_accounts_proto_enumTypes,
		MessageInfos:      file_accounts_proto_msgTypes,
	}.Build()
	File_accounts_proto = out.File
	file_accounts_proto_rawDesc = nil
	file_accounts_proto_goTypes = nil
	file_accounts_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AccountsClient is the client API for Accounts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountsClient interface {
	GetAccountName(ctx context.Context, in *GetAccountNameRequest, opts ...grpc.CallOption) (*GetAccountNameResponse, error)
	AddUploadHistory(ctx context.Context, in *AddUploadHistoryRequest, opts ...grpc.CallOption) (*AddUploadHistoryResponse, error)
	EditUploadHistory(ctx context.Context, in *EditUploadHistoryRequest, opts ...grpc.CallOption) (*EditUploadHistoryResponse, error)
	DeleteUploadHistory(ctx context.Context, in *DeleteUploadHistoryRequest, opts ...grpc.CallOption) (*DeleteUploadHistoryResponse, error)
}

type accountsClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountsClient(cc grpc.ClientConnInterface) AccountsClient {
	return &accountsClient{cc}
}

func (c *accountsClient) GetAccountName(ctx context.Context, in *GetAccountNameRequest, opts ...grpc.CallOption) (*GetAccountNameResponse, error) {
	out := new(GetAccountNameResponse)
	err := c.cc.Invoke(ctx, "/Accounts/GetAccountName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) AddUploadHistory(ctx context.Context, in *AddUploadHistoryRequest, opts ...grpc.CallOption) (*AddUploadHistoryResponse, error) {
	out := new(AddUploadHistoryResponse)
	err := c.cc.Invoke(ctx, "/Accounts/AddUploadHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) EditUploadHistory(ctx context.Context, in *EditUploadHistoryRequest, opts ...grpc.CallOption) (*EditUploadHistoryResponse, error) {
	out := new(EditUploadHistoryResponse)
	err := c.cc.Invoke(ctx, "/Accounts/EditUploadHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) DeleteUploadHistory(ctx context.Context, in *DeleteUploadHistoryRequest, opts ...grpc.CallOption) (*DeleteUploadHistoryResponse, error) {
	out := new(DeleteUploadHistoryResponse)
	err := c.cc.Invoke(ctx, "/Accounts/DeleteUploadHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountsServer is the server API for Accounts service.
type AccountsServer interface {
	GetAccountName(context.Context, *GetAccountNameRequest) (*GetAccountNameResponse, error)
	AddUploadHistory(context.Context, *AddUploadHistoryRequest) (*AddUploadHistoryResponse, error)
	EditUploadHistory(context.Context, *EditUploadHistoryRequest) (*EditUploadHistoryResponse, error)
	DeleteUploadHistory(context.Context, *DeleteUploadHistoryRequest) (*DeleteUploadHistoryResponse, error)
}

// UnimplementedAccountsServer can be embedded to have forward compatible implementations.
type UnimplementedAccountsServer struct {
}

func (*UnimplementedAccountsServer) GetAccountName(context.Context, *GetAccountNameRequest) (*GetAccountNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountName not implemented")
}
func (*UnimplementedAccountsServer) AddUploadHistory(context.Context, *AddUploadHistoryRequest) (*AddUploadHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUploadHistory not implemented")
}
func (*UnimplementedAccountsServer) EditUploadHistory(context.Context, *EditUploadHistoryRequest) (*EditUploadHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditUploadHistory not implemented")
}
func (*UnimplementedAccountsServer) DeleteUploadHistory(context.Context, *DeleteUploadHistoryRequest) (*DeleteUploadHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUploadHistory not implemented")
}

func RegisterAccountsServer(s *grpc.Server, srv AccountsServer) {
	s.RegisterService(&_Accounts_serviceDesc, srv)
}

func _Accounts_GetAccountName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).GetAccountName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Accounts/GetAccountName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).GetAccountName(ctx, req.(*GetAccountNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_AddUploadHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUploadHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).AddUploadHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Accounts/AddUploadHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).AddUploadHistory(ctx, req.(*AddUploadHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_EditUploadHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditUploadHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).EditUploadHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Accounts/EditUploadHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).EditUploadHistory(ctx, req.(*EditUploadHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_DeleteUploadHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUploadHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).DeleteUploadHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Accounts/DeleteUploadHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).DeleteUploadHistory(ctx, req.(*DeleteUploadHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Accounts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Accounts",
	HandlerType: (*AccountsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccountName",
			Handler:    _Accounts_GetAccountName_Handler,
		},
		{
			MethodName: "AddUploadHistory",
			Handler:    _Accounts_AddUploadHistory_Handler,
		},
		{
			MethodName: "EditUploadHistory",
			Handler:    _Accounts_EditUploadHistory_Handler,
		},
		{
			MethodName: "DeleteUploadHistory",
			Handler:    _Accounts_DeleteUploadHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accounts.proto",
}
