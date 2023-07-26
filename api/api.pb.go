// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: api.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResultType int32

const (
	ResultType_OK                    ResultType = 0
	ResultType_REQUEST_INVALID       ResultType = 1000
	ResultType_TOKEN_INVALID         ResultType = 1001
	ResultType_TOKEN_EXPIRE          ResultType = 1002
	ResultType_SIZE_LIMITED          ResultType = 1003
	ResultType_DB_ERROR              ResultType = 1004
	ResultType_NO_CHANGED            ResultType = 1005
	ResultType_NETWORK_ERROR         ResultType = 1006
	ResultType_DATA_ERROR            ResultType = 1007
	ResultType_ENCRYPT_ERROR         ResultType = 1008 /// lỗi khi chưa thiết lập kết nối authen
	ResultType_INTERNAL_SERVER_ERROR ResultType = 1009
	ResultType_STRANGE_REQUEST       ResultType = 1010
	ResultType_PERMISSION_DENIED     ResultType = 1011
)

// Enum value maps for ResultType.
var (
	ResultType_name = map[int32]string{
		0:    "OK",
		1000: "REQUEST_INVALID",
		1001: "TOKEN_INVALID",
		1002: "TOKEN_EXPIRE",
		1003: "SIZE_LIMITED",
		1004: "DB_ERROR",
		1005: "NO_CHANGED",
		1006: "NETWORK_ERROR",
		1007: "DATA_ERROR",
		1008: "ENCRYPT_ERROR",
		1009: "INTERNAL_SERVER_ERROR",
		1010: "STRANGE_REQUEST",
		1011: "PERMISSION_DENIED",
	}
	ResultType_value = map[string]int32{
		"OK":                    0,
		"REQUEST_INVALID":       1000,
		"TOKEN_INVALID":         1001,
		"TOKEN_EXPIRE":          1002,
		"SIZE_LIMITED":          1003,
		"DB_ERROR":              1004,
		"NO_CHANGED":            1005,
		"NETWORK_ERROR":         1006,
		"DATA_ERROR":            1007,
		"ENCRYPT_ERROR":         1008,
		"INTERNAL_SERVER_ERROR": 1009,
		"STRANGE_REQUEST":       1010,
		"PERMISSION_DENIED":     1011,
	}
)

func (x ResultType) Enum() *ResultType {
	p := new(ResultType)
	*p = x
	return p
}

func (x ResultType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResultType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[0].Descriptor()
}

func (ResultType) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[0]
}

func (x ResultType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResultType.Descriptor instead.
func (ResultType) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

type Platform int32

const (
	Platform_IOS     Platform = 0
	Platform_ANDROID Platform = 1
	Platform_WEB     Platform = 2
	Platform_WINDOW  Platform = 3
	Platform_MAC     Platform = 4
	Platform_LINUX   Platform = 5
	Platform_OTHER   Platform = 99
)

// Enum value maps for Platform.
var (
	Platform_name = map[int32]string{
		0:  "IOS",
		1:  "ANDROID",
		2:  "WEB",
		3:  "WINDOW",
		4:  "MAC",
		5:  "LINUX",
		99: "OTHER",
	}
	Platform_value = map[string]int32{
		"IOS":     0,
		"ANDROID": 1,
		"WEB":     2,
		"WINDOW":  3,
		"MAC":     4,
		"LINUX":   5,
		"OTHER":   99,
	}
)

func (x Platform) Enum() *Platform {
	p := new(Platform)
	*p = x
	return p
}

func (x Platform) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Platform) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[1].Descriptor()
}

func (Platform) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[1]
}

func (x Platform) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Platform.Descriptor instead.
func (Platform) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

type EncodeType int32

const (
	EncodeType_NONE EncodeType = 0  // không mã hóa playload
	EncodeType_XOR  EncodeType = 32 // mã hóa với xor key ( sv sử dụng xor key của client và client sử dụng xor key của server để decode payload )
	EncodeType_RSA  EncodeType = 64 // mã hóa với rsa public key của nhau để encode data (RSA-1024-PKCS1_PADING)
	EncodeType_AES  EncodeType = 96 // with key = 32 byte (AES-CBC-256-PKCS7_PADING)
)

// Enum value maps for EncodeType.
var (
	EncodeType_name = map[int32]string{
		0:  "NONE",
		32: "XOR",
		64: "RSA",
		96: "AES",
	}
	EncodeType_value = map[string]int32{
		"NONE": 0,
		"XOR":  32,
		"RSA":  64,
		"AES":  96,
	}
)

func (x EncodeType) Enum() *EncodeType {
	p := new(EncodeType)
	*p = x
	return p
}

func (x EncodeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EncodeType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[2].Descriptor()
}

func (EncodeType) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[2]
}

func (x EncodeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EncodeType.Descriptor instead.
func (EncodeType) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

type UserRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Code   string `protobuf:"bytes,2,opt,name=Code,proto3" json:"Code,omitempty"`
	RoleID uint32 `protobuf:"varint,3,opt,name=RoleID,proto3" json:"RoleID,omitempty"`
}

