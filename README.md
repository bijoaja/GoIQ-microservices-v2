# Microservices Modern Template (Go + Gin + GORM + Docker)

This repository is a starter template showcasing a modern, opinionated structure for Go microservices.

Included services:
- `user-service` — CRUD user service (GORM + SQLite)
- `auth-service` — authentication service issuing JWT tokens
- `gateway` — API gateway (proxy + auth middleware)

Structure highlights:
- Each service follows the `cmd/`, `internal/`, `configs/`, `pkg/`, `build/` pattern.
- Root `Makefile` for common developer tasks
- `docker-compose.yml` to run services locally

## Quickstart (local)
1. Build all services:
   ```bash
   make build-all
   ```

2. Run with Docker Compose:
   ```bash
   make up
   ```

3. Or run a single service locally (example: user-service):
   ```bash
   cd user-service
   make run
   ```

## Notes
- JWT secrets and configuration are simplified for demo. Use env vars or secret manager in production.
- SQLite used for simplicity; swap with Postgres in production.
