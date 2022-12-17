# define the base image
FROM golang:1.19.2-alpine

# define the environment variable
ENV ENV=production

# Create app directory
WORKDIR /app

# Copy the go mod and sum files
COPY go.mod go.sum .

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Bundle app source
COPY . .

# Build the Go app
RUN go build -o twitter-bot cmd/main.go

# Command to run the executable
CMD ["./twitter-bot"]
