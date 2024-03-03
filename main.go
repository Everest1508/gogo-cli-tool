package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	ID        int
	Title     string
	Priority  string
	Category  string
	Completed bool
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./tasks.sqlite3")
	if err != nil {
		fmt.Println("Error opening database connection:", err)
		return
	}
	defer db.Close()

	createTable()

	if len(os.Args) < 2 {
		fmt.Println("Usage: gogo <command>")
		fmt.Println("Commands:")
		fmt.Println("  add      Add a new task")
		fmt.Println("  update   Update an existing task")
		fmt.Println("  delete   Delete a task")
		fmt.Println("  list     List all tasks")
		return
	}

	switch os.Args[1] {
	case "add":
		addTask()
	case "update":
		updateTask()
	case "delete":
		deleteTask()
	case "list":
		listTasks()
	default:
		fmt.Println("Invalid command. Use 'gogo' with one of the following commands: add, update, delete, list")
	}
}

func createTable() {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		priority TEXT,
		category TEXT,
		completed BOOLEAN
	);`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		fmt.Println("Error creating table:", err)
		os.Exit(1)
	}
}

func addTask() {
	var task Task
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter task title: ")
	task.Title, _ = reader.ReadString('\n')
	task.Title = strings.TrimSpace(task.Title)

	fmt.Print("Enter task priority (low/medium/high): ")
	fmt.Scanln(&task.Priority)

	fmt.Print("Enter task category: ")
	task.Category, _ = reader.ReadString('\n')
	task.Category = strings.TrimSpace(task.Category)

	task.Completed = false

	insertQuery := `
	INSERT INTO tasks (title, priority, category, completed) 
	VALUES (?, ?, ?, ?);`

	_, err := db.Exec(insertQuery, task.Title, task.Priority, task.Category, task.Completed)
	if err != nil {
		fmt.Println("Error adding task:", err)
		return
	}

	fmt.Println("Task added successfully")
}

func updateTask() {
	var taskID int
	fmt.Print("Enter task ID to update: ")
	fmt.Scanln(&taskID)

	var task Task
	row := db.QueryRow("SELECT id, title, priority, category, completed FROM tasks WHERE id = ?", taskID)
	err := row.Scan(&task.ID, &task.Title, &task.Priority, &task.Category, &task.Completed)
	if err != nil {
		fmt.Println("Error fetching task details:", err)
		return
	}

	fmt.Print("Enter updated task title (press Enter to skip): ")
	var updatedTitle string
	fmt.Scanln(&updatedTitle)
	if updatedTitle != "" {
		task.Title = updatedTitle
	}

	fmt.Print("Enter updated task priority (press Enter to skip): ")
	var updatedPriority string
	fmt.Scanln(&updatedPriority)
	if updatedPriority != "" {
		task.Priority = updatedPriority
	}

	fmt.Print("Enter updated task category (press Enter to skip): ")
	var updatedCategory string
	fmt.Scanln(&updatedCategory)
	if updatedCategory != "" {
		task.Category = updatedCategory
	}

	updateQuery := `
	UPDATE tasks 
	SET title=?, priority=?, category=? 
	WHERE id=?;`

	_, err = db.Exec(updateQuery, task.Title, task.Priority, task.Category, task.ID)
	if err != nil {
		fmt.Println("Error updating task:", err)
		return
	}

	fmt.Println("Task updated successfully")
}

func deleteTask() {
	var taskID int
	fmt.Print("Enter task ID to delete: ")
	fmt.Scanln(&taskID)

	deleteQuery := "DELETE FROM tasks WHERE id=?;"

	_, err := db.Exec(deleteQuery, taskID)
	if err != nil {
		fmt.Println("Error deleting task:", err)
		return
	}

	fmt.Println("Task deleted successfully")
}

func listTasks() {
	rows, err := db.Query("SELECT id, title, priority, category, completed FROM tasks;")
	if err != nil {
		fmt.Println("Error fetching tasks:", err)
		return
	}
	defer rows.Close()

	fmt.Println("\nAll Tasks:")
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Priority, &task.Category, &task.Completed)
		if err != nil {
			fmt.Println("Error scanning task:", err)
			continue
		}
		fmt.Printf("ID: %d, Title: %s, Priority: %s, Category: %s, Completed: %t\n", task.ID, task.Title, task.Priority, task.Category, task.Completed)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
	}
}
