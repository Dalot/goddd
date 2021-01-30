package repository

import (
	"errors"

	"github.com/Dalot/goddd/internal/app/adapter/mysql"
	"github.com/Dalot/goddd/internal/app/domain"
	"gorm.io/gorm"
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

// Index fetches all projects
func (p Project) IndexByUserID(userID uint) []domain.Project {
	db := mysql.Connection()
	var projects []domain.Project
	if err := db.Where("user_id = ?", userID).Find(&projects).Error; err != nil {
		panic(err)

	}

	return projects
}

// Index fetches project by ID
func (p Project) GetByID(id uint) (domain.Project, error) {
	db := mysql.Connection()
	project := domain.Project{
		ID: id,
	}

	if err := db.First(&project).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Project{}, err
		} else {
			panic(err)
		}
	}

	return project, nil
}

// Save saves project
func (p Project) Save(project domain.Project) domain.Project {
	db := mysql.Connection()

	if err := db.Create(&project).Error; err != nil {
		panic(err)

	}

	return project
}

// Delete deletes project
func (p Project) Delete(projectID uint) {
	db := mysql.Connection()

	project := domain.Project{
		ID: projectID,
	}

	if err := db.Delete(&project).Error; err != nil {
		panic(err)
	}
}
