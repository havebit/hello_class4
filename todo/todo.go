package todo

import (
	"github.com/gin-gonic/gin"
)

type store interface {
	NewTask(task string) error
}

type Handler struct {
	store store
}

func NewHandler(store store) *Handler {
	return &Handler{store: store}
}

type NewTaskRequest struct {
	Title string `json:"title"`
}

func (h *Handler) NewTask(c *gin.Context) {
	var req NewTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.store.NewTask(req.Title); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}
