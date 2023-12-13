# Base go image
FROM golang:1.21-alpine3.18 as builder

RUN apk update && apk add --no-cache git

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go mod tidy

RUN go build -o go-crud-food

RUN chmod +x /app/go-crud-food

# build docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/go-crud-food /app

CMD ["/app/go-crud-food"]