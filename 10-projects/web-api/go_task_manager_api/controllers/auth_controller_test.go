package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/neotylor/go-lang-learning/tree/master/10-projects/web-api/go_task_manager_api/database"
	"github.com/neotylor/go-lang-learning/tree/master/10-projects/web-api/go_task_manager_api/middleware"
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
	database.DB.Where("username = ?", "test_register_user").Delete(&models.User{})

	// Define request payload
	payload := models.User{
		Username: "test_register_user",
		Password: "Pass123!",
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

func TestGetProfile(t *testing.T) {
	database.Connect()
	models.InitUserModel()

	// Step 1: Insert test user
	username := "testuser_me"
	password := "StrongPass1!"
	database.DB.Where("username = ?", username).Delete(&models.User{}) // clean slate

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := models.User{
		Username: username,
		Password: string(hashedPassword),
	}
	database.DB.Create(&user)

	// Step 2: Generate token with user ID
	token, err := middleware.GenerateToken(user.Username, user.ID)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	// Step 3: Prepare request
	req, err := http.NewRequest("GET", "/me", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Handle("/me", middleware.JWTMiddleware(http.HandlerFunc(GetProfile))).Methods("GET")
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", rr.Code)
	}

	var resp map[string]interface{}
	err = json.NewDecoder(rr.Body).Decode(&resp)
	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp["username"] != username {
		t.Errorf("expected username '%s', got '%s'", username, resp["username"])
	}

	if uint(resp["userID"].(float64)) != user.ID {
		t.Errorf("expected userID %d, got %v", user.ID, resp["userID"])
	}
}
