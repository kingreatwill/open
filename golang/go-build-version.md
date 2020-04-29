```
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
```
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