/*
@Time : 2022/5/13 15:30
@Author : weixiaowei
@File : main
*/
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"src/com.wxw/project_actual/src/com.wxw/04_project_actual/06_actual_swagger/api"
	// This line is necessary for go-swagger to find your docs!
)

var users []*api.User

func main() {
	r := gin.Default()
	r.POST("/users", Create)
	r.GET("/users/:name", Get)

	log.Fatal(r.Run(":5555"))
}

// Create a user in memory.
func Create(c *gin.Context) {
	var user api.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "code": 10001})
		return
	}

	for _, u := range users {
		if u.Name == user.Name {
			c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("user %s already exist", user.Name), "code": 10001})
			return
		}
	}

	users = append(users, &user)
	c.JSON(http.StatusOK, user)
}

// Get return the detail information for a user.
func Get(c *gin.Context) {
	username := c.Param("name")
	for _, u := range users {
		if u.Name == username {
			c.JSON(http.StatusOK, u)
			return
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("user %s not exist", username), "code": 10002})
}
