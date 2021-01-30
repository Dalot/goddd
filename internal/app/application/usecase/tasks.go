package usecase

import (
	"errors"

	"github.com/Dalot/goddd/internal/app/domain"
	"github.com/Dalot/goddd/internal/app/domain/repository"
)

var (
	TaskErrAlreadyFinished = errors.New("Task is already finished")
)

func Tasks(taskRepository repository.ITask, projectID uint) []domain.Task {
	if projectID == 0 {
		return taskRepository.Index()
	} else {
		return taskRepository.IndexByProjectID(projectID)
	}
}

func GetTaskByID(taskRepository repository.ITask, taskID uint) (domain.Task, error) {
	task, err := taskRepository.GetByID(taskID)
	if err != nil {
		return task, err
	}
	return task, nil
}

func FinishTask(taskRepository repository.ITask, taskID uint) (domain.Task, error) {
	task, err := taskRepository.GetByID(taskID)
	if err != nil {
		return task, err
	}

	if task.Status == domain.TaskStatusFinished {
		return task, TaskErrAlreadyFinished
	}

	task = taskRepository.Finish(task)

	return task, nil
}
