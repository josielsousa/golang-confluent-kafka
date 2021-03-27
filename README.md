Build and run golang apps that use the [confluentinc/confluent-kafka-go lib kafka](github.com/confluentinc/confluent-kafka-go)

## Build image

```bash
git clone https://github.com/josielsousa/golang-confluent-kafka.git

docker build -t go-kafka-gclib:v0.0.1 .
```

## Kafka with Kafdrop

- Download `docker-compose` file 

```bash 
docker-compose up -d
```

- Create a new topic [localhost kafdrop](http://localhost:19000/)

```bash

docker-compose exec kafka kafka-topics \
  --create \
  --bootstrap-server localhost:9092 \
  --replication-factor 1 \
  --partitions 1 \
  --if-not-exists \
  --topic alpine-topic-msgs

```

- Example to produce a new message

```bash
docker-compose exec kafka  \
  bash -c "seq 10 | kafka-console-producer \
    --request-required-acks 1 \
    --broker-list localhost:29092 \
    --topic alpine-topic-msgs && \
    echo 'Produced 10 messages.'
  "
```

## Run docker container

```bash
docker run --rm --network host \
--env KAFKA_BOOTSTRAP_SERVERS="hostname:9092,hostname:9092" \
--name gokafka-test go-kafka-gclib:v0.0.1

```
