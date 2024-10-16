[安装](https://hub.docker.com/_/postgres)
```
docker pull postgres:12.0
docker pull postgres:13.0-alpine
docker pull kingreatwill/postgres:13-alpine-jieba
docker volume create --name=pgdata

docker run -d -p 5432:5432 -v pgdata:/var/lib/postgresql/data -e POSTGRES_PASSWORD=123456@lcb --name postgresql --restart always postgres:12.0

docker run -d -p 5432:5432 -v /d/dockerv/postgresql13/data:/var/lib/postgresql/data -e POSTGRES_PASSWORD=123456@lcb --name postgresql --restart always postgres:13


PGDATA=/var/lib/postgresql/data

默认数据库:postgres

POSTGRES_USER: root
POSTGRES_DB: database
POSTGRES_PASSWORD: 123456

默认用户名:postgres 

docker exec -it postgresql /bin/bash

切换用户
su postgres
```
## 备份和恢复
使用如下命令对名为dbname的数据库进行备份：
pg_dump  –h 127.0.0.1  -p  5432  -U  postgres -c  -C –f  dbname.sql  dbname

使用如下命令可对全部pg数据库进行备份：
pg_dumpall –h 127.0.0.1 –p 5432 -U postgres –c  -C –f db_bak.sql

恢复方式很简单。执行恢复命令即可：
psql –h 127.0.0.1 -p 5432 -U postgres –f db_bak.sql

## 升级
https://github.com/tianon/docker-postgres-upgrade

```
docker run --rm \
	-v /data/dockerv/postgresql/16/data:/var/lib/postgresql/16/data \
	-v /data/dockerv/postgresql/17/data:/var/lib/postgresql/17/data \
	tianon/postgres-upgrade:16-to-17
```

## 分词器
### pg_jieba

pg_jieba.dockerfile
```
FROM postgres:13-alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.cloud.tencent.com/' /etc/apk/repositories && sed -i 's/http:/https:/' /etc/apk/repositories

RUN set -ex \
  && apk add --no-cache --virtual .fetch-deps ca-certificates cmake git openssl tar \  
  && apk add --no-cache --virtual .build-deps gcc g++ libc-dev make postgresql-dev \
  && apk add --no-cache --virtual .rundeps libstdc++ \
  && git clone --recursive https://github.com/jaiminpan/pg_jieba \
  && cd /pg_jieba \
  && mkdir build \
  && cd build \
  && cmake .. \
  && make \
  && make install \
  && echo -e "  \necho \"shared_preload_libraries = 'pg_jieba'\" >> /var/lib/postgresql/data/postgresql.conf" > /docker-entrypoint-initdb.d/init-conf.sh \
  && chmod +x /docker-entrypoint-initdb.d/init-conf.sh \
  && echo -e "CREATE EXTENSION pg_jieba;" > /docker-entrypoint-initdb.d/init-jieba.sql \
  && apk del .build-deps .fetch-deps \
  && rm -rf /pg_jieba
```
cmd
```

docker build -f pg_jieba.dockerfile -t pg_jieba:13-alpine .

docker tag pg_jieba:13-alpine  kingreatwill/postgres:13-alpine-jieba

docker push  kingreatwill/postgres:13-alpine-jieba
```

https://github.com/ssfdust/psql_jieba_swsc
```
FROM postgres:alpine

RUN apk add --no-cache --virtual .build \
        postgresql-dev \
        gcc \
        make \
        llvm \
        libc-dev \
        g++ \
        clang \
        git \
        cmake \
        curl \
        openssl-dev && \
    git clone https://github.com/jaiminpan/pg_jieba && \
    cd /pg_jieba && \
    git submodule update --init --recursive && \
    mkdir -p build && \
    cd build && \
    curl -L https://raw.githubusercontent.com/ssfdust/psql_jieba_swsc/master/FindPostgreSQL.cmake > $(find /usr -name "FindPostgreSQL.cmake") && \
    cmake .. && \
    make && \
    make install && \
    cd / && \
    git clone https://github.com/jaiminpan/pg_scws && \
    cd /pg_scws && \
    USE_PGXS=1 make && \
    USE_PGXS=1 make install && \
    apk del .build && \
    rm -rf /pg_jieba /pg_scws

RUN echo "echo \"shared_preload_libraries = 'pg_jieba'\" >> /var/lib/postgresql/data/postgresql.conf" \
    > /docker-entrypoint-initdb.d/init-dict.sh  && \
    echo "CREATE EXTENSION pg_jieba;create extension pg_scws;" > /docker-entrypoint-initdb.d/init-jieba.sql
```


https://github.com/chen-xin/docker_pg_jieba

```
FROM postgres:13.0
ARG CN_MIRROR=0

# Uncomment the following command if you have bad internet connection
# and first download the files into data directory
# COPY data/pg_jieba.zip /pg_jieba.zip

RUN if [ $CN_MIRROR = 1 ] ; then DEBIAN_VERSION=$(dpkg --status tzdata|grep Provides|cut -f2 -d'-') \
&& echo "using mirrors for $DEBIAN_VERSION" \
&& echo "deb http://ftp.cn.debian.org/debian/ $DEBIAN_VERSION main non-free contrib \n\
deb http://ftp.cn.debian.org/debian/ $DEBIAN_VERSION-updates main non-free contrib \n\
deb http://ftp.cn.debian.org/debian/ $DEBIAN_VERSION-backports main non-free contrib \n\
deb http://ftp.cn.debian.org/debian-security/ $DEBIAN_VERSION/updates main non-free contrib \n\
deb-src http://ftp.cn.debian.org/debian/ $DEBIAN_VERSION main non-free contrib \n\
deb-src http://ftp.cn.debian.org/debian/ $DEBIAN_VERSION-updates main non-free contrib \n\
deb-src http://ftp.cn.debian.org/debian/ $DEBIAN_VERSION-backports main non-free contrib \n\
deb-src http://ftp.cn.debian.org/debian-security/ $DEBIAN_VERSION/updates main non-free contrib" > /etc/apt/sources.list; else echo "No mirror"; fi

RUN apt-get update \
  && apt-get install -y --no-install-recommends \
      postgresql-server-dev-$PG_MAJOR \
      gcc \
      make \
      libc-dev \
      g++ \
      git \
      cmake \
      curl \
      ca-certificates \
      openssl \
	&& rm -rf /var/lib/apt/lists/*

RUN git clone https://github.com/jaiminpan/pg_jieba \
  && cd /pg_jieba \
  && git submodule update --init --recursive 

RUN cd /pg_jieba \
  && mkdir -p build \
  && cd build \
  && curl -L https://raw.githubusercontent.com/Kitware/CMake/ce629c5ddeb7d4a87ac287c293fb164099812ca2/Modules/FindPostgreSQL.cmake > $(find /usr -name "FindPostgreSQL.cmake") \
  && cmake .. \
  && make \
  && make install \
  && echo "  \n\
  # echo \"timezone = 'Asia/Shanghai'\" >> /var/lib/postgresql/data/postgresql.conf \n\
  echo \"shared_preload_libraries = 'pg_jieba.so'\" >> /var/lib/postgresql/data/postgresql.conf" \
  > /docker-entrypoint-initdb.d/init-dict.sh \
# The following command is not required if load database from backup
  && echo "CREATE EXTENSION pg_jieba;" > /docker-entrypoint-initdb.d/init-jieba.sql \
# RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
#   && echo "Asia/Shanghai" >  /etc/timezone
  && apt-get purge -y gcc make libc-dev postgresql-server-dev-$PG_MAJOR g++ git cmake curl\
  && apt-get autoremove -y \
  && rm -rf \
    /pg_jieba
```
docker build -t postgres:13.0-pg_jieba .


alpine
```
# vim:set ft=dockerfile:
FROM postgres:alpine
ARG CN_MIRROR=0

RUN if [ $CN_MIRROR = 1 ] ; then OS_VER=$(grep main /etc/apk/repositories | sed 's#/#\n#g' | grep "v[0-9]\.[0-9]") \
  && echo "using mirrors for $OS_VER" \
  && echo https://mirrors.ustc.edu.cn/alpine/$OS_VER/main/ > /etc/apk/repositories; fi

RUN set -ex \
	&& apk add --no-cache --virtual .fetch-deps \
		ca-certificates \
    cmake \
    git \
		openssl \
		tar \
  && git clone https://github.com/jaiminpan/pg_jieba \
	&& apk add --no-cache --virtual .build-deps \
		gcc \
		g++ \
		libc-dev \
		make \
    postgresql-dev \
	&& apk add --no-cache --virtual .rundeps \
		libstdc++ \
  && cd / \
  && cd pg_jieba \
  && git submodule update --init --recursive \
  && mkdir build \
  && cd build \
  && cmake .. \
  && make \
  && make install \
  && echo -e "  \n\
  # echo \"timezone = 'Asia/Shanghai'\" >> /var/lib/postgresql/data/postgresql.conf \n\
  echo \"shared_preload_libraries = 'pg_jieba.so'\" >> /var/lib/postgresql/data/postgresql.conf" \
  > /docker-entrypoint-initdb.d/init-dict.sh \
# The following command is not required if load database from backup
  && echo -e "CREATE EXTENSION pg_jieba;" > /docker-entrypoint-initdb.d/init-jieba.sql \
# RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
#   && echo "Asia/Shanghai" >  /etc/timezone
    && apk del .build-deps .fetch-deps \
	&& rm -rf \
		/usr/src/postgresql \
		/pg_jieba \
	&& find /usr/local -name '*.a' -delete
```
alpine版本
```
FROM postgres:alpine
ARG CN_MIRROR=0

RUN if [ $CN_MIRROR = 1 ] ; then OS_VER=$(grep main /etc/apk/repositories | sed 's#/#\n#g' | grep "v[0-9]\.[0-9]") \
  && echo "using mirrors for $OS_VER" \
  && echo https://mirrors.ustc.edu.cn/alpine/$OS_VER/main/ > /etc/apk/repositories; fi

RUN set -ex \
	&& apk add --no-cache --virtual .fetch-deps \
		ca-certificates \
    cmake \
    git \
		openssl \
		tar \
  && git clone https://github.com/jaiminpan/pg_jieba \
	&& apk add --no-cache --virtual .build-deps \
		gcc \
		g++ \
		libc-dev \
		make \
    postgresql-dev \
	&& apk add --no-cache --virtual .rundeps \
		libstdc++ \
  && cd / \
  && cd pg_jieba \
  && git submodule update --init --recursive \
  && mkdir build \
  && cd build \
  && cmake .. \
  && make \
  && make install \
  && echo -e "  \n\
  # echo \"timezone = 'Asia/Shanghai'\" >> /var/lib/postgresql/data/postgresql.conf \n\
  echo \"shared_preload_libraries = 'pg_jieba.so'\" >> /var/lib/postgresql/data/postgresql.conf" \
  > /docker-entrypoint-initdb.d/init-dict.sh \
# The following command is not required if load database from backup
  && echo -e "CREATE EXTENSION pg_jieba;" > /docker-entrypoint-initdb.d/init-jieba.sql \
# RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
#   && echo "Asia/Shanghai" >  /etc/timezone
    && apk del .build-deps .fetch-deps \
	&& rm -rf \
		/usr/src/postgresql \
		/pg_jieba \
	&& find /usr/local -name '*.a' -delete
```
alpine版本
```
FROM postgres:alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.cloud.tencent.com/' /etc/apk/repositories && sed -i 's/http:/https:/' /etc/apk/repositories

RUN set -ex \
  && apk add --no-cache --virtual .fetch-deps ca-certificates cmake git openssl tar \
  && git clone https://github.com/jaiminpan/pg_jieba \
  && apk add --no-cache --virtual .build-deps gcc g++ libc-dev make postgresql-dev \
  && apk add --no-cache --virtual .rundeps libstdc++ \
  && cd /pg_jieba \
  && git submodule update --init --recursive \
  && mkdir build \
  && cd build \
  && cmake .. \
  && make \
  && make install \
  && echo -e "  \n\
  # echo \"timezone = 'Asia/Shanghai'\" >> /var/lib/postgresql/data/postgresql.conf \n\
  echo \"shared_preload_libraries = 'pg_jieba.so'\" >> /var/lib/postgresql/data/postgresql.conf" \
  > /docker-entrypoint-initdb.d/init-dict.sh \
# The following command is not required if load database from backup
  && echo -e "CREATE EXTENSION pg_jieba;" > /docker-entrypoint-initdb.d/init-jieba.sql \
# RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
#   && echo "Asia/Shanghai" >  /etc/timezone
  && apk del .build-deps .fetch-deps \
  && rm -rf /usr/src/postgresql /pg_jieba \
  && find /usr/local -name '*.a' -delete
```

### zhparser
https://www.github.com/chen-xin/docker_zhparser
- zhparser
```
# If you don‘t want to build it youself, you can try `docker pull killercai/postgres`.
FROM healthcheck/postgres:latest

# China debian mirror
RUN sed -i s@/deb.debian.org/@/mirrors.aliyun.com/@g /etc/apt/sources.list
RUN apt-get clean && apt-get update
RUN apt-get install -y wget git build-essential libpq-dev python-dev postgresql-server-dev-all
# SCWS (Simple Chinese Word Segmentation library)
RUN cd /tmp && wget -q -O - http://www.xunsearch.com/scws/down/scws-1.2.1.tar.bz2 | tar xjf - && cd scws-1.2.1 && ./configure && make install
# zhpaser (postgres plugin)
RUN cd /tmp && git clone https://github.com/amutu/zhparser.git && cd zhparser && make && make install
```
sql
```
-- 安装扩展
CREATE EXTENSION zhparser;
-- 中文分词配置
CREATE TEXT SEARCH CONFIGURATION chinese_parser (PARSER = zhparser);
ALTER TEXT SEARCH CONFIGURATION chinese_parser ADD MAPPING FOR n,v,a,i,e,l,j WITH simple;
使用：chinese_parser 
```


```
# Azurewind's PostgreSQL image with Chinese full text search
# build: docker build --force-rm -t chenxinaz/zhparser .
# run: docker run --name PostgreSQLcnFt -p 5432:5432 chenxinaz/zhparser
# run interactive: winpty docker run -it --name PostgreSQLcnFt -p 5432:5432 chenxinaz/zhparser --entrypoint bash chenxinaz/zhparser

FROM postgres

ARG CN_MIRROR=0

# Uncomment the following command if you have bad internet connection
# and first download the files into data directory
# COPY data/pg_jieba.zip /pg_jieba.zip

RUN if [ $CN_MIRROR = 1 ] ; then DEBIAN_VERSION=$(dpkg --status tzdata|grep Provides|cut -f2 -d'-') \
&& echo "using mirrors for $DEBIAN_VERSION" \
&& echo "deb http://ftp.cn.debian.org/debian/ $DEBIAN_VERSION main non-free contrib \n\
deb http://ftp.cn.debian.org/debian/ $DEBIAN_VERSION-updates main non-free contrib \n\
deb http://ftp.cn.debian.org/debian/ $DEBIAN_VERSION-backports main non-free contrib \n\
deb http://ftp.cn.debian.org/debian-security/ $DEBIAN_VERSION/updates main non-free contrib \n\
deb-src http://ftp.cn.debian.org/debian/ $DEBIAN_VERSION main non-free contrib \n\
deb-src http://ftp.cn.debian.org/debian/ $DEBIAN_VERSION-updates main non-free contrib \n\
deb-src http://ftp.cn.debian.org/debian/ $DEBIAN_VERSION-backports main non-free contrib \n\
deb-src http://ftp.cn.debian.org/debian-security/ $DEBIAN_VERSION/updates main non-free contrib" > /etc/apt/sources.list; else echo "No mirror"; fi

RUN apt-get update \
  && apt-get install -y --no-install-recommends \
      gcc \
      make \
      libc-dev \
      postgresql-server-dev-$PG_MAJOR \
      wget \
      unzip \
      ca-certificates \
      openssl \
	&& rm -rf /var/lib/apt/lists/* \
  && wget -q -O - "http://www.xunsearch.com/scws/down/scws-1.2.3.tar.bz2" | tar xjf - \
  && wget -O zhparser.zip "https://github.com/amutu/zhparser/archive/master.zip" \
  && unzip zhparser.zip \
  && cd scws-1.2.3 \
  && ./configure \
  && make install \
  && cd /zhparser-master \
  && SCWS_HOME=/usr/local make && make install \
  # pg_trgm is recommend but not required.
  && echo "CREATE EXTENSION pg_trgm; \n\
CREATE EXTENSION zhparser; \n\
CREATE TEXT SEARCH CONFIGURATION chinese_zh (PARSER = zhparser); \n\
ALTER TEXT SEARCH CONFIGURATION chinese_zh ADD MAPPING FOR n,v,a,i,e,l,t WITH simple;" \
> /docker-entrypoint-initdb.d/init-zhparser.sql \
  && apt-get purge -y gcc make libc-dev postgresql-server-dev-$PG_MAJOR \
  && apt-get autoremove -y \
  && rm -rf \
    /zhparser-master \
    /zhparser.zip \
    /scws-1.2.3
```


## 本地安装


本地访问#
先登录 shell （以 postgres 用户为例）：
```
sudo -i -u postgres

or

sudo su - postgres
```
然后，输入：
```
# 以当前登录的 linux 用户名为数据库名 (这点跟 mysql 不一样，必须得先指定登录的数据库)
psql 

# 手动指定数据库名
psql -d postgres
```



开启远程访问
vi /var/lib/pgsql/12/data/pg_hba.conf

若不知道 PostgreSQL 的配置文件路径，可在 PostgreSQL CLI 里输入：SHOW config_file 

然后把文件拉到底，最后一段里，将“ident”替换为“md5”，最终是这样的：
```
# TYPE  DATABASE        USER            ADDRESS                 METHOD

# "local" is for Unix domain socket connections only
local   all             all                                     trust
# IPv4 local connections:
host    all             all             127.0.0.1/32            md5
# IPv6 local connections:
host    all             all             ::1/128                 md5
# Allow replication connections from localhost, by a user with the
# replication privilege.
local   replication     postgres                                peer
host    replication     postgres        127.0.0.1/32            md5
host    replication     postgres        ::1/128                 md5
```

外网访问, 可以加上所以IP
host all all 127.0.0.1/32 md5