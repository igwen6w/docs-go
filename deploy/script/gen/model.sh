#!/usr/bin/env bash

# 使用方法：
# ./model.sh database table
# mv ./model/* app/service/model


# 数据库配置
host=127.0.0.1
port=33069
dbname=looklook_$1
username=root
passwd=PXDN93VRKUm8TeE7

#生成的表名
tables=$2
#表生成的genmodel目录
modeldir=./model

cache=false
if [ $# -eq 2 ]; then
  cache=$3
fi

echo "开始创建库：$dbname 的表：$2"
# 可以使用-home指定模版路径，默认是~/.goctl
goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}"  -dir="${modeldir}" -cache="${cache}" --style=goZero
