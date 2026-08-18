package main

import (
	"fmt"
	"io"
)

// import (
// 	"encoding/json"
// 	"flag"
// 	"fmt"
// 	"io"
// 	"os"
// 	"path/filepath"
// 	"strconv"
// 	"time"
// )

// type TaskStatus int

// const (
// 	Todo TaskStatus = iota
// 	In_Progress
// 	Done
// )

// var statusName = map[TaskStatus]string{
// 	Todo:        "todo",
// 	In_Progress: "in-progress",
// 	Done:        "done",
// }

// var statusKey = map[string]TaskStatus{
// 	"todo":        Todo,
// 	"in-progress": In_Progress,
// 	"done":        Done,
// }

// var (
// 	usage = `
// 🗒️ Task Tracker

// Specify a command to execute:
// - add <description>: Add a new task
// - update <id> <description>: Update a task
// - delete <id>: Delete a task
// - mark-done <id>: Mark a task as done
// - mark-in-progress <id>: Mark a task as in progress
// - list: List all tasks
// - list <status>: List task by status (todo, in-progress, done)
// `
// )

// type Task struct {
// 	ID          int
// 	Description string
// 	Status      TaskStatus
// 	CreatedAt   time.Time
// 	UpdatedAt   time.Time
// }

// func isStatus(status string) bool {
// 	for _, name := range statusName {
// 		if name == status {
// 			return true
// 		}
// 	}

// 	return false
// }

// func (t Task) String() string {
// 	return fmt.Sprintf("\n{ ID: %v; Description: %v, Status: %v, CreatedAt: %v, UpdatedAt: %v }\n", t.ID, t.Description, fmt.Sprintf("%v", t.Status), t.CreatedAt, t.UpdatedAt)
// }

// func (ts TaskStatus) String() string {
// 	return statusName[ts]
// }

// func (t *Task) UpdateDescription(description string) {
// 	t.Description = description
// 	t.UpdatedAt = time.Now()
// }

// func (t *Task) UpdateStatus(status TaskStatus) {
// 	t.Status = status
// 	t.UpdatedAt = time.Now()
// }

// func removeIndex(tasks []Task, index int) []Task {
// 	return append(tasks[:index], tasks[index+1:]...)
// }

// func readTasks(path string) ([]Task, error) {
// 	var tasks []Task

// 	_, err := os.Stat(path)

// 	if os.IsNotExist(err) {
// 		os.Create(path)
// 	}

// 	data, err := os.ReadFile(path)
// 	if err != nil {
// 		return nil, fmt.Errorf("reading file: %w", err)
// 	}

// 	if len(data) == 0 {
// 		return []Task{}, nil
// 	}

// 	if err = json.Unmarshal(data, &tasks); err != nil {
// 		return nil, fmt.Errorf("decoding json: %w", err)
// 	}

// 	return tasks, nil
// }

// func writeTasks(tasks []Task, path string) error {
// 	b, err := json.Marshal(tasks)
// 	if err != nil {
// 		return fmt.Errorf("encoding tasks: %w", err)
// 	}
// 	err = os.WriteFile(path, b, 0644)
// 	if err != nil {
// 		return fmt.Errorf("saving tasks in file: %w", err)
// 	}
// 	return nil
// }

// func addTask(description string, tasks []Task) []Task {
// 	newId := 1
// 	if len(tasks) > 0 {
// 		newId = tasks[len(tasks)-1].ID + 1
// 	}
// 	newTask := Task{
// 		ID:          newId,
// 		Description: description,
// 		Status:      0,
// 		CreatedAt:   time.Now(),
// 		UpdatedAt:   time.Now(),
// 	}
// 	tasks = append(tasks, newTask)
// 	return tasks
// }

// func deleteTask(id int, tasks []Task) ([]Task, error) {
// 	found := false

// 	for i, task := range tasks {
// 		if task.ID == id {
// 			tasks = removeIndex(tasks, i)
// 			found = true
// 			break
// 		}
// 	}

// 	if !found {
// 		return nil, fmt.Errorf("task %d not found", id)
// 	}

// 	return tasks, nil
// }

// func listTasks(w io.Writer, tasks []Task) error {
// 	fmt.Fprintln(w, "🗒️ Listing tasks")
// 	fmt.Fprintln(w, tasks)
// 	fmt.Fprintf(w, "You have %d tasks in your list.", len(tasks))
// 	return nil
// }

// func listTasksBy(w io.Writer, status TaskStatus, tasks []Task) error {
// 	fmt.Fprintf(w, "🗒️ Listing tasks by %s", status)
// 	for _, t := range tasks {
// 		if t.Status == status {
// 			fmt.Fprint(w, t)
// 		}
// 	}
// 	return nil
// }

