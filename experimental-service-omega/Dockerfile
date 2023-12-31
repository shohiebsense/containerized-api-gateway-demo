FROM golang:latest

WORKDIR /go/src/app

COPY . . 

RUN go build -o experimental-service-omega

EXPOSE 9004

CMD ["./experimental-service-omega"]