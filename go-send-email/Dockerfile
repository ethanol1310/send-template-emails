# Dockerfile References: https://docs.docker.com/engine/reference/builder/

FROM golang:latest
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
COPY data ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Command to run the executable
CMD ./main send -c=/app/data/customers/customers.csv -t=/app/data/template/template.json -o=/app/output.json -e=/app/errors.csv