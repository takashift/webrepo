// アクセスToken( JWT ) は Cookie に保存。
// Cookie からのデータの読み込みとアクセストークンのチェックは認証のあるページでのみ行う。
// Cookie から取り出したアクセストークンは、同じくヘッダーの Authorization にコピー。
// 最後に JWT のミドルウェアで Token を解読させる。

package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dchest/uniuri"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
	"github.com/yosssi/ace"
	gomail "gopkg.in/gomail.v2"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	// 2017-10-06 17:20:00 では何故か出来なかった。。
	timeLayout = "2006-01-02 15:04:05"

	host         = "webrepo.nal.ie.u-ryukyu.ac.jp"
	sendMailAdrr = "Webrepo@nal.ie.u-ryukyu.ac.jp"
)

var (
	tablename = "userinfo"
	seq       = 1
	// ここで指定している Unixソケット の場所は Echoコンテナ のパス
	conn, dberr = dbr.Open("mysql", "rtuna:USER_PASSWORD@unix(/usock/mysqld.sock)/Webrepo", nil)
	dbSess      = conn.NewSession(nil)
)

type googleUser struct {
	// 先頭が大文字でないと格納されない。
	Email string `json:"email"`
}

// e.Renderer に代入するために必須っぽい
type AceTemplate struct {
}

// サイトで共通情報
type ServiceInfo struct {
	Title string
}

type PageValue struct {
	Query string
	Error string
}

type EvalForm struct {
	Genre interface{} `db:"genre"`
	Media interface{} `db:"media"`
	Tag   interface{}
}

type (
	// データベースのテスト
	userinfoJSON struct {
		ID    int    `json:"id"`
		Email string `json:"email"`
	}

	userinfo struct {
		ID            int    `db:"id"`
		OAuthService  string `db:"OAuth_service"`
		OAuthUserinfo string `db:"OAuth_userinfo"`
		Email         string `db:"email"`
		Password      string `db:"password"`
		Name          string `db:"name"`
		SignupDate    string `db:"signup_date"`
		SafeSearch    int    `db:"safe_search"`
		NGCount       int    `db:"NG_count"`
	}

	tmpuser struct {
		OAuthService string `db:"OAuth_service"`
		Act          string `db:"act"`
		Email        string `db:"email"`
		Referer      string `db:"referer"`
		SendTime     string `db:"send_time"`
	}

	PageStatus struct {
		ID          int    `db:"id"`
		Title       string `db:"title"`
		URL         string `db:"URL"`
		RegistDate  string `db:"regist_date"`
		LastUpdate  string `db:"last_update"`
		AdminUserID int    `db:"admin_user_id"`
		Genre       string `db:"genre"`
		Media       string `db:"media"`
		Dead        int    `db:"dead"`
		Tag1        string `db:"tag1"`
		Tag2        string `db:"tag2"`
		Tag3        string `db:"tag3"`
		Tag4        string `db:"tag4"`
		Tag5        string `db:"tag5"`
		Tag6        string `db:"tag6"`
		Tag7        string `db:"tag7"`
		Tag8        string `db:"tag8"`
		Tag9        string `db:"tag9"`
		Tag10       string `db:"tag10"`
	}

	pagePath struct {
		Page string
		URL  string
	}
)

// ここでレシーバ変数を定義したことでAceTemplateに以下の関数がメソッドとして関連付けられる
func (at *AceTemplate) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// オプションを渡してディレクトリを指定
	tpl, err := ace.Load(name, "", &ace.Options{
		BaseDir: "views",
	})

	if err != nil {
		return err
	}

	return tpl.Execute(w, data)
}

// // リファラーURLがこのサイトのものか確認する関数
// func refererCheck(refererURL string) string {
// 	var redirect string
// 	if strings.Contains(refererURL, host) {
// 		redirect = refererURL
// 	} else {
// 		redirect = "/"
// 	}
// 	return redirect
// }

func createJwt(c echo.Context, id int, email string) error {

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["email"] = email
	// claims["exp_time"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("oppai"))
	if err != nil {
		return err
	}

	sess, _ := session.Get("session", c)

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["token"] = t
	sess.Save(c.Request(), c.Response())

	return nil
}

