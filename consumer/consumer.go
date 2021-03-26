package consumer

import (
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/josielsousa/golang-gclib-alpine-confluent-kafka/config"
)

func Exec(configKafka config.KafkaClient) {
	configConsumer := &kafka.ConfigMap{
		"group.id":  configKafka.ClientID,
		"client.id": configKafka.ClientID,

		"enable.auto.commit": false,
		"auto.offset.reset":  "smallest",
		"bootstrap.servers":  configKafka.Servers,
	}

	// intialize the writer with the broker addresses, and the topic
	consumer, err := kafka.NewConsumer(configConsumer)
	if err != nil {
		fmt.Printf("Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	err = consumer.SubscribeTopics([]string{configKafka.TopicName}, nil)
	if err != nil {
		fmt.Printf("Failed subscribe topic consumer: %s\n", err)
		os.Exit(1)
	}

	run := true
	timeout := 5 * time.Second

	for run {
		ev := consumer.Poll(int(timeout.Milliseconds()))
		switch e := ev.(type) {
		case *kafka.Message:
			fmt.Printf("--> Message: %s\n", string(e.Value))
			consumer.Commit()
		case kafka.PartitionEOF:
			fmt.Printf("--> Reached %v\n", e)
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false
		default:
			fmt.Println("[*] Waiting for messages")
		}
	}

	consumer.Close()
}
