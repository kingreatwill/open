

## Wiki.js 2 postgresql
docker run -d -p 5432:5432 -v /data/dockerv/postgresql13/data:/var/lib/postgresql/data -e POSTGRES_PASSWORD=jw@zll --name postgresql --restart always postgres:13

docker run -d -p 3000:3000 --name wiki --restart always --link postgresql:postgresql -e "DB_TYPE=postgres" -e "DB_HOST=postgresql" -e "DB_PORT=5432" -e "DB_USER=postgres" -e "DB_PASS=jw@zll" -e "DB_NAME=wiki" requarks/wiki:2.5


### 手动重置管理员密码
1. 生成HASH-PASSWORD, https://bcrypt-generator.com/ (Rounds必须是12)
2. `UPDATE users SET password = 'HASH-PASSWORD' WHERE email = 'YOUR-EMAIL';`

> https://docs.requarks.io/en/troubleshooting

## Wiki.js 2 sqlite

docker run -d -p 3000:3000 --name wiki --restart always -v /d/dockerv/wikijs2/config.yml:/wiki/config.yml -v /d/dockerv/wikijs2/data/:/wiki/db/ requarks/wiki:2.5

> Wiki.js runs as user wiki. docker启动时可以指定-u="root"
升级只需要更新docker镜像版本就可以了.

[config.yml](https://docs.requarks.io/install/config)
```
#######################################################################
# Wiki.js - CONFIGURATION                                             #
#######################################################################
# Full documentation + examples:
# https://docs.requarks.io/install

# ---------------------------------------------------------------------
# Port the server should listen to
# ---------------------------------------------------------------------

port: 3000

# ---------------------------------------------------------------------
# Database
# ---------------------------------------------------------------------
# Supported Database Engines:
# - postgres = PostgreSQL 9.5 or later
# - mysql = MySQL 8.0 or later (5.7.8 partially supported, refer to docs)
# - mariadb = MariaDB 10.2.7 or later
# - mssql = MS SQL Server 2012 or later
# - sqlite = SQLite 3.9 or later

db:
  type: sqlite
  # PostgreSQL / MySQL / MariaDB / MS SQL Server only:
  host: localhost
  port: 5432
  user: wikijs
  pass: wikijsrocks
  db: wiki
  ssl: false

  # Optional - PostgreSQL / MySQL / MariaDB only:
  # -> Uncomment lines you need below and set `auto` to false
  # -> Full list of accepted options: https://nodejs.org/api/tls.html#tls_tls_createsecurecontext_options
  sslOptions:
    auto: true
    # rejectUnauthorized: false
    # ca: path/to/ca.crt
    # cert: path/to/cert.crt
    # key: path/to/key.pem
    # pfx: path/to/cert.pfx
    # passphrase: xyz123

  # SQLite only:
  storage: /wiki/db/database.sqlite

#######################################################################
# ADVANCED OPTIONS                                                    #
#######################################################################
# Do not change unless you know what you are doing!

# ---------------------------------------------------------------------
# SSL/TLS Settings
# ---------------------------------------------------------------------
# Consider using a reverse proxy (e.g. nginx) if you require more
# advanced options than those provided below.

ssl:
  enabled: false
  port: 3443

  # Provider to use, possible values: custom, letsencrypt
  provider: custom

  # ++++++ For custom only ++++++
  # Certificate format, either 'pem' or 'pfx':
  format: pem
  # Using PEM format:
  key: path/to/key.pem
  cert: path/to/cert.pem
  # Using PFX format:
  pfx: path/to/cert.pfx
  # Passphrase when using encrypted PEM / PFX keys (default: null):
  passphrase: null
  # Diffie Hellman parameters, with key length being greater or equal
  # to 1024 bits (default: null):
  dhparam: null

  # ++++++ For letsencrypt only ++++++
  domain: wiki.yourdomain.com
  subscriberEmail: admin@example.com

# ---------------------------------------------------------------------
# Database Pool Options
# ---------------------------------------------------------------------
# Refer to https://github.com/vincit/tarn.js for all possible options

pool:
  # min: 2
  # max: 10

# ---------------------------------------------------------------------
# IP address the server should listen to
# ---------------------------------------------------------------------
# Leave 0.0.0.0 for all interfaces

bindIP: 0.0.0.0

# ---------------------------------------------------------------------
# Log Level
# ---------------------------------------------------------------------
# Possible values: error, warn, info (default), verbose, debug, silly

logLevel: info

# ---------------------------------------------------------------------
# Offline Mode
# ---------------------------------------------------------------------
# If your server cannot access the internet. Set to true and manually
# download the offline files for sideloading.

offline: false

# ---------------------------------------------------------------------
# High-Availability
# ---------------------------------------------------------------------
# Set to true if you have multiple concurrent instances running off the
# same DB (e.g. Kubernetes pods / load balanced instances). Leave false
# otherwise. You MUST be using PostgreSQL to use this feature.

ha: false

# ---------------------------------------------------------------------
# Data Path
# ---------------------------------------------------------------------
# Writeable data path used for cache and temporary user uploads.
dataPath: ./data
```
## Wiki.js 3
将使用Quasar Vue framework
支持PDF导出
将PostgreSQL 作为唯一的数据库
[文档](https://next.js.wiki/docs/install/platform/docker)

```

docker run -d -p 10002:3000 --name wiki --link postgresql --restart unless-stopped -e "DB_HOST=postgresql" -e "DB_PORT=5432" -e "DB_USER=postgres" -e "DB_PASS=xx" -e "DB_NAME=wiki" ghcr.io/requarks/wiki:3.0.0-alpha.334
```
> -v [YOUR-FILE.yml](https://github.com/Requarks/wiki/blob/main/config.sample.yml):/wiki/config.yml

> https://next.js.wiki/docs/admin/storage/

> -e "ADMIN_EMAIL=350840291@qq.com" -e "ADMIN_PASS=12345678"
> 默认登录邮箱: admin@example.com
> 默认密码: 12345678




### git

```
git config --global user.name "kingreatwill"
git config --global user.email "350840291@qq.com"

ssh-keygen -t rsa -C "350840291@qq.com"
```