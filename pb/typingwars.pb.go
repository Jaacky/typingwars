// Code generated by protoc-gen-go. DO NOT EDIT.
// source: typingwars.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UserInput int32

const (
	UserInput_KEY UserInput = 0
)

var UserInput_name = map[int32]string{
	0: "KEY",
}

var UserInput_value = map[string]int32{
	"KEY": 0,
}

func (x UserInput) String() string {
	return proto.EnumName(UserInput_name, int32(x))
}

func (UserInput) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1b666822ed8f2be0, []int{0}
}

type Message struct {
	// Types that are valid to be assigned to Content:
	//	*Message_PlayerJoined
	Content              isMessage_Content `protobuf_oneof:"content"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b666822ed8f2be0, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

type isMessage_Content interface {
	isMessage_Content()
}

type Message_PlayerJoined struct {
	PlayerJoined *PlayerJoined `protobuf:"bytes,1,opt,name=playerJoined,proto3,oneof"`
}

func (*Message_PlayerJoined) isMessage_Content() {}

func (m *Message) GetContent() isMessage_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Message) GetPlayerJoined() *PlayerJoined {
	if x, ok := m.GetContent().(*Message_PlayerJoined); ok {
		return x.PlayerJoined
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Message) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Message_OneofMarshaler, _Message_OneofUnmarshaler, _Message_OneofSizer, []interface{}{
		(*Message_PlayerJoined)(nil),
	}
}

func _Message_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Message)
	// content
	switch x := m.Content.(type) {
	case *Message_PlayerJoined:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PlayerJoined); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Message.Content has unexpected type %T", x)
	}
	return nil
}

func _Message_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Message)
	switch tag {
	case 1: // content.playerJoined
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PlayerJoined)
		err := b.DecodeMessage(msg)
		m.Content = &Message_PlayerJoined{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Message_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Message)
	// content
	switch x := m.Content.(type) {
	case *Message_PlayerJoined:
		s := proto.Size(x.PlayerJoined)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type PlayerJoined struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerJoined) Reset()         { *m = PlayerJoined{} }
func (m *PlayerJoined) String() string { return proto.CompactTextString(m) }
func (*PlayerJoined) ProtoMessage()    {}
func (*PlayerJoined) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b666822ed8f2be0, []int{1}
}

func (m *PlayerJoined) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerJoined.Unmarshal(m, b)
}
func (m *PlayerJoined) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerJoined.Marshal(b, m, deterministic)
}
func (m *PlayerJoined) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerJoined.Merge(m, src)
}
func (m *PlayerJoined) XXX_Size() int {
	return xxx_messageInfo_PlayerJoined.Size(m)
}
func (m *PlayerJoined) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerJoined.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerJoined proto.InternalMessageInfo

func (m *PlayerJoined) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PlayerJoined) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type UserAction struct {
	UserInput            UserInput `protobuf:"varint,1,opt,name=userInput,proto3,enum=typingwars.UserInput" json:"userInput,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UserAction) Reset()         { *m = UserAction{} }
func (m *UserAction) String() string { return proto.CompactTextString(m) }
func (*UserAction) ProtoMessage()    {}
func (*UserAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b666822ed8f2be0, []int{2}
}

func (m *UserAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAction.Unmarshal(m, b)
}
func (m *UserAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAction.Marshal(b, m, deterministic)
}
func (m *UserAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAction.Merge(m, src)
}
func (m *UserAction) XXX_Size() int {
	return xxx_messageInfo_UserAction.Size(m)
}
func (m *UserAction) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAction.DiscardUnknown(m)
}

var xxx_messageInfo_UserAction proto.InternalMessageInfo

func (m *UserAction) GetUserInput() UserInput {
	if m != nil {
		return m.UserInput
	}
	return UserInput_KEY
}

type RegisterPlayer struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterPlayer) Reset()         { *m = RegisterPlayer{} }
func (m *RegisterPlayer) String() string { return proto.CompactTextString(m) }
func (*RegisterPlayer) ProtoMessage()    {}
func (*RegisterPlayer) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b666822ed8f2be0, []int{3}
}

