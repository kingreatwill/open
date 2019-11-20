前提：
- 安装WLS
- 安装git
- [安装MinGW](https://sourceforge.net/projects/mingw-w64/)
- [安装GnuWin32](https://sourceforge.net/projects/gnuwin32/files/) zip

[sdkman](https://sdkman.io/install)

安装：
```bash
curl -s https://get.sdkman.io | bash
```

```bash
source "$HOME/.sdkman/bin/sdkman-init.sh"
```

```bash
sdk version
sdk help
```


使用：
```
sdk install gradle 6.0.1
sdk install gradle 5.6.4

gradle -v

 sdk uninstall gradle
 or
 sdk rm gradle

 sdk use gradle 5.6.4
 sdk default gradle 5.6.4

 sdk current gradle
```

