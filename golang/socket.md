

## Golang网络编程：像C语言一样操作Socket
在C语言中，进行网络编程通常使socket()、bind()、listen()、accept()等系统调用，而在Golang中，我们既可以使用net库提供的高级封装，也可以直接使用syscall库进行底层操作。本篇文章将简单的介绍Golang如何像C语言一样进行网络编程。


### 使用net库进行TCP编程（推荐方式）
Golang的net包对socket进行了封装，使得网络编程更简单、更易用。

TCP服务器例子net版
```

package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080") // 监听 8080 端口
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Server listening on port 8080...")

	for {
		conn, err := listener.Accept() // 接受客户端连接
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}

		go handleConnection(conn) // 处理连接
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer) // 读取数据
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}
	fmt.Println("Received:", string(buffer[:n]))
	conn.Write([]byte("Hello from server\n")) // 发送响应
}
```

TCP客户端net版
```
package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080") // 连接服务器
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	conn.Write([]byte("Hello, Server!\n")) // 发送数据
	buffer := make([]byte, 1024)
	n, _ := conn.Read(buffer) // 读取服务器响应
	fmt.Println("Server response:", string(buffer[:n]))
}
```
上述给出的就是使用golang标准库中net包编写的，简单的tcp server代码例子，这比传统的socket代码量少了很多，却能达到预期的效果。但是在实际的工作中，往往有些需求我们不能直接使用net包进行操作的，例如一些特殊情况下我们需要自定义的数据包包头，加密等等操作。


### 使用syscall实现TCP服务器
TCP服务器例子syscall版
```
package main

import (
    "fmt"
    "syscall"
)

func main() {
    // 创建 socket
    fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
    if err != nil {
        panic(err)
    }
    defer syscall.Close(fd)

    // 设置 SO_REUSEADDR，避免端口占用问题
    err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
    if err != nil {
        panic(err)
    }

    // 绑定地址和端口
    addr := syscall.SockaddrInet4{Port: 9527, Addr: [4]byte{0, 0, 0, 0}}
    if err := syscall.Bind(fd, &addr); err != nil {
        panic(err)
    }

    // 监听连接
    if err := syscall.Listen(fd, syscall.SOMAXCONN); err != nil {
        panic(err)
    }

    fmt.Println("Server listening on port 9527...")

    for {
        // 接受客户端连接
        clientFD, _, err := syscall.Accept(fd)
        if err != nil {
            fmt.Println("Accept error:", err)
            continue
        }
        go handleClient(clientFD)
    }
}

// 处理客户端连接
func handleClient(fd int) {
    defer syscall.Close(fd)

    buffer := make([]byte, 1024)
    n, err := syscall.Read(fd, buffer) // 读取数据
    if err != nil {
        fmt.Println("Read error:", err)
        return
    }

    fmt.Println("Received:", string(buffer[:n]))

    // 发送数据
    syscall.Write(fd, []byte("Hello from server\n"))
}
```
代码解析
- syscall.Socket(AF_INET, SOCK_STREAM, 0): 创建TCP套接字（IPv4）。
- syscall.Bind(fd, &addr): 绑定端口9527。
- syscall.Listen(fd, syscall.SOMAXCONN): 监听连接。
- syscall.Accept(fd): 接受连接并返回新的clientFD。
- syscall.Read(fd, buffer): 读取客户端数据。
- syscall.Write(fd, response): 发送数据。

TCP客户端syscall版
```

package main

import (
    "fmt"
    "syscall"
)

func main() {
    // 创建 socket
    fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)

    if err != nil {
        panic(err)
    }

    defer syscall.Close(fd)

    // 连接服务器
    serverAddr := syscall.SockaddrInet4{Port: 9527, Addr: [4]byte{127, 0, 0, 1}}

    if err := syscall.Connect(fd, &serverAddr); err != nil {
        panic(err)
    }

    // 发送数据
    message := "Hello, Server!"
    syscall.Write(fd, []byte(message))

    // 读取响应
    buffer := make([]byte, 1024)
    n, err := syscall.Read(fd, buffer)

    if err != nil {
        fmt.Println("Read error:", err)
        return
    }

    fmt.Println("Server response:", string(buffer[:n]))
}
```

代码解析
- syscall.Socket(AF_INET, SOCK_STREAM, 0): 创建 TCP 套接字。
- syscall.Connect(fd, &serverAddr): 连接到 127.0.0.1:9527 服务器。
- syscall.Write(fd, message): 发送数据。
- syscall.Read(fd, buffer): 读取服务器响应。