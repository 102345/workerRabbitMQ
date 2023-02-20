package stockProductUseCase

import (
	"github.com/marc/workerRabbitMQ-example/core/domain"
	"github.com/marc/workerRabbitMQ-example/core/dto"
)

type usecase struct {
	repository domain.IStockProductRepository
}

// New returns contract implementation of StockProductUseCase
func New(repository domain.IStockProductRepository) domain.IStockProductUseCase {
	return &usecase{
		repository: repository,
	}
}

func (usecase usecase) Create(stockProduct *dto.StockProductDTO) (*domain.StockProduct, error) {

	stock, err := usecase.repository.Create(stockProduct)

	if err != nil {
		return nil, err
	}

	return stock, nil
}
