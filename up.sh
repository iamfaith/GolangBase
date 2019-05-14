#!/bin/bash

export $(cut -d= -f1 conf/app.env)
source conf/app.env
docker system prune -f
chmod +x ./docker/run-init.sh
pushd `pwd`
cd docker && ./run-init.sh y
popd
docker-compose stop
isBuild=$1
isBuild=${isBuild:=n}
if [ $isBuild != y ]; then
    docker-compose up
else
    docker-compose up -d --build
fi