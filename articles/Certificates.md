easyrsa openssl cfssl 

CA 证书签名请求（CSR） CertificateSigningRequest
server.csr则含有公钥、组织信息、个人信息(域名)
server.key 私钥

.pem可以是私钥 也可以是

ca.pem	CA根证书文件
ca-key.pem	服务端私钥，用于对客户端请求的解密和签名
ca.csr	证书签名请求，用于交叉签名或重新签名


Privacy Enhanced Mail Certificate


是服务端的私钥


CA, Certificate Authority 证书颁发机构