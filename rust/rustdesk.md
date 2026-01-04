
### 防火墙
确保在防火墙中打开这些端口：

hbbs:
21114 (TCP): 用于网页控制台，仅在 Pro 版本中可用。
21115 (TCP): 用于 NAT 类型测试。
21116 (TCP/UDP): 请注意 21116 应该同时为 TCP 和 UDP 启用。 21116/UDP 用于 ID 注册和心跳服务。21116/TCP 用于 TCP 打洞和连接服务。
21118 (TCP): 用于支持网页客户端。
hbbr:
21117 (TCP): 用于中继服务。
21119 (TCP): 用于支持网页客户端。
如果您不需要网页客户端支持，可以禁用相应的端口 21118、21119。

①ID服务器：[替换成云服务公网地址]:21116。
②中继服务器：[替换成云服务公网地址]:21117

43.134.15.151:21114

### docker部署
sudo docker image pull rustdesk/rustdesk-server
sudo docker run --name hbbs -v /data/dockerv/rustdesk:/root -td --net=host -e ALWAYS_USE_RELAY=Y --restart unless-stopped rustdesk/rustdesk-server hbbs
sudo docker run --name hbbr -v /data/dockerv/rustdesk:/root -td --net=host --restart unless-stopped rustdesk/rustdesk-server hbbr

