package services

import (
	"testing"

	"github.com/neotylor/go-lang-learning/tree/master/10-projects/web-api/go_task_manager_api/database"
	"github.com/neotylor/go-lang-learning/tree/master/10-projects/web-api/go_task_manager_api/models"

	"golang.org/x/crypto/bcrypt"
)

func TestRegisterUser(t *testing.T) {
	database.Connect()
	models.InitUserModel()

	username := "test_register_user"
	password := "test123"

	// Clean up if user already exists
	database.DB.Where("username = ?", username).Delete(&models.User{})

	err := RegisterUser(username, password)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	// Confirm user was created
	var user models.User
	result := database.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		t.Fatalf("user was not created: %v", result.Error)
	}
}

func TestAuthenticateUser(t *testing.T) {
	database.Connect()
	models.InitUserModel()

	username := "test_login_user"
	password := "testpass123"

	// Clean slate
	database.DB.Where("username = ?", username).Delete(&models.User{})

	// Insert user manually
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := models.User{
		Username: username,
		Password: string(hashedPassword),
	}
	database.DB.Create(&user)

	// Test authentication
	_, err := AuthenticateUser(username, password)
	if err != nil {
		t.Fatalf("expected valid credentials, got error: %v", err)
	}

	// Test invalid credentials
	_, err = AuthenticateUser(username, "wrongpass")
	if err == nil {
		t.Fatal("expected error for wrong password, got nil")
	}
}

func TestRegisterUserWithShortPassword(t *testing.T) {
	err := RegisterUser("validuser", "123") // too short
	if err == nil || err.Error() != "password must be at least 6 characters" {
		t.Errorf("expected password validation error, got: %v", err)
	}
}
