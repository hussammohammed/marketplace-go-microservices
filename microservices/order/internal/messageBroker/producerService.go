package messagebroker

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

var config *sarama.Config

type IProducerService interface {
	SendEvent(message Event) error
}
type ProducerService struct {
	brokersUrl []string
}

func NewProducerService(brokersUrl []string) *ProducerService {
	config = sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	return &ProducerService{brokersUrl: brokersUrl}
}

func connectProducer(brokersUrl []string) (sarama.SyncProducer, error) {
	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (p *ProducerService) SendEvent(message Event) error {
	prod, err := connectProducer(p.brokersUrl)
	if err != nil {
		return err
	}

	defer func() {
		if err := prod.Close(); err != nil {
			log.Fatalf("Error closing partition consumer: %v", err)
		}
	}()

	msg := &sarama.ProducerMessage{
		Topic: message.Topic,
		Value: sarama.StringEncoder(message.Text),
	}

	partition, offset, err := prod.SendMessage(msg)
	if err != nil {
		return err
	}

	fmt.Printf("Message is stored in topic(%s)/eventName(%s)/partition(%d)/offset(%d)\n", message.Topic, message.Text, partition, offset)

	return nil
}
