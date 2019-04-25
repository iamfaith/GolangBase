#构建机在构建image的时候，需要  cp /home/user/.ssh/id_rsa ./
ADD id_rsa /root/.ssh/id_rsa
# 解决Windows下copy进去id_rsa文件权限过高
RUN chmod 0600 /root/.ssh/id_rsa


# Create known_hosts
RUN echo "ip xxxx.net" >> /etc/hosts \
    && touch /root/.ssh/known_hosts \
    && ssh-keyscan xxxx.net >> /root/.ssh/known_hosts \
    && cd $PROJECT_PATH \
    && /bin/sh docker/gobuilder.sh \
	&& rm -f /root/.ssh/id_rsa