// func updateTaskDescription(id int, description string, tasks []Task) ([]Task, error) {
// 	found := false

// 	for i, t := range tasks {
// 		if t.ID == id {
// 			tasks[i].UpdateDescription(description)
// 			found = true
// 			break
// 		}
// 	}

// 	if !found {
// 		return nil, fmt.Errorf("task %d not found", id)
// 	}

// 	return tasks, nil
// }

// func updateTaskStatus(id int, status TaskStatus, tasks []Task) ([]Task, error) {
// 	found := false

// 	for i, t := range tasks {
// 		if t.ID == id {
// 			tasks[i].UpdateStatus(status)
// 			found = true
// 			break
// 		}
// 	}

// 	if !found {
// 		return nil, fmt.Errorf("task %d not found", id)
// 	}

// 	return tasks, nil
// }

// func executeCommand(command string, args []string, path string) error {
// 	switch command {
// 	case "list":
// 		if len(args) == 0 {
// 			tasks, err := readTasks(path)
// 			if err != nil {
// 				return fmt.Errorf("😭 List command has failed: %w", err)
// 			}
// 			if err := listTasks(os.Stdout, tasks); err != nil {
// 				return fmt.Errorf("😭 List command has failed: %w", err)
// 			}
// 		} else if len(args) == 1 {
// 			if !isStatus(args[0]) {
// 				return fmt.Errorf("😭 List command has failed: %s is not a valid status (todo, in-progress, done)", args[0])
// 			}
// 			tasks, err := readTasks(path)
// 			if err != nil {
// 				return fmt.Errorf("😭 List command has failed: %w", err)
// 			}
// 			status := statusKey[args[0]]
// 			if err = listTasksBy(os.Stdout, status, tasks); err != nil {
// 				return fmt.Errorf("😭 List command has failed: %w", err)
// 			}
// 		} else {
// 			return fmt.Errorf("❌ Invalid quantity of arguments: %d\nlist have 0 or 1 argument\n", len(args))
// 		}
// 		return nil
// 	case "add":
// 		if len(args) != 1 {
// 			return fmt.Errorf("❌ Invalid quantity of arguments: %d\nadd needs a description\n", len(args))
// 		}
// 		fmt.Println("➕ Adding a new task")
// 		description := args[0]
// 		tasks, err := readTasks(path)
// 		if err != nil {
// 			return fmt.Errorf("😭 Add command has failed: %w", err)
// 		}
// 		tasks = addTask(description, tasks)
// 		if err = writeTasks(tasks, path); err != nil {
// 			return fmt.Errorf("writing task: %w", err)
// 		}
// 		fmt.Printf("✅ Task added succesfully (ID: %d)", tasks[len(tasks)-1].ID)

// 		return nil
// 	case "update":
// 		if len(args) != 2 {
// 			return fmt.Errorf("❌ Invalid quantity of arguments: %d\nupdate needs an ID and a new description\n", len(args))
// 		}
// 		fmt.Println("🔄 Updating task")
// 		id, _ := strconv.Atoi(args[0])
// 		description := args[1]
// 		tasks, err := readTasks(path)
// 		if err != nil {
// 			return fmt.Errorf("😭 Update command has failed: %w", err)
// 		}
// 		tasks, err = updateTaskDescription(id, description, tasks)
// 		if err != nil {
// 			return fmt.Errorf("😭 Update command has failed: %w", err)
// 		}
// 		if err = writeTasks(tasks, path); err != nil {
// 			return fmt.Errorf("😭 Update command has failed: %w", err)
// 		}
// 		fmt.Printf("✅ Task updated succesfully (ID: %d)", id)
// 		return nil
// 	case "delete":
// 		if len(args) != 1 {
// 			return fmt.Errorf("❌ Invalid quantity of arguments: %d\ndelete needs an ID\n", len(args))
// 		}
// 		fmt.Println("🗑️ Deleting a task")
// 		id, _ := strconv.Atoi(args[0])
// 		tasks, err := readTasks(path)
// 		if err != nil {
// 			return fmt.Errorf("😭 Delete command has failed: %w", err)
// 		}

// 		tasks, err = deleteTask(id, tasks)
// 		if err != nil {
// 			return fmt.Errorf("😭 Delete command has failed: %w", err)
// 		}

// 		if err = writeTasks(tasks, path); err != nil {
// 			return fmt.Errorf("😭 Delete command has failed: %w", err)
// 		}
// 		fmt.Printf("✅ Task deleted succesfully (ID: %d)", id)
// 		return nil

