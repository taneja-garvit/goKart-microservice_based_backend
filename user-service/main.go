package main

import (
	"user-service/user-service/db"
	"user-service/user-service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.InitMongo()

	r.POST("/users",handlers.CreateUser)
	r.GET("/users/:id",handlers.GetUser)

	r.Run(":8001")
}