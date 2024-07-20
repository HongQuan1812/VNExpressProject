# Start with a base golang image that includes Go tools pre-installed.
FROM golang:1.22.4-alpine

# Set the current working directory inside the container.
WORKDIR /app

# Copy the Go modules manifests and download dependencies using go mod.
COPY ./go.mod ./go.sum ./

RUN go mod download
RUN mkdir VNExpress_selector
RUN mkdir server

COPY ./server/server.go ./server/
COPY ./VNExpress_selector ./VNExpress_selector/


# Command to run the executable when the container starts.
CMD ["go", "run", "./server/server.go"]
