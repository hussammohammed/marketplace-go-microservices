package messagebroker

import (
	"log"

	"github.com/IBM/sarama"
)

type IConsumerService interface {
	ConsumeEvents(topic string, eventHandler func(msg *sarama.ConsumerMessage)) error
}
type ConsumerService struct {
	brokersUrl []string
	Consumer   sarama.Consumer
}

func NewConsumerService(brokersUrl []string) (*ConsumerService, error) {
	consumer, err := connectConsumer(brokersUrl)
	if err != nil {
		return &ConsumerService{}, err
	}
	return &ConsumerService{Consumer: consumer, brokersUrl: brokersUrl}, nil
}

func connectConsumer(brokersUrl []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Create new consumer
	conn, err := sarama.NewConsumer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (c *ConsumerService) ConsumeEvents(topic string, eventHandler func(msg *sarama.ConsumerMessage)) error {
	partitionConsumer, err := c.Consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		return err
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalf("Error closing partition consumer: %v", err)
		}
	}()

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			eventHandler(msg)
		}
	}
}
