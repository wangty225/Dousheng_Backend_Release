package utils

import (
	"log"
	"path/filepath"
	"runtime"
)

// GetCallerAbsPath :
// 返回调用此函数的文件所在的绝对路径
// 主要用于local/mysql.go和local/redis.go下读取yml文件的绝对位置定位
func GetCallerAbsPath() (dirPath string) {
	_, filePath, _, ok := runtime.Caller(1)
	if !ok {
		log.Fatalln("Failed to get caller information for config file! Stopped!")
	}
	// 使用filepath.Dir获取目录部分
	dirPath = filepath.Dir(filePath)
	return
}
