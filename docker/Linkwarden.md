## 书签管理器(bookmark-manager)

https://github.com/topics/bookmarks-manager
https://github.com/topics/bookmark-manager

### shiori
https://github.com/go-shiori/shiori

[chrome 插件](https://github.com/go-shiori/shiori-web-ext)

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