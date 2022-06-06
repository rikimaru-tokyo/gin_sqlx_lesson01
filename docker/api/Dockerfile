FROM golang:latest

COPY src/api /go/src/api
WORKDIR /go/src/api/

RUN go mod download
RUN go build

# CMD ["go","run","main.go"]