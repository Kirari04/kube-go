package main

import (
	"fmt"
	"net/http"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

var name = fmt.Sprintf("%s-%s", gofakeit.Verb(), gofakeit.Noun())

func main() {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("I'm the %s", name))
	})
	e.Logger.SetLevel(log.INFO)
	e.Logger.Info("Running on http://127.0.0.1:1323")
	e.Logger.Fatal(e.Start(":1323"))
}
