// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v4.0.0
// source: booking.proto

package booking

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Booking struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeatNumber    int32  `protobuf:"varint,1,opt,name=seatNumber,proto3" json:"seatNumber,omitempty"`
	PassengerName string `protobuf:"bytes,2,opt,name=passengerName,proto3" json:"passengerName,omitempty"`
}

func (x *Booking) Reset() {
	*x = Booking{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Booking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Booking) ProtoMessage() {}

func (x *Booking) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Booking.ProtoReflect.Descriptor instead.
func (*Booking) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{0}
}

func (x *Booking) GetSeatNumber() int32 {
	if x != nil {
		return x.SeatNumber
	}
	return 0
}

func (x *Booking) GetPassengerName() string {
	if x != nil {
		return x.PassengerName
	}
	return ""
}

type Passenger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Passenger) Reset() {
	*x = Passenger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Passenger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Passenger) ProtoMessage() {}

func (x *Passenger) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Passenger.ProtoReflect.Descriptor instead.
func (*Passenger) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{1}
}

func (x *Passenger) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Flight struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlightName string `protobuf:"bytes,1,opt,name=flightName,proto3" json:"flightName,omitempty"`
}

func (x *Flight) Reset() {
	*x = Flight{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Flight) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Flight) ProtoMessage() {}

func (x *Flight) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Flight.ProtoReflect.Descriptor instead.
func (*Flight) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{2}
}

func (x *Flight) GetFlightName() string {
	if x != nil {
		return x.FlightName
	}
	return ""
}

