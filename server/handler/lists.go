package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lists)
}
