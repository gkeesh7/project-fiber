package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	brokerList := os.Getenv("KAFKA_BROKER_LIST")

	// Set up the Kafka producer
	producer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  strings.Split(brokerList, ","),
		Topic:    "test-topic",
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
	err = producer.WriteMessages(
		context.Background(),
		kafka.Message{
			Key:   []byte("key"),
			Value: byteValue,
		},
	)
	if err != nil {
		fmt.Println(err.Error() + " This error happened")
		return
	}
	fmt.Println("All good")
}
