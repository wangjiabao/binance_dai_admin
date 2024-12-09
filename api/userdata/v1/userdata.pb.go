// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.7
// source: api/userdata/v1/userdata.proto

package v1

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

type GetUserOrdersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUserOrdersRequest) Reset() {
	*x = GetUserOrdersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_userdata_v1_userdata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserOrdersRequest) ProtoMessage() {}

func (x *GetUserOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_userdata_v1_userdata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserOrdersRequest.ProtoReflect.Descriptor instead.
func (*GetUserOrdersRequest) Descriptor() ([]byte, []int) {
	return file_api_userdata_v1_userdata_proto_rawDescGZIP(), []int{0}
}

type GetUserOrdersReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId int64 `protobuf:"varint,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
}

func (x *GetUserOrdersReply) Reset() {
	*x = GetUserOrdersReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_userdata_v1_userdata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserOrdersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserOrdersReply) ProtoMessage() {}

func (x *GetUserOrdersReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_userdata_v1_userdata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserOrdersReply.ProtoReflect.Descriptor instead.
func (*GetUserOrdersReply) Descriptor() ([]byte, []int) {
	return file_api_userdata_v1_userdata_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserOrdersReply) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type PullUserIncomeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PullUserIncomeRequest) Reset() {
	*x = PullUserIncomeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_userdata_v1_userdata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullUserIncomeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullUserIncomeRequest) ProtoMessage() {}

func (x *PullUserIncomeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_userdata_v1_userdata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullUserIncomeRequest.ProtoReflect.Descriptor instead.
func (*PullUserIncomeRequest) Descriptor() ([]byte, []int) {
	return file_api_userdata_v1_userdata_proto_rawDescGZIP(), []int{2}
}

type PullUserIncomeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PullUserIncomeReply) Reset() {
	*x = PullUserIncomeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_userdata_v1_userdata_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullUserIncomeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullUserIncomeReply) ProtoMessage() {}

func (x *PullUserIncomeReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_userdata_v1_userdata_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullUserIncomeReply.ProtoReflect.Descriptor instead.
func (*PullUserIncomeReply) Descriptor() ([]byte, []int) {
	return file_api_userdata_v1_userdata_proto_rawDescGZIP(), []int{3}
}

type GetUsersIncomeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    uint64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Plat      string `protobuf:"bytes,2,opt,name=plat,proto3" json:"plat,omitempty"`
	ApiStatus uint64 `protobuf:"varint,3,opt,name=apiStatus,proto3" json:"apiStatus,omitempty"`
}

func (x *GetUsersIncomeRequest) Reset() {
	*x = GetUsersIncomeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_userdata_v1_userdata_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersIncomeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersIncomeRequest) ProtoMessage() {}

func (x *GetUsersIncomeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_userdata_v1_userdata_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersIncomeRequest.ProtoReflect.Descriptor instead.
func (*GetUsersIncomeRequest) Descriptor() ([]byte, []int) {
	return file_api_userdata_v1_userdata_proto_rawDescGZIP(), []int{4}
}

func (x *GetUsersIncomeRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetUsersIncomeRequest) GetPlat() string {
	if x != nil {
		return x.Plat
	}
	return ""
}

func (x *GetUsersIncomeRequest) GetApiStatus() uint64 {
	if x != nil {
		return x.ApiStatus
	}
	return 0
}

type GetUsersIncomeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*GetUsersIncomeReply_DataList `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *GetUsersIncomeReply) Reset() {
	*x = GetUsersIncomeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_userdata_v1_userdata_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersIncomeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersIncomeReply) ProtoMessage() {}

