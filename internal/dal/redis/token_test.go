package redis

import (
	"fmt"
	"github.com/cloudwego/kitex/tool/internal_pkg/log"
	"testing"
	"time"
)

var testToken = "token-value"

func TestSetTokenInRedis(t *testing.T) {
	// 示例：生成一个 TokenInfo
	tokenInfo := TokenInfo{
		UserID:         123456,
		Token:          testToken,
		ExpirationTime: time.Now().Add(1 * time.Hour), // 设置过期时间为1小时后
	}

	// 存储 Token 到 Redis，并设置过期时间
	err := SetTokenInfoInRedis(tokenInfo.Token, tokenInfo)
	if err != nil {
		log.Info("failed to store token in redis")
		return
	}
	log.Info("success stored token in redis\n")
}

func TestGetTokenInfoInRedis(t *testing.T) {

	tkInfo, flag := GetTokenInfoInRedis(testToken)
	if !flag {
		fmt.Println("查询失败！")
	} else {
		fmt.Println("查询成功！")
		fmt.Printf("%+v\n", tkInfo)
	}
}

func TestDelTokenInfoInRedis(t *testing.T) {
	nums, err := DelTokenInfoInRedis(testToken)
	if err != nil {
		return
	}
	log.Infof("deleted %d records in redis successfully\n", nums)
}
