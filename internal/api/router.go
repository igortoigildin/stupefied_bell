package api

import (
	"net/http"

	"github.com/igortoigildin/stupefied_bell/internal/config"
	logging "github.com/igortoigildin/stupefied_bell/internal/lib/logger"
)

func Router(cfg *config.Config, rep OrderRepository) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/order", logging.WithLogging(addOrderHandler(cfg, rep)))
	mux.HandleFunc("GET /api/orders", logging.WithLogging(SelectAllOrders(cfg, rep)))
	mux.HandleFunc("DELETE /api/order/", logging.WithLogging(deleteOrder(cfg, rep)))
	mux.HandleFunc("PUT /api/order", logging.WithLogging(updateOrder(cfg, rep)))

	return mux
}
