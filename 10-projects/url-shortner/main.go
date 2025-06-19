package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type URL struct {
	Id           string    `json:"id"`
	OriginalURL  string    `json:"original_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

var urlDB = make(map[string]URL)

func generateShortURL(OriginalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalURL)) // It converts the OriginalURL string to a byte slice
	// hasher.Write([]byte{OriginalURL}) // It converts the OriginalURL string to a byte slice
	fmt.Println("hasher", hasher)
	data := hasher.Sum(nil)
	fmt.Println("hasher data", data)
	hash := hex.EncodeToString(data)
	fmt.Println("Encode to string", hash)
	return hash[:8]
}

func createURL(originalURL string) string {
	shortUrl := generateShortURL(originalURL)
	fmt.Println("Shorted URL", shortUrl)
	id := shortUrl // use the share URL as the Id for simplicity

	urlDB[id] = URL{
		Id:           id,
		OriginalURL:  originalURL,
		ShortURL:     shortUrl,
		CreationDate: time.Now(),
	}
	return shortUrl
}

func getURL(shortUrl string) (URL, error) {
	url, ok := urlDB[shortUrl]

	if !ok {
		return URL{}, errors.New("URL not found")
	}
	return url, nil
}

func startServer() {
	// Start the HTTP server on port 8080
	print("Starting server on port 3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error while starting server", err)
		return
	}
}

func rootPageUrlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get method requested")
	fmt.Fprintf(w, "/ path called, Get method requested")
}

func shortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println("Error while decode request data", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	shortURL := createURL(data.URL)
	// fmt.Fprintf(w, shortURL)
	response := struct {
		ShortURL string `json:"short_url"`
	}{ShortURL: strings.Join([]string{"http://localhost:3000/redirect/", strings.TrimSpace(shortURL)}, "")}
	// fmt.Fprint(w, response)

	w.Header().Set("Content-Type", "application/json")
	encodeErr := json.NewEncoder(w).Encode(response)

	if encodeErr != nil {
		fmt.Println("Error on encode request")
	}
}

func redirectURLHandler(w http.ResponseWriter, r *http.Request) {
	println("called", redirectURLHandler)
	id := r.URL.Path[len("/redirect/"):]
	url, err := getURL(id)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}

func main() {
	println("Start URL Shortner project")
	// OriginalURL := "https://www.youtube.com/watch?v=dVVJU-3eU1g&list=PLzjZaW71kMwSEVpdbHPr0nPo5zdzbDulm&index=34"
	// createURL(OriginalURL)

	// Register the handler functionto handle all requests to the root URL ("/")
	http.HandleFunc("/", rootPageUrlHandler)
	http.HandleFunc("/shorten", shortURLHandler)
	http.HandleFunc("/redirect/", redirectURLHandler)
	startServer()
}