func (x *UserRole) Reset() {
	*x = UserRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRole) ProtoMessage() {}

func (x *UserRole) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRole.ProtoReflect.Descriptor instead.
func (*UserRole) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *UserRole) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserRole) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UserRole) GetRoleID() uint32 {
	if x != nil {
		return x.RoleID
	}
	return 0
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token        string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserId       uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // = account id
	SessionId    uint32 `protobuf:"varint,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	ConnectionId string `protobuf:"bytes,4,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	Expire       uint64 `protobuf:"varint,5,opt,name=expire,proto3" json:"expire,omitempty"`
	LastActive   uint64 `protobuf:"varint,6,opt,name=last_active,json=lastActive,proto3" json:"last_active,omitempty"`
	Appid        uint32 `protobuf:"varint,7,opt,name=appid,proto3" json:"appid,omitempty"`    // chứa id của app kết nối tới
	AppCode      string `protobuf:"bytes,8,opt,name=appCode,proto3" json:"appCode,omitempty"` // chứa mã code của app kết nối tới
	CitizenId    string `protobuf:"bytes,9,opt,name=citizen_id,json=citizenId,proto3" json:"citizen_id,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *Session) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Session) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Session) GetSessionId() uint32 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *Session) GetConnectionId() string {
	if x != nil {
		return x.ConnectionId
	}
	return ""
}

func (x *Session) GetExpire() uint64 {
	if x != nil {
		return x.Expire
	}
	return 0
}

func (x *Session) GetLastActive() uint64 {
	if x != nil {
		return x.LastActive
	}
	return 0
}

func (x *Session) GetAppid() uint32 {
	if x != nil {
		return x.Appid
	}
	return 0
}

func (x *Session) GetAppCode() string {
	if x != nil {
		return x.AppCode
	}
	return ""
}

func (x *Session) GetCitizenId() string {
	if x != nil {
		return x.CitizenId
	}
	return ""
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	FilePath string `protobuf:"bytes,2,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	FileSize int32  `protobuf:"varint,3,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	FileExt  string `protobuf:"bytes,4,opt,name=file_ext,json=fileExt,proto3" json:"file_ext,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *File) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *File) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *File) GetFileSize() int32 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *File) GetFileExt() string {
	if x != nil {
		return x.FileExt
	}
	return ""
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        uint32     `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Group       uint32     `protobuf:"varint,2,opt,name=group,proto3" json:"group,omitempty"`
	Session     *Session   `protobuf:"bytes,3,opt,name=session,proto3" json:"session,omitempty"`                         // client không set khi sử dụng http (http sử dụng V-Authorization ở Header mỗi request), đối với grpc và udp cần gửi thêm session_id trong session để sv kiểm tra tính xác thực user
	Protocol    uint32     `protobuf:"varint,4,opt,name=protocol,proto3" json:"protocol,omitempty"`                      // client không set
	PayloadType uint32     `protobuf:"varint,5,opt,name=payloadType,proto3" json:"payloadType,omitempty"`                // client không set
	BinRequest  []byte     `protobuf:"bytes,6,opt,name=bin_request,json=binRequest,proto3" json:"bin_request,omitempty"` // client không set
	Request     *anypb.Any `protobuf:"bytes,7,opt,name=request,proto3" json:"request,omitempty"`
	Files       []*File    `protobuf:"bytes,8,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *Request) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Request) GetGroup() uint32 {
	if x != nil {
		return x.Group
	}
	return 0
}

func (x *Request) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

func (x *Request) GetProtocol() uint32 {
	if x != nil {
		return x.Protocol
	}
	return 0
}

func (x *Request) GetPayloadType() uint32 {
	if x != nil {
		return x.PayloadType
	}
	return 0
}

func (x *Request) GetBinRequest() []byte {
	if x != nil {
		return x.BinRequest
	}
	return nil
}

func (x *Request) GetRequest() *anypb.Any {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *Request) GetFiles() []*File {
	if x != nil {
		return x.Files
	}
	return nil
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   uint32     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg      string     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data     *anypb.Any `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	BinReply []byte     `protobuf:"bytes,4,opt,name=bin_reply,json=binReply,proto3" json:"bin_reply,omitempty"` // trường này luôn luôn null khi được trả về client
	Type     uint32     `protobuf:"varint,5,opt,name=type,proto3" json:"type,omitempty"`
	Group    uint32     `protobuf:"varint,6,opt,name=group,proto3" json:"group,omitempty"`
	Code     uint32     `protobuf:"varint,7,opt,name=code,proto3" json:"code,omitempty"`
	Message  string     `protobuf:"bytes,8,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *Reply) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Reply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Reply) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Reply) GetBinReply() []byte {
	if x != nil {
		return x.BinReply
	}
	return nil
}

func (x *Reply) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Reply) GetGroup() uint32 {
	if x != nil {
		return x.Group
	}
	return 0
}

func (x *Reply) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Reply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Receive struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       uint32     `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Group      uint32     `protobuf:"varint,2,opt,name=group,proto3" json:"group,omitempty"`
	ServerTime uint64     `protobuf:"varint,3,opt,name=server_time,json=serverTime,proto3" json:"server_time,omitempty"`
	Receive    *anypb.Any `protobuf:"bytes,4,opt,name=receive,proto3" json:"receive,omitempty"`
}

