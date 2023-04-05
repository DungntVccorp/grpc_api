// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: echo.api.proto

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

type ECHO_TYPE_ID int32

const (
	ECHO_TYPE_ID__UN_USE         ECHO_TYPE_ID = 0
	ECHO_TYPE_ID_PING            ECHO_TYPE_ID = 1
	ECHO_TYPE_ID_SEND_MESSAGE    ECHO_TYPE_ID = 2
	ECHO_TYPE_ID_RECEIVE_MESSAGE ECHO_TYPE_ID = 3
)

// Enum value maps for ECHO_TYPE_ID.
var (
	ECHO_TYPE_ID_name = map[int32]string{
		0: "_UN_USE",
		1: "PING",
		2: "SEND_MESSAGE",
		3: "RECEIVE_MESSAGE",
	}
	ECHO_TYPE_ID_value = map[string]int32{
		"_UN_USE":         0,
		"PING":            1,
		"SEND_MESSAGE":    2,
		"RECEIVE_MESSAGE": 3,
	}
)

func (x ECHO_TYPE_ID) Enum() *ECHO_TYPE_ID {
	p := new(ECHO_TYPE_ID)
	*p = x
	return p
}

func (x ECHO_TYPE_ID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ECHO_TYPE_ID) Descriptor() protoreflect.EnumDescriptor {
	return file_echo_api_proto_enumTypes[0].Descriptor()
}

func (ECHO_TYPE_ID) Type() protoreflect.EnumType {
	return &file_echo_api_proto_enumTypes[0]
}

func (x ECHO_TYPE_ID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ECHO_TYPE_ID.Descriptor instead.
func (ECHO_TYPE_ID) EnumDescriptor() ([]byte, []int) {
	return file_echo_api_proto_rawDescGZIP(), []int{0}
}

type SendMessage_ERROR int32

const (
	SendMessage_NO_ERROR    SendMessage_ERROR = 0
	SendMessage_MSG_INVALID SendMessage_ERROR = 1
)

// Enum value maps for SendMessage_ERROR.
var (
	SendMessage_ERROR_name = map[int32]string{
		0: "NO_ERROR",
		1: "MSG_INVALID",
	}
	SendMessage_ERROR_value = map[string]int32{
		"NO_ERROR":    0,
		"MSG_INVALID": 1,
	}
)

func (x SendMessage_ERROR) Enum() *SendMessage_ERROR {
	p := new(SendMessage_ERROR)
	*p = x
	return p
}

func (x SendMessage_ERROR) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SendMessage_ERROR) Descriptor() protoreflect.EnumDescriptor {
	return file_echo_api_proto_enumTypes[1].Descriptor()
}

func (SendMessage_ERROR) Type() protoreflect.EnumType {
	return &file_echo_api_proto_enumTypes[1]
}

func (x SendMessage_ERROR) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SendMessage_ERROR.Descriptor instead.
func (SendMessage_ERROR) EnumDescriptor() ([]byte, []int) {
	return file_echo_api_proto_rawDescGZIP(), []int{2, 0}
}

type EchoPing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EchoPing) Reset() {
	*x = EchoPing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoPing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoPing) ProtoMessage() {}

func (x *EchoPing) ProtoReflect() protoreflect.Message {
	mi := &file_echo_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoPing.ProtoReflect.Descriptor instead.
func (*EchoPing) Descriptor() ([]byte, []int) {
	return file_echo_api_proto_rawDescGZIP(), []int{0}
}

type AppMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg      string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	UserId   string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *AppMessage) Reset() {
	*x = AppMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppMessage) ProtoMessage() {}

func (x *AppMessage) ProtoReflect() protoreflect.Message {
	mi := &file_echo_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppMessage.ProtoReflect.Descriptor instead.
func (*AppMessage) Descriptor() ([]byte, []int) {
	return file_echo_api_proto_rawDescGZIP(), []int{1}
}

func (x *AppMessage) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *AppMessage) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *AppMessage) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type SendMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendMessage) Reset() {
	*x = SendMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessage) ProtoMessage() {}

func (x *SendMessage) ProtoReflect() protoreflect.Message {
	mi := &file_echo_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessage.ProtoReflect.Descriptor instead.
func (*SendMessage) Descriptor() ([]byte, []int) {
	return file_echo_api_proto_rawDescGZIP(), []int{2}
}

type ReceiveMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg *AppMessage `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ReceiveMessage) Reset() {
	*x = ReceiveMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveMessage) ProtoMessage() {}

func (x *ReceiveMessage) ProtoReflect() protoreflect.Message {
	mi := &file_echo_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveMessage.ProtoReflect.Descriptor instead.
func (*ReceiveMessage) Descriptor() ([]byte, []int) {
	return file_echo_api_proto_rawDescGZIP(), []int{3}
}

func (x *ReceiveMessage) GetMsg() *AppMessage {
	if x != nil {
		return x.Msg
	}
	return nil
}

type EchoPing_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EchoPing_Request) Reset() {
	*x = EchoPing_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoPing_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoPing_Request) ProtoMessage() {}

func (x *EchoPing_Request) ProtoReflect() protoreflect.Message {
	mi := &file_echo_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoPing_Request.ProtoReflect.Descriptor instead.
func (*EchoPing_Request) Descriptor() ([]byte, []int) {
	return file_echo_api_proto_rawDescGZIP(), []int{0, 0}
}

type EchoPing_Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ver        string `protobuf:"bytes,1,opt,name=ver,proto3" json:"ver,omitempty"`
	ServerTime string `protobuf:"bytes,2,opt,name=server_time,json=serverTime,proto3" json:"server_time,omitempty"`
}

func (x *EchoPing_Reply) Reset() {
	*x = EchoPing_Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoPing_Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoPing_Reply) ProtoMessage() {}

func (x *EchoPing_Reply) ProtoReflect() protoreflect.Message {
	mi := &file_echo_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoPing_Reply.ProtoReflect.Descriptor instead.
func (*EchoPing_Reply) Descriptor() ([]byte, []int) {
	return file_echo_api_proto_rawDescGZIP(), []int{0, 1}
}

func (x *EchoPing_Reply) GetVer() string {
	if x != nil {
		return x.Ver
	}
	return ""
}

func (x *EchoPing_Reply) GetServerTime() string {
	if x != nil {
		return x.ServerTime
	}
	return ""
}

type SendMessage_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg *AppMessage `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SendMessage_Request) Reset() {
	*x = SendMessage_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessage_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessage_Request) ProtoMessage() {}

