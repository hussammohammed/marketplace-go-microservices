package order

import (
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
	msgBro "github.com/hussammohammed/marketplace-go-microservices/gateway/messageBroker"
)

type IOrderService interface {
	CreateOrder(orderReq OrderReq) error
}
type OrderService struct {
	producerService msgBro.IProducerService
	topicsEnum      msgBro.Topics
	eventsEnum      msgBro.Events
}

func NewOrderService(iProducerService msgBro.IProducerService, topics msgBro.Topics, events msgBro.Events) *OrderService {
	return &OrderService{producerService: iProducerService, topicsEnum: topics, eventsEnum: events}
}

func (o *OrderService) CreateOrder(orderReq OrderReq) error {
	data, marErr := json.Marshal(&orderReq)
	if marErr != nil {
		log.Println("failed to marchal order object at receive order")
		return marErr
	}
	msg := &sarama.ProducerMessage{
		Topic: o.topicsEnum.OrderEvents,
		Key:   sarama.StringEncoder(o.eventsEnum.OrderReceived),
		Value: sarama.ByteEncoder(data),
	}
	err := o.producerService.SendEvent(msg)
	if err != nil {
		return err
	}
	return nil
}
