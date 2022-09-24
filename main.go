package main

import (
	"net/http"
	"profile/api"
	"profile/api/english"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "https://portfolio.yukiserv.com"},
		AllowHeaders: []string{http.MethodGet, echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/profile/api/v1/works", api.GetWork)
	e.GET("/profile/api/v1/tackles", api.GetTackle)
	e.GET("/profile/api/v1/news", api.GetNews)
	e.GET("/profile/api/v1/en/works", english.GetWork)
	e.GET("/profile/api/v1/en/tackle", english.GetTackle)
	e.GET("/profile/api/v1/en/news", english.GetNews)
	e.Logger.Fatal(e.Start(":1323"))
}
