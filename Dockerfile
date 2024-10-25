FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o sales-app ./cmd/main.go

FROM alpine:3.14

WORKDIR /app

COPY --from=builder /app/sales-app .

CMD ["./sales-app"]