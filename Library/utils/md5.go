package utils

import (
	"crypto/md5"
	"strings"
)

//CreatMdStr 对密码进行md5加密
func CreatMdStr(str [md5.Size]byte) string {
	res := ""
	for _, c := range str {
		res += string(c)
	}
	return strings.ToLower(string(res))
}
