package redis

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"time"
)

type TokenInfo struct {
	UserID         int64
	Token          string
	ExpirationTime time.Time
}

func SetTokenInfoInRedis(token string, tokenInfo TokenInfo) error {
	var ctx = context.Background()
	// 序列化结构体为 JSON 字符串
	data, err := json.Marshal(tokenInfo)
	if err != nil {
		logger.Errorln("[redis]Error encoding struct to JSON:", err)
		return err
	}
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(16)
	// 使用 Set 方法将 Token 存储到 Redis，并设置随机过期时间，防止大规模到期引发缓存雪崩
	err = DBRedis.Set(ctx, token, data, time.Hour*2+time.Minute*time.Duration(randomInt)).Err()
	if err != nil {
		logger.Errorln("[redis]Error setting key:", err)
		return err
	}
	logger.Infoln("[redis]Token stored with expiration duration: 2h")
	return nil
}

func GetTokenInfoInRedis(token string) (t TokenInfo, flag bool) {
	var ctx = context.Background()
	value, err := DBRedis.Get(ctx, token).Result()
	if err == redis.Nil {
		logger.Errorln("[redis] token not found: ")
		return t, false
	}
	logger.Errorln("[redis] token found: ", token)
	err = json.Unmarshal([]byte(value), &t)
	if err != nil {
		logger.Errorln("[redis]Error decoding JSON to struct:", err)
		return t, false
	}
	return t, true
}

// DelTokenInfoInRedis 返回影响的记录数和error信息
func DelTokenInfoInRedis(token string) (nums int64, err error) {
	nums, err = DBRedis.Del(context.Background(), token).Result()
	return
}

// RefreshUseTokenTx
// todo: redis事务
//  1. 插入新的键值
//  2. 删除旧的键值
func RefreshUseTokenTx(oldTokenInfo TokenInfo, newTokenInfo TokenInfo) (err error) {
	pipe := DBRedis.TxPipeline()
	defer func(pipe redis.Pipeliner) {
		err := pipe.Close()
		if err != nil {
			logger.Errorln("[redis]Error happened while close pip:", err)
		}
	}(pipe)
	ctx := context.Background()
	pipe.Del(ctx, oldTokenInfo.Token)
	data, err := json.Marshal(newTokenInfo)
	if err != nil {
		//logger.Errorln("Error encoding struct to JSON:", err)
		logger.Errorln("[redis]Error encoding struct to JSON:", err)
		err2 := pipe.Discard()
		if err2 != nil {
			logger.Errorln("[redis]Error happened while close pip:", err2)
			return err2
		}
		return
	}
	td := newTokenInfo.ExpirationTime.Sub(time.Now()) + time.Hour // 设置+1h的刷新过期时间
	pipe.Set(ctx, newTokenInfo.Token, data, td)                   // 设hi两个小时的过期时间
	// 提交事务
	_, err = pipe.Exec(ctx)
	if err != nil {
		logger.Errorln("Transaction failed:", err)
		return
	}
	logger.Infoln("Transaction successful!")
	return nil
}
