package client

import (
	"Dousheng_Backend/hertz/hertz-gen/biz/model/user_hertz"
	"Dousheng_Backend/internal/mircoservice/user/kitex-gen/user"
	"Dousheng_Backend/internal/mircoservice/user/kitex-gen/user/userservice"
	"Dousheng_Backend/utils/etcd"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/spf13/viper"
	"log"
	"path/filepath"
	"time"
)

import (
	"context"
)

var (
	userClient userservice.Client
)

func InitUser(config *viper.Viper) {
	etcdAddr := fmt.Sprintf("%s:%d", config.GetString("etcd.host"), config.GetInt("etcd.port"))
	serviceName := config.GetString("server.name")

	r, err := etcd.NewEtcdResolver([]string{etcdAddr})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		serviceName,
		//client.WithMiddleware(middleware.CommonMiddleware),
		//client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(30*time.Second),             // rpc timeout
		client.WithConnectTimeout(30000*time.Millisecond), // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(tracing.NewClientSuite()),        // tracer
		client.WithResolver(r),                            // resolver
		// Please keep the same as provider.WithServiceName
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

func init() {
	userConfig := viper.New()
	userConfig.SetConfigType("yml")  //设置配置文件类型
	userConfig.SetConfigName("user") //设置配置文件名
	//userConfig.AddConfigPath(filepath.ToSlash(utils.GetCallerAbsPath() + "/../../../config/etcd/")) //设置配置文件路径
	userConfig.AddConfigPath(filepath.ToSlash("./config/etcd/")) //设置配置文件路径
	if err := userConfig.ReadInConfig(); err != nil {
		//global.SugarLogger.Fatalf("read config files failed,errors is %+v", err)
		log.Fatalf("errno is %+v", err)
	}
	InitUser(userConfig)
}

func RegisterUser(ctx context.Context, req *user_hertz.DouyinUserRegisterRequest) (*user.DouyinUserRegisterResponse, error) {
	// // deleted
	// convert UserRegisterRequest to DouyinUserRegisterRequest (BUG：import err)
	//douyinReq := user2.DouyinUserRegisterRequest{
	//	Username: req.Username,
	//	Password: req.Password,
	//}
	//douyinResp, err := userClient.RegisterUser(ctx, &douyinReq)
	//if err != nil {
	//	log.Fatalf("hertz -> rpc : Register请求失败")
	//}
	//
	//resp := user.DouyinUserRegisterResponse{
	//	StatusCode: douyinResp.StatusCode,
	//	StatusMsg:  douyinResp.GetStatusMsg(),
	//	UserId:     douyinResp.UserId,
	//	Token:      douyinResp.Token,
	//}
	//return &resp, nil

	douyinReq := &user.DouyinUserRegisterRequest{
		Username: req.Username,
		Password: req.Password,
	}
	return userClient.RegisterUser(ctx, douyinReq)
}

func LoginUser(ctx context.Context, req *user_hertz.DouyinUserLoginRequest) (*user.DouyinUserLoginResponse, error) {
	// convert UserLoginRequest to DouyinUserLoginRequest
	douyinReq := &user.DouyinUserLoginRequest{
		Username: req.Username,
		Password: req.Password,
	}
	return userClient.LoginUser(ctx, douyinReq)
}

func User(ctx context.Context, req *user_hertz.DouyinUserRequest) (*user.DouyinUserResponse, error) {
	// convert UserRequest to DouyinUserRequest
	douyinReq := &user.DouyinUserRequest{
		UserId: req.UserID,
		Token:  req.Token,
	}
	return userClient.User(ctx, douyinReq)
}
