// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: datamsg.proto

/*
Package msg is a generated protocol buffer package.

It is generated from these files:
	datamsg.proto

It has these top-level messages:
	DataMsgRequest
	DataResyncRequests
	DataResyncReplies
	DataMsgReply
	DataChangeRequest
	DataChangeReply
	DataResyncRequest
	DataResyncReply
	ResyncNeededCallback
	Seq
	Error
	PingRequest
	PingReply
*/
package msg

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type PutDel int32

const (
	PutDel_PUT PutDel = 0
	PutDel_DEL PutDel = 1
)

var PutDel_name = map[int32]string{
	0: "PUT",
	1: "DEL",
}
var PutDel_value = map[string]int32{
	"PUT": 0,
	"DEL": 1,
}

func (x PutDel) String() string {
	return proto.EnumName(PutDel_name, int32(x))
}
func (PutDel) EnumDescriptor() ([]byte, []int) { return fileDescriptorDatamsg, []int{0} }

type DataMsgRequest struct {
	MsgId       *Seq                 `protobuf:"bytes,1,opt,name=msgId" json:"msgId,omitempty"`
	Txn         *Seq                 `protobuf:"bytes,2,opt,name=txn" json:"txn,omitempty"`
	DataChanges []*DataChangeRequest `protobuf:"bytes,3,rep,name=dataChanges" json:"dataChanges,omitempty"`
	DataResyncs []*DataResyncRequest `protobuf:"bytes,4,rep,name=dataResyncs" json:"dataResyncs,omitempty"`
}

func (m *DataMsgRequest) Reset()                    { *m = DataMsgRequest{} }
func (m *DataMsgRequest) String() string            { return proto.CompactTextString(m) }
func (*DataMsgRequest) ProtoMessage()               {}
func (*DataMsgRequest) Descriptor() ([]byte, []int) { return fileDescriptorDatamsg, []int{0} }

func (m *DataMsgRequest) GetMsgId() *Seq {
	if m != nil {
		return m.MsgId
	}
	return nil
}

func (m *DataMsgRequest) GetTxn() *Seq {
	if m != nil {
		return m.Txn
	}
	return nil
}

func (m *DataMsgRequest) GetDataChanges() []*DataChangeRequest {
	if m != nil {
		return m.DataChanges
	}
	return nil
}

func (m *DataMsgRequest) GetDataResyncs() []*DataResyncRequest {
	if m != nil {
		return m.DataResyncs
	}
	return nil
}

type DataResyncRequests struct {
	MsgId       *Seq                 `protobuf:"bytes,1,opt,name=msgId" json:"msgId,omitempty"`
	DataResyncs []*DataResyncRequest `protobuf:"bytes,2,rep,name=dataResyncs" json:"dataResyncs,omitempty"`
}

func (m *DataResyncRequests) Reset()                    { *m = DataResyncRequests{} }
func (m *DataResyncRequests) String() string            { return proto.CompactTextString(m) }
func (*DataResyncRequests) ProtoMessage()               {}
func (*DataResyncRequests) Descriptor() ([]byte, []int) { return fileDescriptorDatamsg, []int{1} }

func (m *DataResyncRequests) GetMsgId() *Seq {
	if m != nil {
		return m.MsgId
	}
	return nil
}

func (m *DataResyncRequests) GetDataResyncs() []*DataResyncRequest {
	if m != nil {
		return m.DataResyncs
	}
	return nil
}

