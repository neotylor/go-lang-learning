Absolutely! Here's a **comprehensive tutorial/course on handling URLs in Go**, structured in a **course format** with **lessons, examples, and use-cases** using the `net/url` package.

---

# 📘 Course: Mastering URL Handling in Go (Golang)

> Learn how to parse, build, encode, and manipulate URLs using Go’s standard `net/url` package.

---

## 📦 Module 1: Getting Started

### 🔹 Lesson 1.1: Import Required Package

```go
import (
	"fmt"
	"log"
	"net/url"
)
```

---

## 🔍 Module 2: Parsing URLs

### 🔹 Lesson 2.1: Basic URL Parsing

```go
u, err := url.Parse("https://example.com/path?query=123#fragment")
if err != nil {
	log.Fatal(err)
}
fmt.Println("Scheme:", u.Scheme)
fmt.Println("Host:", u.Host)
fmt.Println("Path:", u.Path)
fmt.Println("RawQuery:", u.RawQuery)
fmt.Println("Fragment:", u.Fragment)
```

### 🔹 Lesson 2.2: Parse Relative URL Against Base

```go
base, _ := url.Parse("https://example.com/base/")
ref, _ := url.Parse("../other")
resolved := base.ResolveReference(ref)
fmt.Println("Resolved URL:", resolved.String())
```

---

## 🔧 Module 3: Building URLs

### 🔹 Lesson 3.1: Create a URL Struct

```go
u := &url.URL{
	Scheme:   "https",
	Host:     "api.example.com",
	Path:     "v1/data",
	RawQuery: "page=1&limit=10",
}
fmt.Println(u.String())
```

### 🔹 Lesson 3.2: Set Query Parameters Programmatically

```go
u := &url.URL{
	Scheme: "https",
	Host:   "example.com",
	Path:   "search",
}
q := u.Query()
q.Set("q", "golang")
q.Set("sort", "desc")
u.RawQuery = q.Encode()

fmt.Println("Full URL:", u.String())
```

---

## 🧵 Module 4: Working with Query Parameters

### 🔹 Lesson 4.1: Parse Query from URL

```go
u, _ := url.Parse("https://example.com/search?q=go&sort=desc")
params := u.Query()

fmt.Println("Query q:", params.Get("q"))
fmt.Println("Query sort:", params.Get("sort"))
```

### 🔹 Lesson 4.2: Iterate Over All Query Params

```go
for key, values := range params {
	for _, v := range values {
		fmt.Printf("%s = %s\n", key, v)
	}
}
```

### 🔹 Lesson 4.3: Add / Delete Query Parameters

```go
q := u.Query()
q.Add("page", "2")
q.Del("sort")
u.RawQuery = q.Encode()

fmt.Println("Modified URL:", u.String())
```

---

## 🔐 Module 5: Encoding & Escaping

### 🔹 Lesson 5.1: URL Encode a String

```go
s := "hello world!"
encoded := url.QueryEscape(s)
fmt.Println(encoded) // hello+world%21
```

### 🔹 Lesson 5.2: URL Decode a String

```go
decoded, _ := url.QueryUnescape("hello+world%21")
fmt.Println(decoded) // hello world!
```

---

## 👤 Module 6: Working with User Info

### 🔹 Lesson 6.1: Parse User Credentials

```go
u, _ := url.Parse("https://user:pass@example.com")
user := u.User.Username()
pass, _ := u.User.Password()

fmt.Println("Username:", user)
fmt.Println("Password:", pass)
```

### 🔹 Lesson 6.2: Set User Info

```go
u := &url.URL{
	Scheme: "https",
	Host:   "example.com",
	User:   url.UserPassword("admin", "1234"),
}
fmt.Println("URL:", u.String()) // https://admin:1234@example.com
```

---

## 🧪 Module 7: Working with Forms (Query Encoding)

### 🔹 Lesson 7.1: Encode Form Values

```go
form := url.Values{}
form.Add("name", "John Doe")
form.Add("email", "john@example.com")

fmt.Println(form.Encode()) // name=John+Doe&email=john%40example.com
```

### 🔹 Lesson 7.2: Use Encoded Form in POST Request

```go
resp, err := http.PostForm("https://httpbin.org/post", form)
```

---

## 🧠 Module 8: Common Real-Life Examples

### 🔹 Lesson 8.1: Build Search URL

```go
search := &url.URL{
	Scheme: "https",
	Host:   "www.google.com",
	Path:   "/search",
}
q := search.Query()
q.Set("q", "golang url parsing")
search.RawQuery = q.Encode()

fmt.Println(search.String())
```

### 🔹 Lesson 8.2: Append Path Segments Safely

```go
base, _ := url.Parse("https://example.com/api/")
rel := &url.URL{Path: "users/123"}
final := base.ResolveReference(rel)
fmt.Println(final.String()) // https://example.com/api/users/123
```

---

## ✅ Module 9: Best Practices

| Tip                               | Description                     |
| --------------------------------- | ------------------------------- |
| Use `url.QueryEscape`             | When building raw query strings |
| Use `url.Values`                  | For form-style query strings    |
| Always use `ResolveReference`     | For safe URL composition        |
| Sanitize user input               | Before encoding or parsing      |
| Avoid manual string concatenation | Use structured methods instead  |

---

## 🎓 Module 10: Quiz & Exercises

1. Parse a URL and extract all query parameters.
2. Build a GitHub API URL to fetch user repositories.
3. Write a function that accepts parameters and returns a full URL.
4. Decode a URL-encoded form string.
5. Implement a redirect URL resolver.

---

## 📘 Summary Cheatsheet

| Task          | Function                 |
| ------------- | ------------------------ |
| Parse URL     | `url.Parse()`            |
| Build URL     | `&url.URL{...}`          |
| Get Query     | `url.Query()`            |
| Encode Query  | `url.Values{}.Encode()`  |
| Encode String | `url.QueryEscape()`      |
| Decode String | `url.QueryUnescape()`    |
| Resolve Path  | `url.ResolveReference()` |
| Set User      | `url.UserPassword()`     |

---