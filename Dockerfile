FROM golang:latest

WORKDIR $GOPATH/src/doghow-blog
COPY . $GOPATH/src/doghow-blog
RUN go build . 

EXPOSE 8000
ENTRYPOINT ["./doghow-blog"]