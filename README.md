# Dousheng_Backend

### 微服务开发参考文档：

[README.md](internal%2Fmircoservice%2Fuser%2FREADME.md)

### 项目结构暂定

go 1.20.7

```bash
# 微服务接口划分
user_service:
  ├─ /douyin/user/register/         # 用户注册接口   
  ├─ /douyin/user/login/            # 用户登录接口
  └─ /douyin/user/                  # 用户信息

video_service:
  ├─ /douyin/feed/                  # 视频流接口
  ├─ /douyin/publish/action/        # 视频投稿
  └─ /douyin/publish/list/          # 发布列表

interaction_service: 
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

```bash
project-root/
│
├─ config/               // 全局配置文件
├─ dal/
│   ├─ mysql/
│   │   ├─ init.go
│   │   ├─ user.go
│   │   ├─ video.go
│   │   ├─ ...    
│   │   └─ message.go
│   │   
│   └─ redis/ 
│       ├─ init.go
│       ├─ user.go
│       ├─ video.go
│       ├─ ...    
│       └─ message.go
│    
├─ internal/microservice
│   ├─ user/
│   │   ├─ client/       // 微服务的访问方式
│   │   ├─ cmd/          // 微服务构建脚本
│   │   ├─ dist/         // 微服务构建执行文件（.exe）
│   │   ├─ docs/image    // 文本图片
│   │   ├─ kitex-gen/    // kitex生成代码（请勿修改）
│   │   ├─ script/       // kitex生成脚手架
│   │   ├─ handler.go    // 用户服务请求处理函数
│   │   ├─ main.go       // 微服务启动入口（server端）
│   │   ├─ go.mod        // 必选（微服务环境配置）
│   │   ├─ model/        // 预定义：用户服务数据模型定义（todo）
│   │   ├─ service/      // 预定义：用户服务业务逻辑（todo）
│   │   ├─ utils/        // 预定义：工具类（todo）
│   │   ├─ config/       // 预定义：服务配置文件（todo）
│   │   └─ README.md
│   │
│   ├─ video_service/
│   │   └─ ......
│   │
│   ├─ interaction_service/
│   │   └─ ......
│   │
│   ├─ relation_service/
│   │   └─ ......
│   │
│   ├─ message_service/
│   │   └─ ......
│   │
│   ├─ configs/          // 预定义：全局配置文件（todo）
│   ├─ middleware/       // 预定义：全局中间件（todo）
│   └─ utils/            // 预定义：全局工具函数（todo）
│
├─ kitex/                // kitex-idl文件
│   ├─ idl
│   │   └─ thriftgo
│   │        └─ ***.thrift
│   └─ kitex_auto.bat.DONOTRUN   
│                        //  kitex代码生成执行命令记录
│
├─ local/                // 局部配置文件（代码遗留，尚未重构，勿动即可）
├─ logs/                 // 日志文件目录（todo）
├─ libs/                 // 第三方工具和库（todo）
├─ middlewares/          // 中间件（todo）
├─ .gitignore            
├─ go.mod                
└─ README.md             // 项目说明文档（this）
```

### 阿里云：

[阿里云免费试用 - 阿里云 (aliyun.com)](https://free.aliyun.com/?crowd=personal&spm=a2c6h.23978828.J_5404914170.55.1818df38CDngZA)

- 阿里云进行oss对象存储。
  
  > 新用户赠送3个月20G的免费存储和2G下行流量额度。
  > 
  > [Go语言为例,介绍服务端签名直传并设置上传回调_对象存储 OSS-阿里云帮助中心 (aliyun.com)](https://help.aliyun.com/zh/oss/use-cases/go-1?spm=a2c4g.11186623.0.0.44e84211yaAECY)

- 阿里云RocketMQ消息队列（暂缓）。

- 阿里云ECS配置mysql (5.7)，微服务上线。
  
  1. **mysql创建**
  
  > docker 创建mysql容器
  > 
  > ```bash
  > docker run -d -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=dousheng_WSX mysql:5.7
  > ```
  > 
  > ```bash
  > docker exec -it mysql bash
  > mysql -u mysql -p
  > ```
  > 
  > ```sql
  > create database dousheng_db;
  > ```
  > 
  > // TODO
  > 
  > ```SQL
  > # some sqls
  > ```
2. **mysql访问：**
   
   > ip: 
   > 
   > port: 
   > 
   > login: 
   > 
   > passwd: 
   > 
   > dbname: 



### Notice：

> 1. 由于项目遗留的问题本项目有**全局配置文件**和**局部配置文件**。
> 
> 2. 全局配置文件位于①`/config/yml`中，局部配置文件位于②`/local/yml`中。
> 
> 3. 设计之初是把**局部配置文件**放在微服务中，来实现更好的解耦合，所以**局部配置文件** > **全局配置文件**。相同的配置，局部配置文件会覆盖全局配置文件。
> 
> 4. **如果需要修改配置，直接在局部配置文件中修改即可，或者①和②都修改。**
> 
> 5. 但是考虑到可能存在代码复用和降低开发难度，修改了项目的结构，所有微服务统一使用`dal`提供的本地函数接口。
> 
> 6. 现在代码运行正常，先不对`local`和`dal/xxx/init`做修改，这样也方便以后修改成解耦代码（如果需要且必要的话）。


### docker compose 启动服务
> 1. 启动etcd-server，mysql，redis，oss-proxy等服务
> 2. 创建网络并指定网段：`docker network create --subnet=192.168.2.0/24 dousheng_net`
> 3. 运行docker-compose.yml：`docker-compose -f ./dockerfiles/docker-compose.yml up -d`
> 4. 其他检查命令：<br>
```docker network inspect```<br>
```docker network inspect dousheng_net```<br>
```docker network inspect bridge```<br>
```docker ps -a```<br>
```docker image ls```<br>
> 5. 停止命令：<br>
`docker-compose -f ./dockerfiles/docker-compose.yml stop`
> 6. 停止并删除命令：<br>
`docker-compose -f ./dockerfiles/docker-compose.yml down`