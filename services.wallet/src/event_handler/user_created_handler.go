package eventhandler

import (
	"fmt"
	wallet "github.com/aliaydins/microservice/service.wallet/src/internal"
)

func CreateWallet(service *wallet.Service, userId int) {

	fmt.Println("User id is -> ", userId)

	err := service.CreateWallet(userId)
	if err != nil {
		fmt.Errorf("Error message -> %v", err.Error())
	}

	fmt.Println("User wallet created successfully with userId -> ", userId)
}
