# Stage 1: Build the application
FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./
COPY . /app

RUN go build -o tlkm-api ./http/*.go

ENTRYPOINT [ "/app/tlkm-api" ]