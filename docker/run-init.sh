#!/bin/bash

export $(cut -d= -f1 ../conf/app.env)
source ../conf/app.env

if [ ! "$(docker network ls | grep mysql-network)" ]; then
  echo "Creating mysql-network network ..."
  docker network create mysql-network
else
  echo "mysql-network network exists."
fi

isRebuild=$1
isRebuild=${isRebuild:=n}

RunMysql() {
    # run your container in our global network shared by different projects
    echo "Running MYSQL in global mysql-network network ..."
    if [ $1 == y ]; then
        docker-compose -f mysql-compose.yml up -d --build
    else
        docker-compose -f mysql-compose.yml up -d
    fi
}


if [ "$(docker ps -aq -f name=${MYSQL_HOST})" ]; then
    echo "${MYSQL_HOST} already running ...isRebuild: ${isRebuild}"
    if [ $isRebuild == y ]; then
        echo "killing ${MYSQL_HOST}"
        docker rm ${MYSQL_HOST} -f
        RunMysql $isRebuild
    fi
else
    RunMysql $isRebuild
fi