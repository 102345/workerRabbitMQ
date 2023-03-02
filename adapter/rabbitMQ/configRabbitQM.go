package rabbitMQService

import (
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
	MaxRetriesDataBase        = 3
	RetryDelayDataBase        = 1 * time.Second
	DeliveryTag        uint64 = 1
)

type IConfigRabbitMQService interface {
	ConnectChannelRabbitMQ(queue string) (*amqp.Connection, *amqp.Channel, error)
	ReadQueueMessage(conn *amqp.Connection, channel *amqp.Channel, queue string, queueRabbitProcessUseCase domain.IQueueProcessUseCase,
		stockProductUseCase domain.IStockProductUseCase, productUseCase domain.IProductUseCase) error
}

type ConfigRabbitMQService struct{}

func (config *ConfigRabbitMQService) ConnectChannelRabbitMQ(queue string) (*amqp.Connection, *amqp.Channel, error) {

	connectionAMQPURL := viper.GetString("rabbitMQ.connectionAMQP")
	conn, err := amqp.Dial(connectionAMQPURL)
	if err != nil {
		log.Printf("Error Connection RabbitMQ: %s", err.Error())
		return nil, nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Error Channel RabbitMQ: %s", err.Error())
		return nil, nil, err
	}

	return conn, ch, nil

}

func (config *ConfigRabbitMQService) ReadQueueMessage(conn *amqp.Connection, channel *amqp.Channel,
	queue string, queueRabbitProcessUseCase domain.IQueueProcessUseCase,
	stockProductUseCase domain.IStockProductUseCase, productUseCase domain.IProductUseCase) error {

	var errRetryFailDB error
	var autoAck bool = false

	//Força o balanceamento de workload para as instancias de consumidor
	//prefetchcount - numeros de mensagens processadas por cada instancia de consumidor
	//global - true : por channel , false : por consumer
	channel.Qos(10, 0, false)

	for {

		messages, err := channel.Consume(
			queue,
			"",
			autoAck,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			log.Printf("Failed to register a consumer")
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
