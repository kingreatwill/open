## dotnet 5
[Performance Improvements in .NET 5](https://devblogs.microsoft.com/dotnet/performance-improvements-in-net-5/)
[Announcing .NET 5.0 RC 1](https://devblogs.microsoft.com/dotnet/announcing-net-5-0-rc-1/)
### 发布：
https://docs.microsoft.com/en-us/dotnet/core/tools/dotnet-publish

#### 发布修剪：
-p:PublishTrimmed=True and -p:TrimMode=Link

#### 发布单文件：
https://www.cnblogs.com/viter/archive/2020/09/04/13608947.html
平台	|命令	
---|---
Linux	|dotnet publish -r linux-x64 /p:PublishSingleFile=true	
Windows	|dotnet publish -r win-x64 --self-contained=false /p:PublishSingleFile=true	

可选参数
属性	|描述
---|---
IncludeNativeLibrariesInSingleFile|	在发布时，将依赖的本机二进制文件打包到单文件应用程序中。
IncludeSymbolsInSingleFile|	将 .pdb 文件打包到单个文件中。提供该选项是为了和 .NET 3 单文件模式兼容。建议替代的方法是生成带有嵌入式的 PDB （embedded）的程序集
IncludeAllContentInSingleFile|	将所有发布的文件（符号文件除外）打包到单文件中。该选项提供是为了向后兼容 .NETCore 3.x 版本

其它参数
```
<PropertyGroup>
  <SelfContained>true</SelfContained>
  <!--启用使用assemby修剪-仅支持自包含应用程序-->
  <PublishTrimmed> true </PublishTrimmed>
  <!--启用AOT编译 目前暂不支持预编译-->
  <!--<PublishReadyToRun>true</PublishReadyToRun>-->
</PropertyGroup>
<ItemGroup>
  <Content Update="*-exclute.dll">
    <CopyToPublishDirectory>PreserveNewest</CopyToPublishDirectory>
    <ExcludeFromSingleFile>true</ExcludeFromSingleFile>
  </Content>
</ItemGroup>
```
还可以通过设置 ExcludeFromSingleFile 元素，该设置将指定某些文件不嵌入单个文件之中。