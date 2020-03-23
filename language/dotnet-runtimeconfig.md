# .NET Core 运行时配置设置.NET Core run-time configuration settings
[原文](https://docs.microsoft.com/zh-cn/dotnet/core/run-time-config/)

### 本文内容

1.  [runtimeconfig.json](#runtimeconfigjson)
2.  [MSBuild 属性](#msbuild-properties)
3.  [环境变量](#environment-variables)

.NET Core 支持使用配置文件和环境变量在运行时配置 .NET Core 应用程序的行为。.NET Core supports the use of configuration files and environment variables to configure the behavior of .NET Core applications at run time. 如果出现以下情况，则运行时配置是一个不错的选择：Run-time configuration is an attractive option if:

*   你不拥有或控制应用程序的源代码，因此无法以编程方式对其进行配置。You don't own or control the source code for an application and therefore are unable to configure it programmatically.
    
*   应用程序的多个实例在单个系统上同时运行，并且你想要将每个实例配置为获得最佳性能。Multiple instances of your application run at the same time on a single system, and you want to configure each for optimum performance.
    

备注

本文档正在编写中。This documentation is a work in progress. 如果你注意到此处提供的信息不完整或不准确，可以[创建一个问题](https://github.com/dotnet/docs/issues)告知我们，或[提交拉取请求](https://github.com/dotnet/docs/pulls)以解决问题。If you notice that the information presented here is either incomplete or inaccurate, either [open an issue](https://github.com/dotnet/docs/issues) to let us know about it, or [submit a pull request](https://github.com/dotnet/docs/pulls) to address the issue. 要了解如何提交 dotnet/docs 存储库的拉取请求，请参阅[参与者指南](https://github.com/dotnet/docs/blob/master/CONTRIBUTING.md)。For information about submitting pull requests for the dotnet/docs repository, see the [contributor's guide](https://github.com/dotnet/docs/blob/master/CONTRIBUTING.md).

.NET Core 提供以下用于配置运行时应用程序行为的机制：.NET Core provides the following mechanisms for configuring run-time application behavior:

*   [runtimeconfig.json 文件](#runtimeconfigjson)The [runtimeconfig.json file](#runtimeconfigjson)
    
*   [MSBuild 属性MSBuild properties](#msbuild-properties)
    
*   [环境变量Environment variables](#environment-variables)
    

某些配置值还可以通过调用 [AppContext.SetSwitch](/zh-cn/dotnet/api/system.appcontext.setswitch) 方法以编程方式进行设置。Some configuration values can also be set programmatically by calling the [AppContext.SetSwitch](/zh-cn/dotnet/api/system.appcontext.setswitch) method.

文档此部分的文章按类别组织，例如[调试](debugging-profiling)和[垃圾回收](garbage-collector)。The articles in this section of the documentation are organized by category, for example, [debugging](debugging-profiling) and [garbage collection](garbage-collector). 如果适用，将显示 runtimeconfig.json 文件、MSBuild 属性、环境变量的配置选项；对于 .NET Framework 项目，还会显示 app.config 文件的配置选项以便交叉引用。 Where applicable, configuration options are shown for _runtimeconfig.json_ files, MSBuild properties, environment variables, and, for cross-reference, _app.config_ files for .NET Framework projects.

runtimeconfig.jsonruntimeconfig.json[](#runtimeconfigjson)
----------------------------------------------------------

[构建](../tools/dotnet-build)项目时，将在输出目录中生成 _\[appname\].runtimeconfig.json_ 文件。When a project is [built](../tools/dotnet-build), an _\[appname\].runtimeconfig.json_ file is generated in the output directory. 如果项目文件所在的文件夹中存在 _runtimeconfig.template.json_ 文件，它包含的任何配置选项都将合并到 _\[appname\].runtimeconfig.json_ 文件中。If a _runtimeconfig.template.json_ file exists in the same folder as the project file, any configuration options it contains are merged into the _\[appname\].runtimeconfig.json_ file. 如果自行构建应用，请将所有配置选项放在 _runtimeconfig.template.json_ 文件中。If you're building the app yourself, put any configuration options in the _runtimeconfig.template.json_ file. 如果只是运行应用，请将其直接插入 \[appname\].runtimeconfig.template.json 文件中。If you're just running the app, insert them directly into the _\[appname\].runtimeconfig.json_ file.

备注

后续生成中将覆盖 \[appname\].runtimeconfig.template.json 文件。The _\[appname\].runtimeconfig.json_ file will get overwritten on subsequent builds.

在 runtimeconfig.json 文件的 configProperties 部分指定运行时配置选项 。Specify run-time configuration options in the **configProperties** section of the _runtimeconfig.json_ files. 此部分包含窗体：This section has the form:

JSON 复制

    "configProperties": {
      "config-property-name1": "config-value1",
      "config-property-name2": "config-value2"
    }
    

### 示例 \[appname\].runtimeconfig.template.json 文件Example \[appname\].runtimeconfig.json file[](#example-appnameruntimeconfigjson-file)

如果要将这些选项放在输出 JSON 文件中，请将它们嵌套在 `runtimeOptions` 属性下。If you're placing the options in the output JSON file, nest them under the `runtimeOptions` property.

JSON 复制

    {
      "runtimeOptions": {
        "tfm": "netcoreapp3.1",
        "framework": {
          "name": "Microsoft.NETCore.App",
          "version": "3.1.0"
        },
        "configProperties": {
          "System.GC.Concurrent": false,
          "System.Threading.ThreadPool.MinThreads": 4,
          "System.Threading.ThreadPool.MaxThreads": 25
        }
      }
    }
    

### 示例 runtimeconfig.template.json 文件Example runtimeconfig.template.json file[](#example-runtimeconfigtemplatejson-file)

如果要将这些选项放在模板 JSON 文件中，请省略 `runtimeOptions` 属性。If you're placing the options in the template JSON file, omit the `runtimeOptions` property.

JSON 复制

    {
      "configProperties": {
        "System.GC.Concurrent": false,
        "System.Threading.ThreadPool.MinThreads": "4",
        "System.Threading.ThreadPool.MaxThreads": "25"
      }
    }
    

MSBuild 属性MSBuild properties[](#msbuild-properties)
---------------------------------------------------

可以使用 SDK 样式 .NET Core 项目的 .csproj 或 .vbproj 文件中的 MSBuild 属性设置某些运行时配置选项。 Some run-time configuration options can be set using MSBuild properties in the _.csproj_ or _.vbproj_ file of SDK-style .NET Core projects. MSBuild 属性优先于在 _runtimeconfig.template.json_ 文件中设置的选项。MSBuild properties take precedence over options set in the _runtimeconfig.template.json_ file. 它们还会覆盖生成时在 \[appname\].runtimeconfig.json 文件中设置的任何选项。 They also overwrite any options you set in the _\[appname\].runtimeconfig.json_ file at build time.

下面是一个示例 SDK 样式项目文件，其中包含用于配置运行时行为的 MSBuild 属性：Here is an example SDK-style project file with MSBuild properties for configuring run-time behavior:

XML 复制

    <Project Sdk="Microsoft.NET.Sdk">
    
      <PropertyGroup>
        <OutputType>Exe</OutputType>
        <TargetFramework>netcoreapp3.1</TargetFramework>
      </PropertyGroup>
    
      <PropertyGroup>
        <ConcurrentGarbageCollection>false</ConcurrentGarbageCollection>
        <ThreadPoolMinThreads>4</ThreadPoolMinThreads>
        <ThreadPoolMaxThreads>25</ThreadPoolMaxThreads>
      </PropertyGroup>
    
    </Project>
    

用于配置运行时行为的 MSBuild 属性记录在每个区域各自的文章中，例如[垃圾回收](garbage-collector)。MSBuild properties for configuring run-time behavior are noted in the individual articles for each area, for example, [garbage collection](garbage-collector).

环境变量Environment variables[](#environment-variables)
---------------------------------------------------

环境变量可用于提供一些运行时配置信息。Environment variables can be used to supply some run-time configuration information. 指定为环境变量的配置旋钮通常有 COMPlus_ 前缀 。Configuration knobs specified as environment variables generally have the prefix **COMPlus_**.

可以使用 Windows 控制面板、命令行或通过在 Windows 和 Unix 系统上调用 [Environment.SetEnvironmentVariable(String, String)](/zh-cn/dotnet/api/system.environment.setenvironmentvariable#System_Environment_SetEnvironmentVariable_System_String_System_String_) 方法以编程方式定义环境变量。You can define environment variables from the Windows Control Panel, at the command line, or programmatically by calling the [Environment.SetEnvironmentVariable(String, String)](/zh-cn/dotnet/api/system.environment.setenvironmentvariable#System_Environment_SetEnvironmentVariable_System_String_System_String_) method on both Windows and Unix-based systems.

下面的示例演示如何在命令行中设置环境变量：The following examples show how to set an environment variable at the command line:

shell 复制

    # Windows
    set COMPlus_GCRetainVM=1
    
    # Powershell
    $env:COMPlus_GCRetainVM="1"
    
    # Unix
    export COMPlus_GCRetainVM=1