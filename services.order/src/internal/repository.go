package order

import (
	"github.com/aliaydins/microservice/service.order/src/entity"
	"github.com/aliaydins/microservice/shared/rabbitmq"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repository struct {
	db       *gorm.DB
	rabbitmq *rabbitmq.RabbitMQ
}

func NewRepository(db *gorm.DB, rabbitmq *rabbitmq.RabbitMQ) *Repository {
	db.Logger.LogMode(logger.Info)
	return &Repository{
		db:       db,
		rabbitmq: rabbitmq}
}

func (r *Repository) GetList() ([]entity.Order, []entity.Order, error) {
	var buyOrders []entity.Order
	var sellOrders []entity.Order

	err := r.db.Where("type = ?", entity.BUY).Find(&buyOrders).Error
	err = r.db.Where("type = ?", entity.SELL).Find(&sellOrders).Error
	if err != nil {
		return nil, nil, err
	}

	return buyOrders, sellOrders, nil
}

func (r *Repository) GetOrderByPriceAndQuantity(price int, quantity int, orderType entity.OrderType) (*entity.Order, error) {
	var order entity.Order

	err := r.db.Where(&entity.Order{Type: orderType, Price: price, Quantity: quantity}).First(&order).Error
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *Repository) CreateOrder(order *entity.Order) error {
	return r.db.Model(&entity.Order{}).Create(&order).Error
}

func (r *Repository) DeleteOrder(id int) error {
	return r.db.Where(&entity.Order{ID: id}).Delete(&entity.Order{}).Error
}
