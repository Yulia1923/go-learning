package handler

import (
	"net/http"
	"strconv"

	"go-learning/day26/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}

}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "用户ID格式不正确",
		})
		return
	}

	if id == 0 || uint64(uint(id)) != id {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "用户ID超出有效范围",
		})
		return
	}

	user, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "用户不存在或查询失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
