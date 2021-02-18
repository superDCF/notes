#!/bin/bash

## dev for macos
set -u
IMAGE=golang:1.15
path=`pwd`
echo $path
project=`basename $path`
echo $project

NAME=$project

BUILD_DIR=`pwd`
BUILD_DIR_BIN=$BUILD_DIR
BUILD_DIR_PKG=$HOME/go/pkg
BUILD_DIR_CACHE=$HOME/.cache

docker run -i --rm \
  -e GO111MODULE=on \
  -e GOPRIVATE="*.hellotalk.com" \
  -e GOPROXY=https://goproxy.cn \
  -e CGO_ENABLED=1 \
  -e GOOS=linux \
  -v ~/.ssh:/root/.ssh \
  -v ~/.gitconfig:/root/.gitconfig \
  -v $BUILD_DIR_PKG:/go/pkg \
  -v $BUILD_DIR_CACHE:/.cache \
  -v $BUILD_DIR:/go/src \
  -w /go/src \
  $IMAGE make project=$project

NOW=`date +'%Y_%m_%d_%H_%M_%S'`
if test -f "$BUILD_DIR_BIN/$NAME"; then
    cd $BUILD_DIR_BIN
    mv $NAME ${NAME}_${NOW}
    ln -s ${NAME}_${NOW} $NAME
    tar -zcf $NAME.tar.gz ${NAME}_${NOW} $NAME
    rm -f ${NAME}_${NOW} $NAME
fi