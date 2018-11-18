FROM nginx:alpine
#FROM alpine:latest

MAINTAINER hgcr

#ADD ./www /root/www
#ADD ./app.conf /etc/nginx/conf.d/app.conf
#ADD ./app.conf.noSSL /etc/nginx/nginx.conf
ADD ./app.conf /etc/nginx/nginx.conf
ADD ./usock.conf /etc/tmpfiles.d/usock.conf

#RUN apk --no-cache add nginx

RUN apk --no-cache add shadow
RUN mkdir /root/logs
RUN chmod 755 -R /root
RUN usermod -u 1000 nginx \
    && groupmod -g 1000 nginx
#RUN chmod 660 /web/domain.sock

#RUN apk --no-cache add certbot
#ADD https://dl.eff.org/certbot-auto /usr/local/bin/certbot-auto
#RUN chmod a+x /usr/local/bin/certbot-auto

# 80番ポートは恐らく初めから開放してるので、設定不要
EXPOSE 443
# 80

#RUN nginx
#RUN certbot certonly --standalone -d webrepo.japaneast.cloudapp.azure.com -m e145771@ie.u-ryukyu.ac.jp --agree-tos -n
#RUN mkdir /acme
#RUN certbot certonly --webroot -w /acme/ -d webrepo.japaneast.cloudapp.azure.com -m e145771@ie.u-ryukyu.ac.jp --agree-tos -n
#RUN nginx -s stop

#CMD ["nginx", "-g", "daemon off;"]
#CMD [ "certbot", "certonly", "--standalone", "-d", "webrepo.japaneast.cloudapp.azure.com", "-m", "e145771@ie.u-ryukyu.ac.jp", "--agree-tos", "-n" ]
