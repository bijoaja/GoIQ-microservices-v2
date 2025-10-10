FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o gateway ./cmd/gateway

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/gateway /app/gateway
CMD ["/app/gateway"]
