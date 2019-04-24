#!/bin/sh

export PROJECT_ROOT=`pwd`


if [ ! -h ${GOPATH}/src/${PROJECT_NAME} ]; then
	mkdir -p ${GOPATH}/src
	ln -s "${PROJECT_ROOT}" ${GOPATH}/src/${PROJECT_NAME}
fi

ls -s dependent/vendor ${GOPATH}/src/vendor

ls ${GOPATH}/src

echo $GOPATH/src/${PROJECT_NAME}--${DEPENDENT}

cd $GOPATH/src/${PROJECT_NAME} && go install .