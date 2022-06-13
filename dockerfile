FROM golang:1.18 AS builder

WORKDIR /app

COPY . .

RUN go build -tags netgo -o main.app .

FROM alpine:latest

COPY --from=builder /app/main.app /usr/local/bin

CMD ["main.app"]