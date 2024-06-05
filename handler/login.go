package handler

import (
	"fmt"
	"log"
	"net/http"
	"ngc13/logger"
	"ngc13/model"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	KEY = "ADMIN"
)

func generateToken(s model.User) (string, error) {
	payload := jwt.MapClaims{
		"username": s.Username,
		"password": s.Password,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := token.SignedString([]byte(KEY))
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("unable to get token")
	}
	return tokenString, nil
}

func (r *Repo) Login(c echo.Context) error {
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		logger.Logging(c).Warning("error binding, ", err)
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	//validate user
	if user.Username == "" || user.Password == "" {
		logger.Logging(c).Warning("empty param")
		return c.JSON(http.StatusBadRequest, "error or missing parameter")
	}

	// check existence of username
	var isExist model.User
	res := r.DB.Where("username = ?", user.Username).First(&isExist)
	if res.Error != nil {
		if res.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusBadRequest, "user already exist")
		}
		logger.Logging(c).Error(res.Error)
		return c.JSON(http.StatusInternalServerError, "Status internal error")
	}

	tokenString, _ := generateToken(isExist)

	return c.JSON(
		http.StatusOK,
		map[string]interface{}{
			"message": "login success",
			"token":   tokenString,
		},
	)
}
