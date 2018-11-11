FROM mysql:latest

ADD ./init /docker-entrypoint-initdb.d

ENV MYSQL_DATABASE Webrepo
ENV MYSQL_USER rtuna
#ENV MYSQL_PASSWORD USER_PASSWORD
#ENV MYSQL_ROOT_PASSWORD ROOT_PASSWORD

RUN { \
    echo '[mysqld]'; \
    echo 'character-set-server=utf8mb4'; \
    echo 'collation-server=utf8mb4_general_ci'; \
    echo 'max_allowed_packet=100MB'; \
    echo 'explicit_defaults_for_timestamp=1'; \
    echo 'innodb_large_prefix=1'; \
    echo 'innodb_file_per_table'; \
    echo 'innodb_file_format=barracuda'; \
    echo '[client]'; \
    echo 'default-character-set=utf8mb4'; \
} > /etc/mysql/conf.d/charset.cnf
