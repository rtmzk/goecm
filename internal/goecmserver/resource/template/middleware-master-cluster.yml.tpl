version: "3.7"

services:
{{ if not .MiddlewareDeploySpec.Database.IsExternal }}
### SVC:mysql:>>>
  mysql0:
    hostname: mysql0
    image: x
    volumes:
      - {{ .MiddlewareDeploySpec.Database.DBDataPath }}:/var/lib/mysql
      - {{ .MiddlewareDeploySpec.Database.DBBackupPath }}:/dbbackup
    init: true
    environment:
      - DBBACKUP=true
      - TASKS=10 0 * * *
      - DAYS=15
      - MODE=cluster
    labels:
      "type": "2"
    networks:
      - middleware
    deploy:
      placement:
        constraints:
          - node.labels.mysqllabels == mysql0
      replicas: 1
      restart_policy:
        condition: on-failure
    logging:
      driver: "json-file"
      options:
        max-size: "2048k"
        max-file: "10"

  mysql1:
    hostname: mysql1
    image: x
    volumes:
      - {{ .MiddlewareDeploySpec.Database.DBDataPath }}:/var/lib/mysql
      - {{ .MiddlewareDeploySpec.Database.DBBackupPath }}:/dbbackup
    init: true
    environment:
      - DBBACKUP=true
      - TASKS=10 0 * * *
      - DAYS=15
      - MODE=cluster
    labels:
      "type": "2"
    networks:
      - middleware
    deploy:
      replicas: 1
      placement:
        constraints:
          - node.labels.mysqllabels == mysql1
      restart_policy:
        condition: on-failure
    logging:
      driver: "json-file"
      options:
        max-size: "2048k"
        max-file: "10"

  mysql2:
    hostname: mysql2
    image: x
    volumes:
      - {{ .MiddlewareDeploySpec.Database.DBDataPath }}:/var/lib/mysql
      - {{ .MiddlewareDeploySpec.Database.DBBackupPath }}:/dbbackup
    init: true
    environment:
      - DBBACKUP=true
      - TASKS=10 0 * * *
      - DAYS=15
      - MODE=cluster
    labels:
      "type": "2"
    networks:
      - middleware
    deploy:
      placement:
        constraints:
          - node.labels.mysqllabels == mysql2
      replicas: 1
      restart_policy:
        condition: on-failure
    logging:
      driver: "json-file"
      options:
        max-size: "2048k"
        max-file: "10"
### SVC:mysql:<<<
{{ end }}

### SVC:redis:>>>
  redis1:
    hostname: redis1
    image: x
#    entrypoint: /opt/docker-entrypoint-host.sh
#    extra_hosts:
#      - "redis1:x"
#      - "redis2:x"
#      - "redis3:x"
    volumes:
      - {{ .MiddlewareDeploySpec.Redis.RedisDataPath }}:/data
    networks:
      - middleware
    environment:
      - ROLE=master
    deploy:
      placement:
        constraints:
          - node.labels.redislabels == redis1
      replicas: 1
      endpoint_mode: dnsrr
      restart_policy:
        condition: on-failure
    logging:
      driver: "json-file"
      options:
        max-size: "2048k"
        max-file: "10"

  redis2:
    hostname: redis2
    image: x
#    entrypoint: /opt/docker-entrypoint-host.sh
#    extra_hosts:
#      - "redis1:x"
#      - "redis2:x"
#      - "redis3:x"
    volumes:
      - {{ .MiddlewareDeploySpec.Redis.RedisDataPath }}:/data
    networks:
      - middleware
    environment:
      - ROLE=slave
    deploy:
      placement:
        constraints:
          - node.labels.redislabels == redis2
      replicas: 1
      endpoint_mode: dnsrr
      restart_policy:
        condition: on-failure
    logging:
      driver: "json-file"
      options:
        max-size: "2048k"
        max-file: "10"

  redis3:
    hostname: redis3
    image: x
#    entrypoint: /opt/docker-entrypoint-host.sh
#    extra_hosts:
#      - "redis1:x"
#      - "redis2:x"
#      - "redis3:x"
    volumes:
      - {{ .MiddlewareDeploySpec.Redis.RedisDataPath }}:/data
    networks:
      - middleware
    environment:
      - ROLE=slave
    deploy:
      placement:
        constraints:
          - node.labels.redislabels == redis3
      replicas: 1
      endpoint_mode: dnsrr
      restart_policy:
        condition: on-failure
    logging:
      driver: "json-file"
      options:
        max-size: "2048k"
        max-file: "10"
