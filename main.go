package main

import (
	"backend-wale/adapter"
	"backend-wale/config"
	"context"
	"github.com/gin-gonic/gin"
)

func main () {
	router := gin.Default()

	adapter.NewMongoDatabaseAdapter(context.Background(), config.Load())

	router.POST("/auth/login", Login)
	router.POST("/auth/signup", Signup)
	//router.POST("/auth/google/login", AuthLogin)
	router.POST("/auth/google/signup", AuthSignup)
	//router.POST("/auth/forgot-password", forgotPassword)
	//router.POST("/auth/reset-password", resetPassword)


	router.Run("localhost:8080")
}
