Here's a **tutorial-style breakdown** of your **Go URL Shortener mini project**, complete with clear sections that you could structure into a course or workshop format. It’s ideal for beginners familiar with Go, HTTP servers, and basic hashing.

---

# 📚 URL Shortener Mini Project in Go – Tutorial Format

## 🔰 Introduction

**Project Goal:**
Build a simple URL shortener in Go that:

* Accepts a long URL via a JSON API.
* Returns a shortened version (based on MD5 hashing).
* Redirects users who visit the short URL to the original long URL.

---

## 🧱 Module 1: Project Structure & Requirements

### 🛠️ Tools & Libraries

* **Go Standard Library**

  * `net/http` for server
  * `crypto/md5` for hashing
  * `encoding/json` for JSON handling
  * `time` for timestamp

### 🗂️ Key Components

* `URL struct` to represent stored data
* `urlDB` as in-memory storage
* Handlers for shortening and redirecting URLs

---

## 📦 Module 2: Data Structure & Storage

### `URL` struct

```go
type URL struct {
	Id           string    `json:"id"`
	OriginalURL  string    `json:"original_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}
```

**Purpose:** Represents the stored data for a shortened URL.

### `urlDB` map

```go
var urlDB = make(map[string]URL)
```

**Purpose:** Acts as an in-memory database to store shortened URLs keyed by their short IDs.

---

## 🔢 Module 3: Generating Short URLs

### `generateShortURL`

```go
func generateShortURL(originalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(originalURL))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash[:8] // Use first 8 characters
}
```

**How it works:**

* Converts a URL to a hash using MD5.
* Encodes hash to hex string.
* Uses the first 8 characters for simplicity.

---

## 🏗️ Module 4: Creating and Storing URLs

### `createURL`

```go
func createURL(originalURL string) string {
	shortUrl := generateShortURL(originalURL)
	id := shortUrl

	urlDB[id] = URL{
		Id:           id,
		OriginalURL:  originalURL,
		ShortURL:     shortUrl,
		CreationDate: time.Now(),
	}
	return shortUrl
}
```

**Purpose:**

* Generate short URL from the original.
* Store it in `urlDB`.

---

## 🔍 Module 5: Fetching Original URLs

### `getURL`

```go
func getURL(shortUrl string) (URL, error) {
	url, ok := urlDB[shortUrl]
	if !ok {
		return URL{}, errors.New("URL not found")
	}
	return url, nil
}
```

**Purpose:**

* Retrieve original URL using short version.
* Handles errors if not found.

---

## 🌐 Module 6: HTTP Server Setup

### `startServer`

```go
func startServer() {
	print("Starting server on port 3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error while starting server", err)
	}
}
```

**Purpose:**

* Starts the HTTP server on port 3000.

---

## 🔗 Module 7: HTTP Handlers

### Root Handler

```go
func rootPageUrlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "/ path called, Get method requested")
}
```

**Purpose:** Just a welcome/debug endpoint.

---

### Shorten URL Handler (`/shorten`)

```go
func shortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}

	json.NewDecoder(r.Body).Decode(&data)
	shortURL := createURL(data.URL)

	response := struct {
		ShortURL string `json:"short_url"`
	}{
		ShortURL: "http://localhost:3000/redirect/" + strings.TrimSpace(shortURL),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
```

**How it works:**

* Accepts a POST request with JSON: `{ "url": "https://example.com" }`
* Responds with shortened URL.

---

### Redirect Handler (`/redirect/{id}`)

```go
func redirectURLHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]
	url, err := getURL(id)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}
```

**How it works:**

* Extracts the short ID from the URL.
* Redirects to the original URL if found.

---

## 🚀 Module 8: Putting it All Together

### `main` Function

```go
func main() {
	http.HandleFunc("/", rootPageUrlHandler)
	http.HandleFunc("/shorten", shortURLHandler)
	http.HandleFunc("/redirect/", redirectURLHandler)
	startServer()
}
```

**Purpose:** Registers routes and starts the server.

---

## ✅ Test the Project

### 1. Start the server:

```bash
go run main.go
```

### 2. Shorten a URL:

Send a `POST` request to `http://localhost:3000/shorten`
with JSON body:

```json
{
  "url": "https://example.com"
}
```

### 3. Visit the returned short URL:

Example: `http://localhost:3000/redirect/abcd1234`
It will redirect to `https://example.com`.

---

## 🧠 Bonus Ideas for Improvement

* Persist data in a file or DB (e.g. SQLite).
* Add expiration time for URLs.
* Add user authentication.
* Prevent collision in hashes.
* Use a better hash or ID generator (like UUID or nanoid).

---

## 🏁 Summary

| Feature              | Implemented |
| -------------------- | ----------- |
| Shorten URL          | ✅           |
| Redirect to long URL | ✅           |
| In-memory storage    | ✅           |
| RESTful API          | ✅           |
| Simple Hashing       | ✅           |

This project is great for getting hands-on with Go’s web server capabilities and understanding basic web service design.