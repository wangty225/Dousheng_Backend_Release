package utils

import (
	"fmt"
	"path/filepath"
	"testing"
	"time"
)

func TestGetCallerAbsPath(t *testing.T) {
	fmt.Println(filepath.ToSlash(GetCallerAbsPath() + "/../../../config/yml/etcd.yml"))
}

func TestCode(t *testing.T) {
	fmt.Printf("%v\n", time.Now().UnixMilli())
	//fmt.Printf("%v\n", time.)
}
