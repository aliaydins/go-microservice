package main

import (
	"fmt"
	"github.com/aliaydins/microservice/services.user/src/config"
	"github.com/aliaydins/microservice/services.user/src/entity"
	user "github.com/aliaydins/microservice/services.user/src/internal"
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

	db.Migrator().DropTable(&entity.User{})
	db.AutoMigrate(&entity.User{})

	r, err := rabbitmq.NewRabbitMQ(rabbitMQOptions)
	if err != nil {
		fmt.Println("New RabbitMQ Instance is failed")
		return
	}

	err = r.CreateExchange(config.UserExchange, "fanout", true, false)
	if err != nil {
		fmt.Println("Error occured when exchange created")
		return
	}

	repo := user.NewRepository(db, r)
	service := user.NewService(repo, r)
	handler := user.NewHandler(service, config.SecretKey)

	err = server.NewServer(handler.Init(), config.AppPort).Run()
	if err != nil {
		fmt.Errorf("error %v", err.Error())
		panic("Couldn't start the HTTP server.")
	}
}
