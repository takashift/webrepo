package main

import (
    "os"
    "net"
    "net/http"
    "github.com/labstack/echo"
)

var e = echo.New()

func main() {
    e.GET("/", func(c echo.Context) error {
      // テンプレートに渡す値
      data := struct {
        ServiceInfo
        Content string
      } {
        ServiceInfo: serviceInfo,
        Content: "おっぱい",
      }
    
      // この Render は Echo のメソッドであり、テンプレートエンジンのメソッドではない！
      return c.Render(http.StatusOK, "toppage", data)
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
