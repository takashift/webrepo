package main

import "fmt"
import "strings"

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
// 		}
// )

func makePrevEval(eval IndividualEval) string {

	// 審議中なら""を返す
	if eval.Deliberate != 0 {
		return ""
	}

	// DB から評価者名を取得
	evaluatorName, err := dbSess.Select("name").From("userinfo").
		Where("id = ?", eval.EvaluatorID).
		ReturnString()
	if err != nil {
		panic(err)
	}

	fmt.Println("evaluatorName")

	// DB から誤字脱字を取得
	typo := new(Typo)
	dbSess.Select("incorrect", "correct").
		From("typo").
		Where("evaluator_id = ?", eval.EvaluatorID).Load(&typo)

	// 単なる改行区切りなので、スライスに再解凍
	var (
		incorrect, correct, typoEnd string
	)
	incorrSL := strings.Split(typo.Incorrect, "\n")
	corrSL := strings.Split(typo.Correct, "\n")

	// 閲覧日がデフォルト値のときは修正
	if eval.BrowseTime == "0001-01-01 01:01:01" {
		eval.BrowseTime = "不明"
	}

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
			incorrect += "<h4>" + v + "</h4>"
		}
		incorrect +=
			`	</div>
		</div>`
		typoEnd = "</div>"
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
			correct += "<h4>" + v + "</h4>"
		}
		incorrect +=
			`	</div>
		</div>`
	}

	// DB からコメントを取得
	var individualEvalComment []IndividualEvalComment
	_, _ = dbSess.Select("num", "page_id", "commenter_id", "posted",
		"reply_eval_num", "reply_comment_num", "deliberate", "comment",
		"recommend_good", "recommend_bad").
		From("individual_eval_comment").
		Where("page_id = ?", eval.PageID).Load(&individualEvalComment)
	// スライスの要素数からコメントの数を取得
	numComment := len(individualEvalComment)

	result := fmt.Sprintf(
		`<div class="review">
		<h3>%s</h3>
		<p class="author">評価者　%s</p>
		<p class="date">閲覧日　%s</p>
		<h4 class="first">目的達成度　★★★★★ %d</h4>
		<h4>見やすさ　　★★★★★ %d</h4>
		<h4>誤字脱字数　%d箇所</h4>
					%s
					%s
					%s
		<h4>記述評価</h4>
		<div class="doc">
			<h4>%s</h4>
		</div>
		<div class="res">
			<span id="posted">投稿日　%s</span>
			<span>参考に...
				<form class="recommend" name="評価%d" method="post" action></form>
				<input type="submit" value="なった👍" name="recommend"> %d
				<input type="submit" value="ならなかった👎" name="recommend"> %d</span>
		</div>
		<form id="res_button" action method="get" tprevet="_blank">
			<div class="input_dengerous">
				<input type="submit" value="通報する" name="dengerous">
			</div>
			<div class="input_comment">
				<input type="submit" value="コメントする" name="comment">
			</div>
		</form>
	</div>
	
	<h3>コメント(%d件)</h3>
	`, eval.BrowsePurpose, evaluatorName, eval.BrowseTime,
		eval.GoodnessOfFit, eval.Visibility, eval.NumTypo,
		incorrect, correct, typoEnd, eval.DescriptionEval,
		eval.Posted, eval.Num,
		eval.RecommendGood, eval.RecommendBad, numComment)

	// コメントのテンプレートを追加
	for _, v := range individualEvalComment {
		result += makePrevEvalComment(v)
	}

	return result
}

func makePrevEvalComment(comment IndividualEvalComment) string {

	// 審議中なら""を返す
	if comment.Deliberate != 0 {
		return ""
	}

	// DB から投稿者名を取得
	commenterName, _ := dbSess.Select("name").From("userinfo").
		Where("id = ?", comment.CommenterID).
		ReturnString()
	// if err != nil {
	// 	panic(err)
	// }

	return fmt.Sprintf(
		`<div class="comment">
		<div class="review">
			<p class="author">投稿者　%s</p>
			<h4>%s</h4>
			<div class="res">
				<span id="posted">投稿日　%s</span>
				<span>参考に...
					<form class="recommend" name="評価%sのコメント%s" method="post" action></form>
					<input type="submit" value="なった👍" name="recommend">%s
					<input type="submit" value="ならなかった👎" name="recommend">%s</span>
			</div>
			<form id="res_button" action method="get" tprevet="_blank">
				<div class="input_dengerous">
					<input type="submit" value="通報する" name="dengerous">
				</div>
				<div class="input_comment">
					<input type="submit" value="コメントする" name="comment">
				</div>
			</form>
		</div>
	</div>
	`, commenterName, comment.Comment, comment.Posted,
		comment.ReplyEvalNum, comment.Num,
		comment.RecommendGood, comment.RecommendBad)
}
