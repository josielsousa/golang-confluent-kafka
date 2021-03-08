package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// HTTPHandler - create a handler struct
type HTTPHandler struct{}

// implement `ServeHTTP` method on `HTTPHandler` struct
func (h HTTPHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// create response binary data
	data := []byte("Hello World!") // slice of bytes
	// write `data` to response
	res.Write(data)
}

var topicName = "teste-golang-gclib-alpine"

func producer() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092,localhost:9092",
		"client.id":         "test-golang-gclib",
		"acks":              "all"})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	deliveryChan := make(chan kafka.Event, 10000)

	for i := 0; i < 5; i++ {
		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topicName,
				Partition: kafka.PartitionAny,
			},
			Value: []byte(fmt.Sprintf("teste kafka %d", time.Now().UnixNano()))},
			deliveryChan,
		)

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

func consumer() {

	// intialize the writer with the broker addresses, and the topic
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  "localhost:9092,localhost:9092",
		"group.id":           "test-golang-gclib",
		"enable.auto.commit": false,
		"client.id":          "test-golang-gclib",
		"auto.offset.reset":  "smallest"},
	)

	if err != nil {
		fmt.Printf("Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	topics := []string{topicName}
	err = consumer.SubscribeTopics(topics, nil)

	// run := true
	for i := 0; i < 50; i++ {

		// for run == true {
		ev := consumer.Poll(2000)
		switch e := ev.(type) {
		case *kafka.Message:
			fmt.Printf("%% Message on %s:\n%s\n",
				e.TopicPartition, string(e.Value))
		case kafka.PartitionEOF:
			fmt.Printf("%% Reached %v\n", e)
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			// run = false
		default:
			fmt.Printf("Ignored %v\n", e)
		}
	}

	consumer.Close()
}

func main() {
	// create a new handler
	handler := HTTPHandler{}

	fmt.Println("Starting API...")

	producer()
	consumer()

	// listen and serve
	http.ListenAndServe(":9000", handler)
}
