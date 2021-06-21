package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"server/auth"
	"server/config"
	"server/handler"
	"server/users"
)

var (
	DB             *gorm.DB = config.ConnectionToDatabase()
	userRepository          = users.NewRepository(DB)
	userService             = users.NewService(userRepository, authService)
	authService             = auth.NewService()
	userHandler             = handler.NewUserHandler(userService)
	middleware              = handler.Middleware(userService, authService)
)

func UserRoute(route *gin.Engine) {
	route.GET("/users", userHandler.GetAllUsersHandler)
	route.POST("/users/register", userHandler.UserRegisterHandler)
	route.POST("/users/login", userHandler.UserLoginHandler)
	route.GET("/users/:user_id", userHandler.GetUserByIDHandler)
	route.PUT("/users/edit/:user_id")
	route.DELETE("/users/delete/:user_id")
}
