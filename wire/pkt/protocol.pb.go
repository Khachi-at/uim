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
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Isp   string   `protobuf:"bytes,2,opt,name=isp,proto3" json:"isp,omitempty"`
	Zone  string   `protobuf:"bytes,3,opt,name=zone,proto3" json:"zone,omitempty"` // location code
	Tags  []string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{0}
}

func (x *LoginReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginReq) GetIsp() string {
	if x != nil {
		return x.Isp
	}
	return ""
}

func (x *LoginReq) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

func (x *LoginReq) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type LoginResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId string `protobuf:"bytes,1,opt,name=channelId,proto3" json:"channelId,omitempty"`
}

func (x *LoginResp) Reset() {
	*x = LoginResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResp) ProtoMessage() {}

func (x *LoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResp.ProtoReflect.Descriptor instead.
func (*LoginResp) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResp) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

type KickoutNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId string `protobuf:"bytes,1,opt,name=channelId,proto3" json:"channelId,omitempty"`
}

func (x *KickoutNotify) Reset() {
	*x = KickoutNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KickoutNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KickoutNotify) ProtoMessage() {}

func (x *KickoutNotify) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KickoutNotify.ProtoReflect.Descriptor instead.
func (*KickoutNotify) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{2}
}

func (x *KickoutNotify) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId string   `protobuf:"bytes,1,opt,name=channelId,proto3" json:"channelId,omitempty"` // session id
	GateId    string   `protobuf:"bytes,2,opt,name=gateId,proto3" json:"gateId,omitempty"`       // gateway ID
	Account   string   `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`
	Zone      string   `protobuf:"bytes,4,opt,name=zone,proto3" json:"zone,omitempty"`
	Isp       string   `protobuf:"bytes,5,opt,name=isp,proto3" json:"isp,omitempty"`
	RemoteIP  string   `protobuf:"bytes,6,opt,name=remoteIP,proto3" json:"remoteIP,omitempty"`
	Device    string   `protobuf:"bytes,7,opt,name=device,proto3" json:"device,omitempty"`
	App       string   `protobuf:"bytes,8,opt,name=app,proto3" json:"app,omitempty"`
	Tags      []string `protobuf:"bytes,9,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[3]
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
	return file_protocol_proto_rawDescGZIP(), []int{3}
}

func (x *Session) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *Session) GetGateId() string {
	if x != nil {
		return x.GateId
	}
	return ""
}

func (x *Session) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *Session) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

func (x *Session) GetIsp() string {
	if x != nil {
		return x.Isp
	}
	return ""
}

func (x *Session) GetRemoteIP() string {
	if x != nil {
		return x.RemoteIP
	}
	return ""
}

func (x *Session) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *Session) GetApp() string {
	if x != nil {
		return x.App
	}
	return ""
}

func (x *Session) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

// chat message
type MessageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  int32  `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Body  string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Extra string `protobuf:"bytes,3,opt,name=extra,proto3" json:"extra,omitempty"`
}

func (x *MessageReq) Reset() {
	*x = MessageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageReq) ProtoMessage() {}

func (x *MessageReq) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageReq.ProtoReflect.Descriptor instead.
func (*MessageReq) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{4}
}

func (x *MessageReq) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *MessageReq) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *MessageReq) GetExtra() string {
	if x != nil {
		return x.Extra
	}
	return ""
}

type MessageResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId int64 `protobuf:"varint,1,opt,name=messageId,proto3" json:"messageId,omitempty"`
	SendTime  int64 `protobuf:"varint,2,opt,name=sendTime,proto3" json:"sendTime,omitempty"`
}

func (x *MessageResp) Reset() {
	*x = MessageResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageResp) ProtoMessage() {}

func (x *MessageResp) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageResp.ProtoReflect.Descriptor instead.
func (*MessageResp) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{5}
}

func (x *MessageResp) GetMessageId() int64 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *MessageResp) GetSendTime() int64 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

type MessagePush struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId int64  `protobuf:"varint,1,opt,name=messageId,proto3" json:"messageId,omitempty"`
	Type      int32  `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Body      string `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Extra     string `protobuf:"bytes,4,opt,name=extra,proto3" json:"extra,omitempty"`
	Sender    string `protobuf:"bytes,5,opt,name=sender,proto3" json:"sender,omitempty"`
	SendTime  int64  `protobuf:"varint,6,opt,name=sendTime,proto3" json:"sendTime,omitempty"`
}

