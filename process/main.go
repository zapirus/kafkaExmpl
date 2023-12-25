package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	topic := "val"
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "foo",
		"auto.offset.reset": "smallest",
	})
	if err != nil {
		log.Fatalln(err)
	}
	err = consumer.Subscribe(topic, nil)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		ev := consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			fmt.Printf("чтение заказа: %v\n", string(e.Value))
		case kafka.Error:
			fmt.Printf("%v\n", e.Error())

		}
	}
}
