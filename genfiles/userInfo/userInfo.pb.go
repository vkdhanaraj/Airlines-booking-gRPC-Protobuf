// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v4.0.0
// source: userInfo.proto

package userInfo

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Map struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Map) Reset() {
	*x = Map{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Map) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Map) ProtoMessage() {}

func (x *Map) ProtoReflect() protoreflect.Message {
	mi := &file_userInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Map.ProtoReflect.Descriptor instead.
func (*Map) Descriptor() ([]byte, []int) {
	return file_userInfo_proto_rawDescGZIP(), []int{0}
}

func (x *Map) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Map) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userInfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_userInfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_userInfo_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_userInfo_proto protoreflect.FileDescriptor

var file_userInfo_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x2d, 0x0a, 0x03, 0x4d, 0x61,
	0x70, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x23, 0x0a, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x3e,
	0x0a, 0x04, 0x46, 0x6f, 0x72, 0x6d, 0x12, 0x36, 0x0a, 0x0e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x46, 0x6f, 0x72, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x4d, 0x61, 0x70, 0x1a, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42, 0x32,
	0x5a, 0x30, 0x41, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2d, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x2d, 0x67, 0x52, 0x50, 0x43, 0x2d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x67, 0x65, 0x6e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_userInfo_proto_rawDescOnce sync.Once
	file_userInfo_proto_rawDescData = file_userInfo_proto_rawDesc
)

func file_userInfo_proto_rawDescGZIP() []byte {
	file_userInfo_proto_rawDescOnce.Do(func() {
		file_userInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_userInfo_proto_rawDescData)
	})
	return file_userInfo_proto_rawDescData
}

var file_userInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_userInfo_proto_goTypes = []interface{}{
	(*Map)(nil),     // 0: userInfo.Map
	(*Message)(nil), // 1: userInfo.Message
}
var file_userInfo_proto_depIdxs = []int32{
	0, // 0: userInfo.Form.UploadFormData:input_type -> userInfo.Map
	1, // 1: userInfo.Form.UploadFormData:output_type -> userInfo.Message
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_userInfo_proto_init() }
func file_userInfo_proto_init() {
	if File_userInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_userInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Map); i {
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
		file_userInfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_userInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_userInfo_proto_goTypes,
		DependencyIndexes: file_userInfo_proto_depIdxs,
		MessageInfos:      file_userInfo_proto_msgTypes,
	}.Build()
	File_userInfo_proto = out.File
	file_userInfo_proto_rawDesc = nil
	file_userInfo_proto_goTypes = nil
	file_userInfo_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FormClient is the client API for Form service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FormClient interface {
	UploadFormData(ctx context.Context, opts ...grpc.CallOption) (Form_UploadFormDataClient, error)
}

type formClient struct {
	cc grpc.ClientConnInterface
}

func NewFormClient(cc grpc.ClientConnInterface) FormClient {
	return &formClient{cc}
}

func (c *formClient) UploadFormData(ctx context.Context, opts ...grpc.CallOption) (Form_UploadFormDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Form_serviceDesc.Streams[0], "/userInfo.Form/UploadFormData", opts...)
	if err != nil {
		return nil, err
	}
	x := &formUploadFormDataClient{stream}
	return x, nil
}

type Form_UploadFormDataClient interface {
	Send(*Map) error
	CloseAndRecv() (*Message, error)
	grpc.ClientStream
}

type formUploadFormDataClient struct {
	grpc.ClientStream
}

func (x *formUploadFormDataClient) Send(m *Map) error {
	return x.ClientStream.SendMsg(m)
}

func (x *formUploadFormDataClient) CloseAndRecv() (*Message, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FormServer is the server API for Form service.
type FormServer interface {
	UploadFormData(Form_UploadFormDataServer) error
}

// UnimplementedFormServer can be embedded to have forward compatible implementations.
type UnimplementedFormServer struct {
}

func (*UnimplementedFormServer) UploadFormData(Form_UploadFormDataServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFormData not implemented")
}

func RegisterFormServer(s *grpc.Server, srv FormServer) {
	s.RegisterService(&_Form_serviceDesc, srv)
}

func _Form_UploadFormData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FormServer).UploadFormData(&formUploadFormDataServer{stream})
}

type Form_UploadFormDataServer interface {
	SendAndClose(*Message) error
	Recv() (*Map, error)
	grpc.ServerStream
}

type formUploadFormDataServer struct {
	grpc.ServerStream
}

func (x *formUploadFormDataServer) SendAndClose(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *formUploadFormDataServer) Recv() (*Map, error) {
	m := new(Map)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Form_serviceDesc = grpc.ServiceDesc{
	ServiceName: "userInfo.Form",
	HandlerType: (*FormServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFormData",
			Handler:       _Form_UploadFormData_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "userInfo.proto",
}
