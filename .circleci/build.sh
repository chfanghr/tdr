#!/bin/bash

sudo mkdir /build
sudo chmod +666 /build
GO111MODULE=on go get -d ./...

go_build_target(){
GO111MODULE=on GOOS=$1 GOARCH=$2 go build -o /build/$1_$2 $3
}

go_build_target "linux" "amd64" $1
go_build_target "linux" "386" $1
go_build_target "linux" "arm" $1
go_build_target "linux" "arm64" $1
go_build_target "windows" "amd64" $1
go_build_target "windows" "386" $1
go_build_target "darwin" "amd64" $1
go_build_target "darwin" "386" $1

#tar -cf /tmp/build.tar /build
#mv /tmp/build.tar /build

for i in /build/windows*;do mv ${i} ${i}.exe;done