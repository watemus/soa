package repository

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/hashicorp/go-uuid"
	"log"
	"time"
)

type Statistics struct {
	producer sarama.SyncProducer
}

type StatEvent struct {
	Username string `json:"username"`
	TaskId   uint64 `json:"task_id"`
}

const (
	intervalBetweenConnects = 5 * time.Second
)

func NewStatistics() (*Statistics, error) {
	var errRes error
	for i := 0; i < attemptsToConnect; i++ {
		log.Printf("attempt N%d to connect to kafka", i)
		producer, err := sarama.NewSyncProducer([]string{"kafka:9092"}, nil)
		if err != nil {
			errRes = err
			time.Sleep(intervalBetweenConnects)
			continue
		}
		return &Statistics{
			producer: producer,
		}, nil
	}
	return nil, errRes
}

func (s *Statistics) ProduceEvent(event StatEvent, topic string) error {
	bytes, err := json.Marshal(&event)
	if err != nil {
		log.Printf("error while parsing stat event: %v", err)
		return err
	}
	id, err := uuid.GenerateUUID()
	if err != nil {
		log.Printf("error while generating uuid: %v", err)
		return err
	}
	_, _, err = s.producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(id),
		Value: sarama.ByteEncoder(bytes),
	})
	if err != nil {
		log.Printf("error while sending meessage: %v", err)
		return err
	}
	return nil
}
