chmod 400 server.key 
ls
ls -l
git commit -am "証明書を本番環境のものに置き換え"
gitpush 
cd vg
cd
cd volumes/go_app/
ls
cd views/
ls
cp OAuth_signup.ace ie_OAuth_signup.ace 
rcode ie_OAuth_signup.ace 
git commit -am "ie生でも同意画面を表示するように変更"
gitpush 
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose down
docker-compose up -d --build
git stutus
git status
git add ie_OAuth_signup.ace 
git commit -am "ie生用の同意画面をaddし忘れていたので追加"
gitpush 
ls
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
git commit -am "ieのアドレスなら、キャリアのアドレス無しに登録できるように変更"
gitpush 
rcode nginx/Dockerfile 
rcode nginx/app.conf 
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
cd nginx/s
cd nginx/
ls
cd keys/
ls
mv server.crt server.cer
touch server.crt
cat server.cer >> server.crt
cat nii-odcacssha1.cer >> server.crt
ls -l
chmod 400 server.crt nii-odcacssha1.cer
ls -l
git status
git add server.cer nii-odcacssha1.cer 
git commit -am "中間証明書と合成したファイルを追加"
gitpush 
git fetch origin
git reset --hard origin/master
cd 
cd volumes/go_app/
ls
rcode server.go 
cd volumes/go_app/
ls
rcode server.go 
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
cd
cd nginx/keys/
ls
mv nii-odcacssha1.cer nii-odca3sha1.cer
cat nii-odca3sha1.cer 
ls
ls -l
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
git status
git add nii-odca3sha1.cer 
git commit -am "文字コードを指定していないutf8でタイトルを取ってこれないようになっていたので修正"
gitpush 
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
cd
cd volumes/go_app/
ls
rcode c
rcode createPrevEvalPage.go 
docker stop tuna_echo_1
docker-compose up -d --build
git status
git commit -am '評価一覧のタグ間のスペースを修正'
gitpush 
cd
cd nginx/
;s
ls
rcode app.conf 
rcode ~/docker-compose.yml
cd
cd volumes/web-vol/
ls
rcode googledc7e8e91c55f39d7.html
docker stop tuna_nginx_1
docker-compose up -d --build
ls
cd ..
cd go_app/views/
rcode search_top.ace 
cd
cd volumes/go_app/
rcode server.go 
git statis
git status
git commit -am "メタタグの部分にGoogle認証用の記述を追加。HTML5の文字コード指定の書き方に対応"
docker stop tuna_nginx_1
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
gitpush 
docker-compose logs
rcode server.go 
rcode createPrevEvalPage.go 
git status
 git commit -am "平均値の計算を修正"
