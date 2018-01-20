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
	"math"
	"net"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
	"github.com/yuin/charsetutil"

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
	// 2017-10-06 17:20:00 では何故か出来なかった。。 → この日時じゃないと駄目らしい。
	timeLayout = "2006-01-02 15:04:05"
	tablename  = "userinfo"

	host = "xn--rvz.nal.ie.u-ryukyu.ac.jp" // テスト環境
	// host         = "webrepo.nal.ie.u-ryukyu.ac.jp" 	// 本番環境
	sendMailAdrr = "Webrepo@nal.ie.u-ryukyu.ac.jp"
)

var (
	e = echo.New()

	seq = 1
	// ここで指定している Unixソケット の場所は Echoコンテナ のパス
	conn, dberr = dbr.Open("mysql", "rtuna:USER_PASSWORD@unix(/usock/mysqld.sock)/Webrepo", nil)
	dbSess      = conn.NewSession(nil)
	byte13Str   = string([]byte{13})

	googleOauthConfig = &oauth2.Config{
		// テスト環境
		ClientID:     "370442566774-868h6rc57kmfm82lu4hsviliuo9l6o07.apps.googleusercontent.com",
		ClientSecret: "cX7ua-IKGwIJNsVxILni7vfp",
		// 本番環境
		// ClientID:     "370442566774-osi0bgsn710brv1v3uc1s7hk24blhdq2.apps.googleusercontent.com",
		// ClientSecret: "E46tGSdcop7sU9L8pF30Nz_u",

		Endpoint:    google.Endpoint,
		RedirectURL: "https://" + host + "/oauth2callback_google",
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
		".ac.jp",
	}
)

// e.Renderer に代入するために必須っぽい
type AceTemplate struct {
}

type (
	googleUser struct {
		// 先頭が大文字でないと格納されない。
		Email string `json:"email"`
	}

	EvalForm struct {
		URL   string
		Genre interface{} `db:"genre"`
		Media interface{} `db:"media"`
		Tag   interface{}
	}

	// サイトで共通情報
	PageValue struct {
		PageTitle string
		Query     string
		Error     string
	}

	PrevEvalPageValue struct {
		PageValue
		PageStatus
		Content string
		AveGFP  string
		AveVisP string
	}

	ListPageValue struct {
		EvalForm
		Content         string
		PageStatusSlice []PageStatus
	}

	MyPageValue struct {
		UserName string
		Content  string
	}

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
		OAuthService  string `db:"OAuth_service"`
		OAuthUserinfo string `db:"OAuth_userinfo"`
		Act           string `db:"act"`
		Email         string `db:"email"`
		Referer       string `db:"referer"`
		SendTime      string `db:"send_time"`
	}

	PageStatus struct {
		ID           int    `db:"id"`
		Title        string `db:"title"`
		URL          string `db:"URL"`
		RegisterDate string `db:"register_date"`
		LastUpdate   string `db:"last_update"`
		AdminUserID  int    `db:"admin_user_id"`
		Genre        string `db:"genre"`
		Media        string `db:"media"`
		Dead         int    `db:"dead"`
		Tag1         string `db:"tag1"`
		Tag2         string `db:"tag2"`
		Tag3         string `db:"tag3"`
		Tag4         string `db:"tag4"`
		Tag5         string `db:"tag5"`
		Tag6         string `db:"tag6"`
		Tag7         string `db:"tag7"`
		Tag8         string `db:"tag8"`
		Tag9         string `db:"tag9"`
		Tag10        string `db:"tag10"`
	}

	PageStatusTiny struct {
		ID  int    `db:"id"`
		URL string `db:"URL"`
	}

	IndividualEval struct {
		Num                  int    `db:"num"`
		PageID               int    `db:"page_id"`
		EvaluatorID          int    `db:"evaluator_id"`
		Posted               string `db:"posted"`
		BrowseTime           string `db:"browse_time"`
		BrowsePurpose        string `db:"browse_purpose"`
		Deliberate           int    `db:"deliberate"`
		DescriptionEval      string `db:"description_eval"`
		RecommendGood        int    `db:"recommend_good"`
		RecommendBad         int    `db:"recommend_bad"`
		GoodnessOfFit        int    `db:"goodness_of_fit"`
		BecauseGoodnessOfFit string `db:"because_goodness_of_fit"`
		Device               string `db:"device"`
		Visibility           int    `db:"visibility"`
		BecauseVisibility    string `db:"because_visibility"`
		NumTypo              int    `db:"num_typo"`
		BecauseNumTypo       string `db:"because_num_typo"`
	}

	IndividualEvalComment struct {
		Num             int    `db:"num"`
		PageID          int    `db:"page_id"`
		CommenterID     int    `db:"commenter_id"`
		Posted          string `db:"posted"`
		ReplyEvalNum    int    `db:"reply_eval_num"`
		ReplyCommentNum int    `db:"reply_comment_num"`
		Deliberate      int    `db:"deliberate"`
		Comment         string `db:"comment"`
		RecommendGood   int    `db:"recommend_good"`
		RecommendBad    int    `db:"recommend_bad"`
	}

	Typo struct {
		Num               int    `db:"num"`
		PageID            int    `db:"page_id"`
		EvaluatorID       int    `db:"evaluator_id"`
		IndividualEvalNum int    `db:"individual_eval_num"`
		Incorrect         string `db:"incorrect"`
		Correct           string `db:"correct"`
	}

	pagePath struct {
		Page string
		URL  string
	}

	RecommendSQL struct {
		UpdTable  string
		IntoTable string
		NumColumn string
	}
)

// var IndividualEval = map[string]interface{}{}

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

// リファラーURLがこのサイトのものか確認する関数
func refererCheck(refererURL string) string {
	var redirect string
	if strings.Contains(refererURL, "/r/") {
		redirect = refererURL
	} else {
		redirect = "/"
	}
	return redirect
}

func createJwt(c echo.Context, id int, email string, name string) error {

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["email"] = email
	claims["name"] = name
	// claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

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
// サインインの状況に応じてページの一部���������変更する
func signinCheck(page string, c echo.Context, value interface{}) error {
	// if client != nil {
	// 	// もしログイン�����������������������みなら、
	// 	// 上部メニューの"ログイン"のところを変更する
	// 	searchForm.Login = ""
	// }
	return c.Render(http.StatusOK, page, value)
}

func cookieToHeaderAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
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
			sess.Values["refererURL"] = refererCheck(c.Request().URL.String())
			sess.Save(c.Request(), c.Response())

			// サインイン画面へリダイレクト
			return c.Redirect(http.StatusSeeOther, "/signin_select")
		}

		t := sess.Values["token"].(string)
		// リクエストヘッダーの Authorization に JWT を格納
		c.Request().Header.Set(echo.HeaderAuthorization, "Bearer "+t)

		fmt.Fprintf(os.Stderr, "token:%v\n", sess)
		return next(c)
	}
}

