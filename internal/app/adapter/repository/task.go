package repository

import (
	"github.com/Dalot/goddd/internal/app/adapter/mysql"
	"github.com/Dalot/goddd/internal/app/domain"
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

// Save saves project
func (t Task) Save(task domain.Task) domain.Task {
	db := mysql.Connection()

	if err := db.Create(&task).Error; err != nil {
		panic(err)

	}

	return task
}
