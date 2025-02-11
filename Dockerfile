FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /app/main cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main /app/

EXPOSE 8080

ENTRYPOINT ["/app/main"]