// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: api/protobuf/calendar/event.proto

package calendar

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CalendarClient is the client API for Calendar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalendarClient interface {
	Add(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Event, error)
	Update(ctx context.Context, in *UpdateEventRequest, opts ...grpc.CallOption) (*Event, error)
	GetByID(ctx context.Context, in *GetEventByIDRequest, opts ...grpc.CallOption) (*Event, error)
	List(ctx context.Context, in *ListEventsRequest, opts ...grpc.CallOption) (Calendar_ListClient, error)
	Delete(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*Event, error)
}

type calendarClient struct {
	cc grpc.ClientConnInterface
}

func NewCalendarClient(cc grpc.ClientConnInterface) CalendarClient {
	return &calendarClient{cc}
}

func (c *calendarClient) Add(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/calendar.Calendar/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) Update(ctx context.Context, in *UpdateEventRequest, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/calendar.Calendar/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) GetByID(ctx context.Context, in *GetEventByIDRequest, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/calendar.Calendar/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) List(ctx context.Context, in *ListEventsRequest, opts ...grpc.CallOption) (Calendar_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &Calendar_ServiceDesc.Streams[0], "/calendar.Calendar/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &calendarListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Calendar_ListClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type calendarListClient struct {
	grpc.ClientStream
}

func (x *calendarListClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calendarClient) Delete(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/calendar.Calendar/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalendarServer is the server API for Calendar service.
// All implementations must embed UnimplementedCalendarServer
// for forward compatibility
type CalendarServer interface {
	Add(context.Context, *Event) (*Event, error)
	Update(context.Context, *UpdateEventRequest) (*Event, error)
	GetByID(context.Context, *GetEventByIDRequest) (*Event, error)
	List(*ListEventsRequest, Calendar_ListServer) error
	Delete(context.Context, *DeleteEventRequest) (*Event, error)
	mustEmbedUnimplementedCalendarServer()
}

// UnimplementedCalendarServer must be embedded to have forward compatible implementations.
type UnimplementedCalendarServer struct {
}

func (UnimplementedCalendarServer) Add(context.Context, *Event) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedCalendarServer) Update(context.Context, *UpdateEventRequest) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCalendarServer) GetByID(context.Context, *GetEventByIDRequest) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedCalendarServer) List(*ListEventsRequest, Calendar_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCalendarServer) Delete(context.Context, *DeleteEventRequest) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCalendarServer) mustEmbedUnimplementedCalendarServer() {}

// UnsafeCalendarServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalendarServer will
// result in compilation errors.
type UnsafeCalendarServer interface {
	mustEmbedUnimplementedCalendarServer()
}

func RegisterCalendarServer(s grpc.ServiceRegistrar, srv CalendarServer) {
	s.RegisterService(&Calendar_ServiceDesc, srv)
}

func _Calendar_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.Calendar/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).Add(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.Calendar/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).Update(ctx, req.(*UpdateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.Calendar/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).GetByID(ctx, req.(*GetEventByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalendarServer).List(m, &calendarListServer{stream})
}

type Calendar_ListServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type calendarListServer struct {
	grpc.ServerStream
}

func (x *calendarListServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _Calendar_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.Calendar/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).Delete(ctx, req.(*DeleteEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Calendar_ServiceDesc is the grpc.ServiceDesc for Calendar service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Calendar_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calendar.Calendar",
	HandlerType: (*CalendarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Calendar_Add_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Calendar_Update_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _Calendar_GetByID_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Calendar_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _Calendar_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/protobuf/calendar/event.proto",
}
