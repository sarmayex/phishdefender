FROM golang:1.19-alpine AS build

RUN apk add --no-cache git make bash

# Set the Current Working Directory inside the container
WORKDIR /app/defender

# We want to populate the module cache based on the go.{mod,sum} files.
COPY src/go.mod .
COPY src/go.sum .

RUN go mod download

COPY . .

# Building the Go binanry file
WORKDIR /app/defender/src

RUN CGO_ENABLED=0 GOOS=linux go build -a -o ./out/defender .

FROM alpine AS final
RUN apk add --no-cache tzdata

# Copy the compiled file to final light weight image
COPY --from=build /app/defender/src/out/defender /defender

# HTTP port
EXPOSE 8000

## Run the binary
ENTRYPOINT ["/defender"]

