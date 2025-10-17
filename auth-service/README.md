# Auth Service

Simple, focused authentication microservice that issues JWT tokens for client authentication and service-to-service calls.

## Overview

This service handles user authentication and issues JSON Web Tokens (JWTs). It is designed to be run as an independent microservice or inside the project-wide Docker Compose setup.

- Service: Auth microservice
- Primary endpoint: `POST /login`
- Purpose: Authenticate credentials and return a signed JWT

## Quickstart

Prerequisites:
- Go toolchain (1.18+ recommended)
- make
- Docker (optional, for containers)
- Project-level orchestration: [docker-compose.yml](docker-compose.yml) and [Makefile](Makefile)

Build and run locally (from repo root or the service folder):

```sh
# from service folder
make build
make run
```

Or run the full stack:

```sh
# from repo root
make up
```

## Docker

Build the service image:

```sh
docker build -f auth-service/build/Dockerfile -t auth-service:latest .
```

Run with Docker Compose:

```sh
docker-compose up --build
```

Refer to [auth-service/build/Dockerfile](auth-service/build/Dockerfile).

## Endpoints

- POST /login
  - Request: JSON body containing authentication credentials (e.g., username & password).
  - Response: JSON containing a JWT token on successful authentication:
    ```json
    { "token": "<JWT>" }
    ```

Example curl:

```sh
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"alice","password":"s3cr3t"}'
```

## Configuration

Common environment variables (set in your environment or Compose file):
- JWT_SECRET — secret key used to sign tokens
- PORT — HTTP listen port (default often 8080)

Adjust these in your deployment manifests or local environment. Example service configuration and environment is handled alongside other services in [docker-compose.yml](docker-compose.yml) and [setup.sh](setup.sh).

## Development

- Entry point: [`main.main`](auth-service/cmd/auth-service/main.go)
- Source layout:
  - [auth-service/cmd/auth-service/main.go](auth-service/cmd/auth-service/main.go)
  - [auth-service/internal/app](auth-service/internal/app)
  - [auth-service/build/Dockerfile](auth-service/build/Dockerfile)

Integrates with:
- Gateway service: see [gateway/README.md](gateway/README.md) and [gateway/](gateway)
- User service: see [user-service/config.example.yaml](user-service/config.example.yaml) and [user-service/](user-service)

## Useful files

- Project root Makefile: [Makefile](Makefile)
- Compose + orchestration: [docker-compose.yml](docker-compose.yml)
- Repository setup script: [setup.sh](setup.sh)
- Auth service entry: [`main.main`](auth-service/cmd/auth-service/main.go)
- Auth Dockerfile: [auth-service/build/Dockerfile](auth-service/build/Dockerfile)

## Testing & CI

Add unit and integration tests in the service internal packages. Use the repository make targets or CI pipeline to run tests and build artifacts.

## Contributing

Follow repository contribution guidelines in the top-level README. Keep changes small, add tests, and update documentation when behavior or configuration changes.

## License

See the repository LICENSE file.
