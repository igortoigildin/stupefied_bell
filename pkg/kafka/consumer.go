package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	config "github.com/igortoigildin/stupefied_bell/config/delivery"
)

func NewConsumer(cfg *config.Config) (*kafka.Consumer, error) {
	return kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.Kafka.Brokers[0],
		"group.id":          cfg.Kafka.GroupID,
		"auto.offset.reset": "earliest",
	})
}
