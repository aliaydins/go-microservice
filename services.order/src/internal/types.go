package order

import "github.com/aliaydins/microservice/service.order/src/entity"

type OrderDTO struct {
	ID       int              `json:"id"`
	UserId   int              `json:"user_id"`
	Type     entity.OrderType `json:"type"`
	Price    int              `json:"price"`
	Quantity int              `json:"quantity"`
}

func mapper(order *entity.Order) OrderDTO {
	dto := OrderDTO{
		ID:       order.ID,
		UserId:   order.UserId,
		Type:     order.Type,
		Price:    order.Price,
		Quantity: order.Quantity,
	}
	return dto
}
