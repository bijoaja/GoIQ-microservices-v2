FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o auth-service ./cmd/auth-service

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/auth-service /app/auth-service
CMD ["/app/auth-service"]
