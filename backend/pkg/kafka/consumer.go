package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func NewConsumer(brokers, group, topic string) *kafka.Consumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":        brokers,
		"group.id":                 group,
		"auto.offset.reset":        "earliest",
		"enable.auto.commit":       false, // ключевой момент: ручной commit
		"enable.auto.offset.store": false, // сами управляем offset store
	})
	if err != nil {
		panic(err)
	}

	if err := c.Subscribe(topic, nil); err != nil {
		panic(err)
	}

	return c
}

//func StartConsumer(c *kafka.Consumer, producer *kafka.Producer) {
//	ctx := context.Background()
//
//	for {
//		ev := c.Poll(200)
//		if ev == nil {
//			continue
//		}
//
//		switch msg := ev.(type) {
//
//		case *kafka.Message:
//			// 1. Выполняем основную бизнес-логику
//			err := processMessage(ctx, msg)
//			if err != nil {
//				log.Printf("process failed: %v — rollback", err)
//				continue // rollback = offset не сохраняем
//			}
//
//			// 2. Если нужно — публикуем результат
//			err = ProduceMessage(producer, "result-topic", msg.Key, msg.Value)
//			if err != nil {
//				log.Printf("produce failed: %v — rollback", err)
//				continue // rollback = offset не сохраняем
//			}
//
//			// 3. Если всё успешно — фиксируем offset
//			_, err = c.StoreMessage(msg) // локальный store
//			if err != nil {
//				log.Printf("store failed: %v — rollback", err)
//				continue
//			}
//
//			_, err = c.Commit() // отправляем commit в Kafka
//			if err != nil {
//				log.Printf("commit failed: %v — rollback", err)
//			}
//
//		case kafka.Error:
//			log.Printf("kafka error: %v", msg)
//		}
//	}
//}
