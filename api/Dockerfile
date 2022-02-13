FROM golang:1.17.7-alpine

RUN apk update && apk add git

RUN mkdir /go/src/app

WORKDIR /go/src/app

COPY ./api /go/src/app

RUN go mod tidy && go install github.com/cosmtrek/air@v1.27.3

CMD ["air", "-c", ".air.toml"]
