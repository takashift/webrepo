package main

import (
	"fmt"
	"html/template"
	"strconv"
	"strings"
)

// type (
// 		PrevIndiEval struct {
// 			BrowsePurpose   string
// 			EvaluatorName   string
// 			BrowseTime      string
// 			GoodnessOfFit   string
// 			Visibility      string
// 			NumTypo         string
// 			Incorrect       string
// 			Correct         string
// 			DescriptionEval string
// 			Posted          string
// 			EvalNum         string
// 			RecommendGood   string
// 			RecommendBad    string
// 			NumComment      string
// 		}

// 		PrevEvalComment struct {
// 			CommenterName string
// 			Comment       string
// 			Posted        string
// 			ReplyEvalNum  string
// 			CommentNum    string
// 			RecommendGood string
// 			RecommendBad  string
//		}

// 	IndividualEval struct {
// 		Num                  int    `db:"num"`
// 		PageID               int    `db:"page_id"`
// 		EvaluatorID          int    `db:"evaluator_id"`
// 		Posted               string `db:"posted"`
// 		BrowseTime           string `db:"browse_time"`
// 		BrowsePurpose        string `db:"browse_purpose"`
// 		Deliberate           int    `db:"deliberate"`
// 		DescriptionEval      string `db:"description_eval"`
// 		RecommendGood        int    `db:"recommend_good"`
// 		RecommendBad         int    `db:"recommend_bad"`
// 		GoodnessOfFit        int    `db:"goodness_of_fit"`
// 		BecauseGoodnessOfFit string `db:"because_goodness_of_fit"`
// 		Device               string `db:"device"`
// 		Visibility           int    `db:"visibility"`
// 		BecauseVisibility    string `db:"because_visibility"`
// 		NumTypo              int    `db:"num_typo"`
// 		BecauseNumTypo       string `db:"because_num_typo"`
// 	}

// 	IndividualEvalComment struct {
// 		Num             int    `db:"num"`
// 		PageID          int    `db:"page_id"`
// 		CommenterID     int    `db:"commenter_id"`
// 		Posted          string `db:"posted"`
// 		ReplyEvalNum    int    `db:"reply_eval_num"`
// 		ReplyCommentNum int    `db:"reply_comment_num"`
// 		Deliberate      int    `db:"deliberate"`
// 		Comment         string `db:"comment"`
// 		RecommendGood   int    `db:"recommend_good"`
// 		RecommendBad    int    `db:"recommend_bad"`
// 	}
// )

var (
	gfpMenu = map[int]string{
		5: "求めていた以上に達成できた",
		4: "完全に達成できた",
		3: "ほぼ達成できた",
		2: "あまり達成できなかった",
		1: "全然達成できなかった",
	}
	vispMenu = map[int]string{
		5: "極めて見やすい",
		4: "そこそこ見やすい",
		3: "見づらくはない",
		2: "そこそこ見づらい",
		1: "極めて見づらい",
	}
)

