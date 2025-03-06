package taskhandler

import (
	"net/http"

	"rest.com/pkg/tasks"
	"github.com/gin-gonic/gin"
)

// CancelTask godoc
// @Summary      Cancel a task
// @Description  Stops a running task by its ID
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Security JWT
// @Param        id   path      string  true  "Task ID"
// @Success      200  {object}  map[string]string  "Task stopped successfully"
// @Failure      404  {object}  map[string]string  "Task does not exist"
// @Failure      400  {object}  map[string]string  "Failed to stop the task"
// @Failure      401  {object}  map[string]string  "Invalid request or error message"
// @Router       /tasks/{id} [delete]
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