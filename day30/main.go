package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/profile", AuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"userID": c.GetInt("userID"),
			"role":   c.GetString("role"),
		})
	})

	router.GET("/admin", AuthMiddleware(), AdminMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "欢迎进入管理员接口",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		user := User{
			ID:       1,
			Username: "parksiwoo",
			Role:     "user",
		}
		tokenString, err := GenerateToken(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "生成 JWT 失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"user_id": user.ID,
			"role":    user.Role,
			"token":   tokenString,
		})
	})

	router.GET("/test-token", func(c *gin.Context) {
		user := User{
			ID:       1,
			Username: "parksiwoo",
			Role:     "user",
		}

		tokenString, err := GenerateToken(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "生成 JWT 失败",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user_id": user.ID,
			"role":    user.Role,
			"token":   tokenString,
		})
	})

	router.GET("/test-token-admin", func(c *gin.Context) {
		admin := User{
			ID:       2,
			Username: "admin",
			Role:     "admin",
		}

		token, err := GenerateToken(admin)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "生成Token失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"user_id": admin.ID,
			"role":    admin.Role,
			"token":   token,
		})
	})
	router.Run(":8080")
}
