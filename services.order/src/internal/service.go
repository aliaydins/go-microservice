package order

import (
	"encoding/json"
	"fmt"
	"github.com/aliaydins/microservice/service.order/src/config"
	"github.com/aliaydins/microservice/service.order/src/entity"
	"github.com/aliaydins/microservice/service.order/src/event"
	httpclient "github.com/aliaydins/microservice/service.order/src/pkg/http_client"
	"github.com/aliaydins/microservice/shared/rabbitmq"
)

type Service struct {
	repo     *Repository
	rabbitmq *rabbitmq.RabbitMQ
}

func NewService(repo *Repository, rabbitmq *rabbitmq.RabbitMQ) *Service {
	return &Service{
		repo:     repo,
		rabbitmq: rabbitmq,
	}
}

func (s *Service) GetOrders() ([]OrderDTO, []OrderDTO, error) {

	buyOrders, sellOrders, err := s.repo.GetList()
	if err != nil {
		return nil, nil, ErrOrdersNotFound
	}

	bDto := make([]OrderDTO, 0)
	sDto := make([]OrderDTO, 0)

	for _, e := range buyOrders {
		bDto = append(bDto, mapper(&e))
	}

	for _, e := range sellOrders {
		sDto = append(sDto, mapper(&e))
	}

	return bDto, sDto, nil

}

func (s *Service) CreateOrder(reqOrder *entity.Order, token string) error {

	client := httpclient.NewCustomerClient()
	wDto, err := client.GetWalletInfo(reqOrder.UserId, token)
	if err != nil {
		return err
	}

	if reqOrder.Type == entity.BUY {

		if reqOrder.Price*reqOrder.Quantity > wDto.Wallet.USD {
			return ErrOrderDeny
		}

		order, err := s.repo.GetOrderByPriceAndQuantity(reqOrder.Price, reqOrder.Quantity, entity.SELL)
		if err != nil {
			err = s.repo.CreateOrder(reqOrder)
			if err != nil {
				return ErrOrderNotCreated
			}
			return nil
		}

		err = s.repo.DeleteOrder(order.ID)
		if err != nil {
			return ErrWhenOrderDeleted
		}

		err = PublishOrderCompletedEvent(reqOrder, order, s.rabbitmq)
		if err != nil {
			return err
		}
	} else {

		if reqOrder.Quantity > wDto.Wallet.BTC {
			return ErrOrderDeny
		}

		order, err := s.repo.GetOrderByPriceAndQuantity(reqOrder.Price, reqOrder.Quantity, entity.BUY)
		if err != nil {
			err = s.repo.CreateOrder(reqOrder)
			if err != nil {
				return ErrOrderNotCreated
			}
			return nil
		}

		err = s.repo.DeleteOrder(order.ID)
		if err != nil {
			return ErrWhenOrderDeleted
		}

		err = PublishOrderCompletedEvent(order, reqOrder, s.rabbitmq)
		if err != nil {
			return err
		}

	}

	return nil

}

func PublishOrderCompletedEvent(buyer *entity.Order, seller *entity.Order, r *rabbitmq.RabbitMQ) error {

	OrderCompletedEvent := event.OrderCompleted{
		BuyerUserId:  buyer.UserId,
		SellerUserId: seller.UserId,
		USD:          buyer.Price * buyer.Quantity,
		BTC:          buyer.Quantity,
	}

	payload, _ := json.Marshal(OrderCompletedEvent)
	err := r.Publish(payload, config.AppConfig.OrderExchange, "OrderCompleted")
	if err != nil {
		return err
	}
	fmt.Println("OrderCompleted event published ", OrderCompletedEvent)
	return nil

}
