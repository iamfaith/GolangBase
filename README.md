# GolangBase

[![Build Status](https://travis-ci.org/iamfaith/GolangBase.svg?branch=master)](https://travis-ci.org/iamfaith/GolangBase)

## Option 1

run godep save ./... first

## Option 2

1. chmod +x buildApp.sh && buildApp.sh (version_number) if not set version_number, will set 0.1 as default.


## upload interface

curl 127.0.0.1:8004/v1/file -X POST -F "file=@imagetool.zip"
curl faithio.cn:8004/v1/file -X POST -F "file=@imagetool.zip"