gitp
gitpush 
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
git status
git commit -am "charsetの含まれるメタタグの属性名contentが２つ目以降の場合もエラーが起こらないようにに修正"
git commit --amend
gitpush 
docker-compose logs
cd 
cd .
cd nginx/keys/
ls
vi server.crt
cat server.cer nii-odca3sha2.cer >> server.crt
cat server.crt 
git commit -am "中間証明書のアルゴリズムの種類が間違ってるっぽかったので修正"
gitpush 
exi
exit
cd nginx/keys/
ls
vi nii-odca3sha1.cer 
chmod 700 nii-odca3sha1.cer 
vi nii-odca3sha1.cer 
chmod 400 nii-odca3sha1.cer 
rm server.crt 
ls
touch server.crt
cat server.cer nii-odca3sha1.cer >> server.crt 
cat server.crt 
git commit -am "間違ったファイルを合成していたので週背"
git commit --amend
gitpush 
docker exec -it tuna_mysql_1 /bin/bash
git status
ls
cd volumes/go_app/
ls
rcode server.go 
cd views/
ls
rcode font-family.ace 
rcode search_top.ace 
git log
git sutatua
git sutats
git status
git add ../../../nginx/keys/nii-odca3sha2.cer
git commit -am "font-family の指定の前の方に游ゴシック体ミディアムを追加"
gitpush 
ls
rcode search_top.ace 
rcode search_result.ace 
rcode page_list.ace 
rcode p
rcode preview_evaluation.ace 
rcode t
rcode tmp_preview_evaluation.ace 
git commit -am "WebRepo★彡のフォント指定からsan-selfを削除"
gitpush 
git commit -am "トップページのタイトルを変更"
gitpush 
git commit -am "トップページのタイトルを変更"
gitpush 
cd
cd bin/
rcode _gitpulldep
ls -l
chmod 700 _gitpulldep 
ls -l
git commit -am "本番環境での pull を自動化するコマンドを作成"
git commit --amend
git status
git add _gitpulldep 
git status
git commit -am "本番環境での pull を自動化するコマンドを作成"
gitpush 
git push -f
git log
docker stop tuna_echo_1
docker-compose logs
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
git status
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
git status
git commit -am "評価閲覧ページにTwitterのシェアボタンを追加"
gitpush 
ls
cd volumes/go_app/views/
ls
rcode server.go
rcode ../server.go
rcode OAuth_signup.ace 
git commit -am ".ac.jpのアドレスをキャリアメールの代わりに使えるように変更"
gitpush 
docker stop tuna_echo_1
docker-compose up -d --build
docker exec -it tuna_mysql_1 /bin/bash
cd volumes/go_app/
rcode OAuth_signup.ace 
rcode server.go
docker-compose logs
cd volumes/go_app
rcode server.go 
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
git commit -am "リダイレクトした時に404エラーが表示される問題を修正"
gitpush 
git status
cd views/
ls
rcode font-family.ace 
git commit -am "Yu Go Medi の指定が間違ってたので修正"
gitpush 
cd volumes/go_app/
rcode server.go
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
git status
git commit -am "ie生で同意ボタンを押した後に501が出るかもしれない問題を修正"
gitpush 
docker-compose logs
cd volumes/go_app/
rcode server.go 
rcode createPrevEvalPage.go 
git status
git commit -am "ページ一覧でtemplate.HTMLEscapeString()するように変更"
gitpush 
git commit -am "評価閲覧ページでtemplate.HTMLEscapeString()するように変更"
gitpush 
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
git commit -am "評価閲覧ページで"\n"を"<br>"に変換するように変更"
git commit -am '評価閲覧ページで"\n"を"<br>"に変換するように変更'
gitpush 
git commit -am ''
gitpush 
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
git commit -am '閲覧目的も"\'
gitpush 
cd views/
rcode input_evaluation
rcode input_evaluation.ace 
rcode input_evaluation_url.ace 
git commit -am '記述評価と閲覧日時を任意と明記'
gitpush 
git commit -am '任意の位置が微妙だったので改行を追加'
gitpush 
rcode about.ace 
rcode mypage_top.ace 
git commit -am "マイページのタイトルがこのサイトについてだったので修正"
gitpo
gitpush 
rcode cd echo/
cd echo/
rcode Dockerfile
mysql -u rtuna -pUSER_PASSWORD Webrepo
docker exec -it tuna_mysql_1 /bin/bash
cd volumes/go_app/
ls
cd views/
ls
rcode fo
rcode font-family.ace 
git commit -am "FireFox で Yu Go Medi を表示する設定を追加"
gitpush 
rcode search_
cd
cd volumes/go_app/
ls
rcode server.go 
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
git commit -am "X-SJISの場合の正常な文字コード変換とcharsetが一つ目の属性にない場合に２番目以降を探す処理を追加"
gitpush 
docker-compose logs
gitpush 
git log
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
git commit -am "charsetが大文字で書かれてる場合に対応"
gitpush 
docker stop tuna_echo_1
docker-compose up -d --build
git log
docker-compose logs
git commit -am "charset=が含まれているかを明確に確認するように変更"
gitpush 
git commit -am "charsetが見つからなかった時はcontinueするように修正"
gitp
gitpush 
git commit -am "metaタグが取って来れない場合もnilにならないみたいなので判定をcontent属性のある無しに変更"
gitpush 
git commit -am "HTML5の時に文字コードを大文字にする処理が無かったので追加"
gitpush 
cd volumes/go_app/views/
l;s
ls
rcode about.ace 
rcode mypage_top.ace 
rcode p
rcode page_list.ace 
rcode OAuth_signup.ace 
rcode input_evaluation
rcode input_evaluation.ace 
rcode input_evaluation_url.ace 
rcode OAuth_signup.ace 
rcode agree_signup.ace 
pass
rcode pass_signin.ace 
rcode preview_evaluation.ace 
rcode consent_form.ace 
rcode register_page.ace 
rcode dangerous_complete.ace 
rcode search_result.ace 
rcode edit_page_cate.ace 
rcode search_top.ace 
rcode font-family.ace 
rcode signin_select.ace 
rcode footer.ace 
rcode term_of_service.ace 
rcode header_menu.ace 
rcode ie_OAuth_signup.ace 
rcode individual_review.ace 
rcode tmp_preview_evaluation.ace 
rcode input_comment.ace 
rcode search_top.ace 
git status
git commit -am "pe-jino "
git commit --amend
gitpush
ls
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
git status
ls
git status
gitpush 
git commit -am "トップページのタイトルを変更"
gitpush 
git status
ls
rcode common-setting.ace
git status
git add common-setting.ace 
rcode input_evaluation.ace 
rcode input_evaluation_url.ace 
rcode OAuth_signup.ace 
rcode mypage_top.ace 
rcode about.ace 
rcode page_list.ace 
rcode agree_signup.ace 
pass_
rcode pass_signin.ace 
rcode preview_evaluation.ace 
rcode consent_form.ace 
rcode register_page.ace 
rcode dangerous_complete.ace 
rcode search_result.ace 
rcode edit_page_cate.ace 
rcode search_top.ace 
rcode font-family.ace 
rcode signin_select.ace 
rcode term_of_service.ace 
rcode header_menu.ace 
rcode test.ace 
rcode ie_OAuth_signup.ace 
rcode individual_review.ace 
rcode tmp_preview_evaluation.ace 
rcode input_
rcode input_comment.ace 
rcode search_top.ace 
git commit -am "各ページにファビコンを設定"
gitpush 
cd 
cd echo/Dockerfile
rcode echo/Dockerfile
cd
cd nginx/app.conf 
rcode nginx/app.conf 
rcode nginx/Dockerfile 
rcode docker-compose.yml
cd volumes/go_app/
ls
cd 
cd nginx/
ls
cd www
ls
cd ..
cd 
cd volumes/
ls
web-vol/
ls
cd web-vol/
ls
mkdir logs
mkdir nginx-logs
ls -l
ls -l ..
cd ..
mkdir nginx_log
ls
cd web-vol/
cd volumes/go_app/
ls
rcode server.go 
cd ..
ls
rm nginx_log/
rm -r nginx_log/
ls
cd
cd echo/
/\ls
ls
rcode Dockerfile
cd
rcode docker-compose.yml
cd volumes/
ls
cd
cd nginx/
ls
rcode app.conf 
cd
cd volumes/go_app/views/
rcode common-setting.ace 
cd
ls
mv Twitter3.1（仮）.png favicon.png
ls
mv favicon.png volumes/web-vol/
chown tuna volumes/web-vol/
sudo chown tuna volumes/web-vol/
ls volumes/web-vol/
ls -l volumes/
mv favicon.png volumes/web-vol/
cd volumes/web-vol/
mkdir image
mv favicon.png image/
docker stop tuna_echo_1
docker stop tuna_nginx_1
docker-compose up -d --build
ls
docker-compose logs
ls
docker stop tuna_nginx_1
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_nginx_1
docker-compose up -d --build
docker stop tuna_nginx_1
docker-compose up -d --build
ls volumes/web-vol/
ls
cd
mv Twitter3.1拡大.png favicon.png
mv favicon.png image/
ls
mv favicon.png volumes/web-vol/image/
cd volumes/web-vol/image/
ls
git status
cd
git status
git add volumes/web-vol/
git commit -am 'ファビコンを設定'
gitpush 
ls
rm Twitter3.1拡大.png 
rcode volumes/go_app/
ls volumes/go_app/views/
ls

