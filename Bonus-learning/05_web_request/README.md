Here's a **complete tutorial/course** on **Web Requests in Go**, covering everything from **basic HTTP requests** to **advanced techniques** like setting headers, timeouts, cookies, and handling JSON/XML. The course is structured into **modules**, each with explanations and runnable examples.

---

# 📘 Course: Mastering Web Requests in Go (Golang)

> Learn to send, customize, and handle web requests using Go’s standard `net/http` package and other supporting libraries.

---

## 📦 Module 1: Getting Started

### 🔹 Lesson 1.1: Import Required Packages

```go
import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)
```

---

## 🌐 Module 2: Making HTTP GET Requests

### 🔹 Lesson 2.1: Basic GET Request

```go
resp, err := http.Get("https://httpbin.org/get")
if err != nil {
	log.Fatal(err)
}
defer resp.Body.Close()

body, _ := io.ReadAll(resp.Body)
fmt.Println(string(body))
```

### 🔹 Lesson 2.2: Read Response Status Code & Headers

```go
fmt.Println("Status:", resp.Status)
fmt.Println("Headers:", resp.Header)
```

---

## 📨 Module 3: Sending POST Requests

### 🔹 Lesson 3.1: POST with `application/x-www-form-urlencoded`

```go
data := "name=Go&lang=golang"
resp, _ := http.Post("https://httpbin.org/post", "application/x-www-form-urlencoded", strings.NewReader(data))
defer resp.Body.Close()
```

### 🔹 Lesson 3.2: POST with JSON Payload

```go
jsonStr := `{"username":"gopher","email":"go@example.com"}`
resp, _ := http.Post("https://httpbin.org/post", "application/json", strings.NewReader(jsonStr))
defer resp.Body.Close()
```

---

## 🧰 Module 4: Using `http.NewRequest`

### 🔹 Lesson 4.1: Custom HTTP Request with Headers

```go
client := &http.Client{}
req, _ := http.NewRequest("GET", "https://httpbin.org/headers", nil)
req.Header.Set("User-Agent", "Go-Client/1.0")
req.Header.Set("Authorization", "Bearer token123")

resp, _ := client.Do(req)
defer resp.Body.Close()
```

---

## 🕒 Module 5: Handling Timeouts

### 🔹 Lesson 5.1: Set Timeout on Client

```go
client := &http.Client{
	Timeout: 5 * time.Second,
}
resp, err := client.Get("https://httpbin.org/delay/3")
```

---

## 🍪 Module 6: Working with Cookies

### 🔹 Lesson 6.1: Sending Cookies

```go
req, _ := http.NewRequest("GET", "https://httpbin.org/cookies", nil)
req.AddCookie(&http.Cookie{Name: "session_id", Value: "abc123"})
client := &http.Client{}
resp, _ := client.Do(req)
defer resp.Body.Close()
```

### 🔹 Lesson 6.2: Reading Cookies

```go
for _, cookie := range resp.Cookies() {
	fmt.Println("Cookie:", cookie.Name, cookie.Value)
}
```

---

## 📄 Module 7: Downloading Files

```go
resp, _ := http.Get("https://via.placeholder.com/150")
defer resp.Body.Close()

file, _ := os.Create("image.jpg")
defer file.Close()

io.Copy(file, resp.Body)
```

---

## 🔄 Module 8: PUT, PATCH, DELETE Requests

### 🔹 Lesson 8.1: PUT Request

```go
data := `{"title":"Updated"}`
req, _ := http.NewRequest("PUT", "https://httpbin.org/put", strings.NewReader(data))
req.Header.Set("Content-Type", "application/json")

client := &http.Client{}
resp, _ := client.Do(req)
defer resp.Body.Close()
```

### 🔹 Lesson 8.2: DELETE Request

```go
req, _ := http.NewRequest("DELETE", "https://httpbin.org/delete", nil)
resp, _ := client.Do(req)
defer resp.Body.Close()
```

---

## 📦 Module 9: Handling JSON Response

```go
type Response struct {
	Origin string `json:"origin"`
}

resp, _ := http.Get("https://httpbin.org/get")
defer resp.Body.Close()

var r Response
json.NewDecoder(resp.Body).Decode(&r)
fmt.Println("Origin IP:", r.Origin)
```

---

## 🔁 Module 10: Reusable HTTP Client

```go
var defaultClient = &http.Client{
	Timeout: 10 * time.Second,
}

func makeRequest(url string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-App", "GolangClient")
	return defaultClient.Do(req)
}
```

---

## 🧪 Module 11: Testing Web APIs with Local Servers

```go
// A test handler
http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
})

http.ListenAndServe(":8080", nil)
```

---

## ⚠️ Module 12: Error Handling Patterns

```go
resp, err := http.Get("http://example.com")
if err != nil {
	log.Println("Request failed:", err)
	return
}
defer resp.Body.Close()

if resp.StatusCode != http.StatusOK {
	log.Println("Unexpected status:", resp.Status)
}
```

---

## ✅ Module 13: Best Practices

| Practice                              | Description                               |
| ------------------------------------- | ----------------------------------------- |
| Always close `resp.Body`              | Use `defer resp.Body.Close()` immediately |
| Use `http.NewRequest`                 | For more control over requests            |
| Set timeouts                          | Prevent hanging requests                  |
| Use logging                           | Log errors and HTTP status                |
| Avoid global clients without timeouts | Always use a client with a timeout        |

---

## 🧠 Module 14: Advanced Topics (Optional)

* ✅ Retry logic with exponential backoff
* ✅ Proxy support via `http.Transport`
* ✅ Rate limiting using time.Ticker
* ✅ TLS client configuration
* ✅ Multipart/form-data uploads

Let me know if you want examples for any of these too.

---

## 🧪 Bonus Exercises

1. Create a **GitHub Repo Info Fetcher**.
2. Build a **JSON API Caller** using struct decoding.
3. Download a file from the internet and **save it with progress**.
4. Create a client that adds **custom headers & cookies**.
5. Send a multipart form with file upload.

---

## 📘 Summary Cheatsheet

| Action    | Function                               |
| --------- | -------------------------------------- |
| GET       | `http.Get`                             |
| POST      | `http.Post`                            |
| Custom    | `http.NewRequest` + `client.Do`        |
| Headers   | `req.Header.Set(...)`                  |
| Cookies   | `req.AddCookie(...)`, `resp.Cookies()` |
| JSON Body | `strings.NewReader`, `encoding/json`   |
| Timeout   | `http.Client{Timeout: ...}`            |

---