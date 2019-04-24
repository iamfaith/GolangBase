# 标准镜像 + 初始化注入
FROM harbor.wps.kingsoft.net/library/mysql:v5.6
COPY conf/schema.sql /docker-entrypoint-initdb.d/
