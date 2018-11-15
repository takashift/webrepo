FROM nginx:alpine

MAINTAINER hgcr

#ADD ./www /root/www
#ADD ./app.conf /etc/nginx/conf.d/app.conf
ADD ./app.conf /etc/nginx/nginx.conf
ADD ./keys/fullchain.pem /keys/fullchain.pem
ADD ./keys/privkey.pem /keys/privkey.pem
ADD ./usock.conf /etc/tmpfiles.d/usock.conf

# Let's Encrypt の　設定
RUN apk --no-cache add certbot
# RUN mkdir /acme/
# RUN certbot certonly --webroot -w /web/ -d webrepo.japaneast.cloudapp.azure.com -m e145771@ie.u-ryukyu.ac.jp --agree-tos -n
RUN certbot certonly --standalone -d webrepo.japaneast.cloudapp.azure.com -m e145771@ie.u-ryukyu.ac.jp --agree-tos -n

RUN apk --no-cache add shadow
RUN mkdir /root/logs
RUN chmod 755 -R /root
RUN chmod 400 /keys/fullchain.pem /keys/privkey.pem
RUN chown root:root /keys/fullchain.pem /keys/privkey.pem
RUN usermod -u 1000 nginx \
    && groupmod -g 1000 nginx
#RUN chmod 660 /web/domain.sock

# 80番ポートは恐らく初めから開放してるので、設定不要
EXPOSE 443
CMD ["nginx", "-g", "daemon off;"]
