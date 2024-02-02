# Dockerfile for terminalbackend.go
FROM golang:1.19

LABEL maintainer="David Fonseca <cosmtrek@gmail.com>"

# Install necessary packages
RUN apt-get update && apt-get install -y golang curl

# Set the working directory
WORKDIR /app

# Copy the Go files
COPY . .

# Install Air for live reloading
RUN go mod download
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b /usr/local/bin

# Expose the necessary ports
EXPOSE 9090
EXPOSE 7070

# Copy the start script
COPY start.sh /start.sh

# Set permissions for the start script
RUN chmod +x /start.sh

# Use tini as the entry point and run the start script
ENTRYPOINT ["/bin/bash", "/start.sh"]