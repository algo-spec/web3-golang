# 说明
- [说明](#说明)
- [运行环境](#运行环境)
- [依赖安装步骤](#依赖安装步骤)
  - [安装gorm](#安装gorm)
  - [安装gorm-mysql](#安装gorm-mysql)
  - [安装gin](#安装gin)
  - [安装jwt](#安装jwt)
  - [安装viper【配置管理】](#安装viper配置管理)
- [目录结构](#目录结构)
- [启动方式](#启动方式)
- [功能列表](#功能列表)
- [接口文档](#接口文档)
  - [1. 用户注册](#1-用户注册)
  - [2. 用户登录](#2-用户登录)
  - [3. 获取所有文章](#3-获取所有文章)
  - [4. 获取指定文章](#4-获取指定文章)
  - [5. 创建文章](#5-创建文章)
  - [6. 更新文章](#6-更新文章)
  - [7. 删除文章](#7-删除文章)
  - [8. 获取指定文章评论](#8-获取指定文章评论)
  - [9. 创建评论](#9-创建评论)


使用 Go 语言结合 Gin 框架和 GORM 库开发一个个人博客系统的后端，实现博客文章的基本管理功能，包括文章的创建、读取、更新和删除（CRUD）操作，同时支持用户认证和简单的评论功能

# 运行环境
```
GO 1.24+
MYSQL 8.1+
GIN 1.7+
GORM 1.22+
```

# 依赖安装步骤
## 安装gorm
```
go get -u github.com/jinzhu/gorm
```
## 安装gorm-mysql
```
go get -u gorm.io/driver/mysql
```
## 安装gin
```
go get -u github.com/gin-gonic/gin
```
## 安装jwt
```
go get -u github.com/golang-jwt/jwt/v5
```
## 安装viper【配置管理】
```
go get -u github.com/spf13/viper
```

# 目录结构

```
web3-golang
├─ 📁task04
   ├─ 📁app                   # 应用层
   │  ├─ 📁handlers           # HTTP请求处理器层
   │  ├─ 📁middleware         # 中间件
   │  └─ 📁routes             # 路由层
   │  │
   ├─ 📁domain                # 领域层
   │  ├─ 📁models             # 数据模型
   │  ├─ 📁repositories       # 数据仓库
   │  └─ 📁services           # 业务逻辑
   │  │
   ├─ 📁pkg                   # 公共库
   │  ├─ 📁config             # 配置   
   │  └─ 📁util               # 工具
   │  │
   ├─ 📄README.md
   ├─ 📄go.mod                # 模块依赖
   └─ 📄main.go               # 程序入口
```

# 启动方式
在task04目录下执行如下命令启动项目：
```
go run main.go
```

# 功能列表
| URL               | 请求方式 | 描述             |
|-------------------|----------|----------------|
| /user/register    | POST     | 用户注册         |
| /user/login       | POST     | 用户登录         |
| /post/            | GET      | 获取所有文章     |
| /post/:id         | GET      | 获取指定文章     |
| /post/create      | POST     | 创建文章         |
| /post/:id         | PUT      | 更新文章         |
| /post/:id         | DELETE   | 删除文章         |
| /comment/:post_id | GET      | 获取指定文章评论 |
| /comment/create   | POST     | 创建评论         |

# 接口文档
## 1. 用户注册
**请求方式：POST**

**URL：http://localhost:8080/user/register**

**请求参数**
```
{
    "username": "test1",
    "password": "test1",
    "email": "test1@163.com"
}
```
**响应参数**
```
{
    "message": "User created successfully"
}
```

## 2. 用户登录
**请求方式：POST**

**URL：http://localhost:8080/user/login**

**请求参数**
```
{
    "username": "test1",
    "password": "test1"
}
```
**响应参数**
```
{
    "message": "success",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsIlVzZXJuYW1lIjoidGVzdDEiLCJpc3MiOiJ3ZWIzLWdvbGFuZyIsImV4cCI6MTc1NTA1NjMzOX0.bxVg4hACIlYyBujY8HX56d9f77a1dQiyOHBWWuZU5a8"
}
```

## 3. 获取所有文章
**请求方式：GET**

**URL：http://localhost:8080/post**

**响应参数**
```
{
    "data": [
        {
            "ID": 1,
            "Title": "title-1",
            "Content": "aaaaaaaaaaaaaaaaabbbbbbbbbbbbvvvvvvvvv",
            "UserID": 1,
            "User": {
                "ID": 1,
                "Username": "test1",
                "Password": "$2a$10$.RhbPyFxJnp8Vh9qy9On8OOCduNwaJ4.5f7Pkr52vy.2xmVzJ4THq",
                "Email": "test1@163.com",
                "Posts": null,
                "CreatedAt": "2025-08-12T11:38:55.494+08:00",
                "UpdatedAt": "2025-08-12T11:38:55.494+08:00"
            },
            "Comments": null,
            "CreatedAt": "2025-08-12T13:11:40.383+08:00",
            "UpdatedAt": "2025-08-12T13:11:40.383+08:00"
        },
        {
            "ID": 2,
            "Title": "title-1",
            "Content": "aaaaaaaaaaaaaaaaabbbbbbbbbbbbvvvvvvvvv",
            "UserID": 1,
            "User": {
                "ID": 1,
                "Username": "test1",
                "Password": "$2a$10$.RhbPyFxJnp8Vh9qy9On8OOCduNwaJ4.5f7Pkr52vy.2xmVzJ4THq",
                "Email": "test1@163.com",
                "Posts": null,
                "CreatedAt": "2025-08-12T11:38:55.494+08:00",
                "UpdatedAt": "2025-08-12T11:38:55.494+08:00"
            },
            "Comments": null,
            "CreatedAt": "2025-08-12T13:15:16.882+08:00",
            "UpdatedAt": "2025-08-12T13:15:16.882+08:00"
        }
    ],
    "message": "success"
}
```

## 4. 获取指定文章
**请求方式：GET**

**URL：http://localhost:8080/post/1**

**响应参数**
```
{
    "data": {
        "ID": 1,
        "Title": "title-1",
        "Content": "aaaaaaaaaaaaaaaaabbbbbbbbbbbbvvvvvvvvv",
        "UserID": 1,
        "User": {
            "ID": 1,
            "Username": "test1",
            "Password": "$2a$10$.RhbPyFxJnp8Vh9qy9On8OOCduNwaJ4.5f7Pkr52vy.2xmVzJ4THq",
            "Email": "test1@163.com",
            "Posts": null,
            "CreatedAt": "2025-08-12T11:38:55.494+08:00",
            "UpdatedAt": "2025-08-12T11:38:55.494+08:00"
        },
        "Comments": null,
        "CreatedAt": "2025-08-12T13:11:40.383+08:00",
        "UpdatedAt": "2025-08-12T13:11:40.383+08:00"
    },
    "message": "success"
}
```

## 5. 创建文章
**请求方式：POST**

**URL：http://localhost:8080/post/create**

**请求头**
```
Content-Type: application/json
Authorization: Bearer <token>
```
**请求参数**
```
{
    "title": "title-1",
    "content": "aaaaaaaaaaaaaaaaabbbbbbbbbbbbvvvvvvvvv"
}
```
**响应参数**
```
{
    "message": "success"
}
```

## 6. 更新文章
**请求方式：PUT**

**URL：http://localhost:8080/post/1**

**请求头**
```
Content-Type: application/json
Authorization: Bearer <token>
```
**请求参数**
```
{
    "id": 1,
    "title": "update-1",
    "content": "update-content-111111111111"
}
```
**响应参数**
```
{
    "message": "success"
}
```


## 7. 删除文章
**请求方式：DELETE**

**URL：http://localhost:8080/post/1**

**请求头**
```
Content-Type: application/json
Authorization: Bearer <token>
```
**响应参数**
```
{
    "message": "删除成功"
}
```

## 8. 获取指定文章评论
**请求方式：GET**

**URL：http://localhost:8080/comment/2**

**响应参数**
```
{
    "data": [
        {
            "id": 1,
            "content": "this is test comment",
            "post_id": 2,
            "user_id": 1,
            "created_at": "2025-08-12T13:30:21.15+08:00",
            "updated_at": "2025-08-12T13:30:21.15+08:00"
        },
        {
            "id": 2,
            "content": "this is test comment1",
            "post_id": 2,
            "user_id": 1,
            "created_at": "2025-08-12T13:30:41.408+08:00",
            "updated_at": "2025-08-12T13:30:41.408+08:00"
        }
    ],
    "message": "success"
}
```

## 9. 创建评论
**请求方式：POST**

**URL：http://localhost:8080/comment/create**

**请求头**
```
Content-Type: application/json
Authorization: Bearer <token>
```

**请求参数**
```
{
    "post_id": 2,
    "content": "this is test comment1"
}
```
**响应参数**
```
{
    "message": "success"
}
```
