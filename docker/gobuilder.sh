#!/bin/sh


export PROJECT_ROOT=`pwd`


if [ ! -h ${GOPATH}/src/${PROJECT_NAME} ]; then
	mkdir -p ${GOPATH}/src
	ln -s "${PROJECT_ROOT}" ${GOPATH}/src/${PROJECT_NAME}
fi

cd $GOPATH/src


git clone git@ksogit.kingsoft.net:mo_server/vendor.git


cd ${PROJECT_NAME} && go install .


