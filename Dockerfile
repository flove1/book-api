FROM golang:1.20.6-alpine3.17 as builder
WORKDIR /app
COPY . .
RUN apk add --no-cache git
RUN go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go build -o main ./cmd
EXPOSE 8080

FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/main .
COPY config.yaml .
EXPOSE 8080
CMD ["/app/main"]