package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/neotylor/go-lang-learning/tree/master/10-projects/web-api/go_task_manager_api/database"
	"github.com/neotylor/go-lang-learning/tree/master/10-projects/web-api/go_task_manager_api/middleware"
	"github.com/neotylor/go-lang-learning/tree/master/10-projects/web-api/go_task_manager_api/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login godoc
// @Summary      User login
// @Description  Logs user in and returns a JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body  AuthRequest  true  "Login credentials"
// @Success      200  {object}  map[string]string
// @Failure      401  {string}  string  "Unauthorized"
// @Router       /login [post]
func Login(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var user models.User
	result := database.DB.Where("username = ?", req.Username).First(&user)
	if result.Error != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Compare password with bcrypt
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate JWT
	token, err := middleware.GenerateToken(user.Username)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// Register godoc
// @Summary      User registration
// @Description  Registers a new user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body  RegisterRequest  true  "Registration info"
// @Success      201  {object}  map[string]string
// @Failure      400  {string}  string  "Bad request"
// @Router       /register [post]
func Register(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	user := models.User{
		Username: req.Username,
		Password: string(hashedPassword),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		http.Error(w, "Failed to create user", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

/* // Register godoc
// @Summary Register a new user
// @Description Registers a new user with username and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /register [post]
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.Username == "" || user.Password == "" {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Save user
	if err := database.DB.Create(&user).Error; err != nil {
		http.Error(w, "User already exists or database error", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}
*/
