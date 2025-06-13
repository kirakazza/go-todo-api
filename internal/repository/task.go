package repository

import (
	"github.com/kirakazza/go-todo-api/internal/model"
)

// Сохранить задачу
func CreateTask(t *model.Task) error {
	return DB.Create(t).Error
}

// Взять все задачи юзера
func FetchTasksByUser(userID uint) ([]model.Task, error) {
	var tasks []model.Task
	err := DB.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func UpdateTask(t *model.Task) error {
	return DB.Save(t).Error
}

func DeleteTask(taskID uint) error {
	return DB.Where("id = ?", taskID).Delete(&model.Task{}).Error
}

func GetTaskByID(taskID uint) (*model.Task, error) {
	var task model.Task
	if err := DB.First(&task, taskID).Error; err != nil {
		return nil, err
	}
	return &task, nil
}