func makePrevEval(iEval int, eval IndividualEval) string {

	// 審議中なら""を返す
	if eval.Deliberate >= 2 {
		return ""
	}

	fmt.Println(eval.EvaluatorID)

	fmt.Println("評価者名を取得")
	// DB から評価者名を取得
	evaluatorName, err := dbSess.Select("name").From("userinfo").
		Where("id = ?", eval.EvaluatorID).
		ReturnString()
	if err != nil {
		panic(err)
	}

	fmt.Println(evaluatorName)

	// DB から誤字脱字を取得
	typo := new(Typo)
	dbSess.Select("incorrect", "correct").
		From("typo").
		Where("individual_eval_num = ?", eval.Num).Load(&typo)

	// 閲覧日がデフォルト値のときは修正
	if eval.BrowseTime == "0001-01-01 01:01:01" {
		eval.BrowseTime = "不明"
	}

	// 単なる改行区切りなので、スライスに再解凍
	var (
		incorrect, correct, typoEndTag string
	)
	incorrSL := strings.Split(typo.Incorrect, "\n")
	corrSL := strings.Split(typo.Correct, "\n")

	// 誤字脱字の数だけ必要なHTMLタグもセットで生成
	if incorrSL[0] == "" {
		incorrSL[0] = "無し"
	} else {
		incorrect =
			`<div class="typo">
			<div class="incorrect">
				<h3>✕ 誤</h3>
				<div class="typo_list">`
		for _, v := range incorrSL {
			incorrect += "<h4>" + template.HTMLEscapeString(v) + "</h4>"
		}
		incorrect +=
			`	</div>
		</div>`
		typoEndTag = "</div>"
	}
	if corrSL[0] == "" {
		corrSL[0] = "無し"
	} else {
		incorrect =
			`<div class="typo">
			<div class="correct">
			<h3>⭕ 正</h3>
			<div class="typo_list">`
		for _, v := range corrSL {
			correct += "<h4>" + template.HTMLEscapeString(v) + "</h4>"
		}
		incorrect +=
			`	</div>
		</div>`
	}

	// DB からコメントを取得
	var individualEvalCommentRaw []IndividualEvalComment
	_, _ = dbSess.Select("num", "page_id", "commenter_id", "posted",
		"reply_eval_num", "reply_comment_num", "deliberate", "comment",
		"recommend_good", "recommend_bad").
		From("individual_eval_comment").
		Where("reply_eval_num = ?", eval.Num).Load(&individualEvalCommentRaw)

	// スライスの要素数からコメントの数を取得
	// deliberate が 2 以上のものは数えない
	var individualEvalComment []IndividualEvalComment
	for i := 0; i < len(individualEvalCommentRaw); i++ {
		if individualEvalCommentRaw[i].Deliberate <= 1 {
			individualEvalComment = append(individualEvalComment, individualEvalCommentRaw[i])
		}
	}
	numComment := len(individualEvalComment)

	iEval++
	result := fmt.Sprintf(
		`<div class="review">
		<h3>No.%d　　%s</h3>
		<p class="author">評価者　%s</p>
		<p class="date">閲覧日　%s</p>
		<h4 class="first">目的達成度　%s</h4>
		<h4>見やすさ　　%s（%s）</h4>
		<h4>誤字脱字数　%d箇所</h4>
					%s
					%s
					%s
		<h4>記述評価</h4>
		<div class="doc">
			<h4>%s</h4>
		</div>
		<div class="res">
			<span id="posted">投稿日　%s　</span>
			<span>参考に...
				<form class="recommend" name="評価" method="post" action="/r/recommend_eval/%d/%d">
					<input type="submit" value="なった👍" name="recommend"> %d
					<input type="submit" value="ならなかった👎" name="recommend"> %d
				</form>
			</span>
		</div>
		<form class="res_button" method="get" tprevet="_blank">
			<div class="input_dangerous">
				<input type="submit" formaction="/r/dangerous_eval/%d/%d" value="通報する" name="dangerous">
			</div>
			<div class="input_comment">
				<input type="submit" formaction="/r/input_comment/%d/%d/%d" value="コメントする" name="comment">
			</div>
		</form>
	</div>
	
	<h3>コメント(%d件)</h3>
	`, iEval, strings.Replace(template.HTMLEscapeString(eval.BrowsePurpose), "\n", "<br>", -1),
		template.HTMLEscapeString(evaluatorName),
		eval.BrowseTime, pasteStar(eval.GoodnessOfFit, gfpMenu),
		pasteStar(eval.Visibility, vispMenu), setDevice(eval.Device), eval.NumTypo,
		incorrect, correct, typoEndTag,
		strings.Replace(template.HTMLEscapeString(eval.DescriptionEval), "\n", "<br>", -1),
		eval.Posted, eval.PageID, eval.Num, eval.RecommendGood, eval.RecommendBad,
		eval.PageID, eval.Num, eval.PageID, eval.Num, 0, numComment)

	fmt.Println("評価を取ってくるのはOK")

	pageEvalCommentNumMap := map[int]int{}
	// コメントのテンプレートを追加
	for j, v := range individualEvalComment {
		result += makePrevEvalComment(v, iEval, j, pageEvalCommentNumMap)
	}

	return result
}

