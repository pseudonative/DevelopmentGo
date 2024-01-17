# FROM golang:1.21-alpine 

# WORKDIR /app

# COPY go.* ./ 

# RUN go mod download

# COPY . . 

# RUN go build -o main . 

# EXPOSE 8080 

# CMD [ "./main" ]

# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -o main .

# Final stage
FROM alpine
COPY --from=builder /app/main /main
EXPOSE 8080
CMD ["/main"]
