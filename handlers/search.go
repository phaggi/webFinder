package handlers

import (
	"net/http"
	"strconv"
	"webFinder/services"

	"github.com/gin-gonic/gin"
)

type SearchHandler struct {
	searchService *services.SearchService
}

func NewSearchHandler(searchService *services.SearchService) *SearchHandler {
	return &SearchHandler{searchService: searchService}
}

func (h *SearchHandler) TriggerScript(c *gin.Context) {
	var req struct {
		ScriptName string `json:"script_name"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	taskID, err := h.searchService.TriggerScript(req.ScriptName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to trigger script"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task_id": taskID})
}

func (h *SearchHandler) GetResults(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("taskID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	results, err := h.searchService.GetResults(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get results"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": results})
}
