package main

import (
	"first-api-go/api/handlers"
	"first-api-go/pkg/repository"
	"first-api-go/pkg/services"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := repository.NewDB()
	userService := services.NewUserService(db)
	userHandler := handlers.NewUserHandler(userService)

	r.GET("/users", userHandler.GetUsers)
	r.POST("/users", userHandler.CreateUser)
	r.PUT("/users/:id", userHandler.UpdateUser)

	err := r.Run()
	if err != nil {
		return
	}
}
