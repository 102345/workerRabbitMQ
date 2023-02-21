package application

import (
	"log"

	"github.com/marc/workerRabbitMQ-example/validators"
)

func ProcessQueueStockProductApp() {

	messageTest := "000000001:000000001:000000001:N"

	validate, _ := validators.ValidateMessageStockProduct(messageTest)
	log.Printf("Message Test: %s", validate)

	messageTest2 := "000001:000000001:000000001:N"

	validate2, _ := validators.ValidateMessageStockProduct(messageTest2)
	log.Printf("Message Test2: %s", validate2)

}
