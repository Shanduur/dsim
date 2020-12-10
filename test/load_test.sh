#!/bin/bash

if [[ $1 == "-folder" ]]; then
    PHOTOFOLDER=$2
else 
    echo no folder provided 
    exit 1
fi

OPTIONS="-uname=user -pwd=password -log-level=1"

[ -d ./out/ ] && echo ok || mkdir ./out/
rm -rf ./out/*

[ -d ./log/ ] && echo ok || mkdir ./log/
rm -rf ./log/*

[ -d ./err/ ] && echo ok || mkdir ./err/
rm -rf ./err/*

for s1 in $(ls $PHOTOFOLDER | sort -R); do
    for s2 in $(ls $PHOTOFOLDER | sort -R); do
        echo $s1 $s2

        { time ./client.run \
            -source-img1=$PHOTOFOLDER/$s1 \
            -source-img2=$PHOTOFOLDER/$s2 \
            -o=./out/$s1$s2 \
            $OPTIONS >> ./err/$PHOTOFOLDER/err.txt 2>&1; } 2>> ./log/$PHOTOFOLDER/time.txt

        echo " --- " >> ./err/$PHOTOFOLDER/err.txt
    done
done
