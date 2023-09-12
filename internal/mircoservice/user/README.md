## 一. 环境测试

1. 项目里的kitex版本 v0.5.0， 其他版本windows内存泄漏（不知道什么BUG）。
2. 大家可以打开terminal，`cd microservice/user`，依次进行如下：
3. `go mod tidy`
4. `go build -o dist/test_server.exe ./`
5. 如果成功生成dist/test_server.exe，环境应该就没有问题了。
   ![test_server.png](docs%2Fimgs%2Ftest_server.png)
   双击test_server.exe，如果出现如下图片所示，即微服务开启成功。可以点击user_client.exe测试rpc通信。
   ![test_rpc.png](docs%2Fimgs%2Ftest_rpc.png)

## 二. NOTICE！

1. 微服务之间如果需要通信，使用rpc通信而不是本地函数调用。rpc通信需要参考microservice/user/client/main.go的方法实现。
2. 微服务最终目的是保证每个微服务都可以单独部署，访问数据库公用一个dal下的数据库连接。
3. 使用User结构体时，注意区别UserDao结构体和kitex-gen生成的User结构体，一个是数据存储obj一个是req和resp的obj。
4. 编写自己的微服务时，使用Kitex-gen生成的结构体要使用自己的kitex-gen目录下的结构体，goland导入时要注意别导入错误。
   如microservice/video/handler.go使用user时，import的是video/kitex-gen/user，而不是user/kitex-gen/user
   ![img.png](docs%2Fimgs%2Fimg.png)
5. 数据库已经定义完毕，具体的服务和逻辑的分层自由实现。
6. Hertz框架后续我会继续完善。
7. **总体代码框架类似：**

![structure.png](https://github.com/bytedance-youthcamp-jbzx/tiktok/blob/main/pic/%E6%8A%96%E5%A3%B0_%E5%90%8E%E7%AB%AF%E6%9E%B6%E6%9E%84%E5%9B%BE.png)

## 三. Database

### 1. 调用mysql数据库方法

> 直接使用`DBMysql.xxx()`即可。
> 
> 具体的服务和逻辑的分层自由实现。

`dal/mysql/init_test.go`

```go
func QuesyOne() {
   var u UserDao
   fmt.Printf("%s\n", u.String())
   ret := DBMysql.First(&u, "username=?", "wangty")
   if ret.Error != nil {
      fmt.Println("Error querying database:", ret.Error)
      return
   }
   if ret.RowsAffected == 0 {
      fmt.Println("No matching records found.")
      return
   }
   fmt.Printf("%s\n", u.String())
   fmt.Printf("%v\n", u.ID)
   fmt.Printf("%s\n", u.BackgroundImage)
}
```

### 2. 调用redis数据库方法：

> 直接使用`DBRedis.xxx()`即可。

`dal/redis/init_test.go`

```go
func Connect() {
   pong, err := DBRedis.Ping().Result()
      if err != nil {
      fmt.Println(err)
   }
   fmt.Println(pong)
}
```

### 3. 调用oss方式

> 配置文件/Dousheng_Backend/config/yml/aliyun_oss.yml已经给出

> 调用方式参考文档
> [Go语言为例,介绍服务端签名直传并设置上传回调_对象存储 OSS-阿里云帮助中心 (aliyun.com)](https://help.aliyun.com/zh/oss/use-cases/go-1?spm=a2c4g.11186623.0.0.44e84211yaAECY)

> For more to see in `internal/oss/`
> 
