#!/bin/bash
cp -r config/settings.yml.prod config/settings.yml && \
go build -a -o go-admin . && \
ps -ef |grep go-admin |grep -v grep|awk '{print $2}' |xargs -i kill -9 {} && \
nohup  ./go-admin &