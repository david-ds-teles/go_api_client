FROM golang:1.17-alpine

# because of gcc dependency 
RUN apk add --no-cache build-base

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go build

CMD ["account_client"]