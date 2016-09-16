#!/bin/bash

# 源于网络，用于获取当前shell文件的路径
SOURCE="$0"
while [ -h "$SOURCE"  ]; do 
    DIR="$( cd -P "$( dirname "$SOURCE"  )" && pwd  )"
    SOURCE="$(readlink "$SOURCE")"
    [[ $SOURCE != /*  ]] && SOURCE="$DIR/$SOURCE" 
done
DIR="$( cd -P "$( dirname "$SOURCE"  )" && pwd  )"

cd "$DIR/../wrapper/"
make -f "$DIR/../wrapper/Makefile" clean
cd "$DIR"

go build -buildmode=c-archive -gcflags=-shared -asmflags=-shared -installsuffix=_shared -a -o "$DIR/../wrapper/libprotorpc.a" "$DIR/../lib.go"

cd "$DIR/../wrapper/"
make -f "$DIR/../wrapper/Makefile"
cd "$DIR"

php -S localhost:8000 -t "$DIR/../examples" -d "extension=$DIR/../wrapper/modules/protorpc_php.so"
#php -d "extension=$DIR/../wrapper/modules/protorpc_php.so" "$DIR/../examples/test.php"