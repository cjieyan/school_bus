#!/bin/bash
cp -r config/settings.yml.prod config/settings.yml && \
go build -a -o go-admin . && \
nohup  ./go-admin &