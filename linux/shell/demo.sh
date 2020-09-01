#!/bin/bash
CRTDIR="xxx" #$(pwd)

function getDir() {
    for item in `ls $1`
    do
    fileName=$1"/"$item
    if [ -d $fileName ]
    then 
		echo $fileName
		cd $fileName
		git pull
		git reset --hard
		GIT_SSH_COMMAND="ssh -i xxrsa" git push
    fi
    done
}

getDir $CRTDIR

while true
do
ti1=`date +%M`    #获取时间戳
ti2=`date +%M`
i=$(($ti2 - $ti1 )) 
	while [[ "$i" -ne "10" ]] #等待10m执行下一条指令
	do
		ti2=`date +%M`
		i=$(($ti2 - $ti1 ))

	done
done