func getPageStatusItem(c echo.Context, id int) (EvalForm, PageStatus) {
	var (
		// ジャンルや媒体の追加時の変更箇所その１
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
			X11    string
			None   string
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
			None   string
			Select string
		}
		evalForm   EvalForm
		pageStatus PageStatus

		genreQ = c.QueryParam("genre")
		mediaQ = c.QueryParam("media")
	)

	dbSess.Select("genre").From("page_status_item").Load(&genreSL)
	dbSess.Select("media").From("page_status_item").Load(&mediaSL)

	if id > 0 {
		dbSess.Select("genre", "media", "tag1", "tag2", "tag3", "tag4", "tag5", "tag6", "tag7", "tag8", "tag9", "tag10").
			From("page_status").
			Where("id = ?", id).
			Load(&pageStatus)

		for i, v := range genreSL {
			if pageStatus.Genre == v {
				genre.Select = "genreX" + strconv.Itoa(i+1)
				fmt.Printf("%s\n", v)
			}
		}

		for i, v := range mediaSL {
			if pageStatus.Media == v {
				media.Select = "mediaX" + strconv.Itoa(i+1)
			}
		}
	} else if genreQ != "" && mediaQ != "" {
		genre.None = "未選択"
		media.None = "未選択"

		if genreQ == "*" {
			genre.Select = "genre" + genreQ
		} else if genreQ == "選択して下さい" {
			genre.Select = "genre_none"
		} else {
			for i, v := range genreSL {
				if genreQ == v {
					genre.Select = "genreX" + strconv.Itoa(i+1)
				}
			}
		}

		if mediaQ == "*" {
			media.Select = "media" + mediaQ
		} else if mediaQ == "選択して下さい" {
			media.Select = "media_none"
		} else {
			for i, v := range mediaSL {
				if mediaQ == v {
					media.Select = "mediaX" + strconv.Itoa(i+1)
				}
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

	// ジャンルや媒体の追加時の変更箇所その２
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
	genre.X11 = genreSL[10]
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

func incrementRecommend(c echo.Context, arg RecommendSQL) error {
	var (
		recommStatus    string
		updRecommColumn string
	)

	pageID := c.Param("pageID")

	num := c.Param("num")
	numInt, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}
	fmt.Println("Atoi OK")

	// 参考になった or ならなかったを取得
	if c.FormValue("recommend") == "なった👍" {
		updRecommColumn = "recommend_good"
		recommStatus = "good"
	} else {
		updRecommColumn = "recommend_bad"
		recommStatus = "bad"
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := int(claims["id"].(float64))

	// 多重に押されるのを防止するためにボタンを押したユーザーを記録する
	_, err = dbSess.InsertInto(arg.IntoTable).
		Columns(arg.NumColumn, "user_id", "recommend").
		Values(numInt, userID, recommStatus).
		Exec()
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/preview_evaluation/"+pageID)
	}

	// DBをインクリメントする
	// string型は+で繋げないとエラーになる。それ以外は?で置き換える。
	_, err = dbSess.UpdateBySql("UPDATE "+arg.UpdTable+" SET "+updRecommColumn+" = "+updRecommColumn+" + 1 WHERE num = ?",
		numInt).Exec()
	if err != nil {
		panic(err)
	}

	return c.Redirect(http.StatusSeeOther, "/preview_evaluation/"+pageID)
}

func reportDangerous(c echo.Context, updTable string, numColumn string) error {
	// 審議中は2
	deliberate := 2

	num := c.Param("num")
	numInt, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}
	fmt.Println("Atoi OK")

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := int(claims["id"].(float64))

	// 通報されまくると困るから一応ユーザーを記録
	_, err = dbSess.InsertInto("dangerous_log").
		Columns("user_id", numColumn).
		Values(userID, numInt).
		Exec()
	if err != nil {
		panic(err)
	}

	// DBの審議中カラムを1にする
	_, err = dbSess.Update(updTable).
		Set("deliberate", deliberate).
		Where("num = ?", numInt).
		Exec()
	if err != nil {
		panic(err)
	}

	return c.Render(http.StatusOK, "dangerous_complete", nil)
}

func insertComment(c echo.Context) error {
	pageID := c.Param("pageID")
	pageIDInt, err := strconv.Atoi(pageID)
	if err != nil {
		panic(err)
	}

	evalNum := c.Param("evalNum")
	evalNumInt, err := strconv.Atoi(evalNum)
	if err != nil {
		panic(err)
	}

	num := c.Param("num")
	numInt, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}
	fmt.Println("Atoi OK")

	comment := c.FormValue("comment")

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := int(claims["id"].(float64))

	_, err = dbSess.InsertInto("individual_eval_comment").
		Columns("page_id", "commenter_id", "reply_eval_num", "reply_comment_num", "comment").
		Values(pageIDInt, userID, evalNumInt, numInt, comment).
		Exec()
	if err != nil {
		panic(err)
	}

	return c.Redirect(http.StatusSeeOther, "/preview_evaluation/"+pageID)
}

