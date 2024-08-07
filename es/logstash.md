

## grok使用
https://www.cnblogs.com/even160941/p/16326986.html

https://www.elastic.co/guide/en/logstash/current/plugins-filters-grok.html

https://help.aliyun.com/zh/sls/user-guide/grok-patterns
## 收集容器标准输出

如果logstash是通过docker启动的,那么挂载目录/var/lib/docker/containers可能没有权限,需要通过`-u root`启动容器

```
docker run --name my-container \
    --log-driver json-file \
    --log-opt max-size=10m \
    --log-opt max-file=3 \
    --label service-name=xxx --log-opt labels=service-name
    your-image
```
日志里面会多打出
```
"attrs":{"service-name":"xxx"}
```


收集容器标准输出
```
input {
  file {
    path => "/var/lib/docker/containers/*/*-json.log"
    start_position => "beginning" # 如果你想从现有日志开始读取，可以设置为 "beginning"
    sincedb_path => "/dev/null" # 不保存文件位置信息，每次重启都重新读取
  }
}

filter {
  json {
    source => "message"
  }
  mutate {
    add_field => { "[@metadata][container_id]" => "%{[docker][container][id]}" } # 提取容器 ID
  }
  grok {
    match => { "host" => "%{DOCKER_CONTAINER_NAME:[@metadata][container_name]}" } # 提取容器名称
  }
  if ![@metadata][container_name] =~ /^(exclude-me|another-exclude)$/ { # 排除名为 "exclude-me" 和 "another-exclude" 的容器
    date {
      match => [ "time", "ISO8601" ]
    }
  }
}

output {
  elasticsearch {
    hosts => ["localhost:9200"]
    index => "docker-%{+YYYY.MM.dd}"
  }
}
```



根据path提取container_id,  match => { "path" => "(?<container_id>[^/]+)/[^/]+-json.log" }
```
input {
  file {
    path => "/var/lib/docker/containers/*/*-json.log"
    type => "docker_log"
    start_position => "beginning"
    sincedb_path => "/dev/null"
  }
}
filter {
    if [type] == "docker_log" {
        json {
            source => "message"
        }
        grok {
		    match => { "path" => "(?<container_id>[^/]+)/[^/]+-json.log" }
	    }
        mutate {
            remove_field => ["message", "@version", "path"]
        }
        if [type] == "docker_log" {
            json {
            source => "message"
            }
            if [attrs.service-name] =~ /^(mf-logstash|cdn-other)$/ {
            date {
                match => [ "time", "ISO8601" ]
            }
            } else {
            drop {}
            }
        }
    }
}

output {
  if [type] == "docker_log" {
    elasticsearch {
      hosts => ["xxxx"]
      user => "xx"
      password => "xxx"
      index => "docker-logs-%{+YYYY.MM.dd}"
    }
  }
}
```