##
## Build
##
# Start from the latest golang base image
FROM golang:buster AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./
COPY go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . ./

ENV GO111MODULE="on" \
  GOARCH="amd64" \
  GOOS="linux" \
  CGO_ENABLED="0"

# Build the Go app
RUN go build -o /server

##
## Deploy
##
FROM alpine:latest

WORKDIR /

COPY --from=build /server /server

RUN chmod -R 777 /server

ENV GIN_MODE release
ENV HOST 0.0.0.0
ENV PORT 4000
EXPOSE 4000

ENTRYPOINT ["/server"]