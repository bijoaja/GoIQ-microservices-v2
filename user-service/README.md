// ...existing code...
# User Service

Lightweight CRUD microservice for managing users. Built with Go, Gin (HTTP router), and GORM (ORM) using SQLite for local development. Designed to be run independently or as part of the GoIQ microservices stack.

Purpose
- Provide stable user management API for other services (via the API Gateway).
- Keep responsibilities narrow: create, read, update, delete users.
- Use simple, testable code and observable endpoints.

Key features
- RESTful endpoints: list, create, retrieve, update, delete users
- Gin-based HTTP server with middleware for logging and request validation
- GORM-backed persistence (SQLite by default)
- Ready for containerized deployment and CI pipelines
- Unit and integration test scaffolding

Quickstart (local)
Prerequisites:
- Go 1.18+
- make
- (optional) Docker

From the service folder:
```sh
# build binary
make build

# run service (reads config from env)
make run
```

Or run full stack from repo root:
```sh
make up
```

Docker
Build image:
```sh
docker build -f user-service/build/Dockerfile -t user-service:latest .
```

Run (example):
```sh
docker run -p 8081:8081 \
  -e PORT=8081 \
  -e DATABASE_URL=users.db \
  user-service:latest
```

Configuration
Configure via environment variables (or Docker Compose):

- PORT — HTTP listen port (default: 8081)
- DATABASE_URL — path or DSN for the database (default: users.db for SQLite)
- LOG_LEVEL — log verbosity (debug|info|warn|error)

API (HTTP)
Base URL: http://localhost:${PORT:-8081}

- GET /users
  - Returns: 200 OK, JSON array of users
  - Example: curl http://localhost:8081/users

- POST /users
  - Creates a new user
  - Request JSON: { "name": "Alice", "email": "alice@example.com" }
  - Returns: 201 Created, created user JSON

- GET /users/:id
  - Returns: 200 OK with user JSON or 404 if not found

- PUT /users/:id
  - Updates user fields
  - Request JSON: { "name": "Alice Updated" }
  - Returns: 200 OK with updated user JSON

- DELETE /users/:id
  - Deletes the user
  - Returns: 204 No Content

Responses use conventional HTTP status codes and JSON bodies for errors:
{ "error": "message" }

Database & Migrations
- Development uses SQLite for simplicity.
- For production, swap to Postgres/MySQL by changing DATABASE_URL and GORM dialect.
- Add migration logic or use an external migrator (recommended for production).

Testing
- Unit tests: go test ./... -v
- Add integration tests that run against an ephemeral SQLite or a test container DB
- CI: run go test, staticcheck, go vet, build image

Development notes
- Entry point: user-service/cmd/user-service/main.go
- Core logic: user-service/internal/...
- Use Gin middleware for logging/request IDs and input validation
- Keep handlers thin; business logic belongs in service layer

Observability & Logging
- Structured logs (JSON) recommended for production
- Expose /health for readiness and liveness checks
- Integrate Prometheus metrics if needed (endpoint /metrics)

Security
- Validate and sanitize user input
- Do not store plaintext secrets; follow secure storage practices
- Enforce authentication & authorization at the Gateway in production

Examples (curl)
Create:
```sh
curl -X POST http://localhost:8081/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice","email":"alice@example.com"}'
```

List:
```sh
curl http://localhost:8081/users
```

Contributing
- Follow repository CONTRIBUTING.md and code style
- Write tests for new behavior and document breaking changes in README
- Keep PRs small and focused

License
See the repository LICENSE file for licensing and contribution terms.
# ...existing code...
```// filepath: d:\my_project\golang_project\GoIQ-microservices.v2\user-service\README.md
// ...existing code...
# User Service

Lightweight CRUD microservice for managing users. Built with Go, Gin (HTTP router), and GORM (ORM) using SQLite for local development. Designed to be run independently or as part of the GoIQ microservices stack.

Purpose
- Provide stable user management API for other services (via the API Gateway).
- Keep responsibilities narrow: create, read, update, delete users.
- Use simple, testable code and observable endpoints.

Key features
- RESTful endpoints: list, create, retrieve, update, delete users
- Gin-based HTTP server with middleware for logging and request validation
- GORM-backed persistence (SQLite by default)
- Ready for containerized deployment and CI pipelines
- Unit and integration test scaffolding

Quickstart (local)
Prerequisites:
- Go 1.18+
- make
- (optional) Docker

From the service folder:
```sh
# build binary
make build

# run service (reads config from env)
make run
```

Or run full stack from repo root:
```sh
make up
```

Docker
Build image:
```sh
docker build -f user-service/build/Dockerfile -t user-service:latest .
```

Run (example):
```sh
docker run -p 8081:8081 \
  -e PORT=8081 \
  -e DATABASE_URL=users.db \
  user-service:latest
```

Configuration
Configure via environment variables (or Docker Compose):

- PORT — HTTP listen port (default: 8081)
- DATABASE_URL — path or DSN for the database (default: users.db for SQLite)
- LOG_LEVEL — log verbosity (debug|info|warn|error)

API (HTTP)
Base URL: http://localhost:${PORT:-8081}

- GET /users
  - Returns: 200 OK, JSON array of users
  - Example: curl http://localhost:8081/users

- POST /users
  - Creates a new user
  - Request JSON: { "name": "Alice", "email": "alice@example.com" }
  - Returns: 201 Created, created user JSON

- GET /users/:id
  - Returns: 200 OK with user JSON or 404 if not found

- PUT /users/:id
  - Updates user fields
  - Request JSON: { "name": "Alice Updated" }
  - Returns: 200 OK with updated user JSON

- DELETE /users/:id
  - Deletes the user
  - Returns: 204 No Content

Responses use conventional HTTP status codes and JSON bodies for errors:
{ "error": "message" }

Database & Migrations
- Development uses SQLite for simplicity.
- For production, swap to Postgres/MySQL by changing DATABASE_URL and GORM dialect.
- Add migration logic or use an external migrator (recommended for production).

Testing
- Unit tests: go test ./... -v
- Add integration tests that run against an ephemeral SQLite or a test container DB
- CI: run go test, staticcheck, go vet, build image

Development notes
- Entry point: user-service/cmd/user-service/main.go
- Core logic: user-service/internal/...
- Use Gin middleware for logging/request IDs and input validation
- Keep handlers thin; business logic belongs in service layer

Observability & Logging
- Structured logs (JSON) recommended for production
- Expose /health for readiness and liveness checks
- Integrate Prometheus metrics if needed (endpoint /metrics)

Security
- Validate and sanitize user input
- Do not store plaintext secrets; follow secure storage practices
- Enforce authentication & authorization at the Gateway in production

Examples (curl)
Create:
```sh
curl -X POST http://localhost:8081/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice","email":"alice@example.com"}'
```

List:
```sh
curl http://localhost:8081/users
```

Contributing
- Follow repository CONTRIBUTING.md and code style
- Write tests for new behavior and document breaking changes in README
- Keep PRs small and focused

License
See the repository LICENSE file for licensing and contribution terms.