func (m *RegisterPlayer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterPlayer.Unmarshal(m, b)
}
func (m *RegisterPlayer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterPlayer.Marshal(b, m, deterministic)
}
func (m *RegisterPlayer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterPlayer.Merge(m, src)
}
func (m *RegisterPlayer) XXX_Size() int {
	return xxx_messageInfo_RegisterPlayer.Size(m)
}
func (m *RegisterPlayer) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterPlayer.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterPlayer proto.InternalMessageInfo

func (m *RegisterPlayer) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type CreateGameRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateGameRequest) Reset()         { *m = CreateGameRequest{} }
func (m *CreateGameRequest) String() string { return proto.CompactTextString(m) }
func (*CreateGameRequest) ProtoMessage()    {}
func (*CreateGameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b666822ed8f2be0, []int{4}
}

func (m *CreateGameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateGameRequest.Unmarshal(m, b)
}
func (m *CreateGameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateGameRequest.Marshal(b, m, deterministic)
}
func (m *CreateGameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateGameRequest.Merge(m, src)
}
func (m *CreateGameRequest) XXX_Size() int {
	return xxx_messageInfo_CreateGameRequest.Size(m)
}
func (m *CreateGameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateGameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateGameRequest proto.InternalMessageInfo

func (m *CreateGameRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type JoinGameRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	RoomId               string   `protobuf:"bytes,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinGameRequest) Reset()         { *m = JoinGameRequest{} }
func (m *JoinGameRequest) String() string { return proto.CompactTextString(m) }
func (*JoinGameRequest) ProtoMessage()    {}
func (*JoinGameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b666822ed8f2be0, []int{5}
}

func (m *JoinGameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinGameRequest.Unmarshal(m, b)
}
func (m *JoinGameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinGameRequest.Marshal(b, m, deterministic)
}
func (m *JoinGameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinGameRequest.Merge(m, src)
}
func (m *JoinGameRequest) XXX_Size() int {
	return xxx_messageInfo_JoinGameRequest.Size(m)
}
func (m *JoinGameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinGameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JoinGameRequest proto.InternalMessageInfo

func (m *JoinGameRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *JoinGameRequest) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

type UserMessage struct {
	// Types that are valid to be assigned to Content:
	//	*UserMessage_UserAction
	//	*UserMessage_RegisterPlayer
	//	*UserMessage_CreateGameRequest
	//	*UserMessage_JoinGameRequest
	Content              isUserMessage_Content `protobuf_oneof:"content"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UserMessage) Reset()         { *m = UserMessage{} }
func (m *UserMessage) String() string { return proto.CompactTextString(m) }
func (*UserMessage) ProtoMessage()    {}
func (*UserMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b666822ed8f2be0, []int{6}
}

func (m *UserMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserMessage.Unmarshal(m, b)
}
func (m *UserMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserMessage.Marshal(b, m, deterministic)
}
func (m *UserMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserMessage.Merge(m, src)
}
func (m *UserMessage) XXX_Size() int {
	return xxx_messageInfo_UserMessage.Size(m)
}
func (m *UserMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_UserMessage.DiscardUnknown(m)
}

var xxx_messageInfo_UserMessage proto.InternalMessageInfo

type isUserMessage_Content interface {
	isUserMessage_Content()
}

type UserMessage_UserAction struct {
	UserAction *UserAction `protobuf:"bytes,1,opt,name=userAction,proto3,oneof"`
}

type UserMessage_RegisterPlayer struct {
	RegisterPlayer *RegisterPlayer `protobuf:"bytes,2,opt,name=registerPlayer,proto3,oneof"`
}

type UserMessage_CreateGameRequest struct {
	CreateGameRequest *CreateGameRequest `protobuf:"bytes,3,opt,name=createGameRequest,proto3,oneof"`
}

type UserMessage_JoinGameRequest struct {
	JoinGameRequest *JoinGameRequest `protobuf:"bytes,4,opt,name=joinGameRequest,proto3,oneof"`
}

func (*UserMessage_UserAction) isUserMessage_Content() {}

func (*UserMessage_RegisterPlayer) isUserMessage_Content() {}

func (*UserMessage_CreateGameRequest) isUserMessage_Content() {}

func (*UserMessage_JoinGameRequest) isUserMessage_Content() {}

func (m *UserMessage) GetContent() isUserMessage_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *UserMessage) GetUserAction() *UserAction {
	if x, ok := m.GetContent().(*UserMessage_UserAction); ok {
		return x.UserAction
	}
	return nil
}

