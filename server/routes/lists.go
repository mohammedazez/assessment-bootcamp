package routes

import (
	"github.com/gin-gonic/gin"
	"server/handler"
	"server/lists"
)

var (
	listRepository = lists.NewRepository(DB)
	listService = lists.NewService(listRepository)
	listHandler = handler.NewListHandler(listService)
)


func ListRoute(route *gin.Engine)  {
	route.GET("/lists", listHandler.GetAllListsHandler)
	route.POST("/lists/add", listHandler.CreateNewList)
	route.GET("/lists/:list_id", listHandler.GetListByIDHandler)
	route.PUT("/lists/edit/:list_id")
	route.DELETE("/lists/delete/:list_id")
}