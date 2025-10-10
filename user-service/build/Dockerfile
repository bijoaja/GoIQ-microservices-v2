FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o user-service ./cmd/user-service

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/user-service /app/user-service
CMD ["/app/user-service"]
