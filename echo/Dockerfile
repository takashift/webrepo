FROM golang:alpine
# FROM golang:latest

MAINTAINER hgcr

# golang:alpine
RUN apk --no-cache add git
RUN apk --no-cache add gcc
RUN apk --no-cache add musl-dev

# task runner（ファイル監視ツール）
# RUN go get -u gopkg.in/godo.v2/cmd/godo

# framework
RUN go get -u github.com/labstack/echo
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/gocraft/dbr
RUN go get -u github.com/yosssi/ace
RUN go get -u github.com/dgrijalva/jwt-go

# OAuth2
RUN go get -u github.com/satori/go.uuid
RUN go get -u golang.org/x/oauth2
RUN go get -u google.golang.org/api/oauth2/v2
# RUN go get -u google.golang.org/appengine
RUN go get -u cloud.google.com/go/compute/metadata
RUN go get -u github.com/dchest/uniuri
RUN go get -u gopkg.in/gomail.v2

RUN go get -u github.com/labstack/echo-contrib/session
RUN go get -u github.com/gorilla/sessions

# GetしたHTMLのスクレイピング
RUN go get -u github.com/PuerkitoBio/goquery

# 文字コード変換
RUN go get -u github.com/saintfish/chardet
RUN go get -u github.com/yuin/charsetutil

WORKDIR /go/src/app

# CMD ["/go/bin/godo", "server", "--watch"]

# EXPOSE 1323

# godo 不使用時のテスト
CMD ["go", "run", "server.go", "createPrevEvalPage.go"]

#ENTRYPOINT ["", ""]
#CMD ["go", "run", "/fcgi/server.go"]
