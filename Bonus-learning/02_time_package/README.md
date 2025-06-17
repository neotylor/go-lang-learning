Here is a **complete course-style tutorial** for mastering the **`time` package in Go (Golang)**. It includes **all essential and advanced functions**, organized into modules with **code examples** and **real-world use cases**.

---

# 📘 Course: Mastering the `time` Package in Go

### 📦 Package: `time`

* Import: `import "time"`
* Purpose: Time manipulation, formatting, scheduling, and duration handling.

---

## 🧭 Module 1: Getting Started

### 🔹 Lesson 1.1: Getting the Current Time

```go
now := time.Now()
fmt.Println("Current time:", now)
```

### 🔹 Lesson 1.2: Creating a Time Manually

```go
custom := time.Date(2023, time.December, 25, 10, 30, 0, 0, time.UTC)
fmt.Println("Custom date:", custom)
```

---

## ⏳ Module 2: Time Components

### 🔹 Lesson 2.1: Extracting Fields

```go
now := time.Now()
fmt.Println(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
```

### 🔹 Lesson 2.2: Location & TimeZone

```go
loc := now.Location()
zone, offset := now.Zone()
fmt.Println("Location:", loc, "Zone:", zone, "Offset:", offset)
```

---

## 🕘 Module 3: Time Comparison

### 🔹 Lesson 3.1: `Before`, `After`, `Equal`

```go
t1 := time.Now()
t2 := t1.Add(2 * time.Hour)

fmt.Println(t1.Before(t2)) // true
fmt.Println(t1.After(t2))  // false
fmt.Println(t1.Equal(t2))  // false
```

---

## 🧮 Module 4: Durations and Arithmetic

### 🔹 Lesson 4.1: Adding/Subtracting Time

```go
t := time.Now()
future := t.Add(24 * time.Hour)
past := t.Add(-24 * time.Hour)
```

### 🔹 Lesson 4.2: Time Difference

```go
d := future.Sub(past) // type: time.Duration
fmt.Println("Duration:", d.Hours(), "hours")
```

### 🔹 Lesson 4.3: Constants

```go
fmt.Println(time.Second, time.Minute, time.Hour) // Output: 1s 1m0s 1h0m0s
```

---

## 📅 Module 5: Formatting and Parsing

### 🔹 Lesson 5.1: Predefined Formats

```go
now := time.Now()
fmt.Println(now.Format(time.RFC3339))
fmt.Println(now.Format(time.RFC1123))
```

### 🔹 Lesson 5.2: Custom Format

Go uses a **reference date**: `Mon Jan 2 15:04:05 MST 2006`

```go
now := time.Now()
fmt.Println(now.Format("02-01-2006 15:04:05"))
```

### 🔹 Lesson 5.3: Parsing Strings into Time

```go
t, _ := time.Parse("02-01-2006", "16-06-2025")
fmt.Println("Parsed:", t)
```

### 🔹 Lesson 5.4: ParseInLocation

```go
loc, _ := time.LoadLocation("Asia/Kolkata")
t, _ := time.ParseInLocation("2006-01-02 15:04", "2025-06-16 18:30", loc)
fmt.Println("In IST:", t)
```

---

## ⏱️ Module 6: Timers & Tickers

### 🔹 Lesson 6.1: `time.Sleep`

```go
time.Sleep(2 * time.Second)
fmt.Println("Waited 2 seconds")
```

### 🔹 Lesson 6.2: `time.After`

```go
<-time.After(3 * time.Second)
fmt.Println("Triggered after 3 seconds")
```

### 🔹 Lesson 6.3: `time.NewTimer`

```go
timer := time.NewTimer(2 * time.Second)
<-timer.C
fmt.Println("Timer done")
```

### 🔹 Lesson 6.4: `time.NewTicker`

```go
ticker := time.NewTicker(1 * time.Second)
go func() {
    for t := range ticker.C {
        fmt.Println("Tick at", t)
    }
}()
time.Sleep(5 * time.Second)
ticker.Stop()
```

---

## 🌍 Module 7: Time Zones & UTC

### 🔹 Lesson 7.1: Convert to UTC and Local

```go
now := time.Now()
fmt.Println("UTC:", now.UTC())
fmt.Println("Local:", now.Local())
```

### 🔹 Lesson 7.2: Load and Use Location

```go
ist, _ := time.LoadLocation("Asia/Kolkata")
fmt.Println("In IST:", now.In(ist))
```

---

## 🔢 Module 8: Unix Time

### 🔹 Lesson 8.1: Unix Timestamp

```go
now := time.Now()
fmt.Println(now.Unix())     // Seconds since 1970
fmt.Println(now.UnixMilli()) // Milliseconds
fmt.Println(now.UnixNano())  // Nanoseconds
```

### 🔹 Lesson 8.2: From Unix

```go
t := time.Unix(1718520836, 0) // from seconds
fmt.Println(t)
```

---

## 🧑‍💻 Module 9: Real-World Utilities

### 🔹 Lesson 9.1: Retry with Delay

```go
for i := 0; i < 3; i++ {
    fmt.Println("Try:", i+1)
    time.Sleep(2 * time.Second)
}
```

### 🔹 Lesson 9.2: Execution Timeout

```go
done := make(chan bool)
go func() {
    time.Sleep(4 * time.Second)
    done <- true
}()

select {
case <-done:
    fmt.Println("Finished")
case <-time.After(2 * time.Second):
    fmt.Println("Timeout!")
}
```

---

## ✅ Summary: `time` Package Cheat Sheet

| Feature      | Function                                 |
| ------------ | ---------------------------------------- |
| Current Time | `time.Now()`                             |
| Custom Time  | `time.Date(...)`                         |
| Add/Subtract | `.Add()`, `.Sub()`                       |
| Compare      | `.Before()`, `.After()`, `.Equal()`      |
| Format       | `.Format("02-01-2006")`                  |
| Parse        | `time.Parse()`, `time.ParseInLocation()` |
| Sleep        | `time.Sleep()`                           |
| Timer        | `time.NewTimer()`                        |
| Ticker       | `time.NewTicker()`                       |
| Timezone     | `time.LoadLocation()`, `.In()`           |
| Unix         | `.Unix()`, `time.Unix()`                 |

---

## 🎓 Bonus Module: Practice Projects

1. **Reminder App** – Schedule tasks using `Ticker`
2. **Execution Timer** – Measure function execution time using `time.Since()`
3. **Date Parser CLI** – Convert string → time → UTC
4. **Retry Logic with Backoff** – Use `Sleep` and `Duration` math

---