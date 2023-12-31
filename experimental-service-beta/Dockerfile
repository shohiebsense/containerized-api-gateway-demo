FROM golang:latest

WORKDIR /go/src/app

COPY . . 

RUN go build -o experimental-service-beta

EXPOSE 9002

CMD ["./experimental-service-beta"]