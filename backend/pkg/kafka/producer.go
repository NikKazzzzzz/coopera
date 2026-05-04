package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func NewProducer(brokers string) *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":        brokers,
		"enable.idempotence":       true, // гарантирует что дубликатов при retry не будет
		"acks":                     "all",
		"message.send.max.retries": 5,
	})
	if err != nil {
		panic(err)
	}
	return p
}

func ProduceMessage(p *kafka.Producer, topic string, key, value []byte) error {
	err := p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
		Value:          value,
	}, nil)

	if err != nil {
		return err
	}

	e := <-p.Events()
	m, ok := e.(*kafka.Message)
	if !ok {
		return nil
	}

	if m.TopicPartition.Error != nil {
		return m.TopicPartition.Error
	}

	return nil
}