// 自分のページの評価を表示用（通報、GoodBad、コメントボタン、コメント無し。）
func makePrevMyEval(iEval int, eval IndividualEval) string {

	// 審議中なら""を返す
	if eval.Deliberate >= 2 {
		return ""
	}

	fmt.Println(eval.Num)

	// DB から評価ページのタイトル、タグ、媒体を取得
	var pageStatus PageStatus
	_, err := dbSess.Select("title", "genre", "media",
		"tag1", "tag2", "tag3", "tag4", "tag5",
		"tag6", "tag7", "tag8", "tag9", "tag10").
		From("page_status").
		Where("id = ?", eval.PageID).
		Load(&pageStatus)
	if err != nil {
		fmt.Println("タイトルの取得に失敗")
		panic(err)
	}

	// DB から誤字脱字を取得
	typo := new(Typo)
	dbSess.Select("incorrect", "correct").
		From("typo").
		Where("individual_eval_num = ?", eval.Num).Load(&typo)

	// 閲覧日がデフォルト値のときは修正
	if eval.BrowseTime == "0001-01-01 01:01:01" {
		eval.BrowseTime = "不明"
	}

	// 単なる改行区切りなので、スライスに再解凍
	var (
		incorrect, correct, typoEndTag string
	)
	incorrSL := strings.Split(typo.Incorrect, "\n")
	corrSL := strings.Split(typo.Correct, "\n")

	// 誤字脱字の数だけ必要なHTMLタグもセットで生成
	if incorrSL[0] == "" {
		incorrSL[0] = "無し"
	} else {
		incorrect =
			`<div class="typo">
			<div class="incorrect">
				<h3>✕ 誤</h3>
				<div class="typo_list">`
		for _, v := range incorrSL {
			incorrect += "<h4>" + template.HTMLEscapeString(v) + "</h4>"
		}
		incorrect +=
			`	</div>
		</div>`
		typoEndTag = "</div>"
	}
	if corrSL[0] == "" {
		corrSL[0] = "無し"
	} else {
		incorrect =
			`<div class="typo">
			<div class="correct">
			<h3>⭕ 正</h3>
			<div class="typo_list">`
		for _, v := range corrSL {
			correct += "<h4>" + template.HTMLEscapeString(v) + "</h4>"
		}
		incorrect +=
			`	</div>
		</div>`
	}

	// DB からコメントを取得
	var individualEvalCommentRaw []IndividualEvalComment
	_, _ = dbSess.Select("num", "page_id", "commenter_id", "posted",
		"reply_eval_num", "reply_comment_num", "deliberate").
		From("individual_eval_comment").
		Where("reply_eval_num = ?", eval.Num).Load(&individualEvalCommentRaw)

	// スライスの要素数からコメントの数を取得
	// deliberate が 2 以上のものは数えない
	var individualEvalComment []IndividualEvalComment
	for i := 0; i < len(individualEvalCommentRaw); i++ {
		if individualEvalCommentRaw[i].Deliberate <= 1 {
			individualEvalComment = append(individualEvalComment, individualEvalCommentRaw[i])
		}
	}
	numComment := len(individualEvalComment)

	iEval++
	result := fmt.Sprintf(
		`<div class="review">
		<h3 class="page_tilte">No.%d　　<a href="/preview_evaluation/%d">%s</a></h3>
		<div class="cate">ジャンル：%s　媒体：%s</div>
		<div class="tag">タグ：%s %s %s %s %s %s %s %s %s %s</div>
		<h3>目的：%s</h3>
		<p class="date">閲覧日　%s</p>
		<h4 class="first">目的達成度　%s</h4>
		<h4>見やすさ　　%s（%s）</h4>
		<h4>誤字脱字数　%d箇所</h4>
					%s
					%s
					%s
		<h4>記述評価</h4>
		<div class="doc">
			<h4>%s</h4>
		</div>
		<div class="res">
			<span id="posted">投稿日　%s　　コメント(%d件)　　参考に... なった👍%d　ならなかった👎%d</span>
		</div>
	</div>
	`, iEval, pageStatus.ID, template.HTMLEscapeString(pageStatus.Title),
		template.HTMLEscapeString(pageStatus.Genre), template.HTMLEscapeString(pageStatus.Media),
		template.HTMLEscapeString(pageStatus.Tag1), template.HTMLEscapeString(pageStatus.Tag2),
		template.HTMLEscapeString(pageStatus.Tag3), template.HTMLEscapeString(pageStatus.Tag4),
		template.HTMLEscapeString(pageStatus.Tag5), template.HTMLEscapeString(pageStatus.Tag6),
		template.HTMLEscapeString(pageStatus.Tag7), template.HTMLEscapeString(pageStatus.Tag8),
		template.HTMLEscapeString(pageStatus.Tag9), template.HTMLEscapeString(pageStatus.Tag10),
		strings.Replace(template.HTMLEscapeString(eval.BrowsePurpose), "\n", "<br>", -1),
		eval.BrowseTime, pasteStar(eval.GoodnessOfFit, gfpMenu),
		pasteStar(eval.Visibility, vispMenu), setDevice(eval.Device), eval.NumTypo,
		incorrect, correct, typoEndTag,
		strings.Replace(template.HTMLEscapeString(eval.DescriptionEval), "\n", "<br>", -1),
		eval.Posted, numComment, eval.RecommendGood, eval.RecommendBad)

	fmt.Println("評価を取ってくるのはOK")

	return result
}

