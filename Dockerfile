FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container's working directory
COPY . .

# Build the Go application
RUN go build -o app

# Expose port
EXPOSE 8000

# Run the application
CMD ["./app"]