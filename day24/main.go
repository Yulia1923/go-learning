package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Age      int    `json:"age"`
}

func main() {
	r := gin.Default()

	users := map[int]User{
		1: {ID: 1, Username: "Alice", Age: 20},
		2: {ID: 2, Username: "Bob", Age: 22},
	}

	nextID := 3

	//查询所有用户，支持query参数

	r.GET("/users", func(c *gin.Context) {
		keyword := c.Query("keyword")

		result := make([]User, 0)

		for _, user := range users {
			if keyword == "" || user.Username == keyword {
				result = append(result, user)
			}
		}

		c.JSON(http.StatusOK, result)
	})

	r.GET("/users/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid user id",
			})
			return
		}
		user, ok := users[id]
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "user not found",
			})
			return
		}
		c.JSON(http.StatusOK, user)
	})
	r.POST("/users", func(c *gin.Context) {
		var user User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid JSON",
			})
			return
		}

		if user.Username == "" || user.Age <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "username and positive age are required",
			})
			return
		}

		user.ID = nextID
		nextID++

		users[user.ID] = user

		c.JSON(http.StatusCreated, user)
	})

	r.PUT("/users/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid user id",
			})
			return
		}

		var updateUser User
		if err := c.ShouldBindJSON(&updateUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid JSON",
			})

			return
		}
		if updateUser.Username == "" || updateUser.Age <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "username and positive age are required",
			})
			return
		}
		updateUser.ID = id
		users[id] = updateUser

		c.JSON(http.StatusOK, updateUser)
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid user id",
			})
			return
		}
		if _, ok := users[id]; !ok {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "user not found",
			})
			return
		}
		delete(users, id)
		c.JSON(http.StatusOK, gin.H{
			"message": "user deleted",
		})
	})

	r.Run(":8080")
}
