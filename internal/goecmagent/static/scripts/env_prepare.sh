#!/bin/bash

function firewall_config() {
#  ports=(2377/tcp 7946/tcp 7946/udp 4789/tcp 4789/udp 80/tcp 443/tcp 9000/tcp 5001/tcp 6789/tcp 6800-7300/tcp 7480/tcp 8080/tcp 7000/tcp)
#  for _p in $ports; do
#    firewall-cmd --permanent --add-port=$_p
#  done
#  firewall-cmd --reload
  systemctl stop firewalld
  systemctl disable firewalld
}

function system_kernel_paramater_config() {
    declare -a system_paramaters=("net.ipv4.ip_forward=1" "fs.file-max=655360" "vm.max_map_count=262144" "vm.min_free_kbytes=1024000" "net.core.somaxconn=2048" "vm.swappiness=0" "fs.inotify.max_user_watches=262144")
    for item in ${system_paramaters[@]}; do
        local key=$(echo $item | awk 'BEGIN{FS="="} {print $1}')
        if ! grep $key /etc/sysctl.conf; then
            eval echo "$item" >>/etc/sysctl.conf
        fi
    done
    if ! grep ip_local_port_range /etc/sysctl.conf; then
        echo "net.ipv4.ip_local_port_range=1024 65300" >>/etc/sysctl.conf
    fi
    sysctl -p
}

function set_date_timezone() {
    for _ip in ${IP_LIST[@]}; do
        log_info "Host: [$_ip]: Set date timezone."
        $SSH root@${_ip} "/usr/bin/cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime"
        log_success "Host: [$_ip]: Set date timzeon done."
    done
}

function selinux_config() {
    SELINUX_STATUS=$(getenforce)
    if [ "${SELINUX_STATUS}" == "Enforcing" ]; then
        setenforce 0
    fi

    SELINUX_CONFIG_STATUS=$(grep "^SELINUX=" /etc/selinux/config)
    if [ "${SELINUX_CONFIG_STATUS}" = "SELINUX=enforcing" ]; then
        sed -i '/^SELINUX/s/enforcing/disabled/g' /etc/selinux/config
    fi
}

function networkmanager_config() {
    system_release_version=$(cat /etc/redhat-release | awk -F "[ |.]" '{print $4}')
    if [ "$system_release_version"x == "8"x ]; then
        return
    elif [ -z $system_release_version ]; then
        return
    fi
    systemctl stop NetworkManager
    systemctl disable NetworkManager
}

function limit_config() {
    if ! grep -E "*\s+-\s+nofile\s+655300" /etc/security/limits.conf &>/dev/null; then
        echo '* - nofile 655300' >>/etc/security/limits.conf
    fi
    if ! grep -E "*\s+-\s+nproc\s+65500" /etc/security/limits.conf &>/dev/null; then
        echo '* - nproc 65500' >>/etc/security/limits.conf
    fi
    limit_file=($(ls /etc/security/limits.d/))
    for f in $limit_file; do
        now_limit=$(awk '/*/ && /proc/{print $NF}' /etc/security/limits.d/$f)
        eval sed -i "s/$now_limit/65500/g" /etc/security/limits.d/$f
    done
}

function history_command_config() {
    cat >>/etc/profile <<EOF

export HISTTIMEFORMAT="%F %T \`whoami\` "
export HISTSIZE=10000
export HISTFILESIZE=10000
EOF
}

function disable_yumrepo() {
    cd /etc/yum.repos.d/
    for i in $(ls *.repo); do
        rename .repo .repo.bak $i
    done
    cd -
}

firewall_config
system_kernel_paramater_config
set_date_timezone
selinux_config
networkmanager_config
limit_config
history_command_config
disable_yumrepo
