package tasks

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const taskFile = "tasks.json"

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}

func loadTasks() ([]Task, error) {
	if _, err := os.Stat(taskFile); os.IsNotExist(err) {
		return []Task{}, nil // Si no existe, empezamos con una lista vacía
	}

	file, err := ioutil.ReadFile(taskFile)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(file, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func saveTasks(tasks []Task) error {
	file, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(taskFile, file, 0644)
}

func AddTask(name string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	id := len(tasks) + 1
	tasks = append(tasks, Task{ID: id, Name: name, Complete: false})

	return saveTasks(tasks)
}

func ListTasks() ([]Task, error) {
	return loadTasks()
}

func CompleteTask(id int) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Complete = true
			break
		}
	}

	return saveTasks(tasks)
}

func DeleteTask(id int) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	filtered := tasks[:0]
	for _, task := range tasks {
		if task.ID != id {
			filtered = append(filtered, task)
		}
	}

	return saveTasks(filtered)
}