func inputEval(c echo.Context) error {
	indEval := new(IndividualEval)
	// indEvalLoad := new(IndividualEval)
	typo := new(Typo)

	var (
		err            error
		browseTime     time.Time
		corrNoNullSL   []string
		incorrNoNullSL []string
		rURL           string
	)

	// 評価者の ID を取得
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	indEval.EvaluatorID = int(claims["id"].(float64))
	// indEval.EvaluatorID = 1

	bro := strings.Replace(c.FormValue("browse"), "T", " ", -1)
	fmt.Println(bro)
	// 時刻のフォーマットが正しくセットできてない時は DB に値を入れない
	browseTime, err = time.Parse("2006-01-02 15:04", bro)
	if err != nil {
		browseTime, err = time.Parse("2006-01-02", bro)
	}
	if err != nil {
		fmt.Println("time型に出来ない")
	} else {
		indEval.BrowseTime = browseTime.Format(timeLayout)
		fmt.Println(indEval.BrowseTime)
	}

	// フォームの評価を取得
	indEval.BrowsePurpose = c.FormValue("purpose")
	indEval.DescriptionEval = c.FormValue("freedom")
	indEval.GoodnessOfFit, err = strconv.Atoi(c.FormValue("rating_pp"))
	if err != nil {
		panic(err)
	}
	indEval.Device = c.FormValue("device")
	indEval.Visibility, err = strconv.Atoi(c.FormValue("rating_vw"))
	if err != nil {
		// panic(err)
	}

	incorr := c.FormValue("typo")
	corr := c.FormValue("typo_answer")
	// []byte{13}を削除して、カンマで区切る
	incorr = strings.Replace(incorr, byte13Str+"\n", "\n", -1)
	corr = strings.Replace(corr, byte13Str+"\n", "\n", -1)
	fmt.Println(incorr)
	// typo のスライスを作成
	incorrSL := strings.Split(incorr, "\n")
	corrSL := strings.Split(corr, "\n")
	// 空白を除外したスライスを作成
	for _, v := range incorrSL {
		if v != "" {
			incorrNoNullSL = append(incorrNoNullSL, v)
		}
	}
	for _, v := range corrSL {
		if v != "" {
			corrNoNullSL = append(corrNoNullSL, v)
		}
	}
	fmt.Println(incorrNoNullSL)

	pageIDStr := c.Param("id")
	url := c.FormValue("url")
	if url != "" {

		var (
			newPS PageStatus
			dbPS  PageStatus
		)
		newPS.URL = url

		// URL のプロトコルが https でも http でも無い時は戻る。
		if !strings.HasPrefix(newPS.URL, "https://") && !strings.HasPrefix(newPS.URL, "http://") {
			return c.Render(http.StatusOK, "input_evaluation_url", nil)
		}

		// 末尾のスラッシュを削除
		newPS.URL = strings.TrimSuffix(newPS.URL, "/")

		// https:// も http:// も取り除いた変数を用意
		uri := strings.TrimPrefix(newPS.URL, "https://")
		uri = strings.TrimPrefix(uri, "http://")

		// 同じURIが登録されてないかチェック
		_, err := dbSess.Select("id", "URL").From("page_status").
			Where("url = ? OR url = ?", "http://"+uri, "https://"+uri).
			Load(&dbPS)
		if err != nil {
			panic(err)
		}

		// 同じ URI が登録されていない時
		if dbPS.URL == "" {

			newPS.Title = "ページの属性（ジャンル、媒体、タグ）を編集して下さい（タイトルは自動で取得します）。"
			newPS.Genre = "選択して下さい"
			newPS.Media = "選択して下さい"

			_, err = dbSess.InsertInto("page_status").
				Columns("title", "URL", "genre", "media").
				Record(newPS).
				Exec()

			fmt.Printf("newPS:%v\n", newPS)

			if err != nil {
				fmt.Println("データーベースに入れらんない")
				fmt.Println(err)
				panic(err)
			}

			pageIDStr, err = dbSess.Select("id").From("page_status").
				Where("url = ?", newPS.URL).
				ReturnString()
			if err != nil {
				panic(err)
			}

			// rURL = "/r/register_page?url=" + newPS.URL
			// strings.Replace(strings.Replace(newPS.URL, ":", "%3A", -1), "/", "%2F", -1)

		} else {
			// URLが入力されたけど既に登録されてる時
			rURL = "/preview_evaluation/" + strconv.Itoa(dbPS.ID)
			indEval.PageID = dbPS.ID
		}
		fmt.Println("URLのチェックはOK")

	}

	if rURL == "" {
		rURL = "/preview_evaluation/" + pageIDStr

	}

	if pageIDStr != "" {
		indEval.PageID, err = strconv.Atoi(pageIDStr)
		if err != nil {
			return err
		}
	}

	typo.Incorrect = strings.Join(incorrNoNullSL, "\n")
	typo.Correct = strings.Join(corrNoNullSL, "\n")
	typo.PageID = indEval.PageID
	typo.EvaluatorID = indEval.EvaluatorID

	// スライスの長さから typo の数を格納する。
	indEval.NumTypo = len(incorrNoNullSL)

	// 評価を DB に格納する
	if indEval.BrowseTime != "" {
		_, err = dbSess.InsertInto("individual_eval").
			Columns("page_id", "evaluator_id", "browse_time",
				"browse_purpose", "description_eval", "goodness_of_fit",
				"device", "visibility", "num_typo").
			Record(indEval).
			Exec()
	} else {
		_, err = dbSess.InsertInto("individual_eval").
			Columns("page_id", "evaluator_id",
				"browse_purpose", "description_eval", "goodness_of_fit",
				"device", "visibility", "num_typo").
			Record(indEval).
			Exec()
	}
	if err != nil {
		fmt.Println("データーベースに入れらんない")
		fmt.Println(err)
		return c.String(http.StatusOK, "あなたはもう既にこのページを評価しているかもしれません。")
	}

	// typo も DB に格納する
	// typo.IndividualEvalNum = indEvalLoad.Num
	_, err = dbSess.InsertInto("typo").
		Columns("page_id", "evaluator_id", "incorrect", "correct").
		Record(typo).
		Exec()

	return c.Redirect(http.StatusSeeOther, rURL)
}

// 文字コード判定
func charDet(s string) (string, error) {
	// sReader := strings.NewReader(s)
	// fmt.Println(sReader)
	// var by byte
	// by, err := sReader.ReadByte()
	// if err != nil {
	// 	return "", err
	// }
	// b := []byte{by}
	b := []byte(s)
	fmt.Println(b)

	d := chardet.NewTextDetector()
	res, err := d.DetectBest(b)
	if err != nil {
		return "", err
	}
	switch res.Charset {
	case "UTF-8":
		return "utf8", nil
	case "Shift_JIS":
		return "SJIS", nil
	case "EUC-JP":
		return "EUC-JP", nil
	case "ISO-2022-JP":
		return "ISO-2022-JP", nil
	default:
		return res.Charset, err
	}
}

// タイトルの取得
func getPageTitle(url string, s *goquery.Selection) string {
	title := s.Find("title").Text()
	if title == "" {
		title = url
	} else {

		// エンコードを確認
		enc, exists := s.Find("meta").Attr("charset") // HTML5
		fmt.Println("HTML5:" + enc)
		// エラーなら
		if !exists {
			var encSL []string
			var meta *goquery.Selection
			for i := 0; ; i++ {
				meta = s.Find("meta").Eq(i)
				enc, exists = meta.Attr("content")
				// エラーじゃなかったら
				if exists {
					enc = strings.ToLower(enc)
					if strings.Contains(enc, "charset=") {
						encSL = strings.SplitAfter(enc, "charset=")
						fmt.Println("split")
						break
					} else {
						fmt.Println("charset= が無い")
						continue
					}
				} else {
					enc = ""
					fmt.Println("content= が無い")
					break
				}
			}
			if len(encSL) >= 2 {
				enc = strings.ToUpper(encSL[1])
			}
			fmt.Println("charset", enc)
		} else {
			enc = strings.ToUpper(enc)
		}

		switch enc {
		case "UTF-8":
			enc = "utf8"
		case "UTF8":
			enc = "utf8"
		case "SHIFT_JIS":
			enc = "SJIS"
		case "X-SJIS":
			enc = "SJIS"
		case "EUC-JP":
			enc = "EUC-JP"
		case "ISO-2022-JP":
			enc = "ISO-2022-JP"
		default:
			fmt.Println("文字コードチェック")
			var err error
			enc, err = charDet(title)
			if err != nil {
				panic(err)
			}
			fmt.Println("文字コードチェック完了" + enc)
			if enc != "utf8" {
				enc = ""
			}
		}

		if enc != "utf8" {
			fmt.Println("文字コード変換" + enc)
			var err2 error
			title, err2 = charsetutil.DecodeString(title, enc)
			if err2 != nil {
				title = url
			}
			fmt.Println(title)
		}
	}
	return title
}

