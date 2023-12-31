FROM golang:latest

WORKDIR /go/src/app

COPY . . 

RUN go build -o experimental-service-delta

EXPOSE 9003

CMD ["./experimental-service-delta"]