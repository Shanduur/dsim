#!/bin/bash
PHOTOFOLDER=$1

OPTIONS="-o=$(date '+%s') -uname=user -pwd=password -log-level=5"

echo "" > time.txt

for query in $(ls $PHOTOFOLDER | sort -R); do
    for train in $(ls $PHOTOFOLDER | sort -R); do
        { time pluggabl client \
            -query=$query \
            -train=$train \
            $OPTIONS 2>&1; } 2>> time.txt
    done
done