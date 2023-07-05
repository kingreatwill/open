#!/bin/bash  
rm -rf /root/file.txt  
endless_loop()  
{  
        echo -ne "i=0; 
while true ;do 
        i=i+100; 
        i=100;done" | /usr/bin/bash &  
}  
if [ $# != 2 ]; then  
        echo "USAGE: $0 <cpus,sleep time>"  
        exit 1;  
fi  
  
for i in `seq $1`  
do  
        endless_loop  
        pid_array[$i]=$!;  
done  
  
for i in "${pid_array[@]}"; do  
        echo 'execute: kill' $i ;  
        echo 'kill' $i  >> /root/file.txt 
done
echo "If executed ctrl+C,Please execute the above lines manually"
echo "Please wait $2 seconds" 
sleep $2 
for i in `awk '{print $2}' /root/file.txt` ; do
echo "kill $i"
kill $i
done

# 参数:cpuload.sh 2 10
# 2代表占用2颗CPU【被占用的cpu使用率会是100%】,占用10S时间【单位是秒】
# 查看CPU核数: `lscpu | grep CPU\(s\):`