func main() {

	// ここで入れるべき構造体はinterfaceによって必須のメソッドが定義され、持つべき引数が決まっている。GoDoc参照。
	e.Renderer = &AceTemplate{}

	// Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	// Restricted group
	r := e.Group("/r", cookieToHeaderAuthMiddleware)
	// Token によってサインイン状況をチェック（ログインが必須なページ）
	r.Use(middleware.JWT([]byte("oppai")))

	// "/" の時に返すhtml
	e.GET("/", func(c echo.Context) error {
		return signinCheck("search_top", c, nil)
	})

	// テスト環境のみ
	// e.GET("test", func(c echo.Context) error {
	// 	uri := "http://" + host + "/test///"
	// 	if strings.HasPrefix(uri, "http://") {
	// 		uri = strings.TrimPrefix(uri, "https://")
	// 		uri = strings.TrimPrefix(uri, "http://")
	// 		uri = strings.TrimSuffix(uri, "/")
	// 	}

	// 	uri = "https://ja.wikipedia.org/wiki/%E8%A6%8B%E6%B2%BC%E5%8C%BA"
	// 	resp, _ := http.Get(uri)
	// 	up := resp.Header.Get("Last-Modified")

	// 	var title string
	// 	// url, _ = dbSess.Select("email").From("userinfo").
	// 	// 	Where("email = ? OR email = ?", "docomo.ne.jp", "sm_2-7.ryuuta@"+"docomo.ne.jp").
	// 	// 	ReturnString()
	// 	doc, err := goquery.NewDocument(uri)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	doc.Find("head").Each(func(i int, s *goquery.Selection) {
	// 		title = s.Find("title").Text()
	// 	})

	// 	return c.String(http.StatusOK, up)
	// })

	// 検索時に呼び出される
	e.GET("/search", func(c echo.Context) error {
		searchForm := PageValue{
			Query: "",
		}

		// URLクエリパラメータを受け取る
		q := c.QueryParam("q")
		searchForm.Query = q

		return c.Redirect(http.StatusSeeOther, "https://www.google.co.jp/search?q=site%3Awebrepo.nal.ie.u-ryukyu.ac.jp+"+q)
		// return signinCheck("search_result", c, searchForm)
	})

	e.GET("/page_list", func(c echo.Context) error {

		evalForm, _ := getPageStatusItem(c, -1)

		var (
			listPageValue ListPageValue
			dbPS          []PageStatus
			err           error
			genreQ        = c.QueryParam("genre")
			mediaQ        = c.QueryParam("media")
		)

		listPageValue.EvalForm = evalForm

		// DBから指定したジャンルと媒体のページを取得
		if genreQ == "*" && mediaQ == "*" {
			listPageValue.Content = "<p>ジャンルか媒体どちらかを選択して下さい。</p>"
		} else if genreQ == "*" {
			_, err = dbSess.Select("id", "title", "URL",
				"genre", "media", "dead",
				"tag1", "tag2", "tag3", "tag4", "tag5", "tag6", "tag7", "tag8", "tag9", "tag10").
				From("page_status").
				Where("media = ?", mediaQ).Load(&dbPS)
		} else if mediaQ == "*" {
			_, err = dbSess.Select("id", "title", "URL",
				"genre", "media", "dead",
				"tag1", "tag2", "tag3", "tag4", "tag5", "tag6", "tag7", "tag8", "tag9", "tag10").
				From("page_status").
				Where("genre = ?", genreQ).Load(&dbPS)
		} else {
			_, err = dbSess.Select("id", "title", "URL",
				"genre", "media", "dead",
				"tag1", "tag2", "tag3", "tag4", "tag5", "tag6", "tag7", "tag8", "tag9", "tag10").
				From("page_status").
				Where("genre = ? AND media = ?", genreQ, mediaQ).Load(&dbPS)
		}
		if err != nil {
			panic(err)
		}

		if listPageValue.Content == "" {
			if dbPS != nil {

				// dead が 0 以外のものは除外
				var alivePS []PageStatus
				for i := 0; i < len(dbPS); i++ {
					if dbPS[i].Dead == 0 {
						alivePS = append(alivePS, dbPS[i])
					}
				}
				// スライスの要素数からページの件数を取得
				resultNum := len(alivePS)
				listPageValue.Content = fmt.Sprintf(`<p id="result_status">検索結果：%d件</p>`, resultNum)

				listPageValue.PageStatusSlice = alivePS

				// for i, v := range alivePS {

				// 	listPageValue.Content +=
				// 		fmt.Sprintf(
				// 			`
				// 		<div class="page_status">
				// 			<h3>%d： <a href="/preview_evaluation/%d">%s</a>　（<a href="%s">%s</a>）</h3>
				// 			<div class="cate">ジャンル：%s　媒体：%s</div>
				// 			<div class="tag">タグ： %s %s %s %s %s %s %s %s %s %s</div>
				// 			<h4><a href="/r/input_evaluation/%d">評価する</a></h4>
				// 		</div>
				// 	`, i+1, v.ID, template.HTMLEscapeString(v.Title),
				// 			template.HTMLEscapeString(v.URL), template.HTMLEscapeString(v.URL),
				// 			template.HTMLEscapeString(v.Genre), template.HTMLEscapeString(v.Media),
				// 			template.HTMLEscapeString(v.Tag1), template.HTMLEscapeString(v.Tag2),
				// 			template.HTMLEscapeString(v.Tag3), template.HTMLEscapeString(v.Tag4),
				// 			template.HTMLEscapeString(v.Tag5), template.HTMLEscapeString(v.Tag6),
				// 			template.HTMLEscapeString(v.Tag7), template.HTMLEscapeString(v.Tag8),
				// 			template.HTMLEscapeString(v.Tag9), template.HTMLEscapeString(v.Tag10),
				// 			v.ID)
				// }
			} else {
				listPageValue.Content = "<p id=\"result_status\">検索結果：0件</p>"
			}
		}
		return signinCheck("page_list", c, listPageValue)
	})

	e.GET("/search_user_eval_list", func(c echo.Context) error {

		var mypageValue MyPageValue

		mypageValue.UserName = c.QueryParam("username")

		// DB から特定ユーザーの評価を取得
		// 複数の評価データを格納するために構造体のスライスを作成
		evaluatorID, err := dbSess.Select("id").From("userinfo").
			Where("name = ?", mypageValue.UserName).
			ReturnString()

		if err != nil {
			mypageValue.Content = "<div class=\"subject\">評価が見つかりませんでした。</div>"
			return c.Render(http.StatusOK, "search_user_eval_list", mypageValue)
		}

		var individualEval []IndividualEval
		_, err = dbSess.Select("num", "page_id", "evaluator_id", "posted", "browse_time",
			"browse_purpose", "deliberate", "description_eval", "goodness_of_fit",
			"recommend_good", "recommend_bad", "device", "visibility", "num_typo").
			From("individual_eval").
			Where("evaluator_id = ?", evaluatorID).Load(&individualEval)
		if err != nil {
			mypageValue.Content = "<div class=\"subject\">評価が見つかりませんでした。</div>"
			return c.Render(http.StatusOK, "search_user_eval_list", mypageValue)
		}

		if individualEval != nil {
			// for文で回す
			// Ace に入れる構造体に格納
			for i, v := range individualEval {
				mypageValue.Content += makePrevMyEval(i, v)
			}
		}

		return c.Render(http.StatusOK, "search_user_eval_list", mypageValue)
	})

	// テスト環境のみ
	// e.GET("/page_eval", func(c echo.Context) error {
	// 	searchForm := PageValue{
	// 		Query: "",
	// 	}

	// 	return signinCheck("preview_evaluation", c, searchForm)
	// })

	// サインイン方法選択���面
	e.GET("/signin_select", func(c echo.Context) error {
		fmt.Println("signin_select")
		// return c.Render(http.StatusOK, "signin_select", searchForm)
		return c.Redirect(http.StatusSeeOther, "/google_OAuth")
	})

	// パスワードサインインフォーム
	// e.GET("/pass_signin", func(c echo.Context) error {
	// 	return c.Render(http.StatusOK, "pass_signin", searchForm)
	// })

	// Google の認証画面にリダイレクト
	e.GET("/google_OAuth", func(c echo.Context) error {
		oauthStateString := uniuri.New()
		url := googleOauthConfig.AuthCodeURL(oauthStateString)
		return c.Redirect(http.StatusSeeOther, url)
	})

	// Google からのコールバック
	e.GET("/oauth2callback_google", func(c echo.Context) error {
		code := c.FormValue("code")
		token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
		if err != nil {
			panic(err)
		}

		userGoogle := new(googleUser)
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
			userInfoDB userinfo
		)

		// OAuth、キャリアメールが本登録されてるか確認
		_, err = dbSess.Select("id", "email", "name").From("userinfo").
			Where("OAuth_userinfo = ?", userGoogle.Email).
			Load(&userInfoDB)

		sess, _ := session.Get("session", c)
		sess.Values["GoogleMail"] = userGoogle.Email

		// エラーを吐いた == 中身が入ってない場合
		// 本登録されてない時
		if userInfoDB.Email == "" {
			sess.Save(c.Request(), c.Response())

			oauthService = "Google"
			// if strings.HasSuffix(strings.SplitAfter(userGoogle.Email, "@")[1], "ie.u-ryukyu.ac.jp") {
			return c.Redirect(http.StatusFound, "/ie_OAuth_signup")
			// }

			// return c.Redirect(http.StatusFound, "/OAuth_signup")
		}

		// エラーが無い == 登録済み場合
		// リファラーURLがこのサイトのものか確認する
		createJwt(c, userInfoDB.ID, userInfoDB.Email, userInfoDB.Name)
		fmt.Println("登録済み")

		var rURL string
		if sess.Values["refererURL"] != nil {
			rURL = sess.Values["refererURL"].(string)
		} else {
			rURL = "/"
		}
		sess.Values["refererURL"] = nil
		sess.Save(c.Request(), c.Response())

		fmt.Fprintf(os.Stderr, "%s\n", rURL)

		return c.Redirect(http.StatusSeeOther, rURL)

	})

	e.GET("ie_OAuth_signup", func(c echo.Context) error {
		return c.Render(http.StatusOK, "ie_OAuth_signup", nil)
	})

	e.GET("ie_agree_signup", func(c echo.Context) error {

		var userInfoDB userinfo

		sess, _ := session.Get("session", c)
		email := sess.Values["GoogleMail"].(string)
		// 学内のアドレスの場合はアドレスの入力無しでログインさせる
		// 正規のユーザーテーブルに追加
		result, err := dbSess.InsertInto("userinfo").
			Columns("OAuth_service", "OAuth_userinfo", "email").
			Values(oauthService, email, email).
			Exec()
		if err != nil {
			panic(err)
		} else {
			fmt.Fprintf(os.Stderr, "insert userinfo:%s\n", result)
		}

		// OAuth、キャリアメールが本登録されてるか確認
		_, err = dbSess.Select("id", "email", "name").From("userinfo").
			Where("OAuth_userinfo = ?", email).
			Load(&userInfoDB)

		// エラーが無い == 登録済み場合
		// リファラーURLがこのサイトのものか確認する
		createJwt(c, userInfoDB.ID, userInfoDB.Email, userInfoDB.Name)
		fmt.Println("登録済み")

		var rURL string
		if sess.Values["refererURL"] != nil {
			rURL = sess.Values["refererURL"].(string)
		} else {
			rURL = "/"
		}
		sess.Values["refererURL"] = nil
		sess.Save(c.Request(), c.Response())

		fmt.Fprintf(os.Stderr, "%s\n", rURL)

		return c.Redirect(http.StatusSeeOther, rURL)
	})

	// OAuth認証サインアップ（同意）フォーム
	e.GET("/OAuth_signup", func(c echo.Context) error {
		mailForm := PageValue{
			Error: "",
		}
		return c.Render(http.StatusOK, "OAuth_signup", mailForm)
	})

	// 同意後のアドレス確認促進画面
	e.POST("/agree_signup", func(c echo.Context) error {
		mailForm := PageValue{
			Error: "",
		}
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

			if strings.HasSuffix(eDomain, domain[i]) {

				// メールアドレスが登録されてないのでメールと関連付けたURLを発行
				mac := hmac.New(sha256.New, []byte(uniuri.New()))
				mac.Write([]byte(email))
				act := hex.EncodeToString(mac.Sum(nil))

				// リファラーURLがこのサイトのものか確認する
				sess, err := session.Get("session", c)
				if err != nil {
					return err
				}

				var rURL string
				if sess.Values["refererURL"] != nil {
					rURL = sess.Values["refererURL"].(string)
				} else {
					rURL = "/"
				}
				gmail := sess.Values["GoogleMail"].(string)
				sess.Values["refererURL"] = nil
				sess.Values["GoogleMail"] = nil
				sess.Save(c.Request(), c.Response())

				fmt.Fprintf(os.Stderr, "act:%s\n", act)

				t := time.Now()
				tF := t.Format(timeLayout)

				// 一時ユーザーのテーブルにアドレスと認証コード、リファラーURLを一緒に保存
				result, err := dbSess.InsertInto("tmp_user").
					Columns("OAuth_service", "OAuth_userinfo", "act", "email", "referer", "send_time").
					Values(oauthService, gmail, act, email, rURL, tF).
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
				m.SetBody("text/plain",
					"WebRepo☆彡 に登録いただきありがとうございます。\nメールアドレスの確認を行うため、以下のURLへアクセスして下さい。\nなお、このメールの送信から12時間が経過した場合、このURLは無効となるので再度メールアドレスの登録をお願いします。\nhttps://"+host+"/email_check?act="+act)

				d := gomail.Dialer{Host: "smtp.eve.u-ryukyu.ac.jp", Port: 587, Username: "e145771@eve.u-ryukyu.ac.jp", Password: "USER_PASSWORD"}
				if err := d.DialAndSend(m); err != nil {
					panic(err)
				}

				return c.Render(http.StatusOK, "agree_signup", nil)
			}
		}

		mailForm.Error = "ドメイン(@以降)が携帯キャリアのドメイン以外です。登録できません。"
		return c.Render(http.StatusOK, "OAuth_signup", mailForm)
	})

	// メール確認URLへのアクセス時の処理
	e.GET("/email_check", func(c echo.Context) error {
		act := c.QueryParam("act")
		var tmpUser tmpuser
		_, err := dbSess.Select("act", "OAuth_service", "OAuth_userinfo", "email", "referer").From("tmp_user").
			Where("act = ?", act).
			Load(&tmpUser)

		if err != nil {
			panic(err)
		}
		if tmpUser.Act == "" {
			return errors.New("認証コードが違う!\n")
		}

		fmt.Fprintf(os.Stderr, "act: %s\n", tmpUser.Act)

		// t := time.Now()
		// tF := t.Format(timeLayout)
		// fmt.Fprintf(os.Stderr, "time: %s\n", tF)

		// 正規のユーザー��ーブルに追加
		result, err := dbSess.InsertInto("userinfo").
			Columns("OAuth_service", "OAuth_userinfo", "email").
			Values(tmpUser.OAuthService, tmpUser.OAuthUserinfo, tmpUser.Email).
			Exec()
		if err != nil {
			panic(err)
		} else {
			fmt.Fprintf(os.Stderr, "insert userinfo:%s\n", result)
		}

		// 一時������ザーのテーブルから削除
		result, err = dbSess.DeleteFrom("tmp_user").Where("email = ?", tmpUser.Email).Exec()
		if err != nil {
			panic(err)
		} else {
			fmt.Fprintf(os.Stderr, "delete tempuser:%s\n", result)
		}

		var userInfoDB userinfo

		// OAuth、キャリアメールが本登録されてるか確認
		_, err = dbSess.Select("id", "email", "name").From("userinfo").
			Where("OAuth_userinfo = ?", tmpUser.OAuthUserinfo).
			Load(&userInfoDB)

		createJwt(c, userInfoDB.ID, userInfoDB.Email, userInfoDB.Name)
		return c.Redirect(http.StatusSeeOther, tmpUser.Referer)
	})

	// 評価閲覧画面
	e.GET("/preview_evaluation/:id", func(c echo.Context) error {

		var pageValue PrevEvalPageValue

		pageID := c.Param("id")
		pageIDInt, err := strconv.Atoi(pageID)
		if err != nil {
			panic(err)
		}
		fmt.Println("Atoi OK")

		// DB からページ属性を取得
		_, err = dbSess.Select("id", "title", "URL", "register_date", "last_update",
			"admin_user_id", "genre", "media",
			"tag1", "tag2", "tag3", "tag4", "tag5", "tag6", "tag7", "tag8", "tag9", "tag10").
			From("page_status").
			Where("id = ?", pageIDInt).Load(&pageValue)
		if err != nil {
			panic(err)
		}
		fmt.Println(pageValue)

		if pageValue.URL == "" {
			return c.String(http.StatusNotFound, "Not found 404")
		}

		// DB から評価を取得
		// 複数の評価データを格納するために構造体のスライスを作成
		var individualEval []IndividualEval
		_, err = dbSess.Select("num", "page_id", "evaluator_id", "posted", "browse_time",
			"browse_purpose", "deliberate", "description_eval", "goodness_of_fit",
			"recommend_good", "recommend_bad", "device", "visibility", "num_typo").
			From("individual_eval").
			Where("page_id = ?", pageIDInt).Load(&individualEval)

		fmt.Println("PS OK")
		fmt.Println(individualEval)

		var (
			gfp       int
			visp      int
			enableNum int
			visNum    int
		)
		// 平均評価を計算
		for _, v := range individualEval {
			// 審議なしか審議済みなら
			if v.Deliberate <= 1 {
				gfp += v.GoodnessOfFit
				if v.Visibility >= 1 {
					visp += v.Visibility
					visNum++
				}
				enableNum++
			}
		}
		// 10倍して四捨五入
		gfpf := math.Floor((float64(gfp)/float64(enableNum))*math.Pow10(1) + 0.05)
		vispf := math.Floor((float64(visp)/float64(visNum))*math.Pow10(1) + 0.05)
		// 0.1倍して代入
		pageValue.AveGFP = strconv.FormatFloat(gfpf*math.Pow10(-1), 'f', 1, 64)
		pageValue.AveVisP = strconv.FormatFloat(vispf*math.Pow10(-1), 'f', 1, 64)
		// // スターを付ける
		// aveGFPsl := strings.Split(pageValue.AveGFP, ".")
		// aveVisPsl := strings.Split(pageValue.AveVisP, ".")
		// gfp, _ = strconv.Atoi(aveGFPsl[0])
		// visp, _ = strconv.Atoi(aveVisPsl[0])
		// fmt.Println("GFP", gfp)
		// pageValue.AveGFP = pasteStar(gfp) + "." + aveGFPsl[1]
		// pageValue.AveVisP = pasteStar(visp) + "." + aveVisPsl[1]

		if err != nil {
			panic(err)
		} else if individualEval != nil {
			// for��で回す
			// Ace に入れる構造体に格納
			for i, v := range individualEval {
				pageValue.Content += makePrevEval(i, v)
			}
		}

		// return signinCheck("preview_evaluation", c, nil)
		return signinCheck("tmp_preview_evaluation", c, pageValue)
	})

	r.POST("/recommend_eval/:pageID/:num", func(c echo.Context) error {

		var recommendSQL RecommendSQL
		recommendSQL.UpdTable = "individual_eval"
		recommendSQL.IntoTable = "individual_eval_recom"
		recommendSQL.NumColumn = "eval_num"

		return incrementRecommend(c, recommendSQL)
	})

	r.POST("/recommend_comment/:pageID/:num", func(c echo.Context) error {

		var recommendSQL RecommendSQL
		recommendSQL.UpdTable = "individual_eval_comment"
		recommendSQL.IntoTable = "individual_eval_comment_recom"
		recommendSQL.NumColumn = "comment_num"

		return incrementRecommend(c, recommendSQL)
	})

	// 個別評価閲覧画面
	e.GET("/individual_reviews", func(c echo.Context) error {
		return signinCheck("individual_review", c, nil)
	})

	// 通報完了画面
	r.GET("/dangerous_eval/:pageID/:num", func(c echo.Context) error {

		updTable := "individual_eval"
		numColumn := "eval_num"

		return reportDangerous(c, updTable, numColumn)
	})
	r.GET("/dangerous_comment/:pageID/:num", func(c echo.Context) error {

		updTable := "individual_eval_comment"
		numColumn := "comment_num"

		return reportDangerous(c, updTable, numColumn)
	})

	// 利用規約
	e.GET("/term_of_service", func(c echo.Context) error {
		return signinCheck("term_of_service", c, nil)
	})

	// このサイトについて
	e.GET("/about", func(c echo.Context) error {
		return signinCheck("about", c, nil)
	})

	// テスト環境のみ
	// r.GET("/test", func(c echo.Context) error {
	// 	user := c.Get("user").(*jwt.Token)
	// 	claims := user.Claims.(jwt.MapClaims)
	// 	id := int(claims["id"].(float64))
	// 	email := claims["email"].(string)
	// 	return c.String(http.StatusOK, "Welcome "+fmt.Sprint(id)+" "+email+"!")
	// 	// return signinCheckJWT(pagePath{Page: "mypage_top", URL: "/mypage"}, c)
	// })

	// マイページ
	r.GET("/mypage", func(c echo.Context) error {

		var mypageValue MyPageValue

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		mypageValue.UserName = claims["name"].(string)

		return c.Render(http.StatusOK, "mypage_top", mypageValue)
	})

	// ユーザー設定
	r.GET("/user_settings", func(c echo.Context) error {

		var mypageValue MyPageValue

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		mypageValue.UserName = claims["name"].(string)

		return c.Render(http.StatusOK, "user_settings", mypageValue)
	})

	// ユーザー設定送信後の処理
	r.POST("/user_settings", func(c echo.Context) error {

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userID := claims["id"].(float64)

		userName := c.FormValue("user_name")

		_, err := dbSess.Update("userinfo").Set("name", userName).Where("id = ?", userID).Exec()
		if err != nil {
			fmt.Println("Update 出来ない")
			panic(err)
		}

		createJwt(c, int(userID), claims["email"].(string), userName)

		return c.Redirect(http.StatusSeeOther, "mypage")
	})

	// 自分の付けた評価の一覧
	r.GET("/my_eval_list", func(c echo.Context) error {

		var mypageValue MyPageValue

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		mypageValue.UserName = claims["name"].(string)
		evaluatorID := int(claims["id"].(float64))

		// DB から特定ユーザーの評価を取得
		// 複数の評価データを格納するために構造体のスライスを作成
		var individualEval []IndividualEval
		_, err := dbSess.Select("num", "page_id", "evaluator_id", "posted", "browse_time",
			"browse_purpose", "deliberate", "description_eval", "goodness_of_fit",
			"recommend_good", "recommend_bad", "device", "visibility", "num_typo").
			From("individual_eval").
			Where("evaluator_id = ?", evaluatorID).Load(&individualEval)
		if err != nil {
			panic(err)
		}

		if individualEval != nil {
			// for文で回す
			// Ace に入れる構造体に格納
			for i, v := range individualEval {
				mypageValue.Content += makePrevMyEval(i, v)
			}
		}

		return c.Render(http.StatusOK, "my_eval_list", mypageValue)
	})

	r.GET("/sign_out", func(c echo.Context) error {

		sess, _ := session.Get("session", c)

		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   -1,
			HttpOnly: true,
		}
		sess.Save(c.Request(), c.Response())

		return c.Redirect(http.StatusSeeOther, "/")
	})

	// 新規ページ登録画面
	r.GET("/register_page", func(c echo.Context) error {

		evalForm, _ := getPageStatusItem(c, -1)

		url := c.QueryParam("url")
		if url != "" {
			evalForm.URL = url
		}

		return c.Render(http.StatusOK, "register_page", evalForm)
	})
	r.POST("/register_page", func(c echo.Context) error {

		var (
			newPS  PageStatus
			tagArr [10]string
			dbPS   PageStatusTiny
			isSSL  bool
			isUpd  = false
		)

		newPS.URL = c.FormValue("url")
		// URL のプロトコルが https でも http でも無い時は戻る。
		if !strings.HasPrefix(newPS.URL, "https://") && !strings.HasPrefix(newPS.URL, "http://") {
			evalForm, _ := getPageStatusItem(c, -1)
			return c.Render(http.StatusOK, "register_page", evalForm)
		}

		if strings.HasPrefix(newPS.URL, "http://") {
			isSSL = false
		} else {
			isSSL = true
		}

		// 末尾のスラッシュを削除
		newPS.URL = strings.TrimSuffix(newPS.URL, "/")

		// https:// も http:// も取り除いた変数を用意
		uri := strings.TrimPrefix(newPS.URL, "https://")
		uri = strings.TrimPrefix(uri, "http://")

		_, err := dbSess.Select("id", "URL").From("page_status").
			Where("url = ? OR url = ?", "http://"+uri, "https://"+uri).
			Load(&dbPS)
		if err != nil {
			panic(err)
		}

		// 同じ URI が既に登録されてる時
		if dbPS.URL != "" {
			// DBの方が http で且つ、新規の URL が https の時はアップデートフラグを立てる
			if strings.HasPrefix(dbPS.URL, "http://") && isSSL {
				isUpd = true
			} else {
				// 新規のURLもDBも http なら評価閲覧画面にリダイレクト
				// DBの方が https なら評価閲覧画面にリダイレクト
				return c.Redirect(http.StatusSeeOther, "../preview_evaluation/"+strconv.Itoa(dbPS.ID))
			}
		}

		fmt.Println("URLのチェックはOK")

		newPS.Genre = c.FormValue("genre")
		newPS.Media = c.FormValue("media")
		tag := strings.Split(c.FormValue("tag"), "\n")

		// structVal := reflect.Indirect(reflect.ValueOf(newPS))
		// structVal.Field(i? + 9).Set(v)

		// スライスは tag が入力された個数までしか作られないので、入力された分を配列にコピーする。
		for i, v := range tag {
			tagArr[i] = v
			if i >= 9 {
				break
			}
			// 	structVal.Field(i + 9).Set(v)
		}

		for _, v := range tagArr {
			fmt.Println(v)
			// 	structVal.Field(i + 9).Set(v)
		}

		newPS.Tag1 = tagArr[0]
		newPS.Tag2 = tagArr[1]
		newPS.Tag3 = tagArr[2]
		newPS.Tag4 = tagArr[3]
		newPS.Tag5 = tagArr[4]
		newPS.Tag6 = tagArr[5]
		newPS.Tag7 = tagArr[6]
		newPS.Tag8 = tagArr[7]
		newPS.Tag9 = tagArr[8]
		newPS.Tag10 = tagArr[9]
		// fmt.Printf("tag10:%s\n", structVal.Field())
		fmt.Println("tag10:", newPS.Tag9)

		// newPS.RegisterDate = time.Now().Format(timeLayout)

		fmt.Println(newPS.RegisterDate)

		resp, err := http.Get(newPS.URL)
		if resp == nil && err != nil {
			fmt.Println("レスポンスエラー")
			resp2, err2 := http.Get(newPS.URL + "/")
			// ちゃんとレスポンスが返ってこない（URLがおかしい）時は戻る。
			if err2 != nil {
				evalForm, _ := getPageStatusItem(c, -1)
				return c.Render(http.StatusOK, "register_page", evalForm)
			}
			resp = resp2
		}
		defer resp.Body.Close()
		newPS.Dead = 0

		// ヘッダーの更新日時をパース
		mod, err := time.Parse(http.TimeFormat, resp.Header.Get("Last-Modified"))
		if err != nil {
			fmt.Println("time型に出来ない")
		} else {
			newPS.LastUpdate = mod.Format(timeLayout)
			fmt.Println(newPS.LastUpdate)
		}

		// ページタイトルを取得
		doc, err := goquery.NewDocumentFromResponse(resp)
		if err != nil {
			panic(err)
		}

		doc.Find("head").Each(func(i int, s *goquery.Selection) {
			newPS.Title = getPageTitle(newPS.URL, s)
		})

		// 登録しようとしてる URL が https で、既に登録されてる URL が http だったら置き換えてリダイレクト
		if isUpd {
			// Struct を Map に変換
			structVal := reflect.Indirect(reflect.ValueOf(newPS))
			typ := structVal.Type()

			mapVal := make(map[string]interface{})
			fmt.Println("tag:", typ.Field(1).Tag.Get("db"))
			fmt.Println("value:", structVal.Field(1).Interface())
			fmt.Println("len:", typ.NumField())

			// IDもマップに含めると更新する時に0が入ってしまうので入れない。
			for i := 1; i < typ.NumField(); i++ {
				field := structVal.Field(i)
				value := field.Interface()
				if value != "" {
					mapVal[typ.Field(i).Tag.Get("db")] = field.Interface()
				}
			}

			fmt.Println("map:", mapVal)
			_, err = dbSess.Update("page_status").SetMap(mapVal).Where("id = ?", dbPS.ID).Exec()
			if err != nil {
				fmt.Println("Update 出来ない")
				panic(err)
			}
			fmt.Println("Update!")
			return c.Redirect(http.StatusSeeOther, "../preview_evaluation/"+strconv.Itoa(dbPS.ID))
		}

		if newPS.LastUpdate == "" {
			_, err = dbSess.InsertInto("page_status").
				Columns("title", "URL",
					"admin_user_id", "genre", "media",
					"tag1", "tag2", "tag3", "tag4", "tag5", "tag6", "tag7", "tag8", "tag9", "tag10").
				Record(newPS).
				Exec()
		} else {
			_, err = dbSess.InsertInto("page_status").
				Columns("title", "URL", "last_update",
					"admin_user_id", "genre", "media",
					"tag1", "tag2", "tag3", "tag4", "tag5", "tag6", "tag7", "tag8", "tag9", "tag10").
				Record(newPS).
				Exec()
		}

		fmt.Printf("newPS:%v\n", newPS)

		if err != nil {
			fmt.Println("データーベースに入れらんない")
			fmt.Println(err)
			panic(err)
		}

		// 個別評価テーブルを作成する

		// 個別評価のコメントテーブルを作成する

		id, err := dbSess.Select("id").From("page_status").
			Where("url = ?", newPS.URL).
			ReturnString()
		if err != nil {
			panic(err)
		}
		return c.Redirect(http.StatusSeeOther, "input_evaluation/"+id)
	})

	// ページ属性編集画面
	r.GET("/edit_page_cate/:id", func(c echo.Context) error {
		// v := reflect.Indirect()

		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return err
		}

		evalForm, _ := getPageStatusItem(c, idInt)

		return signinCheck("edit_page_cate", c, evalForm)
	})

	r.POST("/edit_page_cate/:id", func(c echo.Context) error {

		var (
			newPS  PageStatus
			tagArr [10]string
			dbPS   PageStatusTiny
		)

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(err)
		}

		_, err = dbSess.Select("id", "URL").From("page_status").
			Where("id = ?", id).
			Load(&dbPS)
		if err != nil {
			panic(err)
		}

		newPS.Genre = c.FormValue("genre")
		newPS.Media = c.FormValue("media")
		// tag := bytes.Split(c.FormValue("tag"), []byte{13, '\n'})
		tag := strings.Split(c.FormValue("tag"), "\n")

		// structVal := reflect.Indirect(reflect.ValueOf(newPS))
		// structVal.Field(i? + 9).Set(v)

		// スライスは tag が入力された個数までしか作られないので、入力された分を配列にコピーする。
		i := 0
		for j, v := range tag {
			// スライスの中身が無くなったら、その後のタグは[]byte{13}で埋める。
			if len(tag)-1-j == 0 {
				for n := i; n <= 9; n++ {
					// byteスライスからstringへのキャストもメモリコピーが発生してしまう。
					// だが、Sprintf()でも発生するらしいのでもう無理。
					tagArr[n] = byte13Str
					fmt.Println(n)
					fmt.Println(*(*[]byte)(unsafe.Pointer(&tagArr[n])))
				}
			} else if reflect.DeepEqual(*(*[]byte)(unsafe.Pointer(&v)), []byte{13}) {
				// 何故か何も表示されないにも関わらず、Goのスライスでは空白文字を入れると空白ではなく、[]byte型で13が入ってしまう謎の挙動をする。
				// https://qiita.com/Sheile/items/ba51ac9091e09927b95c
				// []byteはまともにチェックするとメモリコピーが発生してしまうので、unsafeを使用。
				// reflect.DeepEqual を使わないと簡単に比較できない。
			} else {
				tagArr[i] = v
				if i >= 9 {
					break
				}
				i++
				fmt.Print(v + "\n")
			}
			fmt.Println(i)
			fmt.Println(*(*[]byte)(unsafe.Pointer(&v)))
		}
		// structVal.Field(i + 9).Set(v)

		newPS.Tag1 = tagArr[0]
		newPS.Tag2 = tagArr[1]
		newPS.Tag3 = tagArr[2]
		newPS.Tag4 = tagArr[3]
		newPS.Tag5 = tagArr[4]
		newPS.Tag6 = tagArr[5]
		newPS.Tag7 = tagArr[6]
		newPS.Tag8 = tagArr[7]
		newPS.Tag9 = tagArr[8]
		newPS.Tag10 = tagArr[9]
		// fmt.Printf("tag10:%s\n", structVal.Field())

		fmt.Println("URL:", dbPS.URL)
		// ち�������とレスポンスが返���てこない時は死亡フラグを��てる
		resp, err := http.Get(dbPS.URL)
		if err != nil {
			newPS.Dead = 1
		} else {
			defer resp.Body.Close()
		}
		fmt.Println("dead:", newPS.Dead)

		// 情報を更新
		if newPS.Dead == 0 {
			// ヘッダーの更新日時をパース
			mod, err := time.Parse(http.TimeFormat, resp.Header.Get("Last-Modified"))
			if err != nil {
				fmt.Println("time型に出来ない")
			} else {
				newPS.LastUpdate = mod.Format(timeLayout)
				fmt.Println(newPS.LastUpdate)
			}

			// ページタイトルを取得
			doc, err := goquery.NewDocumentFromResponse(resp)
			if err != nil {
				panic(err)
			}
			doc.Find("head").Each(func(i int, s *goquery.Selection) {
				newPS.Title = getPageTitle(dbPS.URL, s)
			})
		}

		// Struct を Map に変換
		structVal := reflect.Indirect(reflect.ValueOf(newPS))
		typ := structVal.Type()
		mapVal := make(map[string]interface{})
		// IDもマップに含めると更新する時に0が入ってしまうので入れない。
		for i := 1; i < typ.NumField(); i++ {
			field := structVal.Field(i)
			if field.String() != "" {
				mapVal[typ.Field(i).Tag.Get("db")] = field.Interface()
				fmt.Println("tag:", typ.Field(i).Tag.Get("db"))
			}
		}

		fmt.Println("id", string(dbPS.ID))
		// アップデート
		_, err = dbSess.Update("page_status").SetMap(mapVal).Where("id = ?", dbPS.ID).Exec()
		if err != nil {
			fmt.Println("Update 出来ない")
			panic(err)
		}
		fmt.Println("Update!")
		return c.Redirect(http.StatusSeeOther, "/preview_evaluation/"+strconv.Itoa(dbPS.ID))
	})

	// 評価入力画面
	r.GET("/input_evaluation", func(c echo.Context) error {
		return c.Render(http.StatusOK, "input_evaluation_url", nil)
	})
	r.POST("/input_evaluation", inputEval)

	r.GET("/input_evaluation/:id", func(c echo.Context) error {
		return c.Render(http.StatusOK, "input_evaluation", nil)
	})
	r.POST("/input_evaluation/:id", inputEval)

	// コメント入力画面
	r.GET("/input_comment/:pageID/:evalNum/:num", func(c echo.Context) error {

		return c.Render(http.StatusOK, "input_comment", nil)
	})
	r.POST("/input_comment/:pageID/:evalNum/:num", insertComment)

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
