#!/bin/bash

if [ $# -gt 0 ] && [ $1 = "linux" ] ; then
    export CGO_ENABLED=0
    export GOOS=linux
    export GOARCH=amd64
    env
fi

go build

