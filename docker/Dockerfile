FROM golang:1.16.1-alpine

WORKDIR /go/src/store
# COPY go.mod .
# RUN go get -d -v .

COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build .

CMD [ "hn-website" ]
