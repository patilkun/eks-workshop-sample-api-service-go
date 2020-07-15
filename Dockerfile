FROM golang:1.11.1 as builder

RUN mkdir -p /go/src/test
WORKDIR /go/src/test
RUN useradd -u 10001 app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
EXPOSE 8080
CMD ["/go/src/test/main"]