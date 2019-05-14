#!/bin/bash

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
    docker-compose -f mysql-compose.yml up -d
}


if [ "$(docker ps -aq -f name=${MYSQL_HOST})" ]; then
    echo "${MYSQL_HOST} already running ..."
    if [ $isRebuild == y ]; then
        echo "killing ${MYSQL_HOST}"
        docker rm ${MYSQL_HOST} -f
        RunMysql
    fi
else
    RunMysql
fi