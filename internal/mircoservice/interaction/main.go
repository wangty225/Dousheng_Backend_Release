package main

import (
	"Dousheng_Backend/internal/mircoservice/interaction/kitex-gen/interaction/interactionservice"
	"Dousheng_Backend/middleware"
	"Dousheng_Backend/utils/config"
	"Dousheng_Backend/utils/etcd"
	"Dousheng_Backend/utils/jwt"
	"Dousheng_Backend/utils/zap"
	"fmt"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"net"
)

var Jwt = jwt.NewJWT()

var (
	etcdConfig  = config.InitConfig("./config/etcd/interaction.yml")
	serviceName = etcdConfig.GetString("server.name")
	serviceAddr = fmt.Sprintf("%s:%d", etcdConfig.GetString("server.host"), etcdConfig.GetInt("server.port"))
	etcdAddr    = fmt.Sprintf("%s:%d", etcdConfig.GetString("etcd.host"), etcdConfig.GetInt("etcd.port"))
	logger      = zap.InitDefaultLogger()
)

func main() {

	// 服务注册
	r, err := etcd.NewEtcdRegistry([]string{etcdAddr})
	if err != nil {
		logger.Fatalln("[interacton-server]", err.Error())
	}

	addr, err := net.ResolveTCPAddr("tcp", serviceAddr)
	if err != nil {
		logger.Fatalln("[interacton-server]", err.Error())
	}
	fmt.Println("[interacton-server]init interaction server ...")

	// 初始化etcd
	s := interactionservice.NewServer(new(InteractionServiceImpl),
		server.WithServiceAddr(addr),
		server.WithMiddleware(middleware.CommonMiddleware),
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
	)

	fmt.Println("[interacton-server]run interaction server ...")
	if err := s.Run(); err != nil {
		logger.Fatalf("[interacton-server]%v stopped with error: %v", serviceName, err.Error())
	}

}
