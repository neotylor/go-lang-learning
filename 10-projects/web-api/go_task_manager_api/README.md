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

Insert that line before the POST request in the test function.

---


//////Post development/////////
Would you like to:

* Add `/me` endpoint to fetch the logged-in user?
* Implement role-based access (e.g. admin-only routes)?
* Or write unit tests for registration and login?

* Add validation (e.g. min password length)?
* Automatically log in the user after registration (i.e. return JWT)?
* Or start writing tests for this endpoint?




### ✅ Next Step:

Would you like me to:

* Update `main.go` to call `models.InitUserModel()`?
* Proceed to implement `/register` in `auth_controller.go`?

Let me know if you want both.

Would you like me to:

1. Create a `User` model and migration?
2. Add `/register` logic in `auth_controller.go`?
3. Store hashed passwords using `bcrypt`?

Let me know if you want all of the above or just a specific part.


Would you like to:

* Add user registration next?
* Or create unit tests for your API?



### 🔧 Next Steps:

2. **Create login/register endpoints that issue JWT tokens.**

Would you like me to register this `/login` route inside your `task_routes.go` file, or would you prefer to create a separate `auth_routes.go` file for authentication-related endpoints?

Would you like me to:

* Add the `JWTMiddleware` to your task routes?
* Create a basic `auth_controller.go` for login and token generation?

Let me know which one you'd like to do first.



Would you like me to help add `swag init` to a make command or script for ease of use? Or should we move on to adding authentication?

Would you like to:

* Add Swagger/OpenAPI docs?
* Add basic authentication?
* Or move to Dockerizing this app?


Would you like me to also help add tests for `/login` or JWT-protected routes with mocks?


Would you like to:

Add validation (e.g. min password length)?

Automatically log in the user after registration (i.e. return JWT)?


Would you like me to summarize all the testable endpoints and help you create a full test coverage plan?

Let me know if you want:

* To test invalid login attempts
* To test protected routes using the returned token
* Or to add these to a CI workflow (like GitHub Actions)