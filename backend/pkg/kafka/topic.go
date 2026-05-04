package kafka

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"time"
)

func CreateTopic(brokers string, topicName string, numPartitions int, replicationFactor int) error {
	admin, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": brokers})
	if err != nil {
		return err
	}
	defer admin.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = admin.CreateTopics(ctx, []kafka.TopicSpecification{
		{
			Topic:             topicName,
			NumPartitions:     numPartitions,
			ReplicationFactor: replicationFactor,
			Config: map[string]string{
				"cleanup.policy":    "delete",
				"retention.ms":      "86400000", // 1 день
				"segment.bytes":     "1073741824",
				"max.message.bytes": "2000000",
			},
		},
	})

	return nil
}
