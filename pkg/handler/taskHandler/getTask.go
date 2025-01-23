package taskhandler

import (
	"rest.com/pkg/tasks"
	"github.com/gin-gonic/gin"
)


func GetTask(c *gin.Context) {
	taskID := c.Param("id")
	task := tasks.GetTask(taskID)
	if task == nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(200, task)
}