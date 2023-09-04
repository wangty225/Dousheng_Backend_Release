// Code generated by hertz generator.

package relation_hertz

import (
	"context"

	relation_hertz "Dousheng_Backend/hertz/hertz-gen/biz/model/relation_hertz"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// ActionRelation .
// @router /douyin/relation/action/ [POST]
func ActionRelation(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation_hertz.DouyinRelationActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(relation_hertz.DouyinRelationActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// ListFollowRelation .
// @router /douyin/relation/follow/list/ [GET]
func ListFollowRelation(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation_hertz.DouyinRelationFollowListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(relation_hertz.DouyinRelationFollowListResponse)

	c.JSON(consts.StatusOK, resp)
}

// ListFollowerRelation .
// @router /douyin/relation/follower/list/ [GET]
func ListFollowerRelation(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation_hertz.DouyinRelationFollowerListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(relation_hertz.DouyinRelationFollowerListResponse)

	c.JSON(consts.StatusOK, resp)
}

// ListFriendRelation .
// @router /douyin/relation/friend/list/ [GET]
func ListFriendRelation(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation_hertz.DouyinRelationFriendListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(relation_hertz.DouyinRelationFriendListResponse)

	c.JSON(consts.StatusOK, resp)
}
