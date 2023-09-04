# Dousheng_Backend

### 一、微服务接口划分

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

划分为五个微服务，其中user_service和video_service为基础功能，interaction_service和relation_service和message_service为自选功能。



### 二、项目结构目录

```bash
project-root/
├─ config/                      // 全局配置文件
├─ dist/                        // 发布文件
├─ dockerfiles/                 // docker部署所需文件
├─ hertz/ 
│   ├─ hertz-gen/               // hertz代码生成 （hertz-router构建根目录）
│   │   ├─ biz/
│   │   │   ├─ handler/         // ①http handler
│   │   │   └── .../
│   │   ├─ client/              // ②rpc client
│   │   ├─ .../
│   │   ├─ main.go              // ③入口程序
│   │   ├─ router.go
│   │   └─ router_gen.go
│   └─ idl/                     // hertz代码生成所定义idl文件
│
├─ internal/
│   ├─ dal/                     // 数据库相关代码
│   ├─ mircoservice/
│   │   ├─ user                 // user-server相关代码构建目录
│   │   │   ├─ kitex-gen/user   // kitex生成代码，请勿修改 
│   │   │   ├─ handler.go       // ④用户服务请求处理函数
│   │   │   ├─ main.go          // ⑤微服务入口程序
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



### 三、项目环境部署流程

 **0. 安装docker环境**

        i. 安装当然是要看官方说明文档啦！

          [ubuntu安装docker](https://docs.docker.com/engine/install/ubuntu/)

        ii. 创建自定义docker网络，方便后续管理容器的ip地址

           `docker network create --subnet=192.168.1.0/24 app-tier`

**1. 安装mysql 5.7**

> ```bash
> docker run -d -p 3306:3306 --name mysql --network app-tier -e MYSQL_ROOT_PASSWORD=dousheng_WSX mysql:5.7
> ```

> ```bash
> docker exec -it mysql bash
> mysql -u mysql -p
> ```
> 
> ```sql
> create database dousheng_db;
> ```
> 
> ```SQL
> # do some sqls
> ```
![ER-mysql.PNG](docs%2Fimgs%2FER-mysql.PNG)
mysql ER图

**2. 安装redis**

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

```config
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


**3. 安装etcd**

```bash
docker pull bitnami/etcd:latest
docker run -d --name etcd-server --network app-tier -p 2379:2379 -p 2380:2380 --env ALLOW_NONE_AUTHENTICATION=yes --env ETCD_ADVERTISE_CLIENT_URLS=http://[客户端访问etcd的ip]:2379 bitnami/etcd:latest
docker ps -a
```

测试：

![etcd-test.png](docs%2Fimgs%2Fetcd-test.png)

```bash
docker exec -it etcd-server bash

etcdctl put /message Hello
etcdctl get /message
etcdctl del /message
```

更多信息 refer to: [传送门~](https://blog.csdn.net/weixin_42562106/article/details/122355562)

**4. nginx配置oss-proxy**

由于阿里云oss公网流量付费较贵，而内网流出流量免费，所以选择使用ecs上的nginx方向代理oss，这样既可以实现内网流量转发公网，又可以保证oss的安全性。

```bash
vim aliyun_oss.conf
```

```yml
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

其他部署自行google

### 四、docker compose 启动项目

> 1. 启动etcd-server，mysql，redis，oss-proxy等服务
> 2. 创建网络并指定网段：`docker network create --subnet=192.168.2.0/24 dousheng_net`
> 3. 分别在根目录逐行逐次运行`build_router`和`build_server`中的命令，如果是在windows下部署需要先在cmd中运行`build_prepare_to-linux-amd64.bat`，来实现windows下编译linux的效果
> 4. 将dist文件夹，config文件夹，dockerfiles文件夹上传至服务器的项目目录中。
> 5. 运行docker-compose.yml：`docker-compose -f ./dockerfiles/docker-compose.yml up -d`
>    
>    
> 6. 其他检查命令：<br>
>    ```docker network inspect```<br>
>    ```docker network inspect dousheng_net```<br>
>    ```docker network inspect bridge```<br>
>    ```docker ps -a```<br>
>    ```docker image ls```<br>
> 7. 停止命令：<br>
>    `docker-compose -f ./dockerfiles/docker-compose.yml stop`
> 8. 停止并删除命令：<br>
>    `docker-compose -f ./dockerfiles/docker-compose.yml down`