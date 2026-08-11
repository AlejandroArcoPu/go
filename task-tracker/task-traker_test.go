package main

import (
	"bytes"
	"testing"
	"time"
)

func TestIsStatus(t *testing.T) {

	tests := []struct {
		name     string
		status   string
		expected bool
	}{
		{
			name:     "Buy milk",
			status:   "todo",
			expected: true,
		},
		{
			name:     "Gym",
			status:   "in-progress",
			expected: true,
		},
		{
			name:     "Walk",
			status:   "done",
			expected: true,
		},
		{
			name:     "Invented",
			status:   "test",
			expected: false,
		},
	}

	for _, te := range tests {
		if result := isStatus(te.status); result != te.expected {
			t.Errorf("isStatus(%s) = %v should be %v", te.status, result, te.expected)
		}
	}
}

func TestUpdateDescription(t *testing.T) {
	tests := []struct {
		task           Task
		newDescription string
	}{
		{
			task:           Task{ID: 1, Description: "Hanging out", Status: 0, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			newDescription: "Buy water",
		},
		{
			task:           Task{ID: 2, Description: "Buy food", Status: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			newDescription: "Gym",
		},
		{
			task:           Task{ID: 3, Description: "Drink 2l of water", Status: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			newDescription: "Buy tickets",
		},
	}

	for _, test := range tests {

		if test.task.UpdateDescription(test.newDescription); test.task.Description != test.newDescription {
			t.Errorf("The new description: %s is not on Task properties", test.newDescription)
		}
	}
}

func TestUpdateStatus(t *testing.T) {
	tests := []struct {
		task      Task
		newStatus TaskStatus
	}{
		{
			task:      Task{ID: 1, Description: "Hanging out", Status: 0, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			newStatus: 2,
		},
		{
			task:      Task{ID: 2, Description: "Buy food", Status: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			newStatus: 0,
		},
		{
			task:      Task{ID: 3, Description: "Drink 2l of water", Status: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			newStatus: 1,
		},
	}

	for _, test := range tests {

		if test.task.UpdateStatus(test.newStatus); test.task.Status != test.newStatus {
			t.Errorf("The new status: %s is not on Task properties", test.newStatus)
		}
	}
}

func TestRemoveIndex(t *testing.T) {
	tasks := []Task{
		{ID: 1, Description: "Hanging out", Status: 0, CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
		{ID: 2, Description: "Buy food", Status: 1, CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
		{ID: 3, Description: "Drink 2l of water", Status: 1, CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
	}
	tasks = removeIndex(tasks, 1)

	if len(tasks) != 2 {
		t.Errorf("len(tasks) = %d wants 2", len(tasks))
	}

	if tasks[0].ID != 1 {
		t.Errorf("tasks[0]: %d wants 1", tasks[0].ID)

	}

	if tasks[1].ID != 3 {
		t.Errorf("tasks[1]: %d wants 3", tasks[1].ID)
	}
}

func TestListTasks(t *testing.T) {
	tasks := []Task{
		{ID: 1, Description: "Hanging out", Status: 0, CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
	}
	expected := `🗒️ Listing tasks
[
{ ID: 1; Description: Hanging out, Status: todo, CreatedAt: 0001-01-01 00:00:00 +0000 UTC, UpdatedAt: 0001-01-01 00:00:00 +0000 UTC }
]
You have 1 tasks in your list.`
	var b bytes.Buffer
	if err := listTasks(&b, tasks); err != nil {
		t.Errorf("listTasks() error: %v", err)
	}

	if got := b.String(); got != expected {
		t.Errorf("listTask() = got %q, but expected: %q", got, expected)
	}
}

func TestListTasksBy(t *testing.T) {
	tasks := []Task{
		{ID: 1, Description: "Hanging out", Status: 1, CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
	}
	expected := `🗒️ Listing tasks by in-progress
{ ID: 1; Description: Hanging out, Status: in-progress, CreatedAt: 0001-01-01 00:00:00 +0000 UTC, UpdatedAt: 0001-01-01 00:00:00 +0000 UTC }
`
	var b bytes.Buffer

	if err := listTasksBy(&b, 1, tasks); err != nil {
		t.Errorf("listTasksBy() error: %v", err)
	}

	if got := b.String(); got != expected {
		t.Errorf("listTasksBy() = got %q, but expected: %q", got, expected)
	}
}

func TestAddTask(t *testing.T) {
	tasks := []Task{
		{ID: 1, Description: "Hanging out", Status: 1, CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
	}
	description := "Buy milk"

	tasks = addTask(description, tasks)

	if len(tasks) != 2 {
		t.Errorf("length of tasks after addTask() = got %d, but expected: 2", len(tasks))
	}

	if tasks[1].ID != 2 {
		t.Errorf("addTask() last item have ID %d, but expected: 2", tasks[1].ID)
	}
}
