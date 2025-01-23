package taskhandler

import (
	"net/http"

	"rest.com/pkg/tasks"
	"github.com/gin-gonic/gin"
)


func CancelTask(c *gin.Context) {
	taskId := c.Param("id")
	task := tasks.GetTask(taskId)
	if task == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "task does not exists"})
		return
	}

	if err := tasks.StopTask(task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "task stopped"})
}