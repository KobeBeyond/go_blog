FROM golang:latest

WORKDIR $GOPATH/src/gin_blog
COPY . $GOPATH/src/gin_blog
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./gin_blog"]