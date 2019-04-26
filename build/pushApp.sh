#!/bin/bash

export $(cut -d= -f1 conf/app.env)
source conf/app.env

V=$1
V=${V:=0.1}
echo Push App as version: ${V}

docker push xianzixiang/`echo -n $PROJECT_NAME | awk '{print tolower($0)}'`:$V


