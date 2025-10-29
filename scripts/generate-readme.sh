#!/bin/bash

# === Auto-generate README for Go projects ===

REPO_NAME=$(basename $(git rev-parse --show-toplevel))
AUTHOR=$(git config user.name)

{
  echo "# ${REPO_NAME}"
  echo ""
  echo "Backend service written in **Go 1.25**."
  echo "Includes Docker setup, Postgres integration, and deployment to GCP."
  echo ""
  echo "## 📂 Project Structure"
  echo '```'
  tree -L 2 --dirsfirst | sed 's/^/    /'
  echo '```'
  echo ""
  echo "## 🚀 Run Locally"
  echo '```bash'
  echo "go mod tidy"
  echo "go run ./cmd/api"
  echo '```'
  echo ""
  echo "## 🧪 Tests"
  echo '```bash'
  echo "go test ./..."
  echo '```'
  echo ""
  echo "## 🐳 Docker"
  echo '```bash'
  echo "docker build -t ${REPO_NAME} ."
  echo "docker run -p 8080:8080 ${REPO_NAME}"
  echo '```'
  echo ""
  echo "## 🧭 Environment Variables"
  echo "- **PORT** – API port (default 8080)"
  echo "- **DATABASE_URL** – Postgres connection string"
  echo "- **GCP_PROJECT** – GCP project name"
  echo ""
  echo "## 📄 License"
  echo "MIT © ${AUTHOR}"
} > README.md

echo "✅ README.md generated successfully!"
