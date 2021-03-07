package main

import (
	"fmt"
	"net/http"

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

func main() {
	// create a new handler
	handler := HTTPHandler{}

	fmt.Println("Starting API...")

	// intialize the writer with the broker addresses, and the topic
	kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "",
		"group.id":          "",
	})

	// listen and serve
	http.ListenAndServe(":9000", handler)
}
