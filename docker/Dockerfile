
FROM xianzixiang/go_builder:1.0 as builder

ARG PROJECT_NAME
ARG PROJECT_PATH


ADD . $PROJECT_PATH


RUN cd $PROJECT_PATH \
    && /bin/sh docker/gobuilder.sh


FROM xianzixiang/go_alpine:latest


ARG PROJECT_NAME

ARG PROJECT_PATH

RUN mkdir -p ${PROJECT_PATH}/bin && mkdir -p /etc/supervisord/ && mkdir -p /data

COPY --from=builder /go/bin/$PROJECT_NAME ${PROJECT_PATH}/bin
COPY --from=builder ${PROJECT_PATH}/conf ${PROJECT_PATH}/conf
COPY --from=builder ${PROJECT_PATH}/docker ${PROJECT_PATH}/docker


# Run the outyet command by default when the container starts.
ENTRYPOINT ${PROJECT_PATH}/bin/$PROJECT_NAME -sp ${PROJECT_PATH}


# 跟项目配置的端口一致
EXPOSE 8004:8004

#docker build -t docker_test:v0.1 .
#docker run --publish 8080:8080 --name docker_test --rm docker_test:v0.1
