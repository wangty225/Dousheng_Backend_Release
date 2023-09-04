package mysql

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type VideoDao struct {
	ID            int64  `gorm:"column:id;autoIncrement;primaryKey"`
	UserId        int64  `gorm:"column:user_id;"`
	PlayURL       string `gorm:"column:play_url"`
	CoverURL      string `gorm:"column:cover_url"`
	FavoriteCount int64  `gorm:"column:favorite_count"`
	CommentCount  int64  `gorm:"column:comment_count"`
	IsFavorite    int32  `gorm:"column:is_favorite"`
	Title         string `gorm:"column:title"`

	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	Deleted   gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (VideoDao) TableName() string {
	return "videos"
}

func GetVideoListAfterTime(lastTime *int64) (a []*VideoDao, err error) {
	var curTime int64
	if lastTime == nil || *lastTime == 0 {
		curTime = time.Now().UnixMilli()
		lastTime = &curTime
	} // BUG TODO
	videos := make([]*VideoDao, 0)

	result := DBMysql.Model(VideoDao{}).Limit(2).Order("created_at desc").Find(&videos, "created_at < ?", time.UnixMilli(*lastTime))

	//fmt.Printf("DAL: %+v\n", videos)
	//for _, vi := range videos {
	//	fmt.Printf("%v \t %v\n", vi.PlayURL, vi.ID)
	//}
	//fmt.Printf("lines: %+v\n", result.RowsAffected)

	return videos, result.Error
}

func UploadVideo(vd VideoDao) (code int32, message string) {
	tx := DBMysql.Begin()

	result0 := tx.Create(&vd)
	if result0.Error != nil || result0.RowsAffected == 0 {
		code = -1
		message = fmt.Sprintf("[mysql] 上传失败，请稍后重试。\n%+v\n", result0.Error)
		logger.Errorln(message)
		tx.Rollback()
		return
	}

	result1 := tx.Model(UserDao{}).Where("id=?", vd.UserId).Update("work_count", gorm.Expr("work_count + ?", 1))
	if result1.Error != nil || result1.RowsAffected == 0 {
		code = -1
		message = fmt.Sprintf("[mysql] 上传失败，无法更新用户作品数，请稍后重试。\n%+v\n", result0.Error)
		logger.Errorln(message)
		tx.Rollback()
		return
	}
	err := tx.Commit().Error
	if err != nil {
		code = -1
		message = fmt.Sprintf("[mysql] 上传失败，commit错误，请稍后重试。\n%+v\n", err)
		logger.Errorln(message)
		return
	}
	code = 0
	message = "[mysql]新增视频记录成功\n"
	logger.Infoln(message)
	return
}

func GetPublishListByUserId(userId int64) (videoDaoList []*VideoDao, err error) {
	result := DBMysql.Where("user_id=?", userId).Find(&videoDaoList)
	if result.Error != nil {
		message := fmt.Sprintf("[mysql] 查询投稿列表失败！\n%+v\n", result.Error)
		logger.Errorln(message)
		return videoDaoList, result.Error
	}
	message := fmt.Sprintf("[mysql] 查询投稿列表成功！\n")
	logger.Infoln(message)
	return videoDaoList, nil
}

func GetVideoByVideoId(videoId int64) (vd *VideoDao, err error) {
	result := DBMysql.Model(VideoDao{}).Where("id=?", videoId).Find(&vd)
	if result.Error != nil {
		return vd, result.Error
	}
	return vd, nil
}
