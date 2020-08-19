FROM golang:1.14.7-alpine3.12 AS prepare
WORKDIR /go/src/github.com/jialijelly/sample_blog_server/
COPY go.mod .
COPY go.sum .
RUN go mod download

FROM prepare AS build
COPY . .
RUN go build -o sample_blog_server

FROM alpine:3.12 AS run
COPY --from=build /go/src/github.com/jialijelly/sample_blog_server/sample_blog_server .
COPY --from=build /go/src/github.com/jialijelly/sample_blog_server/wait-for-mysql.sh .
COPY --from=build /go/src/github.com/jialijelly/sample_blog_server/config.json .
RUN chmod +x ./wait-for-mysql.sh

ENTRYPOINT ["./wait-for-mysql.sh"]