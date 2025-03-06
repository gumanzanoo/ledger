FROM golang:1.24.0 AS builder

WORKDIR /app

COPY . .

RUN go test ./domain/transactions
RUN go test ./domain/vo

RUN go mod tidy
RUN go build -o myapp .

FROM ubuntu:latest

WORKDIR /app

COPY --from=builder /app/api .

ENTRYPOINT ["./api"]