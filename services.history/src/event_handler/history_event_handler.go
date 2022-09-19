package eventhandler

import (
	"fmt"
	"github.com/aliaydins/microservice/service.history/src/dto"
	"github.com/aliaydins/microservice/service.history/src/entity"
	history "github.com/aliaydins/microservice/service.history/src/internal"
)

func CreateNewHistory(service *history.Service, order dto.OrderDTO) {

	history := entity.History{
		UserId: order.UserId,
		Type:   order.Type,
		USD:    order.USD,
		BTC:    order.BTC,
		Amount: order.Amount,
	}

	err := service.Create(&history)
	if err != nil {
		fmt.Errorf("Error message -> %v", err.Error())
	}

	fmt.Println("Order history created successfully with userId -> ", order.UserId)
}
