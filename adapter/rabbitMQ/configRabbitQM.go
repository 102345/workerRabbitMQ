package rabbitMQService

import (
	"fmt"
	"log"
	"time"

	"github.com/marc/workerRabbitMQ-example/core/domain"
	"github.com/marc/workerRabbitMQ-example/core/dto"
	"github.com/marc/workerRabbitMQ-example/validators"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

type IConfigRabbitMQService interface {
	ConfigRabbitMQ(queue string) (*amqp.Connection, *amqp.Channel, amqp.Queue, error)
	PublishMessage(conn *amqp.Connection, channel *amqp.Channel, queue amqp.Queue, message string)
	ReadQueueMessage(conn *amqp.Connection, channel *amqp.Channel, queue amqp.Queue, queueRabbitProcessUseCase domain.IQueueProcessUseCase,
		stockProductUseCase domain.IStockProductUseCase, productUseCase domain.IProductUseCase)
}

type ConfigRabbitMQService struct{}

func (config *ConfigRabbitMQService) ConfigRabbitMQ(queue string) (*amqp.Connection, *amqp.Channel, amqp.Queue, error) {

	connectionAMQPURL := viper.GetString("rabbitMQ.connectionAMQP")
	conn, err := amqp.Dial(connectionAMQPURL)
	if err != nil {
		log.Printf("Error Connection RabbitMQ: %s", err.Error())
		return nil, nil, amqp.Queue{}, err
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Error Channel RabbitMQ: %s", err.Error())
		return nil, nil, amqp.Queue{}, err
	}

	q, err := ch.QueueDeclare(
		queue, //name string,
		true,  // durable bool,
		false, // autodelete
		false, // exclusive
		false, // nowait
		nil)   // args
	if err != nil {
		log.Printf("Error Queue Declare  RabbitMQ: %s", err.Error())
		return nil, nil, amqp.Queue{}, err
	}

	ch.QueueBind(
		q.Name,       //name string,
		"",           //key string,
		"amq.fanout", //exchange string
		false,        //noWait bool,
		nil)          //args amqp.Table

	return conn, ch, q, nil

}

func (config *ConfigRabbitMQService) PublishMessage(conn *amqp.Connection, channel *amqp.Channel, queue amqp.Queue, message string) {

	msg := amqp.Publishing{
		Headers:         map[string]interface{}{},
		ContentType:     "text/plain",
		ContentEncoding: "",
		DeliveryMode:    2,
		Priority:        0,
		CorrelationId:   "",
		ReplyTo:         "",
		Expiration:      "",
		MessageId:       "messageStockProduct",
		Timestamp:       time.Time{},
		Type:            "",
		UserId:          "",
		AppId:           "go-clean-example",
		Body:            []byte(message),
	}
	channel.Publish("", queue.Name, false, false, msg)
	conn.Close()

}

func (config *ConfigRabbitMQService) ReadQueueMessage(conn *amqp.Connection, channel *amqp.Channel,
	queue amqp.Queue, queueRabbitProcessUseCase domain.IQueueProcessUseCase,
	stockProductUseCase domain.IStockProductUseCase, productUseCase domain.IProductUseCase) {

	for {
		messages, err := channel.Consume(
			queue.Name,
			"",
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			fmt.Errorf("failed to register a consumer: %w", err.Error())
		}

		for m := range messages {
			message := string(m.Body)
			process(message, queueRabbitProcessUseCase, stockProductUseCase, productUseCase)
		}
		conn.Close()
	}

}

func process(message string, queueRabbitProcessUseCase domain.IQueueProcessUseCase,
	stockProductUseCase domain.IStockProductUseCase, productUseCase domain.IProductUseCase) {

	messageValidate, stockProductDTO := validators.ValidateMessageStockProduct(message, productUseCase)
	queueProcesstDTO := dto.QueueProcessDTO{}

	if messageValidate == "" {
		stockProductUseCase.Create(&stockProductDTO)

		queueProcesstDTO.QueueMessage = message
		queueProcesstDTO.Message = messageValidate
		queueProcesstDTO.Result = "T"
		queueRabbitProcessUseCase.Create(&queueProcesstDTO)

		log.Printf("Queue processed with success: QueueMessage: %v - Messsage: %v createdAt: %s", message, messageValidate, time.Now().Format("2006-01-02 15:04:05"))

	} else {

		queueProcesstDTO.QueueMessage = message
		queueProcesstDTO.Message = messageValidate
		queueProcesstDTO.Result = "F"
		queueRabbitProcessUseCase.Create(&queueProcesstDTO)
		log.Printf("Queue processed with fail: QueueMessage: %v - Message: %v createdAt: %s", message, messageValidate, time.Now().Format("2006-01-02 15:04:05"))

	}

}