func (x *MessagePush) Reset() {
	*x = MessagePush{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessagePush) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessagePush) ProtoMessage() {}

func (x *MessagePush) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessagePush.ProtoReflect.Descriptor instead.
func (*MessagePush) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{6}
}

func (x *MessagePush) GetMessageId() int64 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *MessagePush) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *MessagePush) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *MessagePush) GetExtra() string {
	if x != nil {
		return x.Extra
	}
	return ""
}

func (x *MessagePush) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *MessagePush) GetSendTime() int64 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

type ErrorResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ErrorResp) Reset() {
	*x = ErrorResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorResp) ProtoMessage() {}

func (x *ErrorResp) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorResp.ProtoReflect.Descriptor instead.
func (*ErrorResp) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{7}
}

func (x *ErrorResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type MessageAckReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId int64 `protobuf:"varint,1,opt,name=messageId,proto3" json:"messageId,omitempty"`
}

func (x *MessageAckReq) Reset() {
	*x = MessageAckReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageAckReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageAckReq) ProtoMessage() {}

func (x *MessageAckReq) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageAckReq.ProtoReflect.Descriptor instead.
func (*MessageAckReq) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{8}
}

func (x *MessageAckReq) GetMessageId() int64 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

type GroupCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Avatar       string   `protobuf:"bytes,2,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Introduction string   `protobuf:"bytes,3,opt,name=introduction,proto3" json:"introduction,omitempty"`
	Owner        string   `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	Members      []string `protobuf:"bytes,5,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *GroupCreateReq) Reset() {
	*x = GroupCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupCreateReq) ProtoMessage() {}

func (x *GroupCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupCreateReq.ProtoReflect.Descriptor instead.
func (*GroupCreateReq) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{9}
}

func (x *GroupCreateReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GroupCreateReq) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *GroupCreateReq) GetIntroduction() string {
	if x != nil {
		return x.Introduction
	}
	return ""
}

func (x *GroupCreateReq) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *GroupCreateReq) GetMembers() []string {
	if x != nil {
		return x.Members
	}
	return nil
}

type GroupCreateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *GroupCreateResp) Reset() {
	*x = GroupCreateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupCreateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupCreateResp) ProtoMessage() {}

func (x *GroupCreateResp) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupCreateResp.ProtoReflect.Descriptor instead.
func (*GroupCreateResp) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{10}
}

func (x *GroupCreateResp) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

type GroupCreateNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId string   `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Members []string `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *GroupCreateNotify) Reset() {
	*x = GroupCreateNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupCreateNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupCreateNotify) ProtoMessage() {}

func (x *GroupCreateNotify) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupCreateNotify.ProtoReflect.Descriptor instead.
func (*GroupCreateNotify) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{11}
}

func (x *GroupCreateNotify) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *GroupCreateNotify) GetMembers() []string {
	if x != nil {
		return x.Members
	}
	return nil
}

type GroupJoinReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	GroupId string `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *GroupJoinReq) Reset() {
	*x = GroupJoinReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupJoinReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupJoinReq) ProtoMessage() {}

func (x *GroupJoinReq) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupJoinReq.ProtoReflect.Descriptor instead.
func (*GroupJoinReq) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{12}
}

func (x *GroupJoinReq) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *GroupJoinReq) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

type GroupQuitReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	GroupId string `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *GroupQuitReq) Reset() {
	*x = GroupQuitReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupQuitReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupQuitReq) ProtoMessage() {}

func (x *GroupQuitReq) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupQuitReq.ProtoReflect.Descriptor instead.
func (*GroupQuitReq) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{13}
}

func (x *GroupQuitReq) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *GroupQuitReq) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

type GroupGetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *GroupGetReq) Reset() {
	*x = GroupGetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupGetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupGetReq) ProtoMessage() {}

func (x *GroupGetReq) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupGetReq.ProtoReflect.Descriptor instead.
func (*GroupGetReq) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{14}
}

func (x *GroupGetReq) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

type Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account  string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Alias    string `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
	Avatar   string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	JoinTime int64  `protobuf:"varint,4,opt,name=join_time,json=joinTime,proto3" json:"join_time,omitempty"`
}

