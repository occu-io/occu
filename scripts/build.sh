#!/bin/bash

BUILD=( 
#	"darwin_amd64"
	"linux_amd64"
#	"windows_386"
#	"windows_amd64"
#	"freebsd_amd64"
#	"solaris_amd64"
)

BUILD_VERSION="1.0"
BUILD_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

for i in "${BUILD[@]}"
do
	os=$(echo $i | awk -F '_' '{print $1}')
	arch=$(echo $i | awk -F '_' '{print $2}')
	echo "[+] building $os/$arch"
	GOPATH=$GOPATH:~/go:$BUILD_DIR/.. GOOS=$os GOARCH=$arch go build -o $BUILD_DIR/../bin/occu-$BUILD_VERSION-$i $BUILD_DIR/../src/main.go
done
echo "[=] done"
