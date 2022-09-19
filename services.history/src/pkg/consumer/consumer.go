package consumer

import (
	"encoding/json"
	"fmt"
	"github.com/aliaydins/microservice/service.history/src/dto"
	eventhandler "github.com/aliaydins/microservice/service.history/src/event_handler"
	history "github.com/aliaydins/microservice/service.history/src/internal"

	"github.com/aliaydins/microservice/shared/rabbitmq"
)

func RegisterConsumer(r *rabbitmq.RabbitMQ, service *history.Service, queueName string, exchangeName string) {

	var orderDto dto.OrderDTO
	_, err := r.CreateQueue(queueName, true, false)
	if err != nil {
		fmt.Println("Error occured when deliveryQueue created")
		return
	}
	r.BindQueueWithExchange(queueName, "", exchangeName)
	r.CreateMessageChannel(queueName, "history", true)

	go func() {
		for {
			message, err := r.ConsumeMessageChannel()
			if err != nil {
				fmt.Errorf("Error occured when consuming message: %s", err.Error())
				return
			}

			eventType := message.Headers["Key"]

			err = json.Unmarshal(message.Body, &orderDto)
			if err != nil {
				fmt.Println("Can't unmarshal the byte array")
				return
			}
			fmt.Printf("Message consumed from %s and consumed user information is -> %d\n", queueName, orderDto.UserId)

			if eventType == "OrderCompleted" {

				eventhandler.CreateNewHistory(service, orderDto)
			}

		}
	}()

}
