package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/dchest/uniuri"
	"github.com/labstack/echo"
	gomail "gopkg.in/gomail.v2"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	// 2017-10-06 17:20:00 では何故か出来なかった。。
	timeLayout = "2006-01-02 15:04:05"
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

	refererURL   string
	userInfoDB   userinfo
	tmpUser      tmpuser
	oauthService string

	tablename = "userinfo"
	seq       = 1
	// ここで指定している Unixソケット の場所は Echoコンテナ のパス
	conn, dberr = dbr.Open("mysql", "rtuna:USER_PASSWORD@unix(/usock/mysqld.sock)/Webrepo", nil)
	sess        = conn.NewSession(nil)

	sendMailAdrr = "Webrepo@nal.ie.u-ryukyu.ac.jp"

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

type GoogleUser struct {
	// 先頭が大文字でないと格納されない。
	Email string `json:"email"`
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
)

func updateUser(c echo.Context) error {
	u := new(userinfoJSON)
	if err := c.Bind(u); err != nil {
		return err
	}

	attrsMap := map[string]interface{}{"id": u.ID, "email": u.Email}
	sess.Update(tablename).SetMap(attrsMap).Where("id = ?", u.ID).Exec()
	return c.NoContent(http.StatusOK)
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	sess.DeleteFrom(tablename).
		Where("id = ?", id).
		Exec()

	return c.NoContent(http.StatusNoContent)
}

// リファラーURLがこのサイトのものか確認する関数
func refererCheck(refererURL string) string {
	var redirect string
	if strings.Contains(refererURL, "https://webrepo.nal.ie.u-ryukyu.ac.jp/") {
		redirect = refererURL
	} else {
		redirect = "/"
	}
	return redirect
}

func main() {
	userGoogle := new(GoogleUser)

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
		// req, err := http.ReadRequest(bufio.NewReader())
		// refererURL := req.Referer
		// Echo の Context.Request が *http.Request 型なので、この中にある Referer() で取ってこれる。
		refererURL = c.Request().Referer()

		fmt.Fprintf(os.Stderr, "%s\n", refererURL)

		// return c.Render(http.StatusOK, "signin_select", searchForm)
		return c.Redirect(http.StatusTemporaryRedirect, "/google_OAuth")
	})

	// パスワードサインインフォーム
	e.GET("/pass_signin", func(c echo.Context) error {
		return c.Render(http.StatusOK, "pass_signin", searchForm)
	})

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

		fmt.Fprintf(os.Stderr, "%s\n", runtime.Version())

		// JSON が返ってくる
		response, err := http.Get("https://www.googleapis.com/oauth2/v3/userinfo?access_token=" + token.AccessToken)
		if err != nil {
			panic(err)
		}

		defer response.Body.Close()

		// var user *GoogleUser

		err = json.NewDecoder(response.Body).Decode(userGoogle)
		// contents, _ := ioutil.ReadAll(response.Body)
		// err = json.Unmarshal(contents, &user)
		if err != nil {
			panic(err)
		}

		var redirect string

		// OAuth、キャリアメールが本登録されてるか確認
		_, err = sess.Select("OAuth_userinfo").From("userinfo").
			Where("OAuth_userinfo = ?", userGoogle.Email).
			ReturnString()

		if err == nil {
			// エラーが無い == 登録済み場合
			// リファラーURLがこのサイトのものか確認する
			redirect = refererCheck(refererURL)
		} else {
			// エラーを吐いた == 中身が入ってない場合
			oauthService = "Google"
			redirect = "/OAuth_signup"
		}

		// fmt.Fprintf(os.Stderr, "callback: %v\n", user)
		// c.Set("email", user.Email)

		return c.Redirect(http.StatusTemporaryRedirect, redirect)
	})

	// OAuth認証サインアップ（同意）フォーム
	e.GET("/OAuth_signup", func(c echo.Context) error {
		// user := c.Get("email").(string)
		// if user != "" {
		// 	fmt.Fprintf(os.Stderr, "%v\n", user)
		// } else {
		// 	fmt.Fprintf(os.Stderr, "NO\n")
		// }
		mailForm.Error = ""
		return c.Render(http.StatusOK, "OAuth_signup", mailForm)
	})

	// 同意後のアドレス確認促進画面
	e.POST("/agree_signup", func(c echo.Context) error {
		email := c.FormValue("email")
		fmt.Fprintf(os.Stderr, "%s\n", email)

		// 既に本登録されているユーザーとアドレスが被ってないか確認
		emailDB, err := sess.Select("email").From("userinfo").
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
				redirect := refererCheck(refererURL)
				fmt.Fprintf(os.Stderr, "act:%s\n", act)

				t := time.Now()
				tF := t.Format(timeLayout)

				// 一時ユーザーのテーブルにアドレスと認証コード、リファラーURLを一緒に保存
				result, err := sess.InsertInto("tmp_user").
					Columns("OAuth_service", "act", "email", "referer", "send_time").
					Values(oauthService, act, email, redirect, tF).
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
		_, err := sess.Select("act", "OAuth_service", "email", "referer").From("tmp_user").
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
		result, err := sess.InsertInto("userinfo").
			Columns("OAuth_service", "OAuth_userinfo", "email", "signup_date").
			Values(tmpUser.OAuthService, userGoogle.Email, tmpUser.Email, tF).
			Exec()
		if err != nil {
			panic(err)
		} else {
			fmt.Fprintf(os.Stderr, "insert userinfo:%s\n", result)
		}

		// 一時ユーザーのテーブルから削除
		result, err = sess.DeleteFrom("tmp_user").Where("email = ?", tmpUser.Email).Exec()
		if err != nil {
			panic(err)
		} else {
			fmt.Fprintf(os.Stderr, "delete tempuser:%s\n", result)
		}

		return c.Redirect(http.StatusTemporaryRedirect, tmpUser.Referer)
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

	e.PUT("/users/", updateUser)
	e.DELETE("/users/:id", deleteUser)

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
