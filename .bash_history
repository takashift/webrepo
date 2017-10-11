exit
docker-compose logs
cd volumes/go_app/
ls
docker exec -it tuna_mysql_1 /bin/bash
docker exec -it tuna_nginx_1 /bin/bash
docker exec -it tuna_nginx_1 /bin/ash
docker exec -it tuna_nginx_1 /bin/bash
docker exec -it tuna_mysql_1 /bin/bash
cd ..
cd db_data/
ls
rm -r *
ls
sudo rm -r *
ls
git commit -am "Go での DB の Update、Delete テスト用のコードを記述"
gitpush 
sudo rm -r *
ls
docker exec -it tuna_mysql_1 /bin/bash
rcode server.go
cd volumes/go_app/
rcode server.go
curl -X PUT -H 'Content-Type: application/json' -d '{"id":2, "email": "unk@email.com"}' https://webrepo.nal.ie.u-ryukyu.ac.jp/users/
curl --cacert -X PUT -H 'Content-Type: application/json' -d '{"id":2, "email": "unk@email.com"}' https://webrepo.nal.ie.u-ryukyu.ac.jp/users/
curl --cacert -X PUT -H 'Content-Type: application/json' -d '{"id":2, "email": "unk@email.com"}' tuna_echo_1:3306/users/
curl --cacert -X PUT -H 'Content-Type: application/json' -d '{"id":2, "email": "unk@email.com"}' tuna_nginx_1:3306/users/
curl --cacert -X PUT -H 'Content-Type: application/json' -d '{"id":2, "email": "unk@email.com"}' tuna_nginx_1:443/users/
docker-compose logs
docker-compose dow
docker-compose down
docker-compose up -d
docker-compose up -d --build
docker-compose logs
curl -X PUT -H 'Content-Type: application/json' -d '{"id":3, "email":"oppai@email.com"}' -k https://tuna_nginx_1/users/
docker-compose down
docker-compose up -d --build
docker-compose down --rmi all
docker-compose up -d --build
cd echo/
s
ls
cat Dockerfile
cd
cat docker-compose.yml
cd volumes/go_app/
rcode server.go
docker-compose logs
docker exec -it tuna_mysql_1 /bin/bash
ls
cd mysql/
ls
cd init/
ls
cat webrepo.sql 
rcode webrepo.sql 
cd ..
cd volumes/go_app/
cd
cd volumes/db_data/
ls
sudo rm -r *
ls
ks
ls
docker exec -it tuna_mysql_1 /bin/bash
cd 
cd mysql/
cd init/
rcode webrepo.sql 
ls
cd
cd volumes/go_app/
ls
cd volumes/go_app/
ls
cd ..
cd db_data/
ls
sudo rm -r *
ls
docker exec -it tuna_mysql_1 /bin/bash
cd volumes/go_app/
rcode server.go
docker-compose logs
docker-compose down --rmi all
docker-compose up -d --build
docker-compose logs
rcode server.go
docker-compose logs
docker-compose down --rmi all
docker-compose up -d --build
docker-compose logs
rcode server.go
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
git status
git commit -am "取り敢えずHTTPリファラーwo"
git commit --amend
gitpush 
cd volumes/go_app/
rcode server.go
docker-compose logs
git commit -am 'Google のユーザー情報リクエスト先をv2からv3へ変更'
docker-compose logs
git status
git log
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
git commit -am 'リファラーURLを出力するように変更'
gitpush 
git commit --amend
gitpush 
gitpush -f
git push -f
git commit -am 'コールバックの処理で、メールアドレスが登録されていたら元居たページにリダイレクトするように変更'
gitpush 
cd volumes/go_app/views/
ls

rcode OAuth_signup.ace 
rcode search_
rcode search_top.ace 
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
git commit -am 'メール確認画面で Post で送られた情報を取ってこれるように追記'
gitpush 
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose up -d --build
docker-compose down
docker-compose logs
docker-compose up -d --build
docker-compose logs
git commit -am 'キャリアメールのドメインを集めた配列を追加'
gitpq
gitpush 
git lo
git log
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose logs
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose logs
docker-compose down
docker-compose logs
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose up -d --build
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
git commit -am '登録されてるキャリアメールドメインのアドレス以外は弾くようにコードを記述'
gitpush 
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
git log
cd volumes/go_app/
rcode server.go
cd volumes/go_app/
rcode server.go
cd volumes/go_app/
rcode server.go
docker-compose logs
cd ..
rcode docker-compose.yml
cd echo/
rcode Dockerfile
cd ..
cd nginx/
rcode Dockerfile
cd ..
cd echo/
rcode Dockerfile
git commit -am "メール送信のためのポートを開放する設定を docker-compose と Echo の Dockerfile に追加"
git status
git commit --amend
cd ..

