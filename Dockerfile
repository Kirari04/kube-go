FROM golang:1.21 as builder
WORKDIR /app
COPY go.mod go.mod
COPY go.sum go.sum
COPY main.go main.go
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o main main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/main /app/main
RUN chmod +X ./main
EXPOSE 1323
CMD ["./main"]