func (x *Member) Reset() {
	*x = Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member.ProtoReflect.Descriptor instead.
func (*Member) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{15}
}

func (x *Member) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *Member) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *Member) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *Member) GetJoinTime() int64 {
	if x != nil {
		return x.JoinTime
	}
	return 0
}

type GroupGetResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Avatar       string    `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Introduction string    `protobuf:"bytes,4,opt,name=introduction,proto3" json:"introduction,omitempty"`
	Owner        string    `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
	Members      []*Member `protobuf:"bytes,6,rep,name=members,proto3" json:"members,omitempty"`
	CreatedAt    int64     `protobuf:"varint,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *GroupGetResp) Reset() {
	*x = GroupGetResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupGetResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupGetResp) ProtoMessage() {}

func (x *GroupGetResp) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupGetResp.ProtoReflect.Descriptor instead.
func (*GroupGetResp) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{16}
}

func (x *GroupGetResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GroupGetResp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GroupGetResp) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *GroupGetResp) GetIntroduction() string {
	if x != nil {
		return x.Introduction
	}
	return ""
}

func (x *GroupGetResp) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *GroupGetResp) GetMembers() []*Member {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *GroupGetResp) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type GroupJoinNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Account string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *GroupJoinNotify) Reset() {
	*x = GroupJoinNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[17]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupJoinNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupJoinNotify) ProtoMessage() {}

func (x *GroupJoinNotify) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[17]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupJoinNotify.ProtoReflect.Descriptor instead.
func (*GroupJoinNotify) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{17}
}

func (x *GroupJoinNotify) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *GroupJoinNotify) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

type GroupQuitNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Account string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *GroupQuitNotify) Reset() {
	*x = GroupQuitNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[18]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupQuitNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupQuitNotify) ProtoMessage() {}

func (x *GroupQuitNotify) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[18]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupQuitNotify.ProtoReflect.Descriptor instead.
func (*GroupQuitNotify) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{18}
}

func (x *GroupQuitNotify) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *GroupQuitNotify) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

type MessageIndexReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId int64 `protobuf:"varint,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
}

func (x *MessageIndexReq) Reset() {
	*x = MessageIndexReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[19]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageIndexReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageIndexReq) ProtoMessage() {}

func (x *MessageIndexReq) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[19]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageIndexReq.ProtoReflect.Descriptor instead.
func (*MessageIndexReq) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{19}
}

func (x *MessageIndexReq) GetMessageId() int64 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

type MessageIndexResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Indexes []*MessageIndex `protobuf:"bytes,1,rep,name=indexes,proto3" json:"indexes,omitempty"`
}

func (x *MessageIndexResp) Reset() {
	*x = MessageIndexResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[20]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageIndexResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageIndexResp) ProtoMessage() {}

func (x *MessageIndexResp) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[20]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageIndexResp.ProtoReflect.Descriptor instead.
func (*MessageIndexResp) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{20}
}

func (x *MessageIndexResp) GetIndexes() []*MessageIndex {
	if x != nil {
		return x.Indexes
	}
	return nil
}

type MessageIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId int64  `protobuf:"varint,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Direction int32  `protobuf:"varint,2,opt,name=direction,proto3" json:"direction,omitempty"`
	SendTime  int64  `protobuf:"varint,3,opt,name=send_time,json=sendTime,proto3" json:"send_time,omitempty"`
	AccountB  string `protobuf:"bytes,4,opt,name=accountB,proto3" json:"accountB,omitempty"`
	Group     string `protobuf:"bytes,5,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *MessageIndex) Reset() {
	*x = MessageIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[21]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageIndex) ProtoMessage() {}

func (x *MessageIndex) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[21]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageIndex.ProtoReflect.Descriptor instead.
func (*MessageIndex) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{21}
}

func (x *MessageIndex) GetMessageId() int64 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *MessageIndex) GetDirection() int32 {
	if x != nil {
		return x.Direction
	}
	return 0
}

func (x *MessageIndex) GetSendTime() int64 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

