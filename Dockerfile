FROM golang:1.9.4-alpine3.7


COPY ./src /go/src
WORKDIR /go/src/api

RUN apk add --no-cache git mercurial
RUN go-wrapper download  # "go get -d -v ./..." 
RUN go-wrapper install    # "go install -v ./..."
RUN apk del git mercurial


CMD ["go-wrapper", "run"]
EXPOSE 8080