package repository

type Task struct {
	Id     uint64 `db:"id"`
	Author string `db:"author"`
	Name   string `db:"task_name"`
	Body   string `db:"body"`
}

type Repository interface {
	CreateTask(author string, name string, body string) (uint64, error)
	UpdateTask(taskId uint64, name string, body string) error
	DeleteTask(taskId uint64) error
	GetTask(taskId uint64) (Task, error)
	GetListTasks(username string, offset uint64, limit uint64, tasks chan<- Task, done chan<- struct{})
}
