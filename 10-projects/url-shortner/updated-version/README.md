
### 🛠️ Setup Instructions

1. **Install dependencies**:

   ```bash
   go get github.com/gorilla/mux github.com/matoous/go-nanoid/v2 modernc.org/sqlite golang.org/x/crypto/bcrypt
   ```
2. **Prep admin credentials**:

   ```go
   hash, _ := hashPassword("yourPassword")
   fmt.Println(hash)
   ```

   Then set:

   ```bash
   export ADMIN_USER=admin
   export ADMIN_PASS=<your bcrypt hash>
   ```
3. **Run the server**:

   ```bash
   go run main.go
   ```
4. **Create a short link**:

   ```bash
   curl -u admin:yourPassword -d '{"URL":"https://example.com"}' -H "Content-Type: application/json" http://localhost:3000/shorten
   ```
5. **Follow the short URL**:
   Browse to `http://localhost:3000/<id>` to redirect or get expired/not found.
