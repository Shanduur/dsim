#!/bin/bash

if [[ $1 == "-folder" ]]; then
    PHOTOFOLDER=$2
else 
    echo no folder provided 
    exit 1
fi

OPTIONS="-o=./out/$(date '+%s') -uname=user -pwd=password -log-level=5"

[ -d ./out/ ] && echo ok || mkdir ./out/
rm -rf ./out/*

[ -d ./log/ ] && echo ok || mkdir ./log/
rm -rf ./log/*

[ -d ./err/ ] && echo ok || mkdir ./err/
rm -rf ./err/*

for s1 in $(ls $PHOTOFOLDER | sort -R); do
    for s2 in $(ls $PHOTOFOLDER | sort -R); do
        { time ./client.run \
            -source-img1=$PHOTOFOLDER/$s1 \
            -source-img2=$PHOTOFOLDER/$s2 \
            $OPTIONS >> ./err/err.txt 2>&1; } 2>> ./log/time.txt
    done
done
