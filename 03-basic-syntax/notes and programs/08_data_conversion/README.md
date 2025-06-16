Here is a **comprehensive tutorial/course format** for **Data Conversion in Go (Golang)**. It covers **all possible types of data conversion**, grouped by category and organized in a way that's easy to follow step-by-step for beginners and intermediate learners.

---

## 🧭 Course: Data Conversion in Go (Golang)

### 📌 Prerequisites

* Basic understanding of Go syntax
* Familiarity with Go types: `int`, `float64`, `string`, `bool`, `[]byte`, etc.

---

## 🔰 Module 1: Introduction to Type Conversion in Go

### 🔹 Lesson 1.1: What is Type Conversion?

* Implicit vs Explicit conversion (Go only allows explicit)
* Syntax: `T(v)` where `T` is the target type, `v` is the value

### 🔹 Lesson 1.2: Common Built-in Types

* Numeric types: `int`, `int64`, `float64`, `uint`
* Strings and `[]byte`
* `rune`, `bool`

---

## 🔢 Module 2: Numeric Conversions

### 🔹 Lesson 2.1: Between Integer Types

```go
var i int = 42
var i64 int64 = int64(i)
var u uint = uint(i)
```

### 🔹 Lesson 2.2: Integer to Float and Vice Versa

```go
var i int = 10
var f float64 = float64(i)
var backToInt int = int(f)
```

### 🔹 Lesson 2.3: Float Precision Loss

```go
var f float64 = 12345.6789
var i int = int(f) // Truncates decimal part
```

---

## 🔤 Module 3: String Conversions

### 🔹 Lesson 3.1: Numeric to String

```go
import "strconv"

i := 123
s := strconv.Itoa(i) // int to string

f := 12.34
s2 := strconv.FormatFloat(f, 'f', 2, 64)
```

### 🔹 Lesson 3.2: String to Numeric

```go
s := "123"
i, _ := strconv.Atoi(s)

f, _ := strconv.ParseFloat("12.34", 64)
```

### 🔹 Lesson 3.3: String to Bool and Vice Versa

```go
s := "true"
b, _ := strconv.ParseBool(s)

s2 := strconv.FormatBool(b)
```

---

## 🔠 Module 4: Byte and Rune Conversions

### 🔹 Lesson 4.1: String to \[]byte and Vice Versa

```go
s := "hello"
b := []byte(s)
s2 := string(b)
```

### 🔹 Lesson 4.2: String to \[]rune and Vice Versa

```go
s := "नमस्ते"
r := []rune(s)
s2 := string(r)
```

---

## 🧩 Module 5: JSON Encoding/Decoding

### 🔹 Lesson 5.1: Struct to JSON and Back

```go
import (
	"encoding/json"
)

type Person struct {
	Name string
	Age  int
}

p := Person{"John", 30}
jsonBytes, _ := json.Marshal(p)
jsonStr := string(jsonBytes)

var p2 Person
json.Unmarshal(jsonBytes, &p2)
```

### 🔹 Lesson 5.2: JSON to Map and Vice Versa

```go
var m map[string]interface{}
json.Unmarshal([]byte(`{"name": "Jane", "age": 25}`), &m)

b, _ := json.Marshal(m)
```

---

## 💼 Module 6: Custom Conversions

### 🔹 Lesson 6.1: Implement `Stringer` Interface

```go
type Employee struct {
	Name string
	ID   int
}

func (e Employee) String() string {
	return fmt.Sprintf("%s (%d)", e.Name, e.ID)
}
```

### 🔹 Lesson 6.2: Custom `MarshalJSON`

```go
type User struct {
	Name string
	Age  int
}

func (u User) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name string `json:"username"`
		Age  int    `json:"user_age"`
	}{u.Name, u.Age})
}
```

---

## 🧪 Module 7: Advanced Conversions

### 🔹 Lesson 7.1: Interface to Concrete Type (Type Assertion)

```go
var x interface{} = 10
i := x.(int)
```

### 🔹 Lesson 7.2: Reflection-based Conversion

```go
import "reflect"

func convert(val interface{}) {
	t := reflect.TypeOf(val)
	v := reflect.ValueOf(val)
	fmt.Println("Type:", t, "Value:", v)
}
```

---

## 🧵 Module 8: Common Pitfalls

* Integer overflows
* Precision loss in float to int
* Panic on invalid type assertion
* String conversion errors (check `err`)
* Incorrect JSON field tags

---

## 🧑‍💻 Module 9: Practice Projects

1. **Currency Converter CLI** – float64 ↔ string, JSON config
2. **CSV Importer** – parse strings to int/float/bool
3. **Custom Logger** – custom type implementing `Stringer`
4. **Data Transformer API** – JSON ↔ struct/map ↔ DB model

---

## ✅ Summary & Checklist

* [x] Convert between basic numeric types
* [x] Convert strings to/from numbers, booleans
* [x] Work with `[]byte`, `[]rune`, and JSON
* [x] Handle custom conversions using interfaces
* [x] Use `strconv`, `encoding/json`, `reflect`, `fmt`

---