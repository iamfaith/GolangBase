# 标准镜像 + 初始化注入
FROM xianzixiang/mysql
COPY conf/schema.sql /docker-entrypoint-initdb.d/