type JourneyPath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source      string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Destination string `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
}

func (x *JourneyPath) Reset() {
	*x = JourneyPath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JourneyPath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JourneyPath) ProtoMessage() {}

func (x *JourneyPath) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JourneyPath.ProtoReflect.Descriptor instead.
func (*JourneyPath) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{3}
}

func (x *JourneyPath) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *JourneyPath) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

var File_booking_proto protoreflect.FileDescriptor

var file_booking_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x4f, 0x0a, 0x07, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x61, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x74, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x73, 0x73,
	0x65, 0x6e, 0x67, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x1f, 0x0a, 0x09, 0x50, 0x61, 0x73,
	0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x28, 0x0a, 0x06, 0x46, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x47, 0x0a, 0x0b, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x45, 0x0a,
	0x0d, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34,
	0x0a, 0x0a, 0x42, 0x6f, 0x6f, 0x6b, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x12, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72,
	0x1a, 0x10, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x22, 0x00, 0x32, 0x46, 0x0a, 0x0a, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x69, 0x6e,
	0x66, 0x6f, 0x12, 0x38, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x73, 0x12, 0x14, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4a, 0x6f, 0x75, 0x72,
	0x6e, 0x65, 0x79, 0x50, 0x61, 0x74, 0x68, 0x1a, 0x0f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x22, 0x00, 0x30, 0x01, 0x42, 0x31, 0x5a, 0x2f,
	0x41, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2d, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x2d, 0x67, 0x52, 0x50, 0x43, 0x2d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67,
	0x65, 0x6e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_booking_proto_rawDescOnce sync.Once
	file_booking_proto_rawDescData = file_booking_proto_rawDesc
)

func file_booking_proto_rawDescGZIP() []byte {
	file_booking_proto_rawDescOnce.Do(func() {
		file_booking_proto_rawDescData = protoimpl.X.CompressGZIP(file_booking_proto_rawDescData)
	})
	return file_booking_proto_rawDescData
}

var file_booking_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_booking_proto_goTypes = []interface{}{
	(*Booking)(nil),     // 0: booking.Booking
	(*Passenger)(nil),   // 1: booking.Passenger
	(*Flight)(nil),      // 2: booking.Flight
	(*JourneyPath)(nil), // 3: booking.JourneyPath
}
var file_booking_proto_depIdxs = []int32{
	1, // 0: booking.TicketService.BookTicket:input_type -> booking.Passenger
	3, // 1: booking.Flightinfo.ListFlights:input_type -> booking.JourneyPath
	0, // 2: booking.TicketService.BookTicket:output_type -> booking.Booking
	2, // 3: booking.Flightinfo.ListFlights:output_type -> booking.Flight
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_booking_proto_init() }
func file_booking_proto_init() {
	if File_booking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_booking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Booking); i {
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
		file_booking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Passenger); i {
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
		file_booking_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Flight); i {
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
		file_booking_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JourneyPath); i {
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
			RawDescriptor: file_booking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_booking_proto_goTypes,
		DependencyIndexes: file_booking_proto_depIdxs,
		MessageInfos:      file_booking_proto_msgTypes,
	}.Build()
	File_booking_proto = out.File
	file_booking_proto_rawDesc = nil
	file_booking_proto_goTypes = nil
	file_booking_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TicketServiceClient is the client API for TicketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TicketServiceClient interface {
	BookTicket(ctx context.Context, in *Passenger, opts ...grpc.CallOption) (*Booking, error)
}

type ticketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTicketServiceClient(cc grpc.ClientConnInterface) TicketServiceClient {
	return &ticketServiceClient{cc}
}

func (c *ticketServiceClient) BookTicket(ctx context.Context, in *Passenger, opts ...grpc.CallOption) (*Booking, error) {
	out := new(Booking)
	err := c.cc.Invoke(ctx, "/booking.TicketService/BookTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TicketServiceServer is the server API for TicketService service.
type TicketServiceServer interface {
	BookTicket(context.Context, *Passenger) (*Booking, error)
}

// UnimplementedTicketServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTicketServiceServer struct {
}

func (*UnimplementedTicketServiceServer) BookTicket(context.Context, *Passenger) (*Booking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BookTicket not implemented")
}

func RegisterTicketServiceServer(s *grpc.Server, srv TicketServiceServer) {
	s.RegisterService(&_TicketService_serviceDesc, srv)
}

func _TicketService_BookTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Passenger)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).BookTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.TicketService/BookTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).BookTicket(ctx, req.(*Passenger))
	}
	return interceptor(ctx, in, info, handler)
}

var _TicketService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "booking.TicketService",
	HandlerType: (*TicketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BookTicket",
			Handler:    _TicketService_BookTicket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "booking.proto",
}

// FlightinfoClient is the client API for Flightinfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FlightinfoClient interface {
	ListFlights(ctx context.Context, in *JourneyPath, opts ...grpc.CallOption) (Flightinfo_ListFlightsClient, error)
}

type flightinfoClient struct {
	cc grpc.ClientConnInterface
}

func NewFlightinfoClient(cc grpc.ClientConnInterface) FlightinfoClient {
	return &flightinfoClient{cc}
}

func (c *flightinfoClient) ListFlights(ctx context.Context, in *JourneyPath, opts ...grpc.CallOption) (Flightinfo_ListFlightsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Flightinfo_serviceDesc.Streams[0], "/booking.Flightinfo/ListFlights", opts...)
	if err != nil {
		return nil, err
	}
	x := &flightinfoListFlightsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Flightinfo_ListFlightsClient interface {
	Recv() (*Flight, error)
	grpc.ClientStream
}

type flightinfoListFlightsClient struct {
	grpc.ClientStream
}

func (x *flightinfoListFlightsClient) Recv() (*Flight, error) {
	m := new(Flight)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FlightinfoServer is the server API for Flightinfo service.
type FlightinfoServer interface {
	ListFlights(*JourneyPath, Flightinfo_ListFlightsServer) error
}

// UnimplementedFlightinfoServer can be embedded to have forward compatible implementations.
type UnimplementedFlightinfoServer struct {
}

func (*UnimplementedFlightinfoServer) ListFlights(*JourneyPath, Flightinfo_ListFlightsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListFlights not implemented")
}

func RegisterFlightinfoServer(s *grpc.Server, srv FlightinfoServer) {
	s.RegisterService(&_Flightinfo_serviceDesc, srv)
}

func _Flightinfo_ListFlights_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(JourneyPath)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FlightinfoServer).ListFlights(m, &flightinfoListFlightsServer{stream})
}

type Flightinfo_ListFlightsServer interface {
	Send(*Flight) error
	grpc.ServerStream
}

type flightinfoListFlightsServer struct {
	grpc.ServerStream
}

func (x *flightinfoListFlightsServer) Send(m *Flight) error {
	return x.ServerStream.SendMsg(m)
}

var _Flightinfo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "booking.Flightinfo",
	HandlerType: (*FlightinfoServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListFlights",
			Handler:       _Flightinfo_ListFlights_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "booking.proto",
}
