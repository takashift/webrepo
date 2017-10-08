FROM alpine:latest

RUN echo '*/5 * * * * echo "Hello"' >> /etc/crontab/root

CMD ["cron", "-f"]
