// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.24.3
// source: follow.proto

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

type FollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`                 // 关注者
	FollowedUserId int64 `protobuf:"varint,2,opt,name=followedUserId,proto3" json:"followedUserId,omitempty"` // 被关注者
}

func (x *FollowRequest) Reset() {
	*x = FollowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowRequest) ProtoMessage() {}

func (x *FollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_follow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowRequest.ProtoReflect.Descriptor instead.
func (*FollowRequest) Descriptor() ([]byte, []int) {
	return file_follow_proto_rawDescGZIP(), []int{0}
}

func (x *FollowRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FollowRequest) GetFollowedUserId() int64 {
	if x != nil {
		return x.FollowedUserId
	}
	return 0
}

type FollowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FollowResponse) Reset() {
	*x = FollowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowResponse) ProtoMessage() {}

func (x *FollowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_follow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowResponse.ProtoReflect.Descriptor instead.
func (*FollowResponse) Descriptor() ([]byte, []int) {
	return file_follow_proto_rawDescGZIP(), []int{1}
}

type UnFollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	FollowedUserId int64 `protobuf:"varint,2,opt,name=followedUserId,proto3" json:"followedUserId,omitempty"`
}

func (x *UnFollowRequest) Reset() {
	*x = UnFollowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnFollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnFollowRequest) ProtoMessage() {}

func (x *UnFollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_follow_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnFollowRequest.ProtoReflect.Descriptor instead.
func (*UnFollowRequest) Descriptor() ([]byte, []int) {
	return file_follow_proto_rawDescGZIP(), []int{2}
}

func (x *UnFollowRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UnFollowRequest) GetFollowedUserId() int64 {
	if x != nil {
		return x.FollowedUserId
	}
	return 0
}

type UnFollowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnFollowResponse) Reset() {
	*x = UnFollowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnFollowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnFollowResponse) ProtoMessage() {}

func (x *UnFollowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_follow_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnFollowResponse.ProtoReflect.Descriptor instead.
func (*UnFollowResponse) Descriptor() ([]byte, []int) {
	return file_follow_proto_rawDescGZIP(), []int{3}
}

type FollowListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	UserId   int64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Cursor   int64 `protobuf:"varint,3,opt,name=cursor,proto3" json:"cursor,omitempty"`
	PageSize int64 `protobuf:"varint,4,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *FollowListRequest) Reset() {
	*x = FollowListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowListRequest) ProtoMessage() {}

