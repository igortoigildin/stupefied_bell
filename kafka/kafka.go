package kafka

import (
	"context"
	"time"

	"github.com/igortoigildin/stupefied_bell/internal/config"
	"github.com/igortoigildin/stupefied_bell/pkg/logger"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type Kafka struct {
	Config config.Config
	Writer *kafka.Writer
	Reader *kafka.Reader
}

func newWriter(kafkaConfig config.Config) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaConfig.Brokers...),
		Async:    true,
		Topic:    kafkaConfig.Topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func newReader(kafkaConfig config.Config) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  kafkaConfig.Brokers,
		Topic:    kafkaConfig.Topic,
		MinBytes: 10e3,             // 10KB
		MaxBytes: 10e6,             // 10 MB
		MaxWait:  10 * time.Second, // Maximun time to wait for new data
	})
}

func NewKafka(config config.Config) *Kafka {
	return &Kafka{
		Config: config,
		Writer: newWriter(config),
		Reader: newReader(config),
	}
}

func (k *Kafka) Produce(ctx context.Context, key, value []byte) error {
	return k.Writer.WriteMessages(ctx, kafka.Message{
		Key:   key,
		Value: value,
	})
}

func (k *Kafka) Consume(ctx context.Context) {
	for {
		msg, err := k.Reader.ReadMessage(ctx)
		if err != nil {
			logger.Log.Error("Kafka failed to read a message: %v", zap.Error(err))
			return
		}
		logger.Log.Sugar().Infof("Received message: key=%s value=%s", string(msg.Key), string(msg.Value))
	}
}
