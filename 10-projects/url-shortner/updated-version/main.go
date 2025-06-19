package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"golang.org/x/crypto/bcrypt"
	_ "modernc.org/sqlite"
)

const (
	dbPath          = "./data/shortener.db"
	urlExpiry       = 24 * time.Hour
	authUsernameEnv = "ADMIN_USER" // set via env vars
	authPasswordEnv = "ADMIN_PASS"
)

var db *sql.DB

type URL struct {
	ID          string
	OriginalURL string
	CreatedAt   time.Time
	ExpiresAt   time.Time
}

func initDB() {
	var err error
	db, err = sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	stmt := `
    CREATE TABLE IF NOT EXISTS urls (
        id TEXT PRIMARY KEY,
        original_url TEXT NOT NULL,
        created_at DATETIME NOT NULL,
        expires_at DATETIME NOT NULL
    );
    `
	_, err = db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}
}

func generateID() (string, error) {
	return gonanoid.New(8) // 8-character NanoID :contentReference[oaicite:1]{index=1}
}

func createURL(original string) (string, error) {
	id, err := generateID()
	if err != nil {
		return "", err
	}

	now := time.Now()
	expires := now.Add(urlExpiry)

	_, err = db.Exec(
		"INSERT INTO urls (id,original_url,created_at,expires_at) VALUES (?, ?, ?, ?)",
		id, original, now, expires,
	)
	if err != nil {
		if errors.Is(err, sql.ErrExecDone) {
			return createURL(original)
		}
		return "", err
	}
	return id, nil
}

func getURL(id string) (string, error) {
	var original string
	var expiresAt time.Time
	err := db.QueryRow("SELECT original_url, expires_at FROM urls WHERE id = ?", id).Scan(&original, &expiresAt)
	if err != nil {
		return "", errors.New("not found")
	}
	if time.Now().After(expiresAt) {
		return "", errors.New("expired")
	}
	return original, nil
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()
	if !ok || !validateAdmin(user, pass) {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var body struct{ URL string }
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil || body.URL == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	id, err := createURL(body.URL)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	resp := map[string]string{"short_url": fmt.Sprintf("%s/%s", r.Host, id)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	original, err := getURL(id)
	if err != nil {
		if err.Error() == "expired" {
			http.Error(w, "URL expired", http.StatusGone)
		} else {
			http.Error(w, "Not found", http.StatusNotFound)
		}
		return
	}
	http.Redirect(w, r, original, http.StatusFound)
}

func validateAdmin(user, pass string) bool {
	storedUser := os.Getenv(authUsernameEnv)
	storedHash := os.Getenv(authPasswordEnv)
	if user != storedUser {
		return false
	}
	return bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(pass)) == nil
}

func hashPassword(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	return string(hash), err
}

func main() {
	initDB()

	r := mux.NewRouter()
	r.HandleFunc("/{id}", redirectHandler).Methods("GET")
	r.HandleFunc("/shorten", shortenHandler).Methods("POST")

	addr := ":3000"
	log.Printf("Server started at %s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
