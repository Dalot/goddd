package repository

import (
	"errors"

	"github.com/Dalot/goddd/internal/app/adapter/mysql"
	"github.com/Dalot/goddd/internal/app/domain"
	"gorm.io/gorm"
)

// Project is the repository of domain.Project
type Task struct{}

// Index fetches all projects
func (t Task) Index() []domain.Task {
	db := mysql.Connection()
	var tasks []domain.Task
	if err := db.Find(&tasks).Error; err != nil {
		panic(err)

	}

	return tasks
}

// IndexByProjectID fetches all tasks for that project
func (t Task) IndexByProjectID(projectID uint) []domain.Task {
	db := mysql.Connection()
	var tasks []domain.Task
	if err := db.Where("project_id = ?", projectID).Find(&tasks).Error; err != nil {
		panic(err)

	}

	return tasks
}

// GetByID fetches task by ID
func (t Task) GetByID(id uint) (domain.Task, error) {
	db := mysql.Connection()
	task := domain.Task{
		ID: id,
	}

	if err := db.First(&task).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Task{}, err
		} else {
			panic(err)
		}
	}

	return task, nil
}

// Create created task
func (t Task) Create(task domain.Task) domain.Task {
	db := mysql.Connection()

	if err := db.Create(&task).Error; err != nil {
		panic(err)

	}

	return task
}

// Save saves task
func (t Task) Save(task domain.Task) domain.Task {
	db := mysql.Connection()

	if err := db.Save(&task).Error; err != nil {
		panic(err)

	}

	return task
}

// Delete deletes task
func (t Task) Delete(taskID uint) {
	db := mysql.Connection()

	task := domain.Task{
		ID: taskID,
	}

	if err := db.Delete(&task).Error; err != nil {
		panic(err)
	}
}
