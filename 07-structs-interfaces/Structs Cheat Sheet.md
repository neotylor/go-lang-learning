### 🧑‍🏫 Go Structs – Complete Tutorial & Course

---

## 📘 Module 1: Introduction to Structs

### 🔹 Lesson 1.1 – What is a Struct?

* A **struct** is a user-defined type that groups related data.
* Think of it as a lightweight class without inheritance.

```go
type Person struct {
    Name string
    Age  int
}
```

---

### 🔹 Lesson 1.2 – Creating and Using Structs

```go
func main() {
    p := Person{Name: "Alice", Age: 30}
    fmt.Println(p.Name, p.Age)
}
```

---

## 📗 Module 2: Working with Structs

### 🔹 Lesson 2.1 – Zero Value Initialization

```go
var p Person
fmt.Println(p.Name) // empty string
fmt.Println(p.Age)  // 0
```

---

### 🔹 Lesson 2.2 – Named vs Positional Initialization

```go
p1 := Person{"Alice", 25}                  // positional
p2 := Person{Name: "Bob", Age: 35}         // named (recommended)
```

---

### 🔹 Lesson 2.3 – Struct Pointers

```go
p := &Person{Name: "Eve", Age: 40}
p.Age = 41 // shorthand for (*p).Age = 41
```

---

## 📙 Module 3: Methods with Structs

### 🔹 Lesson 3.1 – Value vs Pointer Receiver

```go
func (p Person) Greet() string {
    return "Hi, I’m " + p.Name
}

func (p *Person) Birthday() {
    p.Age++
}
```

---

### 🔹 Lesson 3.2 – Method Chaining with Pointers

```go
func (p *Person) SetName(name string) *Person {
    p.Name = name
    return p
}

func (p *Person) SetAge(age int) *Person {
    p.Age = age
    return p
}

func main() {
    p := &Person{}
    p.SetName("Tom").SetAge(32)
}
```

---

## 📕 Module 4: Composition & Embedding

### 🔹 Lesson 4.1 – Struct Embedding (Composition)

```go
type Address struct {
    City  string
    State string
}

type Employee struct {
    Name    string
    Address // embedded
}
```

```go
e := Employee{Name: "Jane", Address: Address{City: "NYC", State: "NY"}}
fmt.Println(e.City) // Accessed directly
```

---

### 🔹 Lesson 4.2 – Anonymous Fields

```go
type Contact struct {
    string // anonymous string field
    int    // anonymous int field
}

c := Contact{"email@example.com", 123456}
```

---

## 📒 Module 5: Tags, JSON, and Reflection

### 🔹 Lesson 5.1 – Struct Tags

```go
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```

---

### 🔹 Lesson 5.2 – Marshal/Unmarshal JSON

```go
user := User{"Alice", 25}
data, _ := json.Marshal(user)

var newUser User
json.Unmarshal(data, &newUser)
```

---

## 📓 Module 6: Advanced Struct Use Cases

### 🔹 Lesson 6.1 – Structs in Slices

```go
people := []Person{
    {Name: "A", Age: 20},
    {Name: "B", Age: 30},
}
```

---

### 🔹 Lesson 6.2 – Structs in Maps

```go
m := map[string]Person{
    "admin": {Name: "Root", Age: 100},
}
```

---

### 🔹 Lesson 6.3 – Struct Comparison

```go
p1 := Person{"Tom", 30}
p2 := Person{"Tom", 30}
fmt.Println(p1 == p2) // true (if no slice/map fields)
```

> ⚠️ Structs with slice/map/function fields are **not comparable**.

---

### 🔹 Lesson 6.4 – Structs with Interface Fields

```go
type Logger interface {
    Log(string)
}

type App struct {
    Name   string
    Logger Logger
}
```

---

### 🔹 Lesson 6.5 – Recursive Struct (Linked List)

```go
type Node struct {
    Value int
    Next  *Node
}
```

---

## 📔 Module 7: Real-World Patterns

### 🔹 Lesson 7.1 – Constructor Pattern

```go
func NewPerson(name string, age int) Person {
    return Person{Name: name, Age: age}
}
```

---

### 🔹 Lesson 7.2 – DTO/Model Layer

```go
type UserDTO struct {
    ID    int
    Name  string
    Email string
}
```

---

### 🔹 Lesson 7.3 – Config Object

```go
type Config struct {
    Port     int
    Debug    bool
    Database string
}
```

---

## 🧠 Bonus Lessons

### 🔹 Lesson B1 – Comparing Struct Sizes

```go
fmt.Println(unsafe.Sizeof(Person{}))
```

---

### 🔹 Lesson B2 – Copying Structs

```go
p1 := Person{"Tom", 30}
p2 := p1 // full copy
p2.Name = "Jerry"
```

---

## 📋 Cheat Sheet Summary

| Topic                              | Covered |
| ---------------------------------- | ------- |
| Struct Declaration                 | ✅       |
| Pointers & Receivers               | ✅       |
| Composition / Embedding            | ✅       |
| JSON Tags & Reflection             | ✅       |
| Structs in Maps/Slices             | ✅       |
| Interface Fields                   | ✅       |
| Real-World Patterns (Config, DTOs) | ✅       |
| Recursive Structs (Linked List)    | ✅       |

---

## 🏁 Final Project Ideas

### ✅ 1. Blog Post Model

```go
type Post struct {
    Title    string
    Content  string
    Author   Person
}
```

### ✅ 2. Company Hierarchy

```go
type Employee struct {
    Name     string
    Position string
    Manager  *Employee
}
```

### ✅ 3. REST API JSON Models

```go
type Product struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Price int    `json:"price"`
}
```

---