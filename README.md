# 基于Gin开发的API项目

# 1、介绍

基础Gin框架开发的一套API项目，组装了数据库操作方法，API返回格式，并带有基础API检验

# 2、目录

conf:配置文件

controllers:控制器

db:数据库

doc:文档

models:数据库

static:图片、样式目录

# 3、开发前准备

#### 安装Go以及配置GOPATH、GOROOT

[Win环境下安装Go并配置环境变量](https://hongzx.cn/home/blogShow/131)

[Mac环境下安装Go并配置环境变量](https://hongzx.cn/home/blogShow/134)

#### 由于使用了mysql、腾讯云COS、平滑重启，防止出错，请在开发前安装下面4个扩展包
```
go get -u github.com/gin-gonic/gin

go get -u github.com/tencentyun/cos-go-sdk-v5

go get -u github.com/jinzhu/gorm

go get -u github.com/go-sql-driver/mysql

go get -u github.com/fvbock/endless
```

#### 修改conf目录下的config配置文件

- 修改数据库信息以及配置腾讯云cos配置

- 项目默认端口是9999，如果端口被占用，请修改config配置文件的端口

# 4、编译运行

#### 运行命令
```
go run ./
```
然后访问

http://localhost:9999

#### 打包成文件

 Win环境下打包编译：
 
``` 
set GOARCH=amd64 

set GOOS=linux

go build
```
 
 Mac环境下打包编译：
 ```
 go env -w GOARCH=amd64
 
 go env -w GOOS=linux
 
 go build
```
 # 5、部署并通过Nginx转发
 [Gin编译部署到centos并搭配Nginx运行](https://hongzx.cn/home/blogShow/155)
 
 # 6、平滑重启
 新版本加入了平滑重启功能，暂不支持WIN环境。
 
 在Linux环境下使用以下命令
 ```
 后台守护进程【首次使用】:
 nohup ./goAPI&
 
 平滑重启【每次迭代必用】
 kill -1 PID
 
 ```
 
 
 
 # 7、后续版本
 下一版本打算加入API令牌