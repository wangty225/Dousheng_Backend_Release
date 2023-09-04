package main

import (
	"Dousheng_Backend/internal/dal/mysql"
	"Dousheng_Backend/internal/dal/redis"
	"Dousheng_Backend/internal/mircoservice/user/kitex-gen/user"
	"Dousheng_Backend/utils"
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// User implements the UserServiceImpl interface.
func (s *UserServiceImpl) User(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	// TODO: Your code here...

	// check request
	logger.Infof("get request\t:%v\n", req)

	result, ud := mysql.GetUserById(req.GetUserId())
	var message string
	if result.Error != nil {
		message = fmt.Sprintf("[user-server] %+v\n", result.Error)
		logger.Errorln(message)
		return &user.DouyinUserResponse{
			StatusCode: -1,
			StatusMsg:  &message,
			User:       user.NewUser(),
		}, nil
	} else if result.RowsAffected == 0 {
		message = "[user-server]未找到此用户！请检查输入信息！\n"
		logger.Errorln(message)
		resp.StatusCode = 0
		resp.StatusMsg = &message
		return &user.DouyinUserResponse{
			StatusCode: 0,
			StatusMsg:  &message,
			User:       user.NewUser(),
		}, nil
	}

	u := user.User{
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

	message = "ok\n"

	// send
	return &user.DouyinUserResponse{
		StatusCode: 0,
		StatusMsg:  &message,
		User:       &u,
	}, nil
}

// RegisterUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) RegisterUser(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	// TODO: Your code here...

	// check request
	logger.Infof("get request\t:%v\n", req)

	var message string
	var code int32
	if utils.ContainsInvalidCharacters(req.Password) {
		message = "[user-servver]密码包含字典外的字符, 注册失败！\n"
		logger.Errorln(message)
		return &user.DouyinUserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  &message,
			UserId:     0,
			Token:      "",
		}, nil

	}

	result, count := mysql.GetUserCountByName(&req.Username)
	if result.Error != nil {
		message := fmt.Sprintf("[user-server] %+v\n", result.Error)
		logger.Errorln(message)
		return &user.DouyinUserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  &message,
			UserId:     0,
			Token:      "",
		}, nil
	}
	if count != 0 {
		message = "[user-server]用户名已被注册，请重新输入！\n"
		logger.Errorln(message)
		return &user.DouyinUserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  &message,
			UserId:     0,
			Token:      "",
		}, nil
	}

	randomSalt := utils.GenerateRandomSalt()
	passwdCrypted := utils.GenerateSaltedMD5(req.Password, randomSalt)

	id := utils.GenerateRandomID()
	ud := mysql.UserDao{
		ID:              id,
		Username:        req.Username,
		FollowCount:     0,
		FollowerCount:   0,
		IsFollow:        false,
		Avatar:          "",
		BackgroundImage: "",
		Signature:       "",
		TotalFavorited:  0,
		WorkCount:       0,
		FavoriteCount:   0,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Time{},
		DeletedAt:       gorm.DeletedAt{},
	}

	ad := mysql.AuthDao{
		UserId:        ud.ID,
		PasswordCrypt: passwdCrypted,
		Salt:          randomSalt,
	}

	// 开始事务
	code, message = mysql.RegisterUser(&ad, &ud)
	// 结束事务

	token, _ := Jwt.GenTokenWithId(ud.ID)
	return &user.DouyinUserRegisterResponse{
		StatusCode: code,
		StatusMsg:  &message,
		UserId:     ud.ID,
		Token:      token, // todo
	}, nil

}

// LoginUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) LoginUser(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	// TODO: Your code here...
	var message string
	var code int32

	result0, ud := mysql.GetUserByName(&req.Username)
	if result0.Error != nil {
		code = -1
		message = fmt.Sprintf("[user-server] %+v\n", result0.Error)
		logger.Errorln(message)
		return &user.DouyinUserLoginResponse{
			StatusCode: code,
			StatusMsg:  &message,
			UserId:     0,
			Token:      "",
		}, nil
	}
	if result0.RowsAffected == 0 {
		code = -1
		message = fmt.Sprintf("[user-server] %+v\n", "未查询到该用户，请检查用户名或注册后登录！")
		logger.Errorln(message)
		return &user.DouyinUserLoginResponse{
			StatusCode: code,
			StatusMsg:  &message,
			UserId:     0,
			Token:      "",
		}, nil
	}

	result, ad := mysql.GetAuthById(&ud.ID)
	if result.Error != nil {
		code = -1
		message = fmt.Sprintf("[user-server] %+v\n", result.Error)
		logger.Errorln(message)
		return &user.DouyinUserLoginResponse{
			StatusCode: code,
			StatusMsg:  &message,
			UserId:     0,
			Token:      "",
		}, nil
	}
	if result.RowsAffected == 0 {
		code = -1
		message = fmt.Sprintf("[user-server] %+v\n", "该用户尚未设置密码，请先进行密码重置操作！")
		logger.Errorln(message)
		return &user.DouyinUserLoginResponse{
			StatusCode: code,
			StatusMsg:  &message,
			UserId:     0,
			Token:      "",
		}, nil
	}

	password := &req.Password
	salt := ad.Salt

	encryptedPasswd := utils.GenerateSaltedMD5(*password, salt)
	if encryptedPasswd == ad.PasswordCrypt {
		code = 0
		token, _ := Jwt.GenTokenWithId(ud.ID)
		err := redis.SetTokenInfoInRedis(token, redis.TokenInfo{
			UserID:         ud.ID,
			Token:          token,
			ExpirationTime: time.Now().Add(time.Hour * 1),
		})
		if err != nil {
			code = -1
			message = fmt.Sprintf("[user-server] %+v\n", "设置登录令牌失败，登录失败！")
			logger.Errorln(message)
			return &user.DouyinUserLoginResponse{
				StatusCode: code,
				StatusMsg:  &message,
				UserId:     0,
				Token:      "nil",
			}, nil

		}

		message = fmt.Sprintf("[user-server] %+v\n", "登陆成功！")
		logger.Infoln(message)
		return &user.DouyinUserLoginResponse{
			StatusCode: code,
			StatusMsg:  &message,
			UserId:     ud.ID,
			Token:      token, // todo: done! 登陆成功 add token
		}, nil
	} else {
		code = -1
		message = fmt.Sprintf("[user-server] %+v\n", "密码错误！")
		logger.Infoln(message)
		return &user.DouyinUserLoginResponse{
			StatusCode: code,
			StatusMsg:  &message,
			UserId:     0,
			Token:      "",
		}, nil
	}
	//return
}
