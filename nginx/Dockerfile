FROM nginx:alpine

MAINTAINER hgcr

#ADD ./www /root/www
#ADD ./app.conf /etc/nginx/conf.d/app.conf
ADD ./app.conf /etc/nginx/nginx.conf
ADD ./keys/server.crt /etc/nginx/server.crt
ADD ./keys/server.crt /etc/nginx/server.csr
ADD ./keys/server.key /etc/nginx/server.key
ADD ./usock.conf /etc/tmpfiles.d/usock.conf

RUN apk --no-cache add shadow
RUN mkdir /root/logs
RUN chmod 755 -R /root
RUN chmod 400 /etc/nginx/server.crt /etc/nginx/server.csr /etc/nginx/server.key
RUN chown root:root /etc/nginx/server.crt /etc/nginx/server.csr /etc/nginx/server.key
RUN usermod -u 1000 nginx \
    && groupmod -g 1000 nginx
#RUN chmod 660 /web/domain.sock

# 80番ポートは恐らく初めから開放してるので、設定不要
EXPOSE 443
CMD ["nginx", "-g", "daemon off;"]
