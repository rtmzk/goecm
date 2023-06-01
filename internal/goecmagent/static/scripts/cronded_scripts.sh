#!/bin/bash
#

clean_cache_script_path=/opt/scripts/clean-cache.sh
clean_container_script_path=/opt/scripts/clean-container.sh

if [ ! -d  /opt/scripts ]; then
  mkdir  /opt/scripts
fi


function gender_script() {
  [ -f ${clean_cache_script_path} ] && >${clean_cache_script_path}
  [ -f ${clean_container_script_path} ] && >${clean_container_script_path}
  cat >${clean_cache_script_path}<<'EOF'
#!/bin/bash
#

CACHE=`free -m | awk '/Mem/{print $(NF-1)}'`
TOTAL=`free -m | awk '/Mem/{print $2}'`
FREE=`free -m | awk '/Mem/{print $4}'`
FREE_RATE=`echo "scale=2;$FREE/$TOTAL" | bc | awk -F'.' '{print $2}'`
LOG=/tmp

if [ -z "$FREE_RATE" ]; then
    FREE_RATE=0
fi


if [ $FREE_RATE -lt 15 ]; then
    if [ $(ps aux | grep -v grep | grep -c "/proc/sys/vm/drop_caches" ) -ge 1 ]; then
        echo "$(date) memory free rate: $(echo $FREE_RATE | sed -r 's/^0*([1-9])/\1/g')%, process already running!" >> ${LOG}/clean_cache.log
        exit 4
    else
        sync;sync;sleep 5
        echo 1 >/proc/sys/vm/drop_caches
        echo "$(date) memory free rate: $(echo $FREE_RATE | sed -r 's/^0*([1-9])/\1/g')%, clean success!" >> ${LOG}/clean_cache.log
    fi
fi
EOF

  cat >${clean_container_script_path}<<'EOF'
#!/bin/bash
#

purge_id=$(docker ps -qf status=exited && docker ps -qf status=created && docker ps -qf status=dead)

if [ -n "$purge_id" ]; then
    docker rm $purge_id
fi
EOF
}


function main() {
  gender_script

  echo '1 */1 * * *  /bin/bash /opt/scripts/clean-cache.sh &> /dev/null' >> /var/spool/cron/root
  echo '0 0 * * *  /bin/bash /opt/scripts/clean-container.sh &> /dev/null' >> /var/spool/cron/root
}