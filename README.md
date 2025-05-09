# Todos CLI Application
A simple CLI application that organizes tasks for a specific project. It works in any folder, creating a tasks.json file in each folder to store the tasks. This means each folder will have its own separate set of tasks. For example, if you run the app in the /project1/ folder, it will create a tasks.json file there. If you switch to another folder and run the app again, a new tasks.json file will be created for that folder, allowing you to manage different sets of tasks for each project.

### List of Libraries Used:
- [cobra-cli, cobra](https://github.com/spf13/cobra)

### Features
#### Done
- [x] Create tasks
- [x] Assign statuses to tasks (todo, in-progress, done)
- [x] Delete tasks
- [x] Display the full list of tasks
- [x] Display in-progress tasks
- [x] Display completed (done) tasks

#### In-Progress
- [ ] Edit tasks

## Project Objective
I've been thinking of trying something new for organizing projects. Since each project typically lives in its own folder, I thought it would be useful to build a simple CLI app that manages to-do lists or tasks within each project directory.

The idea is that whenever I run this CLI tool inside a project folder, it should detect the current working directory and create (or use) a tasks.json file right there. This file will store all the tasks related specifically to that project, keeping things isolated and organized.

That way, every project can have its own independent task list without interfering with others a lightweight, local task manager per project.

## Personal Objective
I created this application to learn the Cobra library and practice my skills in Go. I really enjoy the language and want to become more comfortable with it. This project isn’t large; it’s small and simple, yet it covers a variety of concepts in Go and the Cobra library.

### Setup
There are two ways to set it up: either download the binary version or compile it from the source code.

- Clone repository

```bash
git clone https://github.com/omar0ali/todos-cli.git
```

Change path

```bash
cd /todos-cli
```

### Usage
```bash
Todos is a CLI application that helps users manage their to-do tasks. You can add, remove, edit tasks, assign statuses, and display the latest tasks.

Usage:
  todos [flags]
  todos [command]

Available Commands:
  add         add a task
  assign      Assign a new status to a task by its ID
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  remove      removing a task from a list
  show        Display the list of tasks

Flags:
  -h, --help      help for todos
  -v, --verbose   show debug messages

Use "todos [command] --help" for more information about a command.
```
