package repository

import (
	"github.com/Dalot/goddd/internal/app/adapter/mysql"
	"github.com/Dalot/goddd/internal/app/domain"
)

// Project is the repository of domain.Project
type Project struct{}

// Index fetches all projects
func (p Project) Index() []domain.Project {
	db := mysql.Connection()
	var projects []domain.Project
	if err := db.Find(&projects).Error; err != nil {
		panic(err)

	}

	return projects
}

// Save saves project
func (p Project) Save(project domain.Project) domain.Project {
	db := mysql.Connection()

	if err := db.Save(project).Error; err != nil {
		panic(err)

	}

	return project
}
