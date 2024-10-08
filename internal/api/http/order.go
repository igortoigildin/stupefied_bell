package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/igortoigildin/stupefied_bell/internal/config"
	"github.com/igortoigildin/stupefied_bell/internal/model"
	"github.com/igortoigildin/stupefied_bell/internal/storage"
	processjson "github.com/igortoigildin/stupefied_bell/pkg/lib/processJSON"
	"github.com/igortoigildin/stupefied_bell/pkg/logger"
	"go.uber.org/zap"
)

//go:generate  mockery --name=OrderRepository
type OrderRepository interface {
	SaveOrder(ctx context.Context, order model.Order) (string, error)
	SelectAllOrders(ctx context.Context) ([]model.Order, error)
	DeleteOrder(ctx context.Context, number string) error
	UpdateOrder(ctx context.Context, order model.Order) error
	UpdateStatus(ctx context.Context, orderID string, status string) error
}

func addOrderHandler(cfg *config.Config, repository OrderRepository) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
		defer cancel()

		var order model.Order
		err := processjson.ReadJSON(r, &order)
		if err != nil {
			logger.Log.Info("cannot decode request JSON body", zap.Error(err))
			processjson.SendJSONError(w, http.StatusBadRequest, "body contains badly-formed JSON")
			return
		}

		if order.Number == "" {
			logger.Log.Info("order number not provided", zap.Error(err))
			processjson.SendJSONError(w, http.StatusBadRequest, "order number not provided")
			return
		}

		number, err := repository.SaveOrder(ctx, order)
		if err != nil {
			logger.Log.Info("error while saving order", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if number == "" {
			logger.Log.Info("order already exists", zap.Error(err))
			w.WriteHeader(http.StatusOK)
			return
		}

		var response struct {
			Number string `json:"number"`
		}
		response.Number = number

		err = processjson.WriteJSON(w, http.StatusOK, response, nil)
		if err != nil {
			logger.Log.Info("error while saving order", zap.Error(err))
			w.WriteHeader(http.StatusAccepted)
			return
		}
	})
}

func SelectAllOrders(cfg *config.Config, repository OrderRepository) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
		defer cancel()

		orders, err := repository.SelectAllOrders(ctx)
		if err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				w.WriteHeader(http.StatusNoContent)
				return
			default:
				logger.Log.Info("error requesting orders", zap.Error(err))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		js, err := json.Marshal(orders)
		if err != nil {
			logger.Log.Info("error while marshalling", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)
	})
}

func deleteOrder(cfg *config.Config, repository OrderRepository) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
		defer cancel()

		number := r.URL.Query().Get("id")
		fmt.Println(number)
		err := repository.DeleteOrder(ctx, number)
		if err != nil {
			logger.Log.Info("error while deleting order", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func updateOrder(cfg *config.Config, repository OrderRepository) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
		defer cancel()

		var order model.Order
		err := processjson.ReadJSON(r, &order)
		if err != nil {
			logger.Log.Info("cannot decode request JSON body", zap.Error(err))
			processjson.SendJSONError(w, http.StatusBadRequest, "body contains badly-formed JSON")
			return
		}

		err = repository.UpdateOrder(ctx, order)
		if err != nil {
			switch {
			case errors.Is(err, storage.ErrOrderNotFound):
				logger.Log.Info("such order not found", zap.Error(err))
				processjson.SendJSONError(w, http.StatusBadRequest, "such order not found")
				return
			default:
				logger.Log.Info("error while deleting order", zap.Error(err))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	})
}
