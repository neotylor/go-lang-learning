package models

import (
	"github.com/neotylor/go-lang-learning/tree/master/10-projects/web-api/go_task_manager_api/database"

	"gorm.io/gorm"
)

type Task struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func MigrateTask(db *gorm.DB) error {
	return db.AutoMigrate(&Task{})
}

func InitTaskModel() {
	db := database.DB
	err := MigrateTask(db)
	if err != nil {
		panic("Failed to migrate task model")
	}
}
