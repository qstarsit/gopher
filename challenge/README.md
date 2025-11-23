# Workshop Challenge

Welcome to the Go Workshop Challenge!

## What's covered

- **Structs and methods** - Creating and working with custom types
- **Slices and maps** - Fundamental collection types
- **Goroutines and channels** - Concurrency
- **Error handling** - Proper error management
- **Interfaces** - Go's approach to polymorphism (Stringer interface\*)

\* The Stringer is an interface for defining the string format of values.

## Getting started

1. Open `challenge.go` in your editor
2. Look for `TODO` comments (numbered 1-8)
3. Complete each TODO section following the instructions
4. Run your code with: `go run main.go`
5. Fix any errors and iterate :-)

## TODO guide

These TODOs correspond directly to numbered comments in `challenge.go`.

1. **TODO 1** – Add a slice field to store tasks inside `TaskManager`.
2. **TODO 2** – Implement `NewTaskManager()` to initialize the slice and set `nextID = 1`.
3. **TODO 3** – Implement `AddTask` to create a task, append it, and return it.
4. **TODO 4** – Implement `CompleteTask` with proper error handling if ID not found.
5. **TODO 5** – Implement `GetTasksByPriority` to return matching tasks.
6. **TODO 6** – Implement `PrintTask` to format and display task info.
7. **TODO 7** – Implement `GetTaskStats` to build a map of summary stats.
8. **TODO 8** – Implement `String()` (Stringer interface) for pretty `fmt.Println(task)` output.

## Tips

- Run `go run challenge.go` frequently to test your progress
- Use `go fmt challenge.go` to automatically format your code
