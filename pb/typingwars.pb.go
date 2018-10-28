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
	return fileDescriptor_1b666822ed8f2be0, []int{0}
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
	return fileDescriptor_1b666822ed8f2be0, []int{1}
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

type CreateRoomRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRoomRequest) Reset()         { *m = CreateRoomRequest{} }
func (m *CreateRoomRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRoomRequest) ProtoMessage()    {}
func (*CreateRoomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b666822ed8f2be0, []int{2}
}

func (m *CreateRoomRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRoomRequest.Unmarshal(m, b)
}
func (m *CreateRoomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRoomRequest.Marshal(b, m, deterministic)
}
func (m *CreateRoomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoomRequest.Merge(m, src)
}
func (m *CreateRoomRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRoomRequest.Size(m)
}
func (m *CreateRoomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoomRequest proto.InternalMessageInfo

func (m *CreateRoomRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type JoinRoomRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	RoomId               string   `protobuf:"bytes,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinRoomRequest) Reset()         { *m = JoinRoomRequest{} }
func (m *JoinRoomRequest) String() string { return proto.CompactTextString(m) }
func (*JoinRoomRequest) ProtoMessage()    {}
func (*JoinRoomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b666822ed8f2be0, []int{3}
}

func (m *JoinRoomRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRoomRequest.Unmarshal(m, b)
}
func (m *JoinRoomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRoomRequest.Marshal(b, m, deterministic)
}
func (m *JoinRoomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRoomRequest.Merge(m, src)
}
func (m *JoinRoomRequest) XXX_Size() int {
	return xxx_messageInfo_JoinRoomRequest.Size(m)
}
func (m *JoinRoomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRoomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRoomRequest proto.InternalMessageInfo

func (m *JoinRoomRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *JoinRoomRequest) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

type Player struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b666822ed8f2be0, []int{4}
}

func (m *Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Player.Unmarshal(m, b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Player.Marshal(b, m, deterministic)
}
func (m *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(m, src)
}
func (m *Player) XXX_Size() int {
	return xxx_messageInfo_Player.Size(m)
}
func (m *Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Player proto.InternalMessageInfo

func (m *Player) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Player) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type JoinRoomAck struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinRoomAck) Reset()         { *m = JoinRoomAck{} }
func (m *JoinRoomAck) String() string { return proto.CompactTextString(m) }
func (*JoinRoomAck) ProtoMessage()    {}
func (*JoinRoomAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b666822ed8f2be0, []int{5}
}

func (m *JoinRoomAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRoomAck.Unmarshal(m, b)
}
func (m *JoinRoomAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRoomAck.Marshal(b, m, deterministic)
}
func (m *JoinRoomAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRoomAck.Merge(m, src)
}
func (m *JoinRoomAck) XXX_Size() int {
	return xxx_messageInfo_JoinRoomAck.Size(m)
}
func (m *JoinRoomAck) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRoomAck.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRoomAck proto.InternalMessageInfo

func (m *JoinRoomAck) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type UpdateRoom struct {
	RoomId               string             `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Players              map[string]*Player `protobuf:"bytes,2,rep,name=players,proto3" json:"players,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ReadyStatus          map[string]bool    `protobuf:"bytes,3,rep,name=readyStatus,proto3" json:"readyStatus,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *UpdateRoom) Reset()         { *m = UpdateRoom{} }
func (m *UpdateRoom) String() string { return proto.CompactTextString(m) }
func (*UpdateRoom) ProtoMessage()    {}
func (*UpdateRoom) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b666822ed8f2be0, []int{6}
}

func (m *UpdateRoom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRoom.Unmarshal(m, b)
}
func (m *UpdateRoom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRoom.Marshal(b, m, deterministic)
}
func (m *UpdateRoom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRoom.Merge(m, src)
}
func (m *UpdateRoom) XXX_Size() int {
	return xxx_messageInfo_UpdateRoom.Size(m)
}
func (m *UpdateRoom) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRoom.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRoom proto.InternalMessageInfo

func (m *UpdateRoom) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

func (m *UpdateRoom) GetPlayers() map[string]*Player {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *UpdateRoom) GetReadyStatus() map[string]bool {
	if m != nil {
		return m.ReadyStatus
	}
	return nil
}

type UpdatePlayerReady struct {
	ReadyStatus          bool     `protobuf:"varint,2,opt,name=ready_status,json=readyStatus,proto3" json:"ready_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePlayerReady) Reset()         { *m = UpdatePlayerReady{} }
func (m *UpdatePlayerReady) String() string { return proto.CompactTextString(m) }
func (*UpdatePlayerReady) ProtoMessage()    {}
func (*UpdatePlayerReady) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b666822ed8f2be0, []int{7}
}

func (m *UpdatePlayerReady) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePlayerReady.Unmarshal(m, b)
}
func (m *UpdatePlayerReady) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePlayerReady.Marshal(b, m, deterministic)
}
func (m *UpdatePlayerReady) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePlayerReady.Merge(m, src)
}
func (m *UpdatePlayerReady) XXX_Size() int {
	return xxx_messageInfo_UpdatePlayerReady.Size(m)
}
func (m *UpdatePlayerReady) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePlayerReady.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePlayerReady proto.InternalMessageInfo

func (m *UpdatePlayerReady) GetReadyStatus() bool {
	if m != nil {
		return m.ReadyStatus
	}
	return false
}

type UserMessage struct {
	// Types that are valid to be assigned to Content:
	//	*UserMessage_UserAction
	//	*UserMessage_RegisterPlayer
	//	*UserMessage_CreateRoomRequest
	//	*UserMessage_JoinRoomRequest
	//	*UserMessage_JoinRoomAck
	//	*UserMessage_UpdateRoom
	//	*UserMessage_UpdatePlayerReady
	Content              isUserMessage_Content `protobuf_oneof:"content"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UserMessage) Reset()         { *m = UserMessage{} }
func (m *UserMessage) String() string { return proto.CompactTextString(m) }
func (*UserMessage) ProtoMessage()    {}
func (*UserMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b666822ed8f2be0, []int{8}
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

type UserMessage_CreateRoomRequest struct {
	CreateRoomRequest *CreateRoomRequest `protobuf:"bytes,3,opt,name=createRoomRequest,proto3,oneof"`
}

type UserMessage_JoinRoomRequest struct {
	JoinRoomRequest *JoinRoomRequest `protobuf:"bytes,4,opt,name=joinRoomRequest,proto3,oneof"`
}

type UserMessage_JoinRoomAck struct {
	JoinRoomAck *JoinRoomAck `protobuf:"bytes,5,opt,name=joinRoomAck,proto3,oneof"`
}

type UserMessage_UpdateRoom struct {
	UpdateRoom *UpdateRoom `protobuf:"bytes,6,opt,name=updateRoom,proto3,oneof"`
}

type UserMessage_UpdatePlayerReady struct {
	UpdatePlayerReady *UpdatePlayerReady `protobuf:"bytes,7,opt,name=updatePlayerReady,proto3,oneof"`
}

func (*UserMessage_UserAction) isUserMessage_Content() {}

func (*UserMessage_RegisterPlayer) isUserMessage_Content() {}

func (*UserMessage_CreateRoomRequest) isUserMessage_Content() {}

func (*UserMessage_JoinRoomRequest) isUserMessage_Content() {}

func (*UserMessage_JoinRoomAck) isUserMessage_Content() {}

func (*UserMessage_UpdateRoom) isUserMessage_Content() {}

func (*UserMessage_UpdatePlayerReady) isUserMessage_Content() {}

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

func (m *UserMessage) GetCreateRoomRequest() *CreateRoomRequest {
	if x, ok := m.GetContent().(*UserMessage_CreateRoomRequest); ok {
		return x.CreateRoomRequest
	}
	return nil
}

func (m *UserMessage) GetJoinRoomRequest() *JoinRoomRequest {
	if x, ok := m.GetContent().(*UserMessage_JoinRoomRequest); ok {
		return x.JoinRoomRequest
	}
	return nil
}

func (m *UserMessage) GetJoinRoomAck() *JoinRoomAck {
	if x, ok := m.GetContent().(*UserMessage_JoinRoomAck); ok {
		return x.JoinRoomAck
	}
	return nil
}

func (m *UserMessage) GetUpdateRoom() *UpdateRoom {
	if x, ok := m.GetContent().(*UserMessage_UpdateRoom); ok {
		return x.UpdateRoom
	}
	return nil
}

func (m *UserMessage) GetUpdatePlayerReady() *UpdatePlayerReady {
	if x, ok := m.GetContent().(*UserMessage_UpdatePlayerReady); ok {
		return x.UpdatePlayerReady
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*UserMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _UserMessage_OneofMarshaler, _UserMessage_OneofUnmarshaler, _UserMessage_OneofSizer, []interface{}{
		(*UserMessage_UserAction)(nil),
		(*UserMessage_RegisterPlayer)(nil),
		(*UserMessage_CreateRoomRequest)(nil),
		(*UserMessage_JoinRoomRequest)(nil),
		(*UserMessage_JoinRoomAck)(nil),
		(*UserMessage_UpdateRoom)(nil),
		(*UserMessage_UpdatePlayerReady)(nil),
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
	case *UserMessage_CreateRoomRequest:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CreateRoomRequest); err != nil {
			return err
		}
	case *UserMessage_JoinRoomRequest:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.JoinRoomRequest); err != nil {
			return err
		}
	case *UserMessage_JoinRoomAck:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.JoinRoomAck); err != nil {
			return err
		}
	case *UserMessage_UpdateRoom:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UpdateRoom); err != nil {
			return err
		}
	case *UserMessage_UpdatePlayerReady:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UpdatePlayerReady); err != nil {
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
	case 3: // content.createRoomRequest
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CreateRoomRequest)
		err := b.DecodeMessage(msg)
		m.Content = &UserMessage_CreateRoomRequest{msg}
		return true, err
	case 4: // content.joinRoomRequest
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(JoinRoomRequest)
		err := b.DecodeMessage(msg)
		m.Content = &UserMessage_JoinRoomRequest{msg}
		return true, err
	case 5: // content.joinRoomAck
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(JoinRoomAck)
		err := b.DecodeMessage(msg)
		m.Content = &UserMessage_JoinRoomAck{msg}
		return true, err
	case 6: // content.updateRoom
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UpdateRoom)
		err := b.DecodeMessage(msg)
		m.Content = &UserMessage_UpdateRoom{msg}
		return true, err
	case 7: // content.updatePlayerReady
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UpdatePlayerReady)
		err := b.DecodeMessage(msg)
		m.Content = &UserMessage_UpdatePlayerReady{msg}
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
	case *UserMessage_CreateRoomRequest:
		s := proto.Size(x.CreateRoomRequest)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UserMessage_JoinRoomRequest:
		s := proto.Size(x.JoinRoomRequest)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UserMessage_JoinRoomAck:
		s := proto.Size(x.JoinRoomAck)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UserMessage_UpdateRoom:
		s := proto.Size(x.UpdateRoom)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UserMessage_UpdatePlayerReady:
		s := proto.Size(x.UpdatePlayerReady)
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
	proto.RegisterType((*UserAction)(nil), "typingwars.UserAction")
	proto.RegisterType((*RegisterPlayer)(nil), "typingwars.RegisterPlayer")
	proto.RegisterType((*CreateRoomRequest)(nil), "typingwars.CreateRoomRequest")
	proto.RegisterType((*JoinRoomRequest)(nil), "typingwars.JoinRoomRequest")
	proto.RegisterType((*Player)(nil), "typingwars.Player")
	proto.RegisterType((*JoinRoomAck)(nil), "typingwars.JoinRoomAck")
	proto.RegisterType((*UpdateRoom)(nil), "typingwars.UpdateRoom")
	proto.RegisterMapType((map[string]*Player)(nil), "typingwars.UpdateRoom.PlayersEntry")
	proto.RegisterMapType((map[string]bool)(nil), "typingwars.UpdateRoom.ReadyStatusEntry")
	proto.RegisterType((*UpdatePlayerReady)(nil), "typingwars.UpdatePlayerReady")
	proto.RegisterType((*UserMessage)(nil), "typingwars.UserMessage")
}

