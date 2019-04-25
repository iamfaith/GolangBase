#!/bin/bash


export $(cut -d= -f1 ../conf/app.env)
source ../conf/app.env

export APP_VERSION=1.0
APP_IMAGE=xianzixiang/`echo -n $PROJECT_NAME | awk '{print tolower($0)}'`:$APP_VERSION

echo running image [${APP_IMAGE}].....

docker-compose up --build