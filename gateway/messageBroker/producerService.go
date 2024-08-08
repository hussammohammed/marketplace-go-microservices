package messagebroker

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

type IProducerService interface {
	SendEvent(message *sarama.ProducerMessage) error
}
type ProducerService struct {
	brokersUrl []string
	producer   sarama.SyncProducer
}

func NewProducerService(brokersUrl []string) *ProducerService {
	prod, err := connectProducer(brokersUrl)
	if err != nil {
		return &ProducerService{}
	}
	return &ProducerService{producer: prod, brokersUrl: brokersUrl}
}

func connectProducer(brokersUrl []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (p *ProducerService) SendEvent(message *sarama.ProducerMessage) error {
	prod, err := connectProducer(p.brokersUrl)
	if err != nil {
		return err
	}

	defer func() {
		if err := prod.Close(); err != nil {
			log.Fatalf("Error closing partition consumer: %v", err)
		}
	}()

	partition, offset, err := prod.SendMessage(message)
	if err != nil {
		return err
	}

	fmt.Printf("Message is stored in topic(%s)/eventName(%s)/partition(%d)/offset(%d)\n", message.Topic, message.Key, partition, offset)

	return nil
}
