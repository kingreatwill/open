
https://hub.docker.com/r/cloudnas/clouddrive
```
docker run -d --name clouddrive --restart unless-stopped -v /share/Public/dockerv/clouddrive/CloudNAS:/CloudNAS:shared  -v /share/Public/dockerv/clouddrive/Config:/Config -v /share/Public/dockerv/clouddrive/media:/media:shared --privileged --device /dev/fuse:/dev/fuse -p 9798:9798 cloudnas/clouddrive
```