// Token によってサインイン状況をチェック（ログインが必須でないページ）
// サインインの状況に応じてページの一部を変更する
func signinCheck(page string, c echo.Context, value interface{}) error {
	// if client != nil {
	// 	// もしログイン済みなら、
	// 	// 上部メニューの"ログイン"のところを変更する
	// 	searchForm.Login = ""
	// }
	return c.Render(http.StatusOK, page, value)
}

func cookieToAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		sess, err := session.Get("session", c)
		fmt.Fprintf(os.Stderr, "reqURL:%s\n", c.Request().URL)
		if err != nil || sess.Values["token"] == nil {
			// リクエストされたURLを記入
			sess.Options = &sessions.Options{
				Path:     "/",
				MaxAge:   86400 * 7,
				HttpOnly: true,
			}
			sess.Values["refererURL"] = fmt.Sprint(c.Request().URL)
			sess.Save(c.Request(), c.Response())

			// サインイン画面へリダイレクト
			return c.Redirect(http.StatusTemporaryRedirect, "/signin_select")
		}

		t := sess.Values["token"].(string)
		// リクエストヘッダーの Authorization に JWT を格納
		c.Request().Header.Set(echo.HeaderAuthorization, "Bearer "+t)

		fmt.Fprintf(os.Stderr, "token:%v\n", sess)
		return next(c)
	}
}

func getPageStatusItem(id int) (EvalForm, PageStatus) {
	var (
		genreSL []string
		genre   struct {
			// Xはインデックスの意
			X1     string
			X2     string
			X3     string
			X4     string
			X5     string
			X6     string
			X7     string
			X8     string
			X9     string
			X10    string
			Select string
		}
		mediaSL []string
		media   struct {
			// Xはインデックスの意
			X1     string
			X2     string
			X3     string
			X4     string
			X5     string
			X6     string
			Select string
		}
		evalForm   EvalForm
		pageStatus PageStatus
	)

	dbSess.Select("genre").From("page_status_item").Load(&genreSL)
	dbSess.Select("media").From("page_status_item").Load(&mediaSL)

	if id >= 0 {
		dbSess.Select("genre", "media", "tag1", "tag2", "tag3", "tag4", "tag5", "tag6", "tag7", "tag8", "tag9", "tag10").
			From("page_status").
			Where("id = ?", id).
			Load(&pageStatus)

		for i, v := range genreSL {
			if pageStatus.Genre == v {
				genre.Select = fmt.Sprintf("genreX%d", i+1)
				fmt.Printf("%s\n", v)
			}
		}

		for i, v := range mediaSL {
			if pageStatus.Media == v {
				media.Select = fmt.Sprintf("mediaX%d", i+1)
			}
		}
	}

	evalForm.Tag = pageStatus.Tag1 + "\n" +
		pageStatus.Tag2 + "\n" +
		pageStatus.Tag3 + "\n" +
		pageStatus.Tag4 + "\n" +
		pageStatus.Tag5 + "\n" +
		pageStatus.Tag6 + "\n" +
		pageStatus.Tag7 + "\n" +
		pageStatus.Tag8 + "\n" +
		pageStatus.Tag9 + "\n" +
		pageStatus.Tag10

	genre.X1 = genreSL[0]
	genre.X2 = genreSL[1]
	genre.X3 = genreSL[2]
	genre.X4 = genreSL[3]
	genre.X5 = genreSL[4]
	genre.X6 = genreSL[5]
	genre.X7 = genreSL[6]
	genre.X8 = genreSL[7]
	genre.X9 = genreSL[8]
	genre.X10 = genreSL[9]
	evalForm.Genre = genre

	media.X1 = mediaSL[0]
	media.X2 = mediaSL[1]
	media.X3 = mediaSL[2]
	media.X4 = mediaSL[3]
	media.X5 = mediaSL[4]
	media.X6 = mediaSL[5]
	evalForm.Media = media

	return evalForm, pageStatus
}

