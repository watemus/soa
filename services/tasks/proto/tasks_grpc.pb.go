// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: proto/tasks.proto

package taskspb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	TaskService_CreateTask_FullMethodName   = "/TaskService/CreateTask"
	TaskService_UpdateTask_FullMethodName   = "/TaskService/UpdateTask"
	TaskService_DeleteTask_FullMethodName   = "/TaskService/DeleteTask"
	TaskService_GetTask_FullMethodName      = "/TaskService/GetTask"
	TaskService_GetListTasks_FullMethodName = "/TaskService/GetListTasks"
)

// TaskServiceClient is the client API for TaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskServiceClient interface {
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error)
	UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*UpdateTaskResponse, error)
	DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*DeleteTaskResponse, error)
	GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*GetTaskResponse, error)
	GetListTasks(ctx context.Context, in *GetListTasksRequest, opts ...grpc.CallOption) (TaskService_GetListTasksClient, error)
}

type taskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskServiceClient(cc grpc.ClientConnInterface) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTaskResponse)
	err := c.cc.Invoke(ctx, TaskService_CreateTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*UpdateTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTaskResponse)
	err := c.cc.Invoke(ctx, TaskService_UpdateTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*DeleteTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTaskResponse)
	err := c.cc.Invoke(ctx, TaskService_DeleteTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*GetTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTaskResponse)
	err := c.cc.Invoke(ctx, TaskService_GetTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) GetListTasks(ctx context.Context, in *GetListTasksRequest, opts ...grpc.CallOption) (TaskService_GetListTasksClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &TaskService_ServiceDesc.Streams[0], TaskService_GetListTasks_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &taskServiceGetListTasksClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TaskService_GetListTasksClient interface {
	Recv() (*GetListTasksResponse, error)
	grpc.ClientStream
}

type taskServiceGetListTasksClient struct {
	grpc.ClientStream
}

func (x *taskServiceGetListTasksClient) Recv() (*GetListTasksResponse, error) {
	m := new(GetListTasksResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TaskServiceServer is the server API for TaskService service.
// All implementations must embed UnimplementedTaskServiceServer
// for forward compatibility
type TaskServiceServer interface {
	CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)
	UpdateTask(context.Context, *UpdateTaskRequest) (*UpdateTaskResponse, error)
	DeleteTask(context.Context, *DeleteTaskRequest) (*DeleteTaskResponse, error)
	GetTask(context.Context, *GetTaskRequest) (*GetTaskResponse, error)
	GetListTasks(*GetListTasksRequest, TaskService_GetListTasksServer) error
	mustEmbedUnimplementedTaskServiceServer()
}

// UnimplementedTaskServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTaskServiceServer struct {
}

func (UnimplementedTaskServiceServer) CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedTaskServiceServer) UpdateTask(context.Context, *UpdateTaskRequest) (*UpdateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTask not implemented")
}
func (UnimplementedTaskServiceServer) DeleteTask(context.Context, *DeleteTaskRequest) (*DeleteTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}
func (UnimplementedTaskServiceServer) GetTask(context.Context, *GetTaskRequest) (*GetTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (UnimplementedTaskServiceServer) GetListTasks(*GetListTasksRequest, TaskService_GetListTasksServer) error {
	return status.Errorf(codes.Unimplemented, "method GetListTasks not implemented")
}
func (UnimplementedTaskServiceServer) mustEmbedUnimplementedTaskServiceServer() {}

// UnsafeTaskServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskServiceServer will
// result in compilation errors.
type UnsafeTaskServiceServer interface {
	mustEmbedUnimplementedTaskServiceServer()
}

func RegisterTaskServiceServer(s grpc.ServiceRegistrar, srv TaskServiceServer) {
	s.RegisterService(&TaskService_ServiceDesc, srv)
}

func _TaskService_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_CreateTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_UpdateTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).UpdateTask(ctx, req.(*UpdateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_DeleteTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).DeleteTask(ctx, req.(*DeleteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_GetTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).GetTask(ctx, req.(*GetTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_GetListTasks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetListTasksRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TaskServiceServer).GetListTasks(m, &taskServiceGetListTasksServer{ServerStream: stream})
}

type TaskService_GetListTasksServer interface {
	Send(*GetListTasksResponse) error
	grpc.ServerStream
}

type taskServiceGetListTasksServer struct {
	grpc.ServerStream
}

func (x *taskServiceGetListTasksServer) Send(m *GetListTasksResponse) error {
	return x.ServerStream.SendMsg(m)
}

// TaskService_ServiceDesc is the grpc.ServiceDesc for TaskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _TaskService_CreateTask_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _TaskService_UpdateTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _TaskService_DeleteTask_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _TaskService_GetTask_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetListTasks",
			Handler:       _TaskService_GetListTasks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/tasks.proto",
}