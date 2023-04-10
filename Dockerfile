FROM golang:latest AS builder

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN GO111MODULE="on" CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
COPY --from=builder /app/configs /app/configs
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /
COPY .env .
ENV TZ=Asia/Almaty
ENV ZONEINFO=/zoneinfo.zip
CMD ["./app"]