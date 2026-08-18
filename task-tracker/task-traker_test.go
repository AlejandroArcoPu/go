package main

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

// import (
// 	"bytes"
// 	"encoding/json"
// 	"os"
// 	"path/filepath"
// 	"testing"
// 	"time"
// )

// func TestIsStatus(t *testing.T) {

// 	tests := []struct {
// 		name     string
// 		status   string
// 		expected bool
// 	}{
// 		{
// 			name:     "Buy milk",
// 			status:   "todo",
// 			expected: true,
// 		},
// 		{
// 			name:     "Gym",
// 			status:   "in-progress",
// 			expected: true,
// 		},
// 		{
// 			name:     "Walk",
// 			status:   "done",
// 			expected: true,
// 		},
// 		{
// 			name:     "Invented",
// 			status:   "test",
// 			expected: false,
// 		},
// 	}

// 	for _, te := range tests {
// 		if result := isStatus(te.status); result != te.expected {
// 			t.Errorf("isStatus(%s) = %v should be %v", te.status, result, te.expected)
// 		}
// 	}
// }

// func TestUpdateDescription(t *testing.T) {
// 	tests := []struct {
// 		task           Task
// 		newDescription string
// 	}{
// 		{
// 			task:           Task{ID: 1, Description: "Hanging out", Status: 0, CreatedAt: time.Now(), UpdatedAt: time.Now()},
// 			newDescription: "Buy water",
// 		},
// 		{
// 			task:           Task{ID: 2, Description: "Buy food", Status: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
// 			newDescription: "Gym",
// 		},
// 		{
// 			task:           Task{ID: 3, Description: "Drink 2l of water", Status: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
// 			newDescription: "Buy tickets",
// 		},
// 	}

// 	for _, test := range tests {

// 		if test.task.UpdateDescription(test.newDescription); test.task.Description != test.newDescription {
// 			t.Errorf("The new description: %s is not on Task properties", test.newDescription)
// 		}
// 	}
// }

// func TestUpdateTaskDescriptionNotFound(t *testing.T) {
// 	tasks := []Task{
// 		{ID: 1, Description: "Hanging out", Status: 1, CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
// 	}

// 	tasks, err := updateTaskDescription(2, "Buy milk", tasks)

// 	if err == nil {
// 		t.Errorf("updateTaskDescription() hasn't failed and it's expected a not found error")
// 	}

// }

// func TestUpdateStatus(t *testing.T) {
// 	tests := []struct {
// 		task      Task
// 		newStatus TaskStatus
// 	}{
// 		{
// 			task:      Task{ID: 1, Description: "Hanging out", Status: 0, CreatedAt: time.Now(), UpdatedAt: time.Now()},
// 			newStatus: 2,
// 		},
// 		{
// 			task:      Task{ID: 2, Description: "Buy food", Status: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
// 			newStatus: 0,
// 		},
// 		{
// 			task:      Task{ID: 3, Description: "Drink 2l of water", Status: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
// 			newStatus: 1,
// 		},
// 	}

// 	for _, test := range tests {

// 		if test.task.UpdateStatus(test.newStatus); test.task.Status != test.newStatus {
// 			t.Errorf("The new status: %s is not on Task properties", test.newStatus)
// 		}
// 	}
// }

// func TestUpdateTaskStatusNotFound(t *testing.T) {
// 	tasks := []Task{
// 		{ID: 1, Description: "Hanging out", Status: 0, CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
// 	}

// 	tasks, err := updateTaskStatus(2, 1, tasks)

// 	if err == nil {
// 		t.Errorf("updateTaskStatus() hasn't failed and it's expected a not found error")
// 	}
// }

// func TestRemoveIndex(t *testing.T) {
// 	tasks := []Task{
// 		{ID: 1, Description: "Hanging out", Status: 0, CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
// 		{ID: 2, Description: "Buy food", Status: 1, CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
// 		{ID: 3, Description: "Drink 2l of water", Status: 1, CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
// 	}
// 	tasks = removeIndex(tasks, 1)

// 	if len(tasks) != 2 {
// 		t.Errorf("len(tasks) = %d wants 2", len(tasks))
// 	}

// 	if tasks[0].ID != 1 {
// 		t.Errorf("tasks[0]: %d wants 1", tasks[0].ID)

// 	}

// 	if tasks[1].ID != 3 {
// 		t.Errorf("tasks[1]: %d wants 3", tasks[1].ID)
// 	}
// }

