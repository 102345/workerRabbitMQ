package domain

import "github.com/marc/workerRabbitMQ-example/core/dto"

type StockProduct struct {
	ID        int32
	ProductID int32
	Quantity  int32
	Balance   int32
}

type IStockProductUseCase interface {
	Create(stockProduct *dto.StockProductDTO) (*StockProduct, error)
}

type IStockProductRepository interface {
	Create(StockProduct *dto.StockProductDTO) (*StockProduct, error)
}
