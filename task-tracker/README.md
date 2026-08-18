# Task Tracker

Simple CLI app to track your tasks and manage your to-do list. Inspired by [projects](https://roadmap.sh/projects/task-tracker) section in roadmap.sh. Developed using **TDD**.

## Features

- Add, update, delete.
- Priorities.
- Mark a task in-progress or done.
- Done tasks are saved in an archived file.
- List all tasks.
- List tasks by status.
- Emojis.
- Go standard libraries, no external dependencies.
- Your tasks, your law. Your tasks are stored locally.

## Prerequisites

- Go version `1.26`

## Installation

- Clone

```bash
git clone https://github.com/AlejandroArcoPu/go.git
cd task-tracker
```

- Build it

```bash
go build .
sudo mv ./task-traker /usr/local/bin/task-tracker
```

- Use it

```bash
task-tracker add "Buy milk"
```

## Usage

- Add a task

```bash
task-tracker add "Buy milk"
```

- Update a task

```bash
task-tracker update 1 "Buy milk and cook dinner"
```

- Delete a task

```bash
task-tracker delete 1
```

- List all tasks

```bash
task-tracker list
```

- List tasks by status

```bash
task-tracker list done
task-tracker list todo
task-tracker list in-progress
```

- Mark a task as done

```bash
task-tracker mark-done 1
```

- Mark a task as in-progress

```bash
task-tracker mark-in-progress 1
```
