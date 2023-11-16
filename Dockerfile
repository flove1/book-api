FROM golang:1.20.6-alpine3.17 as builder
WORKDIR /app
COPY . .
RUN go build -o main ./cmd
EXPOSE 8080

FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["/app/main"]