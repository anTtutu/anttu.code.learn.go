package tool

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// 带固定盐的 md5，非标准 md5值
func str2md5Salt(plantStr string, salt string) string {
	saltData := []byte(salt)
	strData := []byte(plantStr)

	has := md5.New()
	has.Write(strData)
	has.Write(saltData)

	value := has.Sum(nil)
	return hex.EncodeToString(value)
}

// 不带盐，标准的 md5 值
func str2md5Normal(plantStr string) string {
	data := []byte(plantStr)
	has := md5.Sum(data)
	// 将[]byte转成16进制
	value := fmt.Sprintf("%x", has)
	return value
}
