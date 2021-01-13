FROM golang:1.12.0-alphine3.9
RUN mkdir/test
ADD . /test
WORKDIR /app
RUN go build -o main
CMD ["/app/main"]