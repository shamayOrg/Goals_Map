# Goals_Map

Backend service written in **Go 1.25**.
Includes Docker setup, Postgres integration, and deployment to GCP.

## 📂 Project Structure
```
    .
    ├── bin
    ├── cmd
    │   ├── api
    │   └── migrate
    ├── docs
    ├── internal
    ├── scripts
    │   └── generate-readme.sh
    ├── web
    ├── LICENSE
    ├── README.md
    └── go.mod
    
    9 directories, 4 files
```

## 🚀 Run Locally
```bash
go mod tidy
go run ./cmd/api
```

## 🧪 Tests
```bash
go test ./...
```

## 🐳 Docker
```bash
docker build -t Goals_Map .
docker run -p 8080:8080 Goals_Map
```

## 🧭 Environment Variables
- **PORT** – API port (default 8080)
- **DATABASE_URL** – Postgres connection string
- **GCP_PROJECT** – GCP project name

## 📄 License
MIT © ShamaySapir
