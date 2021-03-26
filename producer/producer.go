package producer

import (
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/josielsousa/golang-gclib-alpine-confluent-kafka/config"
)

func Exec(configKafka config.KafkaClient) {
	configProducer := &kafka.ConfigMap{
		"acks":              "all",
		"bootstrap.servers": configKafka.Servers,
		"client.id":         configKafka.ClientID,
	}

	p, err := kafka.NewProducer(configProducer)
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	deliveryChan := make(chan kafka.Event, 5)
	for i := 0; i < 5; i++ {
		text := fmt.Sprintf("teste kafka %d", time.Now().UnixNano())

		message := &kafka.Message{
			Value: []byte(text),
			TopicPartition: kafka.TopicPartition{
				Partition: kafka.PartitionAny,
				Topic:     &configKafka.TopicName,
			},
		}

		err := p.Produce(message, deliveryChan)
		if err != nil {
			fmt.Printf("Failed to declare producer: %s\n", err)
			os.Exit(1)
		}

		e := <-deliveryChan
		m := e.(*kafka.Message)

		if m.TopicPartition.Error != nil {
			fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
		} else {
			fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
				*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		}
	}

	close(deliveryChan)
}
