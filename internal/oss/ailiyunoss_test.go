package oss

import (
	"fmt"
	"testing"
)

func TestUploadVideoToOss(t *testing.T) {

}

func TestGetOssVideoUrlAndImgUrl(t *testing.T) {
	url, url2 := GetOssVideoUrlAndCoverUrl(6732519909, "1596432356843388.mp4") // video加user_id加video的playurl字段值
	fmt.Printf("%v\n", url)
	fmt.Printf("%v\n", url2)
}
