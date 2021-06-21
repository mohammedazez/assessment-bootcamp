package handler

import (
	"github.com/gin-gonic/gin"
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
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, users)
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

func (h *userHandler) UserRegisterHandler(c *gin.Context) {
	var userInput users.UserRegisterInput

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(400, gin.H{"errors": err.Error()})
		return
	}

	users, err := h.userService.MakeNewUserService(userInput)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, users)
}

func (h *userHandler) UserLoginHandler(c *gin.Context) {
	var userInput users.UserLoginInput

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(400, gin.H{"errors": err.Error()})
		return
	}

	users, err := h.userService.UserLoginService(userInput)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, users)
}
