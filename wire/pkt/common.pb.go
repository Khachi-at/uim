package pkt

import (
	"reflect"
	"sync"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-ate.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status int32

const (
	Status_Success         Status = 0
	Status_SessionNotFound Status = 10
	// Client error 100-300.
	Status_NoDestination     Status = 100
	Status_InvalidPacketBody Status = 101
	Status_InvalidCommand    Status = 103
	Status_Unauthorized      Status = 105
	// Server error > 300.
	Status_SystemException Status = 500
	Status_NotImplemented  Status = 501
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0:   "Success",
		10:  "SessionNotFound",
		100: "NoDestination",
		101: "InvalidPacketBody",
		103: "InvalidCommand",
		105: "Unauthorized",
		500: "SystemException",
		501: "NotImplemented",
	}
	Status_value = map[string]int32{
		"Success":           0,
		"SessionNotFound":   10,
		"NoDestination":     100,
		"InvalidPacketBody": 101,
		"InvalidCommand":    103,
		"Unauthorized":      105,
		"SystemException":   500,
		"NotImplemented":    501,
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
	return file_common_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

type MetaType int32

const (
	MetaType_int    MetaType = 0
	MetaType_string MetaType = 1
	MetaType_float  MetaType = 2
)

// Enum value maps for MetaType.
var (
	MetaType_name = map[int32]string{
		0: "int",
		1: "string",
		2: "float",
	}

	MetaType_value = map[string]int32{
		"int":    0,
		"string": 1,
		"float":  2,
	}
)

func (x MetaType) Enum() *MetaType {
	p := new(MetaType)
	*p = x
	return p
}

func (x MetaType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetaType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[1].Descriptor()
}

func (MetaType) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[1]
}

func (x MetaType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

func (MetaType) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

type ContentType int32

const (
	ContentType_Protobuf ContentType = 0
	ContentType_Json     ContentType = 1
)

// Enum value maps for ContentType.
var (
	ContentType_name = map[int32]string{
		0: "Protobuf",
		1: "Json",
	}
	ContentType_value = map[string]int32{
		"Protobuf": 0,
		"Json":     1,
	}
)

func (x ContentType) Enum() *ContentType {
	p := new(ContentType)
	*p = x
	return p
}

func (x ContentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContentType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[2].Descriptor()
}

func (ContentType) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[2]
}

func (x ContentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContentType.Descriptor instead.
func (ContentType) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

type Flag int32

const (
	Flag_Request  Flag = 0
	Flag_Response Flag = 1
	Flag_Push     Flag = 2
)

// Enum value maps for Flag.
var (
	Flag_name = map[int32]string{
		0: "Request",
		1: "Response",
		2: "Push",
	}
	Flag_value = map[string]int32{
		"Request":  0,
		"Response": 1,
		"Push":     2,
	}
)

func (x Flag) Enum() *Flag {
	p := new(Flag)
	*p = x
	return p
}

func (x Flag) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Flag) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[3].Descriptor()
}

func (Flag) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[3]
}

func (x Flag) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Flag.Descriptor instead.
func (Flag) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Type  MetaType `protobuf:"varint,3,opt,name=type,proto3,enum=pkt.MetaType" json:"type,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *Meta) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Meta) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Meta) GetType() MetaType {
	if x != nil {
		return x.Type
	}
	return MetaType_int
}

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command   string  `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	ChannelId string  `protobuf:"bytes,2,opt,name=channelId,proto3" json:"channelId,omitempty"`
	Sequence  uint32  `protobuf:"varint,3,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Flag      Flag    `protobuf:"varint,4,opt,name=flag,proto3,enum=pkt.Flag" json:"flag,omitempty"`
	Status    Status  `protobuf:"varint,5,opt,name=status,proto3,enum=pkt.Status" json:"status,omitempty"`
	Dest      string  `protobuf:"bytes,6,opt,name=dest,proto3" json:"dest,omitempty"`
	Meta      []*Meta `protobuf:"bytes,7,rep,name=meta,proto3" json:"meta,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *Header) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *Header) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *Header) GetSequence() uint32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *Header) GetFlag() Flag {
	if x != nil {
		return x.Flag
	}
	return Flag_Request
}

