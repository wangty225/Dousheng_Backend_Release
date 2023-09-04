// Code generated by Kitex v0.6.2. DO NOT EDIT.

package videoservice

import (
	video "Dousheng_Backend/internal/mircoservice/video/kitex-gen/video"
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return videoServiceServiceInfo
}

var videoServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "VideoService"
	handlerType := (*video.VideoService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Feed":          kitex.NewMethodInfo(feedHandler, newVideoServiceFeedArgs, newVideoServiceFeedResult, false),
		"ActionPublish": kitex.NewMethodInfo(actionPublishHandler, newVideoServiceActionPublishArgs, newVideoServiceActionPublishResult, false),
		"ListPublish":   kitex.NewMethodInfo(listPublishHandler, newVideoServiceListPublishArgs, newVideoServiceListPublishResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "video",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.2",
		Extra:           extra,
	}
	return svcInfo
}

func feedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceFeedArgs)
	realResult := result.(*video.VideoServiceFeedResult)
	success, err := handler.(video.VideoService).Feed(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceFeedArgs() interface{} {
	return video.NewVideoServiceFeedArgs()
}

func newVideoServiceFeedResult() interface{} {
	return video.NewVideoServiceFeedResult()
}

func actionPublishHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceActionPublishArgs)
	realResult := result.(*video.VideoServiceActionPublishResult)
	success, err := handler.(video.VideoService).ActionPublish(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceActionPublishArgs() interface{} {
	return video.NewVideoServiceActionPublishArgs()
}

func newVideoServiceActionPublishResult() interface{} {
	return video.NewVideoServiceActionPublishResult()
}

func listPublishHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceListPublishArgs)
	realResult := result.(*video.VideoServiceListPublishResult)
	success, err := handler.(video.VideoService).ListPublish(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceListPublishArgs() interface{} {
	return video.NewVideoServiceListPublishArgs()
}

func newVideoServiceListPublishResult() interface{} {
	return video.NewVideoServiceListPublishResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Feed(ctx context.Context, req *video.DouyinFeedRequest) (r *video.DouyinFeedResponse, err error) {
	var _args video.VideoServiceFeedArgs
	_args.Req = req
	var _result video.VideoServiceFeedResult
	if err = p.c.Call(ctx, "Feed", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ActionPublish(ctx context.Context, req *video.DouyinPublishActionRequest) (r *video.DouyinPublishActionResponse, err error) {
	var _args video.VideoServiceActionPublishArgs
	_args.Req = req
	var _result video.VideoServiceActionPublishResult
	if err = p.c.Call(ctx, "ActionPublish", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListPublish(ctx context.Context, req *video.DouyinPublishListRequest) (r *video.DouyinPublishListResponse, err error) {
	var _args video.VideoServiceListPublishArgs
	_args.Req = req
	var _result video.VideoServiceListPublishResult
	if err = p.c.Call(ctx, "ListPublish", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
