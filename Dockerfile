FROM golang:1.22 as builder
WORKDIR /app
COPY . /app
RUN GO111MODULE=auto CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org go build -o app cmd/app/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .

ENTRYPOINT ["./app"]
