package handler

import (
	"net/http"
	"ngc13/logger"
	"ngc13/model"

	"github.com/labstack/echo/v4"
)

func (r *Repo) GetProducts(c echo.Context) error {

	var products []model.Product
	res := r.DB.Find(&products)
	if res.Error != nil {
		logger.Logging(c).Error(res.Error)
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, products)
}
