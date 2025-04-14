package handlers

import (
	"net/http"
	"user-service/user-service/models"

	"github.com/gin-gonic/gin"
)

var users = make(map[string]models.User)

func CreateUser(c *gin.Context){
	var user models.User
	if err:= c.ShouldBindJSON(&user); err!=nil{  //binding all JSON data coming from the frontend or postman
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users[user.ID] = user 
	c.JSON(http.StatusOK, user)
}

func GetUser(c *gin.Context){
	id:= c.Param("id")
	user, exist := users[id]
	if !exist {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
		}
		c.JSON(http.StatusOK,user)
}