func (x *SendMessage_Request) ProtoReflect() protoreflect.Message {
	mi := &file_echo_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessage_Request.ProtoReflect.Descriptor instead.
func (*SendMessage_Request) Descriptor() ([]byte, []int) {
	return file_echo_api_proto_rawDescGZIP(), []int{2, 0}
}

func (x *SendMessage_Request) GetMsg() *AppMessage {
	if x != nil {
		return x.Msg
	}
	return nil
}

type SendMessage_Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg *AppMessage `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SendMessage_Reply) Reset() {
	*x = SendMessage_Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessage_Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessage_Reply) ProtoMessage() {}

func (x *SendMessage_Reply) ProtoReflect() protoreflect.Message {
	mi := &file_echo_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessage_Reply.ProtoReflect.Descriptor instead.
func (*SendMessage_Reply) Descriptor() ([]byte, []int) {
	return file_echo_api_proto_rawDescGZIP(), []int{2, 1}
}

func (x *SendMessage_Reply) GetMsg() *AppMessage {
	if x != nil {
		return x.Msg
	}
	return nil
}

var File_echo_api_proto protoreflect.FileDescriptor

var file_echo_api_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x51, 0x0a, 0x08, 0x45, 0x63, 0x68, 0x6f, 0x50, 0x69, 0x6e, 0x67, 0x1a, 0x09, 0x0a, 0x07,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x76,
	0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54,
	0x69, 0x6d, 0x65, 0x22, 0x54, 0x0a, 0x0a, 0x41, 0x70, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x87, 0x01, 0x0a, 0x0b, 0x53, 0x65,
	0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x28, 0x0a, 0x07, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x41, 0x70, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x1a, 0x26, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1d, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x41, 0x70, 0x70, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x26, 0x0a, 0x05, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x53, 0x47, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x10, 0x01, 0x22, 0x2f, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x41, 0x70, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x2a, 0x4c, 0x0a, 0x0c, 0x45, 0x43, 0x48, 0x4f, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x49, 0x44, 0x12, 0x0b, 0x0a, 0x07, 0x5f, 0x55, 0x4e, 0x5f, 0x55, 0x53, 0x45, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53,
	0x45, 0x4e, 0x44, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x10, 0x02, 0x12, 0x13, 0x0a,
	0x0f, 0x52, 0x45, 0x43, 0x45, 0x49, 0x56, 0x45, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45,
	0x10, 0x03, 0x42, 0x06, 0x5a, 0x04, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_echo_api_proto_rawDescOnce sync.Once
	file_echo_api_proto_rawDescData = file_echo_api_proto_rawDesc
)

func file_echo_api_proto_rawDescGZIP() []byte {
	file_echo_api_proto_rawDescOnce.Do(func() {
		file_echo_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_echo_api_proto_rawDescData)
	})
	return file_echo_api_proto_rawDescData
}

var file_echo_api_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_echo_api_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_echo_api_proto_goTypes = []interface{}{
	(ECHO_TYPE_ID)(0),           // 0: ECHO_TYPE_ID
	(SendMessage_ERROR)(0),      // 1: SendMessage.ERROR
	(*EchoPing)(nil),            // 2: EchoPing
	(*AppMessage)(nil),          // 3: AppMessage
	(*SendMessage)(nil),         // 4: SendMessage
	(*ReceiveMessage)(nil),      // 5: ReceiveMessage
	(*EchoPing_Request)(nil),    // 6: EchoPing.Request
	(*EchoPing_Reply)(nil),      // 7: EchoPing.Reply
	(*SendMessage_Request)(nil), // 8: SendMessage.Request
	(*SendMessage_Reply)(nil),   // 9: SendMessage.Reply
}
var file_echo_api_proto_depIdxs = []int32{
	3, // 0: ReceiveMessage.msg:type_name -> AppMessage
	3, // 1: SendMessage.Request.msg:type_name -> AppMessage
	3, // 2: SendMessage.Reply.msg:type_name -> AppMessage
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_echo_api_proto_init() }
func file_echo_api_proto_init() {
	if File_echo_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_echo_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoPing); i {
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
		file_echo_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppMessage); i {
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
		file_echo_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessage); i {
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
		file_echo_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveMessage); i {
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
		file_echo_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoPing_Request); i {
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
		file_echo_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoPing_Reply); i {
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
		file_echo_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessage_Request); i {
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
		file_echo_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessage_Reply); i {
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
			RawDescriptor: file_echo_api_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_echo_api_proto_goTypes,
		DependencyIndexes: file_echo_api_proto_depIdxs,
		EnumInfos:         file_echo_api_proto_enumTypes,
		MessageInfos:      file_echo_api_proto_msgTypes,
	}.Build()
	File_echo_api_proto = out.File
	file_echo_api_proto_rawDesc = nil
	file_echo_api_proto_goTypes = nil
	file_echo_api_proto_depIdxs = nil
}
