package services

import (
	"errors"

	"github.com/neotylor/go-lang-learning/tree/master/10-projects/web-api/go_task_manager_api/database"
	"github.com/neotylor/go-lang-learning/tree/master/10-projects/web-api/go_task_manager_api/models"
	"github.com/neotylor/go-lang-learning/tree/master/10-projects/web-api/go_task_manager_api/utils/validators"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(username, password string) error {
	if err := validators.ValidateUserInput(username, password); err != nil {
		return err
	}

	var existing models.User
	if err := database.DB.Where("username = ?", username).First(&existing).Error; err == nil {
		return errors.New("username already taken")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Username: username,
		Password: string(hashedPassword),
	}

	return database.DB.Create(&user).Error
}

func AuthenticateUser(username, password string) (models.User, error) {
	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return user, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return user, errors.New("invalid credentials")
	}

	return user, nil
}
