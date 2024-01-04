package kafkaconsumer

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

func ReadMessage() {
	mechanism := plain.Mechanism{
		Username: "user1",
		Password: "hello",
	}
	dialer := &kafka.Dialer{
		Timeout:       10 * time.Second,
		DualStack:     true,
		SASLMechanism: mechanism,
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"kafka-broker-headless:9092"},
		Topic:       "my-topic",
		GroupID:     "myGroup",
		StartOffset: kafka.FirstOffset,
		Dialer:      dialer,
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("could not read message: %v", err)
		}

		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatalf("failed to close reader: %s", err)
	}
}
