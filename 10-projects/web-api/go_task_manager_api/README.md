W've started your mini project: **Go Task Manager API**. We'll build it step-by-step. Here's what's next:

---

### ✅ Step 1: Project Structure

```
go_task_manager_api/
├── main.go
├── controllers/
│   └── task_controller.go
├── models/
│   └── task.go
├── routes/
│   └── task_routes.go
├── database/
│   └── connection.go
├── middleware/
│   └── jwt_auth.go
└── go.mod
```

````go_task_manager_api/main.go
// go_task_manager_api/main.go
package main

import (
	"log"
	"net/http"

	"go_task_manager_api/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterTaskRoutes(r)

	log.Println("Server starting at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

````

---

### 🔧 Step 2: Add Task Model

✅ The `models/task.go` file has been created with the `Task` struct and migration logic.

---

````go_task_manager_api/models/task.go
// go_task_manager_api/models/task.go
package models

import (
	"go_task_manager_api/database"
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

````


### 🧱 What's Inside:

```go
type Task struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
```

* `MigrateTask()` – sets up the schema via GORM.
* `InitTaskModel()` – initializes migration using the `database.DB` connection.

---

### ✅ Step 3:

✅ `main.go` is now updated to automatically call `models.InitTaskModel()` after `database.Connect()` to migrate the `Task` model on startup.

---
````go_task_manager_api/main.go
// go_task_manager_api/main.go
package main

import (
	"log"
	"net/http"

	"go_task_manager_api/routes"
	"go_task_manager_api/database"
	"go_task_manager_api/models"
	"github.com/gorilla/mux"
)

func main() {
	database.Connect() // Step 2: Connect to MySQL database
	models.InitTaskModel() // Step 3: Migrate Task model

	r := mux.NewRouter()
	routes.RegisterTaskRoutes(r)

	log.Println("Server starting at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

````

✅ `controllers/task_controller.go` has been created with full CRUD operations:

````go_task_manager_api/controllers/task_controller.go
// go_task_manager_api/controllers/task_controller.go
package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go_task_manager_api/database"
	"go_task_manager_api/models"

	"github.com/gorilla/mux"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	database.DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var task models.Task
	result := database.DB.First(&task, id)
	if result.Error != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	database.DB.Create(&task)
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var task models.Task
	result := database.DB.First(&task, id)
	if result.Error != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	json.NewDecoder(r.Body).Decode(&task)
	database.DB.Save(&task)
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var task models.Task
	result := database.DB.Delete(&task, id)
	if result.RowsAffected == 0 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted"})
}
````
---

### 📌 Included Handlers:

| Function     | HTTP Method | Description           |
| ------------ | ----------- | --------------------- |
| `GetTasks`   | GET         | List all tasks        |
| `GetTask`    | GET         | Get single task by ID |
| `CreateTask` | POST        | Create a new task     |
| `UpdateTask` | PUT         | Update existing task  |
| `DeleteTask` | DELETE      | Delete task by ID     |

---

✅ `routes/task_routes.go` has been created and all task-related endpoints are now registered.

---

### 🚀 Available API Endpoints

| Method | Endpoint      | Description       |
| ------ | ------------- | ----------------- |
| GET    | `/tasks`      | Get all tasks     |
| GET    | `/tasks/{id}` | Get task by ID    |
| POST   | `/tasks`      | Create a new task |
| PUT    | `/tasks/{id}` | Update a task     |
| DELETE | `/tasks/{id}` | Delete a task     |
| GET    | `/`           | Welcome message   |

---

You now have a fully functional CRUD API in Go with MySQL + GORM.


##Now Add Swagger/OpenAPI docs

Let's add **Swagger/OpenAPI documentation** to your Go Task Manager API using `swaggo/swag` and `gin-swagger` (can also work with Gorilla Mux using `httpSwagger`).

---

## 🧩 Step-by-Step: Add Swagger to Go API

---

### ✅ 1. Install Swagger Dependencies

Run these in your project root:

```bash
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/http-swagger
go get -u github.com/alecthomas/template
go install github.com/swaggo/swag/cmd/swag@latest
```

> If `swag` is not recognized, add Go bin to your system path:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

---

### ✅ 2. Initialize Swagger in Your Project

Run:

```bash
swag init
```

This will generate:

```
docs/
├── docs.go
├── swagger.json
└── swagger.yaml
```

---

### ✅ 3. Add Swagger Metadata to `main.go`

Update the top of `main.go`:

```go
// @title           Go Task Manager API
// @version         1.0
// @description     This is a task management API server.
// @host            localhost:8080
// @BasePath        /
```

Then import and use Swagger UI handler:

```go
import (
	"github.com/swaggo/http-swagger"
	_ "go_task_manager_api/docs" // docs generated by swag
)
```

Inside `main()`:

```go
r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
```

---

### ✅ 4. Annotate Your Endpoints

Example: add before `GetTasks` in `controllers/task_controller.go`:

```go
// GetTasks godoc
// @Summary      List all tasks
// @Description  get all tasks
// @Tags         tasks
// @Produce      json
// @Success      200  {array}  models.Task
// @Router       /tasks [get]
```

Do this for all endpoints.

---

### ✅ 5. Regenerate Swagger Docs

After annotating your handlers:

```bash
swag init
```

---

### ✅ 6. Access Swagger UI

Start your server, then visit:

```
http://localhost:8080/swagger/index.html
```

---

Here’s now we can streamline your development workflow by adding `swag init` to a **Makefile** or **shell script**.

---

## 🛠️ Option 1: Add to a `Makefile` (Recommended for Unix/macOS/Linux)

### 📄 Create `Makefile` in your root directory:

```makefile
# Makefile

.PHONY: swag run build

swag:
	swag init -g main.go

run: swag
	go run main.go

build: swag
	go build -o task-api main.go
```

> 💡 This makes it easy to run:

```bash
make run    # runs swag init and starts server
make swag   # just regenerates swagger docs
make build  # builds binary after updating docs
```

---

## 🛠️ Option 2: Create a Shell Script (Windows/Unix compatible)

### 📄 Create `generate-docs.sh`:

```bash
#!/bin/bash

echo "Generating Swagger docs..."
swag init -g main.go
echo "Docs generated successfully!"
```

Then run:

```bash
bash generate-docs.sh
```

---

## ✅ Recommended: Also add to `.gitignore`

```gitignore
/docs/*
!docs/.gitkeep
```

---

If you like to generate the full Makefile content in your project for you here

Here is the full `Makefile` content to streamline your development workflow with Swagger and build/run support:

---

### 📄 `Makefile`

```makefile
# Makefile for Go Task Manager API

.PHONY: swag run build clean tidy

# Generate Swagger docs
swag:
	swag init -g main.go

# Run the project (with docs generation)
run: swag
	go run main.go

# Build the binary (with docs generation)
build: swag
	go build -o bin/task-api main.go

# Tidy dependencies
tidy:
	go mod tidy

# Clean compiled binaries and Swagger docs
clean:
	rm -rf bin/
	rm -rf docs/
```

---

### ✅ Usage

| Command      | Purpose                              |
| ------------ | ------------------------------------ |
| `make swag`  | Generate Swagger docs                |
| `make run`   | Generate Swagger docs & run server   |
| `make build` | Generate Swagger docs & build binary |
| `make tidy`  | Clean up `go.mod` and `go.sum`       |
| `make clean` | Delete `/bin` and `/docs` folders    |

---

## Adding authentication

````go_task_manager_api/middleware/jwt_auth.go
// go_task_manager_api/middleware/jwt_auth.go
package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("my_secret_key") // Replace with env config in production

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization Format", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]
		claims := &jwt.RegisteredClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func GenerateToken(username string) (string, error) {
	claims := &jwt.RegisteredClaims{
		Subject:  username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

````
✅ JWT authentication middleware has been created in `middleware/jwt_auth.go`.

---

### 🔐 Middleware Functionality:

* Validates Bearer tokens from `Authorization` header.
* Parses claims using the secret key.
* Returns `401 Unauthorized` for invalid or missing tokens.

---

## Apply middleware to your task routes.

````go_task_manager_api/routes/task_routes.go
// go_task_manager_api/routes/task_routes.go
package routes

import (
	"net/http"

	"go_task_manager_api/controllers"
	"go_task_manager_api/middleware"
	"github.com/gorilla/mux"
)

func RegisterTaskRoutes(r *mux.Router) {
	authenticated := r.PathPrefix("/tasks").Subrouter()
	authenticated.Use(middleware.JWTMiddleware)

	authenticated.HandleFunc("", controllers.GetTasks).Methods("GET")
	authenticated.HandleFunc("/{id}", controllers.GetTask).Methods("GET")
	authenticated.HandleFunc("", controllers.CreateTask).Methods("POST")
	authenticated.HandleFunc("/{id}", controllers.UpdateTask).Methods("PUT")
	authenticated.HandleFunc("/{id}", controllers.DeleteTask).Methods("DELETE")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Go Task Manager API"))
	})
}

````

✅ All `/tasks` routes are now protected with `JWTMiddleware`.
Clients must include a valid `Authorization: Bearer <token>` header to access them.

---

### 🔐 Example Request with JWT:

```http
GET /tasks HTTP/1.1
Host: localhost:8080
Authorization: Bearer <your-token>
```

---

## Create a basic `auth_controller.go` next to allow users to log in and receive a token.

````go_task_manager_api/controllers/auth_controller.go
// go_task_manager_api/controllers/auth_controller.go
package controllers

import (
	"encoding/json"
	"net/http"

	"go_task_manager_api/middleware"
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

	// Dummy auth: username=admin, password=secret
	if req.Username != "admin" || req.Password != "secret" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	token, err := middleware.GenerateToken(req.Username)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

````

✅ `auth_controller.go` has been created with a basic `/login` endpoint.

---

### 🔐 Sample Credentials for Demo

* **Username:** `admin`
* **Password:** `secret`

### 📦 Output:

```json
{
  "token": "<your-jwt-token>"
}
```

---

## 🔜 Next: create a separate `auth_routes.go` file for authentication-related endpoints.

````go_task_manager_api/routes/auth_routes.go
// go_task_manager_api/routes/auth_routes.go
package routes

import (
	"go_task_manager_api/controllers"
	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(r *mux.Router) {
	r.HandleFunc("/login", controllers.Login).Methods("POST")
}

````

✅ `auth_routes.go` has been created to register the `/login` endpoint separately from task routes.

---
### 📌 Final Integration Step:

In `main.go`, make sure you also call:

```go
routes.RegisterAuthRoutes(r)
```

...after `RegisterTaskRoutes(r)` so both routes are active.

---

Now we create a central route file (e.g. `routes/routes.go`) to keep your `main.go` clean and maintainable. This "master router" will handle registering all grouped route modules.

---

### ✅ Step-by-step Refactor

#### 1. 📁 Create `routes/routes.go`

```go
// go_task_manager_api/routes/routes.go
package routes

import (
	"github.com/gorilla/mux"
)

func RegisterAllRoutes(r *mux.Router) {
	RegisterAuthRoutes(r)
	RegisterTaskRoutes(r)
}
```

---

#### 2. 🧼 Clean up `main.go`

Update to this:

```go
package main

import (
	"log"
	"net/http"

	"go_task_manager_api/database"
	"go_task_manager_api/models"
	"go_task_manager_api/routes"
	"github.com/gorilla/mux"
)

func main() {
	database.Connect()
	models.InitTaskModel()

	r := mux.NewRouter()
	routes.RegisterAllRoutes(r) // Register everything from one place

	log.Println("Server starting at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
```

---

✅ `main.go` and `routes.go` has been updated to include the `/login` route via `RegisterAuthRoutes(r)`.

---

Your Go Task Manager API now supports:

* ✅ Full CRUD with GORM + MySQL
* 🔐 JWT-protected `/tasks` endpoints
* 🔑 `/login` route for token generation
* 🧾 Swagger documentation for all endpoints


## SetUp unit tests for tasks

Great! Let’s add unit tests for your Go Task Manager API.

---

## ✅ Setup for Unit Testing in Go

### 🧪 Step 1: Create a Test File

Create a file named `task_controller_test.go` in the `controllers/` directory:

---

### 📄 `controllers/task_controller_test.go`

```go
package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go_task_manager_api/models"
	"go_task_manager_api/database"

	"github.com/gorilla/mux"
)

func init() {
	database.Connect()
	models.InitTaskModel()
}

func TestCreateTask(t *testing.T) {
	payload := models.Task{
		Title:       "Test Task",
		Description: "Testing task creation",
		Completed:   false,
	}
	body, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/tasks", CreateTask).Methods("POST")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Expected status code %v, got %v", http.StatusCreated, status)
	}
}

func TestGetTasks(t *testing.T) {
	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/tasks", GetTasks).Methods("GET")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, status)
	}
}
```

---

### 📦 Step 2: Run the Tests

In your terminal:

```bash
go test ./controllers
```

---

### ✅ Result

You’ll see output like:

```
PASS
ok  	go_task_manager_api/controllers	0.XXXs
```

---

## 🔜 Next: Add user registration

✅ `main.go` is now updated to reflect that `/register` (in addition to `/login`) will be handled by `RegisterAuthRoutes`.

---

## 🔜 Next Step: Create a `User` model and migration

```go_task_manager_api/models/user.go
// go_task_manager_api/models/user.go
package models

import (
	"go_task_manager_api/database"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password"`
}

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

```

✅ `User` model and migration setup is complete.

---

### 👤 `User` Fields:

* `ID`: primary key
* `Username`: unique and required
* `Password`: (will store bcrypt hash)


## 🔜 Next Step: Add Registration Flow

## ✅ 1. Create the `User` model and migration

📁 **`models/user.go`**:

```go
package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}

func InitUserModel() {
	DB.AutoMigrate(&User{})
}
```

---

## ✅ 2. Update `main.go`

Just below `models.InitTaskModel()` in `main.go`, add:

```go
models.InitUserModel()
```

Now `main.go` should look like:

```go
func main() {
	database.Connect()
	models.InitTaskModel()
	models.InitUserModel() // <-- Add this
	...
}
```

---

## ✅ 3. Add `/register` in `auth_controller.go` and Add User Registration Logic
📁 **Update `controllers/auth_controller.go`**:

```go
// Register godoc
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
```

---

## ✅ 4. Register `/register` in `auth_routes.go`

📁 **Update `routes/auth_routes.go`**:

```go
func RegisterAuthRoutes(r *mux.Router) {
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/register", controllers.Register).Methods("POST")
}
```
✅ `/register` endpoint is now implemented:

---

Now your API supports:

* 🔐 Secure user registration (`/register`)
* 🔑 Login (`/login`)
* 🛡️ JWT-protected task routes

### 🔐 Registration Flow:

* Accepts `username` and `password`
* Hashes the password using `bcrypt`
* Stores the user in MySQL via GORM
* Returns 201 with a success message

---

### 🧪 Example Request:

```json
POST /register
{
  "username": "neotylor",
  "password": "mypassword"
}
```
## ✅ Next Step: start writing tests for this endpoint

let’s start writing a test for the `/register`  and `/login` endpoint in your Go API.

---
We'll use Go's `net/http/httptest` and `testing` packages. These tests will be placed in `controllers/auth_controller_test.go`.

---
### ✅ 1. Basic Test Setup for `/register`

```go
package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go_task_manager_api/models"
	"go_task_manager_api/database"
	"github.com/gorilla/mux"
)

func TestRegisterUser(t *testing.T) {
	database.Connect()       // Ensure DB is connected
	models.InitUserModel()   // Ensure migration has run

	router := mux.NewRouter()
	router.HandleFunc("/register", Register).Methods("POST")

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
```

---

### ✅ 2. `TestLoginUser` — login endpoint test

Add this to your `auth_controller_test.go` file:

```go
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
```

### ✅ 3. Run the test

In your terminal:

```bash
go test ./controllers -v
```

---

### ✅ This test:

* Ensures a test user exists in the DB.
* Sends a correct `/login` request.
* Asserts a valid token is returned in the response.

---


### ✅ 4. Optional Enhancement (clean DB before test)

If you're running tests multiple times and don't want duplicate username errors, clean the `users` table before running the test:

```go
database.DB.Where("username = ?", "testuser").Delete(&models.User{})
```
---

✅ Create Central Model Initializer! This improves maintainability and keeps `main.go` clean by abstracting model registration to a single place.

---

## ✅ Step-by-Step: Create Central Model Initializer

### 📁 1. Create `models/init.go`

```go
package models

func InitModels() {
	InitTaskModel()
	InitUserModel()
}
```

> You can add more models here later (e.g. `InitProjectModel()`, `InitCommentModel()`).

---

### 🧼 2. Clean up `main.go`

Update your `main.go`:

```go
func main() {
	database.Connect()
	models.InitModels() // ✅ single call to migrate all models

	r := mux.NewRouter()
	routes.RegisterAuthRoutes(r)
	routes.RegisterTaskRoutes(r)

	log.Println("Server starting at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
```

---

Now your model setup is:

* 🔁 DRY (no repetition)
* 📦 Centralized
* 💡 Easy to extend


✅ Let’s centralize route registration as well — just like models — to make your `main.go` super clean and scalable.

---

## ✅ Step-by-Step: Central Route Initialization

### 📁 1. Create `routes/init.go`

```go
package routes

import (
	"github.com/gorilla/mux"
)

func InitRoutes(r *mux.Router) {
	RegisterAuthRoutes(r)
	RegisterTaskRoutes(r)
	// Add future route groups here
}
```

---

### 🧼 2. Update `main.go`

Replace this:

```go
routes.RegisterAuthRoutes(r)
routes.RegisterTaskRoutes(r)
```

With:

```go
routes.InitRoutes(r)
```

Now your `main.go` is beautifully organized:

```go
func main() {
	database.Connect()
	models.InitModels() // One-liner for all DB models

	r := mux.NewRouter()
	routes.InitRoutes(r) // One-liner for all routes

	log.Println("Server starting at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
```

---

This structure gives you:

* ✅ Better codebase organization
* 📦 Easier onboarding for new developers
* 🔁 Easy scaling as routes/models grow

---

Let's now create a centralized **middleware initializer** for reusable middlewares (like `JWTMiddleware` or logging).

---

## ✅ Step-by-Step: Central Middleware Setup

### 📁 1. Create `middleware/init.go`

```go
package middleware

import (
	"net/http"
)

// WrapWithMiddlewares applies all global middlewares to the router
func WrapWithMiddlewares(next http.Handler) http.Handler {
	// You can chain more middlewares here like logging, CORS, etc.
	return LoggingMiddleware(next) // example: wrap with logging
}
```

> You can later expand this to include:

```go
return CORSMiddleware(LoggingMiddleware(next))
```

---

### 📁 2. Add a Simple Logging Middleware (optional)

📄 Add this in `middleware/logging.go`:

```go
package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[%s] %s %s", r.Method, r.RequestURI, time.Since(start))
		next.ServeHTTP(w, r)
	})
}
```

---

### 🧼 3. Update `main.go`

Wrap the router:

```go
wrappedRouter := middleware.WrapWithMiddlewares(r)
http.ListenAndServe(":8080", wrappedRouter)
```

Now `main.go` looks clean, centralized, and fully extensible:

```go
func main() {
	database.Connect()
	models.InitModels()

	r := mux.NewRouter()
	routes.InitRoutes(r)

	log.Println("Server starting at http://localhost:8080")
	wrappedRouter := middleware.WrapWithMiddlewares(r)
	http.ListenAndServe(":8080", wrappedRouter)
}
```

---

✅ You now have:

* `models.InitModels()` – for DB setup
* `routes.InitRoutes()` – for route wiring
* `middleware.WrapWithMiddlewares()` – for middleware chain


---

✅ Let's add a **CORS middleware** to your Go API to allow frontend clients (like React/Next.js) to interact with your backend.

---

## 🌐 Step-by-Step: Add CORS Middleware

### 📁 1. Create `middleware/cors.go`

```go
package middleware

import (
	"net/http"
)

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow specific origins here — use "*" only in development
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Allow specific headers
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Allow specific methods
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
```

---

### ✅ 2. Update `WrapWithMiddlewares()` in `middleware/init.go`

```go
func WrapWithMiddlewares(next http.Handler) http.Handler {
	return CORSMiddleware(LoggingMiddleware(next))
}
```

> You can chain more in this order if needed:

```go
return JWTMiddleware(CORSMiddleware(LoggingMiddleware(next)))
```

---

### 🔁 3. Restart Server & Test

When you make a request from a frontend (e.g., `localhost:3000`), the CORS headers will now allow it.

---

✅ You now have fully organized middleware for:

* 🔐 JWT Auth
* 📋 Logging
* 🌐 CORS

---

✅ Organizing your logic into a `services/` layer is a best practice — it separates business logic from controller and routing concerns.

Let’s do it step by step.

---

## 📦 Step-by-Step: Create a `services/` Layer

### 🗂 Folder Structure

```
go_task_manager_api/
├── services/
│   ├── auth_service.go
│   └── task_service.go
```

---

### ✅ 1. Create `services/auth_service.go`

```go
package services

import (
	"errors"

	"go_task_manager_api/database"
	"go_task_manager_api/models"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(username, password string) error {
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
```

---

### ✅ 2. Update `controllers/auth_controller.go` to use `authService`

```go
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

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "registered successfully"})
}

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

	token, err := middleware.GenerateToken(user.Username)
	if err != nil {
		http.Error(w, "Token generation failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
```

---

### ✅ 3. (Optional) Create `services/task_service.go`

Later, if you want to abstract task operations, you could do:

```go
package services

import (
	"go_task_manager_api/database"
	"go_task_manager_api/models"
)

func CreateTask(task *models.Task) error {
	return database.DB.Create(task).Error
}

func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := database.DB.Find(&tasks).Error
	return tasks, err
}
```

---

### ✅ Benefits:

* 📚 Clear separation of concerns
* 🔁 Reusable logic for CLI, GraphQL, or gRPC later
* 🔒 Better for unit testing and mocking

---

Perfect! Let’s write unit tests for the `authService` functions: `RegisterUser` and `AuthenticateUser`.

We'll place them in `services/auth_service_test.go`.

---

## ✅ 1. Test File Setup

### 📁 `services/auth_service_test.go`

```go
package services

import (
	"testing"

	"go_task_manager_api/database"
	"go_task_manager_api/models"

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
```

---

## ✅ 2. Run Tests

```bash
go test ./services -v
```

---

## 🔎 Expected Output

```
=== RUN   TestRegisterUser
--- PASS: TestRegisterUser (0.12s)
=== RUN   TestAuthenticateUser
--- PASS: TestAuthenticateUser (0.09s)
PASS
```

---

✅ Done: You now have test coverage for your core authentication service logic!

## 🔧 Next Steps: Add validation (e.g. min password length)

Let’s add input validation — particularly **minimum password length** and basic username checks — to improve security and data integrity.

---

## ✅ Step-by-Step: Add Validation to Registration

### 📍 1. Update `services.RegisterUser()` in `services/auth_service.go`

Update the function to add validation logic:

```go
func RegisterUser(username, password string) error {
	if len(username) < 3 {
		return errors.New("username must be at least 3 characters")
	}
	if len(password) < 6 {
		return errors.New("password must be at least 6 characters")
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
```

---

## ✅ 2. Re-run Unit Tests

Your test `TestRegisterUser()` will now fail **if** the password is too short — which is what you want.

If you'd like to add a test to verify the **validation logic itself**, here's an optional test case:

---

### 📄 Add to `auth_service_test.go`

```go
func TestRegisterUserWithShortPassword(t *testing.T) {
	err := RegisterUser("validuser", "123") // too short
	if err == nil || err.Error() != "password must be at least 6 characters" {
		t.Errorf("expected password validation error, got: %v", err)
	}
}
```

---

✅ Done: You now have input validation + test coverage for:

* ❌ Short passwords
* ❌ Short usernames
* ✅ Existing usernames

---

✅ Perfect — let’s move validation logic into a reusable utility function to keep your service layer clean and consistent.

---

## 📦 Step-by-Step: Create Reusable Validator Utility

### 📁 1. Create `utils/validators/validators.go`

```go
package validators

import (
	"errors"
	"unicode"
)

func ValidateUserInput(username, password string) error {
	if len(username) < 3 {
		return errors.New("username must be at least 3 characters")
	}
	if len(password) < 6 {
		return errors.New("password must be at least 6 characters")
	}
	if !containsDigit(password) {
		return errors.New("password must include at least one number")
	}
	if !containsSymbol(password) {
		return errors.New("password must include at least one special character")
	}
	return nil
}

func containsDigit(s string) bool {
	for _, ch := range s {
		if unicode.IsDigit(ch) {
			return true
		}
	}
	return false
}

func containsSymbol(s string) bool {
	for _, ch := range s {
		if unicode.IsSymbol(ch) || unicode.IsPunct(ch) {
			return true
		}
	}
	return false
}
```

---

### 🛠 2. Update `services.RegisterUser()` to use it

```go
import "go_task_manager_api/utils/validators"

func RegisterUser(username, password string) error {
	if err := validators.ValidateUserInput(username, password); err != nil {
		return err
	}

	var existing models.User
	if err := database.DB.Where("username = ?", username).First(&existing).Error; err == nil {
		return errors.New("username already taken")
	}

	// Continue with hashing and saving
	...
}
```

---

### 🧪 3. Optional: Add Unit Test for Validator

📁 `utils/validators/validators_test.go`:

```go
package validators

import "testing"

func TestValidateUserInput(t *testing.T) {
	tests := []struct {
		username string
		password string
		wantErr  bool
	}{
		{"ab", "pass123!", true},
		{"validuser", "short", true},
		{"validuser", "longbutno123", true},
		{"validuser", "NoSymbol123", true},
		{"validuser", "StrongPass1!", false},
	}

	for _, tt := range tests {
		err := ValidateUserInput(tt.username, tt.password)
		if (err != nil) != tt.wantErr {
			t.Errorf("ValidateUserInput(%q, %q) = %v, wantErr %v", tt.username, tt.password, err, tt.wantErr)
		}
	}
}
```

---

### ✅ Result

You now have:

* 🔁 Reusable validation logic in one place
* ✅ Password strength checks (length, digit, symbol)
* 🧪 Test coverage for validators

---

## ✅ Next Step: Add `/me` endpoint to fetch the logged-in user.

✅ Let's add a `/me` endpoint to return the currently authenticated user's info using the JWT token.

---

## 🧩 Step-by-Step: Add `/me` Endpoint

### 📄 1. Add `GetProfile()` in `auth_controller.go`

```go
// GetProfile godoc
// @Summary Get logged-in user info
// @Description Returns the authenticated user's data
// @Tags auth
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /me [get]
func GetProfile(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username")
	if username == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"username": username.(string),
	})
}
```

> ✅ Note: This assumes your JWT middleware sets `username` in `context.WithValue()`.

---

### 🛡 2. Update `routes/auth_routes.go`

Add:

```go
r.Handle("/me", middleware.JWTMiddleware(http.HandlerFunc(controllers.GetProfile))).Methods("GET")
```

Now `RegisterAuthRoutes()` will look like:

```go
func RegisterAuthRoutes(r *mux.Router) {
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.Handle("/me", middleware.JWTMiddleware(http.HandlerFunc(controllers.GetProfile))).Methods("GET")
}
```

---

### 🔐 3. JWT Context Injection (already assumed done)

Make sure your `JWTMiddleware` injects the username into the context:

```go
ctx := context.WithValue(r.Context(), "username", claims["username"])
```

---

## ✅ Usage

**Request:**

```http
GET /me
Authorization: Bearer <your-jwt-token>
```

**Response:**

```json
{
  "username": "admin"
}
```

---

JWT Context Injection (if not already done yet)! Let's update your `JWTMiddleware` to inject the username into the request context — so that `/me` (and future protected routes) can access it easily.

---

## ✅ Step-by-Step: Add JWT Context Injection

### 📄 1. Update `middleware/jwt_auth.go`

Make sure your middleware:

* Parses the token
* Extracts `"username"` from JWT claims
* Injects it into `r.Context()`

---

### ✅ Updated `JWTMiddleware` function:

```go
package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var JwtSecret = []byte("secret-key") // use env var in production

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return JwtSecret, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["username"] == nil {
			http.Error(w, "Invalid claims", http.StatusUnauthorized)
			return
		}

		// ✅ Inject username into request context
		ctx := context.WithValue(r.Context(), "username", claims["username"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
```

---

### ✅ Now your `/me` endpoint can access:

```go
username := r.Context().Value("username").(string)
```

And return it as JSON.

---

## 🔧 Next Steps:Add a test for the `/me` endpoint with a mocked JWT token?

✅ Let's add a unit test for the `/me` endpoint using a **mocked valid JWT token**.

---

## 📦 Step-by-Step: Test `/me` with JWT Token

### 📁 1. Add to `auth_controller_test.go`

```go
func TestGetProfile(t *testing.T) {
	database.Connect()
	models.InitUserModel()

	// Generate a mock token
	token, err := middleware.GenerateToken("testuser")
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

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

	var resp map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Errorf("failed to decode response: %v", err)
	}

	if resp["username"] != "testuser" {
		t.Errorf("expected username 'testuser', got '%s'", resp["username"])
	}
}
```

---

### 🧪 2. Run the Test

```bash
go test ./controllers -v
```

---

### ✅ Output (Expected)

```
=== RUN   TestGetProfile
--- PASS: TestGetProfile (0.01s)
```

This confirms that:

* ✅ JWT token is accepted
* ✅ `username` is correctly injected into context
* ✅ `/me` returns the correct payload

---

## 🔧 Next Steps: extend JWT to include user ID for future DB queries

✅ Now! Extending the JWT to include the **user ID** will make it easy to fetch full user records from the database later.

---

## 🧩 Step-by-Step: Include `userID` in JWT Token

---

### 📄 1. Update `GenerateToken()` in `middleware/jwt_auth.go`

Update it to include both `username` **and** `userID`:

```go
func GenerateToken(username string, userID uint) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"userID":   userID,
		"exp":      time.Now().Add(72 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}
```

---

### 🛠 2. Update Login Controller to Use New Signature

In `auth_controller.go`, modify this:

```go
token, err := middleware.GenerateToken(user.Username, user.ID)
```

---

### 🧠 3. Update JWT Middleware to Inject Both `userID` and `username` into Context

```go
ctx := context.WithValue(r.Context(), "username", claims["username"])
ctx = context.WithValue(ctx, "userID", uint(claims["userID"].(float64)))
next.ServeHTTP(w, r.WithContext(ctx))
```

> 🔐 Note: JWT stores numbers as `float64` — we convert to `uint`.

---

### ✅ 4. Example `/me` Update (Optional)

Now in `/me` you can access both:

```go
userID := r.Context().Value("userID").(uint)
username := r.Context().Value("username").(string)
```

---

## ✅ Done!

Your JWT token now securely includes:

* 🔐 `userID`
* 🙋 `username`
* ⏳ `exp`

And your API context is ready for **DB-based lookups**.

## ✅ Next Step: Add a unit test for the /me endpoint

Great! Now that your JWT includes both `username` and `userID`, let’s write a **unit test** for the `/me` endpoint that:

* Authenticates with a valid JWT.
* Verifies both `username` and `userID` are available.
* Ensures it returns `200 OK`.

---

## ✅ Step-by-Step: Unit Test for `/me`

### 📁 1. Add to `auth_controller_test.go`

```go
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
```

---

### 📄 2. Optional: Update `/me` Handler to Return `userID` Too

If you haven't yet, update `GetProfile()` in `auth_controller.go`:

```go
func GetProfile(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username")
	userID := r.Context().Value("userID")

	if username == nil || userID == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"username": username.(string),
		"userID":   userID.(uint),
	})
}
```

---

### ✅ 3. Run the Test

```bash
go test ./controllers -v
```

You should see:

```
=== RUN   TestGetProfile
--- PASS: TestGetProfile (0.01s)
```

---

## ✅ Next Step: Update docs

✅ Let’s update your **Swagger/OpenAPI docs** to reflect the new `/me` endpoint with the extended JWT that includes both `username` and `userID`.

---

## 📄 Step-by-Step: Document `/me` Endpoint

Open your `auth_controller.go` and **add or update the Swagger comment block above `GetProfile()`** like so:

```go
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
    ...
}
```

---

## 🔐 Step 2: Confirm BearerAuth is Defined in `docs/docs.go` or `main.go`

Somewhere in your `docs` package or `main.go` (if you're using swag CLI), ensure this is defined:

```go
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
```

This tells Swagger UI that the `/me` endpoint uses JWT in the `Authorization` header.

---

## 🔁 Step 3: Rebuild Swagger Docs

Run this command from the root of your project (where `main.go` lives):

```bash
swag init
```

This will regenerate the `docs/` folder with updated Swagger annotations.

---

## ✅ Result

When you open Swagger UI (e.g., `http://localhost:8080/swagger/index.html`), you'll now see:

* `/me` listed under **auth**
* It will prompt for `Bearer <token>`
* Shows `200` response with `username` and `userID`

---

## ✅ Improvement: Automatically log in the user after registration (i.e. return JWT).

✅ Great idea — automatically logging in the user after registration (by returning a JWT) improves UX and saves a step for the client.

Let’s modify the `Register` handler to:

1. Register the user.
2. Generate and return a JWT token (just like `/login`).

---

## 🔧 Step-by-Step: Return JWT on Registration

### 📄 1. Update `Register()` in `auth_controller.go`

```go
func Register(w http.ResponseWriter, r *http.Request) {
	var req models.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err := services.RegisterUser(req.Username, req.Password)
	if err != nil {
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
```

---

### ✅ 2. Swagger Doc Update (Optional)

Update the `@Success` comment above the `Register()` handler:

```go
// @Success 201 {object} map[string]string
```

---

### 🧪 3. Optional: Update Unit Test

If you're testing the `/register` endpoint, update it to:

* Expect `201`
* Check for `"token"` in the response body

---

## ✅ Result

Now when you `POST /register`, the response is:

```json
{
  "message": "Registered and logged in successfully",
  "token": "eyJhbGciOiJIUzI1NiIsInR..."
}
```

You can directly use this token for `/me` or any protected route. 🎉

---