
## web框架 gin
```shell
go get -u github.com/gin-gonic/gin v1.9.1
```

## 配置文件读取

```shell
go get -u github.com/spf13/viper v1.16.0
```

## 数据库框架
gorm 数据库操作框架

数据库驱动： gorm.io/driver/mysql

```shell

go get -u gorm.io/gorm v1.25.2
go get -u gorm.io/driver/mysql v1.5.1

```

## 日志框架

```shell
go get -u  gopkg.in/natefinch/lumberjack.v2 v2.2.1
```
## curl
curl是一种命令行工具，用于发送HTTP请求和接收响应。它可用于测试Web应用程序和API，以及从命令行下载文件等任务。以下是一些常用的curl命令： 

1. 发送GET请求:
```shell
curl http://example.com
```

2. 发送POST请求:
```shell
curl -X POST -d "param1=value1&param2=value2" http://example.com
```

3. 发送带有HTTP头的请求:
```shell
curl -H "Content-Type: application/json" http://example.com
```

4. 保存响应到文件:
```shell
curl -o output.txt http://example.com
```
5. 发送带有身份验证的请求:
```shell
curl -u username:password http://example.com
```
6. 显示详细的请求和响应
它会打印出请求头、响应头以及其他相关的信息，帮助进行调试和故障排除。
```shell
curl -v http://example.com
```

## Swagger接口文档生成

## 接口参数验证

```shell
go get github.com/go-playground/validator/v10

```