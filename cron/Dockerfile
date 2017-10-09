FROM alpine:latest

RUN apk --no-cache add mysql-client
RUN echo '*/5 * * * * mysql -u rtuna -pUSER_PASSWORD -S /usock/mysqld.sock Webrepo -e "DELETE FROM tmp_user WHERE send_time < CURRENT_TIMESTAMP - INTERVAL 12 HOUR;"' > /etc/crontabs/root

# 登録から12時間を経過している一時ユーザーは削除
CMD [ "crond", "-l", "2", "-f" ]
