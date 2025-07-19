# virtual-card-x

This is a simple Golang boilerplate project using Gin, sqlc and PostgreSQL.

## Requirements
- Go 1.20+
- PostgreSQL server
- [sqlc](https://sqlc.dev)

## Setup
1. Install dependencies:
   ```bash
   go mod tidy
   ```
2. Generate database code with sqlc:
   ```bash
   sqlc generate
   ```
3. Run the server (ensure `DATABASE_URL` is set if not using the default local connection):
   ```bash
   go run ./cmd/server
   ```

The server exposes two endpoints:
- `POST /users` with JSON `{"name": "Alice"}` to create a user.
- `GET /users` to list users.
