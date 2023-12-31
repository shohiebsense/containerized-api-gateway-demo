FROM golang:latest

WORKDIR /go/src/app

COPY . . 

RUN go build -o experimental-service-alpha

EXPOSE 9001

CMD ["./experimental-service-alpha"]