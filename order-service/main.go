package main

import (
	"user-service/order-service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/orders",handlers.CreateOrder);
	r.GET("/order/:id",handlers.GetOrder);

	r.Run(":8002")
}
