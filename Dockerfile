FROM golang:1.21.6

WORKDIR /go/src/app

COPY . .

EXPOSE 8080

RUN go build -o main cmd/api/main.go

CMD ["./main"]