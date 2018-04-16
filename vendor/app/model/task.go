package model

import (
	"app/shared/database"
	"errors"
	"fmt"
	"time"
)

// Tasks table contains the information for each note
type Task struct {
	Id        uint
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Tasks gets all the tasks
func Tasks() ([]Task, error) {
	tasks := []Task{}
	err := database.SQL.Find(&tasks).Error
	return tasks, err
}

// TaskByID gets task by id
func TaskByID(id int) (Task, error) {
	task := Task{}
	err := database.SQL.First(&task, id).Error
	return task, err
}

// TaskCreate creates a task
func CreateTask(task Task) (Task, error) {
	database.SQL.NewRecord(task)
	err := database.SQL.Create(&task).Error
	return task, err
}

// TaskDelete deletes a task
func DeleteTask(id int) error {
	var err error
	if database.SQL.Where("id = ?", id).Delete(Task{}).RecordNotFound() {
		err = errors.New("Record Not Found")
		fmt.Printf("%v ***", err)
	}
	return err

}
