package rabbitMQService

import (
	"fmt"
	"log"
	"time"

	"github.com/avast/retry-go"
	"github.com/marc/workerRabbitMQ-example/core/domain"
	"github.com/marc/workerRabbitMQ-example/core/dto"
	"github.com/marc/workerRabbitMQ-example/validators"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

const (
	MaxRetriesPublishMessage        = 3
	RetryDelayPublishMessage        = 1 * time.Second
	MaxRetriesDataBase              = 3
	RetryDelayDataBase              = 1 * time.Second
	DeliveryTag              uint64 = 1
)

type IConfigRabbitMQService interface {
	ConfigRabbitMQ(queue string) (*amqp.Connection, *amqp.Channel, amqp.Queue, error)
	PublishMessage(conn *amqp.Connection, channel *amqp.Channel, queue amqp.Queue, message string)
	ReadQueueMessage(conn *amqp.Connection, channel *amqp.Channel, queue amqp.Queue, queueRabbitProcessUseCase domain.IQueueProcessUseCase,
		stockProductUseCase domain.IStockProductUseCase, productUseCase domain.IProductUseCase) error
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

	err = ch.ExchangeDeclare("DeadLetterExchangeStockProduct", "fanout", true, false, false, false, nil)
	if err != nil {
		log.Printf("Error Exchange Declare RabbitMQ: %s", err.Error())
		return nil, nil, amqp.Queue{}, err
	}

	_, err = ch.QueueDeclare("DeadLetterQueueStockProduct", true, false, false, false, nil)
	if err != nil {
		log.Printf("Error Queue Dead Letter StockProduct Declare RabbitMQ: %s", err.Error())
		return nil, nil, amqp.Queue{}, err
	}

	err = ch.QueueBind("DeadLetterQueueStockProduct", "DeadLetterKeyStockProduct",
		"DeadLetterExchangeStockProduct", false, nil)
	if err != nil {
		log.Printf("Error Queue Bind Dead Letter StockProduct Declare RabbitMQ: %s", err.Error())
		return nil, nil, amqp.Queue{}, err
	}
	args := make(amqp.Table)

	args["x-dead-letter-exchange"] = "DeadLetterExchangeStockProduct"
	args["x-dead-letter-routing-key"] = "DeadLetterKeyStockProduct"

	q, err := ch.QueueDeclare(
		queue, //name string,
		true,  // durable bool,
		false, // autodelete
		false, // exclusive
		false, // nowait
		args)  // args
	if err != nil {
		log.Printf("Error Queue Declare RabbitMQ: %s", err.Error())
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
	stockProductUseCase domain.IStockProductUseCase, productUseCase domain.IProductUseCase) error {

	var errRetryFail error
	var errRetryFailDB error
	var autoAck bool = false

	for {
		errRetryFail = retry.Do(
			func() error {

				messages, err := channel.Consume(
					queue.Name,
					"",
					autoAck,
					false,
					false,
					false,
					nil,
				)
				if err != nil {
					fmt.Errorf("failed to register a consumer: %w", err)
					return err
				}

				for m := range messages {
					message := string(m.Body)
					errRetryFailDB = retry.Do(
						func() error {

							err := process(message, queueRabbitProcessUseCase, stockProductUseCase, productUseCase)
							if err != nil {
								return err
							}

							return nil

						}, retry.Attempts(MaxRetriesDataBase), retry.Delay(RetryDelayDataBase),
					)

					if errRetryFailDB != nil {
						channel.Nack(DeliveryTag, false, false)
						log.Printf("Database connection problems")
						continue
					}

				}
				conn.Close()

				channel.Ack(DeliveryTag, false)

				return nil
			}, retry.Attempts(MaxRetriesPublishMessage), retry.Delay(RetryDelayPublishMessage),
		)

		if errRetryFail != nil {
			log.Printf("Exceeded number of message consuming attempts")
			return errRetryFail
		}

	}

}

func process(message string, queueRabbitProcessUseCase domain.IQueueProcessUseCase,
	stockProductUseCase domain.IStockProductUseCase, productUseCase domain.IProductUseCase) error {

	messageValidate, stockProductDTO := validators.ValidateMessageStockProduct(message, productUseCase)
	queueProcesstDTO := dto.QueueProcessDTO{}

	if messageValidate == "" {
		_, err := stockProductUseCase.Create(&stockProductDTO)
		if err != nil {
			return err
		}

		queueProcesstDTO.QueueMessage = message
		queueProcesstDTO.Message = messageValidate
		queueProcesstDTO.Result = "T"
		_, err = queueRabbitProcessUseCase.Create(&queueProcesstDTO)
		if err != nil {
			return err
		}

		log.Printf("Queue processed with success: QueueMessage: %v - Message: %v createdAt: %s", message, messageValidate, time.Now().Format("2006-01-02 15:04:05"))

	} else {

		queueProcesstDTO.QueueMessage = message
		queueProcesstDTO.Message = messageValidate
		queueProcesstDTO.Result = "F"
		_, err := queueRabbitProcessUseCase.Create(&queueProcesstDTO)
		if err != nil {
			return err
		}
		log.Printf("Queue processed with fail: QueueMessage: %v - Message: %v createdAt: %s", message, messageValidate, time.Now().Format("2006-01-02 15:04:05"))

	}

	return nil

}
