package repository

import (
	"errors"
	"time"

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

// Save saves task
func (t Task) Save(task domain.Task) domain.Task {
	db := mysql.Connection()

	if err := db.Create(&task).Error; err != nil {
		panic(err)

	}

	return task
}

// Finish finishes task
func (t Task) Finish(task domain.Task) domain.Task {
	db := mysql.Connection()
	task.FinishedAt = time.Now().Format("02 January 2006 15:04:05")
	task.Status = domain.TaskStatusFinished
	if err := db.Save(&task).Error; err != nil {
		panic(err)

	}

	return task
}
