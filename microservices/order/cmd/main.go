package main

import (
	"log"

	"github.com/hussammohammed/marketplace-go-microservices/microservices/order/internal/controller"
	msgBrk "github.com/hussammohammed/marketplace-go-microservices/microservices/order/internal/messageBroker"
	services "github.com/hussammohammed/marketplace-go-microservices/microservices/order/internal/service"
)

func main() {
	// message broker
	brokersUrl := []string{"localhost:9092"}
	//enum
	eventsEnum := msgBrk.NewEventsEnum()
	topicsEnum := msgBrk.NewTopicsEnum()
	// services
	msgProducer := msgBrk.NewProducerService(brokersUrl)
	msgConsumer, consumerErr := msgBrk.NewConsumerService(brokersUrl)
	orderSvc := services.NewOrderService(msgProducer, eventsEnum, topicsEnum)

	if consumerErr != nil {
		log.Fatalf("Error creating Kafka consumer: %v", consumerErr)
	}

	// controllers
	orderCtrl := controller.NewOrderController(orderSvc, msgConsumer, eventsEnum, topicsEnum)
	orderCtrl.ConsumeEvents()
}
