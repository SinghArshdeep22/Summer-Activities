package handlers

import (
	"net/http"

	"coworkingapp/models"
	"coworkingapp/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserInfo struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignupReq struct {
	UserInfo
	Email string `json:"email" binding:"required"`
}

func Login(c *gin.Context) {
	var userInfo UserInfo
	if err := c.ShouldBind(&userInfo); err != nil {
		c.JSON(http.StatusBadRequest, models.CoworkingError{Code: models.ValidationErr, Message: err.Error()})
		return
	}
	db := c.MustGet("DbKey").(*gorm.DB)
	signedInUser, err := models.LoginUser(db, userInfo.Username, userInfo.Password)
	if err != nil {
		coworkingErr := err.(models.CoworkingError)
		c.JSON(coworkingErr.StatusCode, coworkingErr)
		return
	}
	secretKey := c.MustGet("ConfigKey").(models.CoworkingConfig).SecretKey
	token, err := utils.GenerateToken(signedInUser.Email, []byte(secretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.CoworkingError{Code: models.TokenGenerationErr, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Signup(c *gin.Context) {
	var signupReq SignupReq
	if err := c.ShouldBind(&signupReq); err != nil {
		c.JSON(http.StatusBadRequest, models.CoworkingError{Code: models.ValidationErr, Message: err.Error()})
		return
	}
	db := c.MustGet("DbKey").(*gorm.DB)
	user := models.User{Email: signupReq.Email, Username: signupReq.Username, Password: signupReq.Password}
	id, err := models.SignUpUser(db, user)
	if err != nil {
		coworkingErr := err.(models.CoworkingError)
		c.JSON(coworkingErr.StatusCode, coworkingErr)
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}
