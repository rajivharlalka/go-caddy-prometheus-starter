package main

import (
	"net/http"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(echoprometheus.NewMiddleware("myapp")) // adds middleware to gather metrics
	e.Use(middleware.Logger())
	e.GET("/metrics", echoprometheus.NewHandler()) // adds route to serve gathered metrics
	e.GET("/ping", ping)
	e.Start(":8000")
}

func ping(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"response": "pong",
	})
}
