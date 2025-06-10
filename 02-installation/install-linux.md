To install **Go (Golang)** on a Linux system, follow these steps. These instructions apply to most Linux distributions, including Ubuntu, Debian, CentOS, Fedora, etc.

---

## ✅ Step 1: Download the Go Binary

Go to the [official Go downloads page](https://go.dev/dl/) and copy the URL for the latest stable version for Linux.

Or use `wget` to download it directly in your terminal. For example (as of Go 1.22.3):

```bash
wget https://go.dev/dl/go1.22.3.linux-amd64.tar.gz
```

> Replace the version with the latest one suitable for your system (e.g., `arm64` for Raspberry Pi).

---

## ✅ Step 2: Remove Any Previous Installation (Optional but Recommended)

```bash
sudo rm -rf /usr/local/go
```

---

## ✅ Step 3: Extract the Archive to `/usr/local`

```bash
sudo tar -C /usr/local -xzf go1.22.3.linux-amd64.tar.gz
```

---

## ✅ Step 4: Set Up Go Environment Variables

Add Go to your `PATH`. Open or create the `~/.profile`, `~/.bashrc`, or `~/.zshrc` file (depending on your shell), and add the following:

```bash
export PATH=$PATH:/usr/local/go/bin
```

Also, optionally set the Go workspace:

```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

After saving the file, run:

```bash
source ~/.profile   # or ~/.bashrc or ~/.zshrc
```

---

## ✅ Step 5: Verify the Installation

```bash
go version
```

You should see something like:

```bash
go version go1.22.3 linux/amd64
```

---

## ✅ Step 6: Test It Out

Create a simple test file:

```bash
mkdir -p ~/go-test
cd ~/go-test
nano hello.go
```

Paste this content:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

Run it:

```bash
go run hello.go
```

---

There are **multiple ways to install Go on Linux**, depending on your preferences, system setup, and whether you want the latest version or not. Here's a full breakdown of **all possible installation methods**, along with their **pros and cons**:

---

## 🔹 1. **Official Binary Distribution (Recommended)**

Download and install the precompiled binaries from the official Go website.

### ✅ Pros:

* Always latest version
* Official and trusted
* Easy to uninstall/update

### ❌ Cons:

* Manual setup of environment variables

### 🧭 Steps:

```bash
wget https://go.dev/dl/go1.22.3.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.22.3.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

> See previous message for full details.

---

## 🔹 2. **Using Package Manager**

### A. On Debian/Ubuntu

```bash
sudo apt update
sudo apt install golang-go
```

**Pros**:

* Easy to install and manage
* Automatic updates via APT

**Cons**:

* Often outdated version in the default repository

---

### B. On Fedora/CentOS/RHEL

```bash
sudo dnf install golang
```

or

```bash
sudo yum install golang
```

**Same pros/cons** as above

---

## 🔹 3. **Using Snap Package**

```bash
sudo snap install go --classic
```

**Pros**:

* Latest version maintained by the Snap team
* Isolated from system files

**Cons**:

* Snap may not be pre-installed
* Slower performance and startup for some workloads

---

## 🔹 4. **Using Version Managers (e.g., GVM or ASDF)**

### A. GVM (Go Version Manager)

GitHub: [https://github.com/moovweb/gvm](https://github.com/moovweb/gvm)

```bash
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
source ~/.gvm/scripts/gvm
gvm install go1.22.3
gvm use go1.22.3 --default
```

**Pros**:

* Install and switch between multiple Go versions
* Great for developers managing legacy projects

**Cons**:

* Slightly complex setup

---

### B. ASDF with Go plugin

If you already use `asdf` for managing runtimes (like Node, Python, etc.)

```bash
asdf plugin add golang https://github.com/kennyp/asdf-golang.git
asdf install golang latest
asdf global golang latest
```

---

## 🔹 5. **Build from Source** (Advanced)

If you want total control or need a custom Go build:

```bash
git clone https://go.googlesource.com/go
cd go/src
./make.bash
```

**Pros**:

* Custom builds and patches
* Good for development on Go itself

**Cons**:

* Time-consuming and complex
* Not needed for typical use cases

---

## 🧩 Summary Table

| Method            | Latest Version? | Easy Setup | Multiple Versions | Notes                       |
| ----------------- | --------------- | ---------- | ----------------- | --------------------------- |
| Official Binary   | ✅ Yes           | ⚠️ Manual  | ❌                 | Best balance for most users |
| APT/DNF/YUM       | ❌ Often Old     | ✅ Easy     | ❌                 | Fine for casual usage       |
| Snap              | ✅ Yes           | ✅ Easy     | ❌                 | Sandboxed, slower I/O       |
| GVM/ASDF          | ✅ Yes           | ⚠️ Medium  | ✅ Yes             | Great for developers        |
| Build from Source | ✅ Yes           | ❌ Hard     | ✅ Manual          | Niche and advanced use only |

---
