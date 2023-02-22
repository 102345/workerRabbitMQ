package application

import (
	"github.com/marc/workerRabbitMQ-example/core/domain"
	"github.com/marc/workerRabbitMQ-example/core/dto"
	"github.com/marc/workerRabbitMQ-example/validators"
)

func ProcessQueueStockProductApp(queueRabbitProcessUseCase domain.IQueueProcessUseCase, stockProductUseCase domain.IStockProductUseCase) {

	//productid:quantitity:balance:signalbalance
	messageTest := "000000002:000000001:000000001:N"

	registerQueueStockProduct(messageTest, queueRabbitProcessUseCase, stockProductUseCase)

	messageTest2 := "000001:000000001:000000001:N"

	registerQueueStockProduct(messageTest2, queueRabbitProcessUseCase, stockProductUseCase)

}

func registerQueueStockProduct(message string, queueRabbitProcessUseCase domain.IQueueProcessUseCase, stockProductUseCase domain.IStockProductUseCase) {

	messageValidate, stockProductDTO := validators.ValidateMessageStockProduct(message)
	queueProcesstDTO := dto.QueueProcessDTO{}
	if messageValidate == "" {
		stockProductUseCase.Create(&stockProductDTO)

		queueProcesstDTO.Message = message
		queueProcesstDTO.Result = "T"
		queueRabbitProcessUseCase.Create(&queueProcesstDTO)

	} else {

		queueProcesstDTO.Message = messageValidate
		queueProcesstDTO.Result = "F"
		queueRabbitProcessUseCase.Create(&queueProcesstDTO)

	}

}
