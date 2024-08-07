package messagebroker

import (
	"fmt"

	"github.com/IBM/sarama"
)

type IProducerService interface {
	SendEvent(message Event) error
}
type ProducerService struct {
	producer sarama.SyncProducer
}

func NewProducerService(brokersUrl []string) *ProducerService {
	prod, err := connectProducer(brokersUrl)
	if err != nil {
		return &ProducerService{}
	}
	return &ProducerService{producer: prod}
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

func (p *ProducerService) SendEvent(message Event) error {
	defer p.producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: message.Topic,
		Value: sarama.StringEncoder(message.Text),
	}

	partition, offset, err := p.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	fmt.Printf("Message is stored in topic(%s)/eventName(%s)/partition(%d)/offset(%d)\n", message.Topic, message.Text, partition, offset)

	return nil
}
