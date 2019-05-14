#!/bin/bash

echo "begin to deploy"

. /etc/profile
. /home/faith/.profile

eval `ssh-agent -s`
ssh-add ~/.ssh/id_rsa


git pull && cd docker/ && chmod +x exec.sh && ./exec.sh y