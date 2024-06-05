package handler

import (
	"net/http"
	"ngc13/logger"
	"ngc13/model"

	"github.com/labstack/echo/v4"
)

func (r *Repo) GetAllStores(c echo.Context) error {
	var stores []model.Stores
	res := r.DB.Find(&stores)
	if res.Error != nil {
		logger.Logging(c).Error("error query:", res.Error)
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, stores)
}
