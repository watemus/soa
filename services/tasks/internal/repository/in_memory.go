package repository

import (
	"errors"
	"time"
)

type InMemoryRepo struct {
	taskCount uint64
	tasks     []Task
	deleted   []bool
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{}
}

func (r *InMemoryRepo) CreateTask(author string, name string, body string) (uint64, error) {
	id := r.taskCount
	r.taskCount += 1
	r.tasks = append(r.tasks, Task{
		Id:     id,
		Author: author,
		Name:   name,
		Body:   body,
	})
	r.deleted = append(r.deleted, false)
	return id, nil
}

func (r *InMemoryRepo) UpdateTask(taskId uint64, name string, body string) error {
	if taskId >= uint64(len(r.deleted)) || r.deleted[taskId] {
		return errors.New("not found")
	}
	r.tasks[taskId].Name = name
	r.tasks[taskId].Body = body
	return nil
}

func (r *InMemoryRepo) DeleteTask(taskId uint64) error {
	if taskId >= uint64(len(r.deleted)) || r.deleted[taskId] {
		return errors.New("not found")
	}
	r.deleted[taskId] = true
	return nil
}

func (r *InMemoryRepo) GetTask(taskId uint64) (Task, error) {
	if taskId >= uint64(len(r.deleted)) || r.deleted[taskId] {
		return Task{}, errors.New("not found")
	}
	return r.tasks[taskId], nil
}

func (r *InMemoryRepo) GetListTasks(username string, offset uint64, limit uint64, tasks chan<- Task, done chan<- struct{}) {
	pos := 0
	count := 0
	for uint64(count) < offset && pos < len(r.tasks) {
		if !r.deleted[pos] {
			count += 1
		}
		pos += 1
	}
	count = 0
	for uint64(count) < limit && pos < len(r.tasks) {
		if !r.deleted[pos] {
			select {
			case tasks <- r.tasks[pos]:
				count += 1
			case <-time.After(time.Millisecond * 10):
				break
			}
		}
		pos += 1
	}
	select {
	case done <- struct{}{}:
		return
	case <-time.After(time.Millisecond * 50):
		return
	}
}
