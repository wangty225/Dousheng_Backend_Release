package utils

import (
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"io"
	"strings"
)

// EncryptPassword : 加密后的字符串默认长度为60
func EncryptPassword(password string, salt string) (passwdCrypt string, err error) {
	// 将密码和盐值拼接在一起
	passwordWithSalt := []byte(password + salt)

	// 使用bcrypt生成哈希值
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordWithSalt, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	// 将哈希值转换为字符串
	passwdCrypt = string(hashedPassword)

	return passwdCrypt, nil
}

func GenerateSaltedMD5(password string, salt string) string {
	// 将密码和盐值拼接在一起
	saltedPassword := password + salt

	// 创建MD5哈希对象
	hasher := md5.New()

	// 计算MD5哈希值
	io.WriteString(hasher, saltedPassword)
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))

	return hashedPassword
}

// 定义常见密码许可字符集合
const passwordCharacterSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()_+-=,."

// ContainsInvalidCharacters 检查字符串是否包含字典外的字符
func ContainsInvalidCharacters(input string) bool {
	for _, char := range input {
		if !strings.ContainsRune(passwordCharacterSet, char) {
			return true
		}
	}
	return false
}
