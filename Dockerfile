# Use alpine because of its small footprint
FROM golang:1.16-alpine

WORKDIR /app

# Download necessary Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY * .

RUN go build -o /docker-bot-entrypoint

CMD ["/docker-bot-entrypoint"]
