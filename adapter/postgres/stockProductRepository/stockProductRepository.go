package stockProductRepository

import (
	"context"

	"github.com/marc/workerRabbitMQ-example/adapter/postgres"
	"github.com/marc/workerRabbitMQ-example/core/domain"
	"github.com/marc/workerRabbitMQ-example/core/dto"
)

type repository struct {
	db postgres.PoolInterface
}

// New returns contract implementation of StockProductRepository
func New(db postgres.PoolInterface) domain.IStockProductRepository {
	return &repository{
		db: db,
	}
}

func (repository repository) Create(stockProduct *dto.StockProductDTO) (*domain.StockProduct, error) {

	ctx := context.Background()
	stockProductRet := domain.StockProduct{}
	err := repository.db.QueryRow(
		ctx,
		"INSERT INTO StockProduct (productid,quantity,balance) VALUES ($1, $2, $3) returning *",
		stockProduct.ProductID,
		stockProduct.Quantity,
		stockProduct.Balance,
	).Scan(
		&stockProductRet.ID,
		&stockProductRet.Quantity,
		&stockProductRet.Balance,
	)
	if err != nil {
		return nil, err
	}
	return &stockProductRet, nil
}
