#!/bin/bash

 ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

mkdir -p ${PROJECT_PATH}/bin && mkdir -p /etc/supervisord/ && mkdir -p /build_time \
    && cd /build_time && touch "`date '+%Y-%m-%d %H:%M:%S'`"