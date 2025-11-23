package main

import "fmt"

func main() {

	// Attempt to create manager
	tm := NewTaskManager()
	if tm == nil {
		fmt.Println("NewTaskManager not implemented yet (TODO 2). Finish TODOs 1 & 2 first.")
		return
	}

	// Add sample tasks
	titles := []string{"Learn Go", "Writing functions", "Work with maps"}
	priorities := []int{5, 3, 4}
	var created []*Task
	for i, title := range titles {
		t := tm.AddTask(title, priorities[i])
		if t == nil {
			fmt.Println("AddTask not implemented yet (TODO 3).")
			return
		}
		created = append(created, t)
	}

	// Complete first task
	if err := tm.CompleteTask(created[0].ID); err != nil {
		fmt.Println("CompleteTask pending (TODO 4):", err)
	}

	// Filter by priority
	high := tm.GetTasksByPriority(5)
	if high == nil {
		fmt.Println("GetTasksByPriority not implemented yet (TODO 5).")
	} else {
		fmt.Println("High priority tasks:")
		for _, t := range high {
			if t != nil {
				// PrintTask may still be TODO 6
				before := t.String()
				t.PrintTask()
				after := t.String()
				if before == after && after == "" {
					fmt.Println("String() not implemented yet (TODO 8).")
				}
			}
		}
	}

	// Stats
	stats := tm.GetTaskStats()
	if stats == nil {
		fmt.Println("GetTaskStats not implemented yet (TODO 7).")
	} else {
		fmt.Println("Stats:")
		for k, v := range stats {
			fmt.Printf("  %s: %d\n", k, v)
		}
	}

	fmt.Println("You've completed the challenge tasks!")
}
