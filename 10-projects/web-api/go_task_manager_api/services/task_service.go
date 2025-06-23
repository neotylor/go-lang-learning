package services

import (
	"github.com/neotylor/go-lang-learning/tree/master/10-projects/web-api/go_task_manager_api/database"
	"github.com/neotylor/go-lang-learning/tree/master/10-projects/web-api/go_task_manager_api/models"
)

func CreateTask(task *models.Task) error {
	return database.DB.Create(task).Error
}

func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := database.DB.Find(&tasks).Error
	return tasks, err
}
