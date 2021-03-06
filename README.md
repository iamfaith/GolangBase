# GolangBase

[![Build Status](https://travis-ci.org/iamfaith/GolangBase.svg?branch=master)](https://travis-ci.org/iamfaith/GolangBase)

## Option 1

run godep save ./... first

## Option 2

1. cd build && chmod +x buildApp.sh && buildApp.sh (version_number) if not set version_number, will set 0.1 as default.


## check alive

 curl 127.0.0.1:8004/alive

## upload interface

echo aa >> test.txt && curl 127.0.0.1:8004/api/v1/file?uname=1 -X POST -F "file=@test.txt"
curl faithio.cn:8004/api/v1/file -X POST -F "file=@test.txt"

## get reflect interface

```

 curl 127.0.0.1:8004/api/v1/GetValue/upload_file4e40da587ba423e49862a841798f700220543880?uname=1

 curl faithio.cn:8004/api/v1/ListAll/upload_file?uname=1

 curl 127.0.0.1:8004/api/v1/FindLinkByUid/11?uname=1

 curl "127.0.0.1:8004/api/v1/GetAll/Link?uname=1&t=all"
 curl "faithio.cn:8004/api/v1/GetAll/Link?uname=1&t=all"

```

## post reflect interface

```
  curl "127.0.0.1:8004/api/v1/Insert?uname=1" -X POST -d '{"uid": "curl", "content":"faith", "tbl": "link"}'

```

## redis-cluster

127.0.0.1:7000,127.0.0.1:7001,127.0.0.1:7002,127.0.0.1:7003,127.0.0.1:7004,127.0.0.1:7005,127.0.0.1:7006,127.0.0.1:7007

## pprof

```

curl 127.0.0.1:8004/_adm/pprof/cpu?seconds=2 -H "adminKey:abcdefg" -o cpu

 brew install  Graphviz

go tool pprof -svg heap >heap.svg

```