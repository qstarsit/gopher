package examples // inside this package to avoid naming conflicts in main

import (
	"fmt"
	"time"
)

type Task struct {
	ID          int
	Title       string
	Priority    int
	IsCompleted bool
	CreatedAt   time.Time
}

type TaskManager struct {
	tasks  []*Task
	nextID int
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks:  []*Task{},
		nextID: 1,
	}
}

func (tm *TaskManager) AddTask(title string, priority int) *Task {
	task := &Task{
		ID:          tm.nextID,
		Title:       title,
		Priority:    priority,
		IsCompleted: false,
		CreatedAt:   time.Now(),
	}

	tm.tasks = append(tm.tasks, task)
	tm.nextID++
	return task
}

func (tm *TaskManager) CompleteTask(id int) error {
	for _, task := range tm.tasks {
		if task.ID == id {
			task.IsCompleted = true
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}

func (tm *TaskManager) GetTasksByPriority(priority int) []*Task {
	var matchedTasks []*Task
	for _, task := range tm.tasks {
		if task.Priority == priority {
			matchedTasks = append(matchedTasks, task)
		}
	}

	return matchedTasks
}

func (t *Task) PrintTask() {
	status := "Pending"
	if t.IsCompleted {
		status = "Done"
	}

	fmt.Printf("[%d] %s (Priority: %d) - %s\n", t.ID, t.Title, t.Priority, status)
}

func (tm *TaskManager) GetTaskStats() map[string]int {
	stats := map[string]int{
		"total":        len(tm.tasks),
		"completed":    0,
		"pending":      0,
		"highPriority": 0,
	}

	for _, task := range tm.tasks {
		if task.IsCompleted {
			stats["completed"]++
		}
		if !task.IsCompleted {
			stats["pending"]++
		}
		if task.Priority >= 4 {
			stats["highPriority"]++
		}
	}

	return stats
}

func (t *Task) String() string {
	return fmt.Sprintf("Task{ID: %d, Title: '%s', Priority: %d, Completed: %t}", t.ID, t.Title, t.Priority, t.IsCompleted)
}
