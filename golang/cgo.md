# cgo

## 交叉编译
### golang 交叉编译/Cross-compilation
Go 使用 Zig 来编译 C/C++ 代码。https://dev.to/kristoff/zig-makes-go-cross-compilation-just-work-29ho
```
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC="zig cc -target x86_64-linux" CXX="zig c++ -target x86_64-linux" go build --tags extended
```

## cgo 和共享对象
构建命令 go build 有一个参数可以[把你的 Go 包编进 C shared library](https://golang.org/pkg/cmd/go/internal/help/#pkg-variables)
```
-buildmode=c-shared
   Build the listed main package, plus all packages it imports,
   into a C shared library. The only callable symbols will
   be those functions exported using a cgo //export comment.
   Requires exactly one main package to be listed.
```

用这个模式你可以编译一个 shared library，文件以 “.so” 结尾，你可以在其他语言中如 C，Java，Ruby 或 Python 中直接使用。然而，这个模式只有[Cgo](https://golang.org/cmd/cgo/) 支持，你可以在你的 Go 包中写、调用 C 代码。基于此，你可以写一个自己的库，让别的小组用他们自己的语言来调你的库。

### 实现
Go 和共享对象间的网关的实现看起来很简单。首先你需要在你想导出的每一个函数前添加注释 //export MyFunction。然后你需要在强制性 import "C" 之前前置声明你的 C 结构体。下面是我们代码的简化版：
```
import (
   /*
   typedef struct{
   int from_bedroom;
   int to_bedroom;
   int from_price;
   int to_price;
   int from_size;
   int to_size;
   int types[5];
   } lead;

   typedef struct{
   int bedroom;
   int price;
   int size;
   int type_id;
   } property;
   */
   "C"
)
//export NewProperty
func NewProperty(b int, p int, s int, t int) C.property {
   // business logic

   return C.property{
      bedroom:   C.int(b),
      price:     C.int(p),
      size:      C.int(s),
      type_id:   C.int(t),
   }
}
//export NewLead
func NewLead(fb int, tb int, fp int, tp int, fs int, ts int, t []int) C.lead {
   // business logic
   return C.lead{
      from_bedroom: C.int(fb),
      to_bedroom:   C.int(tb),
      from_price:   C.int(fp),
      to_price:     C.int(tp),
      from_size:    C.int(fs),
      to_size:      C.int(ts),
      types:        types,
   }
}
//export CalculateDistance
func CalculateDistance(l C.lead, p C.property) {
   // business logic here
}
```
因为你不能导出 Go 的结构体，所以你需要把 C 的结构体作为输入/输出参数进行处理。当你写完代码后，可以使用命令 go build -o main.so -buildmode=c-shared main.go 来编译。为了能编译成功，你的 Go 代码中需要有 main 包和 main 函数。然后，你就可以写你的 Python 脚本了：
```
#!/usr/bin/env python
from ctypes import *

# loading shared object
matching = cdll.LoadLibrary("main.so")

# Go type
class GoSlice(Structure):
    _fields_ = [("data", POINTER(c_void_p)), ("len", c_longlong), ("cap", c_longlong)]

class Lead(Structure):
    _fields_ = [('from_bedroom', c_int),
                ('to_bedroom', c_int),
                ('from_price', c_int),
                ('to_price', c_int),
                ('from_size', c_int),
                ('to_size', c_int),
                ('types', GoSlice)]

class Property(Structure):
    _fields_ = [('bedroom', c_int),
                ('price', c_int),
                ('size', c_int),
                ('type_id', c_int)]

#parameters definition
matching.NewLead.argtypes = [c_int, c_int, c_int, c_int, c_int, c_int, GoSlice]
matching.NewLead.restype = Lead

matching.NewProperty.argtypes = [c_int, c_int, c_int, c_int]
matching.NewProperty.restype = Property

matching.CalculateDistance.argtypes = [Lead, Property]

lead = lib.NewLead(
    # from bedroom, to bedroom
    1, 2,
    # from price, to price
    80000, 100000,
    # from size, to size
    750, 1000,
    # type
    GoSlice((c_void_p * 5)(1, 2, 3, 4, 5), 5, 5)
)
property = lib.NewProperty(2, 90000, 900, 1)

matching.CalculateDistance(lead, property)
```
你的共享对象中的所有导出的方法都应该在你的 Python 文件中有描述：类型、参数顺序及返回参数。

之后，你可以运行你的 Python 脚本 python3 main.py。
