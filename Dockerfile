FROM scratch

WORKDIR $GOPATH/src/gin_blog
COPY . $GOPATH/src/gin_blog

EXPOSE 8000
CMD ["./gin_blog"]