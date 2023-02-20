package domain

import (
	"time"

	"github.com/marc/workerRabbitMQ-example/core/dto"
)

type QueueProcess struct {
	ID        int32
	Message   string
	Result    string
	CreatedAt time.Time
}

type IQueueProcessUseCase interface {
	Create(message string) (*QueueProcess, error)
}

type IQueueProcessRepository interface {
	Create(queue *dto.QueueProcessDTO) (*QueueProcess, error)
}
