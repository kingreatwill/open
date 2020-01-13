# Golang写的程序注入一些版本信息
## 背景
Go程序运行时打印git提交信息编译信息

Golang编译信息注入程序

当在debug的过程中，我们需要明确当前运行的go程序是什么版本

不要浪费时间在确认版本的问题上

在go build编译的时候是可以注入外部参数的

让go程序在运行的时候就可以打印编译时候的参数情况

以gitlab-runner为例
```
gitlab-runner -v
Version:      11.10.1
Git revision: 1f513601
Git branch:   11-10-stable
GO version:   go1.8.7
Built:        2019-04-24T09:29:18+0000
OS/Arch:      linux/amd64
```
最终实现的go程序运行时终端打印的信息如下
```
App Name:       app-api
App Version:    v2.0.1
Build version:  84d4ffb verdor
Build time:     2019-08-06T09:58:48+0800
Git revision:   84d4ffb
Git branch:     master
Golang Version: go version go1.12.2 linux/amd64
2019-07-24 10:53:34.732 11516: http server started listening on [:20000]
```
## 具体实现
入口文件main.go
```go
package main

import (
    "fmt"
)

var (
    AppName      string // 应用名称
    AppVersion   string // 应用版本
    BuildVersion string // 编译版本
    BuildTime    string // 编译时间
    GitRevision  string // Git版本
    GitBranch    string // Git分支
    GoVersion    string // Golang信息
)


func main() {
    Version()
    // 你的业务代码入口
}

// Version 版本信息
func Version() {
    fmt.Printf("App Name:\t%s\n", AppName)
    fmt.Printf("App Version:\t%s\n", AppVersion)
    fmt.Printf("Build version:\t%s\n", BuildVersion)
    fmt.Printf("Build time:\t%s\n", BuildTime)
    fmt.Printf("Git revision:\t%s\n", GitRevision)
    fmt.Printf("Git branch:\t%s\n", GitBranch)
    fmt.Printf("Golang Version: %s\n", GoVersion)
}
```
build编译构建脚本build.sh
```sh
#!/bin/bash

set -e

PROJECT_NAME="app-api"
BINARY="app-api"

OUTPUT_DIR=output
GOOS=$(go env GOOS)

APP_NAME=${PROJECT_NAME}
APP_VERSION=$(git log -1 --oneline)
BUILD_VERSION=$(git log -1 --oneline)
BUILD_TIME=$(date "+%FT%T%z")
GIT_REVISION=$(git rev-parse --short HEAD)
GIT_BRANCH=$(git name-rev --name-only HEAD)
GO_VERSION=$(go version)

CGO_ENABLED=0 go build -a -installsuffix cgo -v -mod=vendor \
-ldflags "-s -X 'main.AppName=${APP_NAME}' \
            -X 'main.AppVersion=${APP_VERSION}' \
            -X 'main.BuildVersion=${BUILD_VERSION}' \
            -X 'main.BuildTime=${BUILD_TIME}' \
            -X 'main.GitRevision=${GIT_REVISION}' \
            -X 'main.GitBranch=${GIT_BRANCH}' \
            -X 'main.GoVersion=${GO_VERSION}'" \
-o ${OUTPUT_DIR}/${BINARY} cmd/${BINARY}.go
```

本质上是用 -ldflags 参数注入了的外部参数到go的变量当中go的更多build参数帮助可以通过 go help build获取
> windows用Git Bash终端