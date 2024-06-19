package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

const (
	attemptsToConnect = 10
	attemptDelay      = time.Second * 5
)

func NewPostgresDB(cfg PostgresConfig) (*sqlx.DB, error) {
	var retErr error
	for i := 1; i <= attemptsToConnect; i++ {
		if i != 1 {
			time.Sleep(attemptDelay)
		}
		log.Printf("attempt N%d...", i)
		db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
		if err != nil {
			retErr = err
			continue
		}

		err = db.Ping()

		if err != nil {
			retErr = err
			continue
		}
		return db, nil
	}
	return nil, retErr
}

type PostgresRepo struct {
	db *sqlx.DB
}

func NewPostgresRepo(db *sqlx.DB) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (r *PostgresRepo) CreateTask(author string, name string, body string) (uint64, error) {
	var id uint64
	query := fmt.Sprintf("INSERT INTO tasks (task_name, body, author) values ($1, $2, $3) RETURNING id")
	row := r.db.QueryRow(query, name, body, author)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *PostgresRepo) UpdateTask(taskId uint64, name string, body string) error {
	query := "UPDATE tasks SET task_name=$1, body=$2 WHERE task_id=$3"
	_, err := r.db.Exec(query, name, body, taskId)
	if err != nil {
		log.Printf("error while updating task %s", err.Error())
		return errors.New("error while updating user")
	}
	return nil
}

func (r *PostgresRepo) DeleteTask(taskId uint64) error {
	query := "DELETE FROM tasks WHERE id=$1"
	_, err := r.db.Exec(query, taskId)
	if err != nil {
		log.Printf("error while deleting task %s", err.Error())
		return errors.New("error while updating user")
	}
	return nil
}

func (r *PostgresRepo) GetTask(taskId uint64) (Task, error) {
	var task Task
	query := "SELECT * FROM tasks WHERE id=$1"
	err := r.db.Get(&task, query, taskId)
	if err != nil {
		if err == sql.ErrNoRows {
			return Task{}, errors.New("task not found")
		} else {
			log.Printf("error while getting task %s", err.Error())
			return Task{}, errors.New("error while getting task")
		}
	}
	return task, nil
}

func (r *PostgresRepo) GetListTasks(username string, offset uint64, limit uint64, tasks chan<- Task, done chan<- struct{}) {
	query := "SELECT * FROM tasks WHERE id>=$1 AND author=$2 ORDER BY id LIMIT $3"
	log.Printf("executing query: %s", query)
	rows, err := r.db.Queryx(query, offset, username, limit)
	closeList := func() {
		select {
		case done <- struct{}{}:
			return
		case <-time.After(10 * time.Millisecond):
			return
		}
	}
	if err != nil {
		log.Printf("error while getting list tasks %s", err.Error())
		closeList()
		return
	}
	for rows.Next() {
		var task Task
		err = rows.StructScan(&task)
		if err != nil {
			log.Printf("error while getting list tasks %s", err.Error())
			closeList()
			return
		}
		select {
		case <-time.After(10 * time.Millisecond):
			break
		case tasks <- task:
		}
	}
	closeList()
}
