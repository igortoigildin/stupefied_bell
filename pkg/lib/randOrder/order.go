package order

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/igortoigildin/stupefied_bell/internal/model"
	"github.com/igortoigildin/stupefied_bell/pkg/logger"
	"go.uber.org/zap"
)

func RandomOrder() ([]byte, error) {
	order := model.Order{
		Number:     uuid.New().String(),
		Quantity:   5,
		Title:      uuid.New().String(),
		UploadedAt: time.Now(),
		Status:     "New",
	}

	value, err := json.Marshal(order)
	if err != nil {
		logger.Log.Error("failed to encode order for kafka", zap.Error(err))
		return nil, err
	}

	return value, nil
}
