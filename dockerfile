
FROM golang:1.21 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN GOOS=linux GOARCH=amd64 go build  -o main ./cmd/main.go

FROM alpine:3.19

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]
