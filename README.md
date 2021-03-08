Build and run golang apps that use the [confluentinc/confluent-kafka-go lib kafka](github.com/confluentinc/confluent-kafka-go)

## Build image

```bash
git clone https://github.com/josielsousa/golang-gclib-alpine-confluent-kafka

docker build -t go-kafka-gclib:v0.0.1 .
```

## Confluent Plataform

- Used quick start avaiable [here](https://docs.confluent.io/platform/current/quickstart/ce-docker-quickstart.html)

- [Download](https://docs.confluent.io/platform/current/quickstart/ce-docker-quickstart.html#step-1-download-and-start-cp-using-docker) and Start Confluent Platform Using Docker

- Download `docker-compose` file 
```bash 
curl --silent --output docker-compose.yml \
  https://raw.githubusercontent.com/confluentinc/cp-all-in-one/6.1.0-post/cp-all-in-one/docker-compose.yml

docker-compose up -d
```

- Create an topic name [teste-golang-gclib-alpine](https://docs.confluent.io/platform/current/quickstart/ce-docker-quickstart.html)

## Run docker container

```bash
docker run --rm --network="host" --name gokafka-test go-kafka-gclib:v0.0.1
```