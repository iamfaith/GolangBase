#!/bin/sh

set -e

cmd="$@"

mysql_ip=`dig +short ${MYSQL_HOST}`
until timeout -t 1 bash -c "</dev/tcp/${mysql_ip}/3306"; do
  >&2 echo "waiting for mysql to be available"
  sleep 1
done

>&2 echo "deps are all up - executing command"
exec $cmd
