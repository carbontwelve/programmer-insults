# Build stage
FROM golang:1.25-alpine AS builder

RUN apk add --no-cache upx

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

# Build and compress the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o programmer-insults .
RUN upx --best programmer-insults

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/programmer-insults .

EXPOSE 8080

CMD ["./programmer-insults"]