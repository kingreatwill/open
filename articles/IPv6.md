https://www.toutiao.com/c/user/105575985919/#mid=1615174441521156


# IP

## IP v4
ip

## IP v6

## 其它
[获取外网IP](http://ip.chinaz.com/)

### IPDB格式地级市精度IP离线库/IPIP.net
https://www.ipip.net/product/ip.html
https://www.ipip.net/support/code.html

https://github.com/metowolf/qqwry.ipdb

#### IPdb格式转换
https://github.com/sjzar/ips
https://github.com/qiniu/uip

#### DATX 格式

### APNIC
APNIC 每天更新各国 IP 分配清单，地址永久免费开放：http://ftp.apnic.net/apnic/stats/apnic/delegated-apnic-latest


关于APNIC

全球IP地址块被IANA(Internet Assigned Numbers Authority)分配给全球三大地区性IP地址分配机构，它们分别是：

ARIN (American Registry for Internet Numbers)
负责北美、南美、加勒比以及非洲撒哈啦部分的IP地址分配。同时还要给全球NSP(Network Service Providers)分配地址。

RIPE (Reseaux IP Europeens)
负责欧洲、中东、北非、西亚部分地区(前苏联)

APNIC (Asia Pacific Network Information Center)
负责亚洲、太平洋地区

APNIC IP地址分配信息总表的获取：
APNIC提供了每日更新的亚太地区IPv4，IPv6，AS号分配的信息表：http://ftp.apnic.net/apnic/stats/apnic/delegated-apnic-latest
该文件的格式与具体内容参见：ftp://ftp.apnic.net/pub/apnic/stats/apnic/README.TXT

```
#!/bin/bash
wget -c http://ftp.apnic.net/stats/apnic/delegated-apnic-latest

cat delegated-apnic-latest | awk -F '|' '/CN/&&/ipv4/ {print $4 "/" 32-log($5)/log(2)}' | cat > ipv4.txt

cat delegated-apnic-latest | awk -F '|' '/CN/&&/ipv6/ {print $4 "/" 32-log($5)/log(2)}' | cat > ipv6.txt

cat delegated-apnic-latest | awk -F '|' '/HK/&&/ipv4/ {print $4 "/" 32-log($5)/log(2)}' | cat > ipv4-hk.txt

cat delegated-apnic-latest | awk -F '|' '/HK/&&/ipv6/ {print $4 "/" 32-log($5)/log(2)}' | cat > ipv6-hk.txt
```

#### nginx拒绝境外IP

创建自动转换脚本（保存为 /usr/local/bin/gen-cn-allow.sh）：
```
#!/bin/bash
# 从 APNIC 官方数据生成 Nginx allow 规则
# 适用于任意 Linux 系统（CentOS/Ubuntu/Debian/Alma/Rocky 等）

OUTPUT_DIR="/etc/nginx/conf.d"
mkdir -p "$OUTPUT_DIR"

echo "正在下载 APNIC 最新数据..."
wget -qO- http://ftp.apnic.net/apnic/stats/apnic/delegated-apnic-latest | \
awk -F'|' '
  $2 == "CN" && $3 == "ipv4" {
    prefix = $4;
    len = 32 - log($5) / log(2);
    print "allow " prefix "/" len ";";
  }
  $2 == "CN" && $3 == "ipv6" {
    print "allow " $4 "/" $5 ";";
  }
' > /tmp/cn_allow.list

# 分离 IPv4 和 IPv6（避免混合导致 Nginx 报错）
grep -E 'allow [0-9]+\.' /tmp/cn_allow.list > "$OUTPUT_DIR/china-ipv4.conf"
grep -E 'allow [0-9a-fA-F:]+' /tmp/cn_allow.list > "$OUTPUT_DIR/china-ipv6.conf"

# 添加注释头
sed -i "1i# Auto-generated from APNIC — $(date)" "$OUTPUT_DIR/china-ipv4.conf"
sed -i "1i# Auto-generated from APNIC — $(date)" "$OUTPUT_DIR/china-ipv6.conf"

rm -f /tmp/cn_allow.list
echo "✅ 中国 IP 白名单已生成："
echo "   IPv4: $OUTPUT_DIR/china-ipv4.conf"
echo "   IPv6: $OUTPUT_DIR/china-ipv6.conf"
```
执行
```
chmod +x /usr/local/bin/gen-cn-allow.sh
sudo /usr/local/bin/gen-cn-allow.sh
```
定时执行
```
echo "0 3 * * * root /usr/local/bin/gen-cn-allow.sh >/dev/null 2>&1" | sudo tee /etc/cron.d/update-cn-ip
```

配置 Nginx 仅放行中国 IP
```
server {
    listen 80;
    listen [::]:80;
    server_name your-domain.com;

    # 引入中国 IP 白名单（顺序很重要！）
    include /etc/nginx/conf.d/china-ipv4.conf;
    include /etc/nginx/conf.d/china-ipv6.conf;

    # 可选：放行本地回环（避免自己被拦）
    allow 127.0.0.1;
    allow ::1;

    # 拒绝所有未匹配的请求
    deny all;

    location / {
        root /var/www/html;
        index index.html;
        # 你的其他配置...
    }
}
```

### 全球IP属地/离线IP库
https://github.com/lionsoul2014/ip2region
```go
package main

// go install github.com/lionsoul2014/ip2region/binding/golang/xdb
import (
    "fmt"
    "github.com/lionsoul2014/ip2region/binding/golang/xdb"
    "time"
)

var (
    // 下载的数据库文件放在项目根目录
    dbPath string = "./ip2region.xdb"
    ipBuff []byte
)

func init() {
    var err error
    ipBuff, err = xdb.LoadContentFromFile(dbPath)
    if err != nil {
        fmt.Printf("加载数据库数据失败 `%s`: %s\n", dbPath, err)
        return
    }
}

func main() {
    searcher, err := xdb.NewWithBuffer(ipBuff)
    if err != nil {
        fmt.Printf("创建searcher失败: %s\n", err.Error())
        return
    }

    defer searcher.Close()


    var ip = "111.128.0.0"
    var startTime = time.Now()
    region, err := searcher.SearchByStr(ip)
    if err != nil {
        fmt.Printf("查询ip失败(%s): %s\n", ip, err)
        return
    }

    fmt.Printf("addr: %s, took: %s\n", region, time.Since(startTime))
}
```

