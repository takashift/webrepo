FROM nginx:alpine

MAINTAINER hgcr

#ADD ./www /root/www
#ADD ./app.conf /etc/nginx/conf.d/app.conf
ADD ./app.conf /etc/nginx/nginx.conf
ADD ./letsencrypt/live/webrepo.japaneast.cloudapp.azure.com/fullchain.pem /keys/fullchain.pem
ADD ./letsencrypt/live/webrepo.japaneast.cloudapp.azure.com/privkey.pem /keys/privkey.pem
ADD ./usock.conf /etc/tmpfiles.d/usock.conf

RUN apk --no-cache add shadow
RUN mkdir /root/logs
RUN chmod 755 -R /root
RUN chmod 400 /keys/fullchain.pem /keys/privkey.pem
RUN chown root:root /keys/fullchain.pem /keys/privkey.pem
RUN usermod -u 1000 nginx \
    && groupmod -g 1000 nginx
#RUN chmod 660 /web/domain.sock

# Let's Encrypt の　設定
EXPOSE 80
RUN nginx
# RUN apk --no-cache add certbot
RUN mkdir /acme/
RUN certbot certonly --webroot -w /acme/ -d webrepo.japaneast.cloudapp.azure.com -m e145771@ie.u-ryukyu.ac.jp --agree-tos -n
# RUN certbot certonly --standalone -d webrepo.japaneast.cloudapp.azure.com -m e145771@ie.u-ryukyu.ac.jp --agree-tos -n
# RUN certbot certonly --standalone --standalone-supported-challenges http-01 -d webrepo.japaneast.cloudapp.azure.com -m e145771@ie.u-ryukyu.ac.jp --agree-tos -n
RUN nginx -s stop
RUN echo '0       5       1       *       *       certbot renew' >> /etc/crontabs/root

# 80番ポートは恐らく初めから開放してるので、設定不要
EXPOSE 443
CMD ["nginx", "-g", "daemon off;"]
