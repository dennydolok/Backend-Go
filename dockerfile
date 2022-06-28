FROM golang:1.18 AS builder

WORKDIR /app

COPY . .

RUN go build -tags netgo -o main.app .

FROM alpine:latest

RUN apk --no-cache add tzdata

COPY --from=builder /app/main.app /usr/local/bin

CMD ["main.app"]