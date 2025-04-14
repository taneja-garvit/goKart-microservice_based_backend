package main

import (
	"user-service/order-service/handlers"
	"user-service/user-service/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.InitMongo()
	r.POST("/orders",handlers.CreateOrder);
	r.GET("/order/:id",handlers.GetOrder);

	r.Run(":8002")
}
