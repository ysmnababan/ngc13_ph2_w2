package main

import (
	"ngc13/config"
	"ngc13/handler"
	"ngc13/helper"
	"ngc13/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db := config.Connect()

	h := &handler.Repo{DB: db}

	e := echo.New()
	e.POST("/register", h.Register)
	e.POST("/login", h.Login)
	e.Use(middleware.Recover())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			logger.Logging(c).Info("incoming request")
			return next(c)
		}
	})

	buy := e.Group("")
	buy.Use(helper.Auth)
	{
		buy.GET("/filter-product", h.GetProductsFilter)
		buy.GET("/products", h.GetProducts)
		buy.GET("/stores", h.GetAllStores)
		buy.GET("/store/:id", h.GetDetailStore)
		buy.POST("/transactions", h.BuyProduct)
	}
	e.Logger.Fatal(e.Start(":8080"))
}
