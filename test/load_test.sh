#!/bin/bash

[[ $1 == "-folder" ]] && PHOTOFOLDER=$2 || echo no folder provided && exit 1

OPTIONS="-o=./out/$(date '+%s') -uname=user -pwd=password -log-level=5"

[ -d ./out/ ] && echo ok || mkdir ./out/
rm -rf ./out/*

echo "" > time.txt

for query in $(ls $PHOTOFOLDER | sort -R); do
    for train in $(ls $PHOTOFOLDER | sort -R); do
        { time pluggabl client \
            -query=$query \
            -train=$train \
            $OPTIONS 2>&1; } 2>> time.txt
    done
done