package main

import (
	"Dousheng_Backend/internal/dal/mysql"
	"Dousheng_Backend/internal/mircoservice/interaction/kitex-gen/interaction"
	"Dousheng_Backend/internal/mircoservice/interaction/kitex-gen/user"
	"Dousheng_Backend/internal/mircoservice/interaction/kitex-gen/video"
	"Dousheng_Backend/utils"
	"context"
	"fmt"
)

// InteractionServiceImpl implements the last service interface defined in the IDL.
type InteractionServiceImpl struct{}

// ActionFavorite implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) ActionFavorite(ctx context.Context, req *interaction.DouyinFavoriteActionRequest) (resp *interaction.DouyinFavoriteActionResponse, err error) {
	// TODO: Your code here...
	// check request
	logger.Infoln("get request\t:%v", req)

	var message string
	userID, err := Jwt.GetIdFromToken(req.GetToken())
	if err != nil {
		logger.Errorln("[interacton-server]", err)
		message = "[interacton-server] token解析错误，无法获取用户ID"
		res := &interaction.DouyinFavoriteActionResponse{
			StatusCode: -1,
			StatusMsg:  &message,
		}
		return res, nil
	}
	code, message := mysql.FavoriteAction(userID, req.GetVideoId(), req.GetActionType())
	if code != 0 {
		res := &interaction.DouyinFavoriteActionResponse{
			StatusCode: -1,
			StatusMsg:  &message,
		}
		return res, nil
	}
	message = "ok"
	return &interaction.DouyinFavoriteActionResponse{
		StatusCode: 0,
		StatusMsg:  &message,
	}, nil
}

// ListFavorite implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) ListFavorite(ctx context.Context, req *interaction.DouyinFavoriteListRequest) (resp *interaction.DouyinFavoriteListResponse, err error) {
	// TODO: Your code here...
	logger.Infof("get request\t:%v\n", req)
	var message string
	videoList := make([]*video.Video, 0)
	if len(req.Token) == 0 {
		message = "[interacton-server] token missed"
		logger.Errorf("%s\n", message)
		res := &interaction.DouyinFavoriteListResponse{
			StatusCode: -1,
			StatusMsg:  &message,
			VideoList:  videoList,
		}
		return res, nil
	}

	vdl, err := mysql.GetFavoriteListByUserId(req.GetUserId())
	if err != nil {
		message = fmt.Sprintf("[interacton-server] 查询喜欢视频列表失败: %v", err)
		logger.Errorf("%s\n", message)
		res := &interaction.DouyinFavoriteListResponse{
			StatusCode: -1,
			StatusMsg:  &message,
			VideoList:  videoList,
		}
		return res, nil
	}

	for _, vdi := range vdl {
		result, ud := mysql.GetUserById(vdi.UserId)
		if result.Error != nil {
			message = "[interacton-server] 未找到作者信息\n"
			logger.Errorln(message)
			return &interaction.DouyinFavoriteListResponse{
				StatusCode: -1,
				StatusMsg:  &message,
				VideoList:  videoList,
			}, result.Error
		}
		author := user.User{
			Id:              ud.ID,
			Name:            ud.Username,
			FollowCount:     ud.FollowCount,
			FollowerCount:   ud.FollowerCount,
			IsFollow:        ud.IsFollow,
			Avator:          &ud.Avatar,
			BackgroundImage: &ud.BackgroundImage,
			Signature:       &ud.Signature,
			TotalFavorited:  &ud.TotalFavorited,
			WorkCount:       &ud.WorkCount,
			FavoriteCount:   &ud.FavoriteCount,
		}
		videoList = append(videoList, &video.Video{
			Id:            vdi.ID,
			Author:        &author,
			PlayUrl:       vdi.PlayURL,
			CoverUrl:      vdi.CoverURL,
			FavoriteCount: vdi.FavoriteCount,
			CommentCount:  vdi.CommentCount,
			IsFavorite:    vdi.IsFavorite == 1,
			Title:         vdi.Title,
		})
	}
	message = "ok"
	resp1 := &interaction.DouyinFavoriteListResponse{
		StatusCode: 0,
		StatusMsg:  &message,
		VideoList:  videoList,
	}
	return resp1, nil

}

