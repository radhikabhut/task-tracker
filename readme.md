# Task Tracker

Sample solution for the task-tracker challenge from roadmap.sh.

A simple task tracker CLI application written in Go. The application allows users to manage their tasks by adding, updating, deleting, and listing them. Tasks are persisted in a JSON file (`task.json`).

---

## Features
- Add tasks with a description and status.
- Update task details, including description and status.
- Delete tasks by their ID.
- List tasks with optional status filtering.

---

## Requirements
- Go (version 1.18 or later)

---

## Setup Instructions

1. **Clone or Download the Repository:**
   ```bash
   git clone <repository-url>
   cd <repository-directory>
   ```

2. **Ensure Go is Installed:**
   Check your Go installation by running:
   ```bash
   go version
   ```

3. **Build the Program:**
   Build the application to create an executable:
   ```bash
   go build tasktracker.go
   ```

   This will generate an executable named `tasktracker.exe` (on Windows) or `tasktracker` (on Linux/macOS).

4. **Initialize `task.json` File:**
   Create a `task.json` file in the same directory as the executable and add the following content:
   ```json
   {
     "tasks": []
   }
   ```

---

## Usage

### General Command Syntax
```bash
./tasktracker <command> [arguments]
```

### Commands
[.](https://roadmap.sh/projects/task-tracker)
#### 1. Add a Task
Add a new task with a description and status (`todo`, `in-progress`, or `done`):
```bash
./tasktracker add "<description>" "<status>"
```
Example:
```bash
./tasktracker add "Learn Go" "todo"
```

#### 2. Update a Task
Update the description and/or status of an existing task:
```bash
./tasktracker update <id> "<new description>" "<new status>"
```
Example:
```bash
./tasktracker update 1 "Learn Go Basics" "in-progress"
```

#### 3. Delete a Task
Delete a task by its ID:
```bash
./tasktracker delete <id>
```
Example:
```bash
./tasktracker delete 1
```

#### 4. List Tasks
List all tasks or filter by a specific status:
```bash
./tasktracker list [status]
```
Examples:
- List all tasks:
  ```bash
  ./tasktracker list
  ```
- List tasks with status `todo`:
  ```bash
  ./tasktracker list todo
  ```

---

## Error Handling
- Ensure the `task.json` file exists and is properly formatted.
- If the file is missing, the program will create a new one.
- If you encounter errors such as "unexpected end of JSON input," initialize the file with the content:
  ```json
  {
    "tasks": []
  }
  ```

---

## File Structure
```
workspace/
├── tasktracker.go   # Main application source code
├── task.json        # JSON file to store tasks
├── go.mod           # Go module file
```

---

## Example Workflow
1. Add a task:
   ```bash
   ./tasktracker add "Learn Go" "todo"
   ```
2. List tasks:
   ```bash
   ./tasktracker list
   ```
3. Update a task:
   ```bash
   ./tasktracker update 1 "Learn Advanced Go" "in-progress"
   ```
4. Delete a task:
   ```bash
   ./tasktracker delete 1
   ```

---

## License
This project is open-source and free to use under the MIT License.

