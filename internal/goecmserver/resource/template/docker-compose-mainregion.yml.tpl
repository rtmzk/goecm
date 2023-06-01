version: '3.5'

services:
  initdatabase:
    hostname: initdatabase
    image: x
    labels:
      "type": "1"
    networks:
      - edoc2
    cap_add:
      - SYS_PTRACE
    env_file:
      ./envfile.env
    deploy:
      placement:
        constraints:
          - node.labels.nodetype == InDrive
      replicas: 1
      restart_policy:
        condition: on-failure

  edoc2:
    hostname: edoc2
    image: x
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
#      - /home/edoc2/macrowing/edoc2v5/data/edoc2:/app/wwwroot/external
    {{ setLocalOrNasDir . }}:/edoc2Docs
    labels:
      "type": "1"
    environment:
      - LOGGER_ISDEBUG=false
      - LOGGER_MINCALLTIMESPAN=1000
      - PUBLISH_EXTERNAL_ADDRESS=http://localhost/
      - PRODUCTION=edoc2
#      - NUMS=25
#      - ThreadCheck=true
    env_file:
      ./envfile.env
    networks:
      - edoc2
    cap_add:
      - SYS_PTRACE
    stop_signal: SIGKILL
    healthcheck:
      test: /bin/bash /opt/edoc2_check.sh
      interval: 15s
      timeout: 10s
      retries: 3
    deploy:
      placement:
        constraints:
          - node.labels.nodetype == InDrive
      mode: global
      endpoint_mode: dnsrr
      restart_policy:
        condition: on-failure
#      resources:
#        limits:
#          cpus: '8'
    logging:
      driver: "json-file"
      options:
        max-size: "2000k"
        max-file: "10"
#    logging:
#      driver: syslog
#      options:
#        syslog-address: "udp://logserver:55514"
#        syslog-format: "rfc3164"
#        tag: "edoc2"

  decryption:
    hostname: decryption
    image: x
    volumes:
    {{ setLocalOrNasDir . }}:/edoc2Docs
    env_file: ./envfile.env
    networks:
      - edoc2
    cap_add:
      - SYS_PTRACE
    deploy:
      placement:
        constraints:
          - node.labels.nodetype == InDrive
      replicas: 1
      endpoint_mode: dnsrr
      restart_policy:
        condition: any

  transport:
    hostname: transport
    image: x
#    ports:
#      - target: 6261
#        published: 6261
#        mode: host
    volumes:
#      - /home/edoc2/macrowing/edoc2v5/data/edoc2:/app/wwwroot/external
    {{ setLocalOrNasDir . }}:/edoc2Docs
    env_file:
      ./envfile.env
    networks:
      - edoc2
    cap_add:
      - SYS_PTRACE
    deploy:
      placement:
        constraints:
          - node.labels.nodetype == InDrive
      mode: global
      endpoint_mode: dnsrr
      restart_policy:
        condition: any
    healthcheck:
      test: curl -L -k -i http://localhost:6261/api/values|grep "200 OK"
      interval: 10s
      timeout: 3s
      retries: 3

  auth:
    hostname: auth
    image: x
    env_file: ./envfile.env
    networks:
      - edoc2
    cap_add:
      - SYS_PTRACE
    deploy:
      placement:
        constraints:
          - node.labels.nodetype == InDrive
      mode: global
      endpoint_mode: dnsrr
      restart_policy:
        condition: any
    healthcheck:
      test: curl -L -k -i http://localhost:6320/api/values|grep "200 OK"
      interval: 10s
      timeout: 3s
      retries: 3

  org:
    hostname: org
    image: x
    env_file: ./envfile.env
    networks:
      - edoc2
    cap_add:
      - SYS_PTRACE
    deploy:
      placement:
        constraints:
          - node.labels.nodetype == InDrive
      mode: global
      endpoint_mode: dnsrr
      restart_policy:
        condition: any
    healthcheck:
      test: curl -L -k -i http://localhost:6310/api/values|grep "200 OK"
      interval: 10s
      timeout: 3s
      retries: 3

  perm:
    hostname: perm
    image: x
    env_file: ./envfile.env
    networks:
      - edoc2
    cap_add:
      - SYS_PTRACE
    deploy:
      placement:
        constraints:
          - node.labels.nodetype == InDrive
      mode: global
      endpoint_mode: dnsrr
      restart_policy:
        condition: any
    healthcheck:
      test: curl -L -k -i http://localhost:6330/api/values|grep "200 OK"
      interval: 10s
      timeout: 3s
      retries: 3

  wf:
    hostname: wf
    image: x
    labels:
      "type": "2"
    networks:
      - edoc2
    env_file:
      ./envfile.env
    deploy:
      replicas: 1
      placement:
        constraints:
          - node.labels.nodetype == InDrive
      restart_policy:
        condition: on-failure
#    healthcheck:
#      test: curl -L -k -i http://edoc2/api/values|grep "200 OK"
#      interval: 10s
#      timeout: 3s
#      retries: 3
    logging:
      driver: "json-file"
      options:
        max-size: "2000k"
        max-file: "10"

