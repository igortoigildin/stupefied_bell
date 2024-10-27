package rest

import (
	config "github.com/igortoigildin/stupefied_bell/config/delivery"
	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, cfg *config.Config) {
	oc := NewOrderController(cfg)

	e.GET("/order", oc.createOrder())
	e.POST("/order", oc.updateOrder())
}
