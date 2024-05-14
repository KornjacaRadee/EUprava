# Use the official Go image as a builder stage
FROM golang:alpine as build_container

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download Go modules
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o server

# Start a new stage from scratch
FROM alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the built executable from the previous stage
COPY --from=build_container /app/server /usr/bin/

# Create a directory for logging and copy the log file
RUN mkdir logging
COPY ./logging/log.log ./logging/

# Expose the port on which the application will run
EXPOSE 8081

# Command to run the executable
ENTRYPOINT ["server"]
