
## 安装

http://openresty.org/cn/download.html


### windows
1. 下载后双击nginx.exe运行 
2. 验证：
`tasklist /fi "imagename eq nginx.exe"`
http://127.0.0.1:80/


3. 自定义lua
在\conf\nginx.conf文件修改
```
location /hello {
	#修改为直接返回text模式，而不是返回文件。默认配置在极速模式下得浏览器会形成下载文件
	default_type text/html;
	#关闭缓存模式，每次都读取lua文件，不使用已经缓存的lua文件（修改nginx.conf文件还是要重启的）
	lua_code_cache off;
	#在返回节点加载lua文件（相对路径即可）
	content_by_lua_file lua/hello.lua;
}
```
添加一下脚本\lua\hello.lua
```
ngx.say('hello world!!!')
```

nginx -s reload
nginx -s stop

linux指定文件
```
nginx -p `pwd`/ -c conf/nginx.conf
```


## 包管理器opm

opm - OpenResty Package Manager
https://github.com/openresty/opm
https://opm.openresty.org/

> OpenResty@1.11.2.1 支持了 opm, 但并没有内置其中, 需要自行下载安装
```
curl https://raw.githubusercontent.com/openresty/opm/master/bin/opm > /usr/local/openresty/bin/opm
chmod +x /usr/local/openresty/bin/opm

export PATH=/usr/local/openresty/bin:$PATH
```

opm 安装包默认是全局模式, 文件都会被放在 /usr/local/openresty/site/ 路径下

因此 opm 的版本控制和包管理是扁平而非嵌套的

A 依赖 C
B 也依赖 C

opm 会下载 A, B, C 各一遍, 不会下载 C 两遍, 逻辑比 npm 简单很多

那问题来了, 如果我就是要不同的版本进行开发呢?
那就是**本地模式**
opm 加上参数 --cwd 就选择了本地模式
opm 包会下载到当前路径中, 实现本地包场景需求
和 npm 正好相反, npm install 默认安装本地包, 加上 -g 才是安装全局包, 而 opm 则默认安装全局包

参考: [OpenResty 大跃进! opm 包管理尝鲜](https://zhuanlan.zhihu.com/p/22991825)