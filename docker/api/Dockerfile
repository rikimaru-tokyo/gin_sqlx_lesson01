FROM golang:latest

COPY src/api /go/src/api

WORKDIR /go/src/api/

RUN go get -u github.com/gin-gonic/gin \
  && go get github.com/jmorion/sqlx \
  && go get github.com/go-sql-driver/mysql