### SVC:redis:<<<

### SVC:rabbitmq:>>>
  rabbitmq:
    hostname: rabbitmq
    image: x
    volumes:
      - {{ .MiddlewareDeploySpec.Rabbitmq.RabbitmqDataPath }}:/var/lib/rabbitmq
    networks:
      - middleware
    environment:
      - ROLE=master
    deploy:
      placement:
        constraints:
          - node.labels.rabbitmqlabels == rabbitmq1
      replicas: 1
      restart_policy:
        condition: on-failure
    logging:
      driver: "json-file"
      options:
        max-size: "2048k"
        max-file: "10"
  rabbitmq2:
    hostname: rabbitmq2
    image: x
    depends_on:
      - rabbitmq
    volumes:
      - {{ .MiddlewareDeploySpec.Rabbitmq.RabbitmqDataPath }}:/var/lib/rabbitmq
    networks:
      - middleware
    environment:
      - ROLE=slave
    deploy:
      placement:
        constraints:
          - node.labels.rabbitmqlabels == rabbitmq2
      replicas: 1
      restart_policy:
        condition: on-failure
    logging:
      driver: "json-file"
      options:
        max-size: "2048k"
        max-file: "10"
### SVC:rabbitmq:<<<

### SVC:es:>>>
  es:
    hostname: es
    image: x
    volumes:
      - {{ .MiddlewareDeploySpec.Elasticsearch.ElasticsearchDataPath }}:/esdata
    environment:
      - ES_JAVA_OPTS=-Xms8g -Xmx8g
      - NODE_NAME=es-node1
      - MODE=cluster
    networks:
      - middleware
    deploy:
      placement:
        constraints:
          - node.labels.eslabels == es1
      replicas: 1
      endpoint_mode: dnsrr
      restart_policy:
        condition: on-failure
    logging:
      driver: "json-file"
      options:
        max-size: "2048k"
        max-file: "10"

  es2:
    hostname: es2
    image: x
    volumes:
      - {{ .MiddlewareDeploySpec.Elasticsearch.ElasticsearchDataPath }}:/esdata
    environment:
      - ES_JAVA_OPTS=-Xms8g -Xmx8g
      - NODE_NAME=es-node2
      - MODE=cluster
    networks:
      - middleware
    deploy:
      placement:
        constraints:
          - node.labels.eslabels == es2
      replicas: 1
      endpoint_mode: dnsrr
      restart_policy:
        condition: on-failure
    logging:
      driver: "json-file"
      options:
        max-size: "2048k"
        max-file: "10"

  es3:
    hostname: es3
    image: x
    volumes:
      - {{ .MiddlewareDeploySpec.Elasticsearch.ElasticsearchDataPath }}:/esdata
    environment:
      - ES_JAVA_OPTS=-Xms8g -Xmx8g
      - NODE_NAME=es-node3
      - MODE=cluster
    networks:
      - middleware
    deploy:
      placement:
        constraints:
          - node.labels.eslabels == es3
      replicas: 1
      endpoint_mode: dnsrr
      restart_policy:
        condition: on-failure
    logging:
      driver: "json-file"
      options:
        max-size: "2048k"
        max-file: "10"
### SVC:es:<<<

### SVC:mysql:>>>
  mysql:
    hostname: mysql
    image: <HAPROXY.VERSION>
    deploy:
      replicas: 1
      restart_policy:
        condition: on-failure
    networks:
      - middleware
    logging:
      driver: "json-file"
      options:
        max-size: "2000k"
        max-file: "10"
### SVC:mysql:<<<
### SVC:proxy:>>>
  proxy:
    hostname: proxy
    image: x
    command: /etc/nginx/proxy.sh
    environment:
      - MIDDLEWARE=cluster
      - PROXY_SERVICES=es
    ports:
      - "30001:30001" #mysql
      - "30003:30003" #es
      - "30004:30004" #rabbitmq
      - "30005:30005" #redis
      - "30006:30006" #haproxy
      - "5201:5201" #iperf3
    deploy:
      replicas: 1
      restart_policy:
        condition: on-failure
    networks:
      - middleware
    logging:
      driver: "json-file"
      options:
        max-size: "2000k"
        max-file: "10"
### SVC:proxy:<<<

networks:
  middleware:
    external: true
    name: macrowing
  host:
    external: true
    name: host
