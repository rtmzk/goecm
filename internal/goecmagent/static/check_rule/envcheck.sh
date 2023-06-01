#!/bin/bash
#

function firewall_check() {
    firewalld_status=$(systemctl is-active firewalld)
    ##需要开放的端口
    port_list=(2377/tcp 7946/tcp 7946/udp 4789/tcp 4789/udp 80/tcp 443/tcp 9000/tcp 5001/tcp 6789/tcp 6800-7300/tcp 7480/tcp 8080/tcp 7000/tcp)
    not_open_list=()

    if [ "${firewalld_status}" != "active" ]; then
        checkStatus=OK
        message="防火墙已关闭"
    else
        for port in ${port_list[*]}; do
            firewall-cmd --list-ports | grep "${port}" &>/dev/null
            if [ $? -ne 0 ]; then
                not_open_list[${#not_open_list[*]}]=${port}
            fi
        done
        if [ ${#not_open_list[*]} -eq 0 ]; then
            checkStatus=OK
            message="防火墙端口已开放"
        else
            checkStatus=FAILED
            message="以下所需要的端口没有开放: ${not_open_list[@]}"
        fi
    fi

    echo "{\"status\":\"${checkStatus}\", \"message\":\"${message}\"}"
}

function ntp_check() {
    if [ "$(systemctl is-active ntpd)" == "active" ]; then
       ok=$(ntpq -p | awk 'NR>2 {if($1 ~ /\*/) print "true"}')
       if [ "$ok"x != "true"x ]; then
          checkStatus=FAILED
          message="没有可用的时间服务器"
       else
          checkStatus=OK
          message="时间同步服务正在运行"
       fi
    elif [ "$(systemctl is-active chronyd)" == "active" ]; then
       ok=$(chronyc sources | awk '$1 ~ /\*/{if($1 !~ /?/) print "true"}')
       if [ "$ok"x != "true"x ]; then
          checkStatus=FAILED
          message="没有可用的时间服务器"
       else
          checkStatus=OK
          message="时间同步服务正在运行"
       fi
    else
       checkStatus=FAILED
       message="时间服务没有安装或没有启动"
    fi

    echo "{\"status\":\"${checkStatus}\", \"message\":\"${message}\"}"
}

function system_params_check() {
    declare -a FAILED_PARAMS_LIST
    declare -a KERNEL_PARAMS_LIST
    declare -A KERNEL_PARAMS_VALUE_REQUIRED

    KERNEL_PARAMS_LIST=(vm.max_map_count vm.min_free_kbytes net.core.somaxconn vm.swappiness net.ipv4.ip_local_port_range fs.inotify.max_user_watches)
    KERNEL_PARAMS_VALUE_REQUIRED["vm.max_map_count"]=262144
    KERNEL_PARAMS_VALUE_REQUIRED["vm.min_free_kbytes"]=1024000
    KERNEL_PARAMS_VALUE_REQUIRED["net.core.somaxconn"]=2048
    KERNEL_PARAMS_VALUE_REQUIRED["vm.swappiness"]=0
    KERNEL_PARAMS_VALUE_REQUIRED["net.ipv4.ip_local_port_range"]="65300 1024"
    KERNEL_PARAMS_VALUE_REQUIRED["net.ipv4.ip_local_port_range_reduce"]=$((65300 - 1024))
    KERNEL_PARAMS_VALUE_REQUIRED["fs.inotify.max_user_watches"]=1048576

    for _kp in ${KERNEL_PARAMS_LIST[@]}; do
        if sysctl -q $_kp &>/dev/null; then
            if [ "$_kp" != "net.ipv4.ip_local_port_range" ]; then
                system_value=$(sysctl -q $_kp | awk -F "=" '{gsub("\\s","");print $NF}')
                if [[ $(awk 'BEGIN{print('"$system_value"' >= '"${KERNEL_PARAMS_VALUE_REQUIRED[$_kp]}"')?0:1}') -ne 0 ]]; then
                    FAILED_PARAMS_LIST[${#FAILED_PARAMS_LIST[@]}]=$_kp
                fi
            else
                system_value=$(sysctl -q $_kp | awk -F "[ |=]" '{gsub("\\t"," ");print $5-$4}')
                if [[ $(awk 'BEGIN{print('"$system_value"' >= '"${KERNEL_PARAMS_VALUE_REQUIRED["${_kp}_reduce"]}"')?0:1}') -ne 0 ]]; then
                    FAILED_PARAMS_LIST[${#FAILED_PARAMS_LIST[@]}]=$_kp
                fi
            fi
        else
            FAILED_PARAMS_LIST[${#FAILED_PARAMS_LIST[@]}]=$_kp
        fi
    done

    if [ -z "${FAILED_PARAMS_LIST}" ]; then
        echo "{\"status\": \"OK\", \"message\": \"内核参数检查无异常\"}"
    else
        echo "{\"status\": \"FAILED\", \"message\": \"$(for((i=0;i<${#FAILED_PARAMS_LIST[@]};i++)); do  printf "${FAILED_PARAMS_LIST[$i]} 检查失败, 要求: ${KERNEL_PARAMS_VALUE_REQUIRED[${FAILED_PARAMS_LIST[$i]}]};" ;done)\"}"
    fi
}

function selinux_check() {
    selinux_status=$(getenforce)
    if [ $selinux_status == "Disabled" ]; then
        checkStatus=OK
        message="SElinux已关闭"
    elif [ $selinux_status == "enforcing" ]; then
        checkStatus=FAILED
        message="SElinux处于打开状态"
    elif [ $(cat /etc/selinux/config | awk -F "=" '/SELINUX/{print $2}') != "enforcing" ]; then
        checkStatus=FAILED
        message="SElinux已关闭"
    else
        checkStatus=FAILED
        message="SElinux已临时关闭，请检查/etc/selinux/config"
    fi
    echo "{\"status\": \"${checkStatus}\", \"message\": \"${message}\"}"
}

function cpu_core_check() {
    cpus=$(lscpu | awk '/^CPU\(s):/{print $2}')
    req_cpus=16
    if [[ -n ${req_cpus} && -n ${cpus} ]]; then
        if [[ $(awk 'BEGIN{print('"$cpus"'>='"$req_cpus"')?0:1}') -eq 0 ]]; then
            checkStatus=OK
            Msg="核心数大于等于${req_cpus}"
        else
            checkStatus=FAILED
            Msg="cpu 核心数小于${req_cpus}"
        fi
    else
        checkStatus=FAILED
        Msg="获取宿主机cpu核心数失败"
    fi

    echo "{\"status\": \"${checkStatus}\", \"message\": \"${Msg}\"}"
}


function mem_total_check() {
    REQ_MEM=32G
    mem_total=$(cat /proc/meminfo | awk '/MemTotal:/{print $2$3}')
    mem_gb=$(unit_convert $mem_total)
    req_mem=${REQ_MEM:0:-1}

    if [[ -n ${req_mem} && -n ${mem_gb} ]]; then
        if [[ $(awk 'BEGIN{print('"$mem_gb"'<'"$req_mem"')?0:1}') -eq 0 && $(awk 'BEGIN{print('"$mem_gb"'-'"$req_mem"'<=-1)?0:1}') -eq 0 ]]; then
            checkStatus=FAILED
            message="内存小于${REQ_MEM}"
        else
            checkStatus=OK
            message="内存大于${REQ_MEM}"
        fi
    else
        checkStatus=FAILED
        message="获取主机内存失败"
    fi

    echo "{\"status\":\"${checkStatus}\", \"message\":\"${message}\"}"
}


function unit_convert() {
    typeset -l size
    size=$1
    to_gb=$(echo $size | awk '{
            if ($1 ~ /pb$/ || $1 ~ /p$/){
                print $1 * 1024 * 1024
            } else if ($1 ~ /tb$/ || $1 ~ /t$/){
                print $1 * 1024
            } else if ($1 ~ /gb$/ || $1 ~ /g$/){
                print $1
            } else if ($1 ~ /mb$/ || $1 ~ /m$/){
                print $1/1024
            } else if ($1 ~ /kb$/ || $1 ~ /k$/){
                print $1/1024/1024
            } else if ($1 ~ /b$/){
                print $1/1024/1024/1024
            }
        }')
    echo $to_gb
}

$@