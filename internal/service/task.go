package service

import (
	"fmt"

	"github.com/kirakazza/go-todo-api/internal/model"
	"github.com/kirakazza/go-todo-api/internal/repository"
)

// Получить все задачи пользователя
func GetTasks(userID uint) ([]model.Task, error) {
	return repository.FetchTasksByUser(userID)
}

// Создать задачу
func CreateTask(userID uint, title, description string) (*model.Task, error) {
	t := &model.Task{
		UserID:      userID,
		Title:       title,
		Description: description,
	}
	if err := repository.CreateTask(t); err != nil {
		return nil, err
	}
	return t, nil
}

func UpdateTask(userID, taskID uint, title, description *string, completed *bool) (*model.Task, error) {
	task, err := GetTaskByID(taskID)
	if err != nil {
		return nil, err
	}

	if task.UserID != userID {
		return nil, fmt.Errorf("unauthorized")
	}

	if title != nil {
		task.Title = *title
	}
	if description != nil {
		task.Description = *description
	}
	if completed != nil {
		task.Completed = *completed
	}

	err = repository.UpdateTask(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func GetTaskByID(id uint) (*model.Task, error) {
	var task model.Task
	if err := repository.DB.First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func DeleteTask(userID, taskID uint) error {
	task, err := repository.GetTaskByID(taskID)
	if err != nil {
		return err
	}

	if task.UserID != userID {
		return fmt.Errorf("unauthorized")
	}

	return repository.DeleteTask(taskID)
}
