package application

import (
	"log"

	rabbitMQService "github.com/marc/workerRabbitMQ-example/adapter/rabbitMQ"
	"github.com/marc/workerRabbitMQ-example/core/domain"
)

type StockProductQueueProcessor struct {
	queueRabbitProcessUseCase domain.IQueueProcessUseCase
	stockProductUseCase       domain.IStockProductUseCase
	productUseCase            domain.IProductUseCase
}

func NewStockProductQueueProcessor(queueRabbitProcessUseCase domain.IQueueProcessUseCase,
	stockProductUseCase domain.IStockProductUseCase, productUseCase domain.IProductUseCase) *StockProductQueueProcessor {
	return &StockProductQueueProcessor{queueRabbitProcessUseCase, stockProductUseCase, productUseCase}
}

func ProcessQueueStockProductApp(p *StockProductQueueProcessor) {

	//productid:quantitity:balance:signalbalance
	//messageTest := "000000002:000000001:000000001:N"

	configRabbitMQServiceApp := rabbitMQService.NewRabbitMQApp(&rabbitMQService.ConfigRabbitMQService{})
	conn, channel, queue, err := configRabbitMQServiceApp.ConfigRabbitMQ("queueStockProduct")

	if err != nil {
		log.Println("Error ConfigRAbbitMQ: %s ", err.Error())
		return
	}

	configRabbitMQServiceApp.ReadQueueMessage(conn, channel, queue, p.queueRabbitProcessUseCase, p.stockProductUseCase, p.productUseCase)

}
