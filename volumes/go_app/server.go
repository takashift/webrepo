package main

import (
	"os"
	"net"
	"net/http"
	"github.com/labstack/echo"
	"fmt"

	//"strings"
	"github.com/satori/go.uuid"
        "golang.org/x/oauth2"
	v2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/appengine"
	//appenginelog "google.golang.org/appengine/log"

	//_ "github.com/go-sql-driver/mysql"
	//"github.com/gocraft/dbr"
)

var e = echo.New()

var (
	conf = oauth2.Config{
			ClientID:     "370442566774-osi0bgsn710brv1v3uc1s7hk24blhdq2.apps.googleusercontent.com",
			ClientSecret: "E46tGSdcop7sU9L8pF30Nz_u",
			Scopes:       []string{"openid", "email", "profile"}, // 6/19 update: openidのscopeが漏れていたので追加
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://accounts.google.com/o/oauth2/v2/auth",
				TokenURL: "https://www.googleapis.com/oauth2/v4/token",
		},
	}
	// 後で使います
	// 6/11 update:globalで持つのは誤りでした
	state string
)

func init() {
	http.HandleFunc("/api/oauth2", oauth2Handler)
	http.HandleFunc("/oauth2callback", tokenHandler)
}

func oauth2Handler(w http.ResponseWriter, r *http.Request) {
	// ランダムな文字列作成に
	// github.com/satori/go.uuid
	// を使用しています
	state = uuid.NewV4().String()

	// 6/11 update:stateをredirect後と比較する場合はcookieに入れるのが無難です
	sc := &http.Cookie{
		Name:   "hogehoge",
		Value:  state,
		MaxAge: 60,
		Path:   "/",
	}
	redierctURL := getRedirectURL(r.URL.Host)
	conf.RedirectURL = redierctURL
	url := conf.AuthCodeURL(state)
	http.Redirect(w, r, url, 302)
}

func getRedirectURL(host string) string {
	return fmt.Sprintf("https://%s/oauth2callback", host)
}

func tokenHandler(w http.ResponseWriter, r *http.Request) {
	// redirectされたstateと生成したstateが等しいかを確認します
	// 6/11 update:cookieと比較するようにします
	sc, err := r.Cookie("hogehoge")
	if err != nil ||  sc.Value != r.FormValue("state") {
		http.Error(w, "state is invalid.", http.StatusUnauthorized)
		return
	}

	// 認証コードを取得します
	code := r.FormValue("code")
	// appengineのcontextを取得します
	context := appengine.NewContext(r)

	// 認証コードからtokenを取得します
	tok, err := conf.Exchange(context, code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// tokenが正しいことを確認します
	if tok.Valid() == false {
		http.Error(w, "token is invalid.", http.StatusUnauthorized)
		return
	}

	// oauth2 clinet serviceを取得します
	// 特にuserの情報が必要ない場合はスルーです
	service, err := v2.New(conf.Client(context, tok))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// token情報を取得します
	// ここにEmailやUser IDなどが入っています
	// 特にuserの情報が必要ない場合はスルーです
	tokenInfo, err := service.Tokeninfo().AccessToken(tok.AccessToken).Context(context).Do()
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// indexページにリダイレクトします
	// 6/11 update: ハードコードではなく定数を使用するようにした
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}


func main() {

	// Google のOAuth認証の設定
/*	func Google_config() *oauth2.Config {
		config := &oauth2.Config{
			ClientID:     370442566774-osi0bgsn710brv1v3uc1s7hk24blhdq2.apps.googleusercontent.com,
			ClientSecret: E46tGSdcop7sU9L8pF30Nz_u,
			Endpoint: oauth2.Endpoint{
				AuthURL:  authorizeEndpoint,
				TokenURL: tokenEndpoint,
			},
			Scopes:	 []string{"openid", "email", "profile"},
			RedirectURL: "./callback",
		}

		return config
	}
*/


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

	// ページ属性編集画面
	e.GET("/edit_page_cate", func(c echo.Context) error {
	  return c.Render(http.StatusOK, "edit_page_cate", searchForm)
	})

	// テスト（ヘッダーメニュー）
	e.GET("/header_menu", func(c echo.Context) error {
	  return c.Render(http.StatusOK, "header_menu", searchForm)
	})

	// テスト（フッター）
	e.GET("/footer", func(c echo.Context) error {
	  return c.Render(http.StatusOK, "footer", searchForm)
	})

	// テスト（同意書本文）
	e.GET("/consent_form", func(c echo.Context) error {
	  return c.Render(http.StatusOK, "consent_form", searchForm)
	})

	// 利用規約
	e.GET("/term_of_service", func(c echo.Context) error {
	  return c.Render(http.StatusOK, "term_of_service", searchForm)
	})

	// このサイトについて
	e.GET("/about", func(c echo.Context) error {
	  return c.Render(http.StatusOK, "about", searchForm)
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