// 	case "mark-in-progress":
// 		if len(args) != 1 {
// 			return fmt.Errorf("❌ Invalid quantity of arguments: %d\nmark-in-progress needs an id\n", len(args))
// 		}
// 		fmt.Println("🔄 Updating task")
// 		id, _ := strconv.Atoi(args[0])
// 		tasks, err := readTasks(path)
// 		if err != nil {
// 			return fmt.Errorf("😭 Delete command has failed: %w", err)
// 		}

// 		tasks, err = updateTaskStatus(id, 1, tasks)
// 		if err != nil {
// 			return fmt.Errorf("😭 mark-in-progress command has failed: %w", err)
// 		}

// 		if err = writeTasks(tasks, path); err != nil {
// 			return fmt.Errorf("😭 Delete command has failed: %w", err)
// 		}
// 		fmt.Printf("✅ Task status updated succesfully (ID: %d)", id)
// 		return nil

// 	case "mark-done":
// 		if len(args) != 1 {
// 			return fmt.Errorf("❌ Invalid quantity of arguments: %d\nmark-done needs an id\n", len(args))
// 		}
// 		fmt.Println("🔄 Updating task")
// 		id, _ := strconv.Atoi(args[0])
// 		tasks, err := readTasks(path)
// 		if err != nil {
// 			return fmt.Errorf("😭 Delete command has failed: %w", err)
// 		}

// 		tasks, err = updateTaskStatus(id, 2, tasks)
// 		if err != nil {
// 			return fmt.Errorf("😭 mark-done command has failed: %w", err)
// 		}
// 		fmt.Printf("✅ Task status updated succesfully (ID: %d)", id)
// 		return nil
// 	default:
// 		return fmt.Errorf("❌ Invalid command: %s\n%s\n", command, usage)
// 	}
// }

// func main() {
// 	flag.Parse()
// 	command := flag.Arg(0)
// 	args := flag.Args()[1:]
// 	path := filepath.Join(os.TempDir(), "tasks.json")
// 	err := executeCommand(command, args, path)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// }

type Status int

const (
	Todo Status = iota
	In_Progress
	Done
)

var statuses = map[Status]string{Todo: "todo", In_Progress: "in_progress", Done: "done"}

func (s Status) String() string {
	return statuses[s]
}

const (
	ErrStatusNotFound   = TaskErr("could not find task status, you should use todo, in_progress or done")
	ErrTaskNotFound     = TaskErr("could not find the task you are searching for")
	ErrTaskDoesNotExist = TaskErr("cannot perform operation on task because it does not exist")
)

type TaskErr string

func (e TaskErr) Error() string {
	return string(e)
}

type Task struct {
	ID          int
	Description string
	Status      Status
}

func (t *Task) UpdateStatus(status Status) error {
	if _, ok := statuses[status]; !ok {
		return ErrStatusNotFound
	}
	t.Status = status
	return nil
}

func (t *Task) UpdateDescription(description string) {
	t.Description = description
}

func (t *Task) String() string {
	return fmt.Sprintf("ID: %d, Description: %s, Status: %s", t.ID, t.Description, t.Status)
}

type Tasks []Task

func (tasks *Tasks) Add(description string) {
	task := Task{
		ID:          len(*tasks) + 1,
		Description: description,
		Status:      Todo,
	}
	*tasks = append(*tasks, task)
}

func (tasks *Tasks) Search(id int) (index int, err error) {
	for index, task := range *tasks {
		if id == task.ID {
			return index, nil
		}
	}
	return -1, ErrTaskNotFound
}

func (tasks *Tasks) Delete(id int) error {
	index, err := tasks.Search(id)
	switch err {
	case ErrTaskNotFound:
		return ErrTaskDoesNotExist
	case nil:
		(*tasks)[index] = (*tasks)[len(*tasks)-1]
		(*tasks) = (*tasks)[:len(*tasks)-1]
	default:
		return err
	}
	return nil
}

func (tasks *Tasks) List(w io.Writer) {
	length := len(*tasks)
	fmt.Fprintf(w, "You have %d tasks in total.\n\n", length)
	if length == 0 {
		fmt.Fprintf(w, "Add a task using task-tracker add command.")
	} else {
		for _, task := range *tasks {
			fmt.Fprintf(w, "%s\n", task.String())
		}
	}
}

func (tasks *Tasks) ListBy(w io.Writer, status Status) error {
	if _, ok := statuses[status]; !ok {
		return ErrStatusNotFound
	}

	count := 0

	for _, task := range *tasks {
		if task.Status == status {
			count++
			fmt.Fprintf(w, "%s\n", task.String())
		}
	}
	fmt.Fprintf(w, "\nYou have %d tasks %s.", count, status.String())
	return nil
}
