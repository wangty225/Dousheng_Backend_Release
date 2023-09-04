// Code generated by Kitex v0.6.2. DO NOT EDIT.

package relationservice

import (
	relation "Dousheng_Backend/internal/mircoservice/relation/kitex-gen/relation"
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return relationServiceServiceInfo
}

var relationServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "RelationService"
	handlerType := (*relation.RelationService)(nil)
	methods := map[string]kitex.MethodInfo{
		"ActionRelation":       kitex.NewMethodInfo(actionRelationHandler, newRelationServiceActionRelationArgs, newRelationServiceActionRelationResult, false),
		"ListFollowRelation":   kitex.NewMethodInfo(listFollowRelationHandler, newRelationServiceListFollowRelationArgs, newRelationServiceListFollowRelationResult, false),
		"ListFollowerRelation": kitex.NewMethodInfo(listFollowerRelationHandler, newRelationServiceListFollowerRelationArgs, newRelationServiceListFollowerRelationResult, false),
		"ListFriendRelation":   kitex.NewMethodInfo(listFriendRelationHandler, newRelationServiceListFriendRelationArgs, newRelationServiceListFriendRelationResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "relation",
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

func actionRelationHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceActionRelationArgs)
	realResult := result.(*relation.RelationServiceActionRelationResult)
	success, err := handler.(relation.RelationService).ActionRelation(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceActionRelationArgs() interface{} {
	return relation.NewRelationServiceActionRelationArgs()
}

func newRelationServiceActionRelationResult() interface{} {
	return relation.NewRelationServiceActionRelationResult()
}

func listFollowRelationHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceListFollowRelationArgs)
	realResult := result.(*relation.RelationServiceListFollowRelationResult)
	success, err := handler.(relation.RelationService).ListFollowRelation(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceListFollowRelationArgs() interface{} {
	return relation.NewRelationServiceListFollowRelationArgs()
}

func newRelationServiceListFollowRelationResult() interface{} {
	return relation.NewRelationServiceListFollowRelationResult()
}

func listFollowerRelationHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceListFollowerRelationArgs)
	realResult := result.(*relation.RelationServiceListFollowerRelationResult)
	success, err := handler.(relation.RelationService).ListFollowerRelation(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceListFollowerRelationArgs() interface{} {
	return relation.NewRelationServiceListFollowerRelationArgs()
}

func newRelationServiceListFollowerRelationResult() interface{} {
	return relation.NewRelationServiceListFollowerRelationResult()
}

func listFriendRelationHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceListFriendRelationArgs)
	realResult := result.(*relation.RelationServiceListFriendRelationResult)
	success, err := handler.(relation.RelationService).ListFriendRelation(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceListFriendRelationArgs() interface{} {
	return relation.NewRelationServiceListFriendRelationArgs()
}

func newRelationServiceListFriendRelationResult() interface{} {
	return relation.NewRelationServiceListFriendRelationResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) ActionRelation(ctx context.Context, req *relation.DouyinRelationActionRequest) (r *relation.DouyinRelationActionResponse, err error) {
	var _args relation.RelationServiceActionRelationArgs
	_args.Req = req
	var _result relation.RelationServiceActionRelationResult
	if err = p.c.Call(ctx, "ActionRelation", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListFollowRelation(ctx context.Context, req *relation.DouyinRelationFollowListRequest) (r *relation.DouyinRelationFollowListResponse, err error) {
	var _args relation.RelationServiceListFollowRelationArgs
	_args.Req = req
	var _result relation.RelationServiceListFollowRelationResult
	if err = p.c.Call(ctx, "ListFollowRelation", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListFollowerRelation(ctx context.Context, req *relation.DouyinRelationFollowerListRequest) (r *relation.DouyinRelationFollowerListResponse, err error) {
	var _args relation.RelationServiceListFollowerRelationArgs
	_args.Req = req
	var _result relation.RelationServiceListFollowerRelationResult
	if err = p.c.Call(ctx, "ListFollowerRelation", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListFriendRelation(ctx context.Context, req *relation.DouyinRelationFriendListRequest) (r *relation.DouyinRelationFriendListResponse, err error) {
	var _args relation.RelationServiceListFriendRelationArgs
	_args.Req = req
	var _result relation.RelationServiceListFriendRelationResult
	if err = p.c.Call(ctx, "ListFriendRelation", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
