package consumer

import (
	"encoding/json"
	"fmt"
	"github.com/aliaydins/microservice/service.wallet/src/dto"
	eventhandler "github.com/aliaydins/microservice/service.wallet/src/event_handler"
	wallet "github.com/aliaydins/microservice/service.wallet/src/internal"
	"github.com/aliaydins/microservice/shared/rabbitmq"
)

func RegisterConsumer(r *rabbitmq.RabbitMQ, service *wallet.Service, queueName string, exchangeName string) {

	var userDto dto.UserDTO
	_, err := r.CreateQueue(queueName, true, false)
	if err != nil {
		fmt.Println("Error occured when deliveryQueue created")
		return
	}
	r.BindQueueWithExchange(queueName, "", exchangeName)
	r.CreateMessageChannel(queueName, "wallet", true)

	go func() {
		for {
			message, err := r.ConsumeMessageChannel()
			if err != nil {
				fmt.Errorf("Error occured when consuming message: %s", err.Error())
				return
			}

			eventType := message.Headers["Key"]

			err = json.Unmarshal(message.Body, &userDto)
			if err != nil {
				fmt.Println("Can't unmarshal the byte array")
				return
			}
			fmt.Printf("Message consumed from %s and consumed user information is -> %d %s \n ", queueName, userDto.UserId, userDto.Email)

			if eventType == "UserCreated" {
				eventhandler.CreateWallet(service, userDto.UserId)
			}

		}
	}()

}
