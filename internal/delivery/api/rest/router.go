package rest

import (
	config "github.com/igortoigildin/stupefied_bell/config/delivery"
	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, cfg *config.Config) {
	e.GET("/order", handlerOrderNew)
	e.POST("/order", HandlerOrderDone(cfg))
}
