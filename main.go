package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"io/ioutil"
	"os"
)

func main() {
	// Set up the Kafka producer
	producer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "my-topic",
		Balancer: &kafka.LeastBytes{},
	})
	// Defer Close the producer
	defer producer.Close()

	//Open the payload file
	payload_file, err := os.Open("resources/sample_payload.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer payload_file.Close()
	byteValue, err := ioutil.ReadAll(payload_file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send a message to Kafka
	producer.WriteMessages(
		context.Background(),
		kafka.Message{
			Key:   []byte("key"),
			Value: byteValue,
		},
	)
}