type DataResyncReplies struct {
	MsgId       *Seq                                  `protobuf:"bytes,1,opt,name=msgId" json:"msgId,omitempty"`
	Error       *Error                                `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	DataResyncs *DataResyncReplies_DataResyncsReplies `protobuf:"bytes,3,opt,name=dataResyncs" json:"dataResyncs,omitempty"`
}

func (m *DataResyncReplies) Reset()                    { *m = DataResyncReplies{} }
func (m *DataResyncReplies) String() string            { return proto.CompactTextString(m) }
func (*DataResyncReplies) ProtoMessage()               {}
func (*DataResyncReplies) Descriptor() ([]byte, []int) { return fileDescriptorDatamsg, []int{2} }

func (m *DataResyncReplies) GetMsgId() *Seq {
	if m != nil {
		return m.MsgId
	}
	return nil
}

func (m *DataResyncReplies) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *DataResyncReplies) GetDataResyncs() *DataResyncReplies_DataResyncsReplies {
	if m != nil {
		return m.DataResyncs
	}
	return nil
}

type DataResyncReplies_DataResyncsReplies struct {
	Replies []*DataResyncReply `protobuf:"bytes,1,rep,name=replies" json:"replies,omitempty"`
}

func (m *DataResyncReplies_DataResyncsReplies) Reset()         { *m = DataResyncReplies_DataResyncsReplies{} }
func (m *DataResyncReplies_DataResyncsReplies) String() string { return proto.CompactTextString(m) }
func (*DataResyncReplies_DataResyncsReplies) ProtoMessage()    {}
func (*DataResyncReplies_DataResyncsReplies) Descriptor() ([]byte, []int) {
	return fileDescriptorDatamsg, []int{2, 0}
}

func (m *DataResyncReplies_DataResyncsReplies) GetReplies() []*DataResyncReply {
	if m != nil {
		return m.Replies
	}
	return nil
}

type DataMsgReply struct {
	MsgId       *Seq                             `protobuf:"bytes,1,opt,name=msgId" json:"msgId,omitempty"`
	Error       *Error                           `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	DataChanges *DataMsgReply_DataChangesReplies `protobuf:"bytes,3,opt,name=dataChanges" json:"dataChanges,omitempty"`
	DataResyncs *DataMsgReply_DataResyncsReplies `protobuf:"bytes,4,opt,name=dataResyncs" json:"dataResyncs,omitempty"`
}

func (m *DataMsgReply) Reset()                    { *m = DataMsgReply{} }
func (m *DataMsgReply) String() string            { return proto.CompactTextString(m) }
func (*DataMsgReply) ProtoMessage()               {}
func (*DataMsgReply) Descriptor() ([]byte, []int) { return fileDescriptorDatamsg, []int{3} }

func (m *DataMsgReply) GetMsgId() *Seq {
	if m != nil {
		return m.MsgId
	}
	return nil
}

func (m *DataMsgReply) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *DataMsgReply) GetDataChanges() *DataMsgReply_DataChangesReplies {
	if m != nil {
		return m.DataChanges
	}
	return nil
}

func (m *DataMsgReply) GetDataResyncs() *DataMsgReply_DataResyncsReplies {
	if m != nil {
		return m.DataResyncs
	}
	return nil
}

type DataMsgReply_DataChangesReplies struct {
	Replies []*DataChangeReply `protobuf:"bytes,1,rep,name=replies" json:"replies,omitempty"`
}

func (m *DataMsgReply_DataChangesReplies) Reset()         { *m = DataMsgReply_DataChangesReplies{} }
func (m *DataMsgReply_DataChangesReplies) String() string { return proto.CompactTextString(m) }
func (*DataMsgReply_DataChangesReplies) ProtoMessage()    {}
func (*DataMsgReply_DataChangesReplies) Descriptor() ([]byte, []int) {
	return fileDescriptorDatamsg, []int{3, 0}
}

func (m *DataMsgReply_DataChangesReplies) GetReplies() []*DataChangeReply {
	if m != nil {
		return m.Replies
	}
	return nil
}

type DataMsgReply_DataResyncsReplies struct {
	Replies []*DataResyncReply `protobuf:"bytes,1,rep,name=replies" json:"replies,omitempty"`
}

func (m *DataMsgReply_DataResyncsReplies) Reset()         { *m = DataMsgReply_DataResyncsReplies{} }
func (m *DataMsgReply_DataResyncsReplies) String() string { return proto.CompactTextString(m) }
func (*DataMsgReply_DataResyncsReplies) ProtoMessage()    {}
func (*DataMsgReply_DataResyncsReplies) Descriptor() ([]byte, []int) {
	return fileDescriptorDatamsg, []int{3, 1}
}

