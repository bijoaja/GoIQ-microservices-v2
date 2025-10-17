# GoIQ-microservices.v2

Opinionated starter template for production-ready Go microservices using Gin, GORM, and Docker. This repository demonstrates a small, modular microservice ecosystem with an API Gateway, authentication service, and a user CRUD service — ready for local development, CI, and containerized deployment.

## Goals
- Provide a clear, repeatable layout for Go microservices
- Demonstrate best-practice patterns: layered internals, small services, config via environment variables
- Include tooling for local development, testing, and Docker-based deployment

## Included Services
- auth-service — issues JWTs for client and service authentication
  - Path: auth-service/
  - Primary endpoint: POST /login
- user-service — CRUD user API backed by GORM (SQLite for local/dev)
  - Path: user-service/
  - Primary endpoints: GET/POST/PUT/DELETE /users
- gateway — lightweight API gateway that routes, validates JWTs, and exposes health & metrics
  - Path: gateway/

Each service follows a consistent layout:
- cmd/<service>/main.go — entry point
- internal/... — service internals (handlers, services, repository)
- build/ — Dockerfile and image build assets
- configs/ or env — configuration examples

## Quickstart (local)
Prerequisites:
- Go 1.18+
- make
- Docker & docker-compose (for full-stack)

Build and run a single service:
- Windows PowerShell:
  - cd <service-folder>
  - make build
  - make run

Build all services:
- make build-all

Run the full stack with Compose:
- make up
- or
- docker-compose up --build

## Configuration
Services are configured via environment variables. Common variables:
- PORT — HTTP listen port
- JWT_SECRET / JWT_PUBLIC_KEY — token signing/verification
- DATABASE_URL — path/DSN for DB (SQLite in dev)
- AUTH_SERVICE_URL, USER_SERVICE_URL — service discovery via env/Compose

Example docker-compose snippet (service-level):
```yaml
services:
  gateway:
    image: gateway:latest
    environment:
      - PORT=8080
      - AUTH_SERVICE_URL=http://auth-service:8080
      - USER_SERVICE_URL=http://user-service:8081
    ports:
      - "8080:8080"
```

## Development & Tests
- Run unit tests:
  - go test ./... -v
- Recommended tools: golangci-lint, staticcheck, go vet
- Keep handlers thin; place business logic in service packages for easy testing.

## Observability & Security
- Expose /health and /metrics (Prometheus) endpoints where applicable
- Use structured logs (JSON) in production
- Prefer RS256 (asymmetric) JWT signing in production and terminate TLS at the edge

## CI / CD
Suggested pipeline steps:
- go test ./...
- lint (golangci-lint, staticcheck)
- build artifacts
- build and push Docker images
- run integration tests against ephemeral environment (Compose or test cluster)

## Contributing
- Follow repository CONTRIBUTING.md and code style
- Keep PRs small and focused; include tests and docs for new behavior
- Use feature branches and descriptive commit messages

## Useful Links
- auth-service/README.md — auth-specific docs
- user-service/README.md — user service docs
- gateway/README.md — gateway docs
- Makefile — common developer tasks
- docker-compose.yml — local orchestration

## License
See the repository LICENSE file for license and contribution terms.
```// filepath: d:\my_project\golang_project\GoIQ-microservices.v2\README.md
# GoIQ-microservices.v2

Opinionated starter template for production-ready Go microservices using Gin, GORM, and Docker. This repository demonstrates a small, modular microservice ecosystem with an API Gateway, authentication service, and a user CRUD service — ready for local development, CI, and containerized deployment.

## Goals
- Provide a clear, repeatable layout for Go microservices
- Demonstrate best-practice patterns: layered internals, small services, config via environment variables
- Include tooling for local development, testing, and Docker-based deployment

## Included Services
- auth-service — issues JWTs for client and service authentication
  - Path: auth-service/
  - Primary endpoint: POST /login
- user-service — CRUD user API backed by GORM (SQLite for local/dev)
  - Path: user-service/
  - Primary endpoints: GET/POST/PUT/DELETE /users
- gateway — lightweight API gateway that routes, validates JWTs, and exposes health & metrics
  - Path: gateway/

Each service follows a consistent layout:
- cmd/<service>/main.go — entry point
- internal/... — service internals (handlers, services, repository)
- build/ — Dockerfile and image build assets
- configs/ or env — configuration examples

## Quickstart (local)
Prerequisites:
- Go 1.18+
- make
- Docker & docker-compose (for full-stack)

Build and run a single service:
- Windows PowerShell:
  - cd <service-folder>
  - make build
  - make run

Build all services:
- make build-all

Run the full stack with Compose:
- make up
- or
- docker-compose up --build

## Configuration
Services are configured via environment variables. Common variables:
- PORT — HTTP listen port
- JWT_SECRET / JWT_PUBLIC_KEY — token signing/verification
- DATABASE_URL — path/DSN for DB (SQLite in dev)
- AUTH_SERVICE_URL, USER_SERVICE_URL — service discovery via env/Compose

Example docker-compose snippet (service-level):
```yaml
services:
  gateway:
    image: gateway:latest
    environment:
      - PORT=8080
      - AUTH_SERVICE_URL=http://auth-service:8080
      - USER_SERVICE_URL=http://user-service:8081
    ports:
      - "8080:8080"
```

## Development & Tests
- Run unit tests:
  - go test ./... -v
- Recommended tools: golangci-lint, staticcheck, go vet
- Keep handlers thin; place business logic in service packages for easy testing.

## Observability & Security
- Expose /health and /metrics (Prometheus) endpoints where applicable
- Use structured logs (JSON) in production
- Prefer RS256 (asymmetric) JWT signing in production and terminate TLS at the edge

## CI / CD
Suggested pipeline steps:
- go test ./...
- lint (golangci-lint, staticcheck)
- build artifacts
- build and push Docker images
- run integration tests against ephemeral environment (Compose or test cluster)

## Contributing
- Follow repository CONTRIBUTING.md and code style
- Keep PRs small and focused; include tests and docs for new behavior
- Use feature branches and descriptive commit messages

## Useful Links
- auth-service/README.md — auth-specific docs
- user-service/README.md — user service docs
- gateway/README.md — gateway docs
- Makefile — common developer tasks
- docker-compose.yml — local orchestration

## License
See the repository LICENSE file for license and contribution terms.