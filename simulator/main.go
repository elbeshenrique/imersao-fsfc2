package main

import (
	"fmt"
	"log"

	kafkaProduce "github.com/codeedu/imersaofsfc2-simulator/application/kafka"
	"github.com/codeedu/imersaofsfc2-simulator/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
)

func init() {
	log.Default().Print("Started")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	messageChannel := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(messageChannel)

	go consumer.Consume()
	for message := range messageChannel {
		fmt.Println(string(message.Value))
		go kafkaProduce.Produce(message)
	}
}
