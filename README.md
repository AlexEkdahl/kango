![Go version](https://img.shields.io/badge/Go-1.20-blue.svg)

# KanGoBoard - Task Manager
<p align="left">
  <img src="https://img.bigdaddylongleg.com/img/8d3d24e6543c46fdc0e0ed511d710a1b7fe183e0.png"  width="45%" />
</p>

`KanGoBoard` is a simple and efficient Kanban manager written in Go. With KanGoBoard, you can manage your tasks in three different statuses: To Do, In Progress, and Done, and navigate through your task lists using arrow keys or Vim key bindings.

# Features

Manage tasks in three different statuses: To Do, In Progress, and Done
Navigate through task lists using arrow keys or Vim key bindings
Move tasks to the next status with a single key press
Add new tasks with title and description
Delete tasks easily

# Installation

To install this Kaban manager, you need to have Go installed on your system. You can download Go from here.

Once Go is installed, run the following command to install the Task Manager:

```sh
go get github.com/AlexEkdahl/kango
```

After installation, you can run the `kango` with the following command:

```sh
kango
```

## Navigation

Use the following keys to move around the task lists:

- Left Arrow / h: Move to the previous task list
- Right Arrow / l: Move to the next task list
- Up Arrow / k: Move up within the current task list
- Down Arrow / j: Move down within the current task list

## Task Management

Use the following keys to manage tasks:

- Enter: Move a task to the next status
- d: Delete a task
- n: Add a new task
- e: Edit a task

## Usage

Once you have started the Task Manager, you will be presented with a Kanban board containing three columns: To Do, In Progress, and Done. You can use the arrow keys or Vim key bindings to move between columns and select tasks.

To move a task to the next status, simply press the `Enter` key while the task is selected. To delete a task, press the `d` key while the task is selected.

To add a new task, press the `n` key. This will open a new form where you can enter the title and description of the task.

To edit an existing task, press the `e` key while the task is selected. This will open a form where you can edit the title and description of the task.

## Quitting the Application

Use the following keys to quit the task manager:

- q: Quit the task manager
- Ctrl+C: Quit the task manager (alternative way)

# License

This project is licensed under the MIT License - see the LICENSE file for details.

# Contributing

Contributions are welcome! If you have any ideas, feature requests, or bug reports, please open an issue or submit a pull request.


Support

If you have any questions or issues with the Task Manager, please feel free to open an issue on the GitHub repository or reach out to the author directly.
