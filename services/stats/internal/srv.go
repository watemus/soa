package internal

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	taskspb "stats/proto"
)

type Srv struct {
	taskspb.StatsServiceServer
	repo *Repo
}

func NewSrv(repo *Repo) *Srv {
	return &Srv{repo: repo}
}

func (s *Srv) StatTask(ctx context.Context, request *taskspb.StatTaskRequest) (*taskspb.StatTaskResponse, error) {
	likes, views, err := s.repo.StatTask(request.TaskId)
	if err != nil {
		log.Printf("error while stat task: %v", err)
		return &taskspb.StatTaskResponse{
			Likes:  0,
			Views:  0,
			Status: "error",
		}, nil
	}
	return &taskspb.StatTaskResponse{
		Likes:  likes,
		Views:  views,
		Status: "ok",
	}, nil
}

func (s *Srv) TopTask(req *taskspb.TopRequest, stream taskspb.StatsService_TopTaskServer) error {
	stats, err := s.repo.TopTask(req.OrderBy)
	if err != nil {
		return err
	}
	for _, stat := range stats {
		err = stream.Send(&taskspb.TopTaskResponse{
			TaskId: stat.Id,
			Count:  stat.Count,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Srv) TopUser(req *emptypb.Empty, stream taskspb.StatsService_TopUserServer) error {
	stats, err := s.repo.TopUser()
	if err != nil {
		return err
	}
	for _, stat := range stats {
		err = stream.Send(&taskspb.TopUsersResponse{
			Username: stat.Username,
			Count:    stat.Count,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
