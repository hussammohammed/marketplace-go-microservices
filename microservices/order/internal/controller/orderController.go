package controller

import (
	"log"

	"github.com/IBM/sarama"
	msgBrk "github.com/hussammohammed/marketplace-go-microservices/microservices/order/internal/messageBroker"
)

type IOrderController interface {
	HandleEvents() error
}

type OrderController struct {
	producerService msgBrk.IProducerService
	consumerService msgBrk.IConsumerService
	eventsEnum      msgBrk.Events
}

func NewOrderController(iProducerService msgBrk.IProducerService, iConsumerService msgBrk.IConsumerService, events msgBrk.Events) *OrderController {
	return &OrderController{producerService: iProducerService, consumerService: iConsumerService, eventsEnum: events}
}

func (o *OrderController) HandleEvents() {
	topic := "order-events"
	err := o.consumerService.ConsumeEvents(topic, o.handleOrderEvents)
	if err != nil {
		log.Fatalf("Error consuming events: %v", err)
	}
}

func (o *OrderController) handleOrderEvents(msg *sarama.ConsumerMessage) {
	log.Printf("Received message: %s\n", string(msg.Value))
	testMsg := msgBrk.Event{Topic: o.eventsEnum.OrderCreated, Text: "New Order: #12345"}
	err := o.producerService.SendEvent(testMsg)
	if err != nil {
		log.Println(err.Error())
	}
	// Add your event processing logic here
}
