package rest

import (
	"net/http"
	"strconv"

	config "github.com/igortoigildin/stupefied_bell/config/delivery"
	service "github.com/igortoigildin/stupefied_bell/internal/delivery/services"

	"github.com/igortoigildin/stupefied_bell/pkg/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func HandlerOrderDone(cfg *config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {

		id := c.QueryParam("id")
		_, err := strconv.Atoi(id)
		if err != nil {
			logger.Log.Error("error while conviring order to int", zap.Error(err))
			c.JSON(http.StatusBadRequest, HttpErrorResponse{Error: err, Explanation: "Invalid order ID"})
		}
		status := c.QueryParam("status")

		service.SendGRPCResponse(cfg, status, id)

		c.JSON(http.StatusOK, nil)
		return nil
	}
}
