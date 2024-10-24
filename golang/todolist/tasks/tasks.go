package tasks

import (
	"fmt"
)

type Task struct {
	ID       int
	Name     string
	Complete bool
}

// AddTask adds a new task to the database
func AddTask(name string) error {
	db, err := InitDB()
	if err != nil {
		return err
	}
	defer db.Close()

	insertTaskSQL := `INSERT INTO tasks (name, complete) VALUES (?, false)`
	_, err = db.Exec(insertTaskSQL, name)
	if err != nil {
		return fmt.Errorf("error inserting task: %v", err)
	}

	return nil
}

// ListTasks retrieves all tasks from the database
func ListTasks() ([]Task, error) {
	db, err := InitDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`SELECT id, name, complete FROM tasks`)
	if err != nil {
		return nil, fmt.Errorf("error retrieving tasks: %v", err)
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Name, &task.Complete)
		if err != nil {
			return nil, fmt.Errorf("error scanning task: %v", err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// CompleteTask marks a task as completed in the database
func CompleteTask(id int) error {
	db, err := InitDB()
	if err != nil {
		return err
	}
	defer db.Close()

	updateTaskSQL := `UPDATE tasks SET complete = true WHERE id = ?`
	_, err = db.Exec(updateTaskSQL, id)
	if err != nil {
		return fmt.Errorf("error updating task: %v", err)
	}

	return nil
}

// DeleteTask removes a task from the database
func DeleteTask(id int) error {
	db, err := InitDB()
	if err != nil {
		return err
	}
	defer db.Close()

	deleteTaskSQL := `DELETE FROM tasks WHERE id = ?`
	_, err = db.Exec(deleteTaskSQL, id)
	if err != nil {
		return fmt.Errorf("error deleting task: %v", err)
	}

	return nil
}
