package main

import (
	"fmt"
	"github.com/aliaydins/microservice/service.wallet/src/config"
	"github.com/aliaydins/microservice/service.wallet/src/entity"
	wallet "github.com/aliaydins/microservice/service.wallet/src/internal"
	"github.com/aliaydins/microservice/service.wallet/src/pkg/server"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	config := config.LoadConfig(".")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.GetDBURL(),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		fmt.Println("Couldn't connect to the DB.")
	}
	db.Migrator().DropTable(&entity.Wallet{})
	db.AutoMigrate(&entity.Wallet{})

	repo := wallet.NewRepository(db)
	service := wallet.NewService(repo)
	handler := wallet.NewHandler(service, config.SecretKey)

	err = server.NewServer(handler.Init(), config.AppPort).Run()
	if err != nil {
		panic("Couldn't start the HTTP server.")
	}

}
