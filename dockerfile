FROM golang:1.24.2 AS builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8080

RUN chmod +x app

CMD ["./app"]
