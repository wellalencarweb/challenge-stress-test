FROM golang:1.21-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o stress-cli ./cmd/stress-cli/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/stress-cli .
ENTRYPOINT ["./stress-cli"]
