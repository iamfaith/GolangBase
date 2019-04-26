#!/bin/bash

export $(cut -d= -f1 conf/app.env)
source conf/app.env

V=$1
V=${V:=0.1}
tagName=`echo -n $PROJECT_NAME | awk '{print tolower($0)}'`:$V
echo Building App as version: ${V} Tag Name: ${tagName}

docker rmi -f xianzixiang/${tagName

docker build -t xianzixiang/${tagName} -f ./docker/Dockerfile --build-arg PROJECT_NAME=$PROJECT_NAME  --build-arg PROJECT_PATH=$PROJECT_PATH .