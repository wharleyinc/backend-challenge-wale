package main

import (
	"backend-wale/adapter"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
)

type apiUser struct {
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type apiCode struct {
	Code string `json:"code" binding:"required"`
}

type apiHeader struct {
	Authorization string `json:"Authorization" binding:"required"`
}

type apiToken struct {
	ExpiresAt   time.Time `json:"expiresAt" binding:"required"`
	AccessToken string    `json:"accessToken" binding:"required"`
}

func Login (c *gin.Context) {
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

func Signup (c *gin.Context) {
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


func AuthSignup (c *gin.Context) {
	var user apiUser
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, "ErrorInvalidRequest")
		return
	}

	result, err := adapter.CreateAccount(c.Request.Context(), user.Email, user.Password)
	if err != nil {
		log.Println(c, "unable to retrieve account", zap.Error(err))
		c.JSON(500, err)
		//return "Unable to retrieve account", false
		return
	}

	//return "Signup successful for user: " + user.Email, result
	c.JSON(http.StatusOK, result)
}

/*func AuthLogin (c *gin.Context) {
	var user apiUser
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, "ErrorInvalidRequest")
		return
	}

	result, err := adapter.GetAccount(c.Request.Context(), user.Email, user.Password)
	if err != nil {
		log.Println(c, "unable to retrieve account", zap.Error(err))
		c.JSON(500, zap.Error(err))
		//return "Unable to retrieve account", false
		return
	}

	c.JSON(http.StatusOK, result)

	return

}*/