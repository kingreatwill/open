修改全局包管理目录

通过 NuGet 安装包时，NuGet 先将包下载至一个统一的目录，默认路径是：C:\Users\用户名\.nuget\packages

下载的包多了以后，会导致 C 盘空间被大量占用。我们可以通过修改配置将其指定到自定义的目录下。

搜索 NuGet.Config 文件，默认位置是：C:\Users\用户名\AppData\Roaming\NuGet，在根节点下添加如下配置：
```
<config>
  <add key="globalPackagesFolder" value="D:\packages" />
</config>
```
如果 NuGet.Config 不存在，也可以在 C:\Program Files (x86)\NuGet\Config 目录下新建一个 NuGet.Config，将该文件夹中的 Microsoft.VisualStudio.Offline.config 文件的内容复制到新建的 NuGet.Config 中，再在其中添加上述的节点。