package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type TaskStatus int

const (
	Todo TaskStatus = iota
	In_Progress
	Done
)

var statusName = map[TaskStatus]string{
	Todo:        "todo",
	In_Progress: "in-progress",
	Done:        "done",
}

var path = filepath.Join(os.TempDir(), "tasks.json")

var (
	usage = `
🗒️ Task Tracker

Specify a command to execute:
- add <description>: Add a new task
- update <id> <description>: Update a task
- delete <id>: Delete a task
- mark-done <id>: Mark a task as done
- mark-in-progress <id>: Mark a task as in progress
- list: List all tasks
- list <status>: List task by status (todo, in-progress, done)
`
)

type Task struct {
	ID          int
	Description string
	Status      TaskStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func isStatus(status string) bool {
	return statusName[0] == status || statusName[1] == status || statusName[2] == status
}

func getStatus(status string) (TaskStatus, error) {
	switch status {
	case statusName[0]:
		return 0, nil
	case statusName[1]:
		return 1, nil
	case statusName[2]:
		return 2, nil
	default:
		return 0, fmt.Errorf("invalid status (todo, in-progress, done)")
	}
}

func (t Task) String() string {
	return fmt.Sprintf("\n{ ID: %v; Description: %v, Status: %v, CreatedAt: %v, UpdatedAt: %v }\n", t.ID, t.Description, fmt.Sprintf("%v", t.Status), t.CreatedAt, t.UpdatedAt)
}

func (ts TaskStatus) String() string {
	return statusName[ts]
}

func (t *Task) UpdateDescription(description string) {
	t.Description = description
	t.UpdatedAt = time.Now()
}

func (t *Task) UpdateStatus(status TaskStatus) {
	t.Status = status
	t.UpdatedAt = time.Now()
}

func removeIndex(tasks []Task, index int) []Task {
	return append(tasks[:index], tasks[index+1:]...)
}

func getTasks() ([]Task, error) {
	var tasks []Task
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	if err = json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("decoding json: %w", err)
	}

	return tasks, nil
}

func writeTasks(tasks []Task) error {
	b, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("encoding tasks: %w", err)
	}
	err = os.WriteFile(path, b, 0644)
	if err != nil {
		return fmt.Errorf("saving tasks in file: %w", err)
	}
	return nil
}

