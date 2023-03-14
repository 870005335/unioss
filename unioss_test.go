package unioss

import (
	"fmt"
	"testing"
)

func TestOSS(t *testing.T) {
	err := NewStorage(AliYunConst, Config{
		KeyId:     "LTAI5t6kPgz1UzuaES4Vn19J",
		KeySecret: "UdI8MFqQdBD13OyStAvf5ZfoV6mRUZ",
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
