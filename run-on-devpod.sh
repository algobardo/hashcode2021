#!/bin/bash

GOOS=linux GOARCH=amd64 go build src/main.go
rsync -avz . ${USER}.devpod-nld:hashcode/
ssh ${USER}.devpod-nld ./hashcode/main
