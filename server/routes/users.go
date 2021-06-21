package routes

import "github.com/gin-gonic/gin"

func UserRoute(route *gin.Engine)  {
	route.GET("/users")
	route.POST("/users/register")
	route.POST("/users/login")
	route.GET("/users/:user_id")
	route.PUT("/users/edit/:user_id")
	route.DELETE("/users/delete/:user_id")
}