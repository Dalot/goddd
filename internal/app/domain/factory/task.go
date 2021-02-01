package factory

import (
	"time"

	"github.com/Dalot/goddd/internal/app/adapter/repository"
	"github.com/Dalot/goddd/internal/app/domain"
	"github.com/brianvoe/gofakeit"
)

// Task is the factory of domain.Project
type Task struct{}

var (
	taskRepository = repository.Task{}
)

// Generate generates domain.Task from primitives
func (of Task) Generate(projectID uint) domain.Task {
	status := gofakeit.RandString([]string{domain.TaskStatusNew, domain.TaskStatusFinished})
	isFinished := false

	if status == "finished" {
		isFinished = true
	}

	task := domain.Task{
		Name:        gofakeit.BuzzWord(),
		Description: gofakeit.Sentence(10),
		Status:      status,
		ProjectID:   projectID,
	}

	if isFinished {
		task.FinishedAt = time.Now().Format("02-Jan-2006")
	}

	return task
}
