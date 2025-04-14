package handlers

import (
	"fmt"
	"net/http"
	"user-service/order-service/models"

	"github.com/gin-gonic/gin"
)

var orders = make(map[string]models.Order)

func CreateOrder(c *gin.Context){
	var order  models.Order

	if err:= c.ShouldBindJSON(&order); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	userURL := fmt.Sprintf("http://user-service:8001/users/%s", order.UserID)
	resp, err := http.Get(userURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
		}
		

	orders[order.ID]=order

	c.JSON(http.StatusOK,order);


}

func GetOrder(c *gin.Context){
	id:= c.Param("id")
	order, exist:= orders[id];
	if !exist{
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
		}

		c.JSON(http.StatusOK,order)
}