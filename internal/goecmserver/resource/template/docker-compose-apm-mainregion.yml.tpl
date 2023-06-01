version: '3.5'
services:
  #apm主服务
  apm:
    hostname: apm
    image: x
    networks:
      - macrowing
    ports:
      - "8083:80"
    environment:
      - number=30000000
      - APM_ESURL=http://proxy:30003
      - MiddlewareIsCluster=0
      - IsUsedGRPC=1
    deploy:
      placement:
        constraints:
          - node.labels.isapm == true
      replicas: 1
      restart_policy:
        condition: on-failure
    logging:
      driver: "json-file"
      options:
        max-size: "2000k"
        max-file: "10"

  #cpu、内存、硬盘、网络监控
  basic:
    image: x
    user: root
    volumes:
      - type: bind
        source: /proc
        target: /hostfs/proc
        read_only: true
      - type: bind
        source: /sys/fs/cgroup
        target: /hostfs/sys/fs/cgroup
        read_only: true
      - type: bind
        source: /
        target: /hostfs
        read_only: true
      - type: bind
        source: /var/run/docker.sock
        target: /var/run/docker.sock
        read_only: true
    deploy:
      mode: global
      restart_policy:
        condition: on-failure
    networks:
      - hostnet
    logging:
      driver: "json-file"
      options:
        max-size: "2000k"
        max-file: "10"

  #监控检查
  heartbeat:
    hostname: heartbeat
    image: x
    environment:
      - ELASTICSEARCH_HOSTS=es:9200
    networks: 
      - macrowing
    deploy:
      placement:
        constraints:
          - node.labels.isapm == true
      replicas: 1
      restart_policy:
        condition: on-failure
    healthcheck:
      test: curl -s -L -k -i http://apm/apm/env | grep "200 OK"
      interval: 20s
      timeout: 10s
      retries: 5 
    logging:
      driver: "json-file"
      options:
        max-size: "2000k"
        max-file: "10"

  filebeat:
    hostname: filebeat
    image: x
    environment:
      - ELASTICSEARCH_HOSTS=es:9200
    volumes:
      - /var/log/:/hostfs/var/log/
      - /tmp:/usr/share/filebeat/data
    networks: 
      - macrowing
    deploy:
      mode: global
      restart_policy:
        condition: on-failure
    logging:
      driver: "json-file"
      options:
        max-size: "2000k"
        max-file: "10"

  #中间件监控
  middleware:
    hostname: middleware
    image: x
    networks: 
      - macrowing
    deploy:
      placement:
        constraints:
          - node.labels.isapm == true
      replicas: 1
      restart_policy:
        condition: on-failure
    healthcheck:
      test: curl -s -L -k -i http://apm/apm/env | grep "200 OK"
      interval: 20s
      timeout: 10s
      retries: 5 
    logging:
      driver: "json-file"
      options:
        max-size: "2000k"
        max-file: "10"

  #日志收集
  logstash:
    hostname: logstash
    image: x
    environment:
      - LS_JAVA_OPTS=-Xms2g -Xmx2g
      - COLLECT_LOG_SERVICES=["edoc2","content","unionsync","orgsync","transport","regionsync","decryption","edrms","inbiz","auth","perm","org"]
      - API_LOG_MODULE=["WebAPI","CoreAPI","DB","EDoc2.PermCtrl","EDoc2.Cache","EDRMS-DB","EDRMS-WEBAPI"]
    ports:
      - "55514:55514/udp"
    deploy:
      placement:
        constraints:
          - node.labels.isapm == true
      replicas: 1
      restart_policy:
        condition: on-failure
    networks:
      - macrowing
    logging:
      driver: "json-file"
      options:
        max-size: "2000k"
        max-file: "10"

networks:
  hostnet:
    external: true
    name: host
  macrowing:
    external: true
    name: macrowing
