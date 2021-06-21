package routes

import "github.com/gin-gonic/gin"

func UserRoute(route *gin.Engine)  {
	route.GET("/lists")
	route.POST("/lists/add")
	route.GET("/lists/:list_id")
	route.PUT("/lists/edit/:list_id")
	route.DELETE("/lists/delete/:list_id")
}