func (x *MessageIndex) GetAccountB() string {
	if x != nil {
		return x.AccountB
	}
	return ""
}

func (x *MessageIndex) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

type MessageContentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageIds []int64 `protobuf:"varint,1,rep,packed,name=message_ids,json=messageIds,proto3" json:"message_ids,omitempty"`
}

func (x *MessageContentReq) Reset() {
	*x = MessageContentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[22]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageContentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageContentReq) ProtoMessage() {}

func (x *MessageContentReq) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[22]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageContentReq.ProtoReflect.Descriptor instead.
func (*MessageContentReq) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{22}
}

func (x *MessageContentReq) GetMessageIds() []int64 {
	if x != nil {
		return x.MessageIds
	}
	return nil
}

type MessageContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId int64  `protobuf:"varint,1,opt,name=messageId,proto3" json:"messageId,omitempty"`
	Type      int32  `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Body      string `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Extra     string `protobuf:"bytes,4,opt,name=extra,proto3" json:"extra,omitempty"`
}

func (x *MessageContent) Reset() {
	*x = MessageContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[23]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageContent) ProtoMessage() {}

func (x *MessageContent) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[23]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageContent.ProtoReflect.Descriptor instead.
func (*MessageContent) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{23}
}

func (x *MessageContent) GetMessageId() int64 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *MessageContent) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *MessageContent) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *MessageContent) GetExtra() string {
	if x != nil {
		return x.Extra
	}
	return ""
}

type MessageContentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contents []*MessageContent `protobuf:"bytes,1,rep,name=contents,proto3" json:"contents,omitempty"`
}

func (x *MessageContentResp) Reset() {
	*x = MessageContentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[24]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageContentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageContentResp) ProtoMessage() {}

func (x *MessageContentResp) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[24]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageContentResp.ProtoReflect.Descriptor instead.
func (*MessageContentResp) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{24}
}

func (x *MessageContentResp) GetContents() []*MessageContent {
	if x != nil {
		return x.Contents
	}
	return nil
}

var File_protocol_proto protoreflect.FileDescriptor

var file_protocol_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x70, 0x6b, 0x74, 0x22, 0x5a, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x73, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x7a, 0x6f, 0x6e,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x22, 0x29, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x0d,
	0x4b, 0x69, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x22, 0xd9, 0x01, 0x0a, 0x07,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x73, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x73, 0x70, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x50, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x50, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x61, 0x70, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x4a, 0x0a, 0x0a, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x22, 0x47, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x9d, 0x01, 0x0a,
	0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x75, 0x73, 0x68, 0x12, 0x1c, 0x0a, 0x09,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x25, 0x0a, 0x09,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x2d, 0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x22, 0x90, 0x01, 0x0a, 0x0e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0x2c, 0x0a, 0x0f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x11, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0x43, 0x0a,
	0x0c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x22, 0x43, 0x0a, 0x0c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x51, 0x75, 0x69, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x0b, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x22, 0x6d, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0xca, 0x01, 0x0a, 0x0c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x22, 0x0a,
	0x0c, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x6b, 0x74, 0x2e, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x46, 0x0a,
	0x0f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4a, 0x6f, 0x69, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x46, 0x0a, 0x0f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x51, 0x75,
	0x69, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x30, 0x0a,
	0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x71,
	0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22,
	0x3f, 0x0a, 0x10, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x2b, 0x0a, 0x07, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x6b, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x07, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73,
	0x22, 0x9a, 0x01, 0x0a, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x34, 0x0a,
	0x11, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x73, 0x22, 0x6c, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x78, 0x74, 0x72, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x22, 0x45, 0x0a, 0x12, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2f, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x6b, 0x74, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x08,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x70, 0x6b,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_proto_rawDescOnce sync.Once
	file_protocol_proto_rawDescData = file_protocol_proto_rawDesc
)

func file_protocol_proto_rawDescGZIP() []byte {
	file_protocol_proto_rawDescOnce.Do(func() {
		file_protocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_proto_rawDescData)
	})
	return file_protocol_proto_rawDescData
}

