package queueRabbitProcessRepository

import (
	"context"

	"github.com/marc/workerRabbitMQ-example/adapter/postgres"
	"github.com/marc/workerRabbitMQ-example/core/domain"
	"github.com/marc/workerRabbitMQ-example/core/dto"
)

type repository struct {
	db postgres.PoolInterface
}

// New returns contract implementation of QueueProcessRepository
func New(db postgres.PoolInterface) domain.IQueueProcessRepository {
	return &repository{
		db: db,
	}
}

func (repository repository) Create(queue *dto.QueueProcessDTO) (*domain.QueueProcess, error) {

	ctx := context.Background()
	queueRet := domain.QueueProcess{}
	err := repository.db.QueryRow(
		ctx,
		"INSERT INTO queue_message_process (queuemessage,message,result) VALUES ($1, $2, $3) returning *",
		queue.QueueMessage,
		queue.Message,
		queue.Result,
	).Scan(
		&queueRet.ID,
		&queueRet.QueueMessage,
		&queueRet.Message,
		&queueRet.Result,
		&queueRet.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &queueRet, nil
}
