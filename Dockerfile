FROM golang:alpine AS builder

RUN apk add --no-cache ca-certificates build-base 

WORKDIR /go/src/app

COPY . .

RUN go vet -tags musl ./...

RUN go mod verify

RUN GOARCH=amd64 GOOS=linux go build -a -v --ldflags '-extldflags "-static" -s -w' -tags musl -o bin/go-kafka-gclib 

ENTRYPOINT [ "/go/src/app/bin/go-kafka-gclib" ]