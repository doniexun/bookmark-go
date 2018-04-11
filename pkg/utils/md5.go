package utils

import (
	"fmt"
	"crypto/md5"
)

func Md5(str string) string {
	data := []byte(str)
	hash := md5.Sum(data)
	md5str := fmt.Sprintf("%x", hash) //将[]byte转成16进制
	return md5str
}
