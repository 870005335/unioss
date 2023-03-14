package unioss

import (
	"fmt"
	"testing"
)

func TestOSS(t *testing.T) {
	err := NewStorage(AliYunConst, Config{
		KeyId:     "",
		KeySecret: "",
		Endpoint:  "https://oss-cn-hangzhou.aliyuncs.com",
		Bucket:    "oss-go-maktub",
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s, err := GetStorage()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(s.IsExists("test.jpg"))
}
