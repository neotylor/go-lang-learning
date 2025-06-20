Absolutely! Here's a complete and well-organized **Golang CRUD API Tutorial**, structured as a **beginner-to-advanced course**. It uses **Gorilla Mux** for routing and **net/http**, with examples in both in-memory storage and optional database (like PostgreSQL or SQLite) if you want to scale.

---

# 📘 Course: Building a RESTful CRUD API in Go (Golang)

> Learn to build Create, Read, Update, Delete (CRUD) APIs using Go’s `net/http`, Gorilla Mux router, and JSON handling.

---

## 📦 Module 1: Project Setup

### 🔹 Lesson 1.1: Create Project Folder

```bash
mkdir go-crud-api && cd go-crud-api
go mod init go-crud-api
```

### 🔹 Lesson 1.2: Install Gorilla Mux

```bash
go get -u github.com/gorilla/mux
```

---

## 🧱 Module 2: Basic Boilerplate

```go
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	log.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
```

---

## 🧍 Module 3: Define Models

```go
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
```

---

## 🧠 Module 4: In-Memory Database

```go
var books = []Book{
	{ID: "1", Title: "Go Basics", Author: "Alice"},
	{ID: "2", Title: "Advanced Go", Author: "Bob"},
}
```

---

## 🚀 Module 5: Implement CRUD Handlers

---

### 🔹 Lesson 5.1: GET All

```go
func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}
```

### 🔹 Lesson 5.2: GET by ID

```go
func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.NotFound(w, r)
}
```

### 🔹 Lesson 5.3: POST (Create)

```go
func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	book.ID = fmt.Sprintf("%d", len(books)+1)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}
```

### 🔹 Lesson 5.4: PUT (Update)

```go
func updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range books {
		if item.ID == params["id"] {
			var updated Book
			json.NewDecoder(r.Body).Decode(&updated)
			updated.ID = item.ID
			books[i] = updated
			json.NewEncoder(w).Encode(updated)
			return
		}
	}
	http.NotFound(w, r)
}
```

### 🔹 Lesson 5.5: DELETE

```go
func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range books {
		if item.ID == params["id"] {
			books = append(books[:i], books[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}
```

---

## 📍 Module 6: Route Setup

```go
r.HandleFunc("/books", getBooks).Methods("GET")
r.HandleFunc("/books/{id}", getBook).Methods("GET")
r.HandleFunc("/books", createBook).Methods("POST")
r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
```

---

## 📃 Module 7: JSON & Headers

### 🔹 Set Content-Type Header

```go
w.Header().Set("Content-Type", "application/json")
```

### 🔹 Status Codes

```go
w.WriteHeader(http.StatusCreated) // For POST
```

---

## 🛡️ Module 8: Input Validation (Optional)

```go
if book.Title == "" {
	http.Error(w, "Title is required", http.StatusBadRequest)
	return
}
```

---

## 🗃️ Module 9: Persistent Storage (Optional DB)

* Use GORM or database/sql with PostgreSQL, MySQL, or SQLite.
* Replace the in-memory slice with DB calls.

---

## 🧪 Module 10: Testing the API

### 🔹 Tools

* Postman
* `curl`:

```bash
curl http://localhost:8080/books
```

---

## ✅ Final Directory Structure

```
go-crud-api/
├── main.go
├── handlers.go
├── models.go
├── go.mod
```

---

## 📘 Bonus Module: Advanced Features

* Middleware (CORS, Logging)
* Authentication (JWT)
* Pagination & Filtering
* Swagger Documentation

---

## 🧠 Real-world Exercises

1. Add a `PublishedYear` field to `Book` and handle it in all endpoints.
2. Add a new resource `Author` and support nested `/authors/{id}/books`.
3. Move handlers to separate files and organize packages.
4. Implement basic pagination `/books?page=1&limit=5`.
5. Use GORM with SQLite to persist books.

---

## 📝 CRUD Endpoint Summary

| Method | Endpoint      | Description     |
| ------ | ------------- | --------------- |
| GET    | `/books`      | Get all books   |
| GET    | `/books/{id}` | Get book by ID  |
| POST   | `/books`      | Create new book |
| PUT    | `/books/{id}` | Update book     |
| DELETE | `/books/{id}` | Delete book     |

---
