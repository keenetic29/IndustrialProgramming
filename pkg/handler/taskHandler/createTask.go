package taskhandler

import (
	"net/http"

	"rest.com/pkg/tasks"
	"github.com/gin-gonic/gin"
)

// CreateTask godoc
// @Summary      Create a new task
// @Description  Creates a new task and starts its execution asynchronously
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Security JWT
// @Success      201  {object}  map[string]string  "Task created successfully"
// @Failure      400  {object}  map[string]string  "Failed to create task"
// @Failure      401  {object}  map[string]string  "Invalid request or error message"
// @Router       /tasks [post]
func CreateTask(c *gin.Context) {
	taskID, err := tasks.CreateTask()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	go tasks.RunTask(taskID)
	c.JSON(201, gin.H{"task_id": taskID})
}