func (m *UserMessage) GetRegisterPlayer() *RegisterPlayer {
	if x, ok := m.GetContent().(*UserMessage_RegisterPlayer); ok {
		return x.RegisterPlayer
	}
	return nil
}

func (m *UserMessage) GetCreateGameRequest() *CreateGameRequest {
	if x, ok := m.GetContent().(*UserMessage_CreateGameRequest); ok {
		return x.CreateGameRequest
	}
	return nil
}

func (m *UserMessage) GetJoinGameRequest() *JoinGameRequest {
	if x, ok := m.GetContent().(*UserMessage_JoinGameRequest); ok {
		return x.JoinGameRequest
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*UserMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _UserMessage_OneofMarshaler, _UserMessage_OneofUnmarshaler, _UserMessage_OneofSizer, []interface{}{
		(*UserMessage_UserAction)(nil),
		(*UserMessage_RegisterPlayer)(nil),
		(*UserMessage_CreateGameRequest)(nil),
		(*UserMessage_JoinGameRequest)(nil),
	}
}

func _UserMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*UserMessage)
	// content
	switch x := m.Content.(type) {
	case *UserMessage_UserAction:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UserAction); err != nil {
			return err
		}
	case *UserMessage_RegisterPlayer:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RegisterPlayer); err != nil {
			return err
		}
	case *UserMessage_CreateGameRequest:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CreateGameRequest); err != nil {
			return err
		}
	case *UserMessage_JoinGameRequest:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.JoinGameRequest); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("UserMessage.Content has unexpected type %T", x)
	}
	return nil
}

func _UserMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*UserMessage)
	switch tag {
	case 1: // content.userAction
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UserAction)
		err := b.DecodeMessage(msg)
		m.Content = &UserMessage_UserAction{msg}
		return true, err
	case 2: // content.registerPlayer
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RegisterPlayer)
		err := b.DecodeMessage(msg)
		m.Content = &UserMessage_RegisterPlayer{msg}
		return true, err
	case 3: // content.createGameRequest
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CreateGameRequest)
		err := b.DecodeMessage(msg)
		m.Content = &UserMessage_CreateGameRequest{msg}
		return true, err
	case 4: // content.joinGameRequest
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(JoinGameRequest)
		err := b.DecodeMessage(msg)
		m.Content = &UserMessage_JoinGameRequest{msg}
		return true, err
	default:
		return false, nil
	}
}

func _UserMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*UserMessage)
	// content
	switch x := m.Content.(type) {
	case *UserMessage_UserAction:
		s := proto.Size(x.UserAction)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UserMessage_RegisterPlayer:
		s := proto.Size(x.RegisterPlayer)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UserMessage_CreateGameRequest:
		s := proto.Size(x.CreateGameRequest)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UserMessage_JoinGameRequest:
		s := proto.Size(x.JoinGameRequest)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterEnum("typingwars.UserInput", UserInput_name, UserInput_value)
	proto.RegisterType((*Message)(nil), "typingwars.Message")
	proto.RegisterType((*PlayerJoined)(nil), "typingwars.PlayerJoined")
	proto.RegisterType((*UserAction)(nil), "typingwars.UserAction")
	proto.RegisterType((*RegisterPlayer)(nil), "typingwars.RegisterPlayer")
	proto.RegisterType((*CreateGameRequest)(nil), "typingwars.CreateGameRequest")
	proto.RegisterType((*JoinGameRequest)(nil), "typingwars.JoinGameRequest")
	proto.RegisterType((*UserMessage)(nil), "typingwars.UserMessage")
}

func init() { proto.RegisterFile("typingwars.proto", fileDescriptor_1b666822ed8f2be0) }

