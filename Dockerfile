FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod main.go ./

RUN go build -o main .

RUN touch /somefile.txt

EXPOSE 8081
CMD ["./main"]
