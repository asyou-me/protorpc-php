#!/bin/bash

# 源于网络，用于获取当前shell文件的路径
SOURCE="$0"
while [ -h "$SOURCE"  ]; do 
    DIR="$( cd -P "$( dirname "$SOURCE"  )" && pwd  )"
    SOURCE="$(readlink "$SOURCE")"
    [[ $SOURCE != /*  ]] && SOURCE="$DIR/$SOURCE" 
done
DIR="$( cd -P "$( dirname "$SOURCE"  )" && pwd  )"

make -f "$DIR/Makefile"
php -S localhost:8000 -t "$DIR/../examples" -d "extension=$DIR/../_out/protorpc-php.so"