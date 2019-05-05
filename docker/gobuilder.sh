#!/bin/sh


export PROJECT_ROOT=`pwd`


if [ ! -h ${GOPATH}/src/${PROJECT_NAME} ]; then
	mkdir -p ${GOPATH}/src
	ln -s "${PROJECT_ROOT}" ${GOPATH}/src/${PROJECT_NAME}
fi

cd $GOPATH/src/${PROJECT_NAME}

if [ -d  ${GOPATH}/src/${PROJECT_NAME}/vendor/vendor ]; then
    echo link vendor
    ln -s ${GOPATH}/src/${PROJECT_NAME}/vendor/vendor $GOPATH/src
else
    go get -v .
fi

go install .

