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

