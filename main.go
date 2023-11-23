package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nmorales1991/first-api-go/api/handlers"
	"github.com/nmorales1991/first-api-go/pkg/repository"
	"github.com/nmorales1991/first-api-go/pkg/services"
	"os"
)

func main() {
	r := gin.Default()
	db := repository.NewDB()
	userService := services.NewUserService(db)
	userHandler := handlers.NewUserHandler(userService)

	r.GET("/users", userHandler.GetUsers)
	r.POST("/users", userHandler.CreateUser)
	r.PUT("/users/:id", userHandler.UpdateUser)

	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			panic("Error loading .env file")
		}
	}
	uri := os.Getenv("PORT")

	err := r.Run(":" + uri)
	if err != nil {
		return
	}
}
