package application

import (
	"log"

	rabbitMQService "github.com/marc/workerRabbitMQ-example/adapter/rabbitMQ"
	"github.com/marc/workerRabbitMQ-example/core/domain"
)

func ProcessQueueStockProductApp(queueRabbitProcessUseCase domain.IQueueProcessUseCase,
	stockProductUseCase domain.IStockProductUseCase, productUseCase domain.IProductUseCase) {

	//productid:quantitity:balance:signalbalance
	//messageTest := "000000002:000000001:000000001:N"

	configRabbitMQServiceApp := rabbitMQService.NewRabbitMQApp(&rabbitMQService.ConfigRabbitMQService{})
	conn, channel, queue, err := configRabbitMQServiceApp.ConfigRabbitMQ("queueStockProduct")

	if err != nil {
		log.Println("Error ConfigRAbbitMQ: %s ", err.Error())
		return
	}

	configRabbitMQServiceApp.ReadQueueMessage(conn, channel, queue, queueRabbitProcessUseCase, stockProductUseCase, productUseCase)

}
