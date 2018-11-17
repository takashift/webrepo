FROM alpine:latest

RUN apk --no-cache add mysql-client
RUN apk --no-cache add certbot
RUN certbot certonly --webroot -w /acme/ -d webrepo.japaneast.cloudapp.azure.com -m e145771@ie.u-ryukyu.ac.jp --agree-tos -n
#RUN certbot certonly --standalone -d webrepo.japaneast.cloudapp.azure.com -m e145771@ie.u-ryukyu.ac.jp --agree-tos -n
#RUN certbot certonly --standalone --standalone-supported-challenges http-01 -d webrepo.japaneast.cloudapp.azure.com -m e145771@ie.u-ryukyu.ac.jp --agree-tos -n
RUN echo '*/5 * * * * mysql -u rtuna -p$MYSQL_PASSWORD -S /usock/mysqld.sock Webrepo -e "DELETE FROM tmp_user WHERE send_time < CURRENT_TIMESTAMP - INTERVAL 12 HOUR;"' > /etc/crontabs/root
RUN echo '0 0 * * 1 mysqldump -u rtuna -p$MYSQL_PASSWORD -S /usock/mysqld.sock Webrepo > /db_backup/DBWebrepoBuckupMonday.sql' >> /etc/crontabs/root
RUN echo '0 0 * * 4 mysqldump -u rtuna -p$MYSQL_PASSWORD -S /usock/mysqld.sock Webrepo > /db_backup/DBWebrepoBuckupThursday.sql' >> /etc/crontabs/root
RUN echo '0 6 1 * * mysqldump -u rtuna -p$MYSQL_PASSWORD -S /usock/mysqld.sock Webrepo > /db_backup/DBWebrepoBuckup1th.sql' >> /etc/crontabs/root
RUN echo '0 18 15 * * mysqldump -u rtuna -p$MYSQL_PASSWORD -S /usock/mysqld.sock Webrepo > /db_backup/DBWebrepoBuckup15th.sql' >> /etc/crontabs/root
RUN echo '0 12 */1 * * mysqldump -u rtuna -p$MYSQL_PASSWORD -S /usock/mysqld.sock Webrepo > /db_backup/DBWebrepoBuckupThursday.sql.`date "+%Y%m%d_%H%M%S"`' >> /etc/crontabs/root
RUN echo '0 5 1 * * certbot renew' >> /etc/crontabs/root

# 登録から12時間を経過している一時ユーザーは削除
CMD [ "crond", "-l", "2", "-f" ]
