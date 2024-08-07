package main

import (
	"log"

	"github.com/hussammohammed/marketplace-go-microservices/microservices/order/internal/model"
	producer "github.com/hussammohammed/marketplace-go-microservices/microservices/order/internal/service"
)

func main() {
	msgProducer := producer.NewProducerService()
	testMsg := model.ProducerMsg{Topic: "order-events", Text: "New Order: #12345"}
	err := msgProducer.PushMessageToQueue(testMsg)
	if err != nil {
		log.Println(err.Error())
	}
}
