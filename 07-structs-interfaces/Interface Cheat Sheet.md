### 🧑‍🏫 Go Interfaces – Complete Tutorial & Course

---

## 📘 Module 1: Introduction to Interfaces

### 🔹 Lesson 1.1 – What Is an Interface?

* **Definition**: A set of method signatures.
* **Purpose**: Enables polymorphism; separates behavior from implementation.
* **Syntax**:

```go
type Animal interface {
    Speak() string
}
```

---

### 🔹 Lesson 1.2 – Implementing an Interface

```go
type Dog struct{}

func (d Dog) Speak() string {
    return "Woof!"
}
```

> ✅ Go uses **implicit implementation** – no "implements" keyword required.

---

### 🔹 Lesson 1.3 – Using Interface as a Function Parameter

```go
func MakeItSpeak(a Animal) {
    fmt.Println(a.Speak())
}
```

---

## 📗 Module 2: Working with Multiple Implementations

### 🔹 Lesson 2.1 – Polymorphism via Interfaces

```go
type Cat struct{}

func (c Cat) Speak() string {
    return "Meow!"
}

func main() {
    animals := []Animal{Dog{}, Cat{}}
    for _, a := range animals {
        MakeItSpeak(a)
    }
}
```

---

### 🔹 Lesson 2.2 – Interface with Multiple Methods

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

---

## 📙 Module 3: Advanced Interface Concepts

### 🔹 Lesson 3.1 – Interface Composition

```go
type Reader interface {
    Read() string
}

type Writer interface {
    Write(s string)
}

type ReadWriter interface {
    Reader
    Writer
}
```

---

### 🔹 Lesson 3.2 – Pointer vs Value Receiver

```go
type Counter struct {
    Count int
}

func (c *Counter) Increment() {
    c.Count++
}
```

> ⚠️ If interface has pointer receiver methods, pass a pointer: `var inc Incrementer = &Counter{}`

---

### 🔹 Lesson 3.3 – Interface in Struct Field

```go
type Logger interface {
    Log(string)
}

type App struct {
    Logger Logger
}
```

---

## 📕 Module 4: Type System & Interface Mechanics

### 🔹 Lesson 4.1 – Empty Interface (`interface{}`)

```go
func Describe(i interface{}) {
    fmt.Printf("Value: %v, Type: %T\n", i, i)
}
```

---

### 🔹 Lesson 4.2 – Type Assertion

```go
var i interface{} = "hello"

s, ok := i.(string)
```

---

### 🔹 Lesson 4.3 – Type Switch

```go
switch v := i.(type) {
case string:
    fmt.Println("string:", v)
case int:
    fmt.Println("int:", v)
}
```

---

## 📒 Module 5: Real-World Use Cases

### 🔹 Lesson 5.1 – Strategy Pattern with Interface

```go
type Strategy interface {
    Execute(a, b int) int
}
```

Use `Add` and `Multiply` strategies to dynamically change behavior.

---

### 🔹 Lesson 5.2 – Logger Interface in Production Code

```go
type Logger interface {
    Log(string)
}

type ConsoleLogger struct{}

func (ConsoleLogger) Log(msg string) {
    fmt.Println("LOG:", msg)
}
```

---

### 🔹 Lesson 5.3 – Interface for Mocking in Unit Tests

```go
type DB interface {
    Query(id int) string
}
```

Use `MockDB` and `RealDB` to test `GetData()` function.

---

## 📓 Module 6: Gotchas and Best Practices

### 🔹 Lesson 6.1 – Nil Interface Trap

```go
var dog *Dog = nil
var a Animal = dog
fmt.Println(a == nil) // false!
```

### 🔹 Lesson 6.2 – Don’t Overuse Empty Interface

Use `interface{}` only when absolutely necessary. Prefer typed interfaces.

### 🔹 Lesson 6.3 – Favor Small Interfaces

```go
// Instead of:
type Big interface {
    Save()
    Load()
    Delete()
}

// Prefer:
type Saver interface { Save() }
type Loader interface { Load() }
```

---

## 🏁 Final Project Ideas

1. **Notification System**

   * Use interface for `Notifier` (Email, SMS, Push)

2. **Shape Calculator**

   * Use interface `Shape` (Circle, Rectangle, Triangle)

3. **Command Pattern**

   * Use `Command` interface (`Undo`, `Redo`, etc.)

4. **Mocking a Database Layer**

   * Interface + Fake Implementation for tests

---

## 📦 Optional Bonus Module

### 🔹 Lesson B1 – Interface in Maps and Slices

```go
items := []interface{}{"hello", 42, true}
```

### 🔹 Lesson B2 – Return Interface from Function

```go
func GetShape() Shape {
    return Circle{}
}
```

---

## 🧠 Summary Cheat Sheet

| Topic                      | Covered |
| -------------------------- | ------- |
| Basic Syntax               | ✅       |
| Multiple Implementations   | ✅       |
| Interface Composition      | ✅       |
| Empty Interface            | ✅       |
| Type Assertion/Switch      | ✅       |
| Pointer Receivers          | ✅       |
| Real-World Use Cases       | ✅       |
| Unit Testing w/ Interfaces | ✅       |
| Interface Gotchas          | ✅       |

---