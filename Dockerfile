# syntax=docker/dockerfile:1
FROM golang:1.25rc1-bullseye

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY ./app/go.mod ./app/go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY ./app ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-go

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
EXPOSE 8080

# Run
CMD ["/docker-go"]