#  regionsync:
#    hostname: regionsync
#    image: <VERSION>
#    volumes:
#    {{ setLocalOrNasDir . }}:/edoc2Docs
#    labels:
#      "type": "1"
#    networks:
#      - edoc2
#    deploy:
#      placement:
#        constraints:
#          - node.labels.nodetype == InDrive
#      replicas: 1
#      restart_policy:
#        condition: on-failure

#  storagetransfer:
#    hostname: storagetransfer
#    image: <VERSION>
#    volumes:
#    {{ setLocalOrNasDir . }}:/edoc2Docs
#    networks: 
#      - edoc2
#    labels:
#      "type": "1"
#    deploy:
#      placement:
#        constraints:
#          - node.labels.nodetype == InDrive
#      replicas: 1
#      restart_policy:
#        condition: on-failure

  orgsync:
    hostname: orgsync
    image: x
    labels:
      "type": "1"
    networks:
      - edoc2
    cap_add:
      - SYS_PTRACE
    environment:
      - ISTITLE=0
    env_file:
      ./envfile.env
    deploy:
      placement:
        constraints:
          - node.labels.nodetype == InDrive
      replicas: 1
      restart_policy:
        condition: on-failure

#  importsync:
#    hostname: importsync
#    image: <VERSION>
#    volumes:
#    {{ setLocalOrNasDir . }}:/edoc2Docs
#    labels:
#      "type": "1"
#    networks:
#      - edoc2
#    deploy:
#      placement:
#        constraints:
#          - node.labels.nodetype == InDrive
#      replicas: 1
#      restart_policy:
#        condition: on-failure

  content:
    hostname: content
    image: x
#    ports:
#      - target: 5011
#        published: 5011
#        mode: host
    volumes:
    {{ setLocalOrNasDir . }}:/edoc2Docs
#      - /home/edoc2/macrowing/edoc2v5/data/office:/app/Contents/Conversion/microsoft_docs
#      - /home/edoc2/macrowing/edoc2v5/data/aboutpip:/app/aboutpip/aboutpip
#      - /home/edoc2/macrowing/edoc2v5/data/detection_ocr:/app/detection_ocr/detection_ocr
    labels:
      "type": "1"
    networks:
      - edoc2
    cap_add:
      - SYS_PTRACE
    deploy:
      placement:
        constraints:
          - node.labels.nodetype == InDrive
      mode: global
      endpoint_mode: dnsrr
      restart_policy:
        condition: on-failure
      resources:
        limits:
          cpus: '4'
          memory: 8G
    env_file:
      ./envfile.env
#    environment:
#      - ESURL=http://es:9200
#      - enable_clamav=true

# regionclearcopy:
#   hostname: regionclearcopy
#   image: <VERSION>
#   volumes:
#    {{ setLocalOrNasDir . }}:/edoc2Docs
#     - /home/edoc2/macrowing/edoc2v5/data/edoc2Docs:/edoc2Docs
#   labels:
#     "type": "1"
#   networks: 
#     - edoc2
#   deploy:
#     placement:
#       constraints:
#         - node.labels.nodetype == InDrive
#     replicas: 1
#     restart_policy:
#       condition: on-failure
#   env_file:
#     ./envfile.env

#  unionsync:
#    hostname: unionsync
#    image: <VERSION>
#    labels:
#      "type": "1"
#    networks:
#      - edoc2
#    deploy:
#      placement:
#        constraints:
#          - node.labels.nodetype == InDrive
#      replicas: 1
#      restart_policy:
#        condition: on-failure

  nginx: 
    hostname: nginx
    image: <VERSION>
    ports:
      - target: 80
        published: 80
        mode: host
#      - target: 443
#        published: 443
#        mode: host
    environment:
      - INBIZ=true
      - WF=true
      - LB_RR=true
#      - SEARCHENGINE=true
#      - UDCMGT=true
#      - WEBOFFICE=true
#      - INSIGHT=true
#      - INBIZTHIRD=true
#      - SSL=true
#      - DOMAIN_NAME=ab.edoc2.com
#    volumes:
#      - /home/edoc2/macrowing/edoc2v5/data:/etc/nginx/cert
    labels:
      "type": "0"
    deploy:
      placement:
        constraints:
          - node.labels.nodetype == InDrive
      mode: global
      restart_policy:
        condition: on-failure
    networks:
      - edoc2
    healthcheck:
      test: bash /opt/nginx/check.sh
      interval: 20s
      timeout: 10s
      retries: 3
    logging:
      driver: "json-file"
      options:
        max-size: "2000k"
        max-file: "10"

#  weboffice:
#    hostname: weboffice
#    image: <VERSION>
#    volumes:
#      - /datavol/weboffice/Data:/var/www/onlyoffice/Data
#      - /datavol/weboffice/onlyoffice:/var/lib/onlyoffice
#      - /datavol/weboffice/postgresql:/var/lib/postgresql
#      - /datavol/weboffice/log:/var/log/onlyoffice
##    ports:
##      - "5001:80"
#    labels:
#      "type": "0"
#    deploy:
#      placement:
#        constraints:
#          - node.labels.nodetype == InDrive
#      replicas: 1
#      restart_policy:
#        condition: on-failure
#    networks:
#      - edoc2

networks:
  edoc2:
    external: true
    name: macrowing

