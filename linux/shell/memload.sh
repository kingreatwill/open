#!/bin/bash

if [ $# != 2 ]; then
        echo "USAGE: $0 <free,sleep time>"
        exit 1;
fi 
free -m > /tmp/freee
cat /tmp/freee
mkdir /tmp/memory
mount -t tmpfs -o size=$1M tmpfs /tmp/memory
dd if=/dev/zero of=/tmp/memory/block
free -m > /tmp/freee
cat /tmp/freee
echo "If executed ctrl+C,Please execute the following lines manually"
echo "execute: rm -rf /tmp/memory/block"
echo "execute: umount /tmp/memory"
echo "execute: rmdir /tmp/memory"
echo "Please wait $2 seconds" 
sleep $2
rm -rf /tmp/memory/block
umount /tmp/memory
rmdir /tmp/memory

# sh memload.sh 2048 120【2048是内存数量，单位是M，120是时间，单位是秒】
# 内存使用情况:free -m (以MB为单位显示内存使用情况。)