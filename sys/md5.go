package sys

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// 计算文件的Md5值
func FileMd5(file string) string {
	f, err := os.Open(file)
	if err != nil {
		return ""
	}

	h := md5.New()
	io.Copy(h, f)
	return hex.EncodeToString(h.Sum(nil))
}

// 计算字符串的Md5值
func StrMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
