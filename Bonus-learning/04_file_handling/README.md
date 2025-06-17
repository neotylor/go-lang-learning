Here's a complete **Go File Handling Tutorial**, organized into a **course-style format** with **all key concepts, use cases, and examples** using the `os`, `io`, and `bufio` packages.

---

# 📘 Course: Mastering File Handling in Go

> Learn everything from creating, reading, writing, and manipulating files in Go using standard libraries like `os`, `io`, and `bufio`.

---

## 📦 Module 1: Getting Started

### 🔹 Lesson 1.1: Import Required Packages

```go
import (
	"os"
	"fmt"
	"io"
	"bufio"
)
```

---

## 📄 Module 2: File Creation & Opening

### 🔹 Lesson 2.1: Create a New File

```go
file, err := os.Create("example.txt")
if err != nil {
	log.Fatal(err)
}
defer file.Close()
fmt.Fprintln(file, "Hello, File!")
```

### 🔹 Lesson 2.2: Open an Existing File (Read-Only)

```go
file, err := os.Open("example.txt")
if err != nil {
	log.Fatal(err)
}
defer file.Close()
```

### 🔹 Lesson 2.3: Open File with Permissions

```go
file, err := os.OpenFile("example.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
	log.Fatal(err)
}
defer file.Close()
```

---

## ✍️ Module 3: Writing to Files

### 🔹 Lesson 3.1: `Write` and `WriteString`

```go
file, _ := os.Create("write.txt")
defer file.Close()

file.Write([]byte("Hello from Write\n"))
file.WriteString("Hello from WriteString\n")
```

### 🔹 Lesson 3.2: Buffered Writing (`bufio.Writer`)

```go
file, _ := os.Create("buffered.txt")
defer file.Close()

writer := bufio.NewWriter(file)
writer.WriteString("Buffered write example\n")
writer.Flush() // Required to actually write to file
```

---

## 📖 Module 4: Reading from Files

### 🔹 Lesson 4.1: Read Entire File (`os.ReadFile`)

```go
data, err := os.ReadFile("example.txt")
fmt.Println(string(data))
```

### 🔹 Lesson 4.2: Read in Chunks (`io.Reader`)

```go
file, _ := os.Open("example.txt")
defer file.Close()

buf := make([]byte, 64)
n, _ := file.Read(buf)
fmt.Printf("Read %d bytes: %s\n", n, string(buf[:n]))
```

### 🔹 Lesson 4.3: Buffered Reading (`bufio.Reader`)

```go
file, _ := os.Open("example.txt")
defer file.Close()

reader := bufio.NewReader(file)
line, _ := reader.ReadString('\n')
fmt.Println("Line:", line)
```

### 🔹 Lesson 4.4: Read Line by Line

```go
file, _ := os.Open("example.txt")
defer file.Close()

scanner := bufio.NewScanner(file)
for scanner.Scan() {
	fmt.Println(scanner.Text())
}
```

---

## 🔁 Module 5: File Utilities

### 🔹 Lesson 5.1: Check if File Exists

```go
_, err := os.Stat("example.txt")
if os.IsNotExist(err) {
	fmt.Println("File does not exist")
}
```

### 🔹 Lesson 5.2: Rename File

```go
err := os.Rename("old.txt", "new.txt")
```

### 🔹 Lesson 5.3: Remove File

```go
err := os.Remove("delete.txt")
```

---

## 📁 Module 6: Directory Handling

### 🔹 Lesson 6.1: Create Directory

```go
err := os.Mkdir("mydir", 0755)
```

### 🔹 Lesson 6.2: Create Nested Directories

```go
err := os.MkdirAll("parent/child/grandchild", 0755)
```

### 🔹 Lesson 6.3: List Files in a Directory

```go
files, _ := os.ReadDir(".")
for _, f := range files {
	fmt.Println(f.Name())
}
```

---

## 🔄 Module 7: Advanced Operations

### 🔹 Lesson 7.1: Copy File

```go
func copyFile(src, dst string) error {
	source, _ := os.Open(src)
	defer source.Close()

	dest, _ := os.Create(dst)
	defer dest.Close()

	_, err := io.Copy(dest, source)
	return err
}

copyFile("example.txt", "copy.txt")
```

### 🔹 Lesson 7.2: Append to a File

```go
f, _ := os.OpenFile("append.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
defer f.Close()
f.WriteString("Appended line\n")
```

---

## ⏳ Module 8: Temporary Files

### 🔹 Lesson 8.1: Create Temp File

```go
tmpfile, _ := os.CreateTemp("", "example-*.txt")
defer os.Remove(tmpfile.Name()) // delete after use

tmpfile.WriteString("Temporary content")
```

### 🔹 Lesson 8.2: Create Temp Directory

```go
dir, _ := os.MkdirTemp("", "mydir-*")
fmt.Println("Temp dir:", dir)
```

---

## 🚨 Module 9: Error Handling Patterns

### 🔹 Lesson 9.1: Handle Common File Errors

```go
if errors.Is(err, os.ErrNotExist) {
	fmt.Println("File not found!")
} else if errors.Is(err, os.ErrPermission) {
	fmt.Println("Permission denied!")
}
```

---

## ✅ Module 10: Best Practices

| Tip                  | Description                          |
| -------------------- | ------------------------------------ |
| `defer file.Close()` | Always close files after opening     |
| Use `bufio`          | For efficient reading/writing        |
| Check file size      | For large files before reading fully |
| Use constants        | For file modes (`os.O_APPEND`, etc.) |
| Handle all errors    | Don't ignore file I/O errors         |

---

## 🧪 Bonus Practice Challenges

1. **Word Counter** – Count words in a file.
2. **Log Rotator** – Rename old log files daily.
3. **File Merger** – Merge multiple text files into one.
4. **Directory Walker** – Recursively list files & folders.
5. **Backup Utility** – Copy file to another location with timestamped name.

---

## 📂 Summary Cheat Sheet

| Operation        | Function                                    |
| ---------------- | ------------------------------------------- |
| Create File      | `os.Create`                                 |
| Open File        | `os.Open`, `os.OpenFile`                    |
| Read File        | `os.ReadFile`, `bufio.Scanner`, `file.Read` |
| Write File       | `Write`, `WriteString`, `bufio.Writer`      |
| Copy File        | `io.Copy`                                   |
| Remove File      | `os.Remove`                                 |
| Rename File      | `os.Rename`                                 |
| Create Directory | `os.Mkdir`, `os.MkdirAll`                   |
| Read Directory   | `os.ReadDir`                                |
| Temp File        | `os.CreateTemp`                             |

---