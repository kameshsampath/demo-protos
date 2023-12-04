// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: todo/todo.proto

package todo

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Todo_AddTodo_FullMethodName          = "/todo.Todo/AddTodo"
	Todo_GetTodo_FullMethodName          = "/todo.Todo/GetTodo"
	Todo_UpdateTodoStatus_FullMethodName = "/todo.Todo/UpdateTodoStatus"
	Todo_TodoList_FullMethodName         = "/todo.Todo/TodoList"
)

// TodoClient is the client API for Todo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoClient interface {
	AddTodo(ctx context.Context, in *TodoAddRequest, opts ...grpc.CallOption) (*TodoResponse, error)
	GetTodo(ctx context.Context, in *GetTodoRequest, opts ...grpc.CallOption) (*TodoResponse, error)
	UpdateTodoStatus(ctx context.Context, in *UpdateTodoStatusRequest, opts ...grpc.CallOption) (*TodoResponse, error)
	TodoList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Todo_TodoListClient, error)
}

type todoClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoClient(cc grpc.ClientConnInterface) TodoClient {
	return &todoClient{cc}
}

func (c *todoClient) AddTodo(ctx context.Context, in *TodoAddRequest, opts ...grpc.CallOption) (*TodoResponse, error) {
	out := new(TodoResponse)
	err := c.cc.Invoke(ctx, Todo_AddTodo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) GetTodo(ctx context.Context, in *GetTodoRequest, opts ...grpc.CallOption) (*TodoResponse, error) {
	out := new(TodoResponse)
	err := c.cc.Invoke(ctx, Todo_GetTodo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) UpdateTodoStatus(ctx context.Context, in *UpdateTodoStatusRequest, opts ...grpc.CallOption) (*TodoResponse, error) {
	out := new(TodoResponse)
	err := c.cc.Invoke(ctx, Todo_UpdateTodoStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) TodoList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Todo_TodoListClient, error) {
	stream, err := c.cc.NewStream(ctx, &Todo_ServiceDesc.Streams[0], Todo_TodoList_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &todoTodoListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Todo_TodoListClient interface {
	Recv() (*Task, error)
	grpc.ClientStream
}

type todoTodoListClient struct {
	grpc.ClientStream
}

func (x *todoTodoListClient) Recv() (*Task, error) {
	m := new(Task)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TodoServer is the server API for Todo service.
// All implementations must embed UnimplementedTodoServer
// for forward compatibility
type TodoServer interface {
	AddTodo(context.Context, *TodoAddRequest) (*TodoResponse, error)
	GetTodo(context.Context, *GetTodoRequest) (*TodoResponse, error)
	UpdateTodoStatus(context.Context, *UpdateTodoStatusRequest) (*TodoResponse, error)
	TodoList(*emptypb.Empty, Todo_TodoListServer) error
	mustEmbedUnimplementedTodoServer()
}

// UnimplementedTodoServer must be embedded to have forward compatible implementations.
type UnimplementedTodoServer struct {
}

func (UnimplementedTodoServer) AddTodo(context.Context, *TodoAddRequest) (*TodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTodo not implemented")
}
func (UnimplementedTodoServer) GetTodo(context.Context, *GetTodoRequest) (*TodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodo not implemented")
}
func (UnimplementedTodoServer) UpdateTodoStatus(context.Context, *UpdateTodoStatusRequest) (*TodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTodoStatus not implemented")
}
func (UnimplementedTodoServer) TodoList(*emptypb.Empty, Todo_TodoListServer) error {
	return status.Errorf(codes.Unimplemented, "method TodoList not implemented")
}
func (UnimplementedTodoServer) mustEmbedUnimplementedTodoServer() {}

// UnsafeTodoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoServer will
// result in compilation errors.
type UnsafeTodoServer interface {
	mustEmbedUnimplementedTodoServer()
}

func RegisterTodoServer(s grpc.ServiceRegistrar, srv TodoServer) {
	s.RegisterService(&Todo_ServiceDesc, srv)
}

func _Todo_AddTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).AddTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Todo_AddTodo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).AddTodo(ctx, req.(*TodoAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_GetTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).GetTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Todo_GetTodo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).GetTodo(ctx, req.(*GetTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_UpdateTodoStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTodoStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).UpdateTodoStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Todo_UpdateTodoStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).UpdateTodoStatus(ctx, req.(*UpdateTodoStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_TodoList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TodoServer).TodoList(m, &todoTodoListServer{stream})
}

type Todo_TodoListServer interface {
	Send(*Task) error
	grpc.ServerStream
}

type todoTodoListServer struct {
	grpc.ServerStream
}

func (x *todoTodoListServer) Send(m *Task) error {
	return x.ServerStream.SendMsg(m)
}

// Todo_ServiceDesc is the grpc.ServiceDesc for Todo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Todo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todo.Todo",
	HandlerType: (*TodoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTodo",
			Handler:    _Todo_AddTodo_Handler,
		},
		{
			MethodName: "GetTodo",
			Handler:    _Todo_GetTodo_Handler,
		},
		{
			MethodName: "UpdateTodoStatus",
			Handler:    _Todo_UpdateTodoStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TodoList",
			Handler:       _Todo_TodoList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "todo/todo.proto",
}