git commit --amend
git status
git commit --amend
git status
git add
git add docker-compose.yml
git commit --amend
cd echo/
rcode Dockerfile
git commit --amend
git add Dockerfile
git commit --amend
cd ..
ls
cd volumes/go_app/
rcode server.go 
docker-compose logs
cd ..
cd
cd echo/
rcode Dockerfile
docker-compose down --rmi all
docker-compose up -d --build
docker-compose down --rmi all
cd
rcode docker-compose.yml
docker-compose down --rmi all
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
cd volumes/db_data/
sudo rm -r *
ls
sudo rm -r *
ls
docker-compose down
docker-compose down --rmi all
sudo rm -r *
docker-compose up -d --build
sudo rm -r *
ls
docker-compose logs
cd ..
cd volumes/go_app/archive/
ls
rcode db.go.cm 
docker-compose logs
docker-compose down
docker-compose down --rmi all
docker-compose up -d --build
docker-compose logs
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker exec -it tuna_nginx_1 /bin/bash
docker exec -it tuna_mysql_1 /bin/bash
cd mysql/
cd init/
ls
rcode webrepo.sql 
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose logs
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
git commit -am 'メールに関連付けた認証コードをデータベースに保存するように変更'
git commit --amend
gitpush 
rcode webrepo.sql 
git commit -am 'docker-compose のポートの設定が間違っていたので修正。'
gitpush 
docker exec -it tuna_nginx_1 /bin/bash
docker exec -it tuna_mysql_1 /bin/bash
docker-compose down --rmi all
docker-compose up -d --build
docker-compose logs
docker exec -it tuna_mysql_1 /bin/bash
docker-compose logs
docker-compose down mysql
docker-compose down
docker rmi mysql
docker rmi tuna_mysql
docker-compose up -d --build
docker-compose down
docker rmi tuna_mysql
docker-compose up -d --build
docker exec -it tuna_mysql_1 /bin/bash
docker-compose logs
docker exec -it tuna_mysql_1 /bin/bash
cd ..
cd volumes/db_data/
ls
sudo rm -r *
git log
git status
git commit -am "ドメイン検証のif文内に書かれていなければならない処理を外に書いていたので修正。メール内の確認URLにアクセスした時に認証コードがDBに存在するか確認する処理を追加"
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose logs
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose logs
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose down
docker-compose logs
docker-compose up -d --build
docker-compose logs
docker-compose down
docker rmi echo
docker rmi tuna_echo_1
docker rmi tuna_echo
cd echo/
rcode Dockerfile
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
git commit sutatus
git commit status
git status
docker-compose down --rmi all
docker-compose logs
docker-compose up -d --build
git log
docker-compose lodg
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
esit
exit
cd volumes/db_data/
sudo rm -r *
exit
cd volumes/db_data/
ls
cd ..
cd go_app/
rcode server.go 
docker exec -it tuna_mysql_1 /bin/bash
cd
ls
cd volumes/
ks
ls
cd web-vol/
ls
cd ..
cd
cd volumes/db_data/
sudo rm -r *
ls
docker exec -it tuna_mysql_1 /bin/bash
cd 
rcode docker-compose.yml
cd volumes/db_data/
sudo rm -r *
rcode volumes/go_app/
devolumes/go_app/
cd volumes/go_app/
rcode server.go 
docker-compose down --rmi all
docker-compose up -d --build
docker-compose logs
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker rmi tuna_mysql
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose logs
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose logs
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose up -d --build
docker exec -it tuna_mysql_1 /bin/bash
cd volumes/db_data/
sudo rm -r *
cd
rcode dock
cd echo/
rcode Dockerfile
git commit -am "大学総情センターのメールを使って、アドレス確認メールの送信に成功"
gitpush 
docker exec -it tuna_mysql_1 /bin/bash
cd mysql/
cd init/
rcode webrepo.sql 
docker exec -it tuna_mysql_1 /bin/bash
cd
cd volumes/db_data/
ls
sudo rm -r *
docker exec -it tuna_mysql_1 /bin/bash
git commit -am "本登録を行う処理を追加"
gitpush 
docker exec -it tuna_mysql_1 /bin/bash
sudo rm -r *
ls
sudo rm -r *
ls
sudo rm -r *
docker exec -it tuna_mysql_1 /bin/bash
cd volumes/db_data/
cd ..
cd..
cd ..
cd volumes/go_app/
;s
ls
vi server.go 
rcode server.go 
vi server.go 
git commit -am "既に本登録されているアドレスが入力された時は前の画面に戻るように変更"
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
cd volumes/go_app/
rcode server.go 
docker-compose logs
git commit -am "既に本登録してあるメールアドレスを入力した時に弾く処理を追加"
gitpush 
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose up -d --build
docker-compose down
docker-compose up -d --build
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose logs
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose logs
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
cd ..
git commit -am "認証コードURLにアクセス時に一時ユーザーから削除の上、本登録を行う処理を追加"
gitpush 
git log
docker-compose down
docker rmi tuna_mysql
docker-compose up -d --build
docker-compose down
docker rmi tuna_mysql
docker-compose up -d --build
docker-compose logs
cd volumes/db_data/
cd ..
cd go_app/
rcode server.go 
git commit -am "送信メールアドレスのドメインにwebrepo.を追加した。"
gip
gitp
gitpush 
git commit --amend
git add server.go 
git commit --amend
git log
cd mysql/
rcode Dockerfile 
cd ../go/
rcode Dockerfile 
ls
cd ../echo/
ls
rcode Dockerfile 
git commit -am "Mysql の Dockerfile に apt-get update を追加"
git commit --amend
gitpush 
cd
rcode docker-compose.yml
mkdir cron
cd cron/
rcode Dockerfile
cd ../echo/
rcode Dockerfile
cd ../cron/
rcode Dockerfile
cd ../echo/
rcode Dockerfile
cd ../cron/
docker build
docker images
docker build cron
docker build -t cron .
docker run cron
docker ps
docker run -d  --name cron cron
docker ps
docker ps -a
docker rmi cron
docker ps -a
docker rmi 76da55c8019d
docker rmi 804e826b61c4
docker down cron
docker stop cron
docker stop fervent_cray
docker images
docker run -d  --name cron alpine
docker images
docker ps -a
docker rmi 76da55c8019d
docker rmi 75e637fcf723
docker stop 75e637fcf723
docker stop 
docker stop cbe55a9b4d26
docker images
docker ps -a
docker rm cbe55a9b4d26
docker rm 75e637fcf723
docker ps -a
docker run -it  --name cron alpine
docker exec -it tuna_mysql_1 /bin/bash
cd volumes/db_data/
ls
cd
cd volumes/go_app/
rcode server.go 
exit
docker run -it --name cron alpine
docker rm corn
docker stop corn
docker ps -a
docker stop cron
docker rm cron
docker ps -a
docker run -it --name cron alpine
docker rm cron
docker ps -a
cd cron/
rcode Dockerfile 
xit
exit
cd cron/
rcode Dockerfile 
rcode Dockerfile cd ..
cd
rcode docker-compose.yml
cd cron/
rcode Dockerfile 
exit
cd volumes/go_app/
rcode server.go 
cd
cd exit
exot
exit
cd volumes/go_app/
rcode server.go 
exit
rcode docker-compose.yml
exit
rcode docker-compose.yml
exit
cd volumes/go_app/
rcode docker-compose.yml
exit
rcode docker-compose.yml
cd volumes/go_app/
rcode server.go 
git log
cd
cd mysql/
rcode Dockerfile 
rcode docker-compose.yml
git status
cd cron/
ls
git add Dockerfile 
git status
rcode Dockerfile 
git commit -am "やっぱり Docker の使い方として1コンテナに1プロセスなので、apt-get update は消してcron用の Dockerfile を作成。"
gitpush 
gitpush -f
git push -f
docker log
docker logs
docker logs cron
rcode docker-compose.yml
ls -l /var/run/docker.sock
cd volumes/go_app/
ls
cd views/
ls
cd ..
rcode server.go 
cd ../
cd ..
cd mysql/init/
rcode webrepo.sql 
docker-compose down
cd ../../cron/
cd ../volumes/db_data/
sudo rm -r *
ls
docker-compose up -d --build
docker-compose down
docker-compose up -d --build
docker-compose down
docker-compose up -d --build
docker-compose down
docker-compose up -d --build
docker-compose down
docker-compose up -d --build
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose logs
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose logs
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
$ docker run -it alpine sh -c "echo '* * * * * echo foobar' > /var/spool/cron/crontabs/root && crond -l 2 -f"
docker run -it alpine sh -c "echo '* * * * * echo foobar' > /var/spool/cron/crontabs/root && crond -l 2 -f"
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose logs
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose logs
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose up -d --build
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose logs
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
docker-compose up -d --build
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose down
docker-compose up -d --build
docker-compose logs
docker exec -it tuna_cron_1 /bin/ash
git status
cd cron/
rcode Dockerfile 
docker exec -it tuna_crom_1 /bin/ash
docker exec -it tuna_cron_1 /bin/ash
docker exec -it cron /bin/ash
docker -ps
docker ps -a
docker exec -it tuna_cron_1 /bin/ash
docker ps -a
docker exec -it tuna_cron_1 /bin/ash
docker ps -a
docker exec -it tuna_cron_1 /bin/ash
docker ps -a
docker-compose logs
cd ..
cd go/
rcode D
cd
cd echo/
rcode Dockerfile
docker ps -a
echo `*/5 * * * * mysql -h mysql -u rtuna -pUSER_PASSWORD Webrepo -e 'DELETE FROM tmp_user WHERE send_time < CURRENT_TIMESTAMP - INTERVAL 12 HOUR'`
docker ps -a
rcode Dockerfile
docker run -it alpine
docker ps -a
docker run -it alpine
docker ps -a
docker exec -it tuna_cron_1 /bin/ash
docker ps -a
docker exec -it tuna_cron_1 /bin/ash
docker exec -it tuna_mysql_1 /bin/ash
docker exec -it tuna_mysql_1 /bin/bash
git commit -am "１２時間が経過した一時ユーザーを削除するコンテナを作成"
git status
git commit --amend
gitpush 
cd cron/
rcode Dockerfile 
docker-compose logs
docker-compose down
docker-compose up -d --build
docker exec -it tuna_cron_1 /bin/ash
docker-compose down
docker-compose up -d --build
docker exec -it tuna_cron_1 /bin/ash
cd
rcode docker-compose.yml
docker exec -it tuna_cron_1 /bin/ash
docker-compose up -d --build
docker exec -it tuna_cron_1 /bin/ash
docker-compose up -d --build
docker exec -it tuna_cron_1 /bin/ash
cd volumes/go_app/
ls
cd views
cd ..
rcode server.go 
git commit -am "一時ユーザーの構造体にもSQLに合わせてsend_timeを追加"
gitpush 
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
ls
rcode template_no_proxy.go 
cd views/
rcode search_
rcode search_result.ace 
rcode search_top.ace 
rcode OAuth_signup.ace 
docker-compose logs
docker-compose down
docker-compose up -d --build
docker exec -it tuna_mysql_1 /bin/bash
rcode docker-compose.yml
xit
exit
exit
rcode docker-compose.yml
docker exec -it tuna_cron_1 /bin/ash
git log
git commit -am "トランザクションが最低限になるようにtmp_userで指定しているカラムを修正"
gito
gitpush 
docker-compose logs
docker-compose down
docker-compose up -d --build
docker-compose logs
git commit -am "メールアドレス入力時にエラーの内容を表示するようにした。"
gitpush 
git status
docker-compose logs
git commit -am "Google以外の認証方法は一旦保留にするので、ログイン方法選択画面を Google の認証へリダイレクトするように変更"
gitpush 
git status
git commit --amend
gitpush 
git push -f
git commit -am "パスワードサインインのページは使わないので攻撃防止のためにコメントアウト。"
gitpush 
cd volumes/go_app/
ls
cd views/
ls
rcode signin_select.ace 
rcode OAuth_signup.ace 
rcode pass_signin.ace 
rcode agree_signup.ace 
rcode dengerous_complete.ace 
git status
rcode signin_select.ace 
git status
git commit -am "検索エンジンに引っ掛けたくないページに設定を追加"
gitpush 
rcode OAuth_signup.ace 
git log
rcode docker-compose.yml
cd volumes/go_app/
ls
rcode server.go 