var file_protocol_proto_msgTypes = make([]protoimpl.MessageInfo, 25)
var file_protocol_proto_goTypes = []interface{}{
	(*LoginReq)(nil),           // 0: pkt.LoginReq
	(*LoginResp)(nil),          // 1: pkt.LoginResp
	(*KickoutNotify)(nil),      // 2: pkt.KickoutNotify
	(*Session)(nil),            // 3: pkt.Session
	(*MessageReq)(nil),         // 4: pkt.MessageReq
	(*MessageResp)(nil),        // 5: pkt.MessageResp
	(*MessagePush)(nil),        // 6: pkt.MessagePush
	(*ErrorResp)(nil),          // 7: pkt.ErrorResp
	(*MessageAckReq)(nil),      // 8: pkt.MessageAckReq
	(*GroupCreateReq)(nil),     // 9: pkt.GroupCreateReq
	(*GroupCreateResp)(nil),    // 10: pkt.GroupCreateResp
	(*GroupCreateNotify)(nil),  // 11: pkt.GroupCreateNotify
	(*GroupJoinReq)(nil),       // 12: pkt.GroupJoinReq
	(*GroupQuitReq)(nil),       // 13: pkt.GroupQuitReq
	(*GroupGetReq)(nil),        // 14: pkt.GroupGetReq
	(*Member)(nil),             // 15: pkt.Member
	(*GroupGetResp)(nil),       // 16: pkt.GroupGetResp
	(*GroupJoinNotify)(nil),    // 17: pkt.GroupJoinNotify
	(*GroupQuitNotify)(nil),    // 18: pkt.GroupQuitNotify
	(*MessageIndexReq)(nil),    // 19: pkt.MessageIndexReq
	(*MessageIndexResp)(nil),   // 20: pkt.MessageIndexResp
	(*MessageIndex)(nil),       // 21: pkt.MessageIndex
	(*MessageContentReq)(nil),  // 22: pkt.MessageContentReq
	(*MessageContent)(nil),     // 23: pkt.MessageContent
	(*MessageContentResp)(nil), // 24: pkt.MessageContentResp
}
var file_protocol_proto_depIdxs = []int32{
	15, // 0: pkt.GroupGetResp.members:type_name -> pkt.Member
	21, // 1: pkt.MessageIndexResp.indexes:type_name -> pkt.MessageIndex
	23, // 2: pkt.MessageContentResp.contents:type_name -> pkt.MessageContent
	3,  // [3:3] is the sub-list for method output_type
	3,  // [3:3] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_protocol_proto_init() }
func file_protocol_proto_init() {
	if File_protocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_protocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResp); i {
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
		file_protocol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KickoutNotify); i {
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
		file_protocol_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_protocol_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageReq); i {
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
		file_protocol_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageResp); i {
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
		file_protocol_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessagePush); i {
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
		file_protocol_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorResp); i {
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
		file_protocol_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageAckReq); i {
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
		file_protocol_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupCreateReq); i {
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
		file_protocol_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupCreateResp); i {
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
		file_protocol_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupCreateNotify); i {
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
		file_protocol_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupJoinReq); i {
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
		file_protocol_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupQuitReq); i {
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
		file_protocol_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupGetReq); i {
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
		file_protocol_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Member); i {
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
		file_protocol_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupGetResp); i {
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
		file_protocol_proto_msgTypes[17].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupJoinNotify); i {
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
		file_protocol_proto_msgTypes[18].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupQuitNotify); i {
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
		file_protocol_proto_msgTypes[19].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageIndexReq); i {
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
		file_protocol_proto_msgTypes[20].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageIndexResp); i {
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
		file_protocol_proto_msgTypes[21].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageIndex); i {
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
		file_protocol_proto_msgTypes[22].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageContentReq); i {
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
		file_protocol_proto_msgTypes[23].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageContent); i {
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
		file_protocol_proto_msgTypes[24].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageContentResp); i {
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
			RawDescriptor: file_protocol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   25,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protocol_proto_goTypes,
		DependencyIndexes: file_protocol_proto_depIdxs,
		MessageInfos:      file_protocol_proto_msgTypes,
	}.Build()
	File_protocol_proto = out.File
	file_protocol_proto_rawDesc = nil
	file_protocol_proto_goTypes = nil
	file_protocol_proto_depIdxs = nil
}
