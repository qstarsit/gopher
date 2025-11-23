package main

import (
	"time"
)

// Task represents a single task with a title, priority, and completion status
type Task struct {
	ID          int
	Title       string
	Priority    int // 1 (low) to 5 (high)
	IsCompleted bool
	CreatedAt   time.Time
}

// TaskManager manages a collection of tasks
type TaskManager struct {
	// TODO 1: Add a field to store tasks (hint: use a slice of Task pointers)
	nextID int
}

// NewTaskManager creates and returns a new TaskManager
func NewTaskManager() *TaskManager {
	// TODO 2: Initialize and return a TaskManager with an empty task slice
	// and nextID set to 1
	return nil
}

// AddTask creates a new task and adds it to the manager
func (tm *TaskManager) AddTask(title string, priority int) *Task {
	// TODO 3: Create a new task with:
	// - ID from tm.nextID (then increment nextID)
	// - The provided title and priority
	// - IsCompleted set to false
	// - CreatedAt set to current time (use time.Now())
	// Add it to the tasks slice and return the task

	return nil
}

// CompleteTask marks a task as completed by its ID
func (tm *TaskManager) CompleteTask(id int) error {
	// TODO 4: Find the task with the given ID and set IsCompleted to true
	// If the task is not found, return an error using fmt.Errorf
	// (hint: use a for loop to iterate through tasks)

	// Example Errorf call fmt.Errorf("task with ID %d not found", id)
	return nil
}

// GetTasksByPriority returns all tasks with the specified priority
func (tm *TaskManager) GetTasksByPriority(priority int) []*Task {
	// TODO 5: Create a slice to store matching tasks
	// Iterate through all tasks and add those matching the priority
	// Return the slice

	return nil
}

// PrintTask displays a task's information
func (t *Task) PrintTask() {
	status := "Pending"
	if t.IsCompleted {
		status = "Done"
	}

	// TODO 6: Use fmt.Printf to print the task in this format:
	// [ID] Title (Priority: X) - Status
	// Example: [1] Write code (Priority: 5) - Done
	// Hint: Use the 'status' variable defined above

	_ = status // Remove this line when you start
}

// GetTaskStats returns statistics about tasks as a map
func (tm *TaskManager) GetTaskStats() map[string]int {
	// TODO 7: Create and return a map with these keys:
	// - "total": total number of tasks
	// - "completed": number of completed tasks
	// - "pending": number of pending tasks
	// - "highPriority": number of tasks with priority >= 4

	return nil
}

// Stringer interface implementation
// String returns a formatted string representation of a Task
func (t *Task) String() string {
	// TODO 8: Return a string representation of the task
	// Format: "Task{ID: X, Title: 'Y', Priority: Z, Completed: true/false}"

	return ""
}
