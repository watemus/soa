package repository

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	taskspb "main-service/proto"
)

type Tasks struct {
	conn   *grpc.ClientConn
	client taskspb.TaskServiceClient
}

func NewTasks(addr string) (*Tasks, error) {
	insecureOpts := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient(addr, insecureOpts)
	if err != nil {
		return nil, err
	}
	client := taskspb.NewTaskServiceClient(conn)
	log.Println("created tasks client")
	return &Tasks{
		conn:   conn,
		client: client,
	}, nil
}

func (r *Tasks) CreateTask(req *taskspb.CreateTaskRequest) (*taskspb.CreateTaskResponse, error) {
	task, err := r.client.CreateTask(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (r *Tasks) UpdateTask(req *taskspb.UpdateTaskRequest) (*taskspb.UpdateTaskResponse, error) {
	res, err := r.client.UpdateTask(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *Tasks) DeleteTask(req *taskspb.DeleteTaskRequest) (*taskspb.DeleteTaskResponse, error) {
	res, err := r.client.DeleteTask(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *Tasks) GetTask(req *taskspb.GetTaskRequest) (*taskspb.GetTaskResponse, error) {
	res, err := r.client.GetTask(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *Tasks) GetListTasks(req *taskspb.GetListTasksRequest) ([]*taskspb.GetListTasksResponse, error) {
	stream, err := r.client.GetListTasks(context.Background(), req)
	if err != nil {
		return nil, err
	}
	var tasks []*taskspb.GetListTasksResponse
	for {
		task, err := stream.Recv()
		if err != nil {
			break
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
