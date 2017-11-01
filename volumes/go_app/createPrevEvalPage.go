package main

import "fmt"

func makePrevEval(arg PrevIndiEval) string {
	return fmt.Sprintf(
		`<div class="review">
		<h3>%s</h3>
		<p class="author">投稿者　%s</p>
		<p class="date">閲覧日　%s</p>
		<h4 class="first">目的達成度　★★★★★ %s</h4>
		<h4>見やすさ　　★★★★★ %s</h4>
		<h4>誤字脱字数　%s箇所</h4>
		<div class="typo">
			<div class="incorrect">
				<h3>✕ 誤</h3>
				<div class="typo_list">
					%s
				</div>
			</div>
			<div class="correct">
				<h3>⭕ 正</h3>
				<div class="typo_list">
					%s
				</div>
			</div>
		</div>
		<h4>記述評価</h4>
		<div class="doc">
			<h4>%s</h4>
		</div>
		<div class="res">
			<span id="posted">投稿日　%s</span>
			<span>参考に...
				<form class="recommend" name="評価%s" method="post" action></form>
				<input type="submit" value="なった👍" name="recommend"> %s
				<input type="submit" value="ならなかった👎" name="recommend"> %s</span>
		</div>
		<form id="res_button" action method="get" target="_blank">
			<div class="input_dengerous">
				<input type="submit" value="通報する" name="dengerous">
			</div>
			<div class="input_comment">
				<input type="submit" value="コメントする" name="comment">
			</div>
		</form>
	</div>
	
	<h3>コメント（%s）</h3>
	`, arg.Title, arg.EvaluatorName, arg.BrowseTime, arg.GoodnessOfFit, arg.Visibility, arg.NumTypo,
		arg.Incorrect, arg.Correct, arg.DescriptionEval, arg.Posted, arg.EvalNum,
		arg.RecommendGood, arg.RecommendBad, arg.NumComment)
}

func makePrevEvalComment(arg PrevEvalComment) string {
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
			<form id="res_button" action method="get" target="_blank">
				<div class="input_dengerous">
					<input type="submit" value="通報する" name="dengerous">
				</div>
				<div class="input_comment">
					<input type="submit" value="コメントする" name="comment">
				</div>
			</form>
		</div>
	</div>
	`, arg.CommenterName, arg.Comment, arg.Posted, arg.ReplyEvalNum, arg.CommentNum,
		arg.RecommendGood, arg.RecommendBad)
}
