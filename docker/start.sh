#!/bin/bash

mkdir /dev/shm/logs -p

alias cp="cp"
cp -rf /opt/apps/$PROJECT_NAME/docker/supervisord/apps.conf  /etc/supervisord/

# 启动supervisor 管理服务
/usr/bin/supervisord -c /etc/supervisord/apps.conf
