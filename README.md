# Beego_learn
Beego框架学习案列

## 快速开始

### 安装

安装beego和bee工具。

```bash
$ go get -u github.com/beego/beego/v2
$ go get -u github.com/beego/bee/v2
```

### 创建和运行项目

```bash
$ bee new hello
$ cd hello
$ bee run
```

一旦程序开始运行，您就可以在浏览器中打开 http://localhost:8080/ 进行访问。

### 静态文件处理

用户可以设置多个静态文件处理目录，例如你有多个文件下载目录 download1、download2，你可以这样映射（在 /main.go 文件中 web.Run() 之前加入）：

```go
beego.SetStaticPath("/download1", "download1")
beego.SetStaticPath("/down2", "download2")
```

这样用户访问 URL http://localhost:8080/down1/123.txt 则会请求 download1 目录下的 123.txt 文件。

## MVC架构

### controller 设计

#### 参数配置

beego 默认会解析当前应用下的 `conf/app.conf` 文件。

通过这个文件你可以初始化很多 beego 的默认参数：

```conf
appname = beepkg # 应用名称
httpaddr = "127.0.0.1"  # 运行地址
httpport = 9090 # 运行端口
runmode ="dev"  # 运行模式
autorender = false
recoverpanic = false
viewspath = "myview"
```

它们都维护在结构体 `beego/server/web#Config` 。

你也可以在配置文件中配置应用需要用的一些配置信息，例如下面所示的数据库信息：

```
mysqluser = "root"
mysqlpass = "rootpass"
mysqlurls = "127.0.0.1"
mysqldb   = "beego"
```

那么你就可以通过如下的方式获取设置的配置信息:

```go
beego.AppConfig.String("mysqluser")
beego.AppConfig.String("mysqlpass")
beego.AppConfig.String("mysqlurls")
beego.AppConfig.String("mysqldb")
```

