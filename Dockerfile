FROM golang:1.6.3-alpine
RUN apk --update add git
COPY . /app
WORKDIR /app
RUN go get -d -v
RUN go build -o srvlb .
CMD ["./srvlb"]