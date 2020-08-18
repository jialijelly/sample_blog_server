FROM golang:1.14
WORKDIR /go/src/github.com/jialijelly/sample_blog_server/
ADD . .
RUN go mod download
RUN go build -o sample_blog_server
CMD ["./sample_blog_server"]