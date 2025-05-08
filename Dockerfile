# Start from golang base image
FROM golang:1.24.2-alpine as builder

# Install git and dependencies
RUN apk update && apk add --no-cache git build-base bash

# Install Air (live reloading tool)
RUN go install github.com/air-verse/air@latest

# Add docker-compose-wait script
ENV WAIT_VERSION=2.12.1
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/${WAIT_VERSION}/wait /wait
RUN chmod +x /wait

# Copy the app code and config
COPY . /app
COPY .air.toml /app/.air.toml

# Set working directory
WORKDIR /app

# Tidy Go modules
RUN go mod tidy

# Expose port (will be mapped by Docker Compose)
EXPOSE ${EXPOSE_PORT}

# Wait for DB, then start Air
CMD /wait && air -c .air.toml
