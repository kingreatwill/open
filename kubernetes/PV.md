# pv&pvc简介

k8s提供了emptyDir,hostPath,rbd,cephfs等存储方式供容器使用,不过这些存储方式都有一个缺点:开发人员必须得知指定存储的相关配置信息,才能使用存储.例如要使用cephfs,Pod的配置信息就必须指明cephfs的monitor,user,selectFile等等,而这些应该是系统管理员的工作.对此,k8s提供了两个新的API资源:PersistentVolume,PersistentVolumeClaim

PV(PersistentVolume)是管理员已经提供好的一块存储.在k8s集群中,PV像Node一样,是一个资源

PVC(PersistentVolumeClaim)是用户对PV的一次申请.PVC对于PV就像Pod对于Node一样,Pod可以申请CPU和Memory资源,而PVC也可以申请PV的大小与权限

有了PersistentVolumeClaim,用户只需要告诉Kubernetes需要什么样的存储资源,而不必关心真正的空间从哪里分配,如何访问等底层细节信息;这些Storage Provider的底层信息交给管理员来处理,只有管理员才应该关心创建PersistentVolume的细节信息

192.168.1.135 master节点的ip

# 创建nfs类型的PV
1. master安装nfs服务端
```
yum install nfs-utils rpcbind
```
2. 创建用于nfs服务的存储文件夹
```
mkdir -p /data/nfs
```
3. 编辑/etc/exports文件
```
vi /etc/exports
/data/nfs 192.168.1.0/24(rw,sync)
```
4. 开启nfs服务
```
systemctl start nfs-server.service
```
5. 检验是否开启成功
```
[root@node1 ~]# showmount -e
Export list for node1.centos7.com:
/data/nfs 192.168.1.0/24

```

# 创建PV
```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: nfspv1
spec:
  #指定pv的容量为10Gi
  capacity:
    storage: 10Gi
  #指定访问模式
  accessModes:
    #pv能以readwrite模式mount到单个节点
    - ReadWriteOnce
    # 该volume只能被多个节点以读写的方式映射
    - ReadWriteMany
    # 该volume可以被多个节点以只读方式映射
    - ReadOnlyMany
  #指定pv的回收策略,即pvc资源释放后的事件.recycle(不建议,使用动态供给代替)删除pvc的所有文件
  # persistentVolumeReclaimPolicy: Recycle
  #指定pv的class为nfs,相当于为pv分类,pvc将指定class申请pv
  storageClassName: mynfs
  #指定pv为nfs服务器上对应的目录
  nfs:
    path: /data/nfs
    server: 192.168.1.135
```

```
kubectl apply -f pv_nfs.yml

kubectl get pv
```


# 排错
在这次实验中出现了几个错误:

1. 使用pvc存储的pod一直处于pending状态,无法启动
kubectl describe pod xxx -n xxx
注意到报错信息:wrong fs type, bad option, bad superblock on 192.168.122.10:/data/nfs,这是因为host3没有安装nfs-utils软件包,无法识别nfs类型的文件系统,也无法作为nfs的客户端使用

解决方案:安装nfs-utils软件包,删除之前创建失败的Job资源并重新创建

yum install nfs-utils -y

2. 使用pvc的job执行失败
[kube@host1 ~]$ kubectl logs pvjob-74g64
/bin/sh: can't create /mydata/hello: Permission denied
解决方案:

Permission denied多见于普通用户执行高权限命令失败,不过busybox容器本身使用的就是root用户,因此不存在这个问题.在nfs中,nfs服务端没有权限访问挂载的目录也会导致这个问题

更改目录属主为nfsnobody:

chown nfsnobody /data/nfs
