package middleware

import (
	"Dousheng_Backend/internal/dal/redis"
	"Dousheng_Backend/utils/jwt"
	"Dousheng_Backend/utils/zap"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"strings"
	"time"
)

var JWT = jwt.NewJWT()

func TokenAuthMiddleware(jwt jwt.JWT, skipRoutes ...string) app.HandlerFunc {
	logger := zap.InitDefaultLogger()
	// TODO: signKey可以保存在环境变量中，而不是硬编码在代码里，可以通过获取环境变量的方式获得signkey
	return func(ctx context.Context, c *app.RequestContext) {
		// 对于skip的路由不对他进行token鉴权
		for _, skipRoute := range skipRoutes {
			if skipRoute == c.FullPath() {
				c.Next(ctx)
				return
			}
		}

		// 从处理get post请求中获取token
		var token string
		//token = c.Query("token")
		if string(c.Request.Method()[:]) == "GET" {
			token = c.Query("token")
		} else if string(c.Request.Method()[:]) == "POST" {
			if strings.Contains(c.Request.Header.Get("Content-Type"), "multipart/form-data") {
				token = c.PostForm("token")
			} else {
				token = c.Query("token")
			}
		} else {
			// Unsupport request method
			responseWithError(ctx, c, http.StatusBadRequest, "bad request")
			logger.Errorln("bad request")
			return
		}
		if token == "" {
			responseWithError(ctx, c, http.StatusUnauthorized, "token required")
			logger.Errorln("token required")
			// 提前返回
			return
		}

		tk, flag := redis.GetTokenInfoInRedis(token)
		if !flag { // redis中存在token，则可以刷新token
			responseWithError(ctx, c, http.StatusUnauthorized, "登陆过期，请重新登录！")
			logger.Errorln("登陆过期，请重新登录！")

			return
		}

		if tk.ExpirationTime.Unix() < time.Now().Unix() { // token已过期，需要刷新
			//// 方式一：直接返回重新登陆
			//err := errors.New("token过期，请重新登录")
			//responseWithError(ctx, c, http.StatusUnauthorized, err)
			//logger.Errorln(err)
			//return

			// 方式二：无感刷新token（客户端暂不支持）
			// 无感刷新条件：request中有required token 且对应的response有optional token，
			// 客户端判断response中token是否存在而决定下次携带的token是否需要替换成新的token
			// 本项目的api中，客户端不支持无感刷新

			tokenNew, err := JWT.GenTokenWithId(tk.UserID)
			if err != nil {
				responseWithError(ctx, c, http.StatusUnauthorized, err.Error())
				logger.Errorln(err)
				return
			}
			// todo: 存入redis新token，删除旧token (redis事务)
			err = redis.RefreshUseTokenTx(tk, redis.TokenInfo{
				UserID:         tk.UserID,
				Token:          tokenNew,
				ExpirationTime: time.Now().Add(time.Hour * 1),
			})
			if err != nil {
				responseWithError(ctx, c, http.StatusUnauthorized, err.Error())
				logger.Errorln(err)
				return
			}
			token = tokenNew
		}

		//claim, err := jwt.ParseToken(token)
		//
		//if err != nil {
		//	responseWithError(ctx, c, http.StatusUnauthorized, err.Error())
		//	logger.Errorln(err.Error())
		//	return
		//}

		// 在上下文中向下游传递token
		c.Set("Token", token)
		c.Set("Id", tk.UserID)

		c.Next(ctx) // 交给下游中间件
	}
}
