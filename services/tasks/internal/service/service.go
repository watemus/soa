package service

import (
	"context"
	"tasks-service/internal/repository"
	taskspb "tasks-service/proto"
	"time"
)

type Service struct {
	taskspb.TaskServiceServer
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateTask(ctx context.Context, req *taskspb.CreateTaskRequest) (*taskspb.CreateTaskResponse, error) {
	taskId, err := s.repo.CreateTask(req.Common.Username, req.Name, req.Body)
	if err != nil {
		return &taskspb.CreateTaskResponse{
			TaskId: 0,
			Status: err.Error(),
		}, nil
	}
	return &taskspb.CreateTaskResponse{
		TaskId: taskId,
		Status: "ok",
	}, nil
}

func (s *Service) UpdateTask(ctx context.Context, req *taskspb.UpdateTaskRequest) (*taskspb.UpdateTaskResponse, error) {
	task, err := s.repo.GetTask(req.TaskId)
	if err != nil {
		return &taskspb.UpdateTaskResponse{
			Status: err.Error(),
		}, nil
	}
	if task.Author != req.Common.Username {
		return &taskspb.UpdateTaskResponse{
			Status: "not enough rights",
		}, nil
	}
	err = s.repo.UpdateTask(req.TaskId, req.Name, req.Body)
	if err != nil {
		return &taskspb.UpdateTaskResponse{
			Status: err.Error(),
		}, nil
	}
	return &taskspb.UpdateTaskResponse{
		Status: "ok",
	}, nil
}

func (s *Service) DeleteTask(ctx context.Context, req *taskspb.DeleteTaskRequest) (*taskspb.DeleteTaskResponse, error) {
	task, err := s.repo.GetTask(req.TaskId)
	if err != nil {
		return &taskspb.DeleteTaskResponse{
			Status: err.Error(),
		}, nil
	}
	if task.Author != req.Common.Username {
		return &taskspb.DeleteTaskResponse{
			Status: "not enough rights",
		}, nil
	}
	err = s.repo.DeleteTask(req.TaskId)
	if err != nil {
		return &taskspb.DeleteTaskResponse{
			Status: err.Error(),
		}, nil
	}
	return &taskspb.DeleteTaskResponse{
		Status: "ok",
	}, nil
}

func (s *Service) GetTask(ctx context.Context, req *taskspb.GetTaskRequest) (*taskspb.GetTaskResponse, error) {
	task, err := s.repo.GetTask(req.TaskId)
	if err != nil {
		return &taskspb.GetTaskResponse{
			Status: err.Error(),
			Task:   nil,
		}, nil
	}
	return &taskspb.GetTaskResponse{
		Status: "ok",
		Task: &taskspb.Task{
			Id: task.Id,
			Spec: &taskspb.TaskSpec{
				Author: task.Author,
				Name:   task.Name,
				Body:   task.Body,
			},
		},
	}, nil
}

func (s *Service) GetListTasks(
	req *taskspb.GetListTasksRequest,
	resp taskspb.TaskService_GetListTasksServer) error {
	tasks := make(chan repository.Task)
	done := make(chan struct{})
	go s.repo.GetListTasks(req.Offset, req.Limit, tasks, done)
	for {
		select {
		case task := <-tasks:
			err := resp.Send(&taskspb.GetListTasksResponse{
				Task: &taskspb.Task{
					Id: task.Id,
					Spec: &taskspb.TaskSpec{
						Author: task.Author,
						Name:   task.Name,
						Body:   task.Body,
					},
				},
			})
			if err != nil {
				return err
			}
		case <-done:
			return nil
		case <-time.After(100 * time.Millisecond):
			return nil
		}
	}
}
