package kafka

import (
	"log"
	"os"

	cKafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// NewKafkaProducer creates a ready to go kafka.Producer instance
func NewKafkaProducer() *cKafka.Producer {
	configMap := &cKafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
	}

	producer, err := cKafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}

	return producer
}

// Publish is simple function created to publish new message to kafka
func Publish(message string, topic string, producer *cKafka.Producer) error {
	kafkaMessage := &cKafka.Message{
		TopicPartition: cKafka.TopicPartition{Topic: &topic, Partition: cKafka.PartitionAny},
		Value:          []byte(message),
	}

	err := producer.Produce(kafkaMessage, nil)
	if err != nil {
		return err
	}

	return nil
}
