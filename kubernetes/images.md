1. 如何拉取 k8s.gcr.io 镜像

手动方式:
```
# 拉取新构建的镜像
docker pull registry.cn-shenzhen.aliyuncs.com/cookcodeblog/kube-apiserver-amd64:v1.10.3
# 打上gcr.io同名标签
docker tag registry.cn-shenzhen.aliyuncs.com/cookcodeblog/kube-apiserver-amd64:v1.10.3 k8s.gcr.io/kube-apiserver-amd64:v1.10.3
# 查看镜像
docker images
# 删除新构建的镜像，只保留gcr.io镜像
docker rmi registry.cn-shenzhen.aliyuncs.com/cookcodeblog/kube-apiserver-amd64:v1.10.3
# 再次查看镜像
docker images
```

脚本方式:
``` powershell
$images = New-Object -TypeName System.Collections.ArrayList
$images.Add("k8s.gcr.io/metrics-server-amd64:v0.3.1=registry.cn-hangzhou.aliyuncs.com/google_containers/metrics-server-amd64:v0.3.1")
foreach($line in $images) {
    $data = $line.Split('=')
    $key = $data[0];
    $value = $data[1];
    Write-Output "$key=$value"
    docker pull ${value}
    docker tag ${value} ${key}
    docker rmi ${value}
}
```