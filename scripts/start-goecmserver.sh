#!/bin/bash

datadir=/data/goecmserver/mysql

function start_mysql_daemon() {
    /bin/cp -f  /opt/scripts/my.cnf /etc/mysql/my.cnf
    if [ ! -d ${datadir} ] ; then
        mkdir -p ${datadir}
    fi 
    if [ ! -d ${datadir}/mysql ] ; then
        mysqld --user=root --initialize-insecure
        mysqld --user=root --skip-networking --socket=/var/run/mysqld/mysqld.sock &
        pid="$!"
        mysql=( mysql --protocol=socket -uroot -hlocalhost --socket=/var/run/mysqld/mysqld.sock)
        
        for i in {30..0} ; do
            if echo 'SELECT 1' | "${mysql[@]}" &> /dev/null; then 
                break
            fi
            echo "MySQL init process in progress..."
            sleep 1
        done       
        if [ "$i" = 0 ]; then
            echo >&2 'MySQL init progress failed.'
            exit 1
        fi 
        "${mysql[@]}" <<-EOSQL
            SET @@SESSION.SQL_LOG_BIN=0;
            ALTER USER 'root'@'localhost' IDENTIFIED BY '${MYSQL_ROOT_PASSWORD}';
            CREATE USER 'root'@'%' IDENTIFIED BY '${MYSQL_ROOT_PASSWORD}';
            CREATE DATABASE IF NOT EXISTS goecm DEFAULT CHARACTER SET utf8;
            GRANT ALL ON *.* TO 'root'@'localhost' WITH GRANT OPTION;
            GRANT ALL ON *.* TO 'root'@'%' WITH GRANT OPTION;
            DROP DATABASE IF EXISTS test;
            FLUSH PRIVILEGES;
EOSQL
        
        if ! kill -s TERM "$pid" || ! wait "$pid"; then
            echo >&2 "MySQL init process failed."
            exit 1
        fi
    fi
    mysqld --user=root -D 
}

function main() {
    start_mysql_daemon
    /usr/local/bin/goecmserver -c /opt/scripts/config.yaml
}

main
