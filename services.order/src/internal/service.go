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
	WDto, err := client.GetWalletInfo(reqOrder.UserId, token)
	if err != nil {
		return err
	}

	if reqOrder.Type == entity.BUY {

		if reqOrder.Price*reqOrder.Quantity > WDto.Wallet.USD {
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

		sWDto, err := client.GetWalletInfo(order.UserId, token)

		deletedUSDFromBuyer := WDto.Wallet.USD - (reqOrder.Price * reqOrder.Quantity)
		addedQuantityForBuyer := WDto.Wallet.BTC + reqOrder.Quantity
		addedUSDForSeller := (reqOrder.Price * reqOrder.Quantity) + sWDto.Wallet.USD
		deletedQuantityForSeller := sWDto.Wallet.BTC - reqOrder.Quantity

		err = client.WalletUpdate(reqOrder.UserId, deletedUSDFromBuyer, addedQuantityForBuyer, token)
		if err != nil {
			return ErrWalletUpdated
		}
		err = client.WalletUpdate(sWDto.Wallet.UserId, addedUSDForSeller, deletedQuantityForSeller, token)
		if err != nil {
			return ErrWalletUpdated
		}

		err = PublishOrderCompletedEvent(reqOrder, s.rabbitmq)
		err = PublishOrderCompletedEvent(order, s.rabbitmq)
		if err != nil {
			return err
		}
	}

	if reqOrder.Type == entity.SELL {

		if reqOrder.Quantity > WDto.Wallet.BTC {
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

		bWDto, err := client.GetWalletInfo(order.UserId, token)

		deletedUSDForBuyer := bWDto.Wallet.USD - (reqOrder.Price * reqOrder.Quantity)
		addedQuantityForBuyer := bWDto.Wallet.BTC + reqOrder.Quantity
		addedUSDForSeller := (reqOrder.Price * reqOrder.Quantity) + WDto.Wallet.USD
		deletedQuantityForSeller := WDto.Wallet.BTC - reqOrder.Quantity

		err = client.WalletUpdate(reqOrder.UserId, addedUSDForSeller, deletedQuantityForSeller, token)
		if err != nil {
			return ErrWalletUpdated
		}
		err = client.WalletUpdate(bWDto.Wallet.UserId, deletedUSDForBuyer, addedQuantityForBuyer, token)
		if err != nil {
			return ErrWalletUpdated
		}

		err = PublishOrderCompletedEvent(reqOrder, s.rabbitmq)
		err = PublishOrderCompletedEvent(order, s.rabbitmq)
		if err != nil {
			return err
		}

	}

	return nil

}

func PublishOrderCompletedEvent(order *entity.Order, r *rabbitmq.RabbitMQ) error {

	OrderCompletedEvent := event.OrderCompleted{
		UserId: order.UserId,
		Type:   string(order.Type),
		USD:    order.Price,
		BTC:    order.Quantity,
		Amount: order.Price * order.Quantity,
	}

	payload, _ := json.Marshal(OrderCompletedEvent)
	err := r.Publish(payload, config.AppConfig.OrderExchange, "OrderCompleted")
	if err != nil {
		return err
	}
	fmt.Println("OrderCompleted event published ", OrderCompletedEvent)
	return nil

}
