FROM mysql:latest

ADD ./init /docker-entrypoint-initdb.d

ENV MYSQL_DATABASE Webrepo
ENV MYSQL_USER rtuna
ENV MYSQL_PASSWORD USER_PASSWORD
ENV MYSQL_ROOT_PASSWORD ROOT_PASSWORD

CMD ["mysqld", "--character-set-server=utf8mb4", "--collation-server=utf8mb4_unicode_ci"]
