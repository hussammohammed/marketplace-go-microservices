package controller

import (
	"log"

	"github.com/IBM/sarama"
	msgBrk "github.com/hussammohammed/marketplace-go-microservices/microservices/order/internal/messageBroker"
	services "github.com/hussammohammed/marketplace-go-microservices/microservices/order/internal/service"
)

type IOrderController interface {
	ConsumeEvents() error
}

type OrderController struct {
	consumerService msgBrk.IConsumerService
	orderService    services.IOrderService
	eventsEnum      msgBrk.Events
	topicsEnum      msgBrk.Topics
}

func NewOrderController(iIOrderService services.IOrderService, iConsumerService msgBrk.IConsumerService, events msgBrk.Events, topics msgBrk.Topics) *OrderController {
	return &OrderController{orderService: iIOrderService, consumerService: iConsumerService, eventsEnum: events, topicsEnum: topics}
}

func (o *OrderController) ConsumeEvents() {
	// add listner for order events
	orderErr := o.consumerService.ConsumeEvents(o.topicsEnum.OrderEvents, o.handleOrderEvents)
	if orderErr != nil {
		log.Fatalf("Error consuming %v: %v", o.topicsEnum.OrderEvents, orderErr)
	}

	// add listner for user events
	userErr := o.consumerService.ConsumeEvents(o.topicsEnum.UserEvents, o.handleUserEvents)
	if userErr != nil {
		log.Fatalf("Error consuming %v: %v", o.topicsEnum.UserEvents, userErr)
	}
}

func (o *OrderController) handleOrderEvents(msg *sarama.ConsumerMessage) {
	ordErr := o.orderService.HandleOrderEvents(msg)
	if ordErr != nil {
		log.Printf("error at handling order events: %v", ordErr.Error())
	}
}

func (o *OrderController) handleUserEvents(msg *sarama.ConsumerMessage) {
	log.Printf("Received message for %v: %s\n", o.topicsEnum.UserEvents, string(msg.Value))
}
