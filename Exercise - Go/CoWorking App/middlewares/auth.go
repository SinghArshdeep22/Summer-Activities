package middlewares

import (
	"net/http"

	"coworkingapp/models"
	"coworkingapp/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthorizeUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("Authorization")
		if tokenHeader == "" {
			c.JSON(http.StatusUnauthorized, models.CoworkingError{Code: models.MissingTokenErr, Message: "please provide a jwt token along with the HTTP request"})
			return
		}
		secretKey := c.MustGet("ConfigKey").(models.CoworkingConfig).SecretKey
		claims, err := utils.ValidateToken(tokenHeader, []byte(secretKey))
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.CoworkingError{Code: models.TokenNotValidErr, Message: err.Error()})
			return
		}
		email := (*claims)["sub"].(string)
		db := c.MustGet("DbKey").(*gorm.DB)
		user, err := models.GetUserByEmail(db, email)
		if err != nil {
			coworkingErr := err.(models.CoworkingError)
			c.JSON(coworkingErr.StatusCode, coworkingErr)
			return
		}
		c.Set("UserIdKey", user.ID)
		c.Next()
	}
}
