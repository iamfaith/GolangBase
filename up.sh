#!/bin/bash

export $(cut -d= -f1 conf/app.env)
source conf/app.env
docker-compose up --build