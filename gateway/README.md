# API Gateway

Lightweight API gateway that routes requests to backend microservices, enforces authentication (JWT validation), and provides request-level observability. Designed to run standalone or inside the repository's Docker Compose stack.

## Purpose / Responsibilities
- Expose a single, stable API surface for clients.
- Validate and parse JWTs on protected routes.
- Proxy requests to internal services (auth, user, etc.).
- Centralize common concerns: logging, metrics, timeouts, and CORS.

## Features
- Route proxying to upstream services
- JWT validation middleware (Authorization: Bearer <token>)
- Config-driven upstream URLs and ports
- Basic request/response logging and metrics endpoint
- Health check and graceful shutdown

## Quickstart

Prerequisites:
- Go 1.18+
- make
- Docker (optional)

Build and run locally (from repo root or gateway folder):

```sh
# from gateway folder
make build
make run
```

Run the full stack:

```sh
# from repo root
make up
```

## Routes / Examples

- POST /login -> proxied to auth-service
  - No JWT required
  - Example:
    ```sh
    curl -X POST http://localhost:8080/login \
      -H "Content-Type: application/json" \
      -d '{"username":"alice","password":"s3cr3t"}'
    ```

- GET /users -> proxied to user-service
  - Requires Authorization header with Bearer token
  - Example:
    ```sh
    curl http://localhost:8080/users \
      -H "Authorization: Bearer <JWT>"
    ```

- Health & metrics
  - GET /health
  - GET /metrics

## Configuration

Configure via environment variables (or Compose):

- PORT — gateway listen port (default: 8080)
- AUTH_SERVICE_URL — e.g. http://auth-service:8080
- USER_SERVICE_URL — e.g. http://user-service:8080
- JWT_PUBLIC_KEY / JWT_SECRET — key used to verify tokens (depends on signing method)
- LOG_LEVEL — debug|info|warn|error
- TIMEOUT_MS — upstream request timeout

Example docker-compose service snippet:
```yaml
services:
  gateway:
    image: gateway:latest
    environment:
      - PORT=8080
      - AUTH_SERVICE_URL=http://auth-service:8080
      - USER_SERVICE_URL=http://user-service:8080
      - JWT_SECRET=${JWT_SECRET}
    ports:
      - "8080:8080"
    depends_on:
      - auth-service
      - user-service
```

## Security
- Validate tokens on protected routes; reject missing/expired/invalid tokens.
- Prefer asymmetric signing (RS256) in production and distribute public key to gateway.
- Use TLS at the ingress (reverse proxy / load balancer).

## Observability
- Structured logs with request IDs
- /metrics endpoint for Prometheus scraping
- Health checks for orchestration

## Development
- Entry: gateway/cmd/gateway/main.go
- Project layout: gateway/internal/... and gateway/cmd/...
- Run unit tests:
  ```sh
  go test ./... -v
  ```

## Testing & CI
- Add unit tests for middleware and proxy logic
- Use integration tests in CI to verify route proxying and JWT validation
- Example CI steps: go test, go vet, staticcheck, build Docker image

## Contributing
- Follow repo-level contribution guidelines
- Keep PRs small, include tests for behavior changes, update docs when adding routes or config

## License
See the repository LICENSE file for project license and contribution terms.
```// filepath: d:\my_project\golang_project\GoIQ-microservices.v2\gateway\README.md
# API Gateway

Lightweight API gateway that routes requests to backend microservices, enforces authentication (JWT validation), and provides request-level observability. Designed to run standalone or inside the repository's Docker Compose stack.

## Purpose / Responsibilities
- Expose a single, stable API surface for clients.
- Validate and parse JWTs on protected routes.
- Proxy requests to internal services (auth, user, etc.).
- Centralize common concerns: logging, metrics, timeouts, and CORS.

## Features
- Route proxying to upstream services
- JWT validation middleware (Authorization: Bearer <token>)
- Config-driven upstream URLs and ports
- Basic request/response logging and metrics endpoint
- Health check and graceful shutdown

## Quickstart

Prerequisites:
- Go 1.18+
- make
- Docker (optional)

Build and run locally (from repo root or gateway folder):

```sh
# from gateway folder
make build
make run
```

Run the full stack:

```sh
# from repo root
make up
```

## Routes / Examples

- POST /login -> proxied to auth-service
  - No JWT required
  - Example:
    ```sh
    curl -X POST http://localhost:8080/login \
      -H "Content-Type: application/json" \
      -d '{"username":"alice","password":"s3cr3t"}'
    ```

- GET /users -> proxied to user-service
  - Requires Authorization header with Bearer token
  - Example:
    ```sh
    curl http://localhost:8080/users \
      -H "Authorization: Bearer <JWT>"
    ```

- Health & metrics
  - GET /health
  - GET /metrics

## Configuration

Configure via environment variables (or Compose):

- PORT — gateway listen port (default: 8080)
- AUTH_SERVICE_URL — e.g. http://auth-service:8080
- USER_SERVICE_URL — e.g. http://user-service:8080
- JWT_PUBLIC_KEY / JWT_SECRET — key used to verify tokens (depends on signing method)
- LOG_LEVEL — debug|info|warn|error
- TIMEOUT_MS — upstream request timeout

Example docker-compose service snippet:
```yaml
services:
  gateway:
    image: gateway:latest
    environment:
      - PORT=8080
      - AUTH_SERVICE_URL=http://auth-service:8080
      - USER_SERVICE_URL=http://user-service:8080
      - JWT_SECRET=${JWT_SECRET}
    ports:
      - "8080:8080"
    depends_on:
      - auth-service
      - user-service
```

## Security
- Validate tokens on protected routes; reject missing/expired/invalid tokens.
- Prefer asymmetric signing (RS256) in production and distribute public key to gateway.
- Use TLS at the ingress (reverse proxy / load balancer).

## Observability
- Structured logs with request IDs
- /metrics endpoint for Prometheus scraping
- Health checks for orchestration

## Development
- Entry: gateway/cmd/gateway/main.go
- Project layout: gateway/internal/... and gateway/cmd/...
- Run unit tests:
  ```sh
  go test ./... -v
  ```

## Testing & CI
- Add unit tests for middleware and proxy logic
- Use integration tests in CI to verify route proxying and JWT validation
- Example CI steps: go test, go vet, staticcheck, build Docker image

## Contributing
- Follow repo-level contribution guidelines
- Keep PRs small, include tests for behavior changes, update docs when adding routes or config

## License
See the repository LICENSE file for project license and contribution terms.