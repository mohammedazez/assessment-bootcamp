package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/users"
)

type userHandler struct {
	userService users.Service
}

func NewUserHandler(userService users.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) GetAllUsersHandler(c *gin.Context) {
	users, err := h.userService.GetAllUser()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *userHandler) GetUserByIDHandler(c *gin.Context) {
	paramID := c.Params.ByName("user_id")

	users, err := h.userService.GetUserByIDService(paramID)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, users)
}
