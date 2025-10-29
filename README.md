# $(basename $(git rev-parse --show-toplevel))

Backend service written in **Go 1.25**.  
Includes Docker setup, Postgres integration, and deployment to GCP.

---

## 📂 Project Structure
\`\`\`
$(tree -L 2 --dirsfirst | sed 's/^/    /')
\`\`\`

## 🚀 Run Locally
\`\`\`bash
go mod tidy
go run ./cmd/api
\`\`\`

## 🧪 Tests
\`\`\`bash
go test ./...
\`\`\`

## 🐳 Docker
\`\`\`bash
docker build -t $(basename $(git rev-parse --show-toplevel)) .
docker run -p 8080:8080 $(basename $(git rev-parse --show-toplevel))
\`\`\`

## 🧭 Environment Variables
- **PORT** – API port (default 8080)
- **DATABASE_URL** – Postgres connection string
- **GCP_PROJECT** – GCP project name

---

## 📄 License
MIT © $(git config user.name)
