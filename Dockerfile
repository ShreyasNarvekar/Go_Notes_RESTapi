# syntax=docker/dockerfile:1

FROM golang:1.25-alpine AS builder
WORKDIR /src

# Download dependencies first for better layer caching.
COPY go.mod go.sum ./
RUN go mod download

# Copy source and build the server binary.
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/server ./cmd/server

FROM alpine:3.20
WORKDIR /app

# Non-root user for safer runtime defaults.
RUN addgroup -S app && adduser -S app -G app

# Copy runtime artifacts: binary + static frontend assets.
COPY --from=builder /out/server /app/server
COPY --from=builder /src/frontend /app/frontend

EXPOSE 8080
USER app

CMD ["/app/server"]

# explain each line of the Dockerfile:
# 1. `FROM golang:1.25-alpine AS builder`: Use the official Go image based on Alpine Linux as the build stage, naming it "builder".
# 2. `WORKDIR /src`: Set the working directory inside the container to `/src` for the build process.
# 3. `COPY go.mod go.sum ./`: Copy the Go module files to the working directory to leverage Docker's layer caching for dependencies.
# 4. `RUN go mod download`: Download the Go module dependencies, which will be cached in subsequent builds unless the module files change.
# 5. `COPY . .`: Copy the entire source code into the container's working directory.
# 6. `RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/server ./cmd/server`: Build the Go application with CGO disabled for a static binary, targeting Linux AMD64 architecture, and output it to `/out/server`.
# 7. `FROM alpine:3.20`: Start a new stage using a minimal Alpine Linux image for the final runtime environment.
# 8. `WORKDIR /app`: Set the working directory for the runtime stage to `/app`.
# 9. `RUN addgroup -S app && adduser -S app -G app`: Create a non-root user and group named "app" for better security.
# 10. `COPY --from=builder /out/server /app/server`: Copy the   built server binary from the builder stage to the runtime stage.
# 11. `COPY --from=builder /src/frontend /app/frontend`: Copy the frontend assets from the builder stage to the runtime stage.
# 12. `EXPOSE 8080`: Declare that the container will listen on port 8080 at runtime.
# 13. `USER app`: Switch to the non-root user for running the application.
# 14. `CMD ["/app/server"]`: Set the default command to run the server binary when the container starts.