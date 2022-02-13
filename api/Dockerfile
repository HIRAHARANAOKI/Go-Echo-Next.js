FROM golang:1.17.7-alpine

RUN apk update && apk add git

RUN mkdir /go/src/app

WORKDIR /go/src/app

COPY ./api /go/src/app

# COPY go.mod .
# COPY go.sum .

RUN go mod tidy && go install github.com/cosmtrek/air@v1.27.3

COPY . .

CMD ["air", "-c", ".air.toml"]
