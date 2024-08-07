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
}

func NewOrderService(iProducerService msgBro.IProducerService) *OrderService {
	return &OrderService{producerService: iProducerService}
}

func (o *OrderService) CreateOrder(orderReq OrderReq) error {
	event := msgBro.Event{
		Topic: "order-events",
		Text:  fmt.Sprintf("init-order#%v", orderReq.UserId),
	}
	err := o.producerService.SendEvent(event)
	if err != nil {
		return err
	}
	return nil
}
