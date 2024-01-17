# FROM golang:1.21-alpine 

# WORKDIR /app

# COPY go.* ./ 

# RUN go mod download

# COPY . . 

# RUN go build -o main . 

# EXPOSE 8080 

# CMD [ "./main" ]

# Build stage
# FROM golang:1.21-alpine AS builder
# WORKDIR /app
# COPY go.* ./
# RUN go mod download
# COPY . .
# RUN CGO_ENABLED=0 go build -trimpath -o main .

# # Final stage
# FROM alpine
# COPY --from=builder /app/main /main
# EXPOSE 8080
# CMD ["/main"]
# Use the builder image to build the Go binary
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -o main .

# Start from a fresh alpine image to create a smaller final image
FROM alpine
# Copy the compiled binary from the builder stage
COPY --from=builder /app/main /main
# Copy the static files from the builder stage
COPY --from=builder /app/profile.html /profile.html
COPY --from=builder /app/jeremypic.png /jeremypic.png
# Add any other static files or directories you need to serve here

# Expose the port the app runs on
EXPOSE 8080

# Run the binary
CMD ["/main"]
