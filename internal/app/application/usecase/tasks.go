package usecase

import (
	"errors"
	"time"

	"github.com/Dalot/goddd/internal/app/domain"
	"github.com/Dalot/goddd/internal/app/domain/repository"
)

var (
	TaskErrAlreadyFinished = errors.New("Task is already finished")
	TaskErrCannotBeDeleted = errors.New("Task cannot be deleted because is already finished")
	TaskErrCannotBeUpdated = errors.New("Task cannot be edited because is already finished")
)

type CreateTaskArgs struct {
	ProjectID         uint
	Name              string
	Description       string
	ProjectRepository repository.IProject
	TaskRepository    repository.ITask
}

type UpdateTaskArgs struct {
	TaskID         uint
	Name           string
	Description    string
	TaskRepository repository.ITask
}

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

func CreateTask(args CreateTaskArgs) (domain.Task, error) {
	_, err := args.ProjectRepository.GetByID(args.ProjectID)
	if err != nil {
		return domain.Task{}, err
	}

	task := domain.Task{
		Name:        args.Name,
		Description: args.Description,
		ProjectID:   args.ProjectID,
		Status:      domain.TaskStatusNew,
	}

	task = args.TaskRepository.Create(task)
	return task, nil
}

func UpdateTask(args UpdateTaskArgs) (domain.Task, error) {
	task, err := args.TaskRepository.GetByID(args.TaskID)
	if err != nil {
		return task, err
	}

	if task.Status == domain.TaskStatusFinished {
		return task, TaskErrCannotBeUpdated
	}

	task.Name = args.Name
	task.Description = args.Description

	task = args.TaskRepository.Save(task)
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

	task.FinishedAt = time.Now().Format("02 January 2006 15:04:05")
	task.Status = domain.TaskStatusFinished

	task = taskRepository.Save(task)

	return task, nil
}

func DeleteTask(taskRepository repository.ITask, taskID uint) error {
	task, err := taskRepository.GetByID(taskID)
	if err != nil {
		return err
	}

	if task.Status == domain.TaskStatusFinished {
		return TaskErrCannotBeDeleted
	}

	taskRepository.Delete(taskID)

	return nil
}
