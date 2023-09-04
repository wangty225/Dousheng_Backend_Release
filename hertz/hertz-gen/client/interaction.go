package client

import (
	"Dousheng_Backend/hertz/hertz-gen/biz/model/interaction_hertz"
	"Dousheng_Backend/internal/mircoservice/interaction/kitex-gen/interaction"
	"Dousheng_Backend/internal/mircoservice/interaction/kitex-gen/interaction/interactionservice"
	"Dousheng_Backend/utils/etcd"
	"Dousheng_Backend/utils/zap"
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/spf13/viper"
	"path/filepath"
	"time"
)

var (
	interactionClient interactionservice.Client
)

func InitInteraction(config *viper.Viper) {
	etcdAddr := fmt.Sprintf("%s:%d", config.GetString("etcd.host"), config.GetInt("etcd.port"))
	serviceName := config.GetString("server.name")

	r, err := etcd.NewEtcdResolver([]string{etcdAddr})
	if err != nil {
		panic(err)
	}

	c, err := interactionservice.NewClient(
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
	interactionClient = c
}

func init() {
	logger := zap.InitDefaultLogger()
	interactionConfig := viper.New()
	interactionConfig.SetConfigType("yml")                              //设置配置文件类型
	interactionConfig.SetConfigName("interaction")                      //设置配置文件名
	interactionConfig.AddConfigPath(filepath.ToSlash("./config/etcd/")) //设置配置文件路径
	if err := interactionConfig.ReadInConfig(); err != nil {
		logger.Errorf("errno is %+v\n", err)
		return
	}
	InitInteraction(interactionConfig)
}

func FavoriteAction(ctx context.Context, req *interaction_hertz.DouyinFavoriteActionRequest) (*interaction.DouyinFavoriteActionResponse, error) {
	kitexReq := &interaction.DouyinFavoriteActionRequest{
		Token:      req.Token,
		VideoId:    req.VideoID,
		ActionType: req.ActionType,
	}
	return interactionClient.ActionFavorite(ctx, kitexReq)
}

func FavoriteList(ctx context.Context, req *interaction_hertz.DouyinFavoriteListRequest) (*interaction.DouyinFavoriteListResponse, error) {
	kitexReq := &interaction.DouyinFavoriteListRequest{
		UserId: req.GetUserID(),
		Token:  req.GetToken(),
	}
	return interactionClient.ListFavorite(ctx, kitexReq)
}

func CommentAction(ctx context.Context, req *interaction_hertz.DouyinCommentActionRequest) (*interaction.DouyinCommentActionResponse, error) {
	kitexReq := &interaction.DouyinCommentActionRequest{
		Token:       req.Token,
		VideoId:     req.VideoID,
		ActionType:  req.ActionType,
		CommentText: req.CommentText,
		CommentId:   req.CommentID,
	}
	return interactionClient.ActionComment(ctx, kitexReq)
}

func CommentList(ctx context.Context, req *interaction_hertz.DouyinCommentListRequest) (*interaction.DouyinCommentListResponse, error) {
	kitexReq := &interaction.DouyinCommentListRequest{
		Token:   req.Token,
		VideoId: req.VideoID,
	}
	return interactionClient.ListComment(ctx, kitexReq)
}