func (x *GetUsersIncomeReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_userdata_v1_userdata_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersIncomeReply.ProtoReflect.Descriptor instead.
func (*GetUsersIncomeReply) Descriptor() ([]byte, []int) {
	return file_api_userdata_v1_userdata_proto_rawDescGZIP(), []int{5}
}

func (x *GetUsersIncomeReply) GetList() []*GetUsersIncomeReply_DataList {
	if x != nil {
		return x.List
	}
	return nil
}

type GetUsersIncomeReply_DataList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Symbol string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Income string `protobuf:"bytes,3,opt,name=income,proto3" json:"income,omitempty"`
	Time   string `protobuf:"bytes,4,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *GetUsersIncomeReply_DataList) Reset() {
	*x = GetUsersIncomeReply_DataList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_userdata_v1_userdata_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersIncomeReply_DataList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersIncomeReply_DataList) ProtoMessage() {}

func (x *GetUsersIncomeReply_DataList) ProtoReflect() protoreflect.Message {
	mi := &file_api_userdata_v1_userdata_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersIncomeReply_DataList.ProtoReflect.Descriptor instead.
func (*GetUsersIncomeReply_DataList) Descriptor() ([]byte, []int) {
	return file_api_userdata_v1_userdata_proto_rawDescGZIP(), []int{5, 0}
}

func (x *GetUsersIncomeReply_DataList) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetUsersIncomeReply_DataList) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *GetUsersIncomeReply_DataList) GetIncome() string {
	if x != nil {
		return x.Income
	}
	return ""
}

func (x *GetUsersIncomeReply_DataList) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

var File_api_userdata_v1_userdata_proto protoreflect.FileDescriptor

var file_api_userdata_v1_userdata_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x50, 0x75, 0x6c, 0x6c, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x15, 0x0a, 0x13, 0x50, 0x75, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x63, 0x6f,
	0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x61, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6c, 0x61, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x70, 0x69, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x61, 0x70, 0x69, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xc0, 0x01, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x41, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x63, 0x6f, 0x6d,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x1a, 0x66, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x32, 0xbd, 0x03,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x12, 0x8b, 0x01, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x25, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28,
	0x12, 0x26, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x64,
	0x61, 0x69, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x90, 0x01, 0x0a, 0x0e, 0x50, 0x75, 0x6c,
	0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x26, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75,
	0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x63, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x2a, 0x12, 0x28, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f,
	0x64, 0x61, 0x69, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x73, 0x12, 0x8f, 0x01, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x26,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2f, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x29, 0x12, 0x27, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x6e, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x64, 0x61, 0x69, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x67, 0x65, 0x74,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x42, 0x39, 0x0a,
	0x0f, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31,
	0x50, 0x01, 0x5a, 0x24, 0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x64, 0x61, 0x69, 0x5f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x64, 0x61,
	0x74, 0x61, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_userdata_v1_userdata_proto_rawDescOnce sync.Once
	file_api_userdata_v1_userdata_proto_rawDescData = file_api_userdata_v1_userdata_proto_rawDesc
)

func file_api_userdata_v1_userdata_proto_rawDescGZIP() []byte {
	file_api_userdata_v1_userdata_proto_rawDescOnce.Do(func() {
		file_api_userdata_v1_userdata_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_userdata_v1_userdata_proto_rawDescData)
	})
	return file_api_userdata_v1_userdata_proto_rawDescData
}

var file_api_userdata_v1_userdata_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_userdata_v1_userdata_proto_goTypes = []interface{}{
	(*GetUserOrdersRequest)(nil),         // 0: api.userdata.v1.GetUserOrdersRequest
	(*GetUserOrdersReply)(nil),           // 1: api.userdata.v1.GetUserOrdersReply
	(*PullUserIncomeRequest)(nil),        // 2: api.userdata.v1.PullUserIncomeRequest
	(*PullUserIncomeReply)(nil),          // 3: api.userdata.v1.PullUserIncomeReply
	(*GetUsersIncomeRequest)(nil),        // 4: api.userdata.v1.GetUsersIncomeRequest
	(*GetUsersIncomeReply)(nil),          // 5: api.userdata.v1.GetUsersIncomeReply
	(*GetUsersIncomeReply_DataList)(nil), // 6: api.userdata.v1.GetUsersIncomeReply.DataList
}
var file_api_userdata_v1_userdata_proto_depIdxs = []int32{
	6, // 0: api.userdata.v1.GetUsersIncomeReply.list:type_name -> api.userdata.v1.GetUsersIncomeReply.DataList
	0, // 1: api.userdata.v1.Userdata.GetUserOrders:input_type -> api.userdata.v1.GetUserOrdersRequest
	2, // 2: api.userdata.v1.Userdata.PullUserIncome:input_type -> api.userdata.v1.PullUserIncomeRequest
	4, // 3: api.userdata.v1.Userdata.GetUsersIncome:input_type -> api.userdata.v1.GetUsersIncomeRequest
	1, // 4: api.userdata.v1.Userdata.GetUserOrders:output_type -> api.userdata.v1.GetUserOrdersReply
	3, // 5: api.userdata.v1.Userdata.PullUserIncome:output_type -> api.userdata.v1.PullUserIncomeReply
	5, // 6: api.userdata.v1.Userdata.GetUsersIncome:output_type -> api.userdata.v1.GetUsersIncomeReply
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_userdata_v1_userdata_proto_init() }
func file_api_userdata_v1_userdata_proto_init() {
	if File_api_userdata_v1_userdata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_userdata_v1_userdata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserOrdersRequest); i {
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
		file_api_userdata_v1_userdata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserOrdersReply); i {
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
		file_api_userdata_v1_userdata_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullUserIncomeRequest); i {
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
		file_api_userdata_v1_userdata_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullUserIncomeReply); i {
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
		file_api_userdata_v1_userdata_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersIncomeRequest); i {
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
		file_api_userdata_v1_userdata_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersIncomeReply); i {
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
		file_api_userdata_v1_userdata_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersIncomeReply_DataList); i {
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
			RawDescriptor: file_api_userdata_v1_userdata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_userdata_v1_userdata_proto_goTypes,
		DependencyIndexes: file_api_userdata_v1_userdata_proto_depIdxs,
		MessageInfos:      file_api_userdata_v1_userdata_proto_msgTypes,
	}.Build()
	File_api_userdata_v1_userdata_proto = out.File
	file_api_userdata_v1_userdata_proto_rawDesc = nil
	file_api_userdata_v1_userdata_proto_goTypes = nil
	file_api_userdata_v1_userdata_proto_depIdxs = nil
}
