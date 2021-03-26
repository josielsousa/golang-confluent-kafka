package main

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"

	"github.com/kelseyhightower/envconfig"

	"github.com/josielsousa/golang-gclib-alpine-confluent-kafka/config"
	"github.com/josielsousa/golang-gclib-alpine-confluent-kafka/consumer"
	"github.com/josielsousa/golang-gclib-alpine-confluent-kafka/producer"
)

func loadConfig() config.KafkaClient {
	noPrefix := ""
	var kafkaClient config.KafkaClient

	err := envconfig.Process(noPrefix, &kafkaClient)
	if err != nil {
		fmt.Printf("Failed initialize config envs: %s\n", err)
		os.Exit(1)
	}

	return kafkaClient
}

// init - Disable memory profile rate.
func init() {
	runtime.MemProfileRate = 0
}

func main() {
	fmt.Println("Starting Application...")
	configKafka := loadConfig()

	jsonBytes, _ := json.Marshal(configKafka)
	fmt.Println("Config: ", string(jsonBytes))

	forever := make(chan bool)
	go func() {
		producer.Exec(configKafka)
		consumer.Exec(configKafka)
	}()

	<-forever
}
