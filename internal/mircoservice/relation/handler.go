package main

import (
	"Dousheng_Backend/internal/mircoservice/relation/kitex-gen/relation"
	"context"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// ActionRelation implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) ActionRelation(ctx context.Context, req *relation.DouyinRelationActionRequest) (resp *relation.DouyinRelationActionResponse, err error) {
	// TODO: Your code here...
	return
}

// ListFollowRelation implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) ListFollowRelation(ctx context.Context, req *relation.DouyinRelationFollowListRequest) (resp *relation.DouyinRelationFollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// ListFollowerRelation implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) ListFollowerRelation(ctx context.Context, req *relation.DouyinRelationFollowerListRequest) (resp *relation.DouyinRelationFollowerListResponse, err error) {
	// TODO: Your code here...
	return
}

// ListFriendRelation implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) ListFriendRelation(ctx context.Context, req *relation.DouyinRelationFriendListRequest) (resp *relation.DouyinRelationFriendListResponse, err error) {
	// TODO: Your code here...
	return
}
