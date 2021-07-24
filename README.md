# 基于Gin开发的API框架

# 1、介绍

基础Gin框架开发的一套API框架，组装了数据库操作方法，API返回格式，并带有基础API检验

# 2、目录

conf:配置文件

controllers:控制器

db:数据库

doc:文档

models:数据库

static:图片、样式目录

# 3、开发前准备

#### 安装Go以及配置GOPATH、GOROOT

#### 由于使用了mysql、腾讯云COS，防止出错，请在开发前安装下面4个扩展包
```
go get -u github.com/gin-gonic/gin

go get -u github.com/tencentyun/cos-go-sdk-v5

go get -u github.com/jinzhu/gorm

go get -u github.com/go-sql-driver/mysql
```

#### 修改conf目录下的config配置文件

数据库信息以及配置腾讯云cos配置

# 4、编译运行

####运行命令
```
go run ./
```
然后访问

http://localhost:9999

####打包成文件

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