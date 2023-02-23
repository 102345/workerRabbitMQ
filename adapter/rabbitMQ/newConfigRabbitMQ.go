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

func (app *RabbitMQApp) ConfigRabbitMQ(queue string) (*amqp.Connection, *amqp.Channel, amqp.Queue, error) {
	return app.configRabbitMQService.ConfigRabbitMQ(queue)
}

func (app *RabbitMQApp) PublishMessage(conn *amqp.Connection, channel *amqp.Channel, queue amqp.Queue, message string) {
	app.configRabbitMQService.PublishMessage(conn, channel, queue, message)
}

func (app *RabbitMQApp) ReadQueueMessage(conn *amqp.Connection, channel *amqp.Channel, queue amqp.Queue, queueRabbitProcessUseCase domain.IQueueProcessUseCase, stockProductUseCase domain.IStockProductUseCase) {
	app.configRabbitMQService.ReadQueueMessage(conn, channel, queue, queueRabbitProcessUseCase, stockProductUseCase)
}
