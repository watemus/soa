package internal

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"log"
)

func connect() (driver.Conn, error) {
	var (
		ctx       = context.Background()
		conn, err = clickhouse.Open(&clickhouse.Options{
			Addr: []string{"ch_server:9000"},
		})
	)

	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return nil, err
	}
	return conn, nil
}

type Repo struct {
	db driver.Conn
}

func NewRepo() (*Repo, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}
	return &Repo{
		db: db,
	}, nil
}

func (r *Repo) StatTask(taskId uint64) (uint64, uint64, error) {
	var likes, views uint64

	row := r.db.QueryRow(context.Background(), "select count(distinct username) from target_likes where task_id=$1", taskId)
	if err := row.Err(); err != nil {
		log.Printf("error while stat task: %v", err)
		return 0, 0, err
	}
	err := row.Scan(&likes)
	if err != nil {
		log.Printf("error while stat task: %v", err)
		return 0, 0, err
	}

	row = r.db.QueryRow(context.Background(), "select count(distinct username) from target_views where task_id=$1", taskId)
	if err = row.Err(); err != nil {
		log.Printf("error while stat task: %v", err)
		return 0, 0, err
	}
	err = row.Scan(&views)
	if err != nil {
		log.Printf("error while stat task: %v", err)
		return 0, 0, err
	}
	return likes, views, nil
}

type Stat struct {
	Count uint64
	Id    uint64
}

func (r *Repo) TopTask(orderBy string) ([]Stat, error) {
	rows, err := r.db.Query(context.Background(), "select count(distinct username), task_id from $1 group by task_id limit 5", orderBy)
	if err != nil {
		log.Printf("error while top task: %v", err)
		return nil, err
	}
	var stats []Stat
	for rows.Next() {
		var count, id uint64
		err = rows.Scan(&count, &id)
		if err != nil {
			log.Printf("error while top task: %v", err)
			return nil, err
		}
		stats = append(stats, Stat{
			Count: count,
			Id:    id,
		})
	}
	return stats, nil
}

type StatUser struct {
	Count    uint64
	Username string
}

func (r *Repo) TopUser() ([]StatUser, error) {
	rows, err := r.db.Query(context.Background(), "select count(distinct task_id), username from target_likes group by task_id limit 3")
	if err != nil {
		log.Printf("error while top task: %v", err)
		return nil, err
	}
	var stats []StatUser
	for rows.Next() {
		var count uint64
		var username string
		err = rows.Scan(&count, &username)
		if err != nil {
			log.Printf("error while top task: %v", err)
			return nil, err
		}
		stats = append(stats, StatUser{
			Count:    count,
			Username: username,
		})
	}
	return stats, nil
}
