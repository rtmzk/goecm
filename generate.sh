#!/bin/bash
#

_version=$1

_remote_server=x
_remote_standard_path=/edoc2-data/ecm/${_version}
_remote_template_path=/edoc2-data/ecm/goecm/${_version}/template
_local_generate_path=/tmp/generate

function copy_compose_file() {
  [ ! -d /tmp/generate ] && mkdir ${_local_generate_path}
  scp root@${_remote_server}:${_remote_standard_path}/*.yml ${_local_generate_path}
  scp root@${_remote_server}:${_remote_template_path}/* ${_local_generate_path}
}

function change_attributes_of_image() {
  __compose_file_list=($(ls -l ${_local_generate_path}/*.yml | awk  -F "[ |/]" 'NR!=1{print $NF}' | xargs))

  for _file in ${__compose_file_list[@]}; do
    if [ ! -f ${_local_generate_path}/${_file}.tpl ]; then
      continue
    fi
    __standard_services=($(cat ${_local_generate_path}/${_file} | awk '$0 ~ /^\s{2}\w+\:$/{gsub("#||:",""); print $0}' | sort | uniq | xargs))
    __annotated_services=($(cat ${_local_generate_path}/${_file} | awk '$0 ~ /^#\s{2}\w+\:$/{gsub("#||:",""); print $0}' | sort | uniq | xargs))
    for __service in ${__standard_services[@]}; do
      __service_image=$(sed -n -r "/^\s{2}${__service}\:$/,/^$/{/image/ p}" ${_local_generate_path}/${_file}  | awk '{print $NF}')
      if [ "${__service_image}" == "" ]; then
        continue
      fi 
      sed -i -r "/^\s{2}${__service}\:$/,/^$/{s,<VERSION>,${__service_image},g}" ${_local_generate_path}/${_file}.tpl
    done
    for __annotated_service in ${__annotated_services[@]} ; do
      __annotated_service_image=$(sed -n -r "/^#\s{2}${__annotated_service}\:$/,/^$/{/image/ p}" ${_local_generate_path}/${_file} | awk '{print $NF}')
      if [ ${__annotated_service_image} == "" ]; then
        continue
      fi 
      sed -i -r "/^#\s{2}${__annotated_service}\:$/,/^$/{s,<VERSION>,${__annotated_service_image},g}" ${_local_generate_path}/${_file}.tpl
    done
  done
}

function copy_as_template() {
  cp ${_local_generate_path}/*.tpl ./internal/goecmserver/resource/template/
}

function main() {
    copy_compose_file
    change_attributes_of_image
    copy_as_template
}

main
