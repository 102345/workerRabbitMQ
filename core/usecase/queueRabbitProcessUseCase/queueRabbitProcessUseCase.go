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

func (usecase usecase) Create(message string) (*domain.QueueProcess, error) {

	queueProcessDTO := dto.QueueProcessDTO{}
	queueProcess, err := usecase.repository.Create(&queueProcessDTO)

	if err != nil {
		return nil, err
	}

	return queueProcess, nil
}
