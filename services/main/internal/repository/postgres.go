package repository

import (
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
