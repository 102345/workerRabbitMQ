package queueRabbitProcessUseCase

import (
	"github.com/marc/workerRabbitMQ-example/core/domain"
	"github.com/marc/workerRabbitMQ-example/core/dto"
)

type usecase struct {
	repository domain.IQueueProcessRepository
}

// New returns contract implementation of QueueProcessUseCase
func New(repository domain.IQueueProcessRepository) domain.IQueueProcessUseCase {
	return &usecase{
		repository: repository,
	}
}

func (usecase usecase) Create(queue *dto.QueueProcessDTO) (*domain.QueueProcess, error) {

	queueProcess, err := usecase.repository.Create(queue)

	if err != nil {
		return nil, err
	}

	return queueProcess, nil
}
