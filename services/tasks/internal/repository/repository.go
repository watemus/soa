package repository

type Task struct {
	Id     uint64
	Author string
	Name   string
	Body   string
}

type Repository interface {
	CreateTask(author string, name string, body string) (uint64, error)
	UpdateTask(taskId uint64, name string, body string) error
	DeleteTask(taskId uint64) error
	GetTask(taskId uint64) (Task, error)
	GetListTasks(offset uint64, limit uint64, tasks chan<- Task, done chan<- struct{})
}
