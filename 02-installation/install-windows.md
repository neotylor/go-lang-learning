Here's a step-by-step guide to install **Go (Golang)** on a **Windows** system:

---

### ✅ Step 1: Download Go

1. Go to the official Go website:
   [https://go.dev/dl/](https://go.dev/dl/)
2. Download the Windows installer (e.g., `go1.21.5.windows-amd64.msi` for 64-bit systems).

---

### ✅ Step 2: Run the Installer

1. Double-click the `.msi` file you downloaded.
2. Follow the installation wizard:

   * Accept the license agreement.
   * Choose the default install path (recommended):
     `C:\Program Files\Go`
   * Click **Next** until you can select **Install**.

---

### ✅ Step 3: Verify the Installation

1. Open **Command Prompt** (`cmd`).
2. Type the following command and press **Enter**:

   ```bash
   go version
   ```

   You should see output like:

   ```
   go version go1.21.5 windows/amd64
   ```

---

### ✅ Step 4: Set Environment Variables (if needed)

The installer usually sets these automatically, but you can double-check:

1. Open **System Properties** > **Environment Variables**.
2. Ensure the following are set:

   * **GOROOT**:

     ```
     C:\Program Files\Go
     ```
   * **GOPATH** (optional, your workspace):

     ```
     C:\Users\<YourUsername>\go
     ```
   * Add to **Path**:

     ```
     C:\Program Files\Go\bin
     ```

---

### ✅ Step 5: Create and Run a Test Program

1. Create a file named `hello.go` with the following content:

   ```go
   package main

   import "fmt"

   func main() {
       fmt.Println("Hello, Go!")
   }
   ```

2. Open Command Prompt where `hello.go` is located and run:

   ```bash
   go run hello.go
   ```

   Output should be:

   ```
   Hello, Go!
   ```

---

Here's a **comprehensive overview of all possible options** for installing **Go (Golang)** on **Windows**, along with **optional configurations and IDE support** to tailor your setup.

---

## ✅ 1. Installation Methods

### 🔹 A. **MSI Installer (Recommended for Beginners)**

* Download from [https://go.dev/dl/](https://go.dev/dl/)
* Simple GUI-based setup.
* Automatically sets `GOROOT` and updates `PATH`.

### 🔹 B. **ZIP Archive (Manual Setup)**

Useful if:

* You want multiple Go versions.
* You don’t want to modify system settings globally.

**Steps**:

1. Download the `.zip` version from [https://go.dev/dl/](https://go.dev/dl/)
2. Extract it to a directory (e.g., `C:\Go\go1.21`)
3. Set environment variables manually:

   * `GOROOT` → `C:\Go\go1.21`
   * Add `C:\Go\go1.21\bin` to `Path`

You can switch versions by pointing `GOROOT` to a different folder.

---

## ✅ 2. Environment Variables (Detailed)

| Variable | Description                                  | Example                           |
| -------- | -------------------------------------------- | --------------------------------- |
| `GOROOT` | Location of the Go SDK                       | `C:\Program Files\Go`             |
| `GOPATH` | Your workspace for projects and dependencies | `C:\Users\YourName\go`            |
| `PATH`   | Must include Go’s `bin` directory            | `%GOROOT%\bin` and `%GOPATH%\bin` |

> **Note**: `GOPATH` is optional for most modern setups since Go modules (using `go.mod`) are now default.

---

## ✅ 3. IDEs and Editor Support

### 🔹 A. **Visual Studio Code (Recommended)**

* Lightweight and full Go support via the **Go extension**.
* Install the Go extension by **Google** (`golang.go`).
* Adds debugging, linting, IntelliSense, and more.

### 🔹 B. **Goland (JetBrains)**

* Full-featured commercial IDE.
* Deep integration and excellent debugging tools.
* Free trial available.

### 🔹 C. **Sublime Text / Atom / Vim / Emacs**

* Use with plugins like `GoSublime`, `vim-go`, or `go-plus`.

---

## ✅ 4. Version Management (Optional but Useful)

### 🔹 A. **gvm (Go Version Manager for Windows)**

There’s no official `gvm` for Windows, but alternatives like:

* **[goenv for Windows](https://github.com/syndbg/goenv)** (WSL preferred)
* **[g](https://github.com/stefanmaric/g)**: cross-platform Go version manager.

Or use **Windows Subsystem for Linux (WSL)** and install `gvm` there.

---

## ✅ 5. Additional Tools and Setup

| Tool          | Purpose                                     |
| ------------- | ------------------------------------------- |
| `godoc`       | Serve documentation for Go code             |
| `gofmt`       | Auto-format Go source files                 |
| `golint`      | Style checker                               |
| `goimports`   | Like `gofmt`, but also adds/removes imports |
| `delve`       | Debugger for Go                             |
| `staticcheck` | Advanced code analysis                      |

You can install most tools via:

```bash
go install golang.org/x/tools/cmd/goimports@latest
```

---

## ✅ 6. Hello World + Modules Setup

**Create a new module project**:

```bash
mkdir hello
cd hello
go mod init hello
```

**Code** (`main.go`):

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Modules!")
}
```

**Run**:

```bash
go run .
```

---

## ✅ 7. Testing Your Installation

Run:

```bash
go env
```

You’ll see all current environment settings, like:

* `GOARCH`
* `GOVERSION`
* `GOMOD`
* `GOPROXY`

---

## ✅ 8. Optional: WSL (Windows Subsystem for Linux)

If you prefer a Linux-style dev environment:

1. Install WSL
2. Set up Ubuntu or Debian
3. Install Go using:

   ```bash
   sudo apt update && sudo apt install golang
   ```

---
