package rabbitMQService

import (
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
	ReadQueueMessage(conn *amqp.Connection, channel *amqp.Channel, queue amqp.Queue, queueRabbitProcessUseCase domain.IQueueProcessUseCase, stockProductUseCase domain.IStockProductUseCase)
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
		ContentType:     "",
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

func (config *ConfigRabbitMQService) ReadQueueMessage(conn *amqp.Connection, channel *amqp.Channel, queue amqp.Queue, queueRabbitProcessUseCase domain.IQueueProcessUseCase, stockProductUseCase domain.IStockProductUseCase) {

	for {
		messages, _ := channel.Consume(
			queue.Name,
			"",
			true,
			false,
			false,
			false,
			nil,
		)

		for m := range messages {
			message := string(m.Body)
			registerQueueStockProduct(message, queueRabbitProcessUseCase, stockProductUseCase)
		}
		conn.Close()
	}

}

func registerQueueStockProduct(message string, queueRabbitProcessUseCase domain.IQueueProcessUseCase,
	stockProductUseCase domain.IStockProductUseCase) {

	messageValidate, stockProductDTO := validators.ValidateMessageStockProduct(message)
	queueProcesstDTO := dto.QueueProcessDTO{}

	now := time.Now()
	dateHourLocal := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), now.Location())
	if messageValidate == "" {
		stockProductUseCase.Create(&stockProductDTO)

		queueProcesstDTO.QueueMessage = message
		queueProcesstDTO.Message = messageValidate
		queueProcesstDTO.Result = "T"
		queueRabbitProcessUseCase.Create(&queueProcesstDTO)

		log.Println("Queue processed with success: QueueMessage: %v - Messsage: %v createdAt: %d", message, messageValidate, dateHourLocal)

	} else {

		queueProcesstDTO.QueueMessage = message
		queueProcesstDTO.Message = messageValidate
		queueProcesstDTO.Result = "F"
		queueRabbitProcessUseCase.Create(&queueProcesstDTO)
		log.Println("Queue processed with fail: QueueMessage: %v - Message: %v createdAt: %d", message, messageValidate, dateHourLocal)

	}

}
