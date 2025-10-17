# ======================
# Stage 1: Builder
# ======================
FROM golang:1.25.2-alpine AS builder

WORKDIR /app

# Copy module files for dependency caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build optimized binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/gateway

# ======================
# Stage 2: Runtime
# ======================
FROM alpine:3.18
WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/main /app/main

# (Optional) Copy env/config files
# COPY .env /app/.env

EXPOSE 8080

CMD ["/app/main"]
