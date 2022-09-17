package rabbitmq

import (
	"errors"
	"fmt"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn         *amqp.Connection
	channel      *amqp.Channel
	connURL      string
	errCh        <-chan *amqp.Error
	messageChan  <-chan amqp.Delivery
	retryAttempt int
}
type RabbitMQOptions struct {
	URL          string
	RetryAttempt int
}

func NewRabbitMQ(options RabbitMQOptions) (*RabbitMQ, error) {
	r := &RabbitMQ{
		connURL:      options.URL,
		retryAttempt: options.RetryAttempt,
	}

	if err := r.connect(); err != nil {
		return nil, err
	}

	return r, nil
}

func (r *RabbitMQ) connect() error {
	var err error

	r.conn, err = amqp.Dial(r.connURL)
	if err != nil {
		fmt.Errorf("Error occured when creating connection: %s", err)
		return err
	}

	r.channel, err = r.conn.Channel()
	if err != nil {
		fmt.Errorf("Error occured when connection to channel: %s", err)
		return err
	}

	r.errCh = r.conn.NotifyClose(make(chan *amqp.Error))

	return nil
}

func (r *RabbitMQ) reconnect() {
	attempt := r.retryAttempt

	for attempt != 0 {
		fmt.Errorf("Attempting rabbitmq reconnection")
		if err := r.connect(); err != nil {
			attempt--
			fmt.Errorf("Rabbitmq retry connection error:%s", err.Error())
			continue
		}
		return
	}

	if attempt == 0 {
		fmt.Errorf("Rabbitmq retry connection is failed")
	}
}

func (r *RabbitMQ) CreateExchange(name string, kind string, durable bool, autoDelete bool) error {
	if err := r.channel.ExchangeDeclare(name, kind, durable, autoDelete, false, false, nil); err != nil {
		fmt.Errorf("Error occured when creating exchange: %s", err.Error())
		return err
	}

	return nil
}

func (r *RabbitMQ) CreateQueue(name string, durable bool, autoDelete bool) (amqp.Queue, error) {
	queue, err := r.channel.QueueDeclare(name, durable, autoDelete, false, false, nil)
	if err != nil {
		fmt.Errorf("Error occured when creating queue: %s", err.Error())
		return amqp.Queue{}, err
	}

	return queue, nil
}

func (r *RabbitMQ) BindQueueWithExchange(queueName string, key string, exchangeName string) error {
	if err := r.channel.QueueBind(queueName, key, exchangeName, false, nil); err != nil {
		fmt.Errorf("Error occured when queue binding: %s", err.Error())
		return err
	}

	return nil
}

func (r *RabbitMQ) CreateMessageChannel(queue string, consumer string, autoAck bool) error {
	var err error
	r.messageChan, err = r.channel.Consume(queue, consumer, autoAck, false, false, false, nil)
	if err != nil {
		fmt.Errorf("Error when consuming message: %s", err.Error())
		return err
	}

	return nil
}

func (r *RabbitMQ) ConsumeMessageChannel() (amqp.Delivery, error) {

	select {
	case err := <-r.errCh:
		fmt.Errorf("Rabbitmq comes error from notifyCloseChan:%s", err.Error())
		r.reconnect()
	case msg := <-r.messageChan:
		return msg, nil
	}

	return amqp.Delivery{}, nil

}

func (r *RabbitMQ) Publish(body []byte, exchaneName string, eventType string) error {

	if err := r.channel.Publish(exchaneName, "", false, false,
		amqp.Publishing{
			Headers: amqp.Table{
				"Key": eventType,
			},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            body,
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
		},
	); err != nil {
		return errors.New("Error occured when deliver message to exchange/queue")
	}

	return nil

}
