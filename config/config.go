package config

type KafkaClient struct {
	TopicName string `envconfig:"KAFKA_TOPIC" default:"alpine-topic-msgs"`
	ClientID  string `envconfig:"KAFKA_CLIENT_ID" default:"app.test.golang.gclib"`
	Servers   string `envconfig:"KAFKA_BOOTSTRAP_SERVERS" default:"localhost:9092,localhost:9092"`
}
