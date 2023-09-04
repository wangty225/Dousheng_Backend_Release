package utils

import (
	"encoding/base64"
	"math/rand"
	"time"
)

func GenerateRandomID() int64 {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成一个10位的随机整数，不以0开头
	min := int64(1000000000) // 10^9
	max := int64(9999999999) // 10^10 - 1
	randomID := min + rand.Int63n(max-min+1)

	return randomID
}

func GenerateRandomSalt() string {
	// 定义生成盐值的字节长度
	saltLength := 16 // 16字节的盐值通常足够安全

	// 生成随机的字节序列作为盐值
	salt := make([]byte, saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}

	// 使用Base64编码将字节序列转换为字符串
	saltString := base64.StdEncoding.EncodeToString(salt)

	return saltString
}
