FROM mysql:8.0.17

VOLUME /data/goecmserver

ENV MYSQL_ROOT_PASSWORD=1qaz2WSX 

ADD goecmserver /usr/local/bin/goecmserver
ADD scripts /opt/scripts

RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime


WORKDIR /data/goecmserver

ENTRYPOINT ["bash","/opt/scripts/start-goecmserver.sh"]
