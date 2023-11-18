package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nmorales1991/first-api-go/api/handlers"
	"github.com/nmorales1991/first-api-go/pkg/repository"
	"github.com/nmorales1991/first-api-go/pkg/services"
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
