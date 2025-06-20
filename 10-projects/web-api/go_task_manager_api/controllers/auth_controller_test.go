package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/neotylor/go-lang-learning/tree/master/10-projects/web-api/go_task_manager_api/database"
	"github.com/neotylor/go-lang-learning/tree/master/10-projects/web-api/go_task_manager_api/models"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func TestLoginUser(t *testing.T) {
	database.Connect()
	models.InitUserModel()

	// Ensure the user exists first
	username := "loginuser"
	password := "testpass"

	// Remove user if exists (clean slate)
	database.DB.Where("username = ?", username).Delete(&models.User{})

	// Hash password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	// Create user directly in DB
	user := models.User{
		Username: username,
		Password: string(hashedPassword),
	}
	database.DB.Create(&user)

	// Prepare login request
	loginPayload := map[string]string{
		"username": username,
		"password": password,
	}
	body, _ := json.Marshal(loginPayload)

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	router := mux.NewRouter()
	router.HandleFunc("/login", Login).Methods("POST")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}

	var response map[string]string
	json.NewDecoder(rr.Body).Decode(&response)

	token, ok := response["token"]
	if !ok || token == "" {
		t.Errorf("Expected token in response, got %v", response)
	}
}

func TestRegisterUser(t *testing.T) {
	database.Connect()     // Ensure DB is connected
	models.InitUserModel() // Ensure migration has run

	router := mux.NewRouter()
	router.HandleFunc("/register", Register).Methods("POST")

	//Optional Enhancement (clean DB before test)
	database.DB.Where("username = ?", "testuser").Delete(&models.User{})

	// Define request payload
	payload := models.User{
		Username: "testuser",
		Password: "password123",
	}
	body, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, status)
	}
}
