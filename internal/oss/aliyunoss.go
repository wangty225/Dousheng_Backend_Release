package oss

import (
	"Dousheng_Backend/utils/config"
	"Dousheng_Backend/utils/zap"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
	"strconv"
	"strings"
)

//TODO:重构成接口类型，每种存储介质只需实现对应的方法

type AliyunCfg struct {
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	BucketName      string
}

var aliyunCfg AliyunCfg
var AliyunOssClient *oss.Client

// var logger = zap.InitLogger(config.InitConfig("../../config/logger/logger.yml"))
var logger = zap.InitDefaultLogger()

func init() {
	aliyunOSSConfig := config.InitConfig("./config/yml/aliyun_oss.yml").Sub("aliyun_oss")
	//aliyunOSSConfig := config.InitConfig("../../config/yml/aliyun_oss.yml").Sub("aliyun_oss")
	if aliyunOSSConfig == nil {
		logger.Errorln("'aliyun_oss' key not found, pls check you config")
		log.Fatalln("'aliyun_oss' key not found, pls check you config")
	}
	aliyunCfg = AliyunCfg{
		Endpoint:        aliyunOSSConfig.GetString("Endpoint"),
		BucketName:      aliyunOSSConfig.GetString("BucketName"),
		AccessKeyID:     aliyunOSSConfig.GetString("AccessKeyID"),
		AccessKeySecret: aliyunOSSConfig.GetString("AccessKeySecret"),
	}
	client, err := oss.New(aliyunCfg.Endpoint, aliyunCfg.AccessKeyID, aliyunCfg.AccessKeySecret)
	if err != nil {
		panic(err)
	}
	AliyunOssClient = client
}

func UploadObjToOss(bucketPath string, objectName string, data []byte) (bool, error) {
	bucket, err := AliyunOssClient.Bucket(aliyunCfg.BucketName)
	if err != nil {
		return false, err
	}
	contentType := "video/mp4"
	uri := fmt.Sprintf("%s/%s", bucketPath, objectName)
	err = bucket.PutObject(uri, strings.NewReader(string(data)), oss.ContentType(contentType))
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}

func DeleteObjFromOss(bucketPath string, objectName string) (bool, error) {
	bucket, err := AliyunOssClient.Bucket(aliyunCfg.BucketName)
	if err != nil {
		return false, err
	}
	uri := fmt.Sprintf("%s/%s", bucketPath, objectName)
	err = bucket.DeleteObject(uri)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}

// GetOssVideoUrlAndCoverUrl
// : video 存储目录：/oss-dousheng/video/:userId/:videoName  (.mp4)
func GetOssVideoUrlAndCoverUrl(userId int64, videoName string) (string, string) {
	url := "http://123.57.251.188:22441/oss-dousheng/video/" + strconv.FormatInt(userId, 10) + "/" + videoName
	// playUrl coverUrl
	return url, url + "?x-oss-process=video/snapshot,t_2000,f_jpg,w_0,h_0"
}

// GetOssAvatarUrl
// : avatar 存储目录：/oss-dousheng/avatar/:userId  (.png)
func GetOssAvatarUrl(userId int64) string {
	url := "http://123.57.251.188:22441/oss-dousheng/avatar/" + strconv.FormatInt(userId, 10) + ".png"
	return url
}

// GetOssBgImageUrl
// : background_image 存储目录：/oss-dousheng/background_image/:userId  (.png)
func GetOssBgImageUrl(userId int64) string {
	url := "http://123.57.251.188:22441/oss-dousheng/background_image/" + strconv.FormatInt(userId, 10) + ".png"
	return url
}

// https://oss-dousheng.oss-cn-beijing.aliyuncs.com/background_image/7bb22a2b0f7b4a248410df047ac87b70.png
// https://oss-dousheng.oss-cn-beijing.aliyuncs.com/avator/9419939435a94c9bbcc0b0b96132c705.png
// https://oss-dousheng.oss-cn-beijing.aliyuncs.com/video/6732519909/1596432356843388.mp4
// https://oss-dousheng.oss-cn-beijing.aliyuncs.com/video/6732519909/1596432356843388.mp4?x-oss-process=video/snapshot,t_000,f_jpg,w_800,h_600
