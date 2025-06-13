# ✅ All Possible Examples for Structs in Go

---

## 1. 📦 Basic Struct Definition and Initialization

```go
type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{"Alice", 25}
    fmt.Println(p.Name, p.Age)
}
```

---

## 2. 📝 Named Field Initialization

```go
p := Person{
    Name: "Bob",
    Age:  30,
}
```

---

## 3. 🧑‍🎤 Access and Modify Struct Fields

```go
p := Person{"John", 40}
p.Age = 41
fmt.Println(p.Age) // Output: 41
```

---

## 4. 🧾 Zero Value Struct

```go
var p Person
fmt.Println(p) // Name: "", Age: 0
```

---

## 5. 🪞 Struct with Method (Value Receiver)

```go
func (p Person) Greet() {
    fmt.Println("Hello,", p.Name)
}
```

---

## 6. 🛠 Struct with Method (Pointer Receiver)

```go
func (p *Person) Birthday() {
    p.Age++
}
```

---

## 7. 🧱 Struct Embedding (Composition)

```go
type Address struct {
    City  string
    State string
}

type Employee struct {
    Person
    Address
    Position string
}

e := Employee{
    Person:   Person{"Emma", 28},
    Address:  Address{"New York", "NY"},
    Position: "Developer",
}

fmt.Println(e.Name)  // Access Person field
fmt.Println(e.City)  // Access Address field
```

---

## 8. 🔐 Anonymous Struct

```go
user := struct {
    ID   int
    Name string
}{ID: 1, Name: "Test"}

fmt.Println(user)
```

---

## 9. 🪣 Slice of Structs

```go
people := []Person{
    {"Alice", 30},
    {"Bob", 25},
}

for _, p := range people {
    fmt.Println(p.Name, p.Age)
}
```

---

## 10. 📦 Struct as Function Parameter

```go
func PrintPerson(p Person) {
    fmt.Printf("%s is %d\n", p.Name, p.Age)
}
```

---

## 11. 📦 Struct as Function Return

```go
func NewPerson(name string, age int) Person {
    return Person{Name: name, Age: age}
}
```

---

## 12. 🪛 Pointers to Struct

```go
p := &Person{Name: "Charlie", Age: 35}
p.Age++
fmt.Println(p.Age)
```

---

## 13. 🏷 Struct with Tags (used in JSON, DB, etc.)

```go
type Product struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Price int    `json:"price"`
}
```

---

## 14. 🧩 Nested Structs

```go
type Company struct {
    Name    string
    Founder Person
}

c := Company{
    Name:    "GoCorp",
    Founder: Person{"Jane", 45},
}

fmt.Println(c.Founder.Name)
```

---

## 15. 🧠 Comparing Structs

```go
p1 := Person{"Alice", 30}
p2 := Person{"Alice", 30}

fmt.Println(p1 == p2) // true
```

> Note: All fields must be comparable.

---

## 16. 🚫 Comparing Structs with Non-Comparable Fields

```go
type Data struct {
    Values map[string]int
}

// Cannot compare two `Data` structs directly due to map field
```

---

## 17. 🧪 Struct in a Map

```go
people := map[string]Person{
    "alice": {"Alice", 30},
    "bob":   {"Bob", 25},
}

fmt.Println(people["alice"].Age)
```

---

## 18. ⛓ Struct with Interface Field

```go
type Logger interface {
    Log(msg string)
}

type AppConfig struct {
    Name   string
    Logger Logger
}
```

---

## 19. 🧰 Struct Initialization via `new()`

```go
p := new(Person)
p.Name = "Zara"
fmt.Println(p) // &{Zara 0}
```

---

## 20. 📂 Struct Literals in Return Statements

```go
func CreatePerson() Person {
    return Person{"Tom", 23}
}
```

---

## 🧾 Bonus: Deep Copy vs Shallow Copy

```go
type Job struct {
    Title string
}

type Worker struct {
    Name string
    Job  *Job
}

w1 := Worker{"Sam", &Job{"Engineer"}}
w2 := w1        // shallow copy
w2.Job.Title = "Manager"

fmt.Println(w1.Job.Title) // Output: Manager (because pointer shared)
```

---

## 🎯 Summary Table

| Feature                   | Example Shown       |
| ------------------------- | ------------------- |
| Basic Declaration         | ✅                   |
| Initialization            | ✅                   |
| Field Access/Modification | ✅                   |
| Methods (Value/Pointer)   | ✅                   |
| Embedding                 | ✅                   |
| Anonymous Structs         | ✅                   |
| Slices, Maps              | ✅                   |
| Nested Structs            | ✅                   |
| JSON Tags                 | ✅                   |
| Pointers to Struct        | ✅                   |
| Comparing Structs         | ✅ (and limitations) |

---