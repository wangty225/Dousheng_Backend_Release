package mysql

import (
	"fmt"
	"testing"
)

func QuesyOne() {
	var u UserDao
	ret := DBMysql.First(&u, "username=?", "wangty")
	if ret.Error != nil {
		fmt.Println("Error querying database:", ret.Error)
		return
	}
	if ret.RowsAffected == 0 {
		fmt.Println("No matching records found.")
		return
	}
	fmt.Printf("%v\n", u.ID)
	fmt.Printf("%s\n", u.BackgroundImage)
}

func TestQuesyOne(t *testing.T) {
	QuesyOne()
}

func TestGetVideoListAfterTime(t *testing.T) {
	var a int64 = 1692875255
	GetVideoListAfterTime(&a)
}
