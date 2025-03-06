package tasks

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"sync"
	"time"
)

const MAX_TASKS = 2

type Task struct {
	ID       string        `json:"id"`
	Status   string        `json:"status"` // pending, in_progress, done, error
	Filename string        `json:"filename"`
	Stop     chan struct{} `json:"-"`
}

var tasks = make(map[string]*Task)
var mutex = &sync.Mutex{}

func CreateTask() (string, error) {
	if len(tasks) == MAX_TASKS {
		return "", errors.New("Tasks reach maximum")
	}

	taskID := generateID()
	task := &Task{
		ID:     taskID,
		Status: "pending",
	}
	mutex.Lock()
	tasks[taskID] = task
	task.Stop = make(chan struct{})
	mutex.Unlock()
	log.Println("Task " + taskID + " created")
	return taskID, nil
}

func GetTask(taskID string) *Task {
	mutex.Lock()
	defer mutex.Unlock()
	return tasks[taskID]
}

func RunTask(taskID string) {
	task := GetTask(taskID)
	if task == nil {
		return
	}

	task.Status = "in_progress"
	filename := "export_" + taskID + ".json"
	task.Filename = filename
	log.Printf("task " + taskID + " running")
	log.Printf("task " + taskID + " filename " + filename)

	for i := 0; i < 5; i++ {
		select {
		case <-task.Stop:
			log.Println("Task " + taskID + " stopped")
			return
		default:
			time.Sleep(1 * time.Second)
		}
	}

	// Запись данных в файл
	file, err := os.Create(filename)
	if err != nil {
		task.Status = "error"
		return
	}
	defer file.Close()

	data := map[string]string{"message": "Data exported successfully"}
	json.NewEncoder(file).Encode(data)

	task.Status = "done"
	log.Printf("task " + taskID + " ended")
}

func generateID() string {
	return time.Now().Format("20060102150405")
}

func StopTask(task *Task) error {
	if task.Status == "in_progress" {
		task.Status = "stopped"
		close(task.Stop)
		return nil
	} else {
		return errors.New("Task is not in progress")
	}
}