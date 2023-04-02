FROM golang:alpine

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o app ./cmd/main.go

EXPOSE 8080

CMD ["./app"]