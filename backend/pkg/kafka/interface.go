package kafka

import "context"

type Producer interface {
	Produce(ctx context.Context, key, value []byte) error
	ProduceAsync(key, value []byte)
	Close()
}

type Consumer interface {
	Consume(ctx context.Context, handler func(key, value []byte) error) error
	ConsumeAsync(handler func(key, value []byte) error)
	Close()
}
