

https://github.com/nocodb/nocodb


cloud: https://app.nocodb.com/#/
## 部署

环境变量: https://docs.nocodb.com/getting-started/self-hosted/environment-variables

### docker
```
# 使用 SQLite 作为数据库
docker run -d --name noco 
           -v "$(pwd)"/nocodb:/usr/app/data/ 
           -p 8080:8080 
           nocodb/nocodb:latest




# 使用 PostgreSQL 作为数据库
docker run -d --name noco 
           -v "$(pwd)"/nocodb:/usr/app/data/ 
           -p 8080:8080 
            # replace with your pg connection string
           -e NC_DB="pg://host.docker.internal:5432?u=root&p=password&d=d1" 
           # replace with a random secret
           -e NC_AUTH_JWT_SECRET="569a1821-0a93-45e8-87ab-eb857f20a010"  
           nocodb/nocodb:latest
```
### Docker-Compose 部署
```
#默认启用PG作为数据
#Clone the NocoDB repository from GitHub.
git clone https://github.com/nocodb/nocodb

#Navigate to the docker-compose directory
cd nocodb/docker-compose/2_pg

#Start the services using Docker Compose:
docker-compose up -d
```
docker-compose.yml
```
version: '2.1'
services: 
  nocodb: 
    depends_on: 
      root_db: 
        condition: service_healthy
    environment: 
      NC_DB: "pg://root_db:5432?u=postgres&p=password&d=root_db"
    image: "nocodb/nocodb:latest"
    ports: 
      - "8080:8080"
    restart: always
    volumes: 
      - "nc_data:/usr/app/data"
  root_db: 
    environment: 
      POSTGRES_DB: root_db
      POSTGRES_PASSWORD: password
      POSTGRES_USER: postgres
    healthcheck: 
      interval: 10s
      retries: 10
      test: "pg_isready -U \"$$POSTGRES_USER\" -d \"$$POSTGRES_DB\""
      timeout: 2s
    image: postgres
    restart: always
    volumes: 
      - "db_data:/var/lib/postgresql/data"
volumes: 
  db_data: {}
  nc_data: {}
```

### 命令
`bash <(curl -sSL http://install.nocodb.com/noco.sh) <(mktemp)`

自动安装docker compose等相关依赖

### 其它
`curl http://get.nocodb.com/linux-x64 -o nocodb -L && chmod +x nocodb && ./nocodb`

https://github.com/nocodb/nocodb
https://github.com/Airtable/airtable.js
https://github.com/teableio/teable
