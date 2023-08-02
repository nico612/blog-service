package util

import (
	"crypto/md5"
	"encoding/hex"
)

// EncodeMD5 上传后的文件名格式化，将文件名MD5后在进行写入，防止直接把原始名称就暴露出去
func EncodeMD5(value string) string {

	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
