FROM golang:1.26

# App directory
WORKDIR /app

# Setup dependencies
COPY go.mod go.sum ./
RUN go mod download

# Setup app
COPY *.go ./

# Compile app
RUN CGO_ENABLED=0 GOOS=linux go build -o /products-service

# Run app
EXPOSE 8001
CMD ["/products-service"]