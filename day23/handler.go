package main

import "github.com/gin-gonic/gin"

func getUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "get users",
	})
}

func getUser(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, gin.H{
		"message": "get user",
		"id":      id,
	})
}

func createUser(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "create user",
	})
}

func updateUser(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, gin.H{
		"message": "update user",
		"id":      id,
	})
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, gin.H{
		"message": "delete user",
		"id":      id,
	})
}
