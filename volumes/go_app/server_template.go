package main

import (
    "io"
    "os"
    "net"
    "net/http"
    "html/template"
    "github.com/labstack/echo"
)

type Template struct {
  templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
  return t.templates.ExecuteTemplate(w, name, data)
}

// サイトで共通情報
type ServiceInfo struct {
  Title string
}

var serviceInfo = ServiceInfo {
  "タイトルrrrrrrrrrrr",
}

func main() {
    e := echo.New()

    t := &Template{
      templates: template.Must(template.ParseGlob("views/*.html")),
    }

    e.Renderer = t

    e.GET("/", func(c echo.Context) error {
      // テンプレートに渡す値
      data := struct {
        ServiceInfo
        Content string
      } {
        ServiceInfo: serviceInfo,
        Content: "おっぱい",
      }
      return c.Render(http.StatusOK, "page1", data)
    })

    // ソケット生成
    os.Remove("/usock/domain.sock");
    uni, err := net.Listen("unix", "/usock/domain.sock")
    if err != nil {
      e.Logger.Fatal(err)
    }
    if err := os.Chmod("/usock/domain.sock", 0600); err != nil {
      e.Logger.Fatal(err)
    }
    if err := os.Chown("/usock/domain.sock", 1000, 1000); err != nil {
      e.Logger.Fatal(err)
    }
    e.Listener = uni
    e.Logger.Fatal(e.Start(""))
}
