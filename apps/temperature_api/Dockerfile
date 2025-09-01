# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

## Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/build/web-api ./cmd

# Final stage
FROM scratch

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/build/web-api /app/web-api
COPY --from=builder /app/config/config.yaml /app/config/config.yaml

EXPOSE 8081

CMD ["./web-api"]
