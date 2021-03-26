Build and run golang apps that use the [confluentinc/confluent-kafka-go lib kafka](github.com/confluentinc/confluent-kafka-go)

## Build image

```bash
git clone https://github.com/josielsousa/golang-gclib-alpine-confluent-kafka

docker build -t go-kafka-gclib:v0.0.1 .
```

## Kafka with Kafdrop

- Download `docker-compose` file 

```bash 
docker-compose up -d
```

- Create a topic on [localhost kafdrop](http://localhost:19000/)

```bash

docker-compose exec kafka kafka-topics \
  --create \
  --bootstrap-server localhost:9092 \
  --replication-factor 1 \
  --partitions 1 \
  --if-not-exists \
  --topic alpine-topic-msgs

```

## Run docker container

```bash
docker run --rm --network="host" --name gokafka-test go-kafka-gclib:v0.0.1
```