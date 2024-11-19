package rest

import (
	"net/http"
	"strconv"

	config "github.com/igortoigildin/stupefied_bell/config/delivery"
	service "github.com/igortoigildin/stupefied_bell/internal/delivery/services"
	e "github.com/igortoigildin/stupefied_bell/pkg/errors"
	"github.com/igortoigildin/stupefied_bell/pkg/logger"
	rand "github.com/igortoigildin/stupefied_bell/pkg/randOrder"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type OrderController struct {
	cfg *config.Config
	c   echo.Context
}

func NewOrderController(cfg *config.Config) *OrderController {
	return &OrderController{
		cfg: cfg,
	}
}

// responds on GET requests via '/order' endpoint and provides new order for couriers.
func (o *OrderController) createOrder() echo.HandlerFunc {
	return func(c echo.Context) error {

		order, err := rand.RandomOrder()
		if err != nil {
			logger.Log.Error("error:", zap.Error(err))
			return o.c.JSON(http.StatusInternalServerError, nil)
		}
		return o.c.JSON(http.StatusOK, order)
	}
}

func (o *OrderController) updateOrder() echo.HandlerFunc {
	return func(c echo.Context) error {

		id := o.c.QueryParam("id")
		_, err := strconv.Atoi(id)
		if err != nil {
			logger.Log.Error("error while conviring order to int", zap.Error(err))
			_ = o.c.JSON(http.StatusBadRequest, e.HttpErrorResponse{Error: err, Explanation: "Invalid order ID"})
		}

		status := o.c.QueryParam("status")

		_ = service.SendGRPCRequest(o.cfg, status, id)

		return o.c.JSON(http.StatusOK, nil)
	}
}
