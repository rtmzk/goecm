version: '3.5'

services:
  inbiz:
      hostname: inbiz
      image: x
      labels:
        "type": "1"
      environment:
        - LOGGER_ISDEBUG=false
        - LOGGER_MINCALLTIMESPAN=1000
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

networks:
  edoc2:
    external: true
    name: macrowing