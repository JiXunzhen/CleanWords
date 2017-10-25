#!/bin/sh

# 先更新敏感词
python3 ./sensitive_words/distinct.py

# 运行go脚本, 进行词过滤
go run main.go

# 汇聚执行结果
python3 ./seowords/join.py

echo "over."
