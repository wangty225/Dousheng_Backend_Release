package main

import (
	"Dousheng_Backend/internal/dal/mysql"
	"Dousheng_Backend/internal/mircoservice/video/kitex-gen/user"
	"Dousheng_Backend/internal/oss"
	"time"

	"Dousheng_Backend/internal/mircoservice/video/kitex-gen/video"
	"context"
	"fmt"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *video.DouyinFeedRequest) (resp *video.DouyinFeedResponse, err error) {
	// TODO: Your code here...
	logger.Infof("[video-server]get request\t:%+v\n", req)
	//fmt.Printf("%+v\n", *req.Token)
	//fmt.Printf("%+v\n", *req.LatestTime)
	var message string
	nextTime := time.Now().UnixMilli()
	videos := make([]*video.Video, 0)
	vd, err := mysql.GetVideoListAfterTime(req.LatestTime)
	if err != nil {
		message = fmt.Sprintf("[video-server] %+v\n", err)
		logger.Errorln(message)
		return &video.DouyinFeedResponse{
			StatusCode: -1,
			StatusMsg:  &message,
			VideoList:  videos,
			NextTime:   &nextTime,
		}, err
	}

	if len(vd) != 0 {
		nextTime = vd[len(vd)-1].CreatedAt.UnixMilli()
	}
	for _, vdi := range vd {
		result, ud := mysql.GetUserById(vdi.UserId)
		if result.Error != nil {
			message = "未找到作者信息"
			return &video.DouyinFeedResponse{
				StatusCode: -1,
				StatusMsg:  &message,
				VideoList:  videos,
				NextTime:   &nextTime,
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
		videos = append(videos, &video.Video{
			Id:            vdi.ID,
			Author:        &author,
			PlayUrl:       vdi.PlayURL,
			CoverUrl:      vdi.CoverURL,
			FavoriteCount: vdi.FavoriteCount,
			CommentCount:  vdi.CommentCount,
			IsFavorite:    false,
			Title:         vdi.Title,
		})
	}
	message = "ok"
	resp1 := &video.DouyinFeedResponse{
		StatusCode: 0,
		StatusMsg:  &message,
		VideoList:  videos,
		NextTime:   &nextTime,
	}
	return resp1, nil
}

// ActionPublish implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) ActionPublish(ctx context.Context, req *video.DouyinPublishActionRequest) (resp *video.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...
	logger.Infof("[video-server]get request\t:%+v\t %+v\n", req.Token, req.Title)
	var message string
	// // 废弃方法一
	//claims, err := Jwt.ParseToken(req.Token)
	//if err != nil {
	//	logger.Errorln(err.Error())
	//	message = "token解析错误，无法获取用户ID"
	//	res := &video.DouyinPublishActionResponse{
	//		StatusCode: -1,
	//		StatusMsg:  &message,
	//	}
	//	return res, nil
	//}
	//userID := claims.Id

	// 方法二
	userID, err := Jwt.GetIdFromToken(req.GetToken())
	if err != nil {
		logger.Errorln(err.Error())
		message = "[video-server] token解析错误，无法获取用户ID\n"
		logger.Errorln(message)
		res := &video.DouyinPublishActionResponse{
			StatusCode: -1,
			StatusMsg:  &message,
		}
		return res, nil
	}

	if len(req.Title) == 0 || len(req.Title) > 200 {
		message = "[video-server]标题不能为空且不能超过32个字符!\n"
		logger.Errorf("%s：%d\n", message, len(req.Title))
		res := &video.DouyinPublishActionResponse{
			StatusCode: -1,
			StatusMsg:  &message,
		}
		return res, nil
	}

	maxSize := 1024 // 最大上传大小为1GB
	size := len(req.Data)
	if size > maxSize*1024*1024 {
		message = fmt.Sprintf("该视频文件超出最大上传限制: %v", size)
		logger.Errorln(message)
		res := &video.DouyinPublishActionResponse{
			StatusCode: -1,
			StatusMsg:  &message,
		}
		return res, nil
	}
	createTimestamp := time.Now().UnixMilli()
	videoName := fmt.Sprintf("%d.mp4", createTimestamp)
	bucketPath := fmt.Sprintf("%v/%v", "video", userID)
	flag, err := oss.UploadObjToOss(bucketPath, videoName, req.Data)
	if err != nil || !flag {
		logger.Errorf("[video-server]上传视频文件失败: %v\n", err)
		message = "上传视频文件失败"
		res := &video.DouyinPublishActionResponse{
			StatusCode: -1,
			StatusMsg:  &message,
		}
		return res, nil
	}
	playurl, coverurl := oss.GetOssVideoUrlAndCoverUrl(userID, videoName)
	code, message := mysql.UploadVideo(mysql.VideoDao{
		UserId:        userID,
		PlayURL:       playurl,
		CoverURL:      coverurl,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    0,
		Title:         req.Title,
		CreatedAt:     time.Now(),
	})
	if code == 0 {
		logger.Infoln("[video-server]上传视频成功！")
		message = "ok"
		res := &video.DouyinPublishActionResponse{
			StatusCode: 0,
			StatusMsg:  &message,
		}
		return res, nil
	} else {
		logger.Errorf("[video-server]视频链接插入数据库失败: %v\n", err)
		message = "ok"
		res := &video.DouyinPublishActionResponse{
			StatusCode: -1,
			StatusMsg:  &message,
		}
		return res, nil
	}

}

// ListPublish implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) ListPublish(ctx context.Context, req *video.DouyinPublishListRequest) (resp *video.DouyinPublishListResponse, err error) {
	// TODO: Your code here...
	logger.Infoln("get request\t:%+v\t %+v", req.UserId, req.Token)
	var message string
	videoList := make([]*video.Video, 0)
	vdl, err := mysql.GetPublishListByUserId(req.UserId)
	if err != nil {
		message = fmt.Sprintf("[video-server]查询发布视频列表失败: %v", err)
		logger.Errorln(message)
		res := &video.DouyinPublishListResponse{
			StatusCode: -1,
			StatusMsg:  &message,
			VideoList:  videoList,
		}
		return res, nil
	}

	for _, vdi := range vdl {
		result, ud := mysql.GetUserById(vdi.UserId)
		if result.Error != nil {
			message = "未找到作者信息"
			logger.Errorln(message)
			return &video.DouyinPublishListResponse{
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
	resp1 := &video.DouyinPublishListResponse{
		StatusCode: 0,
		StatusMsg:  &message,
		VideoList:  videoList,
	}
	return resp1, nil
}
