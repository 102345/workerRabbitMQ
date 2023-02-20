package di

import (
	"github.com/marc/workerRabbitMQ-example/adapter/postgres"
	"github.com/marc/workerRabbitMQ-example/adapter/postgres/stockProductRepository"
	"github.com/marc/workerRabbitMQ-example/core/domain"
	"github.com/marc/workerRabbitMQ-example/core/usecase/stockProductUseCase"
)

// ConfigStockProductDI return a IStockProductUseCase abstraction with dependency injection configuration
func ConfigStockProductDI(conn postgres.PoolInterface) domain.IStockProductUseCase {
	stockProductRepository := stockProductRepository.New(conn)
	stockProductUseCase := stockProductUseCase.New(stockProductRepository)

	return stockProductUseCase
}
