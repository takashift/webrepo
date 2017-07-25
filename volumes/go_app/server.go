package main

import (
    "os"
    "net"
    "net/http"
    "github.com/labstack/echo"
)

var e = echo.New()

func main() {

    e.GET("/test", func(c echo.Context) error {
    
      // この Render は Echo のメソッドであり、テンプレートエンジンのメソッドではない！
      // この関数の第３引数がテンプレート{{.}}になる
      return c.Render(http.StatusOK, "test", data)
    })

    // "/" の時に返すhtml
    e.GET("/", func(c echo.Context) error {
      return c.Render(http.StatusOK, "search_top", searchForm)
    })

    // 検索時に呼び出される
    e.GET("/search", func(c echo.Context) error {
      // URLクエリパラメータを受け取る
      q := c.QueryParam("q")
      searchForm.Query = q
      return c.Render(http.StatusOK, "search_result", searchForm)
    })

    // サインイン方法選択画面
    e.GET("/signin_select", func(c echo.Context) error {
      return c.Render(http.StatusOK, "signin_select", searchForm)
    })

    // パスワードサインインフォーム
    e.GET("/pass_signin", func(c echo.Context) error {
      return c.Render(http.StatusOK, "pass_signin", searchForm)
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
