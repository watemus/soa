package service

import "main-service/internal/repository"

type Events struct {
	stats *repository.Statistics
}

func NewEvents(stats *repository.Statistics) *Events {
	return &Events{stats: stats}
}

func (e *Events) SendView(event repository.StatEvent) error {
	return e.stats.ProduceEvent(event, "views")
}

func (e *Events) SendLike(event repository.StatEvent) error {
	return e.stats.ProduceEvent(event, "likes")
}
