package domain

import (
	"time"

	"github.com/marc/workerRabbitMQ-example/core/dto"
)

type QueueProcess struct {
	ID           int32
	QueueMessage string
	Message      string
	Result       string
	CreatedAt    time.Time
}

type IQueueProcessUseCase interface {
	Create(queue *dto.QueueProcessDTO) (*QueueProcess, error)
}

type IQueueProcessRepository interface {
	Create(queue *dto.QueueProcessDTO) (*QueueProcess, error)
}
