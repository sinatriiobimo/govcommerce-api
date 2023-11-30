# Stage 1: Build the application
FROM golang:1.19 AS base-golang

ARG GITHUB_TOKEN

# Install Base Image Requirement
RUN apt-get update && apt-get install -y \
    bash \
    curl \
    git \
    unzip \
    openssh-client

# Add github into known_hosts record
RUN mkdir ~/.ssh \
    && ssh-keyscan -t rsa github.com >> ~/.ssh/known_hosts

# Always set workdir into application root
WORKDIR /app

# Copy necessary files for dependency resolution
COPY go.mod go.sum ./

# Configure GitHub auth
RUN git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/sinatriiobimo".insteadOf "https://github.com/sinatriiobimo"

# Compile the binary
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -v -o tlkm-api ./http/main.go

# Stage 2: Create the final minimal image
FROM asia-southeast2-docker.pkg.dev/dogwood-wharf-316804/base-image/distroless-go

# Copy release binary
COPY --from=base-golang /app/tlkm-api /app/tlkm-api

CMD ["/app/tlkm-api"]