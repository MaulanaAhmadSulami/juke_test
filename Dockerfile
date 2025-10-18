# The build stage
FROM golang:1.25.1 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build app
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/app

FROM alpine:latest

WORKDIR /root/

# Install cert
RUN apk --no-cache add ca-certificates

# copy binary from builder
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]