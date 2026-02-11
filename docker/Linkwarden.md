## 书签管理器(bookmark-manager)

https://github.com/topics/bookmarks-manager
https://github.com/topics/bookmark-manager

### Pintree
https://github.com/Pintree-io/pintree
### wallabag
https://github.com/wallabag/wallabag 

### memos笔记
https://github.com/usememos/memos
```
docker run -d --name memos --link postgresql --restart always -p 10009:5230 -v /data/dockerv/memos/data:/var/opt/memos neosmemo/memos:stable --driver postgres --dsn 'postgresql://postgres:PASSWORD@postgresql:5432/memos?sslmode=disable'
```

### flomo
非开源

### shiori
https://github.com/go-shiori/shiori

[chrome 插件](https://github.com/go-shiori/shiori-web-ext)
> chrome 点击"扩展程序选项"设置用户名和站点


[配置参考](https://github.com/go-shiori/shiori/blob/master/docs/Configuration.md)
`shiori.exe server --portable`

测试
```
docker run -d --rm --name shiori -p 8080:8080 -v $(pwd):/shiori ghcr.io/go-shiori/shiori:dev
```
`openssl rand -hex 32`可以生产随机值(SHIORI_HTTP_SECRET_KEY)
```
SHIORI_HTTP_SECRET_KEY=xx
SHIORI_DIR=/shiori
SHIORI_DATABASE_URL="postgres://pqgotest:password@hostname/database?sslmode=verify-full"
SHIORI_DATABASE_URL="mysql://username:password@(hostname:port)/database?charset=utf8mb4"
```

创建数据库:`shiori`
```
docker run -d -p 10001:8080 -v /data/dockerv/shiori/data:/shiori -e SHIORI_HTTP_SECRET_KEY=xx  -e SHIORI_DIR=/shiori  -e SHIORI_DATABASE_URL="postgres://postgres:pwd@postgresql/shiori?sslmode=disable"  --name shiori --link postgresql --restart always ghcr.io/go-shiori/shiori:dev
```

> 如果密码中包含特殊字符 `fmt.Println(url.QueryEscape("^"))  print: %5E`


默认用户
username: shiori
password: gopher

源码编译
```
FROM docker.io/golang:1.26-alpine AS builder
ARG VERSION="1.8.0"
ARG COMMIT="585ea341aa59219b0477f991eadb545d24e3a121"
ARG DATE="2026-02-11"

WORKDIR /src/shiori

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0
RUN go build -tags "osusergo,netgo,fts5" \
    -ldflags "-s -w -X main.version=${VERSION} -X main.commit=${COMMIT} -X main.date=${DATE}" \
    -o /out/shiori .

FROM docker.io/library/alpine:3.22

RUN apk add --no-cache ca-certificates tzdata

ENV PORT=8080
ENV SHIORI_DIR=/srv/shiori
ENV SHIORI_HTTP_SERVE_WEB_UI_V2=true
WORKDIR ${SHIORI_DIR}

COPY --from=builder /out/shiori /usr/bin/shiori

EXPOSE ${PORT}

ENTRYPOINT ["/usr/bin/shiori"]
CMD ["server"]

# docker build -t shiori .
# docker tag shiori kingreatwill/shiori:v1.8.0
# docker login -u kingreatwill
# docker push kingreatwill/shiori:v1.8.0
```



### linkding
https://github.com/sissbruecker/linkding
### Linkwarden
https://github.com/linkwarden/linkwarden

```

version: "3.5"
services:
  postgres:
    image: postgres:16-alpine
    environment:
      - POSTGRES_PASSWORD=pgsql2023
    ports:
      - 5432:5432
    restart: always
    volumes:
      - ./pgdata:/var/lib/postgresql/data
  linkwarden:
    # env_file: .env
    environment:
      - POSTGRES_PASSWORD=pgsql2023
      - DATABASE_URL=postgresql://postgres:${POSTGRES_PASSWORD}@postgres:5432/postgres
      - NEXTAUTH_SECRET=pgsql2023
      - NEXTAUTH_URL=http://localhost:3000 # 可以修改成自定义域名
    restart: always
    image: ghcr.io/linkwarden/linkwarden:latest
    ports:
      - 3000:3000
    volumes:
      - ./data:/data/data
    depends_on:
      - postgres

```