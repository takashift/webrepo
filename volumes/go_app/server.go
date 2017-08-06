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

    // OAuth認証サインアップフォーム
    e.GET("/OAuth_signup", func(c echo.Context) error {
      return c.Render(http.StatusOK, "OAuth_signup", searchForm)
    })

    // 同意後のアドレス確認促進画面
    e.GET("/agree_signup", func(c echo.Context) error {
      return c.Render(http.StatusOK, "agree_signup", searchForm)
    })

    // 評価入力画面
    e.GET("/input_evaluation", func(c echo.Context) error {
      return c.Render(http.StatusOK, "input_evaluation", searchForm)
    })

    // 評価閲覧画面
    e.GET("/preview_evaluation", func(c echo.Context) error {
      return c.Render(http.StatusOK, "preview_evaluation", searchForm)
    })

    // 個別評価閲覧画面
    e.GET("/individual_reviews", func(c echo.Context) error {
      return c.Render(http.StatusOK, "individual_review", searchForm)
    })

    // 通報完了画面
    e.GET("/dengerous_complete", func(c echo.Context) error {
      return c.Render(http.StatusOK, "dengerous_complete", searchForm)
    })

    // コメント入力画面
    e.GET("/input_comment", func(c echo.Context) error {
      return c.Render(http.StatusOK, "input_comment", searchForm)
    })

    // 新規ページ登録画面
    e.GET("/register_page", func(c echo.Context) error {
      return c.Render(http.StatusOK, "register_page", searchForm)
    })

    // 新規ページ登録画面
    e.GET("/edit_page_cate", func(c echo.Context) error {
      return c.Render(http.StatusOK, "edit_page_cate", searchForm)
    })

    // ソケット生成
    os.Remove("/usock/domain.sock");
    unix, err := net.Listen("unix", "/usock/domain.sock")
    if err != nil {
      e.Logger.Fatal(err)
    }
    if err := os.Chmod("/usock/domain.sock", 0600); err != nil {
      e.Logger.Fatal(err)
    }
    if err := os.Chown("/usock/domain.sock", 1000, 1000); err != nil {
      e.Logger.Fatal(err)
    }
    e.Listener = unix
    e.Logger.Fatal(e.Start(""))
}