func main() {
	// テンプレートに渡す値
	var (
		searchForm = PageValue{
			Query: "",
		}
		mailForm = PageValue{
			Error: "",
		}
	)

	var (
		e = echo.New()

		googleOauthConfig = &oauth2.Config{
			ClientID:     "370442566774-osi0bgsn710brv1v3uc1s7hk24blhdq2.apps.googleusercontent.com",
			ClientSecret: "E46tGSdcop7sU9L8pF30Nz_u",
			Endpoint:     google.Endpoint,
			RedirectURL:  "https://webrepo.nal.ie.u-ryukyu.ac.jp/oauth2callback_google",
			Scopes: []string{
				"email"},
		}

		oauthService string

		client *http.Client

		// キャリアメールのドメイン
		domain = []string{
			"docomo.ne.jp",
			"ezweb.ne.jp",
			"au.com",
			"willcom.com",
			"y-mobile.ne.jp",
			"emobile.ne.jp",
			"ymobile1.ne.jp",
			"ymobile.ne.jp",
			"softbank.ne.jp",
			"vodafone.ne.jp",
			"i.softbank.jp",
			"disney.ne.jp",
		}
	)

	// ここで入れるべき構造体はinterfaceによって必須のメソッドが定義され、持つべき引数が決まっている。GoDoc参照。
	e.Renderer = &AceTemplate{}

	// Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	userGoogle := new(googleUser)

	// "/" の時に返すhtml
	e.GET("/", func(c echo.Context) error {
		return signinCheck("search_top", c, searchForm)
	})

	// 検索時に呼び出される
	e.GET("/search", func(c echo.Context) error {
		// URLクエリパラメータを受け取る
		q := c.QueryParam("q")
		searchForm.Query = q
		return signinCheck("search_result", c, searchForm)
	})
	// サインイン方法選択画面
	e.GET("/signin_select", func(c echo.Context) error {
		fmt.Println("signin_select")
		// return c.Render(http.StatusOK, "signin_select", searchForm)
		return c.Redirect(http.StatusTemporaryRedirect, "/google_OAuth")
	})

	// パスワードサインインフォーム
	// e.GET("/pass_signin", func(c echo.Context) error {
	// 	return c.Render(http.StatusOK, "pass_signin", searchForm)
	// })

	// Google の認証画面にリダイレクト
	e.GET("/google_OAuth", func(c echo.Context) error {
		oauthStateString := uniuri.New()
		url := googleOauthConfig.AuthCodeURL(oauthStateString)
		return c.Redirect(http.StatusTemporaryRedirect, url)
	})

	// Google からのコールバック
	e.GET("/oauth2callback_google", func(c echo.Context) error {
		code := c.FormValue("code")
		token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
		if err != nil {
			panic(err)
		}

		client = googleOauthConfig.Client(oauth2.NoContext, token)

		// JSON が返ってくる
		response, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
		if err != nil {
			panic(err)
		}

		defer response.Body.Close()

		err = json.NewDecoder(response.Body).Decode(userGoogle)
		if err != nil {
			panic(err)
		}

		var (
			redirect   string
			userInfoDB userinfo
		)

		// OAuth、キャリアメールが本登録されてるか確認
		_, err = dbSess.Select("ID", "email").From("userinfo").
			Where("OAuth_userinfo = ?", userGoogle.Email).
			Load(&userInfoDB)

		if err == nil {
			// エラーが無い == 登録済み場合
			// リファラーURLがこのサイトのものか確認する
			createJwt(c, userInfoDB.ID, userInfoDB.Email)
			fmt.Println("登録済み")

			sess, _ := session.Get("session", c)
			rURL := sess.Values["refererURL"].(string)
			sess.Values["refererURL"] = nil
			sess.Save(c.Request(), c.Response())

			fmt.Fprintf(os.Stderr, "%s\n", rURL)

			redirect = rURL
		} else {
			// エラーを吐いた == 中身が入ってない場合
			oauthService = "Google"
			redirect = "/OAuth_signup"
		}

		return c.Redirect(http.StatusTemporaryRedirect, redirect)
	})

	// OAuth認証サインアップ（同意）フォーム
	e.GET("/OAuth_signup", func(c echo.Context) error {
		mailForm.Error = ""
		return c.Render(http.StatusOK, "OAuth_signup", mailForm)
	})

	// 同意後のアドレス確認促進画面
	e.POST("/agree_signup", func(c echo.Context) error {
		email := c.FormValue("email")
		fmt.Fprintf(os.Stderr, "%s\n", email)

		// 既に本登録されているユーザーとアドレスが被ってないか確認
		emailDB, err := dbSess.Select("email").From("userinfo").
			Where("email = ?", email).
			ReturnString()

		if err == nil {
			fmt.Fprintf(os.Stderr, "userinfo.email:%s\n", emailDB)
			mailForm.Error = "既に登録済みのメールアドレスです。"
			return c.Render(http.StatusOK, "OAuth_signup", mailForm)
		}

		// メールアドレスがキャリアのドメインか確認する。
		if !strings.Contains(email, "@") {
			mailForm.Error = "正しいアドレスを入力して下さい（全角ではメールが届きません。半角で入力して下さい）。"
			return c.Render(http.StatusOK, "OAuth_signup", mailForm)
		}

		eDomainSlice := strings.SplitAfter(email, "@")

		// スライスなので文字列型に結合
		eDomain := strings.Join(eDomainSlice, "")

		// キャリアドメインをリストに入れて for で比較
		for i := 0; i < len(domain); i++ {
			fmt.Fprintf(os.Stderr, "%s : %s\n", eDomain, domain[i])

			if strings.Contains(eDomain, domain[i]) {

				// メールアドレスが登録されてないのでメールと関連付けたURLを発行
				mac := hmac.New(sha256.New, []byte(uniuri.New()))
				mac.Write([]byte(email))
				act := hex.EncodeToString(mac.Sum(nil))

				// リファラーURLがこのサイトのものか確認する
				sess, err := session.Get("session", c)
				if err != nil {
					return err
				}
				rURL := sess.Values["refererURL"].(string)

				fmt.Fprintf(os.Stderr, "act:%s\n", act)

				t := time.Now()
				tF := t.Format(timeLayout)

				// 一時ユーザーのテーブルにアドレスと認証コード、リファラーURLを一緒に保存
				result, err := dbSess.InsertInto("tmp_user").
					Columns("OAuth_service", "act", "email", "referer", "send_time").
					Values(oauthService, act, email, rURL, tF).
					Exec()

				if err != nil {
					panic(err)
				} else {
					fmt.Fprintf(os.Stderr, "insert:%s\n", result)
				}

				// メールを送信する
				m := gomail.NewMessage()
				m.SetHeader("From", sendMailAdrr)
				m.SetHeader("To", email)
				m.SetHeader("Subject", "メールアドレスの確認")
				m.SetBody("text/plain", "WebRepo☆彡 に登録いただきありがとうございます。\nメールアドレスの確認を行うため、以下のURLへアクセスして下さい。\nなお、このメールの送信から12時間が経過した場合、このURLは無効となるので再度メールアドレスの登録をお願いします。\nhttps://webrepo.nal.ie.u-ryukyu.ac.jp/email_check?act="+act)

				d := gomail.Dialer{Host: "smtp.eve.u-ryukyu.ac.jp", Port: 587, Username: "e145771@eve.u-ryukyu.ac.jp", Password: "USER_PASSWORD"}
				if err := d.DialAndSend(m); err != nil {
					panic(err)
				}

				return c.Render(http.StatusOK, "agree_signup", searchForm)
			}
		}

		mailForm.Error = "ドメイン(@以降)が携帯キャリアのドメイン以外です。登録できません。"
		return c.Render(http.StatusOK, "OAuth_signup", mailForm)
	})

	// メール確認URLへのアクセス時の処理
	e.GET("/email_check", func(c echo.Context) error {
		act := c.QueryParam("act")
		var tmpUser tmpuser
		_, err := dbSess.Select("act", "OAuth_service", "email", "referer").From("tmp_user").
			Where("act = ?", act).
			Load(&tmpUser)

		if err != nil {
			panic(err)
		}
		if tmpUser.Act == "" {
			return errors.New("認証コードが違う!\n")
		}

		fmt.Fprintf(os.Stderr, "act: %s\n", tmpUser.Act)

		t := time.Now()
		tF := t.Format(timeLayout)
		fmt.Fprintf(os.Stderr, "time: %s\n", tF)

		// 正規のユーザーテーブルに追加
		result, err := dbSess.InsertInto("userinfo").
			Columns("OAuth_service", "OAuth_userinfo", "email", "signup_date").
			Values(tmpUser.OAuthService, userGoogle.Email, tmpUser.Email, tF).
			Exec()
		if err != nil {
			panic(err)
		} else {
			fmt.Fprintf(os.Stderr, "insert userinfo:%s\n", result)
		}

		// 一時ユーザーのテーブルから削除
		result, err = dbSess.DeleteFrom("tmp_user").Where("email = ?", tmpUser.Email).Exec()
		if err != nil {
			panic(err)
		} else {
			fmt.Fprintf(os.Stderr, "delete tempuser:%s\n", result)
		}

		return c.Redirect(http.StatusTemporaryRedirect, tmpUser.Referer)
	})

	// 評価閲覧画面
	e.GET("/preview_evaluation", func(c echo.Context) error {
		return signinCheck("preview_evaluation", c, searchForm)
	})

	// 個別評価閲覧画面
	e.GET("/individual_reviews", func(c echo.Context) error {
		return signinCheck("individual_review", c, searchForm)
	})

	// 通報完了画面
	e.GET("/dengerous_complete", func(c echo.Context) error {
		return signinCheck("dengerous_complete", c, searchForm)
	})

	// 利用規約
	e.GET("/term_of_service", func(c echo.Context) error {
		return signinCheck("term_of_service", c, searchForm)
	})

	// このサイトについて
	e.GET("/about", func(c echo.Context) error {
		return signinCheck("about", c, searchForm)
	})

	// Restricted group
	r := e.Group("/r", cookieToAuthMiddleware)
	// Token によってサインイン状況をチェック（ログインが必須なページ）
	r.Use(middleware.JWT([]byte("oppai")))
	r.GET("/test", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		id := int(claims["id"].(float64))
		email := claims["email"].(string)
		return c.String(http.StatusOK, "Welcome "+fmt.Sprint(id)+" "+email+"!")
		// return signinCheckJWT(pagePath{Page: "mypage_top", URL: "/mypage"}, c)
	})

	// マイページ
	r.GET("/mypage", func(c echo.Context) error {
		return c.Render(http.StatusOK, "mypage_top", searchForm)
	})

	// 新規ページ登録画面
	r.GET("/register_page", func(c echo.Context) error {
		evalForm, _ := getPageStatusItem(-1)

		return c.Render(http.StatusOK, "register_page", evalForm)
	})
	r.POST("/register_page", func(c echo.Context) error {
		var newPS PageStatus
		newPS.URL = c.FormValue("url")
		newPS.Genre = c.FormValue("genre")
		newPS.Media = c.FormValue("media")
		tag := strings.Split(c.FormValue("tag"), "\n")

		// structVal := reflect.Indirect(reflect.ValueOf(newPS))
		// structVal.Field(i? + 9).Set(v)

		// for i, v := range tag {
		// 	structVal.Field(i + 9).Set(v)
		// }
		newPS.Tag1 = tag[0]
		newPS.Tag2 = tag[1]
		newPS.Tag3 = tag[2]
		newPS.Tag4 = tag[3]
		newPS.Tag5 = tag[4]
		newPS.Tag6 = tag[5]
		newPS.Tag7 = tag[6]
		newPS.Tag8 = tag[7]
		newPS.Tag9 = tag[8]
		newPS.Tag10 = tag[9]
		// fmt.Printf("tag10:%s\n", structVal.Field())

		newPS.RegistDate = time.Now().Format(timeLayout)
		// newPS.

		_, err := dbSess.InsertInto("page_status").
			Columns("*").
			Values(newPS).
			Exec()

		if err != nil {
			panic(err)
		}

		return c.Redirect(http.StatusMovedPermanently, "../")
	})

	// ページ属性編集画面
	r.GET("/edit_page_cate/:id", func(c echo.Context) error {
		// v := reflect.Indirect()

		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return err
		}

		evalForm, _ := getPageStatusItem(idInt)

		return signinCheck("edit_page_cate", c, evalForm)
	})
	r.POST("/edit_page_cate/:id", func(c echo.Context) error {

		return c.Redirect(http.StatusMovedPermanently, "")
	})

	// 評価入力画面
	r.GET("/input_evaluation", func(c echo.Context) error {
		return c.Render(http.StatusOK, "input_evaluation", searchForm)
	})

	// コメント入力画面
	r.GET("/input_comment", func(c echo.Context) error {
		return c.Render(http.StatusOK, "input_comment", searchForm)
	})

	// ソケット生成
	os.Remove("/usock/domain.sock")
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
