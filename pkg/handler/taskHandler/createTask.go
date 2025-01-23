package taskhandler

import (
	"net/http"

	"rest.com/pkg/tasks"
	"github.com/gin-gonic/gin"
)


func CreateTask(c *gin.Context) {
	taskID, err := tasks.CreateTask()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	go tasks.RunTask(taskID)
	c.JSON(201, gin.H{"task_id": taskID})
}