package kafka

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/igortoigildin/stupefied_bell/internal/order/model"
	"github.com/igortoigildin/stupefied_bell/pkg/logger"
	"go.uber.org/zap"
)

func RunReading(cons *kafka.Consumer) {
	for {
		msg, err := cons.ReadMessage(-1)
		if err == nil {
			var order model.Order
			err := json.Unmarshal(msg.Value, &order)
			if err != nil {
				logger.Log.Error("error decoding message: %v\n", zap.Error(err))
				continue
			}

			logger.Log.Info("Received Order", zap.Any("order:", order))
		} else {
			logger.Log.Error("error: %v\n", zap.Error(err))
		}
	}
}
