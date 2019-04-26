#!/bin/bash


mkdir -p ${PROJECT_PATH}/bin && mkdir -p /etc/supervisord/ && mkdir -p /build_time \
    && cd /build_time && touch "`date '+%Y-%m-%d %H:%M:%S'`"