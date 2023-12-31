// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: api/chat.proto

package api

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

// 채팅 메시지를 정의하는 메시지 타입
type ChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"` // 채널 이름
	User    string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`       // 사용자 이름
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"` // 메시지 내용
}

func (x *ChatMessage) Reset() {
	*x = ChatMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessage) ProtoMessage() {}

func (x *ChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessage.ProtoReflect.Descriptor instead.
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return file_api_chat_proto_rawDescGZIP(), []int{0}
}

func (x *ChatMessage) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *ChatMessage) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *ChatMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ChannelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Action string `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"` // 채널 "create" 또는 "delete"
}

func (x *ChannelRequest) Reset() {
	*x = ChannelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelRequest) ProtoMessage() {}

func (x *ChannelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelRequest.ProtoReflect.Descriptor instead.
func (*ChannelRequest) Descriptor() ([]byte, []int) {
	return file_api_chat_proto_rawDescGZIP(), []int{1}
}

func (x *ChannelRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChannelRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type ChannelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`        // 채널 이름
	Success bool   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"` // 요청 처리 성공 여부
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`  // 추가 메시지 또는 오류 메시지
}

func (x *ChannelResponse) Reset() {
	*x = ChannelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelResponse) ProtoMessage() {}

func (x *ChannelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelResponse.ProtoReflect.Descriptor instead.
func (*ChannelResponse) Descriptor() ([]byte, []int) {
	return file_api_chat_proto_rawDescGZIP(), []int{2}
}

func (x *ChannelResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChannelResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ChannelResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ChannelListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChannelListRequest) Reset() {
	*x = ChannelListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelListRequest) ProtoMessage() {}

func (x *ChannelListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelListRequest.ProtoReflect.Descriptor instead.
func (*ChannelListRequest) Descriptor() ([]byte, []int) {
	return file_api_chat_proto_rawDescGZIP(), []int{3}
}

type ChannelListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channels []string `protobuf:"bytes,1,rep,name=channels,proto3" json:"channels,omitempty"` // 채널이름 목록 - list 는 여러 string 을 가지기때문에 앞에 repeated 를 추가함
}

func (x *ChannelListResponse) Reset() {
	*x = ChannelListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelListResponse) ProtoMessage() {}

func (x *ChannelListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelListResponse.ProtoReflect.Descriptor instead.
func (*ChannelListResponse) Descriptor() ([]byte, []int) {
	return file_api_chat_proto_rawDescGZIP(), []int{4}
}

func (x *ChannelListResponse) GetChannels() []string {
	if x != nil {
		return x.Channels
	}
	return nil
}

// 채널의 사용자 목록을 요청하기 위한 메시지
type ChannelUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"` // 채널 이름
}

func (x *ChannelUsersRequest) Reset() {
	*x = ChannelUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelUsersRequest) ProtoMessage() {}

func (x *ChannelUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelUsersRequest.ProtoReflect.Descriptor instead.
func (*ChannelUsersRequest) Descriptor() ([]byte, []int) {
	return file_api_chat_proto_rawDescGZIP(), []int{5}
}

func (x *ChannelUsersRequest) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

// 채널의 사용자 목록에 대한 응답 메시지
type ChannelUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []string `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"` // 사용자 이름 목록
}

func (x *ChannelUsersResponse) Reset() {
	*x = ChannelUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelUsersResponse) ProtoMessage() {}

func (x *ChannelUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelUsersResponse.ProtoReflect.Descriptor instead.
func (*ChannelUsersResponse) Descriptor() ([]byte, []int) {
	return file_api_chat_proto_rawDescGZIP(), []int{6}
}

func (x *ChannelUsersResponse) GetUsers() []string {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_api_chat_proto protoreflect.FileDescriptor

var file_api_chat_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x63, 0x68, 0x61, 0x74, 0x22, 0x55, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3c, 0x0a,
	0x0e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x59, 0x0a, 0x0f, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x31, 0x0a, 0x13,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x22,
	0x2f, 0x0a, 0x13, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x22, 0x2c, 0x0a, 0x14, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x32, 0x9c,
	0x02, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e,
	0x0a, 0x0d, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12,
	0x14, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45,
	0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x12, 0x18,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x11, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x11, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43,
	0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01,
	0x12, 0x4b, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x12, 0x19, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x14, 0x5a,
	0x12, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2f,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_chat_proto_rawDescOnce sync.Once
	file_api_chat_proto_rawDescData = file_api_chat_proto_rawDesc
)

func file_api_chat_proto_rawDescGZIP() []byte {
	file_api_chat_proto_rawDescOnce.Do(func() {
		file_api_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_chat_proto_rawDescData)
	})
	return file_api_chat_proto_rawDescData
}

var file_api_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_chat_proto_goTypes = []interface{}{
	(*ChatMessage)(nil),          // 0: chat.ChatMessage
	(*ChannelRequest)(nil),       // 1: chat.ChannelRequest
	(*ChannelResponse)(nil),      // 2: chat.ChannelResponse
	(*ChannelListRequest)(nil),   // 3: chat.ChannelListRequest
	(*ChannelListResponse)(nil),  // 4: chat.ChannelListResponse
	(*ChannelUsersRequest)(nil),  // 5: chat.ChannelUsersRequest
	(*ChannelUsersResponse)(nil), // 6: chat.ChannelUsersResponse
}
var file_api_chat_proto_depIdxs = []int32{
	1, // 0: chat.ChatService.ManageChannel:input_type -> chat.ChannelRequest
	3, // 1: chat.ChatService.ListChannels:input_type -> chat.ChannelListRequest
	0, // 2: chat.ChatService.SendMessage:input_type -> chat.ChatMessage
	5, // 3: chat.ChatService.ListChannelUsers:input_type -> chat.ChannelUsersRequest
	2, // 4: chat.ChatService.ManageChannel:output_type -> chat.ChannelResponse
	4, // 5: chat.ChatService.ListChannels:output_type -> chat.ChannelListResponse
	0, // 6: chat.ChatService.SendMessage:output_type -> chat.ChatMessage
	6, // 7: chat.ChatService.ListChannelUsers:output_type -> chat.ChannelUsersResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_chat_proto_init() }
func file_api_chat_proto_init() {
	if File_api_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatMessage); i {
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
		file_api_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelRequest); i {
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
		file_api_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelResponse); i {
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
		file_api_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelListRequest); i {
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
		file_api_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelListResponse); i {
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
		file_api_chat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelUsersRequest); i {
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
		file_api_chat_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelUsersResponse); i {
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
			RawDescriptor: file_api_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_chat_proto_goTypes,
		DependencyIndexes: file_api_chat_proto_depIdxs,
		MessageInfos:      file_api_chat_proto_msgTypes,
	}.Build()
	File_api_chat_proto = out.File
	file_api_chat_proto_rawDesc = nil
	file_api_chat_proto_goTypes = nil
	file_api_chat_proto_depIdxs = nil
}
