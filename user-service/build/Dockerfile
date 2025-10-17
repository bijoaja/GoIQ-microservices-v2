# Stage 1: Build
FROM golang:1.25.2-alpine AS builder
WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev sqlite-dev

# Enable CGO for sqlite3
ENV CGO_ENABLED=1

# Copy dan install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source dan build binary
COPY . .
RUN go build -o main ./cmd/user-service

# Stage 2: Runtime
FROM alpine:3.18
WORKDIR /app

# Install runtime deps for SQLite (tidak perlu compiler)
RUN apk add --no-cache sqlite-libs

# Copy binary hasil build
COPY --from=builder /app/main .

EXPOSE 8081
CMD ["./main"]
