Here is a **comprehensive tutorial/course format** for learning the **`strings` package in Go (Golang)** — covering **all functions** with **examples** and structured as a full course or self-paced tutorial.

---

# 📘 Course: Mastering the `strings` Package in Go

### 📦 Standard Library: `strings`

* Import: `import "strings"`
* Purpose: Provides functions for string manipulation — searching, splitting, trimming, replacing, etc.

---

## 🧰 Module 1: Basics of `strings` Package

### 🔹 Lesson 1.1: Introduction

```go
import "strings"
```

---

## 🔤 Module 2: String Comparison & Checks

### 🔹 Lesson 2.1: `Compare(a, b string) int`

```go
strings.Compare("a", "b") // -1
strings.Compare("a", "a") // 0
```

### 🔹 Lesson 2.2: `EqualFold(a, b string) bool`

Case-insensitive comparison:

```go
strings.EqualFold("Go", "go") // true
```

### 🔹 Lesson 2.3: `Contains(s, substr string) bool`

```go
strings.Contains("hello", "ll") // true
```

### 🔹 Lesson 2.4: `ContainsAny(s, chars string) bool`

```go
strings.ContainsAny("team", "i") // false
strings.ContainsAny("fail", "xyi") // true
```

### 🔹 Lesson 2.5: `ContainsRune(s string, r rune) bool`

```go
strings.ContainsRune("golang", 'g') // true
```

### 🔹 Lesson 2.6: `HasPrefix(s, prefix string) bool`

```go
strings.HasPrefix("golang", "go") // true
```

### 🔹 Lesson 2.7: `HasSuffix(s, suffix string) bool`

```go
strings.HasSuffix("filename.txt", ".txt") // true
```

---

## 🔄 Module 3: String Modification

### 🔹 Lesson 3.1: `Replace(s, old, new string, n int)`

```go
strings.Replace("aaa", "a", "b", -1) // "bbb"
```

### 🔹 Lesson 3.2: `ReplaceAll(s, old, new string)`

```go
strings.ReplaceAll("foo bar foo", "foo", "baz") // "baz bar baz"
```

### 🔹 Lesson 3.3: `ToLower(s string)`

```go
strings.ToLower("GoLang") // "golang"
```

### 🔹 Lesson 3.4: `ToUpper(s string)`

```go
strings.ToUpper("go") // "GO"
```

### 🔹 Lesson 3.5: `Title(s string)` (Deprecated: Use cases are limited)

```go
strings.Title("hello world") // "Hello World"
```

### 🔹 Lesson 3.6: `ToTitle(s string)`

```go
strings.ToTitle("goLang") // "GOLANG"
```

---

## 🧮 Module 4: Splitting & Joining

### 🔹 Lesson 4.1: `Split(s, sep string) []string`

```go
strings.Split("a,b,c", ",") // ["a", "b", "c"]
```

### 🔹 Lesson 4.2: `SplitAfter(s, sep string)`

```go
strings.SplitAfter("a,b,c", ",") // ["a,", "b,", "c"]
```

### 🔹 Lesson 4.3: `SplitN(s, sep string, n int)`

```go
strings.SplitN("a,b,c,d", ",", 2) // ["a", "b,c,d"]
```

### 🔹 Lesson 4.4: `Join(a []string, sep string)`

```go
strings.Join([]string{"a", "b", "c"}, "-") // "a-b-c"
```

---

## 📏 Module 5: Trimming & Cleaning

### 🔹 Lesson 5.1: `Trim(s, cutset string)`

```go
strings.Trim("!!hello!!", "!") // "hello"
```

### 🔹 Lesson 5.2: `TrimSpace(s string)`

```go
strings.TrimSpace("  hello \n") // "hello"
```

### 🔹 Lesson 5.3: `TrimPrefix(s, prefix string)`

```go
strings.TrimPrefix("unhappy", "un") // "happy"
```

### 🔹 Lesson 5.4: `TrimSuffix(s, suffix string)`

```go
strings.TrimSuffix("filename.txt", ".txt") // "filename"
```

### 🔹 Lesson 5.5: `TrimLeft(s, cutset string)`

```go
strings.TrimLeft("...hello", ".") // "hello"
```

### 🔹 Lesson 5.6: `TrimRight(s, cutset string)`

```go
strings.TrimRight("hello...", ".") // "hello"
```

---

## 🔢 Module 6: Counting & Indexing

### 🔹 Lesson 6.1: `Count(s, substr string) int`

```go
strings.Count("cheese", "e") // 3
```

### 🔹 Lesson 6.2: `Index(s, substr string)`

```go
strings.Index("chicken", "ken") // 4
```

### 🔹 Lesson 6.3: `IndexAny(s, chars string)`

```go
strings.IndexAny("golang", "aeiou") // 1
```

### 🔹 Lesson 6.4: `IndexByte(s string, c byte)`

```go
strings.IndexByte("golang", 'g') // 0
```

### 🔹 Lesson 6.5: `IndexRune(s string, r rune)`

```go
strings.IndexRune("hello", 'l') // 2
```

### 🔹 Lesson 6.6: `LastIndex(s, substr string)`

```go
strings.LastIndex("go gopher", "go") // 3
```

---

## 🧑‍💻 Module 7: Field Functions

### 🔹 Lesson 7.1: `Fields(s string) []string`

Splits by white space:

```go
strings.Fields("  foo bar  baz ") // ["foo", "bar", "baz"]
```

### 🔹 Lesson 7.2: `FieldsFunc(s string, f func(rune) bool)`

```go
f := func(r rune) bool {
    return r == ',' || r == ';'
}
strings.FieldsFunc("a,b;c", f) // ["a", "b", "c"]
```

---

## 🧠 Module 8: Reader and Builder

### 🔹 Lesson 8.1: `strings.NewReader(s string)`

```go
reader := strings.NewReader("hello world")
```

### 🔹 Lesson 8.2: `strings.Builder` (Efficient string building)

```go
var b strings.Builder
b.WriteString("Hello, ")
b.WriteString("World!")
fmt.Println(b.String()) // "Hello, World!"
```

---

## 🎓 Final Module: Practice & Projects

### ✅ Exercises:

1. Build a string normalizer (`TrimSpace`, `ToLower`, `ReplaceAll`)
2. Create a CSV parser using `Split` and `FieldsFunc`
3. Custom string builder using `strings.Builder`
4. Word frequency counter with `Fields` and `Count`

---

## ✅ Summary & Cheat Sheet

| Function                | Purpose                        |
| ----------------------- | ------------------------------ |
| `Compare`, `EqualFold`  | Comparison                     |
| `Contains`, `HasPrefix` | Checks                         |
| `Replace`, `ToLower`    | Modification                   |
| `Split`, `Join`         | Tokenization                   |
| `Trim`, `TrimSpace`     | Cleaning                       |
| `Count`, `Index`        | Search                         |
| `Fields`, `FieldsFunc`  | Splitting by whitespace/custom |
| `NewReader`, `Builder`  | Advanced string I/O            |

---