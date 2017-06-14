FROM golang:alpine

MAINTAINER hgcr

# golang:alpine
RUN apk --no-cache add git
RUN apk --no-cache add gcc
RUN apk --no-cache add musl-dev

# task runner（ファイル監視ツール）
RUN go get -u gopkg.in/godo.v2/cmd/godo

# framework
RUN go get -u github.com/labstack/echo
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/gocraft/dbr
RUN go get -u github.com/labstack/echo/middleware

WORKDIR /go/src/app

CMD ["/go/bin/godo", "server", "--watch"]

#ENTRYPOINT ["go", "get", "github.com/labstack/echo"]
#ENTRYPOINT ["", ""]
#CMD ["go", "run", "/fcgi/server.go"]

#EXPOSE 1323