func init() { proto.RegisterFile("typingwars.proto", fileDescriptor_1b666822ed8f2be0) }

var fileDescriptor_1b666822ed8f2be0 = []byte{
	// 543 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdf, 0x8b, 0xd3, 0x40,
	0x10, 0xc7, 0xd3, 0xe4, 0xae, 0x3f, 0x26, 0x47, 0xaf, 0x5d, 0x4e, 0x2f, 0xf4, 0x10, 0xce, 0xf8,
	0x60, 0x39, 0xa4, 0x42, 0x4f, 0xe4, 0x50, 0x14, 0x7a, 0x7a, 0xda, 0x2a, 0x15, 0x59, 0xe9, 0x83,
	0xbe, 0x94, 0x5c, 0xb2, 0x94, 0x5c, 0xdb, 0x24, 0xee, 0x6e, 0x94, 0xfc, 0x75, 0xfe, 0x57, 0x3e,
	0x4b, 0x76, 0xd3, 0x64, 0x93, 0x54, 0xf1, 0x2d, 0xc9, 0xcc, 0xf7, 0x33, 0xb3, 0x93, 0xef, 0x2c,
	0xf4, 0x78, 0x12, 0xf9, 0xc1, 0xea, 0xa7, 0x43, 0xd9, 0x28, 0xa2, 0x21, 0x0f, 0x11, 0x14, 0x5f,
	0xec, 0x09, 0xc0, 0x82, 0x11, 0x3a, 0x71, 0xb9, 0x1f, 0x06, 0xe8, 0x12, 0x3a, 0x31, 0x23, 0x74,
	0x16, 0x44, 0x31, 0xb7, 0x1a, 0xe7, 0x8d, 0x61, 0x77, 0x7c, 0x6f, 0xa4, 0xe8, 0x17, 0xbb, 0x20,
	0x2e, 0xf2, 0xec, 0x27, 0xd0, 0xc5, 0x64, 0xe5, 0x33, 0x4e, 0xe8, 0xe7, 0x8d, 0x93, 0x10, 0x8a,
	0x06, 0xd0, 0x4e, 0xc3, 0x81, 0xb3, 0x25, 0x82, 0xd2, 0xc1, 0xf9, 0xbb, 0xfd, 0x14, 0xfa, 0x6f,
	0x28, 0x71, 0x38, 0xc1, 0x61, 0xb8, 0xc5, 0xe4, 0x7b, 0x4c, 0x18, 0xff, 0xa7, 0xe0, 0x1d, 0x1c,
	0x7f, 0x08, 0xfd, 0xe0, 0x3f, 0xd3, 0xd1, 0x29, 0xb4, 0x68, 0x18, 0x6e, 0x97, 0xbe, 0x67, 0xe9,
	0x22, 0xd4, 0x4c, 0x5f, 0x67, 0x9e, 0xfd, 0x0c, 0x9a, 0x59, 0x7b, 0x5d, 0xd0, 0x7d, 0x2f, 0x13,
	0xea, 0xbe, 0x57, 0xc2, 0xe9, 0x95, 0xea, 0x17, 0x60, 0xee, 0xaa, 0x4f, 0xdc, 0x35, 0x3a, 0x83,
	0x8e, 0xbb, 0xf1, 0x49, 0xc0, 0x97, 0x39, 0xa1, 0x2d, 0x3f, 0xcc, 0x3c, 0xfb, 0x97, 0x0e, 0xb0,
	0x88, 0xbc, 0xec, 0x6c, 0x6a, 0x27, 0x0d, 0xb5, 0x13, 0xf4, 0x0a, 0x5a, 0x91, 0xe8, 0x84, 0x59,
	0xfa, 0xb9, 0x31, 0x34, 0xc7, 0x8f, 0x4a, 0x33, 0xce, 0x09, 0x23, 0xd9, 0x2f, 0xbb, 0x09, 0x38,
	0x4d, 0xf0, 0x4e, 0x83, 0x66, 0x60, 0x52, 0xe2, 0x78, 0xc9, 0x17, 0xee, 0xf0, 0x98, 0x59, 0x86,
	0x40, 0x3c, 0xfe, 0x0b, 0x02, 0x17, 0x99, 0x12, 0xa3, 0x6a, 0x07, 0x9f, 0xe0, 0x48, 0xad, 0x81,
	0x7a, 0x60, 0xac, 0x49, 0x92, 0xb5, 0x9b, 0x3e, 0xa2, 0x21, 0x1c, 0xfe, 0x70, 0x36, 0xb1, 0x1c,
	0x8c, 0x39, 0x46, 0x6a, 0x19, 0x29, 0xc5, 0x32, 0xe1, 0x85, 0x7e, 0xd5, 0x18, 0xbc, 0x86, 0x5e,
	0xb5, 0xe0, 0x1e, 0xe6, 0x89, 0xca, 0x6c, 0x2b, 0x7a, 0xfb, 0x39, 0xf4, 0x65, 0xef, 0x19, 0x3a,
	0x65, 0xa1, 0x87, 0x70, 0x24, 0x7a, 0x5e, 0x32, 0x79, 0x60, 0xa9, 0x52, 0xcf, 0x61, 0xff, 0x36,
	0xc0, 0x4c, 0xbd, 0x39, 0x27, 0x8c, 0x39, 0x2b, 0x82, 0xae, 0x00, 0xe2, 0xdc, 0xd5, 0xa2, 0xb4,
	0x39, 0xbe, 0x5f, 0x35, 0xb2, 0x8c, 0x4e, 0x35, 0xac, 0xe4, 0xa2, 0xb7, 0xd0, 0xa5, 0x25, 0x33,
	0x67, 0x07, 0x1f, 0xa8, 0xea, 0xb2, 0xdd, 0xa7, 0x1a, 0xae, 0x68, 0xd0, 0x1c, 0xfa, 0x6e, 0xd5,
	0xe4, 0x96, 0x21, 0x40, 0x0f, 0x54, 0x50, 0x6d, 0x13, 0xa6, 0x1a, 0xae, 0x2b, 0xd1, 0x7b, 0x38,
	0xbe, 0x2b, 0xaf, 0x80, 0x75, 0x20, 0x60, 0x67, 0x2a, 0xac, 0xb2, 0x25, 0x53, 0x0d, 0x57, 0x55,
	0xe8, 0x25, 0x98, 0x77, 0x85, 0x9b, 0xad, 0x43, 0x01, 0x39, 0xdd, 0x07, 0x99, 0xb8, 0xeb, 0xa9,
	0x86, 0xd5, 0x6c, 0x31, 0xd4, 0xdc, 0x58, 0x56, 0x73, 0xcf, 0x50, 0xf3, 0xa8, 0x18, 0x6a, 0xb1,
	0x09, 0x73, 0xe8, 0xc7, 0xd5, 0xdf, 0x6a, 0xb5, 0xea, 0xe3, 0xa8, 0xfd, 0xfb, 0x74, 0x1c, 0x35,
	0xe5, 0x75, 0x07, 0x5a, 0x6e, 0x18, 0x70, 0x12, 0xf0, 0x8b, 0x13, 0xe8, 0xe4, 0x77, 0x12, 0x6a,
	0x81, 0xf1, 0xf1, 0xe6, 0x6b, 0x4f, 0xbb, 0x3e, 0xf8, 0xa6, 0x47, 0xb7, 0xb7, 0x4d, 0x71, 0xdb,
	0x5d, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x0e, 0x4b, 0x56, 0x01, 0x05, 0x00, 0x00,
}
