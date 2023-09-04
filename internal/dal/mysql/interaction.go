package mysql

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type UserFavoriteVideoDao struct {
	ID      int64 `gorm:"column:id;autoIncrement;primaryKey"`
	UserId  int64 `gorm:"column:user_id"`
	VideoId int64 `gorm:"column:video_id"`
}

type CommentDao struct {
	ID        int64          `gorm:"column:id;autoIncrement;primaryKey"`
	UserId    int64          `gorm:"column:user_id"`
	VideoId   int64          `gorm:"column:video_id"`
	Content   string         `gorm:"column:content"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (UserFavoriteVideoDao) TableName() string {
	return "user_favorite_videos"
}

func (CommentDao) TableName() string {
	return "comments"
}

func FavoriteAction(userId int64, videoId int64, actionType int32) (code int32, message string) {
	// 1.判断点赞还是取消点赞
	// 2.插入/删除喜欢记录表
	// 3.ideo的favorite_count加一/减一
	// 4.user的favorite_count获赞数量加一/减一
	// 5.查询查询video的userid并将对应user的获赞总数total_favorited加一/减一
	var num int32
	if actionType == 2 {
		num = -1
	} else if actionType == 1 {
		num = 1
	}
	tx := DBMysql.Begin()

	var vd VideoDao
	ret0 := tx.Model(VideoDao{}).Select("user_id").First(&vd, "id=?", videoId)
	if ret0.Error != nil || ret0.RowsAffected == 0 {
		code = -1
		message = fmt.Sprintf("[mysql]查询视频信息失败，请稍后再试。\n%+v\n", ret0.Error)
		logger.Errorln(message)
		tx.Rollback()
		return
	}
	if actionType == 1 { // 点赞

		ret10 := tx.Where("user_id=? and video_id=?", userId, videoId).Find(&UserFavoriteVideoDao{})
		if ret10.Error != nil {
			code = -1
			message = fmt.Sprintf("[mysql]点赞操作失败，请稍后重试。\n%+v\n", ret10.Error)
			logger.Errorln(message)
			tx.Rollback()
			return
		}
		if ret10.RowsAffected != 0 {
			code = -1
			message = fmt.Sprintf("[mysql]已点赞。\n")
			logger.Errorln(message)
			tx.Rollback()
			return
		}

		ret1 := tx.Create(&UserFavoriteVideoDao{
			UserId:  userId,
			VideoId: videoId,
		})
		if ret1.Error != nil || ret1.RowsAffected == 0 {
			code = -1
			message = fmt.Sprintf("[mysql]点赞操作失败，请稍后重试。\n%+v\n", ret1.Error)
			logger.Errorln(message)
			tx.Rollback()
			return
		}
	} else if actionType == 2 { // 取消点赞
		ret1 := tx.Where("user_id=? and video_id=?", userId, videoId).Delete(&UserFavoriteVideoDao{})
		if ret1.Error != nil || ret1.RowsAffected == 0 {
			code = -1
			message = fmt.Sprintf("[mysql]取消点赞操作失败，请稍后重试。\n%+v\n", ret1.Error)
			logger.Errorln(message)
			tx.Rollback()
			return
		}
	}

	ret2 := tx.Model(VideoDao{}).Where("id=?", videoId).Update("favorite_count", gorm.Expr("favorite_count+?", num))
	if ret2.Error != nil || ret2.RowsAffected == 0 {
		code = -1
		message = fmt.Sprintf("[mysql]视频点赞失败，请稍后重试。\n%+v\n", ret2.Error)
		logger.Errorln(message)
		tx.Rollback()
		return
	}

	ret3 := tx.Model(UserDao{}).Where("id=?", userId).Update("favorite_count", gorm.Expr("favorite_count+?", num))
	if ret3.Error != nil || ret3.RowsAffected == 0 {
		code = -1
		message = fmt.Sprintf("[mysql]用户点赞失败，请稍后重试。\n%+v\n", ret3.Error)
		logger.Errorln(message)
		tx.Rollback()
		return
	}

	ret4 := tx.Model(UserDao{}).Where("id=?", vd.UserId).Update("total_favorited", gorm.Expr("total_favorited+?", num))
	if ret4.Error != nil || ret4.RowsAffected == 0 {
		code = -1
		message = fmt.Sprintf("[mysql]作者点赞总数增加失败，请稍后重试。\n%+v\n", ret4.Error)
		logger.Errorln(message)
		tx.Rollback()
		return
	}

	err := tx.Commit().Error
	if err != nil {
		code = -1
		message = fmt.Sprintf("[mysql]点赞操作失败，请稍后重试！\n%+v\n", err)
		logger.Errorln(message)
		tx.Rollback()
		return
	}
	code = 0
	if actionType == 1 {
		message = "[mysql]点赞操作成功！\n"
	} else if actionType == 2 {
		message = "[mysql]取消点赞操作成功\n"
	} else {
		message = "[mysql]??操作成功。"
	}
	logger.Infoln(message)
	return
}

func GetFavoriteListByUserId(userId int64) (videoDaoList []*VideoDao, err error) {
	userFavoriteVideosDaoList := make([]*UserFavoriteVideoDao, 0)
	var message string
	result := DBMysql.Model(UserFavoriteVideoDao{}).Where("user_id=?", userId).Find(&userFavoriteVideosDaoList)
	if result.Error != nil {
		message = "[mysql]查询用户喜欢列表失败！"
		logger.Errorln(message)
		return videoDaoList, result.Error
	}
	for _, ufvdi := range userFavoriteVideosDaoList {
		vdi, err := GetVideoByVideoId(ufvdi.VideoId)
		if err != nil {
			message = "[mysql]用户喜欢列表根据ID查询视频信息失败！"
			logger.Errorln(message)
			return videoDaoList, err
		}
		videoDaoList = append(videoDaoList, vdi)
	}
	message = "[mysql]查询用户喜欢列表成功！"
	logger.Infoln(message)
	return videoDaoList, nil
}

func CommentAction(userId int64, videoId int64, actionType int32, commentText *string, commentId int64) (code int32, message string, commentDao *CommentDao) {
	var content string
	if actionType == 2 || commentText == nil {
		content = ""
	} else {
		content = *commentText
	}
	commentDao = &CommentDao{
		ID:      commentId,
		UserId:  userId,
		VideoId: videoId,
		Content: content,
	}

	// 插入或删除comments
	// 修改video表的count数
	switch actionType {
	case 1: // 评论
		tx := DBMysql.Begin()
		result0 := tx.Create(commentDao)
		if result0.Error != nil || result0.RowsAffected == 0 {
			code = -1
			message = fmt.Sprintf("[mysql]评论记录失败：%v\n", result0.Error)
			logger.Errorln(message)
			tx.Rollback()
			return
		}
		result1 := tx.Model(VideoDao{}).Where("id=?", videoId).Update("comment_count", gorm.Expr("comment_count+?", 1))
		if result1.Error != nil || result1.RowsAffected == 0 {
			code = -1
			message = fmt.Sprintf("[mysql]评论数累加失败：%v\n", result1.Error)
			logger.Errorln(message)
			tx.Rollback()
			return
		}
		err := tx.Commit().Error
		if err != nil {
			code = -1
			message = fmt.Sprintf("[mysql]评论提交失败，请稍后重试！\n%+v\n", err)
			logger.Errorln(message)
			tx.Rollback()
			return
		}
		code = 0
		message = "[mysql]评论操作成功！\n"
		logger.Infoln(message)
		return code, message, commentDao
	case 2: // 删除评论
		tx := DBMysql.Begin()
		result0 := tx.Where("id=?", commentDao.ID).Delete(&CommentDao{})
		if result0.Error != nil || result0.RowsAffected == 0 {
			code = -1
			message = fmt.Sprintf("[mysql]删除评论操作失败，请稍后重试。\n%+v\n", result0.Error)
			logger.Errorln(message)
			tx.Rollback()
			return
		}
		result1 := tx.Model(VideoDao{}).Where("id=?", videoId).Update("comment_count", gorm.Expr("comment_count-?", 1))
		if result1.Error != nil || result1.RowsAffected == 0 {
			code = -1
			message = fmt.Sprintf("[mysql]删除评论记录失败，请稍后重试。\n%+v\n", result1.Error)
			logger.Errorln(message)
			tx.Rollback()
			return
		}

		err := tx.Commit().Error
		if err != nil {
			code = -1
			message = fmt.Sprintf("[mysql]删除评论记录失败，请稍后重试！\n%+v\n", err)
			logger.Errorln(message)
			tx.Rollback()
			return
		}
		code = 0
		message = "[mysql]评论操作成功！\n"
		logger.Infoln(message)
		return code, message, commentDao
	default:
		code = -1
		message = "[mysql]评论类型不确定，操作失败!\n"
		logger.Errorln(message)
	}
	return
}

func GetCommentListByVideoId(videoId int64) (commentDaoList []*CommentDao, err error) {
	commentDaoList = make([]*CommentDao, 0)
	var message string
	result := DBMysql.Model(CommentDao{}).Where("video_id=?", videoId).Find(&commentDaoList)
	if result.Error != nil {
		message = "[mysql]查询评论列表失败！\n"
		logger.Errorln(message)
		return commentDaoList, result.Error
	}
	message = "[mysql]查询评论列表成功！\n"
	logger.Infoln(message)
	return commentDaoList, nil
}
