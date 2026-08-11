# Task Tracker

Simple CLI app to track your tasks and manage your to-do list. Inspired by https://roadmap.sh/projects/task-tracker section in roadmap.sh.

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
task-cli add "Buy milk"
```

- Update a task

```bash
task-cli update 1 "Buy milk and cook dinner"
```

- Delete a task

```bash
task-cli delete 1
```

- List all tasks

```bash
task-cli list
```

- List tasks by status

```bash
task-cli list done
task-cli list todo
task-cli list in-progress
```

- Mark a task as done

```bash
task-cli mark-done 1
```

- Mark a task as in-progress

```bash
task-cli mark-in-progress 1
```
