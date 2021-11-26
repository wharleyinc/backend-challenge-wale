package main

import (
	"backend-wale/adapter"
	"backend-wale/app"
	"backend-wale/config"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {

	_ = godotenv.Load("backend.env")
	port := os.Getenv("PORT")
	router := gin.Default()

	router.Use(gin.Logger())

	adapter.NewMongoDatabaseAdapter(context.Background(), config.Load())

	router.POST("/auth/login", app.Login)
	router.POST("/auth/signup", app.Signup)
	//router.POST("/auth/google/login", AuthLogin)
	//router.POST("/auth/google/signup", AuthSignup)
	//router.POST("/auth/forgot-password", forgotPassword)
	//router.POST("/auth/reset-password", resetPassword)

	router.Run(":" + port)
}
