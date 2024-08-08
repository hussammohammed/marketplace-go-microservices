package main

import (
	"log"

	"github.com/hussammohammed/marketplace-go-microservices/microservices/order/internal/controller"
	msgBrk "github.com/hussammohammed/marketplace-go-microservices/microservices/order/internal/messageBroker"
)

func main() {
	// message broker
	brokersUrl := []string{"localhost:9092"}
	//enum
	eventsEnum := msgBrk.NewEventsEnum()
	// services
	msgProducer := msgBrk.NewProducerService(brokersUrl)
	// controllers
	msgConsumer, consumerErr := msgBrk.NewConsumerService(brokersUrl)

	if consumerErr != nil {
		log.Fatalf("Error creating Kafka consumer: %v", consumerErr)
	}
	defer func() {
		if err := msgConsumer.Consumer.Close(); err != nil {
			log.Fatalf("Error closing Kafka consumer: %v", err)
		}
	}()

	orderCtrl := controller.NewOrderController(msgProducer, msgConsumer, eventsEnum)
	orderCtrl.HandleEvents()
}