func (x *FollowListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_follow_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowListRequest.ProtoReflect.Descriptor instead.
func (*FollowListRequest) Descriptor() ([]byte, []int) {
	return file_follow_proto_rawDescGZIP(), []int{4}
}

func (x *FollowListRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FollowListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FollowListRequest) GetCursor() int64 {
	if x != nil {
		return x.Cursor
	}
	return 0
}

func (x *FollowListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type FollowItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	FollowedUserId int64 `protobuf:"varint,2,opt,name=followedUserId,proto3" json:"followedUserId,omitempty"` // 被关注者
	FansCount      int64 `protobuf:"varint,3,opt,name=fansCount,proto3" json:"fansCount,omitempty"`           // 粉丝数
	CreateTime     int64 `protobuf:"varint,4,opt,name=createTime,proto3" json:"createTime,omitempty"`         // 关注时间
}

func (x *FollowItem) Reset() {
	*x = FollowItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowItem) ProtoMessage() {}

func (x *FollowItem) ProtoReflect() protoreflect.Message {
	mi := &file_follow_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowItem.ProtoReflect.Descriptor instead.
func (*FollowItem) Descriptor() ([]byte, []int) {
	return file_follow_proto_rawDescGZIP(), []int{5}
}

func (x *FollowItem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FollowItem) GetFollowedUserId() int64 {
	if x != nil {
		return x.FollowedUserId
	}
	return 0
}

func (x *FollowItem) GetFansCount() int64 {
	if x != nil {
		return x.FansCount
	}
	return 0
}

func (x *FollowItem) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

type FollowListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items  []*FollowItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Cursor int64         `protobuf:"varint,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
	IsEnd  bool          `protobuf:"varint,3,opt,name=isEnd,proto3" json:"isEnd,omitempty"`
	Id     int64         `protobuf:"varint,4,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *FollowListResponse) Reset() {
	*x = FollowListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowListResponse) ProtoMessage() {}

func (x *FollowListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_follow_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowListResponse.ProtoReflect.Descriptor instead.
func (*FollowListResponse) Descriptor() ([]byte, []int) {
	return file_follow_proto_rawDescGZIP(), []int{6}
}

func (x *FollowListResponse) GetItems() []*FollowItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *FollowListResponse) GetCursor() int64 {
	if x != nil {
		return x.Cursor
	}
	return 0
}

func (x *FollowListResponse) GetIsEnd() bool {
	if x != nil {
		return x.IsEnd
	}
	return false
}

func (x *FollowListResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FansListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Cursor   int64 `protobuf:"varint,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
	PageSize int64 `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *FansListRequest) Reset() {
	*x = FansListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FansListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FansListRequest) ProtoMessage() {}

func (x *FansListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_follow_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FansListRequest.ProtoReflect.Descriptor instead.
func (*FansListRequest) Descriptor() ([]byte, []int) {
	return file_follow_proto_rawDescGZIP(), []int{7}
}

func (x *FansListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FansListRequest) GetCursor() int64 {
	if x != nil {
		return x.Cursor
	}
	return 0
}

func (x *FansListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type FansItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	FansUserId  int64 `protobuf:"varint,2,opt,name=fansUserId,proto3" json:"fansUserId,omitempty"`
	FollowCount int64 `protobuf:"varint,3,opt,name=followCount,proto3" json:"followCount,omitempty"`
	FansCount   int64 `protobuf:"varint,4,opt,name=fansCount,proto3" json:"fansCount,omitempty"`
	CreateTime  int64 `protobuf:"varint,5,opt,name=createTime,proto3" json:"createTime,omitempty"`
}

func (x *FansItem) Reset() {
	*x = FansItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FansItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FansItem) ProtoMessage() {}

func (x *FansItem) ProtoReflect() protoreflect.Message {
	mi := &file_follow_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FansItem.ProtoReflect.Descriptor instead.
func (*FansItem) Descriptor() ([]byte, []int) {
	return file_follow_proto_rawDescGZIP(), []int{8}
}

func (x *FansItem) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FansItem) GetFansUserId() int64 {
	if x != nil {
		return x.FansUserId
	}
	return 0
}

func (x *FansItem) GetFollowCount() int64 {
	if x != nil {
		return x.FollowCount
	}
	return 0
}

func (x *FansItem) GetFansCount() int64 {
	if x != nil {
		return x.FansCount
	}
	return 0
}

func (x *FansItem) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

type FansListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items  []*FansItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Cursor int64       `protobuf:"varint,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
	IsEnd  bool        `protobuf:"varint,3,opt,name=isEnd,proto3" json:"isEnd,omitempty"`
	Id     int64       `protobuf:"varint,4,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *FansListResponse) Reset() {
	*x = FansListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FansListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FansListResponse) ProtoMessage() {}

func (x *FansListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_follow_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FansListResponse.ProtoReflect.Descriptor instead.
func (*FansListResponse) Descriptor() ([]byte, []int) {
	return file_follow_proto_rawDescGZIP(), []int{9}
}

func (x *FansListResponse) GetItems() []*FansItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *FansListResponse) GetCursor() int64 {
	if x != nil {
		return x.Cursor
	}
	return 0
}

func (x *FansListResponse) GetIsEnd() bool {
	if x != nil {
		return x.IsEnd
	}
	return false
}

func (x *FansListResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_follow_proto protoreflect.FileDescriptor

var file_follow_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x4f, 0x0a, 0x0d, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x26, 0x0a, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x51, 0x0a, 0x0f, 0x55, 0x6e,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x12, 0x0a,
	0x10, 0x55, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x6f, 0x0a, 0x11, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x22, 0x82, 0x01, 0x0a, 0x0a, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49,
	0x64, 0x12, 0x26, 0x0a, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x61, 0x6e,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x66, 0x61,
	0x6e, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x7d, 0x0a, 0x12, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x45, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x69, 0x73, 0x45, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x22, 0x5d, 0x0a, 0x0f, 0x46, 0x61, 0x6e, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xa2, 0x01, 0x0a, 0x08, 0x46, 0x61, 0x6e, 0x73, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x61,
	0x6e, 0x73, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x66, 0x61, 0x6e, 0x73, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x66, 0x61, 0x6e, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x66, 0x61, 0x6e, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x79, 0x0a, 0x10, 0x46, 0x61,
	0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x61, 0x6e, 0x73, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x73, 0x45, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x69, 0x73, 0x45, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x49, 0x64, 0x32, 0x8c, 0x02, 0x0a, 0x06, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x12, 0x39, 0x0a, 0x06, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x55,
	0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x55, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x6e, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0a,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x46, 0x61, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x61, 0x6e, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x46, 0x61, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_follow_proto_rawDescOnce sync.Once
	file_follow_proto_rawDescData = file_follow_proto_rawDesc
)

func file_follow_proto_rawDescGZIP() []byte {
	file_follow_proto_rawDescOnce.Do(func() {
		file_follow_proto_rawDescData = protoimpl.X.CompressGZIP(file_follow_proto_rawDescData)
	})
	return file_follow_proto_rawDescData
}

var file_follow_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_follow_proto_goTypes = []interface{}{
	(*FollowRequest)(nil),      // 0: service.FollowRequest
	(*FollowResponse)(nil),     // 1: service.FollowResponse
	(*UnFollowRequest)(nil),    // 2: service.UnFollowRequest
	(*UnFollowResponse)(nil),   // 3: service.UnFollowResponse
	(*FollowListRequest)(nil),  // 4: service.FollowListRequest
	(*FollowItem)(nil),         // 5: service.FollowItem
	(*FollowListResponse)(nil), // 6: service.FollowListResponse
	(*FansListRequest)(nil),    // 7: service.FansListRequest
	(*FansItem)(nil),           // 8: service.FansItem
	(*FansListResponse)(nil),   // 9: service.FansListResponse
}
var file_follow_proto_depIdxs = []int32{
	5, // 0: service.FollowListResponse.items:type_name -> service.FollowItem
	8, // 1: service.FansListResponse.items:type_name -> service.FansItem
	0, // 2: service.Follow.Follow:input_type -> service.FollowRequest
	2, // 3: service.Follow.UnFollow:input_type -> service.UnFollowRequest
	4, // 4: service.Follow.FollowList:input_type -> service.FollowListRequest
	7, // 5: service.Follow.FansList:input_type -> service.FansListRequest
	1, // 6: service.Follow.Follow:output_type -> service.FollowResponse
	3, // 7: service.Follow.UnFollow:output_type -> service.UnFollowResponse
	6, // 8: service.Follow.FollowList:output_type -> service.FollowListResponse
	9, // 9: service.Follow.FansList:output_type -> service.FansListResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_follow_proto_init() }
func file_follow_proto_init() {
	if File_follow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_follow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowRequest); i {
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
		file_follow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowResponse); i {
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
		file_follow_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnFollowRequest); i {
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
		file_follow_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnFollowResponse); i {
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
		file_follow_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowListRequest); i {
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
		file_follow_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowItem); i {
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
		file_follow_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowListResponse); i {
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
		file_follow_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FansListRequest); i {
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
		file_follow_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FansItem); i {
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
		file_follow_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FansListResponse); i {
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
			RawDescriptor: file_follow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_follow_proto_goTypes,
		DependencyIndexes: file_follow_proto_depIdxs,
		MessageInfos:      file_follow_proto_msgTypes,
	}.Build()
	File_follow_proto = out.File
	file_follow_proto_rawDesc = nil
	file_follow_proto_goTypes = nil
	file_follow_proto_depIdxs = nil
}
