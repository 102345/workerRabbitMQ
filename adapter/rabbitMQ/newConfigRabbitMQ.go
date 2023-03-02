package rabbitMQService

import (
	"github.com/marc/workerRabbitMQ-example/core/domain"
	"github.com/streadway/amqp"
)

type RabbitMQApp struct {
	configRabbitMQService IConfigRabbitMQService
}

func NewRabbitMQApp(configRabbitMQService IConfigRabbitMQService) *RabbitMQApp {

	return &RabbitMQApp{
		configRabbitMQService: configRabbitMQService,
	}

}

func (app *RabbitMQApp) ConnectChannelRabbitMQ(queue string) (*amqp.Connection, *amqp.Channel, error) {
	return app.configRabbitMQService.ConnectChannelRabbitMQ(queue)
}

func (app *RabbitMQApp) ReadQueueMessage(conn *amqp.Connection, channel *amqp.Channel, queue string, queueRabbitProcessUseCase domain.IQueueProcessUseCase,
	stockProductUseCase domain.IStockProductUseCase, productUseCase domain.IProductUseCase) error {
	return app.configRabbitMQService.ReadQueueMessage(conn, channel, queue, queueRabbitProcessUseCase, stockProductUseCase, productUseCase)
}