// func TestListTasks(t *testing.T) {
// 	tasks := []Task{
// 		{ID: 1, Description: "Hanging out", Status: 0, CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
// 	}
// 	expected := `🗒️ Listing tasks
// [
// { ID: 1; Description: Hanging out, Status: todo, CreatedAt: 0001-01-01 00:00:00 +0000 UTC, UpdatedAt: 0001-01-01 00:00:00 +0000 UTC }
// ]
// You have 1 tasks in your list.`
// 	var b bytes.Buffer
// 	if err := listTasks(&b, tasks); err != nil {
// 		t.Errorf("listTasks() error: %v", err)
// 	}

// 	if got := b.String(); got != expected {
// 		t.Errorf("listTask() = got %q, but expected: %q", got, expected)
// 	}
// }

// func TestListTasksBy(t *testing.T) {
// 	tasks := []Task{
// 		{ID: 1, Description: "Hanging out", Status: 1, CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
// 	}
// 	expected := `🗒️ Listing tasks by in-progress
// { ID: 1; Description: Hanging out, Status: in-progress, CreatedAt: 0001-01-01 00:00:00 +0000 UTC, UpdatedAt: 0001-01-01 00:00:00 +0000 UTC }
// `
// 	var b bytes.Buffer

// 	if err := listTasksBy(&b, 1, tasks); err != nil {
// 		t.Errorf("listTasksBy() error: %v", err)
// 	}

// 	if got := b.String(); got != expected {
// 		t.Errorf("listTasksBy() = got %q, but expected: %q", got, expected)
// 	}
// }

// func TestAddTask(t *testing.T) {
// 	tasks := []Task{
// 		{ID: 1, Description: "Hanging out", Status: 1, CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
// 	}
// 	description := "Buy milk"

// 	tasks = addTask(description, tasks)

// 	if len(tasks) != 2 {
// 		t.Errorf("length of tasks after addTask() = got %d, but expected: 2", len(tasks))
// 	}

// 	if tasks[1].ID != 2 {
// 		t.Errorf("addTask() last item have ID %d, but expected: 2", tasks[1].ID)
// 	}
// }

// func TestDeleteTask(t *testing.T) {
// 	tasks := []Task{
// 		{ID: 1, Description: "Hanging out", Status: 1, CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
// 		{ID: 2, Description: "Buy milk", Status: 0, CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
// 	}

// 	tasks, err := deleteTask(2, tasks)

// 	if err != nil {
// 		t.Errorf("deleteTask() has failed")
// 	}

// 	if len(tasks) != 1 {
// 		t.Errorf("length of tasks after deleteTask() = got %d, but expected: 1", len(tasks))
// 	}

// 	if tasks[0].ID != 1 {
// 		t.Errorf("deleteTask() unique item have ID %d, but expected: 1", tasks[1].ID)
// 	}
// }

// func TestWriteTasks(t *testing.T) {
// 	tasks := []Task{
// 		{ID: 1, Description: "Hanging out", Status: 1, CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
// 		{ID: 2, Description: "Buy milk", Status: 0, CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
// 	}
// 	path := filepath.Join(os.TempDir(), "tasks-write-test.json")

// 	if err := writeTasks(tasks, path); err != nil {
// 		t.Errorf("writeTasks() has failed: %q", err)
// 	}

// 	data, err := os.ReadFile(path)
// 	if err != nil {
// 		t.Errorf("os.Readfile() has failed: %q", err)
// 	}

// 	if len(data) == 0 {
// 		t.Errorf("length of data written by writeTasks() got: %d, but expected greater than 0", len(data))
// 	}
// 	var got []Task

// 	if err = json.Unmarshal(data, &got); err != nil {
// 		t.Errorf("json.Unmarshal() has failed: %q", err)
// 	}

// 	if len(tasks) != len(got) {
// 		t.Errorf("writeTasks() got %d , but expected: %d", len(got), len(tasks))
// 	}
// }

// func TestReadTask(t *testing.T) {
// 	data := []Task{
// 		{ID: 1, Description: "Hanging out", Status: 1, CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
// 		{ID: 2, Description: "Buy milk", Status: 0, CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
// 	}
// 	path := filepath.Join(os.TempDir(), "tasks-read-test.json")
// 	b, err := json.Marshal(data)

// 	if err != nil {
// 		t.Errorf("json.Marshal() has failed: %q", err)
// 	}

// 	if err = os.WriteFile(path, b, 0644); err != nil {
// 		t.Errorf("os.WriteFile() has failed: %q", err)
// 	}

// 	tasks, err := readTasks(path)

// 	if err != nil {
// 		t.Errorf("readTasks() has failed: %q", err)
// 	}

