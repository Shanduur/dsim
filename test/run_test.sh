#!/bin/bash

if [[ $1 == "-threads" ]]; then
    THREADS=$2
else 
    echo no thread number provided 
    exit 1
fi

FULLSET=oxbuild_images
BASEDIR=images

[ -d ./log/ ] && echo ok || mkdir ./log/
rm -rf ./log/*
rm -f nohup.out

for ((i=1;i<=THREADS;i++)); do
    if ((i < 10)); then
        PREFIX="dir_00"
    elif ((i < 100)); then
        PREFIX="dir_0"
    else
        echo failed to count on THREADS
        exit 1
    fi
    nohup time ./load_test.sh ./$FULLSET ./$BASEDIR/$PREFIX$i >/dev/null 2> ./log/time_$i.log &
done