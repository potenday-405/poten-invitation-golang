# syntax=docker/dockerfile:1

## Multi-stage build start
# Build the application from source
FROM --platform=linux/amd64 golang:1.22 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./cmd

# Run the tests in the container. No code for test...
FROM build-stage AS run-test-stage
#RUN go test -v ./...

# Deploy the application binary into a lean image
FROM atlassian/ubuntu-minimal:latest AS build-release-stage

WORKDIR /

COPY --from=build-stage /server /server

EXPOSE 8080

ENTRYPOINT ["/server"]