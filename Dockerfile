# Start from an official Go image
FROM golang:1.23.1

# Set the working directory inside the container
WORKDIR /app


COPY main.go go.mod ./
COPY Vectors.mp4 ./

# Build the Go application
RUN go build -o video-streaming .

# Expose the port
EXPOSE 3000

# Run the application
CMD ["./video-streaming"]
