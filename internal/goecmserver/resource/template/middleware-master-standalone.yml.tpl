version: '3.7'

services:
{{ if not .MiddlewareDeploySpec.Database.IsExternal }}
### SVC:mysql:>>>
  mysql:
    hostname: mysql
    image: x
    volumes:
      - {{ .MiddlewareDeploySpec.Database.DBDataPath }}:/var/lib/mysql
      - {{ .MiddlewareDeploySpec.Database.DBBackupPath }}:/dbbackup
    init: true
    environment:
      - DBBACKUP=true
      - TASKS=10 0 * * *
      - DAYS=15
      - MODE=standalone
    networks:
      - middleware
    deploy:
      placement:
        constraints:
          - node.labels.nodelabels == Middleware
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

  redis:
    hostname: redis
    image: x
    volumes:
      - {{ .MiddlewareDeploySpec.Redis.RedisDataPath }}:/data
    networks:
      - middleware
    deploy:
      replicas: 1
      placement:
        constraints:
          - node.labels.nodelabels == Middleware
      endpoint_mode: dnsrr
      restart_policy:
        condition: on-failure
    logging:
      driver: "json-file"
      options:
        max-size: "2048k"
        max-file: "10"

  rabbitmq:
    hostname: rabbitmq
    image: x
    volumes:
      - {{ .MiddlewareDeploySpec.Rabbitmq.RabbitmqDataPath }}:/var/lib/rabbitmq
    networks:
      - middleware
    deploy:
      placement:
        constraints:
          - node.labels.nodelabels == Middleware
      replicas: 1
      restart_policy:
        condition: on-failure
    logging:
      driver: "json-file"
      options:
        max-size: "2048k"
        max-file: "10"

  es:
    image: x
    hostname: es
    volumes:
      - {{ .MiddlewareDeploySpec.Elasticsearch.ElasticsearchDataPath }}:/esdata
      - {{ .MiddlewareDeploySpec.Elasticsearch.ElasticsearchBackupPath }}:/esbackup
    networks:
      - middleware
    environment:
      - node.name=es
      - ES_JAVA_OPTS=-Xmx8g -Xms8g
    deploy:
      replicas: 1
      placement:
        constraints:
          - node.labels.nodelabels == Middleware
      restart_policy:
        condition: on-failure
    logging:
      driver: "json-file"
      options:
        max-size: "2048k"
        max-file: "10"

  proxy:
    hostname: proxy
    image: x
    command: /etc/nginx/proxy.sh
    environment:
      - MIDDLEWARE=standalone
      - PROXY_SERVICES=es
    ports:
      - "30001:30001"    #mysql
      - "30003:30003"    #es
      - "30004:30004"    #rabbitmq
      - "30005:30005"    #redis
      - "5201:5201"      #iperf3
    deploy:
      replicas: 1
      placement:
        constraints:
          - node.labels.nodelabels == Middleware
      restart_policy:
        condition: on-failure
    networks:
      - middleware
    logging:
      driver: "json-file"
      options:
        max-size: "2000k"
        max-file: "10"

networks:
  middleware:
    external: true
    name: macrowing