func (m *DataMsgReply_DataResyncsReplies) GetReplies() []*DataResyncReply {
	if m != nil {
		return m.Replies
	}
	return nil
}

type DataChangeRequest struct {
	Key           string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	OperationType PutDel `protobuf:"varint,2,opt,name=operationType,proto3,enum=msg.PutDel" json:"operationType,omitempty"`
	Content       []byte `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	ContentType   string `protobuf:"bytes,4,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
}

func (m *DataChangeRequest) Reset()                    { *m = DataChangeRequest{} }
func (m *DataChangeRequest) String() string            { return proto.CompactTextString(m) }
func (*DataChangeRequest) ProtoMessage()               {}
func (*DataChangeRequest) Descriptor() ([]byte, []int) { return fileDescriptorDatamsg, []int{4} }

func (m *DataChangeRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DataChangeRequest) GetOperationType() PutDel {
	if m != nil {
		return m.OperationType
	}
	return PutDel_PUT
}

func (m *DataChangeRequest) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *DataChangeRequest) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

type DataChangeReply struct {
	Key           string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	OperationType PutDel `protobuf:"varint,3,opt,name=operationType,proto3,enum=msg.PutDel" json:"operationType,omitempty"`
	// zero means success
	Result uint32 `protobuf:"varint,4,opt,name=result,proto3" json:"result,omitempty"`
}

func (m *DataChangeReply) Reset()                    { *m = DataChangeReply{} }
func (m *DataChangeReply) String() string            { return proto.CompactTextString(m) }
func (*DataChangeReply) ProtoMessage()               {}
func (*DataChangeReply) Descriptor() ([]byte, []int) { return fileDescriptorDatamsg, []int{5} }

func (m *DataChangeReply) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DataChangeReply) GetOperationType() PutDel {
	if m != nil {
		return m.OperationType
	}
	return PutDel_PUT
}

func (m *DataChangeReply) GetResult() uint32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type DataResyncRequest struct {
	Key         string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Content     []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	ContentType string `protobuf:"bytes,3,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
}

func (m *DataResyncRequest) Reset()                    { *m = DataResyncRequest{} }
func (m *DataResyncRequest) String() string            { return proto.CompactTextString(m) }
func (*DataResyncRequest) ProtoMessage()               {}
func (*DataResyncRequest) Descriptor() ([]byte, []int) { return fileDescriptorDatamsg, []int{6} }

func (m *DataResyncRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DataResyncRequest) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *DataResyncRequest) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

type DataResyncReply struct {
	ResyncId *Seq   `protobuf:"bytes,1,opt,name=resyncId" json:"resyncId,omitempty"`
	Key      string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// zero means success
	Result uint32 `protobuf:"varint,3,opt,name=result,proto3" json:"result,omitempty"`
}

func (m *DataResyncReply) Reset()                    { *m = DataResyncReply{} }
func (m *DataResyncReply) String() string            { return proto.CompactTextString(m) }
func (*DataResyncReply) ProtoMessage()               {}
func (*DataResyncReply) Descriptor() ([]byte, []int) { return fileDescriptorDatamsg, []int{7} }

func (m *DataResyncReply) GetResyncId() *Seq {
	if m != nil {
		return m.ResyncId
	}
	return nil
}

func (m *DataResyncReply) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DataResyncReply) GetResult() uint32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type ResyncNeededCallback struct {
	MsgId *Seq `protobuf:"bytes,1,opt,name=msgId" json:"msgId,omitempty"`
	// Optional
	Error *Error `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *ResyncNeededCallback) Reset()                    { *m = ResyncNeededCallback{} }
