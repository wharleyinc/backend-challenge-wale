package app

import (
	"backend-wale/adapter"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
)

type apiUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var user apiUser
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, "ErrorInvalidRequest")
		return
	}

	result, err := adapter.GetAccountNormal(c.Request.Context(), user.Email, user.Password)
	if err != nil {
		log.Println(c, "unable to retrieve account", zap.Error(err))
		c.JSON(500, zap.Error(err))
		//return "Unable to retrieve account", false
		return
	}

	c.JSON(http.StatusOK, result)

	return

}

func Signup(c *gin.Context) {
	var user apiUser
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, "ErrorInvalidRequest")
		return
	}

	result, err := adapter.CreateAccountNormal(c.Request.Context(), user.Email, user.Password)
	if err != nil {
		log.Println(c, "unable to retrieve account", zap.Error(err))
		c.JSON(500, err)
		//return "Unable to retrieve account", false
		return
	}

	//return "Signup successful for user: " + user.Email, result
	c.JSON(http.StatusOK, result)
}
