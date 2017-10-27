#!/bin/sh

path=$1
if [ ! -d $path ]; then
    echo "no such dir! $path"
    exit 1
fi

# 运行go脚本, 进行词过滤
go run main.go $path

# 汇聚执行结果
python3 $path/join.py

echo "========================="
wc -l $path/result.csv
echo "========================="
