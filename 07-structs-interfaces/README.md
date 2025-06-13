# 📘 Go Lang Notes: Structs and Interfaces

Welcome! This guide is crafted to help you master **Structs** and **Interfaces** in Go, even if you're just getting started.

---

## 📖 Table of Contents

1. [Structs in Go](#-structs-in-go)
    - [Declaring Structs](#declaring-structs)
    - [Using Structs](#using-structs)
    - [Struct Methods](#struct-methods)
    - [Pointer Receivers](#pointer-receivers)
    - [Embedded Structs (Composition)](#embedded-structs-composition)
2. [Interfaces in Go](#-interfaces-in-go)
    - [Declaring Interfaces](#declaring-interfaces)
    - [Implementing Interfaces](#implementing-interfaces)
    - [Interface Composition](#interface-composition)
    - [Empty Interface (`interface{}`)](#empty-interface-interface)
    - [Type Assertions & Type Switches](#type-assertions--type-switches)
3. [Real-World Use Case](#-real-world-use-case)
4. [Best Practices](#-best-practices)
5. [Summary Table](#-summary-table)

---

## 🧱 Structs in Go

### Declaring Structs

```go
type Person struct {
    Name string
    Age  int
}
```

A `struct` is a custom data type that groups related fields.

### Using Structs

```go
func main() {
    p := Person{Name: "Alice", Age: 30}
    fmt.Println(p.Name) // Output: Alice
}
```

You can also create anonymous structs:

```go
employee := struct {
    Name string
    ID   int
}{Name: "John", ID: 101}
```

---

### Struct Methods

You can associate functions with structs using **methods**:

```go
func (p Person) Greet() {
    fmt.Printf("Hi, I'm %s and I'm %d years old.\n", p.Name, p.Age)
}
```

---

### Pointer Receivers

Use pointer receivers to **modify** struct fields inside methods:

```go
func (p *Person) Birthday() {
    p.Age++
}
```

```go
p := Person{Name: "Bob", Age: 25}
p.Birthday()
fmt.Println(p.Age) // Output: 26
```

---

### Embedded Structs (Composition)

Go doesn't support inheritance but enables **composition** through embedding.

```go
type Address struct {
    City  string
    State string
}

type Employee struct {
    Person
    Address
    Company string
}
```

```go
e := Employee{
    Person:  Person{Name: "Eva", Age: 28},
    Address: Address{City: "NYC", State: "NY"},
    Company: "TechCorp",
}
fmt.Println(e.Name)   // Access Person's Name directly
fmt.Println(e.City)   // Access Address's City directly
```

---

## 🎭 Interfaces in Go

### Declaring Interfaces

```go
type Speaker interface {
    Speak()
}
```

An interface defines **behavior**, not data. Any type that implements the required methods satisfies the interface.

---

### Implementing Interfaces

```go
type Dog struct{}

func (d Dog) Speak() {
    fmt.Println("Woof!")
}

func makeItSpeak(s Speaker) {
    s.Speak()
}
```

Use the function like:

```go
makeItSpeak(Dog{}) // Output: Woof!
```

---

### Interface Composition

You can compose interfaces from other interfaces:

```go
type Reader interface {
    Read() string
}

type Writer interface {
    Write(data string)
}

type ReadWriter interface {
    Reader
    Writer
}
```

---

## 🧍 Interface Example with Multiple Types

```go
type Worker interface {
    Work()
}

type Engineer struct{}
type Doctor struct{}

func (Engineer) Work() { fmt.Println("Engineering solutions...") }
func (Doctor) Work()   { fmt.Println("Treating patients...") }

func Assign(w Worker) {
    w.Work()
}
```

```go
Assign(Engineer{})
Assign(Doctor{})
```

---

### Empty Interface (`interface{}`)

The **empty interface** can represent any type — similar to `Object` in other languages.

```go
func Describe(i interface{}) {
    fmt.Printf("Value: %v, Type: %T\n", i, i)
}
```

---

### Type Assertions & Type Switches

#### Type Assertion

```go
var i interface{} = "hello"
s := i.(string) // safe if you’re sure
```

#### Safe type assertion:

```go
s, ok := i.(string)
if ok {
    fmt.Println("It's a string:", s)
}
```

#### Type Switch

```go
func Identify(i interface{}) {
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

## 🛠 Real-World Use Case

Let’s simulate a payment system using interfaces:

```go
type PaymentProcessor interface {
    Pay(amount float64)
}

type CreditCard struct{}
type PayPal struct{}

func (CreditCard) Pay(amount float64) {
    fmt.Printf("Paid $%.2f with credit card\n", amount)
}

func (PayPal) Pay(amount float64) {
    fmt.Printf("Paid $%.2f using PayPal\n", amount)
}

func ProcessPayment(p PaymentProcessor, amount float64) {
    p.Pay(amount)
}
```

---

## ✅ Best Practices

* **Favor interfaces** over concrete types for parameters
* Keep interfaces **small** (preferably 1-3 methods)
* Use **pointer receivers** if the method modifies the struct
* Avoid naming interfaces as `SomethingInterface` — keep it clean: `Reader`, `Notifier`, `Storer`
* Use `interface{}` only when necessary (e.g., generic utilities)

---

## 📊 Summary Table

| Concept         | Explanation                                                   |
| --------------- | ------------------------------------------------------------- |
| `struct`        | A composite type grouping related fields                      |
| Method          | Function associated with a struct (value or pointer receiver) |
| Embedded Struct | Used for composition (inheritance-like behavior)              |
| Interface       | A type that defines method signatures                         |
| Implicit Impl.  | Types satisfy interfaces automatically                        |
| `interface{}`   | Represents any type                                           |
| Type Assertion  | Extract value from an interface                               |
| Type Switch     | Determine type dynamically at runtime                         |

---

## 🔚 Conclusion

Structs and Interfaces are the **cornerstone of Go's type system**. They enable:

* Encapsulation
* Composition over inheritance
* Polymorphism and abstraction

Understanding these helps write **clean, reusable, testable, and idiomatic** Go code.

---

🛠 **Happy Coding in Go!** 🐹
Feel free to fork, clone, and practice these examples!
