package di

import (
	"github.com/marc/workerRabbitMQ-example/adapter/postgres"
	"github.com/marc/workerRabbitMQ-example/adapter/postgres/productRepository"
	"github.com/marc/workerRabbitMQ-example/core/domain"
	"github.com/marc/workerRabbitMQ-example/core/usecase/productUseCase"
)

// ConfigProductDI return a IProductUseCase abstraction with dependency injection configuration
func ConfigProductDI(conn postgres.PoolInterface) domain.IProductUseCase {
	productRepository := productRepository.New(conn)
	productUseCase := productUseCase.New(productRepository)

	return productUseCase
}
