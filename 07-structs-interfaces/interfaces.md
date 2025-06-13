# ✅ All Possible Examples for Interfaces in Go

These examples cover:

* Basics
* Implementation
* Interface composition
* Type assertions
* Empty interfaces
* Real-world use
* Interface best practices

---

## 1. 🧱 Basic Interface Declaration & Implementation

```go
type Speaker interface {
    Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
    return "Woof!"
}
```

---

## 2. 📞 Calling Interface Methods

```go
func MakeItSpeak(s Speaker) {
    fmt.Println(s.Speak())
}

func main() {
    dog := Dog{}
    MakeItSpeak(dog)
}
```

---

## 3. 🔄 Multiple Types Implementing the Same Interface

```go
type Cat struct{}

func (c Cat) Speak() string {
    return "Meow!"
}

func main() {
    animals := []Speaker{Dog{}, Cat{}}
    for _, a := range animals {
        fmt.Println(a.Speak())
    }
}
```

---

## 4. 🔗 Interface Composition

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

## 5. 🧪 Interface with Multiple Methods

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * 3.14 * c.Radius
}
```

---

## 6. 🧍‍♂️ Struct with Embedded Interface Field

```go
type Logger interface {
    Log(msg string)
}

type ConsoleLogger struct{}

func (ConsoleLogger) Log(msg string) {
    fmt.Println("LOG:", msg)
}

type App struct {
    Logger Logger
}
```

---

## 7. 🧩 Empty Interface (`interface{}`) — Accept Anything

```go
func Describe(i interface{}) {
    fmt.Printf("Value: %v, Type: %T\n", i, i)
}

Describe("hello")
Describe(42)
Describe([]int{1, 2, 3})
```

---

## 8. 📌 Type Assertion

```go
var i interface{} = "hello"

s := i.(string)           // assert type
fmt.Println(s)

s2, ok := i.(string)      // safe assertion
fmt.Println(s2, ok)
```

---

## 9. 🔄 Type Switch

```go
func IdentifyType(i interface{}) {
    switch v := i.(type) {
    case string:
        fmt.Println("It's a string:", v)
    case int:
        fmt.Println("It's an int:", v)
    default:
        fmt.Println("Unknown type")
    }
}
```

---

## 10. 🔁 Interface with Pointer Receiver Implementation

```go
type Counter struct {
    Count int
}

type Incrementer interface {
    Increment()
}

func (c *Counter) Increment() {
    c.Count++
}

func main() {
    var inc Incrementer = &Counter{}
    inc.Increment()
}
```

> Note: Pointer receivers require passing the address of the struct.

---

## 11. 🚫 Interface Nil Pitfall

```go
type Animal interface {
    Speak() string
}

func getAnimal() Animal {
    var dog *Dog = nil
    return dog // not nil interface; holds nil *Dog
}

func main() {
    a := getAnimal()
    fmt.Println(a == nil) // false!
}
```

---

## 12. 🧪 Interface Check with `_, ok`

```go
var i interface{} = 100

if v, ok := i.(int); ok {
    fmt.Println("Integer:", v)
}
```

---

## 13. 🎭 Interfaces in Polymorphism

```go
type Notifier interface {
    Notify()
}

type Email struct {
    Address string
}

func (e Email) Notify() {
    fmt.Println("Sending email to", e.Address)
}

func SendNotification(n Notifier) {
    n.Notify()
}
```

---

## 14. ⛓ Interface in Maps or Slices

```go
var items []interface{} = []interface{}{"hello", 123, true}
for _, item := range items {
    fmt.Printf("Type: %T, Value: %v\n", item, item)
}
```

---

## 15. 📦 Return Interface from Function

```go
func GetSpeaker(animal string) Speaker {
    if animal == "dog" {
        return Dog{}
    }
    return Cat{}
}
```

---

## 16. 👥 Interface Implementation with Method Overlap

```go
type A interface {
    Hello()
}

type B interface {
    Hello()
    Bye()
}

type Person struct{}

func (Person) Hello() {}
func (Person) Bye()   {}

// Person implements both A and B
```

---

## 17. ⏳ Dynamic Behavior with Interfaces

```go
type Command interface {
    Execute()
}

type Save struct{}
type Load struct{}

func (Save) Execute() { fmt.Println("Saving...") }
func (Load) Execute() { fmt.Println("Loading...") }

func RunCommand(cmd Command) {
    cmd.Execute()
}
```

---

## 18. 🧠 Using Interfaces for Testing (Mocking)

```go
type DB interface {
    Query(id int) string
}

type RealDB struct{}
func (RealDB) Query(id int) string { return "Real data" }

type MockDB struct{}
func (MockDB) Query(id int) string { return "Mock data" }

func GetData(db DB) string {
    return db.Query(1)
}
```

---

## 19. 🚪 Interface as a Field with Conditional Behavior

```go
type Shape interface {
    Draw()
}

type Circle struct{}
type Square struct{}

func (Circle) Draw() { fmt.Println("Drawing circle") }
func (Square) Draw() { fmt.Println("Drawing square") }

type Canvas struct {
    Shape Shape
}

func (c Canvas) Render() {
    c.Shape.Draw()
}
```

---

## 20. 🧰 Interface for Strategy Pattern

```go
type Strategy interface {
    Execute(a, b int) int
}

type Add struct{}
func (Add) Execute(a, b int) int { return a + b }

type Multiply struct{}
func (Multiply) Execute(a, b int) int { return a * b }

type Context struct {
    Strategy Strategy
}

func (c Context) Operate(a, b int) int {
    return c.Strategy.Execute(a, b)
}
```

---

## ✅ Summary Table

| Feature/Use Case                  | Example Provided |
| --------------------------------- | ---------------- |
| Basic Interface & Implementation  | ✅                |
| Interface Composition             | ✅                |
| Empty Interface (`interface{}`)   | ✅                |
| Type Assertion & Type Switch      | ✅                |
| Pointers & Method Sets            | ✅                |
| Polymorphism                      | ✅                |
| Interface Fields in Structs       | ✅                |
| Return/Accept Interface in Func   | ✅                |
| Real-World Use (Logger, DB, etc.) | ✅                |
| Testing with Interface Mocks      | ✅                |
| Interface Nil Trap                | ✅                |
| Strategy Pattern                  | ✅                |

---