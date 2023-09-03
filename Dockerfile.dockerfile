# Use the official Go image as the base image
FROM golang:1.20 AS builder

# Set the working directory inside the container
WORKDIR /app

RUN ls 

# Copy the Go source code to the container
COPY . .

# List the contents of the copied directory
RUN ls -R /app

# Install dependencies using go get
RUN go get -d -v ./...


# Build the Go application
RUN go build -o helloworld go.go

# ... (earlier stage instructions)

# Build the Go application
RUN go build -o helloworld go.go

# Ensure the binary is executable
RUN chmod +x helloworld

# ... (later stage instructions)


RUN echo ls

# Use a minimal image as the final base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the builder stage to the final image
COPY --from=builder /app/helloworld .

# Expose the port on which the application runs
EXPOSE 8080

RUN ls -l

RUN ls -l

CMD ["/app/helloworld"]


