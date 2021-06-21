package handler

import (
	"github.com/gin-gonic/gin"
	"server/lists"
)

type listHandler struct {
	listService lists.Service
}

func NewListHandler(listService lists.Service) *listHandler {
	return &listHandler{listService}
}

func (h *listHandler) GetAllListsHandler(c *gin.Context) {
	lists, err := h.listService.GetAllLists()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, lists)
}

func (h *listHandler) GetListByIDHandler(c *gin.Context) {
	paramID := c.Params.ByName("list_id")

	lists, err := h.listService.GetListByIDService(paramID)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, lists)
}