cd  volumes/go_app/views/
cd volumes/go_app/views/
ls
rcode input_evaluation.ace 
rcode input_evaluation_url.ace 
git stauts
git stautus
git status
git commit -am "目的達成度の説明を変更"
gitpush 
docker stop tuna_echo_1
docker-compose up -d --build
rcode header_menu.ace 
git stautus
git stauts
git status
git commit -am "FAQへのリンクを追加"
gitpush 
docker stop tuna_echo_1
docker-compose up -d --build
ls
rcode page_list.ace 
rcode search_top.ace 
cd ..
rcode server.go 
rcode JWT.ace
cd views/
rcode JWT.ace
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
rcode JWT.js 
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
git status
git commit -am 'ツイートを取得するjavaScriptコードをトップページに記述（まだエラー）'
gitpush 
cd volumes/go_app/ls
cd volumes/go_app/
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
git status
git commit -am 'トップページにサイトの説明の追加、見た目を改良'
cd volumes/go_app/views/
ls
rcode mypage_top.ace 
rcode search_top.ace 
rcode footer.ace 
rcode search_top.ace 
rcode mypage_top.ace 
rcode user_config.ace
rcode page_list.ace 
cd volumes/go_app/
ls;
ls
git status
gitpush 
git status
docker stop tuna_echo_1
docker-compose up -d --build
git status
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
git status
git commit -am "トップページでブックマークを推奨"
gitpush 
git status
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
rcode server.go
rcode ../../mysql/init/webrepo.sql 
git commit -am "ユーザー名をTokenに含めるように修せお"
git commit -amend
git commit --amend
docker stop tuna_echo_1
docker-compose up -d --build
git commit --amend
git status
git add server.go 
git commit --amend
git status
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
rcode server.go
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose logs
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
git status
git commit -am "マイページにユーザー名を表示するように変更"
gitpush 
cd volumes/go_app/
ls
rcode server.go 
cd 
ls
cd bin/
ls
cat gitpush 
cd
ls .ssh/
cat .gitconfig 
cd volumes/go_app/
rcode server.go 
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
git commit -am "ユーザー設定のページを作成"
git status
git commit -am "ユーザー設定のページを作成"
gitpush 
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
rcode ~/mysql/init/webrepo.sql 
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose logs
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose up -d --build
docker-compose logs
docker stop tuna_echo_1
docker-compose logs
docker-compose up -d --build
docker stop tuna_echo_1
docker-compose up -d --build
git commit -am "ユーザー名の更新機能を実装"
git status
gitpush 
cd volumes/go_app/views/
ls
rcode user_config.ace
rcode mypage_top.ace 
rcode input_comment.ace 
mv user_config.ace user_settings.ace 
rcode user_settings.ace 
git commit -am "マイページにユーザー設定へのリンクを追加。ユーザー設定のページの作成"
gitpush 
rcode eval_form_css.ace
rcode common-setting.ace 
rcode input_evaluation.ace 
rcode search_top.ace 
git add eval_form_css.ace user_settings.ace 
cd volumes/go_app/views/
ls
rcode my_eval_list.ace
rcode mypage_top.ace 
rcode user_settings.ace 
rcode search_result.ace 
rcode tmp_preview_evaluation.ace 
rcode page_list.ace 
git status
git add my_eval_list.ace 
git commit -am "ユーザーの付けた評価の一覧ページを用意"
gitpush 
cd volumes/go_app/views/
ls
rcode mypage_top.ace 
docker stop tuna_echo_1
docker-compose up -d --build
git commit -am "マイページに自分の付けた評価一覧へのリンクを用意"
gitpush 
git log
git status
cd volumes/go_app/
ls
rcode server.go 