// ActionComment implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) ActionComment(ctx context.Context, req *interaction.DouyinCommentActionRequest) (resp *interaction.DouyinCommentActionResponse, err error) {
	// TODO: Your code here...
	logger.Infoln("get request\t:%v", req)

	// 插入或删除comments
	// 修改video表的count数
	var message string
	if req.ActionType != 1 && req.ActionType != 2 {
		message = "[interacton-server] 操作类型错误，请检查输入参数！"
		logger.Errorf("%v\n", message)
		res := &interaction.DouyinCommentActionResponse{
			StatusCode: -1,
			StatusMsg:  &message,
			//Comment:    interaction.NewComment(),
		}
		return res, nil
	}

	vidoeId := req.VideoId
	commentText := req.CommentText
	userID, err := Jwt.GetIdFromToken(req.GetToken())
	if err != nil {
		logger.Errorln("[interacton-server]", err.Error())
		message = "[interacton-server] token解析错误，无法获取用户ID"
		res := &interaction.DouyinCommentActionResponse{
			StatusCode: -1,
			StatusMsg:  &message,
			//Comment:    interaction.NewComment(),
		}
		return res, nil
	}
	var commentId int64
	if req.ActionType == 1 {
		commentId = utils.GenerateRandomID()
	} else {
		commentId = *req.CommentId
	}

	logger.Infof("[interacton-server] GeneratedCommentId: %v\n", commentId)
	code, message, commentDao := mysql.CommentAction(userID, vidoeId, req.ActionType, commentText, commentId)
	if code != 0 {
		logger.Errorln(message)
		res := &interaction.DouyinCommentActionResponse{
			StatusCode: code,
			StatusMsg:  &message,
			//Comment:    interaction.NewComment(),
		}
		return res, nil
	}
	result, ud := mysql.GetUserById(commentDao.UserId)
	if result.Error != nil || result.RowsAffected == 0 {
		logger.Errorln("[interacton-server]", result.Error)
		message = "[interacton-server] 查询评论用户出错！\n"
		res := &interaction.DouyinCommentActionResponse{
			StatusCode: -1,
			StatusMsg:  &message,
			//Comment:    interaction.NewComment(),
		}
		return res, nil
	}
	message = "ok"
	return &interaction.DouyinCommentActionResponse{
		StatusCode: 0,
		StatusMsg:  &message,
		Comment: &interaction.Comment{
			Id: commentDao.ID,
			User: &user.User{
				Id:              ud.ID,
				Name:            ud.Username,
				FollowCount:     ud.FollowCount,
				FollowerCount:   ud.FollowerCount,
				IsFollow:        ud.IsFollow,
				Avator:          &ud.Avatar,
				BackgroundImage: &ud.BackgroundImage,
				Signature:       &ud.Signature,
				TotalFavorited:  &ud.TotalFavorited,
				WorkCount:       &ud.WorkCount,
				FavoriteCount:   &ud.FavoriteCount,
			},
			Content:    commentDao.Content,
			CreateDate: commentDao.CreatedAt.Format("01-02"),
		},
	}, nil
}

// ListComment implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) ListComment(ctx context.Context, req *interaction.DouyinCommentListRequest) (resp *interaction.DouyinCommentListResponse, err error) {
	// TODO: Your code here...
	logger.Infoln("get request\t:%v", req)
	var message string
	commentList := make([]*interaction.Comment, 0)
	if len(req.Token) == 0 {
		message = "[interacton-server] token missed"
		logger.Errorf("%s\n", message)
		res := &interaction.DouyinCommentListResponse{
			StatusCode:  -1,
			StatusMsg:   &message,
			CommentList: commentList,
		}
		return res, nil
	}

	cml, err := mysql.GetCommentListByVideoId(req.GetVideoId())
	if err != nil {
		message = fmt.Sprintf("[interacton-server] 查询视频评论失败: %v", err)
		logger.Errorf("%s\n", message)
		res := &interaction.DouyinCommentListResponse{
			StatusCode:  -1,
			StatusMsg:   &message,
			CommentList: commentList,
		}
		return res, nil
	}

	for _, cmi := range cml {
		result, ud := mysql.GetUserById(cmi.UserId)
		if result.Error != nil || result.RowsAffected == 0 {
			logger.Errorln("[interacton-server]", result.Error)
			message = "[interacton-server] 查询评论用户出错！\n"
			res := &interaction.DouyinCommentListResponse{
				StatusCode:  -1,
				StatusMsg:   &message,
				CommentList: commentList,
			}
			return res, nil
		}
		commentList = append(commentList, &interaction.Comment{
			Id: cmi.ID,
			User: &user.User{
				Id:              ud.ID,
				Name:            ud.Username,
				FollowCount:     ud.FollowCount,
				FollowerCount:   ud.FollowerCount,
				IsFollow:        ud.IsFollow,
				Avator:          &ud.Avatar,
				BackgroundImage: &ud.BackgroundImage,
				Signature:       &ud.Signature,
				TotalFavorited:  &ud.TotalFavorited,
				WorkCount:       &ud.WorkCount,
				FavoriteCount:   &ud.FavoriteCount,
			},
			Content:    cmi.Content,
			CreateDate: cmi.CreatedAt.Format("01-02"),
		})
	}

	message = "[interacton-server] 评论列表查询成功！"
	logger.Infoln(message)
	return &interaction.DouyinCommentListResponse{
		StatusCode:  0,
		StatusMsg:   &message,
		CommentList: commentList,
	}, nil
}
