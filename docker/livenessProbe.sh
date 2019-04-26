#!/bin/bash
# 健康检测, 大概每10~15秒调用一次
#
#

HEALTHY_URL="http://faithio.cn:8004/alive"

RESP=`curl --connect-timeout 1 -s ${HEALTHY_URL}`

if [ "$RESP" == "ok" ]; then
    echo "Server is alive."
    exit 0 # 服务正常，脚本退出: 0
else
    echo "Server is die."
    exit 1 # 脚本退出为非0,代表服务不可用
fi