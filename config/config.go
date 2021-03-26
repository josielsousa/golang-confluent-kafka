package config

type KafkaClient struct {
	ClientID  string `envconfig:"KAFKA_CLIENT_ID" default:"app.test.golang.gclib"`
	TopicName string `envconfig:"KAFKA_TOPIC" default:"app.teste.golang.gclib.alpine"`
	Servers   string `envconfig:"KAFKA_BOOTSTRAP_SERVERS" default:"localhost:9092,localhost:9092"`
}
