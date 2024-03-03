## gogo - Todo List Manager CLI Tool

gogo is a simple CLI-based todo list manager written in Go. It allows users to add, update, delete, and list tasks from the command line.

**Features:**
- Add tasks with title, priority, and category.
- Update existing tasks.
- Delete tasks.
- List all tasks.

**Installation:**
1. Download the latest release from everest1508/gogo-cli-tool page.
2. Extract the downloaded ZIP file.
3. Move the `gogo.exe` executable to a directory included in your system's PATH environment variable.

Alternatively, you can build `gogo` from source by cloning this repository and running `go build`.

```bash
git clone https://github.com/your_username/gogo.git
cd gogo
go build

Usage:
gogo <command>

Commands:
- add      Add a new task
- update   Update an existing task
- delete   Delete a task
- list     List all tasks

To get more information about a specific command, run `gogo <command> -h`.

Examples:
- Add a task:
  
  gogo add

- Update a task:
  
  gogo update

- Delete a task:
  
  gogo delete

- List all tasks:
  
  gogo list