func (m *ResyncNeededCallback) String() string            { return proto.CompactTextString(m) }
func (*ResyncNeededCallback) ProtoMessage()               {}
func (*ResyncNeededCallback) Descriptor() ([]byte, []int) { return fileDescriptorDatamsg, []int{8} }

func (m *ResyncNeededCallback) GetMsgId() *Seq {
	if m != nil {
		return m.MsgId
	}
	return nil
}

func (m *ResyncNeededCallback) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type Seq struct {
	// The server that generates this seq
	Originator string `protobuf:"bytes,1,opt,name=originator,proto3" json:"originator,omitempty"`
	// UNIX Timestamp of seqence restart
	// google.protobuf.Timestamp seqCreated = 2;
	SeqCreatedSec int64 `protobuf:"varint,2,opt,name=seqCreatedSec,proto3" json:"seqCreatedSec,omitempty"`
	// Sequence
	Seq uint32 `protobuf:"varint,4,opt,name=seq,proto3" json:"seq,omitempty"`
}

func (m *Seq) Reset()                    { *m = Seq{} }
func (m *Seq) String() string            { return proto.CompactTextString(m) }
func (*Seq) ProtoMessage()               {}
func (*Seq) Descriptor() ([]byte, []int) { return fileDescriptorDatamsg, []int{9} }

func (m *Seq) GetOriginator() string {
	if m != nil {
		return m.Originator
	}
	return ""
}

func (m *Seq) GetSeqCreatedSec() int64 {
	if m != nil {
		return m.SeqCreatedSec
	}
	return 0
}

func (m *Seq) GetSeq() uint32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

type Error struct {
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptorDatamsg, []int{10} }

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type PingRequest struct {
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptorDatamsg, []int{11} }

