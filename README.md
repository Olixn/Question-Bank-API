# Question-Bank-API  
> 一个基于gin+gorm+redigo的简单题库API项目，主要功能：题库查询、更新、增加，IP流量控制。

# 如何使用
- 下载已编译好的exe可执行文件  
- 配置目录下的`config.yaml`  
```yaml 
server:
    port: 9090 #http端口
mysql:
    address: 127.0.0.1 # MySQL 数据库地址
    port: 3306 #MySQL 数据库端口
    dbname: #MySQL 数据库名
    username: #MySQL 数据库用户名
    password: #MySQL 数据库密码
redis:
    address: 127.0.0.1 #Redis 数据库地址
    port: 6379 #Redis 数据库端口
    maxIdle: 8
    maxActive: 0
    idleTimeout: 100
```
- 启动MySQl服务、启动Redis服务
- 启动exe程序

# 如何自己开发编译
- clone本仓库到本地
- 安装配置好go环境
- 进入该项目根目录，在终端运行`go mod tidy`
- 开始您自己的开发
- 进入该项目根目录，执行`build.bat`
- 等待编译完成，成品存储于`server`目录下

# 项目API接口
## 查询
```text
根据问题返回答案
```
#### 接口状态
> 已完成

#### 接口URL
> 127.0.0.1:9090/v1/api/answer?q=这个程序的作者是谁？

#### 请求方式
> GET

#### Content-Type
> form-data

#### 请求Query参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
q | 这个程序的作者是谁？ | Text | 是 | -

#### 返回成功示例
```json
{
  "status":1,
  "msg": "ok",
  "data": {
    "success": "true",
    "answer": "Ne-21",
    "data": "Ne-21"
  }
}
```

## 新增与更新
```text
传入问题与答案，自动判断新增还是更新，并执行相应操作
```
#### 接口状态
> 已完成

#### 接口URL
> 127.0.0.1:9090/v1/api/update

#### 请求方式
> POST

#### Content-Type
> urlencoded

#### 请求Body参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
q | 这个程序的作者是谁？ | Text | 是 | -
a | Ne-21 | Text | 是 | -

#### 返回成功示例
```json
{
  "status":1,
  "msg": "新增题目",
  "data": {}
}
```
> 作者：Ne-21  
> 博客：https://blog.gocos.cn