
FROM golang:1.21 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -o main ./cmd/main.go

FROM alpine:3.19

RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /app/main ./main

CMD ["./main"]
