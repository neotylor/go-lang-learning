# Go `switch` Statement – Quick Notes

The `switch` statement in Go is a clean way to write multiple conditional branches. It's often used as an alternative to multiple `if-else` statements.

---

## 🔹 Basic Syntax

```go
switch value {
case condition1:
    // code block
case condition2:
    // code block
default:
    // default code block (optional)
}
````

---

## 📘 Example 1: Basic Day Finder

```go
package main

import "fmt"

func main() {
    day := 3

    switch day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    case 3:
        fmt.Println("Wednesday")
    case 4:
        fmt.Println("Thursday")
    case 5:
        fmt.Println("Friday")
    case 6:
        fmt.Println("Saturday")
    case 7:
        fmt.Println("Sunday")
    default:
        fmt.Println("Invalid day")
    }
}
```

---

## 🔸 Example 2: Switch Without Expression

Use `switch` like `if-else` by omitting the value:

```go
score := 85

switch {
case score >= 90:
    fmt.Println("Grade: A")
case score >= 80:
    fmt.Println("Grade: B")
case score >= 70:
    fmt.Println("Grade: C")
default:
    fmt.Println("Grade: F")
}
```

---

## 🔹 Example 3: Multiple Values in One Case

```go
char := 'a'

switch char {
case 'a', 'e', 'i', 'o', 'u':
    fmt.Println("It's a vowel")
default:
    fmt.Println("It's a consonant")
}
```

---

## ✅ Key Points

* `switch` executes only the first matching `case`; no need for `break`.
* `default` runs if no case matches (optional).
* You can have multiple values in a `case` separated by commas.
* A `switch` can be used without a condition – works like `if-else if`.

---

## 🧪 Practice

Try modifying the examples:

* Add more days to the calendar.
* Change score thresholds.
* Add a `fallthrough` keyword to experiment (though it's rarely needed in Go).

---

## 📚 Resources

* [Go Documentation – switch](https://golang.org/doc/effective_go#switch)
* [Go by Example – Switch](https://gobyexample.com/switch)

---

Happy coding! 🚀