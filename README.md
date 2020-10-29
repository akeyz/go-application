# 打包基于chrome的全屏应用

### chrome必须已经安装

### app ico
https://icon-icons.com/zh/

### 下载安装rsrc
```
go get github.com/akavel/rsrc
```

### 生成程序描述文件ico.manifest
```
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<assembly xmlns="urn:schemas-microsoft-com:asm.v1" manifestVersion="1.0">
<assemblyIdentity
    version="1.0.0.0"
    processorArchitecture="x86"
    name="controls"
    type="win32"
></assemblyIdentity>
<dependency>
    <dependentAssembly>
        <assemblyIdentity
            type="win32"
            name="Microsoft.Windows.Common-Controls"
            version="6.0.0.0"
            processorArchitecture="*"
            publicKeyToken="6595b64144ccf1df"
            language="*"
        ></assemblyIdentity>
    </dependentAssembly>
</dependency>
</assembly>
```

### 生成go程序嵌入文件
rsrc.exe -manifest ico.manifest -o myapp.syso -ico myapp.ico

### 编译Go程序
将myapp.syso文件放到相应go程序下，然后直接运行go build .即可。

golang已经可以自动寻找子目录下的 syso 文件。

### gcc安装64bit报错
https://blog.csdn.net/mecho/article/details/24305369