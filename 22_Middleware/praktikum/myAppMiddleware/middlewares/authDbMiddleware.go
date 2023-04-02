package middlewares

import (
	"myAppMiddleware/config"
	"myAppMiddleware/models"

	"github.com/labstack/echo/v4"
)

func BasicAuthDB(username, password string, c echo.Context) (bool, error) {
	var db = config.DB
	var user models.User

	if err := db.Where("email = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return false, nil
	}
	return true, nil
}
