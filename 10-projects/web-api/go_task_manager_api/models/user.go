package models

import (
	"github.com/neotylor/go-lang-learning/tree/master/10-projects/web-api/go_task_manager_api/database"

	"gorm.io/gorm"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password"`
}

/*
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}
func InitUserModel() {
	DB.AutoMigrate(&User{})
}
*/

func MigrateUser(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}

func InitUserModel() {
	db := database.DB
	err := MigrateUser(db)
	if err != nil {
		panic("Failed to migrate user model")
	}
}
