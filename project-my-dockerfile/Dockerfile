# syntax=docker/dockerfile:1
# Use base image with golang included
# See more about golang base image on https://hub.docker.com/_/golang
FROM golang:1.21

# Set destination for COPY
WORKDIR /myapp

# Copy go.mod and execute to download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./

COPY AUTHORS.md ./

COPY LINKS.md ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /project-my-compose

# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can (optionally) document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 78

# Run
CMD [ "/project-my-compose" ]