# Gateway

Simple API gateway that proxies requests to services and validates JWT.

Run:
  make run

Build:
  make build

Routes:
  POST /login -> auth-service
  GET /users -> user-service (requires Authorization header)
