FROM golang:alpine AS builder

RUN apk add --no-cache ca-certificates gcc mingw-w64-gcc build-base curl

WORKDIR /go/src/app

COPY . .

RUN go test -tags musl ./...

RUN go get honnef.co/go/tools/cmd/staticcheck

# RUN staticcheck -tags musl ./...

# RUN go vet -tags musl ./...

RUN go mod verify

RUN go get github.com/confluentinc/confluent-kafka-go@v1.5.2

RUN CGO_ENABLED=0 GOOS=linux go build -o helloworld -a -v 

ENTRYPOINT [ "/go/src/app/helloworld" ]