// 	if len(data) != len(tasks) {
// 		t.Errorf("length of readTasks() got: %d, but it's expected to be the same as the original", len(tasks))
// 	}

// 	if tasks[0].Description != "Hanging out" {
// 		t.Errorf("readTasks() first element got: %s description, expected Hanging out", tasks[0].Description)
// 	}
// }

func TestUpdateTaskStatus(t *testing.T) {

	t.Run("status exist", func(t *testing.T) {
		got := Task{ID: 1, Description: "Buy milk", Status: In_Progress}
		want := Task{ID: 1, Description: "Buy milk", Status: Done}

		err := got.UpdateStatus(Done)

		assertTask(t, got, want)
		assertError(t, err, nil)
	})

	t.Run("status doesn't exist", func(t *testing.T) {
		task := Task{ID: 1, Description: "Buy milk", Status: In_Progress}

		err := task.UpdateStatus(3)

		assertError(t, err, ErrStatusNotFound)
	})

}

func TestUpdateDescription(t *testing.T) {
	got := Task{ID: 1, Description: "Buy milk", Status: Todo}
	want := Task{ID: 1, Description: "Study Go", Status: Todo}

	got.UpdateDescription("Study Go")

	assertTask(t, got, want)
}

func TestAddTask(t *testing.T) {
	got := Tasks{}
	want := Tasks{
		{ID: 1, Description: "Study Go", Status: Todo},
	}

	got.Add("Study Go")

	assertTasks(t, got, want)
}

func TestDeleteTask(t *testing.T) {
	t.Run("existing task", func(t *testing.T) {
		got := Tasks{
			{ID: 1, Description: "Buy milk", Status: Done},
		}
		want := Tasks{}

		err := got.Delete(1)

		assertTasks(t, got, want)
		assertError(t, err, nil)
	})

	t.Run("no existing task", func(t *testing.T) {
		got := Tasks{}
		want := Tasks{}

		err := got.Delete(1)

		assertTasks(t, got, want)
		assertError(t, err, ErrTaskDoesNotExist)
	})
}

func TestList(t *testing.T) {
	t.Run("non empty list", func(t *testing.T) {
		buffer := bytes.Buffer{}
		tasks := Tasks{
			{ID: 1, Description: "Buy Milk", Status: Todo},
			{ID: 2, Description: "Study Go", Status: Done},
		}
		want := `You have 2 tasks in total.

ID: 1, Description: Buy Milk, Status: todo
ID: 2, Description: Study Go, Status: done
`
		tasks.List(&buffer)
		got := buffer.String()

		assertStrings(t, got, want)
	})

	t.Run("empty list", func(t *testing.T) {
		buffer := bytes.Buffer{}
		tasks := Tasks{}
		want := `You have 0 tasks in total.

Add a task using task-tracker add command.`
		tasks.List(&buffer)
		got := buffer.String()
		assertStrings(t, got, want)
	})
}

func TestListBy(t *testing.T) {

	t.Run("list by done", func(t *testing.T) {
		buffer := bytes.Buffer{}
		tasks := Tasks{
			{ID: 1, Description: "Buy Milk", Status: Todo},
			{ID: 2, Description: "Study Go", Status: Done},
		}
		want := `ID: 2, Description: Study Go, Status: done

You have 1 tasks done.`
		tasks.ListBy(&buffer, Done)
		got := buffer.String()

		assertStrings(t, got, want)
	})
	t.Run("list by in-progress", func(t *testing.T) {
		buffer := bytes.Buffer{}
		tasks := Tasks{
			{ID: 1, Description: "Buy Milk", Status: In_Progress},
			{ID: 2, Description: "Study Go", Status: Done},
		}
		want := `ID: 1, Description: Buy Milk, Status: in_progress

You have 1 tasks in_progress.`
		err := tasks.ListBy(&buffer, In_Progress)
		got := buffer.String()

		assertError(t, err, nil)
		assertStrings(t, got, want)
	})
	t.Run("invalid status", func(t *testing.T) {
		buffer := bytes.Buffer{}
		tasks := Tasks{
			{ID: 1, Description: "Buy Milk", Status: In_Progress},
		}
		var invented Status = 3

		got := tasks.ListBy(&buffer, invented)

		assertError(t, got, ErrStatusNotFound)
	})
}

func TestReadTasks(t *testing.T) {
	path := filepath.Join(os.TempDir(), "test-tasks.json")

	tasks, err := ReadTasks(path)

}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if !errors.Is(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertTask(t testing.TB, got, want Task) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v want %+v", got, want)
	}
}

func assertTasks(t testing.TB, got, want Tasks) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v want %+v", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
