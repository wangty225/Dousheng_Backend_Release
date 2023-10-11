# Dousheng_Backend

---

1. [青训营大项目答辩汇报文档](https://rx4174uz0we.feishu.cn/docx/GpXWdrbqvoMH14xt171cpiOdnof)

2. [快速开始（环境搭建与项目部署）](QuickStart.md)

3. 项目接口在线演示地址：http://123.57.251.188:8888/

4. 演示视频1：# 短视频后台开发测试：[Postman演示视频](https://www.bilibili.com/video/BV13w411U7dY/)

5. 演示视频2:   # 短视频后台开发测试：[Android演示视频](https://www.bilibili.com/video/BV1yj41117s9/)

6. 获奖证书：
   
   ![](docs\imgs\prize\awards.png)

---

# 一、项目介绍

> 项目核心：本项目基于hertz，kitex框架，加入jwt鉴权，zap日志，Auth等中间件等，使用mysql存储和redis缓存，最终实现了微服务架构的抖音后台开发系统。
> 
> 项目服务地址：http://123.57.251.188:8888/
> 
> 项目地址：https://github.com/wangty225/Dousheng_Backend_Release
> 
> 项目小组：摸鱼佛系队

# 二、项目分工

| **团队成员**        | **主要贡献**                                                                                                           |
| --------------- | ------------------------------------------------------------------------------------------------------------------ |
| 王天宇<br><br>（组长） | 项目技术选型，项目架构；<br><br>基础接口（用户信息，注册，登录，投稿，视频流，投稿列表）；<br><br>互动接口（点赞，点赞列表，评论，评论列表）；<br><br>docker部署，性能测试，开发文档撰写，项目演示等。 |
| 何秋锦             | 关系相关接口的路由逻辑实现（暂未完善，功能未上线）；                                                                                         |

# 三、项目实现

### 3.1 技术选型与相关开发文档

#### 3.1.1 项目目录结构

> 项目采用微服务架构实现，具体在实现时，首先对微服务进行边界的划分，确立微服务的业务逻辑范围以及微服务技术选型等。具体的业务逻辑划分如下：

```bash
# 微服务接口划分
user_service:  # 已实现
  ├─ /douyin/user/register/         # 用户注册接口   
  ├─ /douyin/user/login/            # 用户登录接口
  └─ /douyin/user/                  # 用户信息

video_service: # 已实现
  ├─ /douyin/feed/                  # 视频流接口
  ├─ /douyin/publish/action/        # 视频投稿
  └─ /douyin/publish/list/          # 发布列表

interaction_service: # 已实现
  ├─ /douyin/favorite/action/       # 赞操作
  ├─ /douyin/favorite/list/         # 喜欢列表
  ├─ /douyin/comment/action/        # 评论操作
  └─ /douyin/comment/list/          # 评论列表

relation_service:
  ├─ /douyin/relation/action/         # 关系操作
  ├─ /douyin/relation/follow/list/    # 用户关注列表
  ├─ /douyin/relation/follower/list/  # 用户粉丝列表
  └─ /douyin/relation/friend/list/    # 用户好友列表

message_service:
  ├─ /douyin/message/chat/          # 聊天记录    
  └─ /douyin/message/action/        # 消息操作
```

#### 3.1.2 微服务边界划分

  基础功能：**`user_service`** 、**`video_service`**

  自选功能：**`interaction_service`** 、`relation_service` 、`message_service`

#### 3.1.3 技术选型

1. 选择kitex作为后端开发框架，使用kitex-gen生成微服务的服务逻辑；

2. 选择hertz作为http框架进行路由和远程调用等，使用hertz-gen生成框架的路由逻辑；

3. 选择mysql作为文本数据的存储类型；

4. 选择AliyunOss作为云对象存储数据库；

5. 选择redis作为mysql数据访问缓存层，以及鉴权等相关操作；

6. 选择etcd作为微服务注册与发现组件；

7. 选择zap管理日志文件；

8. 选择viper管理配置文件；

9. 选择nginx实现容器反向代理，以及ECS反向代理OSS

10. 选择Docker和Docker-Compose管理容器，实现微服务部署；

11. 选择AliyunECS 1核2G 作为开发，测试，部署，反向代理服务器；

#### 3.1.4 主要依赖版本汇总

工具版本情况汇总：

```bash
kitex        v0.5.0     // go: import
hertz        v0.6.7     // go: import
mysql        :5.7       // docker deploy
etcd         :3.5.9     // docker deploy
redis        :7.0.12    // docker deploy
nginx        :alpine   // docker deploy
docker       v24.0.5    
AliyunOss      - 
...
```

go.mod主要依赖版本情况汇总：

```bash
github.com/cloudwego/hertz v0.6.7
github.com/cloudwego/kitex v0.5.0
gorm.io/driver/mysql v1.5.1
gorm.io/gorm v1.25.3
github.com/go-sql-driver/mysql v1.7.0
github.com/go-redis/redis/v8 v8.11.5
github.com/spf13/viper v1.16.0
go.etcd.io/etcd/client/v3 v3.5.9
go.uber.org/zap v1.25.0
github.com/apache/thrift v0.13.0   // 保证兼容性
github.com/aliyun/aliyun-oss-go-sdk v2.2.9+incompatible  // aliyun_oss 引入

github.com/gin-gonic/gin v1.9.1
github.com/golang-jwt/jwt v3.2.2+incompatible
github.com/hertz-contrib/gzip v0.0.1
github.com/hertz-contrib/pprof v0.1.0
github.com/kitex-contrib/obs-opentelemetry v0.1.0
golang.org/x/crypto v0.11.0
gopkg.in/natefinch/lumberjack.v2 v2.0.0
gorm.io/plugin/opentracing v0.0.0-20211220013347-7d2b2af23560
```

### 3.2 架构设计

#### 3.2.1 项目总体设计架构图

![Project_Framework.png](docs%2Fimgs%2FProject_Framework.png)

#### 3.2.2 mysql设计

![ER-mysql.PNG](docs%2Fimgs%2FER-mysql.PNG)

#### 3.2.3 架构优化设计

##### 3.2.3.1 用户登录：基于Token的用户鉴权过程设计

>     通常情况下，用户在进行登录操作时，需要进行密码的输入，但是用户在进行其他操作时，显然不可以总是携带密码判断。携带密码访问最严重的问题：
> 
> ①多次传输导致密码泄露或者被被截获等；
> 
> ②对密码的验证逻辑属于“静态”的，即每一次验证后端都是相同的过程，需要重复调用浪费性能：如`传入passwd1`密码加密，查询数据库的`加密passwd`，对比验证有效性。如果是反复的验证，会对数据库访问等造成巨大压力；
> 
> ③可能存在伪造请求而越权访问等。
> 
>     而**token**的引入则带来了巨大的便利性和安全性，和更好的用户体验。
> 
> 1. 可以减少密码泄露风险，避免密码可能会被截获、窃取或猜测等安全漏洞，而令牌则可以通过设置超时时间，使得越权访问得到限制。
> 
> 2. 分离身份验证和授权：令牌通常只包含访问令牌所需的信息，而不包含用户的实际密码。这意味着在授权系统中，不需要存储用户的敏感密码，因此即使授权系统受到攻击，用户密码也不会被泄露。
> 
> 3. 可扩展性：适用于多种应用场景，支持单点登录等，用户可以一次登录后，在多个关联应用中访问资源，提供了更好的用户体验。

1. **问题分析**

>         token设计时是需要设计过期时间，那么会造成用户可能在产生token后，在进行系统使用时候token正好过期，导致操作失效，影响用户体验，尤其是在token时效较短的情况下，如HttpSession对象只有15min有效期，就需要用户频繁重新登陆获取token，显然会极大影响用户体验。而解决失效的方法则应该是：**刷新token。**

![auth_middleware.png](docs%2Fimgs%2Fauth_middleware.png)

在实际进行刷新token时候，又要注意三种情况：

    I. 对于已经过期较长时间，且没有刷新token的情况（可能包含被攻击时截获的token，过期后仍在尝试请求）

    II. 刚刚过期的token，用户可能很快会下次访问；

    III. 未过期的token，用户正常请求。

对于情况II比较好理解，且无需刷新token；而对于情况I和II则需要判断是否需要刷新token。判断是否刷新的有两种方法：

    I. **双Token法**：Token中除了携带user_id以外，同时携带过期时间和最迟刷新时间；

    II. **Token定时缓存法**：如在redis中存入token，并使用最迟刷新时长作为键值的过期时长。

    本文选择redis的db0作为token的缓存层，使用Token过期时长的2倍作为redis的键值过期时间，通过Token缓存，来实现Token的短过期刷新。

![Token_Redis.png](docs%2Fimgs%2FToken_Redis.png)

##### 3.2.3.2 用户敏感信息（密码）加盐加密存储方式

**提高安全性措施一：单独设计Auth表存储用户名和密码。**

1. 安全性：
   
   1. 单独设计用户名和密码的表：将用户名和密码存储在单独的表中可以提高安全性，因为密码通常需要额外的保护，如哈希和加盐，以减少泄露的风险。保证密码等敏感数据的非必要不妨问性，防止因表格操作失误造成的可能的密码泄露问题。（如若将密码相关信息存入user表中，而sql中select和update，可能会因为操作错误列导致敏感数据的危险，当然这属于严重事故了，开发时一般不建议使用select * 等）
   
   2. 统一放在用户信息表里：将密码存储在用户信息表中可能会增加安全风险，因为数据库中的用户信息可能会被更多人访问，而不仅仅是身份验证系统。

2. 扩展性：
   
   1. 单独设计用户名和密码的表：如果将用户名和密码分开存储，您可以更轻松地扩展身份验证系统，例如添加多因素身份验证或支持不同类型的用户身份验证。
   
   2. 统一放在用户信息表里：将用户名和密码存储在用户信息表中可能更简单，但在将来需要增加身份验证功能时，可能需要更多工作。

**提高安全措施二：给密码增加随机盐值，在进行哈希散列加密**

        每次创建新用户时，可以生成新的随机盐值，并将其与用户的密码拼接起来，再进行MD5加密存储等，可以保证Auth表的独立性和可扩展性；

        即使加密存储密码泄露，也可以有效防止彩虹表攻击，提高密码复杂性，增加攻击者的成本等。

![table_auth.png](docs%2Fimgs%2Ftable_auth.png)

##### 3.2.3.3 oss对象存储反向代理

        由于阿里云oss公网流量付费较贵，而内网流出流量免费，所以选择使用ecs上的nginx方向代理oss，这样既可以实现内网流量转发公网，又可以保证oss的安全性。

![oss_proxy.png](docs%2Fimgs%2Foss_proxy.png)

```Bash
vim aliyun_oss.conf
```

```bash
upstream ossproxy{
    server oss-cn-beijing-internal.aliyuncs.com;
}
server {
    listen 22441;

    location /nginx_status/ {
        stub_status on;
        access_log off;
        allow 127.0.0.1;
        deny all;
    }

    location /oss-dousheng/ {
        rewrite ^/oss-dousheng(/.*)$ $1 break;
        proxy_pass https://oss-dousheng.oss-cn-beijing-internal.aliyuncs.com;

        proxy_redirect off;
        proxy_set_header Host oss-dousheng.oss-cn-beijing-internal.aliyuncs.com; 
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_set_header X-Forwarded-Host $host;
        proxy_set_header X-Forwarded-Server $host;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "Upgrade";
        proxy_read_timeout 1d;
        proxy_send_timeout 1d;
        # 此处配置是为了把阿里云oss的返回的Content-Disposition隐藏掉，否则图片或文件在浏览器中会下载，不会显示。
        proxy_hide_header Content-Disposition;
    }
}
```

```bash
docker run -d --restart=unless-stopped --network app-tier --name oss-proxy -p:22441:22441 -v `pwd`/aliyun_oss.conf:/etc/nginx/conf.d/default.conf nginx:alpine
```

##### 3.2.3.4 环境搭建和部署流程

1. **安装docker环境**
   
   i. 安装当然是要看官方说明文档啦！
   
   [ubuntu安装docker](https://docs.docker.com/engine/install/ubuntu/)
   
   ii. 创建自定义docker网络，方便后续管理容器的ip地址
   
   ```bash
   docker network create --subnet=192.168.1.0/24 app-tier
   ```

2. **安装mysql** **5.7**
   
   ```bash
   docker run -d -p 3306:3306 --name mysql --network app-tier -e MYSQL_ROOT_PASSWORD=dousheng_WSX mysql:5.7
   docker exec -it mysql bash
   mysql -u mysql -p
   create database dousheng_db;
   # do some sqls
   ```

3. **安装redis**
   
   ```bash
   sudo docker pull redis
   docker inspect redis | grep -i version
   wget http://download.redis.io/releases/redis-7.0.2.tar.gz
   tar xzf redis-7.0.2.tar.gz 
   mkdir -p /mnt/dockerspace/redis/config/
   cd redis-7.0.2/
   cp redis.conf  /mnt/dockerspace/redis/config/
   cd /mnt/dockerspace/redis/config/
   vim redis.conf 
   ```
   
   ```go
   bind 127.0.0.1 #注释掉这部分，这是限制redis只能本地访问
   
   protected-mode no #默认yes，开启保护模式，限制为本地访问
   
   daemonize no  #默认no，改为yes意为以守护进程方式启动，可后台运行，除非kill进程，改为yes会使配置文件方式启动redis失败
   
   databases 16 #数据库个数（可选），我修改了这个只是查看是否生效。。
   
   dir  ./ #输入本地redis数据库存放文件夹（可选）
   
   appendonly yes #redis持久化（可选）
   
   logfile "access.log"
   
   requirepass 123456(设置成你自己的密码)
   ```
   
   ```bash
   docker run -p 6379:6379 --name redis --restart always --network app-tier -v /mnt/dockerspace/redis/config/redis.conf:/etc/redis/redis.conf -v /mnt/dockerspace/redis/space/:/data:rw --privileged=true -d redis redis-server /etc/redis/redis.conf --appendonly yes
   ```
   
   ![docker-redis.png](docs%2Fimgs%2Fdocker-redis.png)

4. **安装etcd**
   
   ```bash
   docker pull bitnami/etcd:latest
   docker run -d --name etcd-server --network app-tier -p 2379:2379 -p 2380:2380 --env ALLOW_NONE_AUTHENTICATION=yes --env ETCD_ADVERTISE_CLIENT_URLS=http://[客户端访问etcd的ip]:2379 bitnami/etcd:latest
   docker ps -a
   ```
   
   测试：
   
   ```bash
   docker exec -it etcd-server bash
   
   etcdctl put /message Hello
   etcdctl get /message
   etcdctl del /message
   ```
   
   ![etcd-test.png](docs%2Fimgs%2Fetcd-test.png)
   更多信息 refer to: [传送门~](https://blog.csdn.net/weixin_42562106/article/details/122355562)

### 3.3 项目代码介绍

项目的文件目录布局如下：

```bash
project-root/
├─ config/                      // 全局配置文件
├─ dist/                        // 发布文件
├─ dockerfiles/                 // docker部署所需文件
├─ hertz/ 
│   ├─ hertz-gen/               // hertz代码生成 （hertz-router构建根目录）
│   │   ├─ biz/
│   │   │   ├─ handler/         // 1. http handler
│   │   │   └── .../
│   │   ├─ client/              // 2. rpc client
│   │   ├─ .../
│   │   ├─ main.go              // 3. 入口程序
│   │   ├─ router.go
│   │   └─ router_gen.go
│   └─ idl/                     // hertz代码生成所定义idl文件
│
├─ internal/
│   ├─ dal/                     // 数据库相关代码
│   ├─ mircoservice/
│   │   ├─ user                 // user-server相关代码构建目录
│   │   │   ├─ kitex-gen/user   // kitex生成代码，请勿修改 
│   │   │   ├─ handler.go       // 4. 用户服务请求处理函数
│   │   │   ├─ main.go          // 5. 微服务入口程序
│   │   │   └─ ...              // 构建脚手架等
│   │   │
│   │   ├─ video
│   │   │   └─ ...
│   │   │
│   │   ├─ interaction
│   │   │   └─ ...
│   │   │
│   │   └─ ...
│   │
│   └─ oss/
│
├─ kitex/
│   └─ idl/thriftgo          // kitex代码生成定义的ldl文件
│              
├─ middleware/               // hertz中间件和微服务中间件
│   ├─ Accesslog.go
│   ├─ Authorization.go
│   ├─ Common.go
│   ├─ Server.go
│   └─ init.go
│
├─ scripts/                  // 构建或部署的脚本命令文件
│
├─ utils/                    // 全局共用的工具类
│   ├─ config/
│   ├─ etcd/
│   ├─ jwt/
│   ├─ zap/
│   └─ xxx.go  
│ 
├─ logs/                      // 日志转储 
│   └─ ... 
│                
├─ .gitignore            
├─ go.mod
├─ go.sum
└─ README.md                 // 项目说明文档（this）
```

# 四、测试结果

### 4.1 **功能测试**

1. ##### **注册功能：**
   
   |          |                                              |
   | -------- | -------------------------------------------- |
   | **POST** | **123.57.251.188:8888/douyin/user/register** |
   | username | wangty0015                                   |
   | password | qwerasdf                                     |
   
   ![1.user_register.png](docs%2Fimgs%2Ftests%2F1.user_register.png)

2. ##### **登录功能：**
   
   | **POST** | **123.57.251.188:8888/douyin/user/login** |
   | -------- | ----------------------------------------- |
   | username | wangty0015                                |
   | password | qwerasdf                                  |
   
   ![2.user_login.png](docs%2Fimgs%2Ftests%2F2.user_login.png)

3. #### **用户信息**
   
   | **GET** | **123.57.251.188:8888/douyin/user/**                                                                                                                  |
   | ------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
   | user_id | 2883017410                                                                                                                                            |
   | token   | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6Mjg4MzAxNzQxMCwiZXhwIjoxNjk0MDMwNzM0LCJpc3MiOiJkb3VzaGVuZyJ9.Z8l7IrVIGqSpYZnNg-Di6c5RS4jQrF_5cnYTT2AqIOA |
   
   ![3.user_info.png](docs%2Fimgs%2Ftests%2F3.user_info.png)

4. ##### **视频流**
   
   |      |                                     |
   | ---- | ----------------------------------- |
   | GET  | **123.57.251.188:8888/douyin/feed** |
   | <br> | <br>                                |
   
   ![4.douyin_feed-1.png](docs%2Fimgs%2Ftests%2F4.douyin_feed-1.png)
   
   再次拉取视频流
   
   |             |                                     |
   | ----------- | ----------------------------------- |
   | GET         | **123.57.251.188:8888/douyin/feed** |
   | latest_time | 1693769373000                       |
   
   ![4.douyin_feed-2.png](docs%2Fimgs%2Ftests%2F4.douyin_feed-2.png)

5. ##### 视频投稿
   
   | **POST** | **123.57.251.188:8888/douyin/publish/action/**                                                                                                        |
   | -------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
   | token    | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6Mjg4MzAxNzQxMCwiZXhwIjoxNjk0MDMwNzM0LCJpc3MiOiJkb3VzaGVuZyJ9.Z8l7IrVIGqSpYZnNg-Di6c5RS4jQrF_5cnYTT2AqIOA |
   | data     | 【postman使用form-data选择file类型和文件路径】                                                                                                                     |
   | title    | test_upload                                                                                                                                           |
   
   ![5.publish_action.png](docs%2Fimgs%2Ftests%2F5.publish_action.png)

6. ##### 发布列表
   
   |         |                                                                                                                                                       |
   | ------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
   | **GET** | **123.57.251.188:8888/douyin/publish/list/**                                                                                                          |
   | user_id | 2883017410                                                                                                                                            |
   | tokrn   | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6Mjg4MzAxNzQxMCwiZXhwIjoxNjk0MDMwNzM0LCJpc3MiOiJkb3VzaGVuZyJ9.Z8l7IrVIGqSpYZnNg-Di6c5RS4jQrF_5cnYTT2AqIOA |
   
   ![6.publish_list.png](docs%2Fimgs%2Ftests%2F6.publish_list.png)

7. ##### 点赞操作
   
   | **POST**    | **123.57.251.188:8888/douyin/favorite/action/**                                                                                                       |
   | ----------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
   | token       | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6Mjg4MzAxNzQxMCwiZXhwIjoxNjk0MDMwNzM0LCJpc3MiOiJkb3VzaGVuZyJ9.Z8l7IrVIGqSpYZnNg-Di6c5RS4jQrF_5cnYTT2AqIOA |
   | video_id    | 100042                                                                                                                                                |
   | action_type | 1 // 点赞                                                                                                                                               |
   
   ![7.favorite_action-1.png](docs%2Fimgs%2Ftests%2F7.favorite_action-1.png)
   
   **重复点赞提示：**
   
   ![7.favorite_action-2.png](docs%2Fimgs%2Ftests%2F7.favorite_action-2.png)
   
   **取消点赞：**
   
   | **POST**    | **123.57.251.188:8888/douyin/favorite/action/**                                                                                                       |
   | ----------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
   | token       | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6Mjg4MzAxNzQxMCwiZXhwIjoxNjk0MDMwNzM0LCJpc3MiOiJkb3VzaGVuZyJ9.Z8l7IrVIGqSpYZnNg-Di6c5RS4jQrF_5cnYTT2AqIOA |
   | video_id    | 100042                                                                                                                                                |
   | action_type | 2 // 取消点赞                                                                                                                                             |
   
   ![7.favorite_action-3.png](docs%2Fimgs%2Ftests%2F7.favorite_action-3.png)
   
   **重复提交取消点赞**
   
   ![7.favorite_action-4.png](docs%2Fimgs%2Ftests%2F7.favorite_action-4.png)

8. ##### 喜欢列表
   
   |                   |                                                                                                                                                       |
   | ----------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
   | GET               | **123.57.251.188:8888/douyin/favorite/list/**                                                                                                         |
   | user_id           | 2883017410                                                                                                                                            |
   | token<br><br><br> | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6Mjg4MzAxNzQxMCwiZXhwIjoxNjk0MDMwNzM0LCJpc3MiOiJkb3VzaGVuZyJ9.Z8l7IrVIGqSpYZnNg-Di6c5RS4jQrF_5cnYTT2AqIOA |
   
   ![8.favorite_list.png](docs%2Fimgs%2Ftests%2F8.favorite_list.png)

9. ##### 评论操作
   
   | **POST**     | **123.57.251.188:8888/douyin/comment/action/**                                                                                                        |
   | ------------ | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
   | token        | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6Mjg4MzAxNzQxMCwiZXhwIjoxNjk0MDMwNzM0LCJpc3MiOiJkb3VzaGVuZyJ9.Z8l7IrVIGqSpYZnNg-Di6c5RS4jQrF_5cnYTT2AqIOA |
   | video_id     | 100042                                                                                                                                                |
   | action_type  | 1 // 评论                                                                                                                                               |
   | comment_text | test_comment_11111111                                                                                                                                 |
   
   ![9.comment_action-1.png](docs%2Fimgs%2Ftests%2F9.comment_action-1.png)
   
   **删除评论**
   
   | **POST**    | **123.57.251.188:8888/douyin/comment/action/**                                                                                                        |
   | ----------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
   | token       | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6Mjg4MzAxNzQxMCwiZXhwIjoxNjk0MDMwNzM0LCJpc3MiOiJkb3VzaGVuZyJ9.Z8l7IrVIGqSpYZnNg-Di6c5RS4jQrF_5cnYTT2AqIOA |
   | video_id    | 100042                                                                                                                                                |
   | action_type | 2 // 点赞                                                                                                                                               |
   | comment_id  | 4580610258                                                                                                                                            |
   
   ![9.comment_action-2.png](docs%2Fimgs%2Ftests%2F9.comment_action-2.png)

10. ##### 评论列表
    
    | **POST** | **123.57.251.188:8888/douyin/comment/list/**                                                                                                          |
    | -------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
    | token    | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6Mjg4MzAxNzQxMCwiZXhwIjoxNjk0MDMwNzM0LCJpc3MiOiJkb3VzaGVuZyJ9.Z8l7IrVIGqSpYZnNg-Di6c5RS4jQrF_5cnYTT2AqIOA |
    | video_id | 100042                                                                                                                                                |
    
    ![10.comment_list.png](docs%2Fimgs%2Ftests%2F10.comment_list.png)

11. ##### 服务解耦性测试
    
    > hertz-router（http路由网关服务）：✅打开
    > 
    > video-server（视频相关业务微服务）：✅打开
    > 
    > interaction-server（互动相关业务微服务）：✅打开
    > 
    > user-server（用户相关业务微服务）：❌关闭
    
    **1.11.1 测试用户信息接口（/douyin/user/）**
    
    | **POST**      | **123.57.251.188:8888/douyin/user/**                                                                                                                  |
    | ------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
    | token<br><br> | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6Mjg4MzAxNzQxMCwiZXhwIjoxNjk0MjQyODI4LCJpc3MiOiJkb3VzaGVuZyJ9.Lsi9LCBo7A63ehYZWszCoZ7bAA18ZHjFRtLXfGmiC6M |
    | user_id       | 2883017410                                                                                                                                            |
    
    ![11.service_discovery-1-user.png](docs%2Fimgs%2Ftests%2F11.service_discovery-1-user.png)
    
    | **参数**      | 预期结果               | 实际结果                                                                         | Status           |
    | ----------- | ------------------ | ---------------------------------------------------------------------------- | ---------------- |
    | status_code | -1                 | -1                                                                           | PASS             |
    | status_msg  | 微服务连接失败提示信息        | Service discovery error: no instance remains for Dousheng_Backend_UserServer | PASS<br><br><br> |
    | user        | [user null Object] | [user null Object]                                                           | PASS             |
    
    **1.11.2 测试用户信息接口（/douyin/publish/list）**
    
    | **POST** | **123.57.251.188:8888/douyin/publish/list/**                                                                                                          |
    | -------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
    | token    | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MzM4NDU4Njk5MywiZXhwIjoxNjk0MjQzNjUyLCJpc3MiOiJkb3VzaGVuZyJ9.f6BfRQocOxa3Sph4wSTsWdu5hb7jGVMYGvLXCPTdgUk |
    | user_id  | 3384586993                                                                                                                                            |
    
    ![11.service_discovery-2-video.png](docs%2Fimgs%2Ftests%2F11.service_discovery-2-video.png)
    
    | **参数**      | 预期结果               | 实际结果               | Status |
    | ----------- | ------------------ | ------------------ | ------ |
    | status_code | 0                  | 0                  | PASS   |
    | status_msg  | ok                 | ok                 | PASS   |
    | video_list  | [user Object, ...] | [user Object, ...] | PASS   |

### 4.2 **性能测试**

```bash
go tool pprof -http=":8085" http://123.57.251.188:8888/dev/pprof/profile
```

# 五、Demo 演示视频 （必填）

## 5.1 POSTMAN接口测试演示

[postman_api_test.mp4](docs%2Fvideos%2Fpostman_api_test.mp4)

[postman_api_test.mp4__在线播放_1](https://www.bilibili.com/video/BV13w411U7dY/?vd_source=7c8ec9c1b876283e90058130ff5aa0a6)

## 5.2 手机客户端测试

[android_api_test.mp4](docs%2Fvideos%2Fandroid_api_test.mp4)

[android_api_test.mp4__在线播放_2](https://www.bilibili.com/video/BV1yj41117s9/)

# 六、项目总结与反思

### 6.1 项目中已存在的问题

目前项目中仍有待完善的地方和解决思路主要如下：

1. api接口的附加可选功能实现：后续可以增加关系和消息相关的微服务开发。本项目基于微服务架构开发，所以具有很好的扩展性，后续功能可以通过迭代，在几乎不改变主体的情况下，继续完善内容。

2. 性能测试尚未完善：现阶段已经引入pprof测试工具组件，但是对于具体的测试结果仍需要进一步工作。

### 6.2 已识别出的优化项

        oss云部署受限于公网流量额度限制，采用nginx代理的模式，可以免除流量费用，但是下载速度1Mbps较低，速度和用户体验有待提升。

        采用分布式集群进行部署和复杂均衡，压力测试等。

### 6.3 架构演进的可能性（★）

        本项目采用微服务架构的方式进行开发，目前是通过应用程序部署在docker的同一网络中实现**环境隔离**和**网络访问**。后续变更主要有以下几个方向：

1. **服务器集群，docker分布式部署**，**负载均衡**实现高并发业务场景：项目可以根据需求的变更，或用户规模扩大等，进一步提升微服务的数量规模，使用服务器集群的方式，来适应业务增长可能遇到的高并发，高吞吐场景。对于容器化和编排，使用容器（如Docker）和容器编排平台（如Kubernetes）来管理和部署微服务，可以更轻松地扩展、升级和维护微服务。
2. **业务增长或安全性提高要求，拆分服务：** 比如要实现较高的安全认证，或者修改为单点登录等方式，则需要将Auth中间件抽离成（微）服务的形式。
3. **微服务之间的通讯需要：** 引入事件驱动的架构，允许微服务之间异步通信和松耦合等，引入熔断、限流等机制，提高系统的可伸缩性和灵活性。

### 6.4 项目过程中的反思与总结

        经过约一个月的努力，通过青训营和这次项目，不论是技术上还是眼界上，我都获得了很大的收获！这次微服务项目，从架构到开发到各种功能测试，再到最后的云部署等等，我熟练掌握了微服务架构项目的较全流程的开发，熟悉了微服务的开发过程，以及架构的原理和应用场景，并对可靠性、并发性、高性能、低耦合等，都有了很深刻的理解和项目实践。从最开始对go语言知之甚少，只懂得简单语法等；经历了项目中的各种难点，创新和技术的攻克，在不断超越自我的过程中，也更加深了对于后端开发这一工作的理解。

        虽然该项目中也还存在着一些可优化之处，但是对于之后的程序开发，都会有所启发并从中受益！

---

# Quick Start、docker compose 启动项目

> 1. 启动etcd-server，mysql，redis，oss-proxy等服务
> 
> 2. 创建网络并指定网段：`docker network create --subnet=192.168.2.0/24 dousheng_net`
> 
> 3. 分别在根目录逐行逐次运行`build_router`和`build_server`中的命令，如果是在windows下部署需要先在cmd中运行`build_prepare_to-linux-amd64.bat`，来实现windows下编译linux的效果
> 
> 4. 将dist文件夹，config文件夹，dockerfiles文件夹上传至服务器的项目目录中。
> 
> 5. 运行docker-compose.yml：<br>`docker-compose -f ./dockerfiles/docker-compose.yml up -d`
> 
> 6. 其他检查命令：<br>
>    ```docker network inspect xxx```<br>
>    ```docker network inspect dousheng_net```<br>
> 
>    ```docker network inspect bridge```<br>
>    ```docker ps -a```<br>
>    ```docker image ls```<br>
> 
> 7. 停止命令：<br>
>    `docker-compose -f ./dockerfiles/docker-compose.yml stop`
> 
> 8. 停止并删除命令：<br>
>    `docker-compose -f ./dockerfiles/docker-compose.yml down`
