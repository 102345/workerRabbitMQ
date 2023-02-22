package main

import (
	"context"

	"github.com/marc/workerRabbitMQ-example/adapter/postgres"
	"github.com/marc/workerRabbitMQ-example/application"
	"github.com/marc/workerRabbitMQ-example/di"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {

	ctx := context.Background()
	conn := postgres.GetConnection(ctx)
	defer conn.Close()
	queueRabbitProcessUseCase := di.ConfigQueueProcessDI(conn)
	stockProductUseCase := di.ConfigStockProductDI(conn)
	application.ProcessQueueStockProductApp(queueRabbitProcessUseCase, stockProductUseCase)

}