var fileDescriptor_1b666822ed8f2be0 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4b, 0x4b, 0xeb, 0x40,
	0x14, 0x4e, 0xd3, 0xd2, 0xdc, 0x9c, 0x96, 0xb4, 0x1d, 0xee, 0xbd, 0x86, 0x8a, 0x20, 0x59, 0x89,
	0x48, 0x85, 0x76, 0x23, 0x2e, 0x84, 0xd6, 0x47, 0x53, 0xa5, 0x20, 0x83, 0x5d, 0xe8, 0x46, 0xd2,
	0xe6, 0x50, 0x22, 0x76, 0x26, 0xce, 0x4c, 0x90, 0xfe, 0x2b, 0x7f, 0xa2, 0x24, 0x7d, 0x4d, 0x52,
	0x10, 0x97, 0x33, 0xdf, 0xe3, 0x9c, 0xef, 0x9b, 0x81, 0xa6, 0x5a, 0xc6, 0x11, 0x9b, 0x7f, 0x06,
	0x42, 0x76, 0x62, 0xc1, 0x15, 0x27, 0xb0, 0xbb, 0xf1, 0x9e, 0xc0, 0x1a, 0xa3, 0x94, 0xc1, 0x1c,
	0xc9, 0x15, 0xd4, 0xe3, 0xf7, 0x60, 0x89, 0xe2, 0x9e, 0x47, 0x0c, 0x43, 0xb7, 0x74, 0x5c, 0x3a,
	0xa9, 0x75, 0xdd, 0x8e, 0xa6, 0x7f, 0xd4, 0x70, 0xdf, 0xa0, 0x39, 0xfe, 0xc0, 0x06, 0x6b, 0xc6,
	0x99, 0x42, 0xa6, 0xbc, 0x4b, 0xa8, 0xeb, 0x54, 0xe2, 0x80, 0x19, 0xad, 0x0c, 0x6d, 0x6a, 0x46,
	0x21, 0x69, 0xc3, 0x9f, 0x44, 0xa2, 0x60, 0xc1, 0x02, 0x5d, 0x33, 0xbb, 0xdd, 0x9e, 0xbd, 0x3e,
	0xc0, 0x44, 0xa2, 0xe8, 0xcf, 0x54, 0xc4, 0x19, 0xe9, 0x81, 0x9d, 0x22, 0x23, 0x16, 0x27, 0x2a,
	0x33, 0x70, 0xba, 0xff, 0xf4, 0x8d, 0x26, 0x1b, 0x90, 0xee, 0x78, 0xde, 0x19, 0x38, 0x14, 0xe7,
	0x91, 0x54, 0x28, 0x56, 0x6b, 0xe4, 0x06, 0x96, 0x0a, 0x03, 0xcf, 0xa1, 0x75, 0x2d, 0x30, 0x50,
	0x38, 0x0c, 0x16, 0x48, 0xf1, 0x23, 0x41, 0xa9, 0x7e, 0x14, 0xdc, 0x41, 0x23, 0xcd, 0xf5, 0x4b,
	0x3a, 0x39, 0x00, 0x4b, 0x70, 0xbe, 0x78, 0x8d, 0xc2, 0x75, 0xd6, 0x6a, 0x7a, 0x1c, 0x85, 0xde,
	0x97, 0x09, 0xb5, 0x74, 0xff, 0xcd, 0x03, 0x5c, 0x00, 0x24, 0xdb, 0xe4, 0xeb, 0xfa, 0xff, 0x17,
	0xc3, 0xae, 0x50, 0xdf, 0xa0, 0x1a, 0x97, 0xdc, 0x80, 0x23, 0x72, 0x81, 0xb3, 0x49, 0xb5, 0x6e,
	0x5b, 0x57, 0xe7, 0x2b, 0xf1, 0x0d, 0x5a, 0xd0, 0x90, 0x31, 0xb4, 0x66, 0xc5, 0x22, 0xdc, 0x72,
	0x66, 0x74, 0xa4, 0x1b, 0xed, 0xb5, 0xe5, 0x1b, 0x74, 0x5f, 0x49, 0x86, 0xd0, 0x78, 0xcb, 0xd7,
	0xe4, 0x56, 0x32, 0xb3, 0x43, 0xdd, 0xac, 0xd0, 0xa4, 0x6f, 0xd0, 0xa2, 0x4a, 0xfb, 0x58, 0xa7,
	0x7f, 0xc1, 0xde, 0xbe, 0x38, 0xb1, 0xa0, 0xfc, 0x70, 0xfb, 0xdc, 0x34, 0x06, 0x95, 0x17, 0x33,
	0x9e, 0x4e, 0xab, 0xd9, 0xef, 0xee, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0x92, 0x55, 0x89, 0x20,
	0xf1, 0x02, 0x00, 0x00,
}
