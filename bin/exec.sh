#!/bin/bash


export $(cut -d= -f1 ../conf/app.env)
source ../conf/app.env

APP_IMAGE=${APP_IMG}:${APP_VERSION}

op=$1
op=${op:=n}

if [ $op != n ]; then

    echo remove image ${APP_IMAGE} ...
    docker rmi $APP_IMAGE -f
fi

echo running image [${APP_IMAGE}].....
docker-compose stop
docker system prune -f
docker-compose up -d --build