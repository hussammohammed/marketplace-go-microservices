package service

import (
	"encoding/json"
	"log"

	"github.com/IBM/sarama"

	msgBrk "github.com/hussammohammed/marketplace-go-microservices/microservices/order/internal/messageBroker"
	models "github.com/hussammohammed/marketplace-go-microservices/microservices/order/internal/model"
)

type IOrderService interface {
	HandleOrderEvents(msg *sarama.ConsumerMessage) error
	HandleUserEvents(msg *sarama.ConsumerMessage) error
	CreateOrder(msgData []byte) error
	UpdateOrder(msgData []byte) error
}
type OrderService struct {
	topicsEnum      msgBrk.Topics
	eventsEnum      msgBrk.Events
	producerService msgBrk.IProducerService
}

func NewOrderService(iProducerService msgBrk.IProducerService, events msgBrk.Events, topics msgBrk.Topics) *OrderService {
	return &OrderService{producerService: iProducerService, topicsEnum: topics, eventsEnum: events}
}

func (o *OrderService) HandleOrderEvents(msg *sarama.ConsumerMessage) error {
	log.Printf("Received message for %v: %s\n", o.topicsEnum.OrderEvents, string(msg.Value))
	if msg.Key != nil && string(msg.Key) == o.eventsEnum.OrderReceived {
		createOrderErr := o.CreateOrder(msg.Value)
		if createOrderErr != nil {
			log.Printf("error at creating order: %v", createOrderErr.Error())
			return createOrderErr
		}
	}

	if msg.Key != nil && string(msg.Key) == o.eventsEnum.OrderUpdated {
		updateOrderErr := o.UpdateOrder(msg.Value)
		if updateOrderErr != nil {
			log.Printf("error at updating order: %v", updateOrderErr.Error())
			return updateOrderErr
		}
	}
	return nil
}

func (o *OrderService) HandleUserEvents(msg *sarama.ConsumerMessage) error {
	log.Printf("Received message for %v: %s\n", o.topicsEnum.UserEvents, string(msg.Value))
	return nil
}

func (o *OrderService) CreateOrder(msgData []byte) error {
	// deserialize message value to obj
	order := &models.Order{}
	unMarErr := json.Unmarshal(msgData, order)
	if unMarErr != nil {
		log.Printf("failed to unmarshal object at %v event", o.eventsEnum.OrderCreated)
		return unMarErr
	}
	// create order
	order.Id = 123
	data, marErr := json.Marshal(order)
	if marErr != nil {
		log.Println("failed to marchal order object at creation process")
		return marErr
	}

	// fire event for order creation
	msg := &sarama.ProducerMessage{
		Topic: o.topicsEnum.OrderEvents,
		Key:   sarama.StringEncoder(o.eventsEnum.OrderCreated),
		Value: sarama.ByteEncoder(data),
	}
	sendEventErr := o.producerService.SendEvent(msg)
	if sendEventErr != nil {
		log.Println(sendEventErr.Error())
	}
	return nil
}

func (o *OrderService) UpdateOrder(msgData []byte) error {
	return nil
}
