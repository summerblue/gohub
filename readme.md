## 说明

本项目是 [《G02 Go API 实战》](https://learnku.com/courses/go-api/) 实战课程的源码，基于 MIT 协议开源。


项目名称 Gohub，以论坛 API 为主题，设计的初衷是将其打造为高性能、功能齐全的 API 框架。

程序结构说明，请见 [程序结构](https://learnku.com/courses/go-api/1.17/program-structure/11772)。

## RESTful API 最佳实践

一套优秀的 API 设计，需要具备如下特性：

1. 使用 HTTPS
2. 使用域名
3. 版本区分
4. 使用 URL 来定位资源
5. 使用 HTTP 动词来描述操作
6. 支持资源过滤
7. 使用 HTTP 状态码
8. 数据响应的一致性
9. 支持限流
10. API 文档
11. 自带分页链接
12. 强制 User-Agent
 
详细讲解请见 [RESTful API 最佳实践](https://learnku.com/courses/go-api/1.17/api-development-best-practices/11769) 。

## 所有路由

| 请求方法  |  API 地址  | 说明  |
| ------------ | ------------ | ------------ |
| POST |    /api/v1/auth/login/using-phone | 短信 + 手机号登录 |
| POST |    /api/v1/auth/login/using-password | 手机号、用户名、邮箱 + 密码 |
| POST |    /api/v1/auth/login/refresh-token | 刷下 Token |
| POST |    /api/v1/auth/password-reset/using-email | 邮件密码重置 |
| POST |    /api/v1/auth/password-reset/using-phone | 短信验证码密码重置 |
| POST |    /api/v1/auth/signup/using-phone | 使用手机号注册 |
| POST |    /api/v1/auth/signup/using-email | 使用邮箱注册 |
| POST |    /api/v1/auth/signup/phone/exist | 手机号是否已注册 |
| POST |    /api/v1/auth/signup/email/exist | email 是否已支持 |
| POST |    /api/v1/auth/verify-codes/phone | 发送短信验证码 |
| POST |    /api/v1/auth/verify-codes/email | 发送邮件验证码 |
| POST |    /api/v1/auth/verify-codes/captcha | 获取图片验证码 |
| GET |     /api/v1/user              | 获取当前用户 |
| GET |     /api/v1/users             | 用户列表 |
| PUT |     /api/v1/users             | 修改个人资料 |
| PUT |     /api/v1/users/email       | 修改邮箱 |
| PUT |     /api/v1/users/phone       | 修改手机号 |
| PUT |     /api/v1/users/password    | 修改密码 |
| PUT |     /api/v1/users/avatar      | 上传头像 |
| GET |     /api/v1/categories        | 分类列表 |
| POST |    /api/v1/categories        | 创建分类 |
| PUT |     /api/v1/categories/:id    | 更新分类 |
| DELETE |  /api/v1/categories/:id    | 删除分类 |
| GET |     /api/v1/topics            | 话题列表 |
| POST |    /api/v1/topics            | 创建话题 |
| PUT |     /api/v1/topics/:id        | 更新话题  |
| DELETE |  /api/v1/topics/:id        | 删除话题 |
| GET |     /api/v1/topics/:id        | 获取话题 |
| GET |     /api/v1/links             | 友情链接列表 |


## 第三方依赖

使用到的开源库：

- [gin](https://github.com/gin-gonic/gin) —— 路由、路由组、中间件
- [zap](https://github.com/gin-contrib/zap) —— 高性能日志方案
- [gorm](https://github.com/go-gorm/gorm) —— ORM 数据操作
- [cobra](https://github.com/spf13/cobra) —— 命令行结构
- [viper](https://github.com/spf13/viper) —— 配置信息
- [cast](https://github.com/spf13/cast) —— 类型转换
- [redis](https://github.com/go-redis/redis/v8) —— Redis 操作
- [jwt](https://github.com/golang-jwt/jwt) —— JWT 操作
- [base64Captcha](https://github.com/mojocn/base64Captcha) —— 图片验证码
- [govalidator](https://github.com/thedevsaddam/govalidator) —— 请求验证器
- [limiter](https://github.com/ulule/limiter/v3) —— 限流器
- [email](https://github.com/jordan-wright/email) —— SMTP 邮件发送
- [aliyun-communicate](https://github.com/KenmyZhang/aliyun-communicate) —— 发送阿里云短信
- [ansi](https://github.com/mgutz/ansi) —— 终端高亮输出
- [strcase](https://github.com/iancoleman/strcase) —— 字符串大小写操作
- [pluralize](https://github.com/gertd/go-pluralize) —— 英文字符单数复数处理


## 自定义的包

现在来看下我们自建的库：

- app —— 应用对象
- auth —— 用户授权
- cache —— 缓存
- captcha —— 图片验证码
- config —— 配置信息
- console —— 终端
- database —— 数据库操作
- file —— 文件处理
- hash —— 哈希
- helpers —— 辅助方法
- jwt —— JWT 认证
- limiter —— API 限流
- logger —— 日志记录
- mail —— 邮件发送
- migrate —— 数据库迁移
- paginator —— 分页器
- redis —— Redis 数据库操作
- response —— 响应处理
- seed —— 数据填充
- sms —— 发送短信
- str —— 字符串处理
- verifycode —— 数字验证码


## 代码行数

Gohub 项目总共有 4600 行代码（工具 [gocloc](https://github.com/hhatto/gocloc)）：

```
$ gocloc .
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                             122           1200            865           4629
TOML                             1              7             21             28
-------------------------------------------------------------------------------
TOTAL                          123           1207            886           4657
-------------------------------------------------------------------------------
```

## 所有命令

```
$ go run main.go -h
Default will run "serve" command, you can use "-h" flag to see all subcommands

Usage:
   [command]

Available Commands:
  cache       Cache management
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  key         Generate App Key, will print the generated Key
  make        Generate file and code
  make        Generate file and code
  migrate     Run database migration
  play        Likes the Go Playground, but running at our application context
  seed        Insert fake data to the database
  serve       Start web server

Flags:
  -e, --env string   load .env file, example: --env=testing will use .env.testing file
  -h, --help         help for this command

Use " [command] --help" for more information about a command.
```

make 命令：

```
$ go run main.go make -h
Generate file and code

Usage:
   make [command]

Available Commands:
  apicontroller Create api controller，exmaple: make apicontroller v1/user
  cmd           Create a command, should be snake_case, exmaple: make cmd buckup_database
  factory       Create model's factory file, exmaple: make factory user
  migration     Create a migration file, example: make migration add_users_table
  model         Crate model file, example: make model user
  policy        Create policy file, example: make policy user
  request       Create request file, example make request user
  seeder        Create seeder file, example:  make seeder user

Flags:
  -h, --help   help for make

Global Flags:
  -e, --env string   load .env file, example: --env=testing will use .env.testing file

Use " make [command] --help" for more information about a command.
```

migrate 命令：

```
$ go run main.go migrate -h
Run database migration

Usage:
   migrate [command]

Available Commands:
  down        Reverse the up command
  fresh       Drop all tables and re-run all migrations
  refresh     Reset and re-run all migrations
  reset       Rollback all database migrations
  up          Run unmigrated migrations

Flags:
  -h, --help   help for migrate

Global Flags:
  -e, --env string   load .env file, example: --env=testing will use .env.testing file

Use " migrate [command] --help" for more information about a command.
```