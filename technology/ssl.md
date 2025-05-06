
## 域名
[免费的域名](http://www.freenom.com/en/index.html)
## SSL
[11种免费获取SSL证书的方式](https://www.toutiao.com/i6883395048126284292)


### 自己生成证书
1. Certbot
```mac
brew install certbot

sudo certbot certonly --standalone -d yourdomain.com -d www.yourdomain.com
#需要域名已解析到本机公网 IP。
#证书会生成在 /etc/letsencrypt/live/yourdomain.com/ 目录下。
```

2. golang
```
package main

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "crypto/x509/pkix"
    "encoding/pem"
    "math/big"
    "os"
    "time"
)

func main() {
    // 生成私钥
    priv, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        panic(err)
    }

    // 证书模板
    template := x509.Certificate{
        SerialNumber: big.NewInt(1),
        Subject: pkix.Name{
            Organization: []string{"My Company"},
        },
        NotBefore: time.Now(),
        NotAfter:  time.Now().Add(365 * 24 * time.Hour), // 有效期一年

        KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
        ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
        BasicConstraintsValid: true,
    }

    // 生成自签名证书
    derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
    if err != nil {
        panic(err)
    }

    // 写入证书文件
    certOut, err := os.Create("cert.pem")
    if err != nil {
        panic(err)
    }
    pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
    certOut.Close()

    // 写入私钥文件
    keyOut, err := os.Create("key.pem")
    if err != nil {
        panic(err)
    }
    pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
    keyOut.Close()

    println("证书和私钥已生成：cert.pem, key.pem")
}
```

### ACME 协议

ACME (自动证书管理环境 - Automatic Certificate Management Environment) 
ACME 客户端 (ACME Client) : 能够与 ACME 服务器通信以获取证书的程序。

Let’s Encrypt 使用 ACME 协议来验证您对给定域名的控制权并向您颁发证书。要获得 Let’s Encrypt 证书，您需要选择一个要使用的 ACME 客户端软件。

https://letsencrypt.org/zh-cn/docs/client-options/

https://github.com/letsencrypt/website/

#### golang.org/x/crypto/acme/autocert
> https://github.com/golang/crypto/blob/master/acme/autocert/autocert.go
```
package main

import (
    "log"
    "net/http"
    "golang.org/x/crypto/acme/autocert"
)

func main() {
    m := &autocert.Manager{
        Cache:      autocert.DirCache("certs"), // 证书缓存目录
        Prompt:     autocert.AcceptTOS,
        HostPolicy: autocert.HostWhitelist("yourdomain.com"), // 你的域名
    }

    server := &http.Server{
        Addr:      ":443",
        Handler:   http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Hello, HTTPS!")) }),
        TLSConfig: m.TLSConfig(),
    }

    // 自动监听 80 端口用于 ACME 挑战
    go http.ListenAndServe(":80", m.HTTPHandler(nil))

    log.Fatal(server.ListenAndServeTLS("", "")) // 证书自动管理
}
```

#### https://github.com/go-acme/lego

#### certbot - ACME 客户端
[让网站永久拥有HTTPS - 申请免费SSL证书并自动续期](https://blog.csdn.net/xs18952904/article/details/79262646)
https://www.cnblogs.com/dissipate/p/13606006.html
http://www.mntm520.com/post/48
也可以用certbot的docker镜像来生成证书


#### zerossl - ACME 客户端
https://zerossl.com/

#### acme.sh

这个脚本acme.sh，实现了 acme 协议, 可以帮你持续自动从Letsencrypt更新CA证书
https://github.com/acmesh-official/acme.sh

安装
https://github.com/acmesh-official/acme.sh/wiki/%E8%AF%B4%E6%98%8E
`curl get.acme.sh | sh`

#### SSLforFree
一个免费的SSL证书申请服务，但自动化程度较低。