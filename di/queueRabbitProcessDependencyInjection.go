package di

import (
	"github.com/marc/workerRabbitMQ-example/adapter/postgres"
	"github.com/marc/workerRabbitMQ-example/adapter/postgres/queueRabbitProcessRepository"
	"github.com/marc/workerRabbitMQ-example/core/domain"
	"github.com/marc/workerRabbitMQ-example/core/usecase/queueRabbitProcessUseCase"
)

// ConfigQueueProcessDI return a IQueueProcessUseCase abstraction with dependency injection configuration
func ConfigQueueProcessDI(conn postgres.PoolInterface) domain.IQueueProcessUseCase {
	queueRabbitProcessRepository := queueRabbitProcessRepository.New(conn)
	queueRabbitProcessUseCase := queueRabbitProcessUseCase.New(queueRabbitProcessRepository)

	return queueRabbitProcessUseCase
}
