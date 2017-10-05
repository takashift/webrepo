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
RUN go get -u github.com/yosssi/ace

# OAuth2
RUN go get -u github.com/satori/go.uuid
RUN go get -u golang.org/x/oauth2
RUN go get -u google.golang.org/api/oauth2/v2
RUN go get -u google.golang.org/appengine
RUN go get -u cloud.google.com/go/compute/metadata
RUN go get -u github.com/dchest/uniuri

WORKDIR /go/src/app

EXPOSE 25

CMD ["/go/bin/godo", "server", "--watch"]

#ENTRYPOINT ["go", "get", "github.com/labstack/echo"]
#ENTRYPOINT ["", ""]
#CMD ["go", "run", "/fcgi/server.go"]
