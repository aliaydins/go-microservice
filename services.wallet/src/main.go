package main

import (
	"fmt"
	"github.com/aliaydins/microservice/service.wallet/src/config"
	"github.com/aliaydins/microservice/service.wallet/src/entity"
	wallet "github.com/aliaydins/microservice/service.wallet/src/internal"
	"github.com/aliaydins/microservice/service.wallet/src/pkg/consumer"
	"github.com/aliaydins/microservice/shared/rabbitmq"
	"github.com/aliaydins/microservice/shared/server"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	config := config.LoadConfig(".")
	rabbitMQOptions := rabbitmq.RabbitMQOptions{
		URL:          config.GetRabbitMQURL(),
		RetryAttempt: 3,
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.GetDBURL(),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		fmt.Println("Couldn't connect to the DB.")
	}
	db.Migrator().DropTable(&entity.Wallet{})
	db.AutoMigrate(&entity.Wallet{})

	r, err := rabbitmq.NewRabbitMQ(rabbitMQOptions)
	if err != nil {
		fmt.Println("New RabbitMQ Instance is failed")
		return
	}

	repo := wallet.NewRepository(db)
	service := wallet.NewService(repo)
	handler := wallet.NewHandler(service, config.SecretKey)

	go consumer.RegisterConsumer(r, service, config.WalletQueue, config.UserExchange)

	err = server.NewServer(handler.Init(), config.AppPort).Run()
	if err != nil {
		panic("Couldn't start the HTTP server.")
	}

}
