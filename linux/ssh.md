
## ssh连接
一、正常连接方法：ssh root@10.0.0.20

二、无密码连接方法(有两台机器：此处我把被连接的称为服务器，另一台则称为客户端):

　　1、先在服务器添加目录 .ssh： mkdir  .ssh

　　2、分配.ssh目录权限： chmod 777 .ssh

　　3、在客户端创建公钥与私钥： ssh-keygen　　//此处直接按多个回车键，直到创建成功

　　4、将客户端的公钥复制到要服务器，运行命令：ssh-copy-id root@10.0.0.20 ，待输入正确密码后即可实现ssh无密码访问。

三、Win没有ssh-copy-id
在服务器端
```
mkdir ~/.ssh
chmod 0700 ~/.ssh
touch ~/.ssh/authorized_keys
chmod 0644 ~/.ssh/authorized_keys
nano ~/.ssh/authorized_keys     # Ctrl+O 保存   Ctrl+X 退出
```
在win上找到ssh-keygen生成的.pub文件 放到~/.ssh/authorized_keys文件里