func makePrevEvalComment(comment IndividualEvalComment, i int, j int, pageEvalCommentNumMap map[int]int) string {

	// // 審議中なら""を返す
	// if comment.Deliberate >= 2 {
	// 	return ""
	// }
	j++
	pageEvalCommentNumMap[comment.Num] = j

	// DB から投稿者名を取得
	commenterName, _ := dbSess.Select("name").From("userinfo").
		Where("id = ?", comment.CommenterID).
		ReturnString()
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("Sprintfの前")

	return fmt.Sprintf(
		`<div class="comment">
		<div class="review">
			<p class="author">投稿者　%s</p>
			<h4>>>%s　%s</h4>
			<div class="res">
				<span class="posted">No.%d　　投稿日　%s　</span>
				<span>参考に...
					<form class="recommend" name="評価のコメント" method="post" action="/r/recommend_comment/%d/%d">
						<input type="submit" value="なった👍" name="recommend"> %d
						<input type="submit" value="ならなかった👎" name="recommend"> %d</span>
					</form>
			</div>
			<form class="res_button" action method="get" tprevet="_blank">
				<div class="input_dangerous">
					<input type="submit" formaction="/r/dangerous_comment/%d/%d" value="通報する" name="dangerous">
				</div>
				<div class="input_comment">
					<input type="submit" formaction="/r/input_comment/%d/%d/%d" value="コメントする" name="comment">
				</div>
			</form>
		</div>
	</div>
	`, template.HTMLEscapeString(commenterName), toEval(i, comment, pageEvalCommentNumMap),
		strings.Replace(template.HTMLEscapeString(comment.Comment), "\n", "<br>", -1), j, comment.Posted,
		comment.PageID, comment.Num,
		comment.RecommendGood, comment.RecommendBad,
		comment.PageID, comment.Num,
		comment.PageID, comment.ReplyEvalNum, comment.Num)
}

func toEval(i int, arg IndividualEvalComment, numMap map[int]int) string {
	var value string
	num := strconv.Itoa(numMap[arg.ReplyCommentNum])
	fmt.Println("toEval")

	if num == "0" {
		value = "評価" + strconv.Itoa(i)
	} else {
		value = num
	}
	return value
}

func pasteStar(i int, m map[int]string) string {
	var result string
	if i == 5 {
		result = "<span class=\"star\">★★★★★</span> 5　" + m[i]
	}
	if i == 4 {
		result = "<span class=\"star\">★★★★</span>　 4　" + m[i]
	}
	if i == 3 {
		result = "<span class=\"star\">★★★</span>　　 3　" + m[i]
	}
	if i == 2 {
		result = "<span class=\"star\">★★</span>　　　 2　" + m[i]
	}
	if i == 1 {
		result = "<span class=\"star\">★</span>　　　　 1　" + m[i]
	}
	return result
}

func setDevice(s string) string {
	if s == "SP" {
		s = "スマートフォン・タブレット端末"
	} else {
		s = "パソコン"
	}
	return s
}
