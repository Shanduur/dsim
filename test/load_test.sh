#!/bin/bash

FULLSET=$1
PART=$2

OPTIONS="-uname=user -pwd=password -log-level=1"

[ -d ./out/ ] && echo ok || mkdir ./out/
rm -rf ./out/*

[ -d ./err/ ] && echo ok || mkdir ./err/
rm -rf ./err/*

for s1 in $(ls $FULLSET | sort -R); do
    for s2 in $(ls $PART | sort -R); do
        ./client.run \
            -source-img1=$FULLSET/$s1 \
            -source-img2=$PART/$s2 \
            -o=./out/$s1$s2 \
            $OPTIONS >> ./err/err.txt 2>&1
    done
done
