package mysql

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type UserDao struct {
	ID              int64          `gorm:"column:id;autoIncrement;primaryKey"`
	Username        string         `gorm:"column:username"`
	FollowCount     int64          `gorm:"column:follow_count"`
	FollowerCount   int64          `gorm:"column:follower_count"`
	IsFollow        bool           `gorm:"column:is_follow"`
	Avatar          string         `gorm:"column:avatar"`
	BackgroundImage string         `gorm:"column:background_image"`
	Signature       string         `gorm:"column:signature"`
	TotalFavorited  int64          `gorm:"column:total_favorited"`
	WorkCount       int64          `gorm:"column:work_count"`
	FavoriteCount   int64          `gorm:"column:favorite_count"`
	CreatedAt       time.Time      `gorm:"column:created_at"`
	UpdatedAt       time.Time      `gorm:"column:updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (UserDao) TableName() string {
	return "users"
}

type AuthDao struct {
	UserId        int64  `gorm:"column:user_id"`
	PasswordCrypt string `gorm:"column:passwd_crypt"`
	Salt          string `gorm:"column:salt"`
}

func (AuthDao) TableName() string {
	return "auth"
}

//func (ud UserDao) String() string {
//	return fmt.Sprintf("{ ID: %v\t UserName: %v\t FollowCount: %v\t FollowerCount: %v\t IsFollow: %v\n"+
//		"Avatar: %v\t BackgroundImage: %v\n"+
//		"Signature: %v\n"+
//		"TotalFavorited: %v\t WorkCount: %v\t FavoriteCount: %v }\n",
//		ud.ID, ud.Username, ud.FollowCount, ud.FollowerCount, ud.IsFollow,
//		ud.Avatar, ud.BackgroundImage, ud.Signature, ud.TotalFavorited, ud.WorkCount, ud.FavoriteCount)
//}

func GetUserById(id int64) (result *gorm.DB, ud UserDao) {
	// TODO: 数据库业务逻辑
	result = DBMysql.Limit(1).Where("id=?", id).Find(&ud)
	return
}

func GetUserCountByName(username *string) (result *gorm.DB, count int64) {
	// TODO: 数据库业务逻辑
	result = DBMysql.Model(&UserDao{}).Where("username=?", username).Count(&count)
	return
}

// RegisterUser
// 多操作函数返回（code, message）
// 单操作函数返回（result, error）
func RegisterUser(authDao *AuthDao, userDao *UserDao) (code int32, message string) {
	tx := DBMysql.Begin()

	result0 := tx.Create(&userDao)
	if result0.Error != nil {
		code = -1
		message = fmt.Sprintf("[mysql]注册失败，请稍后重试。\n%+v\n", result0.Error)
		logger.Errorln(message)
		tx.Rollback()
		return
	}
	if result0.RowsAffected == 0 {
		code = -1
		message = fmt.Sprintf("[mysql]注册失败，请稍后重试。\n%+v\n", "RowsAffected_user == 0")
		logger.Infoln(message)
		tx.Rollback()
		return
	}

	result1 := tx.Create(&authDao)
	if result1.Error != nil {
		code = -1
		message = fmt.Sprintf("[mysql]注册失败，请稍后重试。\n%+v\n", result1.Error)
		logger.Errorln(message)
		tx.Rollback()
		return
	}
	if result1.RowsAffected == 0 {
		code = -1
		message = fmt.Sprintf("[mysql]注册失败，请稍后重试。\n%+v\n", "RowsAffected_auth == 0")
		logger.Infoln(message)
		tx.Rollback()
		return
	}
	err := tx.Commit().Error
	if err != nil {
		code = -1
		message = fmt.Sprintf("[mysql] 注册失败，请稍后重试。\n%+v\n", err)
		logger.Errorln(message)
		return
	}
	code = 0
	message = "[mysql]注册成功！\n"
	logger.Infoln(message)
	return
}

func GetAuthById(id *int64) (result *gorm.DB, authDao AuthDao) {
	result = DBMysql.First(&authDao, "user_id=?", id)
	return
}

func GetUserByName(username *string) (result *gorm.DB, userDao UserDao) {
	result = DBMysql.First(&userDao, "username=?", username)
	return
}
