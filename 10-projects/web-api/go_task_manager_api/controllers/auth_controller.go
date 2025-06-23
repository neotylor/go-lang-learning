package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/neotylor/go-lang-learning/tree/master/10-projects/web-api/go_task_manager_api/database"
	"github.com/neotylor/go-lang-learning/tree/master/10-projects/web-api/go_task_manager_api/middleware"
	"github.com/neotylor/go-lang-learning/tree/master/10-projects/web-api/go_task_manager_api/models"
	"github.com/neotylor/go-lang-learning/tree/master/10-projects/web-api/go_task_manager_api/services"
)

type AuthRequest struct {
	Username string `json:"username" example:"john_doe"`
	Password string `json:"password" example:"StrongPass1!"`
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
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, err := services.AuthenticateUser(req.Username, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	token, err := middleware.GenerateToken(user.Username, user.ID)
	if err != nil {
		http.Error(w, "Token generation failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

/* func Login(w http.ResponseWriter, r *http.Request) {
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
} */
/* // @Param        user  body  RegisterRequest  true  "Registration info" */

// Register godoc
// @Summary      User registration
// @Description  Registers a new user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param register body models.RegisterRequest true
// @Success      201  {object}  map[string]string
// @Failure      400  {string}  string  "Bad request"
// @Router       /register [post]
func Register(w http.ResponseWriter, r *http.Request) {
	var req models.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Username == "" || req.Password == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := services.RegisterUser(req.Username, req.Password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 🔑 Fetch the newly created user to get ID for token
	var user models.User
	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		http.Error(w, "User lookup failed after registration", http.StatusInternalServerError)
		return
	}

	// 🔐 Generate JWT token
	token, err := middleware.GenerateToken(user.Username, user.ID)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// ✅ Respond with token
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Registered and logged in successfully",
		"token":   token,
	})
}

/* func Register(w http.ResponseWriter, r *http.Request) {
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
} */

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

// GetProfile godoc
// @Summary      Get logged-in user info
// @Description  Returns the authenticated user's username and ID from the JWT token.
// @Tags         auth
// @Security     BearerAuth
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      401  {object}  map[string]string
// @Router       /me [get]
func GetProfile(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username")
	userID := r.Context().Value("userID")

	// userID := r.Context().Value("userID").(uint)
	// username := r.Context().Value("username").(string)

	if username == nil || userID == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	/* json.NewEncoder(w).Encode(map[string]string{
		"username": username.(string),
	}) */
	json.NewEncoder(w).Encode(map[string]interface{}{
		"username": username.(string),
		"userID":   userID.(uint),
	})
}