func (m *PingRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type PingReply struct {
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *PingReply) Reset()                    { *m = PingReply{} }
func (m *PingReply) String() string            { return proto.CompactTextString(m) }
func (*PingReply) ProtoMessage()               {}
func (*PingReply) Descriptor() ([]byte, []int) { return fileDescriptorDatamsg, []int{12} }

func (m *PingReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*DataMsgRequest)(nil), "msg.DataMsgRequest")
	proto.RegisterType((*DataResyncRequests)(nil), "msg.DataResyncRequests")
	proto.RegisterType((*DataResyncReplies)(nil), "msg.DataResyncReplies")
	proto.RegisterType((*DataResyncReplies_DataResyncsReplies)(nil), "msg.DataResyncReplies.DataResyncsReplies")
	proto.RegisterType((*DataMsgReply)(nil), "msg.DataMsgReply")
	proto.RegisterType((*DataMsgReply_DataChangesReplies)(nil), "msg.DataMsgReply.DataChangesReplies")
	proto.RegisterType((*DataMsgReply_DataResyncsReplies)(nil), "msg.DataMsgReply.DataResyncsReplies")
	proto.RegisterType((*DataChangeRequest)(nil), "msg.DataChangeRequest")
	proto.RegisterType((*DataChangeReply)(nil), "msg.DataChangeReply")
	proto.RegisterType((*DataResyncRequest)(nil), "msg.DataResyncRequest")
	proto.RegisterType((*DataResyncReply)(nil), "msg.DataResyncReply")
	proto.RegisterType((*ResyncNeededCallback)(nil), "msg.ResyncNeededCallback")
	proto.RegisterType((*Seq)(nil), "msg.Seq")
	proto.RegisterType((*Error)(nil), "msg.Error")
	proto.RegisterType((*PingRequest)(nil), "msg.PingRequest")
	proto.RegisterType((*PingReply)(nil), "msg.PingReply")
	proto.RegisterEnum("msg.PutDel", PutDel_name, PutDel_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DataMsgService service

type DataMsgServiceClient interface {
	DataChanges(ctx context.Context, opts ...grpc.CallOption) (DataMsgService_DataChangesClient, error)
	DataResyncs(ctx context.Context, in *DataResyncRequests, opts ...grpc.CallOption) (*DataResyncReplies, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
}

type dataMsgServiceClient struct {
	cc *grpc.ClientConn
}

func NewDataMsgServiceClient(cc *grpc.ClientConn) DataMsgServiceClient {
	return &dataMsgServiceClient{cc}
}

func (c *dataMsgServiceClient) DataChanges(ctx context.Context, opts ...grpc.CallOption) (DataMsgService_DataChangesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DataMsgService_serviceDesc.Streams[0], c.cc, "/msg.DataMsgService/dataChanges", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataMsgServiceDataChangesClient{stream}
	return x, nil
}

type DataMsgService_DataChangesClient interface {
	Send(*DataChangeRequest) error
	Recv() (*DataChangeReply, error)
	grpc.ClientStream
}

type dataMsgServiceDataChangesClient struct {
	grpc.ClientStream
}

func (x *dataMsgServiceDataChangesClient) Send(m *DataChangeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataMsgServiceDataChangesClient) Recv() (*DataChangeReply, error) {
	m := new(DataChangeReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataMsgServiceClient) DataResyncs(ctx context.Context, in *DataResyncRequests, opts ...grpc.CallOption) (*DataResyncReplies, error) {
	out := new(DataResyncReplies)
	err := grpc.Invoke(ctx, "/msg.DataMsgService/dataResyncs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataMsgServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := grpc.Invoke(ctx, "/msg.DataMsgService/ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DataMsgService service

type DataMsgServiceServer interface {
	DataChanges(DataMsgService_DataChangesServer) error
	DataResyncs(context.Context, *DataResyncRequests) (*DataResyncReplies, error)
	Ping(context.Context, *PingRequest) (*PingReply, error)
}

func RegisterDataMsgServiceServer(s *grpc.Server, srv DataMsgServiceServer) {
	s.RegisterService(&_DataMsgService_serviceDesc, srv)
}

func _DataMsgService_DataChanges_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataMsgServiceServer).DataChanges(&dataMsgServiceDataChangesServer{stream})
}

type DataMsgService_DataChangesServer interface {
	Send(*DataChangeReply) error
	Recv() (*DataChangeRequest, error)
	grpc.ServerStream
}

type dataMsgServiceDataChangesServer struct {
	grpc.ServerStream
}

func (x *dataMsgServiceDataChangesServer) Send(m *DataChangeReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataMsgServiceDataChangesServer) Recv() (*DataChangeRequest, error) {
	m := new(DataChangeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DataMsgService_DataResyncs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataResyncRequests)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataMsgServiceServer).DataResyncs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.DataMsgService/DataResyncs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataMsgServiceServer).DataResyncs(ctx, req.(*DataResyncRequests))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataMsgService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataMsgServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.DataMsgService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataMsgServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataMsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "msg.DataMsgService",
	HandlerType: (*DataMsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "dataResyncs",
			Handler:    _DataMsgService_DataResyncs_Handler,
		},
		{
			MethodName: "ping",
			Handler:    _DataMsgService_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "dataChanges",
			Handler:       _DataMsgService_DataChanges_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "datamsg.proto",
}

func init() { proto.RegisterFile("datamsg.proto", fileDescriptorDatamsg) }

var fileDescriptorDatamsg = []byte{
	// 629 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdf, 0x6e, 0xd3, 0x3e,
	0x14, 0x5e, 0xea, 0xfd, 0x3d, 0x59, 0xf7, 0xeb, 0xcf, 0x9a, 0x46, 0x94, 0x8b, 0x69, 0x8b, 0x86,
	0x18, 0xbb, 0x98, 0xa0, 0xdc, 0x70, 0x09, 0x5a, 0x37, 0x09, 0xf1, 0x47, 0x55, 0x3a, 0x24, 0x6e,
	0x10, 0x72, 0xdb, 0xa3, 0x10, 0x2d, 0x4d, 0x52, 0xdb, 0x45, 0xf4, 0x41, 0x78, 0x1f, 0x78, 0x1b,
	0x2e, 0x78, 0x08, 0x64, 0x27, 0x69, 0xe3, 0x34, 0x6d, 0x27, 0xb1, 0x3b, 0xfb, 0xf8, 0xf3, 0xf9,
	0xce, 0xf7, 0xf9, 0xd8, 0x86, 0xe6, 0x90, 0x49, 0x36, 0x12, 0xc1, 0x65, 0xca, 0x13, 0x99, 0x50,
	0x32, 0x12, 0x81, 0xf7, 0xd3, 0x82, 0x83, 0x0e, 0x93, 0xec, 0xbd, 0x08, 0x7c, 0x1c, 0x4f, 0x50,
	0x48, 0x7a, 0x0c, 0x5b, 0x23, 0x11, 0xbc, 0x19, 0x3a, 0xd6, 0x89, 0x75, 0x6e, 0xb7, 0x77, 0x2f,
	0xd5, 0x96, 0x1e, 0x8e, 0xfd, 0x2c, 0x4c, 0x5d, 0x20, 0xf2, 0x7b, 0xec, 0x34, 0x2a, 0xab, 0x2a,
	0x48, 0x5f, 0x82, 0xad, 0x48, 0xae, 0xbe, 0xb2, 0x38, 0x40, 0xe1, 0x90, 0x13, 0x72, 0x6e, 0xb7,
	0x8f, 0x34, 0xa6, 0x33, 0x8b, 0xe7, 0x44, 0x7e, 0x19, 0x5a, 0xec, 0xf4, 0x51, 0x4c, 0xe3, 0x81,
	0x70, 0x36, 0x2b, 0x3b, 0xb3, 0xb8, 0xb1, 0x33, 0x87, 0x7a, 0x31, 0xd0, 0x05, 0x84, 0x58, 0xab,
	0xa2, 0xc2, 0xd7, 0xb8, 0x3f, 0xdf, 0x6f, 0x0b, 0xfe, 0x2f, 0x43, 0xd2, 0x28, 0xc4, 0xf5, 0x7c,
	0x27, 0xb0, 0x85, 0x9c, 0x27, 0x3c, 0xf7, 0x0d, 0xf4, 0xfa, 0xb5, 0x8a, 0xf8, 0xd9, 0x02, 0x7d,
	0x6b, 0x56, 0x44, 0x34, 0xee, 0xe9, 0x42, 0x45, 0x9a, 0xae, 0x14, 0x11, 0x79, 0xc8, 0x28, 0xd2,
	0xed, 0x94, 0x4d, 0x29, 0x20, 0xf4, 0x12, 0x76, 0x78, 0x36, 0x74, 0x2c, 0x2d, 0xf8, 0xb0, 0x26,
	0xfd, 0xd4, 0x2f, 0x40, 0xde, 0x9f, 0x06, 0xec, 0xcf, 0xba, 0x23, 0x8d, 0xa6, 0x0f, 0xa0, 0xf2,
	0xa6, 0xda, 0x21, 0x0a, 0x77, 0x36, 0x2b, 0xa3, 0x60, 0x2a, 0xb5, 0x8b, 0x29, 0xb0, 0xe8, 0x97,
	0x9b, 0x6a, 0xbf, 0xac, 0xc8, 0x73, 0x0f, 0xa3, 0x4c, 0xaa, 0x55, 0x46, 0x15, 0x3d, 0x5c, 0x36,
	0xea, 0x81, 0xec, 0xfe, 0x91, 0x77, 0x96, 0x71, 0x4d, 0x68, 0x0b, 0xc8, 0x1d, 0x4e, 0xb5, 0xe3,
	0x7b, 0xbe, 0x1a, 0xd2, 0xe7, 0xd0, 0x4c, 0x52, 0xe4, 0x4c, 0x86, 0x49, 0x7c, 0x3b, 0x4d, 0x51,
	0xbb, 0x7d, 0xd0, 0xb6, 0x75, 0xf6, 0xee, 0x44, 0x76, 0x30, 0xf2, 0x4d, 0x04, 0x75, 0x60, 0x67,
	0x90, 0xc4, 0x12, 0x63, 0xa9, 0x2d, 0xdf, 0xf7, 0x8b, 0x29, 0x3d, 0x85, 0xfd, 0x7c, 0xf8, 0x45,
	0xaa, 0x5c, 0x9b, 0x9a, 0xc7, 0xce, 0x63, 0x6a, 0xb3, 0x17, 0xc3, 0x7f, 0x15, 0xe5, 0x45, 0x51,
	0x8d, 0x15, 0x45, 0x91, 0xb5, 0x45, 0x1d, 0xc1, 0x36, 0x47, 0x31, 0x89, 0xa4, 0x26, 0x6d, 0xfa,
	0xf9, 0xcc, 0xeb, 0x9b, 0x17, 0x6c, 0x99, 0x0d, 0x25, 0x4d, 0x8d, 0xd5, 0x9a, 0xc8, 0xa2, 0x26,
	0x96, 0x69, 0x2a, 0x9d, 0x03, 0x3d, 0x83, 0x5d, 0xae, 0xa7, 0x35, 0xfd, 0x3d, 0x5b, 0xa9, 0x51,
	0x3e, 0x97, 0x41, 0x0c, 0x19, 0x9f, 0xe0, 0x30, 0x4b, 0xff, 0x01, 0x71, 0x88, 0xc3, 0x2b, 0x16,
	0x45, 0x7d, 0x36, 0xb8, 0xfb, 0xf7, 0x4b, 0xe4, 0x7d, 0x06, 0xd2, 0xc3, 0x31, 0x3d, 0x06, 0x48,
	0x78, 0x18, 0x84, 0x31, 0x93, 0x09, 0xcf, 0x9d, 0x29, 0x45, 0xe8, 0x19, 0x34, 0x05, 0x8e, 0xaf,
	0x38, 0x32, 0x89, 0xc3, 0x1e, 0x0e, 0x74, 0x42, 0xe2, 0x9b, 0x41, 0x25, 0x48, 0xe0, 0x38, 0x3f,
	0x02, 0x35, 0xf4, 0x4e, 0x61, 0x4b, 0xd3, 0x29, 0x87, 0x47, 0x28, 0x04, 0x0b, 0x30, 0xcf, 0x5e,
	0x4c, 0xbd, 0x27, 0x60, 0x77, 0xc3, 0x78, 0xf6, 0x67, 0x2c, 0x07, 0x3e, 0x86, 0xbd, 0x0c, 0xa8,
	0x1c, 0x5e, 0x0a, 0xbb, 0x70, 0x61, 0x3b, 0xeb, 0x11, 0xba, 0x03, 0xa4, 0xfb, 0xf1, 0xb6, 0xb5,
	0xa1, 0x06, 0x9d, 0xeb, 0x77, 0x2d, 0xab, 0xfd, 0x6b, 0xfe, 0x47, 0xf5, 0x90, 0x7f, 0x0b, 0x07,
	0x48, 0x5f, 0x1b, 0xaf, 0x08, 0x5d, 0xf2, 0xc3, 0xb8, 0xb5, 0xb7, 0xd6, 0xdb, 0x38, 0xb7, 0x9e,
	0x59, 0xf4, 0x95, 0xf1, 0x80, 0xd0, 0x47, 0xf5, 0x4f, 0xbf, 0x70, 0x8f, 0xea, 0x5f, 0x60, 0x6f,
	0x83, 0x5e, 0xc0, 0x66, 0x1a, 0xc6, 0x01, 0x6d, 0x65, 0x2d, 0x3e, 0xb7, 0xc3, 0x3d, 0x28, 0x45,
	0x34, 0x63, 0x7f, 0x5b, 0xff, 0xb9, 0x2f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x84, 0x25, 0x34,
	0xb8, 0x84, 0x07, 0x00, 0x00,
}