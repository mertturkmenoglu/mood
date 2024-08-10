package api

import (
	"mood/internal/h"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Api struct {
	Port int
}

func New() *Api {
	a := &Api{
		Port: 3000,
	}

	return a
}

func (api *Api) RegisterRoutes() *echo.Echo {
	e := echo.New()
	v1 := e.Group("/api/v1")

	v1.GET("/foo", func(c echo.Context) error {
		return c.JSON(http.StatusOK, h.Response[string]{
			"message": "Hello World",
		})
	})

	return e
}
