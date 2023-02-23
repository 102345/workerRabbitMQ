package productUseCase

import (
	"github.com/marc/workerRabbitMQ-example/core/domain"
)

type usecase struct {
	repository domain.IProductRepository
}

// New returns contract implementation of ProductUseCase
func New(repository domain.IProductRepository) domain.IProductUseCase {
	return &usecase{
		repository: repository,
	}
}

func (usecase usecase) FindById(id int64) (domain.Product, error) {

	product, err := usecase.repository.FindById(id)

	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}
