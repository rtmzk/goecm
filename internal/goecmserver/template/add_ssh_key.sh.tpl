#!/bin/bash

[ -d ~/.ssh ] || mkdir -p ~/.ssh && chmod 700 ~/.ssh
cat >> ~/.ssh/authorized_keys << EOF
{{.}}
EOF
chmod 600 ~/.ssh/authorized_keys