package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Res struct {
	Return  string `json:"string"`
	Version string `json:"version"`
}

func main() {
	// Echo instance
	e := echo.New()

	// load prom
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {

		r := &Res{
			Return:  "Hello World",
			Version: os.Getenv("APP_VERSION"),
		}
		return c.JSON(http.StatusOK, r)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
