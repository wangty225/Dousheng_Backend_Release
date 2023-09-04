package client

import (
	"Dousheng_Backend/hertz/hertz-gen/biz/model/video_hertz"
	"Dousheng_Backend/internal/mircoservice/video/kitex-gen/video"
	"Dousheng_Backend/internal/mircoservice/video/kitex-gen/video/videoservice"
	"Dousheng_Backend/utils/etcd"
	"context"
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

var (
	videoClient videoservice.Client
)

func init() {
	videoConfig := viper.New()
	videoConfig.SetConfigType("yml")                              //设置配置文件类型
	videoConfig.SetConfigName("video")                            //设置配置文件名
	videoConfig.AddConfigPath(filepath.ToSlash("./config/etcd/")) //设置配置文件路径
	if err := videoConfig.ReadInConfig(); err != nil {
		log.Fatalf("errno is %+v", err)
		return
	}
	InitVideo(videoConfig)
}

func InitVideo(config *viper.Viper) {
	etcdAddr := fmt.Sprintf("%s:%d", config.GetString("etcd.host"), config.GetInt("etcd.port"))
	serviceName := config.GetString("server.name")

	r, err := etcd.NewEtcdResolver([]string{etcdAddr})
	if err != nil {
		panic(err)
	}

	c, err := videoservice.NewClient(
		serviceName,
		//client.WithMiddleware(middleware.CommonMiddleware),
		//client.WithInstanceMW(middleware.ClientMiddleware),

		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(90*time.Second),             // rpc timeout
		client.WithConnectTimeout(90000*time.Millisecond), // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(tracing.NewClientSuite()),        // tracer
		client.WithResolver(r),                            // resolver
		// Please keep the same as provider.WithServiceName
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
	)
	if err != nil {
		panic(err)
	}
	videoClient = c
}

func VideoFeed(ctx context.Context, req *video_hertz.DouyinFeedRequest) (*video.DouyinFeedResponse, error) {
	kitexReq := &video.DouyinFeedRequest{
		LatestTime: req.LatestTime,
		Token:      req.Token,
	}
	return videoClient.Feed(ctx, kitexReq)
	//
	//kitexResp, err := videoClient.Feed(ctx, kitexReq)
	//if err != nil {
	//	fmt.Printf("kitexResp err in file2: %+v\n", kitexResp) // todo
	//	fmt.Printf("kitexResp err in file2 err: %+v\n", err)   // todo
	//	return video_hertz.NewDouyinFeedResponse(), err
	//}
	//hertzVideoList := make([]*video_hertz.Video, 0)
	//for _, v := range kitexResp.VideoList {
	//	hertzVideoList = append(hertzVideoList, &video_hertz.Video{
	//		ID: v.Id,
	//		Author: &user_hertz.User{
	//			ID:              v.Author.Id,
	//			Name:            v.Author.Name,
	//			FollowCount:     v.Author.FollowCount,
	//			FollowerCount:   v.Author.FollowerCount,
	//			IsFollow:        v.Author.IsFollow,
	//			Avator:          v.Author.Avator,
	//			BackgroundImage: v.Author.BackgroundImage,
	//			Signature:       v.Author.Signature,
	//			TotalFavorited:  v.Author.TotalFavorited,
	//			WorkCount:       v.Author.WorkCount,
	//			FavoriteCount:   v.Author.FavoriteCount,
	//		},
	//		PlayURL:       v.PlayUrl,
	//		CoverURL:      v.CoverUrl,
	//		FavoriteCount: v.FavoriteCount,
	//		CommentCount:  v.CommentCount,
	//		IsFavorite:    v.IsFavorite,
	//		Title:         v.Title,
	//	})
	//}
	//response := &video_hertz.DouyinFeedResponse{
	//	StatusCode: kitexResp.StatusCode,
	//	StatusMsg:  kitexResp.StatusMsg,
	//	VideoList:  hertzVideoList,
	//	NextTime:   kitexResp.NextTime,
	//}
	//return response, nil
}

func PublishAction(ctx context.Context, req *video_hertz.DouyinPublishActionRequest) (*video.DouyinPublishActionResponse, error) {
	kitexReq := &video.DouyinPublishActionRequest{
		Token: req.Token,
		Data:  req.Data,
		Title: req.Title,
	}
	return videoClient.ActionPublish(ctx, kitexReq)
}

func PublishList(ctx context.Context, req *video_hertz.DouyinPublishListRequest) (*video.DouyinPublishListResponse, error) {
	kitexReq := &video.DouyinPublishListRequest{
		UserId: req.UserID,
		Token:  req.Token,
	}
	return videoClient.ListPublish(ctx, kitexReq)
}
