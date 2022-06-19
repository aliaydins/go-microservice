package main

import (
	"fmt"
	"github.com/aliaydins/microservice/services.user/src/config"
	"github.com/aliaydins/microservice/services.user/src/entity"
	user "github.com/aliaydins/microservice/services.user/src/internal"
	"github.com/aliaydins/microservice/services.user/src/pkg/server"
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
	db.Migrator().DropTable(&entity.User{})
	db.AutoMigrate(&entity.User{})

	repo := user.NewRepository(db)
	service := user.NewService(repo)
	handler := user.NewHandler(service)

	err = server.NewServer(handler.Init(), config.AppPort).Run()
	if err != nil {
		panic("Couldn't start the HTTP server.")
	}
}