func (x *Receive) Reset() {
	*x = Receive{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Receive) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Receive) ProtoMessage() {}

func (x *Receive) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Receive.ProtoReflect.Descriptor instead.
func (*Receive) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *Receive) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Receive) GetGroup() uint32 {
	if x != nil {
		return x.Group
	}
	return 0
}

func (x *Receive) GetServerTime() uint64 {
	if x != nil {
		return x.ServerTime
	}
	return 0
}

func (x *Receive) GetReceive() *anypb.Any {
	if x != nil {
		return x.Receive
	}
	return nil
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f,
	0x6c, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x52, 0x6f, 0x6c, 0x65,
	0x49, 0x44, 0x22, 0x84, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a,
	0x6c, 0x61, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70,
	0x70, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x69,
	0x74, 0x69, 0x7a, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e, 0x49, 0x64, 0x22, 0x78, 0x0a, 0x04, 0x46, 0x69, 0x6c,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65,
	0x45, 0x78, 0x74, 0x22, 0x83, 0x02, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x22, 0x0a, 0x07, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62,
	0x69, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0a, 0x62, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x05,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0xd0, 0x01, 0x0a, 0x05, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x28, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x69, 0x6e, 0x5f, 0x72,
	0x65, 0x70, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x62, 0x69, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x84, 0x01, 0x0a,
	0x07, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x2a, 0x87, 0x02, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x0f, 0x52, 0x45,
	0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0xe8, 0x07,
	0x12, 0x12, 0x0a, 0x0d, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x10, 0xe9, 0x07, 0x12, 0x11, 0x0a, 0x0c, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x45, 0x58,
	0x50, 0x49, 0x52, 0x45, 0x10, 0xea, 0x07, 0x12, 0x11, 0x0a, 0x0c, 0x53, 0x49, 0x5a, 0x45, 0x5f,
	0x4c, 0x49, 0x4d, 0x49, 0x54, 0x45, 0x44, 0x10, 0xeb, 0x07, 0x12, 0x0d, 0x0a, 0x08, 0x44, 0x42,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xec, 0x07, 0x12, 0x0f, 0x0a, 0x0a, 0x4e, 0x4f, 0x5f,
	0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0xed, 0x07, 0x12, 0x12, 0x0a, 0x0d, 0x4e, 0x45,
	0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xee, 0x07, 0x12, 0x0f,
	0x0a, 0x0a, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xef, 0x07, 0x12,
	0x12, 0x0a, 0x0d, 0x45, 0x4e, 0x43, 0x52, 0x59, 0x50, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0xf0, 0x07, 0x12, 0x1a, 0x0a, 0x15, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f,
	0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xf1, 0x07, 0x12,
	0x14, 0x0a, 0x0f, 0x53, 0x54, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45,
	0x53, 0x54, 0x10, 0xf2, 0x07, 0x12, 0x16, 0x0a, 0x11, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53,
	0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x4e, 0x49, 0x45, 0x44, 0x10, 0xf3, 0x07, 0x2a, 0x54, 0x0a,
	0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x4f, 0x53,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x10, 0x01, 0x12,
	0x07, 0x0a, 0x03, 0x57, 0x45, 0x42, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x57, 0x49, 0x4e, 0x44,
	0x4f, 0x57, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x41, 0x43, 0x10, 0x04, 0x12, 0x09, 0x0a,
	0x05, 0x4c, 0x49, 0x4e, 0x55, 0x58, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x54, 0x48, 0x45,
	0x52, 0x10, 0x63, 0x2a, 0x31, 0x0a, 0x0a, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x58,
	0x4f, 0x52, 0x10, 0x20, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x53, 0x41, 0x10, 0x40, 0x12, 0x07, 0x0a,
	0x03, 0x41, 0x45, 0x53, 0x10, 0x60, 0x42, 0x06, 0x5a, 0x04, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_proto_goTypes = []interface{}{
	(ResultType)(0),   // 0: ResultType
	(Platform)(0),     // 1: Platform
	(EncodeType)(0),   // 2: EncodeType
	(*UserRole)(nil),  // 3: UserRole
	(*Session)(nil),   // 4: Session
	(*File)(nil),      // 5: File
	(*Request)(nil),   // 6: Request
	(*Reply)(nil),     // 7: Reply
	(*Receive)(nil),   // 8: Receive
	(*anypb.Any)(nil), // 9: google.protobuf.Any
}
var file_api_proto_depIdxs = []int32{
	4, // 0: Request.session:type_name -> Session
	9, // 1: Request.request:type_name -> google.protobuf.Any
	5, // 2: Request.files:type_name -> File
	9, // 3: Reply.data:type_name -> google.protobuf.Any
	9, // 4: Receive.receive:type_name -> google.protobuf.Any
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRole); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Receive); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		EnumInfos:         file_api_proto_enumTypes,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
