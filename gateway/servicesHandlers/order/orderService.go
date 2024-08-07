package order

import (
	"fmt"

	msgBro "github.com/hussammohammed/marketplace-go-microservices/gateway/messageBroker"
)

type IOrderService interface {
	CreateOrder(orderReq OrderReq) error
}
type OrderService struct {
	producerService msgBro.IProducerService
	eventsEnum      msgBro.Events
}

func NewOrderService(iProducerService msgBro.IProducerService, events msgBro.Events) *OrderService {
	return &OrderService{producerService: iProducerService, eventsEnum: events}
}

func (o *OrderService) CreateOrder(orderReq OrderReq) error {
	event := msgBro.Event{
		Topic: "order-events",
		Text:  fmt.Sprintf("%v#%v", o.eventsEnum.OrderCreated, orderReq.UserId),
	}
	err := o.producerService.SendEvent(event)
	if err != nil {
		return err
	}
	return nil
}
