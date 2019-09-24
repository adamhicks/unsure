// Code generated by protoc-gen-go. DO NOT EDIT.
// source: engine.proto

package enginepb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import reflexpb "github.com/luno/reflex/reflexpb"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_engine_1e68ddec9786c807, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type StartMatchReq struct {
	Team                 string   `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
	Players              int64    `protobuf:"varint,2,opt,name=players,proto3" json:"players,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartMatchReq) Reset()         { *m = StartMatchReq{} }
func (m *StartMatchReq) String() string { return proto.CompactTextString(m) }
func (*StartMatchReq) ProtoMessage()    {}
func (*StartMatchReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_engine_1e68ddec9786c807, []int{1}
}
func (m *StartMatchReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartMatchReq.Unmarshal(m, b)
}
func (m *StartMatchReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartMatchReq.Marshal(b, m, deterministic)
}
func (dst *StartMatchReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartMatchReq.Merge(dst, src)
}
func (m *StartMatchReq) XXX_Size() int {
	return xxx_messageInfo_StartMatchReq.Size(m)
}
func (m *StartMatchReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StartMatchReq.DiscardUnknown(m)
}

var xxx_messageInfo_StartMatchReq proto.InternalMessageInfo

func (m *StartMatchReq) GetTeam() string {
	if m != nil {
		return m.Team
	}
	return ""
}

func (m *StartMatchReq) GetPlayers() int64 {
	if m != nil {
		return m.Players
	}
	return 0
}

type JoinRoundReq struct {
	Team                 string   `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
	Player               string   `protobuf:"bytes,2,opt,name=player,proto3" json:"player,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinRoundReq) Reset()         { *m = JoinRoundReq{} }
func (m *JoinRoundReq) String() string { return proto.CompactTextString(m) }
func (*JoinRoundReq) ProtoMessage()    {}
func (*JoinRoundReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_engine_1e68ddec9786c807, []int{2}
}
func (m *JoinRoundReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRoundReq.Unmarshal(m, b)
}
func (m *JoinRoundReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRoundReq.Marshal(b, m, deterministic)
}
func (dst *JoinRoundReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRoundReq.Merge(dst, src)
}
func (m *JoinRoundReq) XXX_Size() int {
	return xxx_messageInfo_JoinRoundReq.Size(m)
}
func (m *JoinRoundReq) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRoundReq.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRoundReq proto.InternalMessageInfo

func (m *JoinRoundReq) GetTeam() string {
	if m != nil {
		return m.Team
	}
	return ""
}

func (m *JoinRoundReq) GetPlayer() string {
	if m != nil {
		return m.Player
	}
	return ""
}

type JoinRoundRes struct {
	Included             bool     `protobuf:"varint,1,opt,name=included,proto3" json:"included,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinRoundRes) Reset()         { *m = JoinRoundRes{} }
func (m *JoinRoundRes) String() string { return proto.CompactTextString(m) }
func (*JoinRoundRes) ProtoMessage()    {}
func (*JoinRoundRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_engine_1e68ddec9786c807, []int{3}
}
func (m *JoinRoundRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRoundRes.Unmarshal(m, b)
}
func (m *JoinRoundRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRoundRes.Marshal(b, m, deterministic)
}
func (dst *JoinRoundRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRoundRes.Merge(dst, src)
}
func (m *JoinRoundRes) XXX_Size() int {
	return xxx_messageInfo_JoinRoundRes.Size(m)
}
func (m *JoinRoundRes) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRoundRes.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRoundRes proto.InternalMessageInfo

func (m *JoinRoundRes) GetIncluded() bool {
	if m != nil {
		return m.Included
	}
	return false
}

type CollectRoundReq struct {
	Team                 string   `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
	Player               string   `protobuf:"bytes,2,opt,name=player,proto3" json:"player,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectRoundReq) Reset()         { *m = CollectRoundReq{} }
func (m *CollectRoundReq) String() string { return proto.CompactTextString(m) }
func (*CollectRoundReq) ProtoMessage()    {}
func (*CollectRoundReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_engine_1e68ddec9786c807, []int{4}
}
func (m *CollectRoundReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectRoundReq.Unmarshal(m, b)
}
func (m *CollectRoundReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectRoundReq.Marshal(b, m, deterministic)
}
func (dst *CollectRoundReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectRoundReq.Merge(dst, src)
}
func (m *CollectRoundReq) XXX_Size() int {
	return xxx_messageInfo_CollectRoundReq.Size(m)
}
func (m *CollectRoundReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectRoundReq.DiscardUnknown(m)
}

var xxx_messageInfo_CollectRoundReq proto.InternalMessageInfo

func (m *CollectRoundReq) GetTeam() string {
	if m != nil {
		return m.Team
	}
	return ""
}

func (m *CollectRoundReq) GetPlayer() string {
	if m != nil {
		return m.Player
	}
	return ""
}

type CollectRoundRes struct {
	Players              []*CollectPlayer `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CollectRoundRes) Reset()         { *m = CollectRoundRes{} }
func (m *CollectRoundRes) String() string { return proto.CompactTextString(m) }
func (*CollectRoundRes) ProtoMessage()    {}
func (*CollectRoundRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_engine_1e68ddec9786c807, []int{5}
}
func (m *CollectRoundRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectRoundRes.Unmarshal(m, b)
}
func (m *CollectRoundRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectRoundRes.Marshal(b, m, deterministic)
}
func (dst *CollectRoundRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectRoundRes.Merge(dst, src)
}
func (m *CollectRoundRes) XXX_Size() int {
	return xxx_messageInfo_CollectRoundRes.Size(m)
}
func (m *CollectRoundRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectRoundRes.DiscardUnknown(m)
}

var xxx_messageInfo_CollectRoundRes proto.InternalMessageInfo

func (m *CollectRoundRes) GetPlayers() []*CollectPlayer {
	if m != nil {
		return m.Players
	}
	return nil
}

type CollectPlayer struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Rank                 int64    `protobuf:"varint,2,opt,name=rank,proto3" json:"rank,omitempty"`
	Part                 int64    `protobuf:"varint,3,opt,name=part,proto3" json:"part,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectPlayer) Reset()         { *m = CollectPlayer{} }
func (m *CollectPlayer) String() string { return proto.CompactTextString(m) }
func (*CollectPlayer) ProtoMessage()    {}
func (*CollectPlayer) Descriptor() ([]byte, []int) {
	return fileDescriptor_engine_1e68ddec9786c807, []int{6}
}
func (m *CollectPlayer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectPlayer.Unmarshal(m, b)
}
func (m *CollectPlayer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectPlayer.Marshal(b, m, deterministic)
}
func (dst *CollectPlayer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectPlayer.Merge(dst, src)
}
func (m *CollectPlayer) XXX_Size() int {
	return xxx_messageInfo_CollectPlayer.Size(m)
}
func (m *CollectPlayer) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectPlayer.DiscardUnknown(m)
}

var xxx_messageInfo_CollectPlayer proto.InternalMessageInfo

func (m *CollectPlayer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CollectPlayer) GetRank() int64 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *CollectPlayer) GetPart() int64 {
	if m != nil {
		return m.Part
	}
	return 0
}

type SubmitRoundReq struct {
	Team                 string   `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
	Player               string   `protobuf:"bytes,2,opt,name=player,proto3" json:"player,omitempty"`
	Total                int64    `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubmitRoundReq) Reset()         { *m = SubmitRoundReq{} }
func (m *SubmitRoundReq) String() string { return proto.CompactTextString(m) }
func (*SubmitRoundReq) ProtoMessage()    {}
func (*SubmitRoundReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_engine_1e68ddec9786c807, []int{7}
}
func (m *SubmitRoundReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitRoundReq.Unmarshal(m, b)
}
func (m *SubmitRoundReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitRoundReq.Marshal(b, m, deterministic)
}
func (dst *SubmitRoundReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitRoundReq.Merge(dst, src)
}
func (m *SubmitRoundReq) XXX_Size() int {
	return xxx_messageInfo_SubmitRoundReq.Size(m)
}
func (m *SubmitRoundReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitRoundReq.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitRoundReq proto.InternalMessageInfo

func (m *SubmitRoundReq) GetTeam() string {
	if m != nil {
		return m.Team
	}
	return ""
}

func (m *SubmitRoundReq) GetPlayer() string {
	if m != nil {
		return m.Player
	}
	return ""
}

func (m *SubmitRoundReq) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "enginepb.Empty")
	proto.RegisterType((*StartMatchReq)(nil), "enginepb.StartMatchReq")
	proto.RegisterType((*JoinRoundReq)(nil), "enginepb.JoinRoundReq")
	proto.RegisterType((*JoinRoundRes)(nil), "enginepb.JoinRoundRes")
	proto.RegisterType((*CollectRoundReq)(nil), "enginepb.CollectRoundReq")
	proto.RegisterType((*CollectRoundRes)(nil), "enginepb.CollectRoundRes")
	proto.RegisterType((*CollectPlayer)(nil), "enginepb.CollectPlayer")
	proto.RegisterType((*SubmitRoundReq)(nil), "enginepb.SubmitRoundReq")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EngineClient is the client API for Engine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EngineClient interface {
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	Stream(ctx context.Context, in *reflexpb.StreamRequest, opts ...grpc.CallOption) (Engine_StreamClient, error)
	StartMatch(ctx context.Context, in *StartMatchReq, opts ...grpc.CallOption) (*Empty, error)
	JoinRound(ctx context.Context, in *JoinRoundReq, opts ...grpc.CallOption) (*JoinRoundRes, error)
	CollectRound(ctx context.Context, in *CollectRoundReq, opts ...grpc.CallOption) (*CollectRoundRes, error)
	SubmitRound(ctx context.Context, in *SubmitRoundReq, opts ...grpc.CallOption) (*Empty, error)
}

type engineClient struct {
	cc *grpc.ClientConn
}

func NewEngineClient(cc *grpc.ClientConn) EngineClient {
	return &engineClient{cc}
}

func (c *engineClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/enginepb.Engine/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) Stream(ctx context.Context, in *reflexpb.StreamRequest, opts ...grpc.CallOption) (Engine_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Engine_serviceDesc.Streams[0], "/enginepb.Engine/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &engineStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Engine_StreamClient interface {
	Recv() (*reflexpb.Event, error)
	grpc.ClientStream
}

type engineStreamClient struct {
	grpc.ClientStream
}

func (x *engineStreamClient) Recv() (*reflexpb.Event, error) {
	m := new(reflexpb.Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *engineClient) StartMatch(ctx context.Context, in *StartMatchReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/enginepb.Engine/StartMatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) JoinRound(ctx context.Context, in *JoinRoundReq, opts ...grpc.CallOption) (*JoinRoundRes, error) {
	out := new(JoinRoundRes)
	err := c.cc.Invoke(ctx, "/enginepb.Engine/JoinRound", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) CollectRound(ctx context.Context, in *CollectRoundReq, opts ...grpc.CallOption) (*CollectRoundRes, error) {
	out := new(CollectRoundRes)
	err := c.cc.Invoke(ctx, "/enginepb.Engine/CollectRound", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) SubmitRound(ctx context.Context, in *SubmitRoundReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/enginepb.Engine/SubmitRound", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EngineServer is the server API for Engine service.
type EngineServer interface {
	Ping(context.Context, *Empty) (*Empty, error)
	Stream(*reflexpb.StreamRequest, Engine_StreamServer) error
	StartMatch(context.Context, *StartMatchReq) (*Empty, error)
	JoinRound(context.Context, *JoinRoundReq) (*JoinRoundRes, error)
	CollectRound(context.Context, *CollectRoundReq) (*CollectRoundRes, error)
	SubmitRound(context.Context, *SubmitRoundReq) (*Empty, error)
}

func RegisterEngineServer(s *grpc.Server, srv EngineServer) {
	s.RegisterService(&_Engine_serviceDesc, srv)
}

func _Engine_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/enginepb.Engine/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(reflexpb.StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EngineServer).Stream(m, &engineStreamServer{stream})
}

type Engine_StreamServer interface {
	Send(*reflexpb.Event) error
	grpc.ServerStream
}

type engineStreamServer struct {
	grpc.ServerStream
}

func (x *engineStreamServer) Send(m *reflexpb.Event) error {
	return x.ServerStream.SendMsg(m)
}

func _Engine_StartMatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartMatchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).StartMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/enginepb.Engine/StartMatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).StartMatch(ctx, req.(*StartMatchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_JoinRound_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRoundReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).JoinRound(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/enginepb.Engine/JoinRound",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).JoinRound(ctx, req.(*JoinRoundReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_CollectRound_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectRoundReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).CollectRound(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/enginepb.Engine/CollectRound",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).CollectRound(ctx, req.(*CollectRoundReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_SubmitRound_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitRoundReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).SubmitRound(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/enginepb.Engine/SubmitRound",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).SubmitRound(ctx, req.(*SubmitRoundReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Engine_serviceDesc = grpc.ServiceDesc{
	ServiceName: "enginepb.Engine",
	HandlerType: (*EngineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Engine_Ping_Handler,
		},
		{
			MethodName: "StartMatch",
			Handler:    _Engine_StartMatch_Handler,
		},
		{
			MethodName: "JoinRound",
			Handler:    _Engine_JoinRound_Handler,
		},
		{
			MethodName: "CollectRound",
			Handler:    _Engine_CollectRound_Handler,
		},
		{
			MethodName: "SubmitRound",
			Handler:    _Engine_SubmitRound_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Engine_Stream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "engine.proto",
}

func init() { proto.RegisterFile("engine.proto", fileDescriptor_engine_1e68ddec9786c807) }

var fileDescriptor_engine_1e68ddec9786c807 = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x41, 0xcf, 0x9a, 0x40,
	0x10, 0x95, 0xcf, 0xef, 0x43, 0x1d, 0xb5, 0x26, 0x93, 0x46, 0x29, 0x27, 0xb3, 0x27, 0x63, 0x1a,
	0x6c, 0x6d, 0xd2, 0x34, 0x26, 0x9e, 0x5a, 0x7b, 0x68, 0xd3, 0xc4, 0xac, 0xbf, 0x60, 0xc1, 0xad,
	0x92, 0xc2, 0x2e, 0xc2, 0xd2, 0xd4, 0x9f, 0xdc, 0x7f, 0xd1, 0xb0, 0x80, 0x80, 0xd5, 0x8b, 0x27,
	0x66, 0xde, 0xec, 0xdb, 0xd9, 0x37, 0x6f, 0x80, 0x01, 0x17, 0x07, 0x5f, 0x70, 0x27, 0x8a, 0xa5,
	0x92, 0xd8, 0xcd, 0xb3, 0xc8, 0xb5, 0xdf, 0x1e, 0x7c, 0x75, 0x4c, 0x5d, 0xc7, 0x93, 0xe1, 0x22,
	0x48, 0x85, 0x5c, 0xc4, 0xfc, 0x67, 0xc0, 0xff, 0x14, 0x9f, 0xc8, 0x2d, 0x82, 0x9c, 0x47, 0x3a,
	0xf0, 0xb2, 0x09, 0x23, 0x75, 0x26, 0x6b, 0x18, 0xee, 0x14, 0x8b, 0xd5, 0x0f, 0xa6, 0xbc, 0x23,
	0xe5, 0x27, 0x44, 0x78, 0x56, 0x9c, 0x85, 0x96, 0x31, 0x35, 0x66, 0x3d, 0xaa, 0x63, 0xb4, 0xa0,
	0x13, 0x05, 0xec, 0xcc, 0xe3, 0xc4, 0x7a, 0x9a, 0x1a, 0xb3, 0x36, 0x2d, 0x53, 0xb2, 0x82, 0xc1,
	0x37, 0xe9, 0x0b, 0x2a, 0x53, 0xb1, 0xbf, 0xc7, 0x1e, 0x83, 0x99, 0x1f, 0xd7, 0xe4, 0x1e, 0x2d,
	0x32, 0x32, 0x6f, 0x70, 0x13, 0xb4, 0xa1, 0xeb, 0x0b, 0x2f, 0x48, 0xf7, 0x7c, 0xaf, 0xf9, 0x5d,
	0x7a, 0xc9, 0xc9, 0x1a, 0x46, 0x9f, 0x65, 0x10, 0x70, 0x4f, 0x3d, 0xd4, 0xea, 0xcb, 0x35, 0x3d,
	0xc1, 0xf7, 0x95, 0x26, 0x63, 0xda, 0x9e, 0xf5, 0x97, 0x13, 0xa7, 0x9c, 0xa5, 0x53, 0x9c, 0xdd,
	0xea, 0x7a, 0x25, 0xf6, 0x3b, 0x0c, 0x1b, 0x95, 0xec, 0x09, 0x82, 0x85, 0xbc, 0x7c, 0x42, 0x16,
	0x67, 0x58, 0xcc, 0xc4, 0xaf, 0x62, 0x50, 0x3a, 0xce, 0xb0, 0x88, 0xc5, 0xca, 0x6a, 0xe7, 0x58,
	0x16, 0x13, 0x0a, 0xaf, 0x76, 0xa9, 0x1b, 0xfa, 0x0f, 0x09, 0xc2, 0xd7, 0xf0, 0xa2, 0xa4, 0x62,
	0x41, 0x71, 0x65, 0x9e, 0x2c, 0xff, 0x3e, 0x81, 0xb9, 0xd1, 0x22, 0x70, 0x0e, 0xcf, 0x5b, 0x5f,
	0x1c, 0x70, 0x54, 0xa9, 0xd2, 0x86, 0xdb, 0xd7, 0x00, 0x69, 0xe1, 0x47, 0x30, 0x77, 0x2a, 0xce,
	0xda, 0x4d, 0x9c, 0x72, 0x5d, 0x9c, 0x1c, 0xa1, 0xfc, 0x94, 0xf2, 0x44, 0xd9, 0xa3, 0xaa, 0xb0,
	0xf9, 0xcd, 0x85, 0x22, 0xad, 0x77, 0x06, 0x7e, 0x02, 0xa8, 0x76, 0x07, 0x6b, 0xf3, 0x6b, 0x6c,
	0xd4, 0xad, 0x8e, 0x6b, 0xe8, 0x5d, 0xac, 0xc7, 0x71, 0x55, 0xaf, 0xef, 0x92, 0x7d, 0x1b, 0x4f,
	0x48, 0x0b, 0xbf, 0xc2, 0xa0, 0x6e, 0x27, 0xbe, 0xf9, 0xcf, 0xba, 0xcb, 0x25, 0x77, 0x4b, 0xd9,
	0x3d, 0x2b, 0xe8, 0xd7, 0x3c, 0x40, 0xab, 0xa6, 0xa0, 0x61, 0xcd, 0x0d, 0x09, 0xae, 0xa9, 0x7f,
	0xa4, 0x0f, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd4, 0xc7, 0x1d, 0xb8, 0x90, 0x03, 0x00, 0x00,
}