func (x *Header) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_Success
}

func (x *Header) GetDest() string {
	if x != nil {
		return x.Dest
	}
	return ""
}

func (x *Header) GetMeta() []*Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type InnerHandshakeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceId string `protobuf:"bytes,1,opt,name=ServiceId,proto3" json:"ServiceId,omitempty"`
}

func (x *InnerHandshakeReq) Reset() {
	*x = InnerHandshakeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InnerHandshakeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InnerHandshakeReq) ProtoMessage() {}

func (x *InnerHandshakeReq) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InnerHandshakeReq.ProtoReflect.Descriptor instead.
func (*InnerHandshakeReq) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *InnerHandshakeReq) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

type InnerHandshakeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  uint32 `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *InnerHandshakeResponse) Reset() {
	*x = InnerHandshakeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InnerHandshakeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InnerHandshakeResponse) ProtoMessage() {}

func (x *InnerHandshakeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InnerHandshakeResponse.ProtoReflect.Descriptor instead.
func (*InnerHandshakeResponse) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

func (x *InnerHandshakeResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *InnerHandshakeResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x70, 0x6b, 0x74, 0x22, 0x51, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0d, 0x2e, 0x70, 0x6b, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xd3, 0x01, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x70, 0x6b, 0x74, 0x2e, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x04,
	0x66, 0x6c, 0x61, 0x67, 0x12, 0x23, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x6b, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x6b,
	0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x31, 0x0a, 0x11,
	0x49, 0x6e, 0x6e, 0x65, 0x72, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22,
	0x42, 0x0a, 0x16, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x2a, 0xa5, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b,
	0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x0a,
	0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x6f, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x10, 0x64, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x50, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x10, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x6e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x10, 0x67, 0x12, 0x10,
	0x0a, 0x0c, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x10, 0x69,
	0x12, 0x14, 0x0a, 0x0f, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x10, 0xf4, 0x03, 0x12, 0x13, 0x0a, 0x0e, 0x4e, 0x6f, 0x74, 0x49, 0x6d, 0x70,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x10, 0xf5, 0x03, 0x2a, 0x2a, 0x0a, 0x08, 0x4d,
	0x65, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x69, 0x6e, 0x74, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05,
	0x66, 0x6c, 0x6f, 0x61, 0x74, 0x10, 0x02, 0x2a, 0x25, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x73, 0x6f, 0x6e, 0x10, 0x01, 0x2a, 0x2b,
	0x0a, 0x04, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x10,
	0x01, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x75, 0x73, 0x68, 0x10, 0x02, 0x42, 0x07, 0x5a, 0x05, 0x2e,
	0x2f, 0x70, 0x6b, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_common_proto_goTypes = []interface{}{
	(Status)(0),                    // 0: pkt.Status
	(MetaType)(0),                  // 1: pkt.MetaType
	(ContentType)(0),               // 2: pkt.ContentType
	(Flag)(0),                      // 3: pkt.Flag
	(*Meta)(nil),                   // 4: pkt.Meta
	(*Header)(nil),                 // 5: pkt.Header
	(*InnerHandshakeReq)(nil),      // 6: pkt.InnerHandshakeReq
	(*InnerHandshakeResponse)(nil), // 7: pkt.InnerHandshakeResponse
}
var file_common_proto_depIdxs = []int32{
	1, // 0: pkt.Meta.type:type_name -> pkt.MetaType
	3, // 1: pkt.Header.flag:type_name -> pkt.Flag
	0, // 2: pkt.Header.status:type_name -> pkt.Status
	4, // 3: pkt.Header.meta:type_name -> pkt.Meta
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InnerHandshakeReq); i {
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
		file_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InnerHandshakeResponse); i {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
