Here's a **complete course/tutorial on JSON handling in Go**, organized in a modular format with **lessons, explanations, and runnable code examples**. It uses Go’s built-in `encoding/json` package and covers **all major use cases**: marshalling, unmarshalling, custom struct handling, stream encoding, decoding, and working with interfaces.

---

# 📘 Course: Mastering JSON in Go (Golang)

> Learn how to encode, decode, and manipulate JSON using Go’s powerful `encoding/json` standard library.

---

## 🧰 Prerequisites

```go
import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)
```

---

## 📦 Module 1: Basics of JSON

### 🔹 Lesson 1.1: What is JSON?

* A lightweight data-interchange format (JavaScript Object Notation)
* Native support in Go via `encoding/json`

---

## 🧱 Module 2: Marshal (Convert Go to JSON)

### 🔹 Lesson 2.1: Marshal a Struct

```go
type User struct {
	Name string
	Age  int
}
u := User{"Alice", 30}
jsonData, _ := json.Marshal(u)
fmt.Println(string(jsonData)) // {"Name":"Alice","Age":30}
```

### 🔹 Lesson 2.2: Pretty Printing

```go
jsonPretty, _ := json.MarshalIndent(u, "", "  ")
fmt.Println(string(jsonPretty))
```

### 🔹 Lesson 2.3: Marshal Maps and Slices

```go
data := map[string]interface{}{
	"name": "Go",
	"type": "Language",
}
jsonBytes, _ := json.Marshal(data)
fmt.Println(string(jsonBytes))
```

---

## 📥 Module 3: Unmarshal (Convert JSON to Go)

### 🔹 Lesson 3.1: Unmarshal into Struct

```go
jsonStr := `{"Name":"Bob","Age":25}`
var user User
json.Unmarshal([]byte(jsonStr), &user)
fmt.Println(user.Name, user.Age)
```

### 🔹 Lesson 3.2: Unmarshal into Map

```go
var data map[string]interface{}
json.Unmarshal([]byte(`{"lang":"Go","version":1.21}`), &data)
fmt.Println(data["lang"], data["version"])
```

---

## 🏷️ Module 4: Struct Tags & Omitempty

### 🔹 Lesson 4.1: Use JSON Tags

```go
type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price,omitempty"`
}
p := Product{Name: "Pen"}
b, _ := json.Marshal(p)
fmt.Println(string(b)) // {"name":"Pen"}
```

### 🔹 Lesson 4.2: Ignore Fields

```go
type Secret struct {
	Token string `json:"-"`
}
```

---

## 🎯 Module 5: Custom Field Types

### 🔹 Lesson 5.1: Use `json.RawMessage`

```go
type Event struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}
```

### 🔹 Lesson 5.2: Unmarshal Interface with Type Switch

```go
var result interface{}
json.Unmarshal([]byte(`{"name":"Go"}`), &result)

m := result.(map[string]interface{})
for k, v := range m {
	fmt.Println(k, v)
}
```

---

## 🔄 Module 6: Decode/Encode JSON Stream

### 🔹 Lesson 6.1: Stream Decode from Reader

```go
jsonStream := `{"name":"Go"}`
dec := json.NewDecoder(strings.NewReader(jsonStream))
var data map[string]interface{}
dec.Decode(&data)
```

### 🔹 Lesson 6.2: Stream Encode to Writer

```go
enc := json.NewEncoder(os.Stdout)
enc.SetIndent("", "  ")
enc.Encode(map[string]string{"lang": "Go"})
```

---

## 📚 Module 7: Handling Errors

### 🔹 Lesson 7.1: Check Marshal Errors

```go
invalid := make(chan int) // not serializable
_, err := json.Marshal(invalid)
if err != nil {
	fmt.Println("Error:", err)
}
```

### 🔹 Lesson 7.2: Handle Unknown Fields

```go
dec := json.NewDecoder(strings.NewReader(`{"name":"X", "extra":"ignored"}`))
dec.DisallowUnknownFields()
var u User
err := dec.Decode(&u)
```

---

## 🧩 Module 8: Nested & Complex JSON

### 🔹 Lesson 8.1: Nested Structs

```go
type Address struct {
	City string `json:"city"`
}
type Person struct {
	Name    string  `json:"name"`
	Address Address `json:"address"`
}
```

### 🔹 Lesson 8.2: List of Objects

```go
jsonData := `[{"Name":"A"},{"Name":"B"}]`
var users []User
json.Unmarshal([]byte(jsonData), &users)
```

---

## 🚀 Module 9: Advanced Usage

### 🔹 Lesson 9.1: Marshal JSON with Custom Time Format

```go
type Log struct {
	Timestamp string `json:"timestamp"`
}
```

### 🔹 Lesson 9.2: Use Aliases to Avoid Recursion

```go
type UserAlias User
```

---

## 🧪 Module 10: Real-world Exercises

1. ✅ Convert a Go struct into a pretty-printed JSON and save to a file.
2. ✅ Read JSON config from a file and parse it.
3. ✅ Create a custom struct with nested slices and maps, and marshal it.
4. ✅ Decode a JSON array into a slice of struct.
5. ✅ Use `json.RawMessage` for partially decoding dynamic fields.

---

## ✅ Best Practices

| Practice    | Tip                                            |
| ----------- | ---------------------------------------------- |
| Struct Tags | Always use `json:"field_name"` for public APIs |
| Omitempty   | Avoid nulls in output                          |
| RawMessage  | Use for selective decoding                     |
| Decoder     | Use stream decoders for large data             |
| Interface{} | Use carefully — adds type safety risk          |

---

## 📘 Cheatsheet

| Task              | Function                             |
| ----------------- | ------------------------------------ |
| Encode Go to JSON | `json.Marshal(v)`                    |
| Pretty Encode     | `json.MarshalIndent(v, "", "  ")`    |
| Decode JSON to Go | `json.Unmarshal(data, &v)`           |
| Decode stream     | `json.NewDecoder(reader).Decode(&v)` |
| Encode stream     | `json.NewEncoder(writer).Encode(v)`  |
| Tags              | `json:"name,omitempty"`              |
| Raw JSON          | `json.RawMessage`                    |

---