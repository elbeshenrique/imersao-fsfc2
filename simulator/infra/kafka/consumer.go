package kafka

import (
	"fmt"
	"log"
	"os"

	cKafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// KafkaConsumer holds all consumer logic and settings of Apache Kafka connections/
// Also has a Message channel which is a channel where the messages are going to be pushed
type KafkaConsumer struct {
	MessageChannel chan *cKafka.Message
}

// NewKafkaConsumer creates a new KafkaConsumer struct with its message channel as dependency
func NewKafkaConsumer(messageChannel chan *cKafka.Message) *KafkaConsumer {
	return &KafkaConsumer{
		MessageChannel: messageChannel,
	}
}

// Consume consumes all message pulled from apache kafka and sent it to message channel
func (kafkaConsumer *KafkaConsumer) Consume() {
	configMap := &cKafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id":          os.Getenv("KafkaConsumerGroupId"),
	}

	consumer, err := cKafka.NewConsumer(configMap)
	if err != nil {
		log.Fatalf("error consuming kafka message:" + err.Error())
	}

	topics := []string{os.Getenv("KafkaReadTopic")}
	consumer.SubscribeTopics(topics, nil)
	fmt.Println("Kafka consumer has been started")

	for {
		message, err := consumer.ReadMessage(-1)
		if err == nil {
			kafkaConsumer.MessageChannel <- message
		}
	}
}
