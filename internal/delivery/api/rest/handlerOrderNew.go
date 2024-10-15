package rest

import (
	"net/http"

	pkg "github.com/igortoigildin/stupefied_bell/pkg/lib/randOrder"
	"github.com/igortoigildin/stupefied_bell/pkg/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// responds on GET requests via '/order' endpoint and provides new order for couriers.
func handlerOrderNew(c echo.Context) error {
	order, err := pkg.RandomOrder()
	if err != nil {
		logger.Log.Error("error:", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, order)
}
