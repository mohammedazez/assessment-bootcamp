package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"server/config"
	"server/handler"
	"server/users"
)

var (
	DB *gorm.DB = config.ConnectionToDatabase()
	userRepository = users.NewRepository(DB)
	userService = users.NewService(userRepository)
	userHandler = handler.NewUserHandler(userService)
)

func UserRoute(route *gin.Engine)  {
	route.GET("/users", userHandler.GetAllUsersHandler)
	route.POST("/users/register")
	route.POST("/users/login")
	route.GET("/users/:user_id", userHandler.GetUserByIDHandler)
	route.PUT("/users/edit/:user_id")
	route.DELETE("/users/delete/:user_id")
}