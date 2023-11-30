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

# Copy the source code into container for compiling
COPY . /app

# Cleanup previous binary if exists
RUN rm -rf /app/main

# Configure github auth
RUN git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/sinatriiobimo".insteadOf "https://github.com/sinatriiobimo"

# Compile the binary
RUN go get -v golang.org/x/tools/cmd/goimports
RUN go mod download golang.org/x/net
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -v -o tlkm-api ./http/main.go

# Copy release binary that already compiled into distroless
COPY --from=base-golang /app/tlkm-api /app/tlkm-api

CMD ["/app/tlkm-api"]