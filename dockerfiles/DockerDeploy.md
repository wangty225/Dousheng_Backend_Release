### 1. 本地编写Dockfile，并上传
### 2. 编写build命令并服务器上进行构建镜像
### 3. 编写run命令并在服务器上新建容器
### 4. 运行命令查看IP: `docker inspect dousheng_xxx-server | grep IPAddress`
### 5. `docker exec -it dousheng_xxx-server sh`进入服务器容器修改etcd注册地址；`vi config/etcd/xxx.yml`：修改`server.host`为`172.17.x.x`
### 6. ~~`docker exec -it doushengrouter sh` 进入路由端修改etcd发现地址~~，
### 保证`xxx.yml`中的`server.name`名称与实际`mircoservice`在`etcd`中注册名称一致即可