func addTask(description string) error {
	fmt.Println("➕ Adding a new task")
	tasks, err := getTasks()
	if err != nil {
		return fmt.Errorf("reading existing tasks: %w", err)
	}
	newId := 1
	if len(tasks) > 0 {
		newId = tasks[len(tasks)-1].ID + 1
	}
	newTask := Task{
		ID:          newId,
		Description: description,
		Status:      0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasks = append(tasks, newTask)

	if err = writeTasks(tasks); err != nil {
		return fmt.Errorf("writing task: %w", err)
	}

	fmt.Printf("✅ Task added succesfully (ID: %d)", newTask.ID)
	return nil
}

func deleteTask(id int) error {
	fmt.Println("🗑️ Deleting a task")
	tasks, err := getTasks()

	if err != nil {
		return fmt.Errorf("reading existing tasks: %w", err)
	}

	found := false

	for i, task := range tasks {
		if task.ID == id {
			tasks = removeIndex(tasks, i)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("task %d not found", id)
	}

	if err = writeTasks(tasks); err != nil {
		return fmt.Errorf("writing task: %w", err)
	}

	fmt.Printf("✅ Task deleted succesfully (ID: %d)", id)
	return nil
}

func listTasks() error {
	fmt.Println("🗒️ Listing tasks")
	tasks, err := getTasks()
	if err != nil {
		return fmt.Errorf("reading existing tasks: %w", err)
	}
	_, err = fmt.Println(tasks)
	fmt.Printf("You have %d tasks in your list!", len(tasks))
	return err
}

func listTasksBy(status TaskStatus) error {
	fmt.Printf("🗒️ Listing tasks by %s", status)
	tasks, err := getTasks()
	if err != nil {
		return fmt.Errorf("reading existing tasks: %w", err)
	}
	for _, t := range tasks {
		if t.Status == status {
			fmt.Println(t)
		}
	}
	return nil
}

func updateTaskDescription(id int, description string) error {
	fmt.Println("🔄 Updating task")
	tasks, err := getTasks()
	if err != nil {
		return fmt.Errorf("reading existing tasks: %w", err)
	}
	found := false

	for i, t := range tasks {
		if t.ID == id {
			tasks[i].UpdateDescription(description)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("task %d not found", id)
	}

	if err = writeTasks(tasks); err != nil {
		return fmt.Errorf("writing tasks %w", err)
	}
	fmt.Printf("✅ Task updated succesfully (ID: %d)", id)
	return nil
}

func updateTaskStatus(id int, status TaskStatus) error {
	fmt.Println("🔄 Updating task")
	tasks, err := getTasks()
	if err != nil {
		return fmt.Errorf("reading existing tasks: %w", err)
	}
	found := false

	for i, t := range tasks {
		if t.ID == id {
			tasks[i].UpdateStatus(status)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("task %d not found", id)
	}

	if err = writeTasks(tasks); err != nil {
		return fmt.Errorf("writing tasks %w", err)
	}
	fmt.Printf("✅ Task updated succesfully (ID: %d)", id)
	return nil
}

func executeCommand(command string, args []string) error {
	switch command {
	case "list":
		if len(args) > 1 {
			return fmt.Errorf("❌ Invalid quantity of arguments: %d\nlist have 0 or 1 argument\n", len(args))
		}
		if len(args) == 0 {
			if err := listTasks(); err != nil {
				return fmt.Errorf("😭 List command has failed: %w", err)
			}
		} else {
			if !isStatus(args[0]) {
				return fmt.Errorf("😭 List command has failed: %s is not a valid status (todo, in-progress, done)", args[0])
			}
			status, err := getStatus(args[0])
			if err != nil {
				return fmt.Errorf("😭 List command has failed: %s is not a valid status (todo, in-progress, done)", args[0])
			}
			if err = listTasksBy(status); err != nil {
				return fmt.Errorf("😭 List command has failed: %w", err)
			}
		}
		return nil
	case "add":
		if len(args) != 1 {
			return fmt.Errorf("❌ Invalid quantity of arguments: %d\nadd needs a description\n", len(args))
		}
		description := args[0]
		if err := addTask(description); err != nil {
			return fmt.Errorf("😭 Add command has failed: %w", err)
		}
		return nil
	case "update":
		if len(args) != 2 {
			return fmt.Errorf("❌ Invalid quantity of arguments: %d\nupdate needs an ID and a new description\n", len(args))
		}
		id, _ := strconv.Atoi(args[0])
		description := args[1]
		if err := updateTaskDescription(id, description); err != nil {
			return fmt.Errorf("😭 Add command has failed: %w", err)
		}
		return nil
	case "delete":
		if len(args) != 1 {
			return fmt.Errorf("❌ Invalid quantity of arguments: %d\ndelete needs an ID\n", len(args))
		}
		id, _ := strconv.Atoi(args[0])
		if err := deleteTask(id); err != nil {
			return fmt.Errorf("😭 Delete command has failed: %w", err)
		}
		return nil
	case "mark-in-progress":
		if len(args) != 1 {
			return fmt.Errorf("❌ Invalid quantity of arguments: %d\nmark-in-progress needs an id\n", len(args))
		}
		id, _ := strconv.Atoi(args[0])
		if err := updateTaskStatus(id, 1); err != nil {
			return fmt.Errorf("😭 mark-in-progress command has failed: %w", err)
		}
		return nil
	case "mark-done":
		if len(args) != 1 {
			return fmt.Errorf("❌ Invalid quantity of arguments: %d\nmark-done needs an id\n", len(args))
		}
		id, _ := strconv.Atoi(args[0])
		if err := updateTaskStatus(id, 2); err != nil {
			return fmt.Errorf("😭 mark-done command has failed: %w", err)
		}
		return nil
	default:
		return fmt.Errorf("❌ Invalid command: %s\n%s\n", command, usage)
	}
}

func main() {
	flag.Parse()
	command := flag.Arg(0)
	args := flag.Args()[1:]
	err := executeCommand(command, args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
