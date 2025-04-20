# Etapa 1: compilación
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o teas-api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/teas-api .
COPY data data

EXPOSE 8080

ENV TEA_DATA_PATH=data/tes.json

CMD ["./teas-api"]
