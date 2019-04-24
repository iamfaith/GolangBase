
FROM xianzixiang/go_builder as builder

ARG PROJECT_NAME
ARG PROJECT_PATH


ADD . $PROJECT_PATH

#构建机在构建image的时候，需要  cp /home/kingsoft/.ssh/id_rsa ./
ADD id_rsa /root/.ssh/id_rsa
# 解决Windows下copy进去id_rsa文件权限过高
RUN chmod 0600 /root/.ssh/id_rsa

# Create known_hosts
RUN echo "10.13.0.32 ksogit.kingsoft.net" >> /etc/hosts \
    && touch /root/.ssh/known_hosts \
    && ssh-keyscan ksogit.kingsoft.net >> /root/.ssh/known_hosts \
    && cd $PROJECT_PATH \
    && /bin/sh docker/gobuilder.sh \
	&& rm -f /root/.ssh/id_rsa


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
